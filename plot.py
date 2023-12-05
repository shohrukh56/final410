import matplotlib.pyplot as plt

# File Read Benchmark results
read_iterations = 88412169
read_time_ms = 13.74

# File Write Benchmark results
write_iterations = 76567645
write_time_ms = 15.53

# Combined File Read and Write Operation Benchmark Plot
plt.figure(figsize=(10, 6))

# Plot File Read Benchmark
plt.plot(read_iterations, read_time_ms, marker='o', linestyle='-', color='blue', label='File Read')

# Plot File Write Benchmark
plt.plot(write_iterations, write_time_ms, marker='o', linestyle='-', color='green', label='File Write')

plt.title('File Read and Write Operation Benchmark')
plt.xlabel('Number of Iterations')
plt.ylabel('Time Taken (ms)')
plt.grid(True)
plt.legend()
plt.text(read_iterations, read_time_ms, f'({read_iterations}, {read_time_ms} ms)')
plt.text(write_iterations, write_time_ms, f'({write_iterations}, {write_time_ms} ms)')
plt.annotate('Start Read', xy=(read_iterations, read_time_ms), xytext=(read_iterations - 20000000, read_time_ms + 1),
             arrowprops=dict(facecolor='black', shrink=0.05))
plt.annotate('Start Write', xy=(write_iterations, write_time_ms), xytext=(write_iterations - 20000000, write_time_ms + 1),
             arrowprops=dict(facecolor='black', shrink=0.05))

plt.show()
