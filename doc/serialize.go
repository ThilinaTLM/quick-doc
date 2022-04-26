package doc

import (
	"io/ioutil"
	"os"
)

func CompiledDocFromFile(path string) (*CompiledDoc, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// read bytes
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	cd := &CompiledDoc{
		Json: b,
	}
	return cd, nil
}

func (cd *CompiledDoc) Save(path string) error {
	err := os.WriteFile(path, cd.Json, 0644)
	return err
}
