package config

import (
	"os"
	"path/filepath"

	"github.com/gnames/gnfmt"
) // Config is a configuration object for the CoLDP archive data processing.
type Config struct {
	// CacheDir is the root path for all temporary files.
	CacheDir string

	// DownloadDir is used to store downloaded files.
	DownloadDir string

	// ExtractDir is used to store extracted files of CoLDP archive.
	ExtractDir string

	// BadRow sets how to process rows with wrong number of fields in CSV
	// files.
	BadRow gnfmt.BadRow

	// WithQuotes sets CSV reader to use `"` as quote. When it is true,
	// RFC-based CSV reader is used, even if delimiter is tab or pipe.
	WithQuotes bool
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

func OptBadRow(br gnfmt.BadRow) Option {
	return func(c *Config) {
		c.BadRow = br
	}
}

func OptWithQuotes(b bool) Option {
	return func(c *Config) {
		c.WithQuotes = b
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
