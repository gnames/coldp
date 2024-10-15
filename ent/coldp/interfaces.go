package coldp

import (
	"github.com/gnames/coldp/config"
)

type Archive interface {
	// ResetCache clears cache directory from files.
	ResetCache() error

	// Extracts zip archive to cache.
	Extract() error

	// DirInfo finds where data and metadata files are located.
	// It also determines if metadata is provided in JSON or YAML
	// format.
	DirInfo() error

	// DataPaths returns map where low-case name of data type is a key and
	// path to corresponding file is the value.
	DataPaths() map[DataType]string

	// Config returns configuration settings of archive.
	Config() config.Config

	// Meta returns coldp.Meta struct. If the struct is empty it populates
	// it with data from meta file first.
	Meta() (*Meta, error)
}

type DataLoader interface {
	Load(header, row []string) (DataLoader, error)
}
