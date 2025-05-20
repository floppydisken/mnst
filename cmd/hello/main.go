package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
)

var (
	seed   = flag.Int64("seed", 42, "random seed")
	epochs = flag.Int("epochs", 10, "training epochs")
	lr     = flag.Float64("lr", 0.01, "learning rate")
)

func main() {
	flag.Parse()

	rand.New(rand.NewSource(*seed))

	log.Printf("Hello, MNIST! seed=%d epochs=%d lr=%g\n", *seed, *epochs, *lr)

	// For now we just print; later this hooks into the training loop.
	fmt.Println("Scaffold OK ✔️")
}
