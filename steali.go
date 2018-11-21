package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	n := flag.Int("n", 1, "how many lines to steal")
	e := flag.Bool("e", false, "output lines to stderr instead of stdout")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [-e] [-n int] command [args]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	b := bufio.NewReader(os.Stdin)
	for i := 0; i < *n; i++ {
		line, err := b.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			// TODO: handle err
			panic(err)
		}
		if *e {
			os.Stderr.Write(line)
		} else {
			os.Stdout.Write(line)
		}
	}

	cmdline := flag.Args()
	if len(cmdline) == 0 {
		return
	}
	name := cmdline[0]
	args := cmdline[1:]
	cmd := exec.Command(name, args...)
	cmd.Stdin = b
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				os.Exit(status.ExitStatus())
			}
		}
		panic(err)
	}
}
