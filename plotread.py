import matplotlib.pyplot as plt

# Results from the benchmark
num_concurrent_calls = [5, 10, 20]
total_iterations = [888324, 487538, 248364]
total_times = [1.084313791, 1.185722708, 1.214991875]
average_time_per_iteration = [time / iterations * 1000 for time, iterations in zip(total_times, total_iterations)]

# Plotting the results
plt.figure(figsize=(10, 6))
plt.plot(total_iterations, average_time_per_iteration, marker='o', linestyle='-', color='blue')

# Annotate each point with the corresponding info
for i, txt in enumerate(num_concurrent_calls):
    plt.annotate(f'Calls={txt}\nAverage Time per Iteration={average_time_per_iteration[i]:.12f} ms',
                 (total_iterations[i], average_time_per_iteration[i]), textcoords="offset points", xytext=(0, 10), ha='center')

plt.xlabel('Number of Iterations')
plt.ylabel('Average Time per Iteration (ms)')
plt.title('Average Time per Iteration for Different Number of Concurrent Calls')
plt.grid(True)
plt.show()
