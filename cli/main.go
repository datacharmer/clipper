package main

import (
	"fmt"
	"os"

	"github.com/supitsdu/clipper/cli/clipper"
	"github.com/supitsdu/clipper/cli/options"
	"github.com/supitsdu/clipper/cli/reader"
)

func RunMain() int {
	config := options.ParseFlags() // Parse command-line flags

	if config.ShowVersion {
		fmt.Printf("Clipper %s\n", options.GetVersion())
		return 0
	}

	writer := clipper.DefaultClipboardWriter{} // Create the default clipboard writer

	reader := reader.ContentReader{Config: config}

	msg, err := clipper.Run(reader, writer) // Run the main Clipper logic
	if err != nil {
		fmt.Printf("Error %s\n", err)
		return 1
	}

	fmt.Println(msg)
	return 0
}

func main() {
	os.Exit(RunMain())
}
