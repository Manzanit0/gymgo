package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/manzanit0/gymgo/pkg/classes"
)

func TestCreateClass_successful(t *testing.T) {
	t.Cleanup(func() {
		classes.DeleteClasses()
	})

	app := fiber.New()
	setupRouter(app)

	body := []byte(`
	{
		"name": "Foo",
		"start_date": "1993-02-24",
		"end_date": "2021-04-21",
		"capacity": 20
	}`)

	req, err := http.NewRequest("PUT", "http://10.0.0.1/classes", bytes.NewReader(body))
	if err != nil {
		t.Errorf("error %v not expected", err.Error())
	}

	req.Header.Add("content-type", "application/json")

	res, err := app.Test(req)
	if err != nil {
		t.Errorf("error %v not expected", err.Error())
	}

	if res.StatusCode != 201 {
		t.Errorf("expected status 201, got %d", res.StatusCode)
	}

	resBody, _ := ioutil.ReadAll(req.Body)
	if string(resBody) != string(body) {
		t.Errorf("response should contain request entity, but got %s", string(resBody))
	}
}

func TestCreateClass_missingName(t *testing.T) {
	app := fiber.New()
	setupRouter(app)

	body := []byte(`
	{
		"start_date": "1993-02-24",
		"end_date": "1993-02-24"
	}`)

	req, err := http.NewRequest("PUT", "http://10.0.0.1/classes", bytes.NewReader(body))
	if err != nil {
		t.Errorf("error %v not expected", err.Error())
	}

	req.Header.Add("content-type", "application/json")

	res, err := app.Test(req)
	if err != nil {
		t.Errorf("error %v not expected", err.Error())
	}

	if res.StatusCode != 400 {
		t.Errorf("expected status 400, got %d", res.StatusCode)
	}

	resBody, _ := ioutil.ReadAll(res.Body)
	if !strings.Contains(string(resBody), "cannot create class without name") {
		t.Errorf("expected missing field error, but got: %s", string(resBody))
	}
}
func TestCreateClass_missingStartDate(t *testing.T) {
	app := fiber.New()
	setupRouter(app)

	body := []byte(`
	{
		"name": "Foo",
		"end_date": "1993-02-24"
	}`)

	req, err := http.NewRequest("PUT", "http://10.0.0.1/classes", bytes.NewReader(body))
	if err != nil {
		t.Errorf("error %v not expected", err.Error())
	}

	req.Header.Add("content-type", "application/json")

	res, err := app.Test(req)
	if err != nil {
		t.Errorf("error %v not expected", err.Error())
	}

	if res.StatusCode != 400 {
		t.Errorf("expected status 400, got %d", res.StatusCode)
	}

	resBody, _ := ioutil.ReadAll(res.Body)
	if !strings.Contains(string(resBody), "cannot create class without start date") {
		t.Errorf("expected missing field error, but got: %s", string(resBody))
	}
}

func TestCreateClass_missingEndDate(t *testing.T) {
	app := fiber.New()
	setupRouter(app)

	body := []byte(`
	{
		"name": "Foo",
		"start_date": "1993-02-24"
	}`)

	req, err := http.NewRequest("PUT", "http://10.0.0.1/classes", bytes.NewReader(body))
	if err != nil {
		t.Errorf("error %v not expected", err.Error())
	}

	req.Header.Add("content-type", "application/json")

	res, err := app.Test(req)
	if err != nil {
		t.Errorf("error %v not expected", err.Error())
	}

	if res.StatusCode != 400 {
		t.Errorf("expected status 400, got %d", res.StatusCode)
	}

	resBody, _ := ioutil.ReadAll(res.Body)
	if !strings.Contains(string(resBody), "cannot create class without end date") {
		t.Errorf("expected missing field error, but got: %s", string(resBody))
	}
}

func TestCreateClass_missingCapacity(t *testing.T) {
	app := fiber.New()
	setupRouter(app)

	body := []byte(`
	{
		"name": "Foo",
		"start_date": "1993-02-24",
		"end_date": "2021-04-21"
	}`)

	req, err := http.NewRequest("PUT", "http://10.0.0.1/classes", bytes.NewReader(body))
	if err != nil {
		t.Errorf("error %v not expected", err.Error())
	}

	req.Header.Add("content-type", "application/json")

	res, err := app.Test(req)
	if err != nil {
		t.Errorf("error %v not expected", err.Error())
	}

	if res.StatusCode != 400 {
		t.Errorf("expected status 400, got %d", res.StatusCode)
	}

	resBody, _ := ioutil.ReadAll(res.Body)
	if !strings.Contains(string(resBody), "cannot create class without capacity") {
		t.Errorf("expected missing field error, but got: %s", string(resBody))
	}
}

func TestCreateClass_invalidDate(t *testing.T) {
	app := fiber.New()
	setupRouter(app)

	body := []byte(`
	{
		"name": "Foo",
		"start_date": "1993-02-24",
		"end_date": "nope"
	}`)

	req, err := http.NewRequest("PUT", "http://10.0.0.1/classes", bytes.NewReader(body))
	if err != nil {
		t.Errorf("error %v not expected", err.Error())
	}

	req.Header.Add("content-type", "application/json")

	res, err := app.Test(req)
	if err != nil {
		t.Errorf("error %v not expected", err.Error())
	}

	if res.StatusCode != 400 {
		t.Errorf("expected status 400, got %d", res.StatusCode)
	}

	resBody, _ := ioutil.ReadAll(res.Body)
	if !strings.Contains(string(resBody), "cannot parse") {
		t.Errorf("expected parsing error, but got: %s", string(resBody))
	}
}
