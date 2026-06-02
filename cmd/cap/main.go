// binary cap prints the capabilities of the archsimd library.
package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	caps := getCaps()
	for _, name := range slices.Sorted(maps.Keys(caps)) {
		capability := caps[name]
		fmt.Printf("%s: %v\n", name, capability)
	}
}
