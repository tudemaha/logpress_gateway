package service

import (
	"compress/gzip"
	"io"
	"os"
)

func CompressGZIP(filename string) error {
	file, err := os.Open("./dump/uncompressed/" + filename + ".sql")
	if err != nil {
		return err
	}
	defer file.Close()

	output, err := os.Create("./dump/compressed/" + filename + ".sql.gz")
	if err != nil {
		return err
	}
	defer output.Close()

	gzipWriter := gzip.NewWriter(output)
	defer gzipWriter.Close()

	_, err = io.Copy(gzipWriter, file)
	if err != nil {
		return err
	}

	return nil
}
