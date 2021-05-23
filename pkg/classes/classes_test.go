package classes

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestCreateClass_invalidNameError(t *testing.T) {
	err := CreateClass("", time.Time{}, time.Time{}, 0)

	want := "cannot create class without name"
	got := err.Error()
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestCreateClass_invalidStartDateError(t *testing.T) {
	err := CreateClass("foo", time.Time{}, time.Time{}, 0)

	want := "cannot create class without start date"
	got := err.Error()
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestCreateClass_invalidEndDateError(t *testing.T) {
	err := CreateClass("foo", time.Now(), time.Time{}, 0)

	want := "cannot create class without end date"
	got := err.Error()
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestCreateClass_invalidCapacityError(t *testing.T) {
	err := CreateClass("foo", time.Now(), time.Now(), 0)

	want := "cannot create class without capacity"
	got := err.Error()
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestCreateClass_overlappingClassError(t *testing.T) {
	t.Cleanup(func() {
		DeleteClasses()
	})

	err := CreateClass("foo", time.Now(), time.Now(), 5)
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}

	err = CreateClass("foo", time.Now(), time.Now(), 5)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	want := fmt.Sprintf("class already exists in day %s", time.Now().Format("2006-01-02"))
	if want != err.Error() {
		t.Errorf("want: %s, got: %s", want, err.Error())
	}
}

func TestCreateClass_endDateIsSmallerError(t *testing.T) {
	err := CreateClass("foo", time.Now(), time.Now().Add(-24*time.Hour), 3)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	want := "end date cannot be smaller than start date"
	got := err.Error()
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestCreateClass_ok(t *testing.T) {
	t.Cleanup(func() {
		DeleteClasses()
	})

	err := CreateClass("foo", time.Now(), time.Now(), 5)
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
}

func TestCreateClass_multipleClassesOk(t *testing.T) {
	t.Cleanup(func() {
		DeleteClasses()
	})

	err := CreateClass("foo", time.Now(), time.Now(), 5)
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}

	err = CreateClass("bar", time.Now().Add(24*time.Hour), time.Now().Add(24*time.Hour), 5)
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}

	err = CreateClass("baz", time.Now().Add(2*24*time.Hour), time.Now().Add(2*24*time.Hour), 5)
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}

	classes := GetClasses()
	if len(classes) != 3 {
		t.Errorf("Expected 3 classes, got: %d", len(classes))
	}

	c := GetClass(time.Now())
	if c.Name != "foo" {
		t.Errorf("Expected class 'foo', got '%s'", c.Name)
	}

	c = GetClass(time.Now().Add(24 * time.Hour))
	if c.Name != "bar" {
		t.Errorf("Expected class 'bar', got '%s'", c.Name)
	}

	c = GetClass(time.Now().Add(2 * 24 * time.Hour))
	if c.Name != "baz" {
		t.Errorf("Expected class 'baz', got '%s'", c.Name)
	}
}

func TestBookClass_ok(t *testing.T) {
	t.Cleanup(func() {
		DeleteClasses()
	})

	err := CreateClass("foo", time.Now(), time.Now(), 5)
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}

	err = BookClass("Ben", time.Now())
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}

	err = BookClass("James", time.Now())
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}

	c := GetClass(time.Now())
	if len(c.BookedInMembers) != 2 {
		t.Errorf("Expected 2 bookings, got %d", len(c.BookedInMembers))
	}

	if !reflect.DeepEqual(c.BookedInMembers, []string{"Ben", "James"}) {
		t.Errorf("Expected booked members to be ['Ben', 'James'], got %v", c.BookedInMembers)
	}
}

func TestBookClass_noMatchingDate(t *testing.T) {
	err := BookClass("foo", time.Now())
	want := fmt.Errorf("there are no classes on %s", time.Now().Format("2006-01-02"))
	if err.Error() != want.Error() {
		t.Errorf("Expected error '%s', got: '%s'", want.Error(), err.Error())
	}
}
