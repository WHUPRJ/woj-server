package unzip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func(r *zip.ReadCloser) {
		_ = r.Close()
	}(r)

	handler := func(f *zip.File, dest string) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func(rc io.ReadCloser) {
			_ = rc.Close()
		}(rc)

		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(path, 0755)
		} else {
			_ = os.MkdirAll(filepath.Dir(path), 0755)

			file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				return err
			}
			defer func(file *os.File) {
				_ = file.Close()
			}(file)

			_, err = io.Copy(file, rc)
			if err != nil {
				return err
			}
		}

		return nil
	}

	for _, f := range r.File {
		err := handler(f, dest)
		if err != nil {
			_ = os.RemoveAll(dest)
			return err
		}
	}

	return nil
}
