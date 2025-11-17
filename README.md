# go-plugin-tryout

trying out plugin mechanisms:

- https://github.com/hashicorp/go-plugin
- https://golang.org/pkg/plugin

on windows, run:

```shell
demo.bat
```

on \*X:

```shell
./demo.sh
```

## Plugin Approaches Comparison

### hashicorp/go-plugin (RPC-based)

**Pros:**

- ✅ **Cross-platform**: Works on all platforms (Windows, Linux, macOS)
- ✅ **Process isolation**: Plugins run in separate processes, providing better isolation and crash protection
- ✅ **Version flexibility**: Can work with different Go versions between host and plugin
- ✅ **Network support**: Can run plugins over network
- ✅ **Better error handling**: Process crashes don't affect the main application

**Cons:**

- ❌ **Performance overhead**: RPC communication adds latency compared to in-process calls
- ❌ **Resource usage**: Each plugin runs as a separate process
- ❌ **Complexity**: Requires more boilerplate code (RPC wrappers, handshake configs)
- ❌ **Startup time**: Process spawning adds initialization overhead

### golang.org/pkg/plugin (Native)

**Pros:**

- ✅ **Performance**: Direct in-process calls, no RPC overhead
- ✅ **Low latency**: Function calls are as fast as regular Go function calls
- ✅ **Resource efficiency**: Shared memory space, no separate processes
- ✅ **Simplicity**: Simpler API, less boilerplate code

**Cons:**

- ❌ **Platform limitations**: **NOT supported on Windows/amd64** (see [demo.bat](demo.bat) line 15)
- ❌ **Go version lock**: Plugin and host must be compiled with the same Go version
- ❌ **No isolation**: Plugin crashes can bring down the main application
- ❌ **Build constraints**: Requires `-buildmode=plugin` flag and platform-specific file extensions (.so, .dylib)

### Platform Support Summary

| Platform | hashicorp/go-plugin | golang.org/pkg/plugin |
| -------- | ------------------- | --------------------- |
| Linux    | ✅                  | ✅                    |
| macOS    | ✅                  | ✅                    |
| Windows  | ✅                  | ❌ (not supported)    |

## Structure

- [interfaces/randomizer.go](interfaces/randomizer.go) - pure Go interface of the plugin functionality
- [randomizer](randomizer) - plugin-independent interface implementation
- [interfaces/randomizer_rpc.go](interfaces/randomizer_rpc.go) - go-plugin RPC wrapper
- [randomizer_rpc](randomizer_rpc) - go-plugin plugin main
- [randomizer_native](randomizer_native) - `golang.org/pkg/plugin` plugin entry point
- [main.go](main.go) - demo
