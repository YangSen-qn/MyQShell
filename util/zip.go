package util

import (
	"archive/zip"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"os"
	"path/filepath"
	"unicode/utf8"
)

func gbk2Utf8(text string) (string, error) {
	var gDecoder = simplifiedchinese.GBK.NewDecoder()
	utf8Dst := make([]byte, len(text)*3)
	_, _, err := gDecoder.Transform(utf8Dst, []byte(text), true)
	if err != nil {
		return "", nil
	}
	gDecoder.Reset()
	utf8Bytes := make([]byte, 0)
	for _, b := range utf8Dst {
		if b != 0 {
			utf8Bytes = append(utf8Bytes, b)
		}
	}
	return string(utf8Bytes), nil
}

func Unzip(zipFilePath string, unzipPath string) error {
	zipReader, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	zipFiles := zipReader.File

	//list dir
	for _, zipFile := range zipFiles {
		fileInfo := zipFile.FileHeader.FileInfo()
		fileName := zipFile.FileHeader.Name

		//check charset utf8 or gbk
		if !utf8.Valid([]byte(fileName)) {
			fileName, err = gbk2Utf8(fileName)
			if err != nil {
				continue
			}
		}

		fullPath := filepath.Join(unzipPath, fileName)
		if fileInfo.IsDir() {
			err = os.MkdirAll(fullPath, 0775)
			if err != nil {
				continue
			}
		}
	}

	//list file
	for _, zipFile := range zipFiles {
		fileInfo := zipFile.FileHeader.FileInfo()
		fileName := zipFile.FileHeader.Name

		//check charset utf8 or gbk
		if !utf8.Valid([]byte(fileName)) {
			fileName, err = gbk2Utf8(fileName)
			if err != nil {
				continue
			}
		}

		fullPath := filepath.Join(unzipPath, fileName)
		if !fileInfo.IsDir() {
			//to be compatible with pkzip(4.5)
			fullPathDir := filepath.Dir(fullPath)
			err = os.MkdirAll(fullPathDir, 0755)
			if err != nil {
				continue
			}

			localFp, err := os.OpenFile(fullPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, zipFile.Mode())
			if err != nil {
				continue
			}
			defer localFp.Close()

			zipFp, err := zipFile.Open()
			if err != nil {
				continue
			}
			defer zipFp.Close()

			_, err = io.Copy(localFp, zipFp)
			if err != nil {
				continue
			}
		}
	}
	return err
}