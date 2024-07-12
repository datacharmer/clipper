package main

import (
	"os"
	"testing"

	"github.com/atotto/clipboard"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestClipper(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir:                 "testdata",
		RequireExplicitExec: true,
		Cmds:                customCommands(),
	})
}

func TestMain(m *testing.M) {

	exitCode := testscript.RunMain(m, map[string]func() int{
		"clipper": RunMain,
	})

	os.Exit(exitCode)
}

func dumpClipboard(ts *testscript.TestScript, neg bool, args []string) {
	text, err := clipboard.ReadAll()
	ts.Check(err)
	ts.Stdout().Write([]byte(text))
}

func clearClipboard(ts *testscript.TestScript, neg bool, args []string) {
	err := clipboard.WriteAll("")
	ts.Check(err)
}

func customCommands() map[string]func(ts *testscript.TestScript, neg bool, args []string) {
	return map[string]func(ts *testscript.TestScript, neg bool, args []string){
		"dump_clipboard":  dumpClipboard,
		"clear_clipboard": clearClipboard,
	}
}
