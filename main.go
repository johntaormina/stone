package main

import (
	"stone/cmd"
)

func main() {
	cmd.Execute()
}

// OPENING VIM FILE

// cmd := exec.Command("vim", "test.txt")
// cmd.Stdin = os.Stdin
// cmd.Stdout = os.Stdout
// err := cmd.Run()
// fmt.Println(err)
