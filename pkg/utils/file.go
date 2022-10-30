package utils

import (
	"io"
	"os"
	"path/filepath"
)

func FileRead(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(f)
}

func FileWrite(filePath string, content []byte) error {
	return os.WriteFile(filePath, content, 0644)
}

func FileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return If(err == nil || os.IsExist(err), true, false)
}

func FileEmpty(filePath string) bool {
	stat, err := os.Stat(filePath)
	if err != nil {
		return true
	}
	return stat.Size() == 0
}

func FileTouch(filePath string) bool {
	base := filepath.Dir(filePath)
	_ = os.MkdirAll(base, 0755)
	_, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
	return If(err == nil, true, false)
}
