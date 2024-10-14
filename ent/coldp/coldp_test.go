package coldp_test

import (
	"path/filepath"
	"testing"

	"github.com/gnames/coldp/config"
	"github.com/gnames/coldp/ent/coldp"
	"github.com/gnames/coldp/io/arcio"
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

func TestMeta(t *testing.T) {
	var err error
	var meta *coldp.Meta
	var arc coldp.Archive
	assert := assert.New(t)
	tests := []struct {
		msg, path, creatorGiven, contactOrganization, license string
	}{
		{"name-yaml", "name/ptero-yaml.zip", "Donald", "Species 2000", "CC0"},
		{"usg-yaml", "nameusage/ptero-yaml.zip", "Donald", "Species 2000", "cc0"},
	}

	for _, v := range tests {
		path := filepath.Join("..", "..", "testdata", v.path)
		arc, err = Extract(path)
		assert.Nil(err)
		err = arc.DirInfo()
		assert.Nil(err)
		meta, err = arc.Meta()
		assert.Nil(err)
		assert.Equal(v.creatorGiven, meta.Creators[0].Given)
		assert.Equal(v.contactOrganization, meta.Contact.Organization)
		assert.Equal(v.license, meta.License)
	}
}

func TestName(t *testing.T) {
	var err error
	var arc coldp.Archive
	assert := assert.New(t)
	path := filepath.Join("..", "..", "testdata", "name/ptero-yaml.zip")
	arc, err = Extract(path)
	assert.Nil(err)
	_ = arc

}

func Extract(path string) (coldp.Archive, error) {
	cfg := config.New()
	c := arcio.New(cfg, path)
	err := c.ResetCache()
	if err != nil {
		return nil, err
	}
	err = c.Extract()
	if err != nil {
		return nil, err
	}
	err = c.DirInfo()
	if err != nil {
		return nil, err
	}

	return c, nil
}
