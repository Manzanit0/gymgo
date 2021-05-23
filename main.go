package main

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/manzanit0/gymgo/pkg/classes"
)

func main() {
	app := fiber.New()
	setupRouter(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}

func setupRouter(app *fiber.App) {
	app.Put("/classes", createClassHandler)
	app.Put("/bookings", createBookingHandler)
}

type createClassPayload struct {
	Name      string `json:"name,omitempty"`
	StartDate date   `json:"start_date,omitempty"`
	EndDate   date   `json:"end_date,omitempty"`
	Capacity  int    `json:"capacity,omitempty"`
}

type date struct {
	time.Time
}

func (dt date) MarshalJSON() ([]byte, error) {
	return json.Marshal(dt.Time.Format("2006-01-02"))
}

func (dt *date) UnmarshalJSON(input []byte) error {
	strInput := strings.Trim(string(input), `"`)
	newTime, err := time.Parse("2006-01-02", strInput)
	if err != nil {
		return err
	}

	dt.Time = newTime
	return nil
}

func createClassHandler(c *fiber.Ctx) error {
	b := &createClassPayload{}
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := classes.CreateClass(b.Name, b.StartDate.Time, b.EndDate.Time, b.Capacity); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(b)
}

type createBookingPayload struct {
	MemberName string `json:"member_name,omitempty"`
	ClassDate  date   `json:"class_date,omitempty"`
}

func createBookingHandler(c *fiber.Ctx) error {
	b := &createBookingPayload{}
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := classes.BookClass(b.MemberName, b.ClassDate.Time); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(b)
}
