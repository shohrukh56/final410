import matplotlib.pyplot as plt

iterations = 1000000

concurrent_goroutines = [2, 5, 20, 100]

write_durations = [1852.557375, 1281.431125, 1082.676166, 1014.744125]  # in milliseconds
read_durations = [809.304125, 669.536417, 571.269708, 621.515958]  # in milliseconds
cpu_usage = [24.91, 10.16, 22.57, 22.65]

# Plot CPU Usage and Write/Read Durations in the same plot
plt.figure(figsize=(12, 6))

# Connect the dots with lines for WriteFile durations
write_line, = plt.plot(concurrent_goroutines, write_durations, marker='o', label='WriteFile Duration', color='blue', linestyle='-', markersize=8)

# Connect the dots with lines for ReadFile durations
read_line, = plt.plot(concurrent_goroutines, read_durations, marker='o', label='ReadFile Duration', color='orange', linestyle='-', markersize=8)

# Add CPU usage and WriteFile duration annotations near each data point
for i, txt in enumerate(cpu_usage):
    plt.annotate(f'CPU: {txt:.2f}%\n {write_durations[i]:.2f} ms',
                 (concurrent_goroutines[i], write_durations[i]), textcoords="offset points", xytext=(0,5), ha='center')

# Add CPU usage and WriteFile duration annotations near each data point
for i, txt in enumerate(cpu_usage):
    plt.annotate(f'CPU: {txt:.2f}%\n {read_durations[i]:.2f} ms',
                 (concurrent_goroutines[i], read_durations[i]), textcoords="offset points", xytext=(0,5), ha='center')

plt.xlabel('Max Concurrent Goroutines')
plt.ylabel('Duration (ms)')
plt.title('Performance Metrics for Different Concurrent Goroutines per 1000000 iteration using GO and goroutines')
plt.legend(handles=[write_line, read_line], loc='upper left')
plt.grid(True)
plt.show()
