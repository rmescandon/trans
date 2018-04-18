package trans

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

func listDirs(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []string{}, errors.WithMessage(err, "Could not list dir subdirs")
	}

	ds := []string{}
	for _, file := range files {
		ds = append(ds, file.Name())
	}
	return ds, nil
}
