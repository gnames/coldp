package arcio

import (
	"github.com/gnames/coldp/config"
	"github.com/gnames/coldp/ent/coldp"
)

type arcio struct {
	// path to the CoLDP archive file
	path string

	// cfg is the configuration
	cfg config.Config

	// metaPath is the path to metafile.
	metaPath string

	// metaType can be YAML or JSON
	metaType coldp.FileType

	// meta contains metadata. It is nil if empty.
	meta *coldp.Meta

	// dataPaths contains file paths to dataPaths files.
	// key is low-case type of data, the value is the file path.
	dataPaths map[string]string

	// dataType explains what kind of archive is detected.
	// Can be Name or NameUsage. If neither is found
	// it will generate error during harvest process.
	dataType coldp.ArchiveType
}

func New(cfg config.Config, path string) coldp.Archive {
	res := arcio{
		path:      path,
		cfg:       cfg,
		dataPaths: make(map[string]string),
	}

	return &res
}

// DataPaths returns map of low-case names of data files without extensions
// and the path to these files.
func (a *arcio) DataPaths() map[string]string {
	return a.dataPaths
}
