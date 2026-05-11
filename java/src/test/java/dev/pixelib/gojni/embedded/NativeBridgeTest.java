package dev.pixelib.gojni.embedded;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class NativeBridgeTest {
    @Test
    void pingReturnsExpectedValue() {
        assertEquals(true, EmbeddedApplication.ping());
    }

    @Test
    void createExampleObjectReturnsExpectedValues() {
        var obj = EmbeddedApplication.createExampleObject();
        assertEquals("Hello from Go!", obj.getMessage());
        assertEquals(42, obj.getCount());
        // We can't assert the exact timestamp, but we can check it's a reasonable value (e.g., within the last 5 minutes)
        long now = System.currentTimeMillis();
        long fiveMinutesAgo = now - 5 * 60 * 1000;
        assert (obj.getTimestamp() >= fiveMinutesAgo && obj.getTimestamp() <= now);
    }
}
