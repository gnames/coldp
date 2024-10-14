package config

import (
	"os"
	"path/filepath"
)

// Config is a configuration object for the CoLDP archive
// data processing.
type Config struct {
	// CacheDir is the root path for all temporary files.
	CacheDir string

	// DownloadDir is used to store downloaded files.
	DownloadDir string

	// ExtractDir is used to store extracted files of DwCA archive.
	ExtractDir string
}

// Option is a function type that allows to standardize how options to
// the configuration are organized.
type Option func(*Config)

// OptCacheDir sets the root path for all temporary files.
func OptCacheDir(s string) Option {
	return func(c *Config) {
		c.CacheDir = s
	}
}

func New(opts ...Option) Config {
	path, err := os.UserCacheDir()
	if err != nil {
		path = os.TempDir()
	}
	path = filepath.Join(path, "coldp_go")

	res := Config{
		CacheDir: path,
	}

	for _, opt := range opts {
		opt(&res)
	}

	res.DownloadDir = filepath.Join(res.CacheDir, "download")
	res.ExtractDir = filepath.Join(res.CacheDir, "extract")

	return res
}