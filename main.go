package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/jasontconnell/crudgeon/configuration"
	"github.com/jasontconnell/crudgeon/process"

	"path/filepath"
)

func main() {
	configFile := flag.String("config", "config.json", "config file")
	file := flag.String("file", "", "source file")
	path := flag.String("path", "", "output location")
	obj := flag.String("obj", "", "object name")
	ns := flag.String("ns", "", "namespace")
	fld := flag.Bool("usefield", false, "use field name for property name")
	flag.Parse()

	n := time.Now()

	if *file == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	pfile, err := process.ParseFile(*file)

	if *obj == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	cfg := configuration.LoadConfig(*configFile)

	tmplRoot, err := filepath.Abs(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	tmplRoot = filepath.Dir(tmplRoot)

	for _, g := range cfg.Generations {
		tmpfile := g.File
		if !filepath.IsAbs(tmpfile) {
			tmpfile = filepath.Join(tmplRoot, g.File)
		}
		gp, err := process.GetGenPackage(*obj, *path, pfile.Fields, g.FileType, tmpfile, *ns, g.OutputPrefix, g.Folder, g.Flags, pfile.GenFlags, *fld)

		if err != nil {
			log.Fatal(err)
		}
		err = process.Generate(gp, g.CreateObjDir)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Finished", *obj, time.Since(n))
}
