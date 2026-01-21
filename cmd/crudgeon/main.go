package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/jasontconnell/crudgeon/configuration"
	"github.com/jasontconnell/crudgeon/data"
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

	cfg := configuration.LoadConfig(*configFile)
	tmplRoot, err := filepath.Abs(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	tmplRoot = filepath.Dir(tmplRoot)

	basetypeMap := make(map[string]data.MappedType)
	for _, mt := range cfg.TypeMap {
		basetypeMap[mt.Name] = data.MappedType{CodeType: mt.CodeType, DbType: mt.DbType, CodeDefault: mt.CodeDefault, DbDefault: mt.DbDefault, Import: mt.Import}
	}

	pfiles := []process.ParsedFile{}

	if *dir == "" {
		pfile, err := process.ParseFile(*file, basetypeMap, cfg.Null, cfg.DbNull)
		if err != nil {
			log.Fatal("in file", *file, err)
		}
		pfiles = append(pfiles, pfile)
	} else {
		paths, err := getFiles(*dir)
		if err != nil {
			log.Fatal(err)
		}

		if len(paths) == 0 {
			log.Fatal("no .txt files in", *dir)
		}

		for _, p := range paths {
			pfile, err := process.ParseFile(p, basetypeMap, cfg.Null, cfg.DbNull)
			if err != nil {
				log.Fatal("parsing file", pfile.Path, err)
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

			skipFlag, ok := pfile.GenFlags.Skip[g.Alias]
			skip := pfile.GenFlags.HasSkip && ok && skipFlag
			if skip {
				continue
			}

			if *ns == "" {
				ns = &g.Namespace
			}

			gp, err := process.GetGenPackage(*obj, *path, pfile.Fields, pfile.Imports, g.Database, tmpfile, *ns, g.FilenameTemplate, g.Folder, g.Extension, g.Flags, cfg.CollectionTemplate, pfile.GenFlags, *fld, g.ConditionFlag)

			if err != nil {
				log.Fatal("getting gen package from file: ", pfile.Path, " error: ", err)
			}
			err = process.Generate(gp, g.CreateObjDir)
			if err != nil {
				log.Fatal("generating from file: ", pfile.Path, " error: ", err)
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
