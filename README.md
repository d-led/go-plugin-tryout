# go-plugin-tryout

[![Build Status](https://travis-ci.org/d-led/go-plugin-tryout.svg?branch=main)](https://travis-ci.org/d-led/go-plugin-tryout)

trying out plugin mechanisms:

- https://github.com/hashicorp/go-plugin
- https://golang.org/pkg/plugin

on windows, run:

```shell
demo.bat
```

on *X:

```shell
./demo.sh
```

## Structure

- [interfaces/randomizer.go](interfaces/randomizer.go) - pure Go interface of the plugin functionality
- [randomizer](randomizer) - plugin-independent interface implementation
- [interfaces/randomizer_rpc.go](interfaces/randomizer_rpc.go) - go-plugin RPC wrapper
- [randomizer_rpc](randomizer_rpc) - go-plugin plugin main
- [randomizer_native](randomizer_native) - `golang.org/pkg/plugin` plugin entry point
- [main.go](main.go) - demo
