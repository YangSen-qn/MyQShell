package qn_util

import (
	"archive/zip"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"os"
	"path/filepath"
	"qshell/qn_error"
	"unicode/utf8"
)

func gbk2Utf8(text string) (string, qn_error.IError) {
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

func Unzip(zipFilePath string, unzipPath string) (err qn_error.IError) {
	zipReader, zipErr := zip.OpenReader(zipFilePath)
	if zipErr != nil {
		err = qn_error.NewFilePathError("Open zip file error" + zipErr.Error())
		return
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
				err = qn_error.NewFilePathError("Unsupported filename encoding")
				continue
			}
		}

		fullPath := filepath.Join(unzipPath, fileName)
		if fileInfo.IsDir() {
			mErr := os.MkdirAll(fullPath, 0775)
			if mErr != nil {
				err = qn_error.NewFilePathError("Mkdir error " + mErr.Error())
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
				err = qn_error.NewFilePathError("Unsupported filename encoding")
				continue
			}
		}

		fullPath := filepath.Join(unzipPath, fileName)
		if !fileInfo.IsDir() {
			//to be compatible with pkzip(4.5)
			fullPathDir := filepath.Dir(fullPath)
			mErr := os.MkdirAll(fullPathDir, 0755)
			if mErr != nil {
				err = qn_error.NewFilePathError("Mkdir error " + mErr.Error())
				continue
			}

			localFp, openErr := os.OpenFile(fullPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, zipFile.Mode())
			if openErr != nil {
				err = qn_error.NewFilePathError("Open local file error " + openErr.Error())
				continue
			}
			defer localFp.Close()

			zipFp, openErr := zipFile.Open()
			if openErr != nil {
				err = qn_error.NewFilePathError("Read zip content error " + openErr.Error())
				continue
			}
			defer zipFp.Close()

			_, wErr := io.Copy(localFp, zipFp)
			if wErr != nil {
				err = qn_error.NewFilePathError("Save zip content error " + wErr.Error())
				continue
			}
		}
	}
	return
}