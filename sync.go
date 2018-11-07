package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("%s\n", path)
	return nil
}


func main() {
	root :=  os.Args[1]
	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Printf("Error when walking through the fs: %s\n", err)
	}
}