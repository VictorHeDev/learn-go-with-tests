package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrWordExists = errors.New("word already exists in the dictionary")

// func Search(dictionary Dictionary, key string) string {
// 	return dictionary[key]
// }

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) (string, error) {
	definition, ok := d[key]

	if !ok {
		d[key] = value
	}

	return definition, nil
}
