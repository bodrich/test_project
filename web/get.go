package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"project/cache"
	"project/database"
	"strconv"
)

// чтение из БД
func Get(c *gin.Context) {
	val := c.Param("id")
	id, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		checkError(c, fmt.Errorf("некорректный параметр id, причина: %s", err.Error()))
		return
	}

	// пробуем сначала загрузить из кеша
	person, success := cache.Load(id)
	if success {
		c.JSON(200, person)
		return
	}

	// чтение из бд
	row := database.SelectRequest.QueryRow(id)
	err = row.Scan(&person.Id, &person.FirstName, &person.SurName, &person.Sex)
	if checkError(c, err) != nil {
		return
	}

	// запишем в кеш, чтобы повторный раз не летать туда
	cache.Store(&person)

	c.JSON(200, person)
}
