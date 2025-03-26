package service

import (
	"io"
	"os"
	"path/filepath"
)

func (s *Service) saveFile(filePath string, src io.Reader) (int64, error) {
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return 0, err
	}

	dst, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
