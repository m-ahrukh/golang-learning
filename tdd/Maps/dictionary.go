package maps

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound   = DictionaryErr("could not fint the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

// var (
// 	ErrNotFound   = errors.New("could not fint the word you were looking for")
// 	ErrWordExists = errors.New("cannot add word because it already exists")
// )

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)
	if err == ErrNotFound {
		d[word] = definition
	} else if err == nil {
		return ErrWordExists
	}
	return nil
}
