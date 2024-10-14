package sysio

import (
	"github.com/gnames/coldp/ent/coldp"
	"github.com/gnames/gnsys"
)

func (s *sysio) ResetCache() error {
	err := s.emptyCacheDir()
	if err != nil {
		return err
	}

	err = gnsys.MakeDir(s.cfg.DownloadDir)
	if err != nil {
		return err
	}

	err = gnsys.MakeDir(s.cfg.ExtractDir)
	if err != nil {
		return err
	}
	return nil
}

func (s *sysio) emptyCacheDir() error {
	switch gnsys.GetDirState(s.cfg.CacheDir) {
	case gnsys.DirAbsent:
		return gnsys.MakeDir(s.cfg.CacheDir)
	case gnsys.DirEmpty:
		return nil
	case gnsys.DirNotEmpty:
		return gnsys.CleanDir(s.cfg.CacheDir)
	default:
		return &coldp.ErrBadDir{Dir: s.cfg.CacheDir}
	}
}
