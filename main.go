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
	dir := flag.String("dir", "", "process all files in a directory. they must have the +class flag in the file, or it'll fail")
	flag.Parse()

	n := time.Now()

	if *file == "" && *dir == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	cfg := configuration.LoadConfig(*configFile)
	tmplRoot, err := filepath.Abs(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	tmplRoot = filepath.Dir(tmplRoot)

	pfiles := []process.ParsedFile{}

	if *dir == "" {
		pfile, err := process.ParseFile(*file)
		if err != nil {
			log.Fatal(err)
		}
		pfiles = append(pfiles, pfile)
	} else {
		paths, err := getFiles(*dir)
		if err != nil {
			log.Fatal(err)
		}

		for _, p := range paths {
			pfile, err := process.ParseFile(p)
			if err != nil {
				log.Fatal(err)
			}
			pfiles = append(pfiles, pfile)
		}
	}

	if len(pfiles) == 0 {
		log.Fatal("no files")
	}

	for _, pfile := range pfiles {
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
	}

	log.Println("Finished", *obj, time.Since(n))
}

func getFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	paths := []string{}
	for _, f := range entries {
		if filepath.Ext(f.Name()) != ".txt" {
			continue
		}
		paths = append(paths, filepath.Join(dir, f.Name()))
	}
	return paths, nil
}
