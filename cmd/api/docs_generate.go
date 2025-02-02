package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func DocsGenerate() {
	fmt.Println("Running docs init.... ")

	generateAt := "/home/vikas/temp/go-lang"
	cmd := exec.Command("swag", "init", "-d", generateAt)
	cmd.Dir = "/home/vikas/temp/go-lang/cmd/api"

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting stdout pipe:", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "Error starting command:", err)
		return
	}

	var b bytes.Buffer
	io.Copy(&b, stdout)

	cmd.Wait()

	fmt.Println(b.String())
	fmt.Println("Docs generated....")
}
