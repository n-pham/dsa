import requests
import json
from datetime import datetime
import time
import matplotlib.pyplot as plt
import pandas as pd
from collections import defaultdict
import numpy as np

def get_user_repositories(username):
    """
    Fetch all public repositories for a given GitHub username
    
    Args:
        username (str): GitHub username
    
    Returns:
        list: List of repository dictionaries with created_date and languages
    """
    repositories = []
    page = 1
    per_page = 100  # Maximum allowed by GitHub API
    
    while True:
        # GitHub API endpoint for user repositories
        url = f"https://api.github.com/users/{username}/repos"
        params = {
            'page': page,
            'per_page': per_page,
            'sort': 'created',
            'direction': 'desc'
        }
        
        try:
            response = requests.get(url, params=params)
            response.raise_for_status()
            
            repos = response.json()
            
            # If no more repositories, break the loop
            if not repos:
                break
            
            for repo in repos:
                # Skip forks
                if repo['fork']:
                    continue
                
                # Get languages for this repository
                languages = get_repository_languages(username, repo['name'])
                
                repo_info = {
                    'name': repo['name'],
                    'full_name': repo['full_name'],
                    'created_date': repo['created_at'],
                    'created_date_formatted': datetime.strptime(repo['created_at'], '%Y-%m-%dT%H:%M:%SZ').strftime('%Y-%m-%d %H:%M:%S'),
                    'languages': languages,
                    'description': repo['description'],
                    'url': repo['html_url'],
                    'is_fork': repo['fork'],
                    'stars': repo['stargazers_count'],
                    'forks': repo['forks_count']
                }
                
                repositories.append(repo_info)
                
                # Small delay to be respectful to GitHub's rate limits
                time.sleep(0.1)
            
            page += 1
            
        except requests.exceptions.RequestException as e:
            print(f"Error fetching repositories: {e}")
            break
    
    return repositories

def get_repository_languages(username, repo_name):
    """
    Get the languages used in a specific repository
    
    Args:
        username (str): GitHub username
        repo_name (str): Repository name
    
    Returns:
        dict: Dictionary with languages and their byte counts
    """
    url = f"https://api.github.com/repos/{username}/{repo_name}/languages"
    
    try:
        response = requests.get(url)
        response.raise_for_status()
        return response.json()
    except requests.exceptions.RequestException as e:
        print(f"Error fetching languages for {repo_name}: {e}")
        return {}

def create_language_growth_chart(repositories, top_n_languages=10, time_period='month'):
    """
    Create a stacked area chart showing language variety growth over time
    
    Args:
        repositories (list): List of repository dictionaries from get_user_repositories
        top_n_languages (int): Number of top languages to show individually (others grouped as 'Other')
        time_period (str): Time aggregation period ('month', 'quarter', 'year')
    
    Returns:
        matplotlib.figure.Figure: The created figure
    """
    if not repositories:
        print("No repositories data provided")
        return None
    
    # Filter out repositories without languages or creation date
    valid_repos = [repo for repo in repositories if repo.get('languages') and repo.get('created_date')]
    
    if not valid_repos:
        print("No repositories with language data found")
        return None
    
    # Convert to DataFrame for easier manipulation
    data = []
    for repo in valid_repos:
        created_date = datetime.strptime(repo['created_date'], '%Y-%m-%dT%H:%M:%SZ')
        
        # Get primary language (most used) for each repository
        if repo['languages']:
            primary_language = max(repo['languages'].items(), key=lambda x: x[1])[0]
        else:
            primary_language = 'Unknown'
        
        data.append({
            'date': created_date,
            'primary_language': primary_language,
            'repo_name': repo['name'],
            'all_languages': list(repo['languages'].keys()) if repo['languages'] else []
        })
    
    df = pd.DataFrame(data)
    df = df.sort_values('date')
    
    # Aggregate by time period
    if time_period == 'month':
        df['period'] = df['date'].dt.to_period('M')
    elif time_period == 'quarter':
        df['period'] = df['date'].dt.to_period('Q')
    elif time_period == 'year':
        df['period'] = df['date'].dt.to_period('Y')
    else:
        raise ValueError("time_period must be 'month', 'quarter', or 'year'")
    
    # Count cumulative unique languages over time
    language_counts = defaultdict(lambda: defaultdict(int))
    all_languages = set()
    
    # Track cumulative language usage
    for period in sorted(df['period'].unique()):
        period_repos = df[df['period'] <= period]
        
        # Count repositories using each language (cumulative)
        for _, repo in period_repos.iterrows():
            for lang in repo['all_languages']:
                all_languages.add(lang)
                language_counts[period][lang] += 1
    
    # Get top N languages by total usage
    total_usage = defaultdict(int)
    for period_data in language_counts.values():
        for lang, count in period_data.items():
            total_usage[lang] = max(total_usage[lang], count)
    
    top_languages = sorted(total_usage.items(), key=lambda x: x[1], reverse=True)[:top_n_languages]
    top_lang_names = [lang for lang, _ in top_languages]
    
    # Prepare data for stacked area chart
    periods = sorted(language_counts.keys())
    period_labels = [str(p) for p in periods]
    
    # Create matrix for stacked area chart
    data_matrix = []
    for lang in top_lang_names:
        lang_data = [language_counts[period].get(lang, 0) for period in periods]
        data_matrix.append(lang_data)
    
    # Calculate "Other" languages
    other_data = []
    for period in periods:
        total_other = sum(count for lang, count in language_counts[period].items() 
                         if lang not in top_lang_names)
        other_data.append(total_other)
    
    if any(count > 0 for count in other_data):
        data_matrix.append(other_data)
        labels = top_lang_names + ['Other']
    else:
        labels = top_lang_names
    
    # Create the stacked area chart
    fig, ax = plt.subplots(figsize=(12, 8))
    
    # Define colors for better visualization
    colors = plt.cm.Set3(np.linspace(0, 1, len(labels)))
    
    # Create stacked area chart
    ax.stackplot(range(len(periods)), *data_matrix, labels=labels, colors=colors, alpha=0.8)
    
    # Customize the chart
    ax.set_xlabel(f'Time Period ({time_period.capitalize()})', fontsize=12)
    ax.set_ylabel('Cumulative Number of Repositories', fontsize=12)
    ax.set_title('Language Variety Growth Over Time\n(Cumulative Repository Count by Language)', fontsize=14, fontweight='bold')
    
    # Set x-axis labels
    ax.set_xticks(range(len(periods)))
    ax.set_xticklabels(period_labels, rotation=45, ha='right')
    
    # Add legend
    ax.legend(bbox_to_anchor=(1.05, 1), loc='upper left')
    
    # Add grid for better readability
    ax.grid(True, alpha=0.3)
    
    # Adjust layout to prevent label cutoff
    plt.tight_layout()
    
    # Add some statistics
    total_repos = len(valid_repos)
    total_languages = len(all_languages)
    
    textstr = f'Total Repositories: {total_repos}\nTotal Languages: {total_languages}'
    props = dict(boxstyle='round', facecolor='wheat', alpha=0.8)
    ax.text(0.02, 0.98, textstr, transform=ax.transAxes, fontsize=10,
            verticalalignment='top', bbox=props)
    
    return fig

def main():
    repositories = get_user_repositories("TODO")
    # print(repositories)
    fig1 = create_language_growth_chart(repositories, top_n_languages=10, time_period='month')
    if fig1:
        plt.show()

if __name__ == "__main__":
    main()