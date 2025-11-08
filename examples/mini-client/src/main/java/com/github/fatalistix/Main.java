package com.github.fatalistix;

import java.net.InetSocketAddress;
import java.net.Socket;
import java.nio.channels.SelectionKey;
import java.nio.channels.Selector;
import java.nio.channels.ServerSocketChannel;
import java.time.LocalDateTime;
import java.util.Set;

public class Main {
    public static void main(String[] args) throws Exception {
        int targetPort = 8880;
        var host = "192.168.50.15";

        try (var socket = new Socket(host, targetPort);
             var in = socket.getInputStream();
             var out = socket.getOutputStream();
        ) {
            var msg = "1".repeat(10_000);
            var msgBytes = msg.getBytes();

            System.out.println(msgBytes.length);

            out.write(msgBytes);
            in.read();
            out.write(msgBytes);
            in.read();
            // etc
        }
    }

    public static void anotherMain() throws Exception {
        var selector = Selector.open();

        var channel = ServerSocketChannel.open();
        channel.configureBlocking(false);
        channel.bind(new InetSocketAddress(5000));

        channel.register(selector, SelectionKey.OP_ACCEPT);

        while (true) {
            int result = selector.select();
            if (result == 0) {
                continue;
            }

            selector.selectedKeys().forEach(key -> {
                // handle channel
            });
        }
    }
}