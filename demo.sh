#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

echo "--== compiling main ==--"
rm -f ./randomizer_rpc_plugin ./randomizer_plugin.so
go build -o ./randomizer_demo .

echo "--== running the demo without the plugins built ==--"
./randomizer_demo || true

echo "--== compiling the rpc plugin ==--"
go build -o ./randomizer_rpc_plugin github.com/d-led/go-plugin-tryout/randomizer_rpc

echo "--== compiling the native plugin ==--"
go build -buildmode=plugin -o randomizer_plugin.so github.com/d-led/go-plugin-tryout/randomizer_native

echo "--== running the demo with the plugin built ==--"
./randomizer_demo
