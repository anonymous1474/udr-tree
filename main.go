package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/anonymous1474/udr-tree/tree"
)

var num, id, rounds, rate, siz int

func init() {
	flag.IntVar(&num, "n", 3, "total number of replicas (default=3)")
	flag.IntVar(&id, "p", 0, "must put id")
	flag.IntVar(&rounds, "r", 100, "number of rounds")
	flag.IntVar(&rate, "o", 2, "ops per sec")
	flag.IntVar(&siz, "s", 100, "tree size")
	flag.Parse()

	log.SetFlags(0) // Turn off timestamps in log output.
	rand.Seed(time.Now().UnixNano())
}

func main() {
	tree.SetupReplica(num, id, rounds, rate, siz)
}
