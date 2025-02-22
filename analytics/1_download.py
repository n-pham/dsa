import os
import requests
from dotenv import load_dotenv
import json

# Load GitHub token from secrets.env
load_dotenv('secrets.env')
GITHUB_TOKEN = os.getenv('GITHUB_TOKEN')

# GitHub API URL for commits
url = 'https://api.github.com/repos/n-pham/dsa/commits'

# Headers for authentication
headers = {
    'Authorization': f'token {GITHUB_TOKEN}'
}

# Function to get all commits with paging
def get_all_commits(url, headers):
    commits = []
    while url:
        response = requests.get(url, headers=headers)
        response.raise_for_status()
        commits.extend(response.json())
        # Get the next page URL from the response headers
        url = response.links.get('next', {}).get('url')
    return commits

# Download all commits
all_commits = get_all_commits(url, headers)
# Save commits to json file
with open('commits_20250222.json', 'w') as f:
    json.dump(all_commits, f, indent=4)
# Print the number of commits downloaded
print(f'Total commits downloaded: {len(all_commits)}')