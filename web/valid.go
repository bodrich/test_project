package web

import (
	"github.com/pkg/errors"
	"project/structures"
)

// проверяем, что по HTTP пришли корректные данные
func CheckValidPerson(person *structures.WritePerson) error {
	// по IEC_5218 пол может быть только 0, 1, 2 или 9
	if person.Sex != 0 && person.Sex != 1 &&
		person.Sex != 2 && person.Sex != 9 {
		return errors.Errorf("такой пол невозможен: %d", person.Sex)
	}

	var lenString = len(person.FirstName)
	if lenString < 4 || lenString > 62 {
		return errors.Errorf("некорретное имя")
	}

	lenString = len(person.SurName)
	if lenString < 4 || lenString > 62 {
		return errors.Errorf("некорретная фамилия")
	}

	return nil
}
