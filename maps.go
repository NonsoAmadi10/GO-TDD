package main
type Dictionary map[string]string
type DictionaryErr string

// errors are meant to be constant and immutable
const (
	ErrNotFound = DictionaryErr("sorry, could not find what you are looking for")
	ErrWordExist = DictionaryErr("sorry, it appears the word already exists!")
	ErrWordNotExist = DictionaryErr("sorry, it appears the word does not exit!")
)

// It looks similar to the errors.errorString implementation that powers errors.New. 
// However unlike errors.errorString this type is a constant expression.
func (e DictionaryErr) Error() string {
    return string(e)
}

// Search map
func (d Dictionary) Search(word string) (string, error){
	result, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return result, nil
}

// Add new key value pairs without duplicates
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

func (d Dictionary) Update(dict, definition string) error{
	d[dict] = definition

	_, err := d.Search(dict)

	switch err {
	case nil :
		d[dict] = definition
	case ErrNotFound:
		return ErrWordNotExist
	default:
		return err
	}

	return nil
}