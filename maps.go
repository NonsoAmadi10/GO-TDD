package main
type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound = DictionaryErr("sorry, could not find what you are looking for")
	ErrWordExist = DictionaryErr("sorry, it appears the word already exists!")
)

func (e DictionaryErr) Error() string {
    return string(e)
}

func (d Dictionary) Search(word string) (string, error){
	result, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return result, nil
}

func (d Dictionary) Add(dict, definition string) error{
	_, err := d.Search(dict)
	switch err {
	case ErrNotFound:
		d[dict] = definition
	case nil:
		return ErrWordExist
	default:
		return err
	}
	return nil
}