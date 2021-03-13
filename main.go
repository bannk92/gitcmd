package main

import (
	"fmt"
	"os/exec"
	"log"
	"bytes"
)

func main() {
	runCmd("git", "fetch","--all")
	runCmd("git", "checkout","main")
}


func runCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Dir = "d:\\branchs\\convert"
	var stderr bytes.Buffer
	var out bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatalln(err.Error(), stderr.String())
	}
	fmt.Println(out.String())
}