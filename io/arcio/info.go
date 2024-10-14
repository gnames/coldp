package arcio

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gnames/coldp/ent/coldp"
	"github.com/gnames/gnsys"
)

func (a *arcio) DirInfo() error {
	paths, err := a.getFiles()
	if err != nil {
		return err
	}
	// find directory where data files reside
	dataDir := getDataDir(paths)
	var metaOK bool

	for _, v := range paths {
		dir, file, ext := gnsys.SplitPath(v)
		metaOK = a.checkMeta(v, file, ext, metaOK)

		if dir != dataDir {
			continue
		}

		file = strings.ToLower(file)
		switch file {
		case "author", "namerelation", "taxon", "synonym", "nameusage",
			"taxonproperty", "taxonconceptrelation", "speciesinteraction",
			"speciesestimate", "reference", "typematerial", "distribution",
			"media", "vernacularname", "treatment":
			a.dataPaths[file] = v
		default:
		}
	}

	return nil
}

func typeMap() map[string]struct{} {
	types := []string{"author", "namerelation", "taxon", "synonym", "nameusage",
		"taxonproperty", "taxonconceptrelation", "speciesinteraction",
		"speciesestimate", "reference", "typematerial", "distribution",
		"media", "vernacularname", "treatment"}
	res := make(map[string]struct{})
	for _, v := range types {
		res[v] = struct{}{}
	}
	return res
}

// getDataDir returns dataDir where data files are residing.
func getDataDir(paths []string) string {
	types := typeMap()

	dirs := make(map[string]int)
	for _, v := range paths {
		dir, file, _ := gnsys.SplitPath(v)
		file = strings.ToLower(file)
		if _, ok := types[file]; ok {
			dirs[dir]++
		}
	}

	var res string
	var count int
	for k, v := range dirs {
		if v > count {
			count = v
			res = k
		}
	}
	return res
}

func (a *arcio) checkMeta(path, file, ext string, metaOK bool) bool {
	if metaOK {
		return true
	}

	file = strings.ToLower(file)
	ext = strings.ToLower(ext)

	if !strings.HasPrefix(file, "meta") {
		return false
	}

	a.metaPath = path
	if strings.HasPrefix(ext, ".json") {
		a.metaType = coldp.JSON
	}
	if ext == "yaml" {
		a.metaType = coldp.YAML
	}
	return true
}

func (a *arcio) getFiles() ([]string, error) {
	var files []string
	root := a.cfg.ExtractDir

	err := filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				files = append(files, path)
			}
			return nil
		})

	if err != nil {
		return nil, err

	}
	return files, nil
}