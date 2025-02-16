package service

import "os"

func DeleteUncompressed(filename string) error {
	err := os.Remove("./dump/uncompressed/" + filename + ".sql")
	if err != nil {
		return err
	}

	return nil
}

func DeleteCompressed(filename string) error {
	err := os.Remove("./dump/compressed/" + filename + ".sql.gz")
	if err != nil {
		return err
	}

	return nil
}
