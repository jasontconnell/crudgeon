package main

import (
	"flag"
	"log"
	"lpgagen/configuration"
	//"lpgagen/data"
	"lpgagen/process"
	"os"
	"time"
)

func main() {
	file := flag.String("file", "", "source file")
	path := flag.String("path", "", "output location")
	obj := flag.String("obj", "", "object name")
	flag.Parse()

	n := time.Now()

	if *file == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	flds, err := process.ParseFields(*file)

	if *obj == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	cfg := configuration.LoadConfig("config.json")

	for _, g := range cfg.Generations {
		gp, err := process.GetGenPackage(*obj, *path, flds, g.FileType, g.File, g.OutputPrefix, g.Folder, g.Instruction)

		if err != nil {
			log.Fatal(err)
		}
		err = process.Generate(gp)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Finished", *obj, time.Since(n))
}
