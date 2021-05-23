package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/manzanit0/gymgo/pkg/classes"
)

func TestCreateClass_ok(t *testing.T) {
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

	resBody, _ := ioutil.ReadAll(res.Body)
	if strings.Contains(string(resBody), `"name": "Foo"`) {
		t.Errorf("expected response to contain name, got %d", resBody)
	}

	if strings.Contains(string(resBody), `"start_date": "1993-02-34"`) {
		t.Errorf("expected response to contain start_date, got %d", resBody)
	}

	if strings.Contains(string(resBody), `"end_date": "2021-04-21"`) {
		t.Errorf("expected response to contain end_date, got %d", resBody)
	}

	if strings.Contains(string(resBody), `"capacity": 20`) {
		t.Errorf("expected response to contain capacity, got %d", resBody)
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

func TestCreateClass_startDateIsSmaller(t *testing.T) {
	app := fiber.New()
	setupRouter(app)

	body := []byte(`
	{
		"name": "Foo",
		"start_date": "1993-02-24",
		"end_date": "1993-02-21",
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

	if res.StatusCode != 400 {
		t.Errorf("expected status 400, got %d", res.StatusCode)
	}

	resBody, _ := ioutil.ReadAll(res.Body)
	if !strings.Contains(string(resBody), "end date cannot be smaller than start date") {
		t.Errorf("expected parsing error, but got: %s", string(resBody))
	}
}

func TestCreateClass_cannotOverwriteAnExistingClass(t *testing.T) {
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

	// Run the same request twice to provoke a duplicate class.
	res, err := app.Test(req)
	res, err = app.Test(req)
	if err != nil {
		t.Errorf("error %v not expected", err.Error())
	}

	resBody, _ := ioutil.ReadAll(req.Body)
	if string(resBody) != string(body) {
		t.Errorf("response should contain request entity, but got %s", string(resBody))
	}

	if res.StatusCode != 400 {
		t.Errorf("expected status 400, got %d", res.StatusCode)
	}

	resBody, _ = ioutil.ReadAll(res.Body)
	if !strings.Contains(string(resBody), "class already exists") {
		t.Errorf("expected parsing error, but got: %s", string(resBody))
	}
}

func TestBookClass_ok(t *testing.T) {
	t.Cleanup(func() {
		classes.DeleteClasses()
	})

	err := classes.CreateClass("Kung-Fu", time.Date(1993, 2, 24, 0, 0, 0, 0, time.UTC), time.Now(), 5)
	if err != nil {
		t.Errorf("was not expecting error, got %s", err.Error())
	}

	app := fiber.New()
	setupRouter(app)

	body := []byte(`
	{
		"member_name": "James",
		"class_date": "1993-02-24"
	}`)

	req, err := http.NewRequest("PUT", "http://10.0.0.1/bookings", bytes.NewReader(body))
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

	resBody, _ := ioutil.ReadAll(res.Body)
	if strings.Contains(string(resBody), `"member_name": "James"`) {
		t.Errorf("expected response to contain member_name, got %d", resBody)
	}

	if strings.Contains(string(resBody), `"class_date": "1994-02-24"`) {
		t.Errorf("expected response to contain class_date, got %d", resBody)
	}
}
