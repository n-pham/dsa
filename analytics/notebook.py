import marimo

__generated_with = "0.11.24"
app = marimo.App(width="medium")


@app.cell
def _():
    import dlt
    from dlt.destinations import duckdb as duckdb_dest
    import duckdb
    from dotenv import load_dotenv
    import os
    import requests

    # Define a function to fetch commits from GitHub
    @dlt.resource
    def fetch_commits(repo: str, token: str):
        url = f"https://api.github.com/repos/{repo}/commits"
        headers = {"Authorization": f"token {token}"}

        while url:
            response = requests.get(url, headers=headers)
            response.raise_for_status()
            commits = response.json()
            yield commits  # Yield the current page of commits

            # Get the next page URL from the response headers
            url = response.links.get("next", {}).get("url")


    # Define the pipeline
    pipeline = dlt.pipeline(
        pipeline_name="github_commits",
        destination=duckdb_dest(credentials=duckdb.connect(":memory:"))
    )
    return (
        dlt,
        duckdb,
        duckdb_dest,
        fetch_commits,
        load_dotenv,
        os,
        pipeline,
        requests,
    )


@app.cell
def _(load_dotenv, os):
    load_dotenv('secrets.env', override=True)
    GITHUB_TOKEN = os.getenv('GITHUB_TOKEN')
    REPO = "n-pham/dsa"
    return GITHUB_TOKEN, REPO


@app.cell
def _(GITHUB_TOKEN, REPO, fetch_commits, pipeline):
    # Load all commits
    info = pipeline.run(fetch_commits(REPO, GITHUB_TOKEN))

    # Print pipeline information
    print(info)
    return (info,)


@app.cell
def _(pipeline):
    dataset = pipeline.dataset()
    return (dataset,)


@app.cell
def _(dataset):
    sql_code = dataset("""
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
            strftime(CAST(commit__committer__date AS TIMESTAMP) + INTERVAL '7 hours', '%H')::INTEGER AS hour
        FROM fetch_commits
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
    """)
    print(sql_code.df().head())
    return (sql_code,)


@app.cell
def _(dataset):
    sql_code_2 = dataset("""
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
            strftime(CAST(commit__committer__date AS TIMESTAMP) + INTERVAL '7 hours', '%w')::INTEGER AS day_of_week
        FROM fetch_commits
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
    """)
    print(sql_code_2.df().head())
    return (sql_code_2,)


@app.cell
def _():
    return


if __name__ == "__main__":
    app.run()
