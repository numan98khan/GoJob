// Package atexit provides handling for functions you want called when
// the program exits unexpectedly due to a signal.
//
// You should also make sure you call Run in the normal exit path.
package atexit

import (
	"fmt"
)

// Run all the at exit functions if they haven't been run already
func Run() {
	fmt.Printf("Bout to Bounce...")
}
