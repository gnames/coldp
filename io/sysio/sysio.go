package sysio

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gnames/coldp/config"
	"github.com/gnames/coldp/ent/coldp"
	"github.com/gnames/gnsys"
)

type sysio struct {
	path string
	cfg  config.Config
}

func New(cfg config.Config, path string) coldp.Sys {
	res := sysio{
		path: path,
		cfg:  cfg,
	}

	return &res
}

func (s *sysio) Extract() error {
	var err error
	if strings.HasPrefix(s.path, "http") {
		s.path, err = gnsys.Download(s.path, s.cfg.DownloadDir, true)
		if err != nil {
			return err
		}
	}
	err = s.unzip()
	if err != nil {
		return err
	}
	return nil
}

func (s *sysio) unzip() error {
	exists, _ := gnsys.FileExists(s.path)
	if !exists {
		return &coldp.ErrorFileMissing{Path: s.path}
	}

	// Open the zip file for reading.
	r, err := zip.OpenReader(s.path)
	if err != nil {
		return &coldp.ErrExtract{Path: s.path, Err: err}
	}
	defer r.Close()

	for _, f := range r.File {
		// Construct the full path for the file/directory and ensure its directory exists.
		fpath := filepath.Join(s.cfg.ExtractDir, f.Name)
		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return &coldp.ErrExtract{Path: fpath, Err: err}
		}

		// If it's a directory, move on to the next entry.
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Open the file within the zip.
		rc, err := f.Open()
		if err != nil {
			return &coldp.ErrExtract{Path: fpath, Err: err}
		}
		defer rc.Close()

		// Create a file in the filesystem.
		outFile, err := os.OpenFile(
			fpath,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			f.Mode(),
		)
		if err != nil {
			return &coldp.ErrExtract{Path: fpath, Err: err}
		}
		defer outFile.Close()

		// Copy the contents of the file from the zip to the new file.
		_, err = io.Copy(outFile, rc)
		if err != nil {
			return &coldp.ErrExtract{Path: fpath, Err: err}
		}
	}

	return nil
}

func (s *sysio) DirInfo() error {
	return nil
}

func (s *sysio) Meta() (*coldp.Meta, error) {
	return nil, nil
}
