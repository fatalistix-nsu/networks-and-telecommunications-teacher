package io.github.fatalistix;

import javax.swing.*;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

public class ThreadPoolsApp {

    @SuppressWarnings("all")
    public static void main(String[] args) throws Exception {
        var client = HttpClient.newHttpClient();
        var request = HttpRequest.newBuilder().build();
        var response = client.sendAsync(request, HttpResponse.BodyHandlers.ofString());

        response.whenComplete((resp, err) -> {
            System.out.println("Response is complete!");
        });



        try (
                var defaultPool = ThreadPoolConfig.newDefaultPool();
                var ioPool = ThreadPoolConfig.newIoPool();
        ) {
            defaultPool.execute(new AlgoTask());
            ioPool.execute(new WriteToFileTask());
            new JButton().addActionListener(e -> {});
            SwingUtilities.invokeLater(() -> {})
        }
    }
}
