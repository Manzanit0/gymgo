package classes

import (
	"fmt"
	"time"
)

// Current storage for classes.
var classes = []*Class{}
var classesByDate = make(map[time.Time]*Class)

type Class struct {
	Name            string
	StartOn         time.Time
	EndOn           time.Time
	Capacity        int
	BookedInMembers []string
}

func CreateClass(name string, startOn time.Time, endOn time.Time, capacity int) error {
	if name == "" {
		return fmt.Errorf("cannot create class without name")
	}

	if startOn == (time.Time{}) {
		return fmt.Errorf("cannot create class without start date")
	}

	if endOn == (time.Time{}) {
		return fmt.Errorf("cannot create class without end date")
	}

	if capacity == 0 {
		return fmt.Errorf("cannot create class without capacity")
	}

	startOn = truncateToDate(startOn)
	endOn = truncateToDate(endOn)

	days := int(endOn.Sub(startOn).Hours() / 24)
	if days < 0 {
		return fmt.Errorf("end date cannot be smaller than start date")
	}

	for i := 0; i <= days; i++ {
		d := startOn.Add(time.Hour * 24 * time.Duration(i))
		if _, exists := classesByDate[d]; exists {
			return fmt.Errorf("class already exists in day %s", d.Format("2006-01-02"))
		}
	}

	for i := 0; i <= days; i++ {
		d := startOn.Add(time.Hour * 24 * time.Duration(i))
		class := &Class{
			Name:     name,
			StartOn:  startOn,
			EndOn:    endOn,
			Capacity: capacity,
		}

		classesByDate[d] = class
		classes = append(classes, class)
	}

	return nil
}

func truncateToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func BookClass(memberName string, classDate time.Time) error {
	classDate = truncateToDate(classDate)
	c, ok := classesByDate[classDate]
	if !ok {
		return fmt.Errorf("there are no classes on %s", classDate.Format("2006-01-02"))
	}

	c.BookedInMembers = append(c.BookedInMembers, memberName)
	return nil
}

func GetClasses() []Class {
	array := []Class{}
	for _, c := range classes {
		array = append(array, *c)
	}
	return array
}

func GetClass(t time.Time) Class {
	t = truncateToDate(t)
	return *classesByDate[t]
}

func DeleteClasses() {
	classesByDate = make(map[time.Time]*Class)
	classes = []*Class{}
}
