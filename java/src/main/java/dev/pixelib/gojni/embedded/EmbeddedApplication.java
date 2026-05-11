package dev.pixelib.gojni.embedded;

import dev.pixelib.gojni.example.HelloWorldContainer;

public class EmbeddedApplication {

    static {
        NativeLoader.load();
    }

    public static native boolean ping();
    public static native HelloWorldContainer createExampleObject();
}
