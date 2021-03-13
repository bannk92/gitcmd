package main

import (
	"fmt"
	"os/exec"
	"log"
	"bytes"
)

func main() {
	cmd := exec.Command("git", "fetch","--all")
	cmd.Dir = "d:\\branchs\\convert"
	var stderr bytes.Buffer
	var out bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatalln(err.Error(), stderr.String())
	}
	fmt.Println(out.String())
	
	checkout := exec.Command("git", "checkout","main")
	checkout.Dir = "d:\\branchs\\convert"
	checkout.Stderr = &stderr
	checkout.Stdout = &out
	if err := checkout.Run();err != nil {
		log.Fatalln(err.Error(), stderr.String())
	}
	fmt.Println(out.String())
}
