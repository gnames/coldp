package coldp

import (
	"github.com/gnames/dwca/pkg/ent/meta"
	"honnef.co/go/tools/config"
)

type Archive interface {
	// Config returns the configuration object of the archive.
	Config() config.Config

	// Meta returns the Meta object of the archive.
	Meta() *meta.Meta

	// Load extracts the archive and loads data for Meta.
	Load() error

	// Close cleans up temporary files.
	Close() error
}
