package coldp

import "sync"

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
	DataPaths() map[string]string

	// LoadMeta returns coldp.Meta struct. If the struct is empty it populates
	// it with data from meta file.
	Meta() (*Meta, error)
}

type Reader[T Data] interface {
	Read(chan<- T, sync.WaitGroup) error
}

type Data interface {
	Load(header, row []string) error
}
