package main

// By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.
import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Intn will return the same number. (To see a different number, seed the number generator; see rand.Seed
	fmt.Println("My Favorite num is", rand.Intn(10))
}
