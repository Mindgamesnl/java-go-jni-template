package dev.pixelib.gojni.embedded;

import java.io.IOException;
import java.io.InputStream;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.StandardCopyOption;
import java.util.Locale;
import java.util.concurrent.atomic.AtomicBoolean;

public class NativeLoader {
    private static final AtomicBoolean LOADED = new AtomicBoolean(false);

    private NativeLoader() {
    }

    static void load() {
        if (!LOADED.compareAndSet(false, true)) {
            return;
        }

        String platform = normalizePlatform(System.getProperty("os.name"));
        String arch = normalizeArch(System.getProperty("os.arch"));
        String libName = platformLibraryName(platform);
        String resourcePath = "/_native/" + platform + "/" + arch + "/" + libName;

        try (InputStream in = NativeLoader.class.getResourceAsStream(resourcePath)) {
            if (in == null) {
                throw new UnsatisfiedLinkError("Native library not found: " + resourcePath);
            }
            Path tempDir = Files.createTempDirectory("tscgo-native-");
            tempDir.toFile().deleteOnExit();
            Path libPath = tempDir.resolve(libName);
            Files.copy(in, libPath, StandardCopyOption.REPLACE_EXISTING);
            libPath.toFile().deleteOnExit();
            System.load(libPath.toAbsolutePath().toString());
        } catch (IOException e) {
            throw new UnsatisfiedLinkError("Failed to load native library: " + e.getMessage());
        }
    }

    private static String normalizePlatform(String rawOs) {
        String os = rawOs.toLowerCase(Locale.ROOT);
        if (os.contains("mac") || os.contains("darwin")) {
            return "macos";
        }
        if (os.contains("win")) {
            return "windows";
        }
        if (os.contains("nux") || os.contains("linux")) {
            return "linux";
        }
        return os;
    }

    private static String normalizeArch(String rawArch) {
        String arch = rawArch.toLowerCase(Locale.ROOT);
        if (arch.equals("x86_64") || arch.equals("amd64")) {
            return "amd64";
        }
        if (arch.equals("aarch64") || arch.equals("arm64")) {
            return "arm64";
        }
        return arch;
    }

    private static String platformLibraryName(String platform) {
        if (platform.equals("windows")) {
            return "native.exe";
        }
        return "native";
    }
}
