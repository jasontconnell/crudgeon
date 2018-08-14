package main

import (
	"flag"
	"log"
	"lpgagen/configuration"
	"lpgagen/data"
	"lpgagen/process"
	"os"
	"time"
)

func main() {
	file := flag.String("file", "", "source file")
	path := flag.String("path", "", "output location")
	obj := flag.String("obj", "", "object name")
	jsonout := flag.String("jsonout", "", "convert to json and save and quit")
	jsonfile := flag.String("json", "", "json file version")
	flag.Parse()

	n := time.Now()

	if *file == "" && *jsonfile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var err error
	var flds []data.Field
	if *file != "" {
		flds, err = process.ParseFields(*file)
	} else if *jsonfile != "" {
		flds, err = process.ParseJsonFields(*jsonfile)
	}

	if *jsonout != "" {
		process.GenerateJson(flds, *jsonout)
		os.Exit(0)
	}

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
		gp := process.GetGenPackage(*obj, *path, flds, g.FileType, g.File, g.OutputPrefix, g.Folder, g.ForInterface)
		err := process.Generate(gp)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Finished", *obj, time.Since(n))
}
