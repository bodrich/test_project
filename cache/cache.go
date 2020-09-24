package cache

import (
	"project/structures"
	"sync"
)

var cache sync.Map

// загрузка Person из кеша по Id, если Id не найден, возвращает false
func Load(id int64) (structures.ReadPerson, bool) {
	var person structures.ReadPerson
	result, ok := cache.Load(id)
	if !ok {
		return person, false
	}

	person = result.(structures.ReadPerson)
	return person, true
}

// сохранение Person в кеш
func Store(person *structures.ReadPerson) {
	cache.Store(person.Id, *person)
}


