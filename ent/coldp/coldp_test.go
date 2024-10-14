package coldp_test

import (
	"path/filepath"
	"testing"

	"github.com/gnames/coldp/config"
	"github.com/gnames/coldp/ent/coldp"
	"github.com/gnames/coldp/io/sysio"
	"github.com/stretchr/testify/assert"
)

func TestNameExtract(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		msg, path string
	}{
		{"name-yaml", "name/ptero-yaml.zip"},
		{"usg-yaml", "nameusage/ptero-yaml.zip"},
	}

	for _, v := range tests {
		path := filepath.Join("..", "..", "testdata", v.path)
		arc, err := Extract(path)
		assert.Nil(err)
		assert.NotNil(arc)
	}
}

func TestNoFile(t *testing.T) {
	assert := assert.New(t)
	arc, err := Extract("notfile")
	assert.Nil(arc)
	assert.IsType(&coldp.ErrorFileMissing{}, err)
}

func TestNoURL(t *testing.T) {
	assert := assert.New(t)
	arc, err := Extract("http://not-a-url")
	assert.Nil(arc)
	assert.NotNil(err)
}

func TestNotZip(t *testing.T) {
	assert := assert.New(t)
	path := filepath.Join("..", "..", "testdata", "notzip.zip")
	arc, err := Extract(path)
	assert.NotNil(err)
	assert.IsType(&coldp.ErrExtract{}, err)
	assert.Nil(arc)
}

func Extract(path string) (coldp.Sys, error) {
	cfg := config.New()
	c := sysio.New(cfg, path)
	err := c.ResetCache()
	if err != nil {
		return nil, err
	}
	err = c.Extract()
	if err != nil {
		return nil, err
	}

	return c, nil
}
