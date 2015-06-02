package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	if err := maine(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func maine(args []string) error {

	name := args[0]
	if _, err := os.Stat(name); err != nil {
		return fmt.Errorf("Can't find target: %s", err)
	}
	base := name + ".BASE"
	if _, err := os.Stat(base); err != nil {
		return fmt.Errorf("Can't find base: %s", err)
	}
	this := name + ".THIS"
	if _, err := os.Stat(this); err != nil {
		return fmt.Errorf("Can't find this: %s", err)
	}
	other := name + ".OTHER"
	if _, err := os.Stat(other); err != nil {
		return fmt.Errorf("Can't find other: %s", err)
	}
	cmd := exec.Command("bcompare", this, other, base, name)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		if err == exec.ErrNotFound {
			return fmt.Errorf("Can't find becompare executable")
		}
		// otherwise, this is probably just an non-zero error return from
		// bcompare, which hopefully has already written to stderr, so don't
		// write our own error.
		os.Exit(1)
	}
	return nil
}

func usage() {
	fmt.Print(`
Launches Beyond Compare to do a 3-way merge.

usage:
	bsmart [filename]
`[1:])
	os.Exit(1)
}
