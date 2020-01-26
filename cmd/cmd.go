package cmd

import (
	"math/rand"
	"time"
)

// Main runs GoJob interpreting flags and commands out of os.Args
func Main() {
	rand.Seed(time.Now().Unix())
}
