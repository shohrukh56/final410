import matplotlib.pyplot as plt

iterations = 1000000

max_concurrent_threads = [2, 5, 20, 100]

write_durations = [5906, 8314, 6060, 5963]  # in milliseconds
read_durations = [5636, 6856, 5658, 5562]  # in milliseconds
cpu_usage = [2.06, 2.36, 2.90, 2.91]

# Plot CPU Usage and Write/Read Durations in the same plot
plt.figure(figsize=(12, 6))

# Connect the dots with lines for WriteFile durations
write_line, = plt.plot(max_concurrent_threads, write_durations, marker='o', label='WriteFile Duration', color='blue', linestyle='-', markersize=8)

# Connect the dots with lines for ReadFile durations
read_line, = plt.plot(max_concurrent_threads, read_durations, marker='o', label='ReadFile Duration', color='orange', linestyle='-', markersize=8)

# Add CPU usage and WriteFile duration annotations near each data point
for i, txt in enumerate(cpu_usage):
    plt.annotate(f'CPU: {txt:.2f}%\n {write_durations[i]:.2f} ms',
                 (max_concurrent_threads[i], write_durations[i]), textcoords="offset points", xytext=(0,5), ha='center')

# Add CPU usage and WriteFile duration annotations near each data point
for i, txt in enumerate(cpu_usage):
    plt.annotate(f'CPU: {txt:.2f}%\n {read_durations[i]:.2f} ms',
                 (max_concurrent_threads[i], read_durations[i]), textcoords="offset points", xytext=(0,5), ha='center')

plt.xlabel('Max Concurrent Threads')
plt.ylabel('Duration (ms)')
plt.title('Performance Metrics for Different Concurrent Threads (per 1000000 iterations) using Java and thread')
plt.legend(handles=[write_line, read_line], loc='upper left')
plt.grid(True)
plt.show()
