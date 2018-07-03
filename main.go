package main

import (
	"flag"
	"fmt"
	"lpgagen/process"
	"os"
)

func main() {
	file := flag.String("file", "", "source file")
	out := flag.String("out", "", "dest file")
	obj := flag.String("obj", "", "class and table name")
	flag.Parse()

	if *file == "" || *out == "" || *obj == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	process.Parse(*file)

	fmt.Println("Hi")
}
