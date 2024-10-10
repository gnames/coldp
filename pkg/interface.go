package coldp

import (
	"github.com/gnames/coldp/ent/meta"
	"github.com/gnames/coldp/pkg/config"
)

type Archive interface {
	// Config returns the configuration object of the archive.
	Config() config.Config

	// Meta returns the Meta object of the archive.
	Meta() *meta.Data

	// Load extracts the archive and loads data for Meta.
	Load() error

	// Close cleans up temporary files.
	Close() error
}
