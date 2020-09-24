package web

import (
	"github.com/gin-gonic/gin"
	"project/database"
	"project/structures"
)

// вставка в БД
func Put(c *gin.Context) {
	var person structures.WritePerson
	if err := c.ShouldBindJSON(&person); checkError(c, err) != nil {
		return
	}

	if err := CheckValidPerson(&person); checkError(c, err) != nil {
		return
	}

	// записываем в бд
	var row = database.PutRequest.QueryRow(person.FirstName, person.SurName, person.Sex)
	// айдишник вставленной записи
	var id int64
	if err := row.Scan(&id); checkError(c, err) != nil {
		return
	}

	c.JSON(200, gin.H{"Id": id})

}
