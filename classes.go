package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

var classes = make(map[string]*class)

type class struct {
	Name      string `json:"name,omitempty"`
	StartDate date   `json:"start_date,omitempty"`
	EndDate   date   `json:"end_date,omitempty"`
	Capacity  int
}

type date struct {
	time.Time
}

func (dt *date) MarshalJSON() ([]byte, error) {
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

func createClass(c *class) error {
	if c.Name == "" {
		return fmt.Errorf("cannot create class without name")
	}

	if c.EndDate == *new(date) {
		return fmt.Errorf("cannot create class without end date")
	}

	if c.StartDate == *new(date) {
		return fmt.Errorf("cannot create class without start date")
	}

	if c.Capacity == 0 {
		return fmt.Errorf("cannot create class without capacity")
	}

	classes[c.Name] = c

	return nil
}
