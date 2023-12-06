package gocon

import (
	"encoding/json"
	"os"
)

// Config ...
type Config interface {
	// Path returns the path it should be saved to.
	Path() string
}

// Load will load a config file and umarshal it.
func Load[T Config]() (T, error) {
	var data T

	b, err := os.ReadFile(data.Path())
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return data, err
	}

	return data, nil
}

// MustLoad will load a config file, or panic.
func MustLoad[T Config]() T {
	data, err := Load[T]()
	if err != nil {
		panic(err)
	}
	return data
}

// Save settings to a JSON file
func Save(data Config) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// 0600 allows only the owner of the file to access it
	return os.WriteFile(data.Path(), b, 0600)
}
