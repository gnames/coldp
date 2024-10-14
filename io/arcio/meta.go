package arcio

import (
	"os"

	"github.com/gnames/coldp/ent/coldp"
	"gopkg.in/yaml.v3"
)

func (a *arcio) Meta() (res *coldp.Meta, err error) {
	if a.meta != nil {
		return a.meta, nil
	}
	if a.metaPath == "" {
		err = a.DirInfo()
		if err != nil {
			return nil, err
		}
	}

	bs, err := os.ReadFile(a.metaPath)
	if err != nil {
		return nil, err
	}

	var meta coldp.Meta
	err = yaml.Unmarshal(bs, &meta)
	if err != nil {
		return nil, err
	}

	a.meta = &meta

	return a.meta, nil
}
