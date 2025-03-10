import plotly.graph_objects as go

categories = ['STR', 'INT', 'DEX']

fig = go.Figure()

fig.add_trace(go.Scatterpolar(
    r=[4, 1, 1],
    theta=categories,
    fill='toself',
    name='Data(Ops) Engineer',
    marker=dict(color='rgba(255, 0, 0, 0.5)')
))
fig.add_trace(go.Scatterpolar(
    r=[1, 4, 1],
    theta=categories,
    fill='toself',
    name='Data Scientist',
    marker=dict(color='rgba(0, 0, 255, 0.5)')
))

fig.add_trace(go.Scatterpolar(
      r=[1, 1, 4],
      theta=categories,
      fill='toself',
      name='Data Analyst',
    marker=dict(color='rgba(0, 255, 0, 0.5)')
))

fig.add_trace(go.Scatterpolar(
      r=[3, 3, 1],
      theta=categories,
      fill='toself',
      name='AI/MLOps Engineer'
))

fig.add_trace(go.Scatterpolar(
      r=[3, 1, 3],
      theta=categories,
      fill='toself',
      name='Analytics Engineer'
))

fig.add_trace(go.Scatterpolar(
      r=[1, 3, 3],
      theta=categories,
      fill='toself',
      name='Decision Scientist'
))

fig.update_layout(
    polar=dict(
        radialaxis=dict(
            visible=True,
            range=[0, 4],
            tickmode='linear',
            tick0=0,
            dtick=1
        ),
        angularaxis=dict(
            direction='clockwise',
            rotation=90
        )
    ),
    showlegend=False,
    # paper_bgcolor='rgb(240, 240, 240)',  # Very light grey background
    # plot_bgcolor='rgb(240, 240, 240)'    # Very light grey plot background
)

fig.show()

# import matplotlib.pyplot as plt
# import numpy as np

# # Define the data roles and their corresponding RPG classes
# labels = [
#     "Data Engineer (Fighter)",
#     "Data Scientist (Wizard)",
#     "Data Analysts (Rogue)",
#     "AI/ML Engineer (Knight)",
#     "Analytics Engineer (Ranger)",
#     "Decision Scientist (Trickster)"
# ]

# # Define the attributes for each role
# attributes = ["STR", "INT", "DEX"]

# # Define the values for each role
# values = [
#     [3, 0, 0],  # Data Engineer (Fighter)
#     [0, 3, 0],  # Data Scientist (Wizard)
#     [0, 0, 3],  # Data Analysts (Rogue)
#     [2, 2, 0],  # AI/ML Engineer (Knight)
#     [2, 0, 2],  # Analytics Engineer (Ranger)
#     [0, 2, 2]   # Decision Scientist (Trickster)
# ]

# # Number of variables
# num_vars = len(attributes)

# # Compute angle for each attribute
# angles = np.linspace(0, 2 * np.pi, num_vars, endpoint=False).tolist()

# # The radar chart is a circle, so we need to "complete the loop"
# # and append the start value to the end.
# values = [v + [v[0]] for v in values]
# angles += angles[:1]

# # Plot the radar chart
# fig, ax = plt.subplots(figsize=(8, 8), subplot_kw=dict(polar=True))

# for i, value in enumerate(values):
#     ax.fill(angles, value, alpha=0.25)
#     ax.plot(angles, value, label=labels[i])

# # Add attribute labels
# ax.set_yticklabels([])
# ax.set_xticks(angles[:-1])
# ax.set_xticklabels(attributes)

# # Add a legend
# plt.legend(loc='upper right', bbox_to_anchor=(1.3, 1.1))

# # Title
# plt.title("Comparison of Data Roles with RPG Classes")

# # Show the plot
# plt.show()
