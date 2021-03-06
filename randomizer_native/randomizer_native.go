package main

import (
	"github.com/d-led/go-plugin-tryout/interfaces"
	"github.com/d-led/go-plugin-tryout/randomizer"
)

// InjectRand is the plugin entry point
func InjectRand(i *interfaces.Randomizer) {
	randomizer.InjectRand(i)
}

// main for the main package (not needed for the plugin itself)
func main() {}
