package coldp

type Sys interface {
	// ResetCache clears cache directory from files.
	ResetCache() error

	// Extracts zip archive to cache.
	Extract() error

	// DirInfo finds where data and metadata files are located.
	// It also determines if metadata is provided in JSON or YAML
	// format.
	DirInfo() error

	// LoadMeta populates coldp.Meta object with data from
	// meta JSON or YAML file.
	Meta() (*Meta, error)
}

type Data interface {
	Load(header, row []string) error
}
