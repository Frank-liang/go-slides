package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-ltah", "/abcd")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
}
