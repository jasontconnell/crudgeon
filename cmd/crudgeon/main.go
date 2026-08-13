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
	path := flag.String("path", "", "output location")
	ns := flag.String("ns", "", "namespace")
	obj := flag.String("obj", "", "object name")
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

	if len(pfiles) == 0 {
		log.Fatal("no files")
	}

	for _, pfile := range pfiles {
		for _, g := range cfg.Generations {
			if g.OneFile {
				continue
			}
			tmpfile := g.File
			if !filepath.IsAbs(tmpfile) {
				tmpfile = filepath.Join(tmplRoot, g.File)
			}

			skipFlag, ok := pfile.GenFlags.Skip[g.Alias]
			skip := pfile.GenFlags.HasSkip && ok && skipFlag
			if skip {
				continue
			}

			lns := g.Namespace
			if lns == "" {
				lns = *ns
			}

			gp, err := process.GetGenPackage(pfile.GenFlags.ClassName, *path, pfile, g.Database, tmpfile, lns, g.FilenameTemplate, g.Folder, g.Extension, g.Flags, cfg.CollectionTemplate, g.ConditionFlag)
			if err != nil {
				log.Fatal("getting gen package from file: ", pfile.Path, " error: ", err)
			}

			err = process.Generate(gp, g.CreateObjDir)

			if err != nil {
				log.Fatal("generating from file: ", pfile.Path, " error: ", err)
			}
		}
	}

	for _, g := range cfg.Generations {
		if !g.OneFile {
			continue
		}

		tmpfile := g.File
		if !filepath.IsAbs(tmpfile) {
			tmpfile = filepath.Join(tmplRoot, g.File)
		}

		lns := g.Namespace
		if lns == "" {
			lns = *ns
		}

		gp, err := process.GetAllGenPackage(*path, pfiles, g.Database, tmpfile, lns, g.FilenameTemplate, g.Folder, g.Extension, g.Flags, cfg.CollectionTemplate, g.ConditionFlag)
		if err != nil {
			log.Fatal("getting gen package from all files. error: ", err)
		}
		err = process.Generate(gp, false)

		if err != nil {
			log.Fatal("generating: error: ", err)
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
