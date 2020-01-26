package cmd

import (
	"fmt"
	// "log"
	"math/rand"
	"os"
	// "os/exec"
	// "path"
	// "regexp"
	"runtime"
	// "runtime/pprof"
	// "strconv"
	// "strings"
	// "sync"
	"time"

	"github.com/pkg/errors"
	"github.com/numan98khan/GoJob/atexit/atexit"
)

// Globals
var (
	// Flags
	version         bool
	// Errors
	errorCommandNotFound    = errors.New("command not found")
	errorUncategorized      = errors.New("uncategorized error")
	errorNotEnoughArguments = errors.New("not enough arguments")
	errorTooManyArguments   = errors.New("too many arguments")
)

// ShowVersion prints the version to stdout
func ShowVersion() {
	fmt.Printf("GoJob %s\n", "Alpha")
	fmt.Printf("- os/arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("- go version: %s\n", runtime.Version())
}

func resolveExitCode(err error) {
	atexit.Run()
}

// Main runs GoJob interpreting flags and commands out of os.Args
func Main() {
	rand.Seed(time.Now().Unix())
	setupRootCommand(Root)
}
