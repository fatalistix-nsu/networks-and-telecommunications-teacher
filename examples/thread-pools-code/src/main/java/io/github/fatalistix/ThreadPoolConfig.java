package io.github.fatalistix;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class ThreadPoolConfig {

    private ThreadPoolConfig() {
    }

    public static ExecutorService newDefaultPool() {
        return Executors.newFixedThreadPool(Runtime.getRuntime().availableProcessors());
    }

    public static ExecutorService newIoPool() {
        return Executors.newFixedThreadPool(Runtime.getRuntime().availableProcessors() * 3);
    }
}
