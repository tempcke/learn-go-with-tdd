package main

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	if !d.Has(word) {
		return "", ErrNotFound
	}

	return d[word], nil
}

func (d Dictionary) Add(word string, definition string) error {
	if d.Has(word) {
		return ErrWordExists
	}

	d[word] = definition

	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	if !d.Has(word) {
		return ErrWordDoesNotExist
	}

	d[word] = definition

	return nil
}

func (d Dictionary) Has(word string) bool {
	_, ok := d[word]

	return ok
}

func (d Dictionary) Delete(word string) {
	if d.Has(word) {
		delete(d, word)
	}
}