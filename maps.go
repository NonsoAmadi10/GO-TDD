package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("sorry, could not find what you are looking for")

func (d Dictionary) Search(word string) (string, error){
	result, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return result, nil
}

func (d Dictionary) Add(dict, definition string){
	d[dict] = definition
}