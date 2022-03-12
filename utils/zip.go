package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/foamzou/audio-get/logger"
)

func Unzip(zipFile, outputPath string) bool {
	archive, err := zip.OpenReader(zipFile)
	if err != nil {
		logger.Error(fmt.Sprintf("unzip %s failed, %v", zipFile, err))
		return false
	}
	defer func(archive *zip.ReadCloser) {
		_ = archive.Close()
	}(archive)

	for _, f := range archive.File {
		filePath := filepath.Join(outputPath, f.Name)
		logger.Debug("unzipping file ", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(outputPath)+string(os.PathSeparator)) {
			logger.Error("invalid file path")
			return false
		}
		if f.FileInfo().IsDir() {
			logger.Debug("creating directory...")
			err := os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				logger.Error(fmt.Sprintf("mkdir failed, path: %s, err: %v", filePath, err))
				return false
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			logger.Error(fmt.Sprintf("mkdir failed, path: %s, err: %v", filePath, err))
			return false
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			logger.Error(fmt.Sprintf("open file failed, path: %s, err: %v", filePath, err))
			return false
		}

		fileInArchive, err := f.Open()
		if err != nil {
			logger.Error(fmt.Sprintf("f.Open failed, err: %v", err))
			return false
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			logger.Error(fmt.Sprintf("copy failed, err: %v", err))
			return false
		}

		_ = dstFile.Close()
		_ = fileInArchive.Close()
	}
	return true
}
