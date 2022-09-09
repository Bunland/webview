import { dlopen, suffix } from "bun:ffi";

const path = `./webview.${suffix}`;

const { symbols } = dlopen(path, {
    CreateWindow: {
        // args: []
    }
})

symbols.CreateWindow()