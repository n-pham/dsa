from __future__ import unicode_literals
import ibis
from datetime import timedelta
import matplotlib.pyplot as plt

# Create an Ibis connection to DuckDB
con = ibis.duckdb.connect()

# Read the JSON file directly into a DuckDB table using Ibis
con.raw_sql("CREATE TABLE commits AS SELECT * FROM read_json_auto('commits_20250222.json')")
table = con.table('commits')

# Query the data using Ibis
result = table[['commit']].execute()

# Convert the result to a list of tuples
data = list(result.itertuples(index=False, name=None))

# Extract dates and messages
datetimes = [item[0]['committer']['date'] + timedelta(hours=7) for item in data]
messages = [item[0]['message'] for item in data]

# Define time ranges
time_ranges = {
    'Early morning 5-9': (5, 9),
    'Late morning 10-12' : (10, 12),
    'Early afternoon 12-15': (12, 15),
    'Late afternoon 15-18': (15, 18),
    'Evening 18-21': (18, 21),
    'Night 21-5': (21, 5)
}

# Initialize counts
time_counts = {key: 0 for key in time_ranges.keys()}

# Count commits in each time range
for dt in datetimes:
    hour = dt.hour
    for time_range, (start, end) in time_ranges.items():
        if start <= end:
            if start <= hour < end:
                time_counts[time_range] += 1
        else:  # For ranges that span midnight
            if hour >= start or hour < end:
                time_counts[time_range] += 1

# Plot the data as a pie chart
# plt.figure(figsize=(10, 6))
# plt.pie(time_counts.values(), labels=time_counts.keys(), autopct='%1.1f%%', startangle=140)
# plt.title('Commits by Time of Day')
# plt.axis('equal')  # Equal aspect ratio ensures that pie is drawn as a circle.
# plt.tight_layout()
# plt.show()

# Calculate percentages for time ranges
total_commits = sum(time_counts.values())
time_percentages = {key: (value / total_commits) * 100 for key, value in time_counts.items()}

# Count commits per day of the week
day_counts = [0] * 7
for dt in datetimes:
    day_counts[dt.weekday()] += 1

# Calculate percentages for days of the week
day_percentages = [(count / total_commits) * 100 for count in day_counts]
days_of_week = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday']

# Plot the data
fig, axs = plt.subplots(2, 1, figsize=(12, 10))

# Bar Chart 1: Number of commits per day part
axs[0].bar(time_counts.keys(), time_counts.values(), color='skyblue')
axs[0].set_title('Number of Commits per Day Part')
axs[0].set_ylabel('Number of Commits')

# Annotate bars with percentages
for i, (key, value) in enumerate(time_counts.items()):
    axs[0].text(i, value + 0.5, f'{time_percentages[key]:.1f}%', ha='center')

# Bar Chart 2: Number of commits per day of the week
axs[1].bar(days_of_week, day_counts, color='lightgreen')
axs[1].set_title('Number of Commits per Day of the Week')
axs[1].set_ylabel('Number of Commits')

# Annotate bars with percentages
for i, value in enumerate(day_counts):
    axs[1].text(i, value + 0.5, f'{day_percentages[i]:.1f}%', ha='center')

plt.tight_layout()
plt.show()