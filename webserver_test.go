package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/magiconair/properties/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"project/config"
	"project/database"
	"project/structures"
	"project/web"
	"testing"
)

// тестируем вставку и чтение из БД
func TestPerson(t *testing.T) {
	// сначала проинициализируем конфиг, бд
	config.InitConfig()
	database.InitDB()

	// теперь проинициализируем веб-сервер для тестов
	ts := httptest.NewServer(web.SetupServer())
	defer ts.Close()

	var testWritePerson = structures.WritePerson{
		FirstName: "TestPersonFirstName",
		SurName:   "TestPersonSurName",
		Sex:       1,
	}

	client := &http.Client{}
	body, err := json.Marshal(testWritePerson)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/insert", ts.URL),
		bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Код ответа должен быть 200, получено %v", resp.StatusCode)
	}

	var response struct{
		Id int64
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(body, &response); err != nil {
		panic(err)
	}

	if response.Id == 0 {
		t.Fatalf("Не получен айди записи после вставки")
	}

	// Теперь получим нашу запись, которую мы вставили
	resp, err = http.Get(fmt.Sprintf("%s/get/%d", ts.URL, response.Id))

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Код ответа должен быть 200, получено %v", resp.StatusCode)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var readPerson structures.ReadPerson

	if err = json.Unmarshal(body, &readPerson); err != nil {
		panic(err)
	}

	assert.Equal(t, readPerson.Id, response.Id)
	assert.Equal(t, readPerson.FirstName, testWritePerson.FirstName)
	assert.Equal(t, readPerson.SurName, testWritePerson.SurName)
	assert.Equal(t, readPerson.Sex, testWritePerson.Sex)
}
