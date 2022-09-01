package mydict

import "errors"

// Type
type Dictionary map[string]string

var errSearch = errors.New("Not Found")
var errAdd = errors.New("It already exists.")
var errUpdate = errors.New("You can't update non-existing word.")
var errDelete = errors.New("You can't delete non-existing word.")

// Search word and if not exist return error.
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errSearch
}

// Add a word and definition in dictionary if word not exists.
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	// 검색되지 않을때 단어와 뜻을 추가한다. 검색된다면 에러를 발생시킨다.
	switch err {
	case errSearch:
		d[word] = def
		return nil
	case nil:
		return errAdd
	}
	return nil
}

// Update word to another word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = def
	case errSearch:
		return errUpdate
	}
	return nil
}

// Delete word
// 존재하지 않는 단어는 삭제할 수 없다.
func (d Dictionary) Worddelete(word string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		delete(d, word)
		return nil
	case errSearch:
		return errDelete
	}
	return nil
}
