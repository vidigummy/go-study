package mydict

import (
	"errors"
)

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("Cant update non-existing word")
	errCantDelete = errors.New("Cant delete non-existing word")
)

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	} else {
		return "-1", errNotFound
	}
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	// if err == errNotFound {
	// 	d[word] = def
	// } else if err != nil {
	// 	return errWordExists
	// }
	// 스위치로도 써보자.
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// UPDATE
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}
