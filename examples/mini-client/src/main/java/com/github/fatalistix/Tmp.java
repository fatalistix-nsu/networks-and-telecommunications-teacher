package com.github.fatalistix;

import java.time.LocalDateTime;

public record Tmp(
        String ip,
        LocalDateTime lastSeen
) {

    public Tmp withNewLastSeen(LocalDateTime newLastSeen) {
        return new Tmp(ip, newLastSeen);
    }
}
