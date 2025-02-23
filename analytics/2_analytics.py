import ibis
from datetime import timedelta, datetime
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
    'Late morning 9-12': (9, 12),
    'Early afternoon 12-15': (12, 15),
    'Late afternoon 15-18': (15, 18),
    'Evening 18-21': (18, 21),
    'Night 21-5': (21, 5)
}

# Calculate time_counts using SQL
time_counts_query = """
SELECT
    CASE
        WHEN hour >= 5 AND hour < 9 THEN 'Early morning 5-9'
        WHEN hour >= 10 AND hour < 12 THEN 'Late morning 10-12'
        WHEN hour >= 12 AND hour < 15 THEN 'Early afternoon 12-15'
        WHEN hour >= 15 AND hour < 18 THEN 'Late afternoon 15-18'
        WHEN hour >= 18 AND hour < 21 THEN 'Evening 18-21'
        ELSE 'Night 21-5'
    END AS time_range,
    COUNT(*) AS count
FROM (
    SELECT
        strftime(CAST(commit->>'committer'->>'date' AS TIMESTAMP) + INTERVAL '7 hours', '%H')::INTEGER AS hour
    FROM commits
)
GROUP BY time_range
ORDER BY
    CASE
        WHEN time_range = 'Early morning 5-9' THEN 1
        WHEN time_range = 'Late morning 10-12' THEN 2
        WHEN time_range = 'Early afternoon 12-15' THEN 3
        WHEN time_range = 'Late afternoon 15-18' THEN 4
        WHEN time_range = 'Evening 18-21' THEN 5
        ELSE 6
    END
"""
time_counts_df = con.raw_sql(time_counts_query).fetchdf()
time_counts = dict(zip(time_counts_df['time_range'], time_counts_df['count']))

# Calculate total commits
total_commits = sum(time_counts.values())

# Calculate time_percentages
time_percentages = {key: (value / total_commits) * 100 for key, value in time_counts.items()}

# Calculate day_counts using SQL
day_counts_query = """
SELECT
    CASE
        WHEN day_of_week = 0 THEN 'Sunday'
        WHEN day_of_week = 1 THEN 'Monday'
        WHEN day_of_week = 2 THEN 'Tuesday'
        WHEN day_of_week = 3 THEN 'Wednesday'
        WHEN day_of_week = 4 THEN 'Thursday'
        WHEN day_of_week = 5 THEN 'Friday'
        ELSE 'Saturday'
    END AS day_of_week,
    COUNT(*) AS count
FROM (
    SELECT
        strftime(CAST(commit->>'committer'->>'date' AS TIMESTAMP) + INTERVAL '7 hours', '%w')::INTEGER AS day_of_week
    FROM commits
)
GROUP BY day_of_week
ORDER BY
    CASE
        WHEN day_of_week = 1 THEN 1
        WHEN day_of_week = 2 THEN 2
        WHEN day_of_week = 3 THEN 3
        WHEN day_of_week = 4 THEN 4
        WHEN day_of_week = 5 THEN 5
        WHEN day_of_week = 6 THEN 6
        ELSE 7
    END
"""
day_counts_df = con.raw_sql(day_counts_query).fetchdf()
day_counts = [day_counts_df.loc[day_counts_df['day_of_week'] == day, 'count'].values[0] for day in ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday']]

# Calculate day_percentages
day_percentages = [(count / total_commits) * 100 for count in day_counts]

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
axs[1].bar(['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'], day_counts, color='lightgreen')
axs[1].set_title('Number of Commits per Day of the Week')
axs[1].set_ylabel('Number of Commits')

# Annotate bars with percentages
for i, value in enumerate(day_counts):
    axs[1].text(i, value + 0.5, f'{day_percentages[i]:.1f}%', ha='center')

# Calculate week_counts using SQL
week_counts_query = """
SELECT
    strftime(CAST(commit->>'committer'->>'date' AS TIMESTAMP) + INTERVAL '7 hours', '%Y-%W') AS week,
    COUNT(*) AS count
FROM commits
GROUP BY week
ORDER BY week
"""
week_counts_df = con.raw_sql(week_counts_query).fetchdf()
weeks = week_counts_df['week'].tolist()
week_counts = week_counts_df['count'].tolist()

# Bar Chart 3: Number of commits per week
fig, ax = plt.subplots(figsize=(12, 5))
ax.bar(weeks, week_counts, color='lightcoral')
ax.set_title('Number of Commits per Week')
ax.set_ylabel('Number of Commits')
ax.set_xlabel('Week')

# Annotate bars with counts
for i, value in enumerate(week_counts):
    ax.text(i, value + 0.5, f'{value}', ha='center', rotation=90)
# Convert week labels to begin and end dates

week_dates = []
for week in weeks:
    year, week_num = map(int, week.split('-'))
    start_date = datetime.strptime(f'{year}-W{week_num}-1', '%Y-W%W-%w')
    end_date = start_date + timedelta(days=6)
    week_dates.append(f'{start_date.strftime("%b %d")} - {end_date.strftime("%b %d")}')

# Update x-axis labels
ax.set_xticks(range(len(week_dates)))
ax.set_xticklabels(week_dates, rotation=45, ha='right')

# Set y-axis to integer values
ax.yaxis.set_major_locator(plt.MaxNLocator(integer=True))

plt.tight_layout()
plt.show()
