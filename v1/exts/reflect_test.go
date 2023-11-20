package exts

import (
	"reflect"
	"testing"
)

type (
	PeopleTestFlag struct {
		Name       string
		BornYear   int
		ParentName []string
	}

	StudentTestFlag struct {
		*PeopleTestFlag

		ClassRoom string
	}
)

func TestReflect_NewSlice(t *testing.T) {
	people := &PeopleTestFlag{}

	fieldValue := reflect.ValueOf(people).Elem().FieldByName("ParentName")
	t.Log(fieldValue.Kind())
	fieldValue.Set(reflect.MakeSlice(fieldValue.Type(), 0, 0))

	if people.ParentName == nil {
		t.Error("parentname is nil")
	}
}

func TestReflect_NewStruct(t *testing.T) {
	student := &StudentTestFlag{}

	fieldValue := reflect.ValueOf(student).Elem().FieldByName("PeopleTestFlag")
	t.Log(fieldValue.Kind())
	fieldValue.Set(reflect.New(fieldValue.Type().Elem()))

	if student.PeopleTestFlag == nil {
		t.Error("PeopleTestFlag is nil")
	}
}
