import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;
import java.lang.management.OperatingSystemMXBean;
import java.lang.management.ManagementFactory;

class FileSystem {

    private final Map<String, FileData> data; // Map to store file data
    private final Object mutex = new Object(); // Mutex for synchronization

    // Constructor to initialize the FileSystem with an empty data map
    public FileSystem() {
        this.data = new HashMap<>();
    }

    // Method to read file content
    public String readFile(String filename) {
        synchronized (mutex) {
            FileData file = data.get(filename);
            if (file == null) {
                throw new RuntimeException("File not found: " + filename);
            }

            // Verify checksum
            String calculatedChecksum = calculateChecksum(file.getContent());
            if (!calculatedChecksum.equals(file.getChecksum())) {
                throw new RuntimeException("Checksum verification failed for file: " + filename);
            }

            return file.getContent();
        }
    }

    // Method to write file content
    public void writeFile(String filename, String content) {
        synchronized (mutex) {
            // Generate checksum for the content
            String checksum = calculateChecksum(content);

            // Store content and checksum in the data map
            data.put(filename, new FileData(content, checksum));
        }
    }

    // Method to calculate MD5 checksum for a given content
    private String calculateChecksum(String content) {
        try {
            MessageDigest md = MessageDigest.getInstance("MD5");
            md.update(content.getBytes());
            byte[] digest = md.digest();
            StringBuilder hexString = new StringBuilder();
            for (byte b : digest) {
                hexString.append(String.format("%02x", b));
            }
            return hexString.toString();
        } catch (NoSuchAlgorithmException e) {
            throw new RuntimeException("Error calculating checksum", e);
        }
    }

    // Inner class representing file data (content and checksum)
    private static class FileData {
        private final String content;
        private final String checksum;

        public FileData(String content, String checksum) {
            this.content = content;
            this.checksum = checksum;
        }

        public String getContent() {
            return content;
        }

        public String getChecksum() {
            return checksum;
        }
    }

    // Main method to run benchmarks
    public static void main(String[] args) {
        int[] iterationsList = {1000000};
        int[] maxNumThreadsList = {2, 5, 20, 100};

        for (int iterations : iterationsList) {
            for (int maxNumThreads : maxNumThreadsList) {
                System.out.printf("iterations: %d\n", iterations);
                System.out.printf("max concurrent threads: %d\n", maxNumThreads);
                FileSystem fileSystem = new FileSystem();

                // Larger content for the file
                String content = "BenchmarkContentBenchmarkContentBenchmarkContentBenchmarkContentBenchmarkContent";

                // Benchmark writeFile operation
                long writeDuration = benchmarkWriteFile(fileSystem, content, iterations, maxNumThreads);
                System.out.printf("writeFile duration: %d\n", writeDuration);

                // Benchmark readFile operation
                long readDuration = benchmarkReadFile(fileSystem, iterations, maxNumThreads);
                System.out.printf("readFile duration: %d\n", readDuration);
                printCPUUsage();
            }
            System.out.println("-----------------");
        }
    }

    // Method to benchmark writeFile operation
    private static long benchmarkWriteFile(FileSystem fileSystem, String content, int iterations, int maxNumThreads) {
        long startTime = System.currentTimeMillis();

        ExecutorService executor = Executors.newFixedThreadPool(maxNumThreads);

        // Concurrent write operations
        for (int i = 0; i < iterations; i++) {
            final int fileIndex = i;
            executor.submit(() -> fileSystem.writeFile("file" + fileIndex + ".txt", content));
        }

        executor.shutdown();
        try {
            executor.awaitTermination(Long.MAX_VALUE, TimeUnit.NANOSECONDS);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        long endTime = System.currentTimeMillis();

        // Calculate the total duration for the benchmarkWriteFile function
        return endTime - startTime;
    }

    // Method to benchmark readFile operation
    private static long benchmarkReadFile(FileSystem fileSystem, int iterations, int maxNumThreads) {
        long startTime = System.currentTimeMillis();

        ExecutorService executor = Executors.newFixedThreadPool(maxNumThreads);

        // Prepare data for benchmark
        for (int i = 0; i < iterations; i++) {
            final int fileIndex = i;
            executor.submit(() -> {
                try {
                    fileSystem.readFile("file" + fileIndex + ".txt");
                } catch (Exception e) {
                    System.out.printf("Error reading file%d.txt: %s%n", fileIndex, e.getMessage());
                }
            });
        }

        executor.shutdown();
        try {
            executor.awaitTermination(Long.MAX_VALUE, TimeUnit.NANOSECONDS);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        long endTime = System.currentTimeMillis();

        // Calculate the total duration for the benchmarkReadFile function
        return endTime - startTime;
    }

    //CPU usage
    private static void printCPUUsage() {
        OperatingSystemMXBean osBean = ManagementFactory.getOperatingSystemMXBean();
        double cpuUsage = osBean.getSystemLoadAverage();
        System.out.printf("CPU Usage: %.2f%%\n\n", cpuUsage);
    }
}