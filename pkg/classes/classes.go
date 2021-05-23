package classes

import (
	"fmt"
	"time"
)

// Current storage for classesByDate.
var classesByDate = make(map[time.Time]*Class)
var uniqueClasses = []*Class{}

type Class struct {
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Capacity  int
}

func CreateClass(name string, startDate time.Time, endDate time.Time, capacity int) error {
	if name == "" {
		return fmt.Errorf("cannot create class without name")
	}

	if startDate == (time.Time{}) {
		return fmt.Errorf("cannot create class without start date")
	}

	if endDate == (time.Time{}) {
		return fmt.Errorf("cannot create class without end date")
	}

	if capacity == 0 {
		return fmt.Errorf("cannot create class without capacity")
	}

	startDate = truncateToDate(startDate)
	endDate = truncateToDate(endDate)

	days := int(endDate.Sub(startDate).Hours() / 24)

	for i := 0; i <= days; i++ {
		d := startDate.Add(time.Hour * 24 * time.Duration(i))
		if _, exists := classesByDate[d]; exists {
			return fmt.Errorf("class already exists in day %s", d.Format("2006-01-02"))
		}
	}

	for i := 0; i <= days; i++ {
		d := startDate.Add(time.Hour * 24 * time.Duration(i))
		class := &Class{
			Name:      name,
			StartDate: startDate,
			EndDate:   endDate,
			Capacity:  capacity,
		}

		classesByDate[d] = class
		uniqueClasses = append(uniqueClasses, class)
	}

	return nil
}

func truncateToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func GetClasses() []Class {
	array := []Class{}
	for _, c := range uniqueClasses {
		array = append(array, *c)
	}
	return array
}

func GetClass(t time.Time) Class {
	t = truncateToDate(t)
	return *classesByDate[t]
}