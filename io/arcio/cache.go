package arcio

import (
	"github.com/gnames/coldp/ent/coldp"
	"github.com/gnames/gnsys"
)

func (a *arcio) ResetCache() error {
	err := emptyCacheDir(a.cfg.CacheDir)
	if err != nil {
		return err
	}
	err = gnsys.MakeDir(a.cfg.DownloadDir)
	if err != nil {
		return err
	}

	err = gnsys.MakeDir(a.cfg.ExtractDir)
	if err != nil {
		return err
	}

	return nil
}

func emptyCacheDir(cacheDir string) error {
	switch gnsys.GetDirState(cacheDir) {
	case gnsys.DirAbsent:
		return gnsys.MakeDir(cacheDir)
	case gnsys.DirEmpty:
		return nil
	case gnsys.DirNotEmpty:
		return gnsys.CleanDir(cacheDir)
	default:
		return &coldp.ErrBadDir{Dir: cacheDir}
	}
}
