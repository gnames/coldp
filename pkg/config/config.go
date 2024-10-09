package config

import (
	"os"
	"path/filepath"

	"github.com/gnames/gnfmt"
)

// Config is a configuration object for the Darwin Core Archive (DwCA)
// data processing.
type Config struct {
	// CacheDir is the root path for all temporary files.
	CacheDir string

	// DownloadDir is used to store downloaded files.
	DownloadDir string

	// ExtractDir is used to store extracted files of ColDP archive.
	ExtractDir string

	// WithSloppyCSV allows to have more fields in a row, than it should have.
	WrongFieldsNum gnfmt.BadRow

	// Delimiter (usually tab \t) delimites between fields.
	Delimiter string

	// Quote (usually empty) used if a field has a delimiter in its content.
	Quote string
}

type Option func(*Config)

// New creates a new Config object with default values, and allows to
// override them with options.
func New(opts ...Option) Config {
	path, err := os.UserCacheDir()
	if err != nil {
		path = os.TempDir()
	}

	path = filepath.Join(path, "coldp_go")
	c := Config{
		CacheDir:       path,
		WrongFieldsNum: gnfmt.ErrorBadRow,
		Delimiter:      "\t",
		Quote:          "",
	}

	for _, opt := range opts {
		opt(&c)
	}

	c.DownloadDir = filepath.Join(c.CacheDir, "download")
	c.ExtractDir = filepath.Join(c.CacheDir, "extract")
	return c
}
