// package gen contains generated code, because unrolling loops is boring.
// See cmd/gen for the actual generation.
package gen

//go:generate go run ../cmd/gen -package-name=gen -unrolls=16 -output=./unrolled.gen.go
