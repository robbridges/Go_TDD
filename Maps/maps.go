package Maps

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

func (d Dictionary) Search(k string) (string, error) {
	var notFoundError = errors.New(fmt.Sprintf("no value found for %s", k))
	if d[k] == "" {
		return "No value found of K", notFoundError
	}
	return d[k], nil
}

// We're modifying a dictionary but not passing in a pointer. This is because a map is a reference type, and it's just a pointer
// to a go data thing called a hmap. So this just creates a copy of the map and the underlying hmap
func (d Dictionary) Add(k, v string) error {
	keyAlreadyExists := errors.New("This key already exists, please use the Update method to change the value")
	_, ok := d[k]
	if ok {
		return keyAlreadyExists
	}
	d[k] = v
	return nil
}

func (d Dictionary) Remove(k string) error {
	err := notFoundHelper(d, k)
	if err != nil {
		return err
	}

	delete(d, k)
	return nil
}

func (d Dictionary) Update(k, v string) error {
	err := notFoundHelper(d, k)
	if err != nil {
		return err
	}
	d[k] = v
	return nil
}

func notFoundHelper(d Dictionary, k string) error {
	notFoundError := errors.New("This value was not in the map")
	_, ok := d[k]
	if !ok {
		return notFoundError
	}
	return nil
}
