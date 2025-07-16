from asciiplot import asciiize

# Method 3: Custom ASCII bar chart
# ========================================
# Commit Counts by Month
# ----------------------------------------
#  2024-12 |██████████████████████████                          41
#  2025-03 |███████████████████████████████████                 56
#  2025-06 |██████████████████████████████████████████████████  78
# ----------------------------------------
#          |    0                                           78

# Your commit counts data
commit_counts = {'2024-12': 41, '2025-03': 56, '2025-06': 78}

# Extract months and counts
months = list(commit_counts.keys())
counts = list(commit_counts.values())

# Method 1: Using asciiplot's asciiize function
print("Method 1: Using asciiplot library")
print("=" * 40)
chart = asciiize(
    counts,
    height=15,
    title='Commit Counts by Month',
    x_axis_description='Month Index',
    y_axis_description='Commits'
)
print(chart)

# Add month labels below the chart
print("\nMonth Labels:")
for i, month in enumerate(months):
    print(f"{i}: {month}")

# Method 2: More customized version with colors (if colored package is available)
print("\n\nMethod 2: Enhanced version")
print("=" * 40)
try:
    from asciiplot import Color
    chart = asciiize(
        counts,
        height=15,
        title='Commit Counts by Month',
        x_axis_description='Month Index',
        y_axis_description='Commits',
        sequence_colors=[Color.BLUE],
        title_color=Color.RED,
        inter_points_margin=2
    )
    print(chart)
except ImportError:
    # Fallback if colored package is not available
    chart = asciiize(
        counts,
        height=15,
        title='Commit Counts by Month',
        x_axis_description='Month Index',
        y_axis_description='Commits',
        inter_points_margin=2
    )
    print(chart)

print("\nMonth Labels:")
for i, month in enumerate(months):
    print(f"{i}: {month}")

# Method 3: Custom bar chart representation
print("\n\nMethod 3: Custom ASCII bar chart")
print("=" * 40)
max_count = max(counts)
bar_width = 50  # Width of the chart

print("Commit Counts by Month")
print("-" * 40)
for month, count in commit_counts.items():
    bar_length = int((count / max_count) * bar_width)
    bar = '█' * bar_length
    print(f"{month:>8} |{bar:<{bar_width}} {count:>3}")

print("-" * 40)
print(f"{'':>8} |{'0':>5}{'':>{bar_width-10}}{max_count:>5}")

# Method 4: Vertical bar chart style
print("\n\nMethod 4: Vertical bar chart")
print("=" * 40)
max_count = max(counts)
height = 10

for level in range(height, 0, -1):
    line = f"{int(max_count * level / height):>3} |"
    for count in counts:
        if count >= (max_count * level / height):
            line += "██"
        else:
            line += "  "
    print(line)

print("    " + "-" * (len(counts) * 2 + 1))
print("    |" + "".join(f"{i:>2}" for i in range(len(counts))))
print("Month labels:", " | ".join(months))