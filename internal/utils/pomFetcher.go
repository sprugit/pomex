package utils

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

// Note : must be in caps to be exported
func FetchPOMs(work_dir string, estimate uint) ([]string, error) {

	if work_dir == "" {
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		work_dir = wd
	}
	if estimate < 10 {
		return nil, errors.New("Estimate can't be lower than 10.")
	}

	var poms []string = make([]string, 10, estimate)

	err := filepath.Walk(work_dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == "pom.xml" {
			poms = append(poms, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return poms, nil
}
