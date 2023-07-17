package cli

import (
	"encoding/json"
	"strings"
	"testing"
)

type exampleFlag struct {
	HasChild    bool `flag:"name:has-child2,usage:hasChild"`
	TestDefault bool
}

func TestAnaylseFlag(t *testing.T) {

	flags := anaylseFlags("", &exampleFlag{
		HasChild: true,
	})

	if len(flags) <= 0 {
		t.Error("flags length <= 0")
	}

	value, _ := json.Marshal(flags[0])
	t.Log(string(value))
	value, _ = json.Marshal(flags[1])
	t.Log(string(value))
}

type arrayFlag struct {
	TargetOS []string
}

func TestFlag_ArrayField(t *testing.T) {

	flags := anaylseFlags("", &arrayFlag{})

	if len(flags) <= 0 {
		t.Error("flags length <= 0")
	}

	if flags[0].Multiple == false {
		t.Error("flags 0 is not multiple")
	}
}

type People struct {
	Name string
}

type student struct {
	*People
	ClassRoom string
}

func TestFlag_EmbedField(t *testing.T) {

	flags := anaylseFlags("", &student{
		People: &People{
			Name: "",
		},
		ClassRoom: "",
	})

	if len(flags) <= 0 {
		t.Error("flags length <= 0")
	}

	for _, flag := range flags {
		value, _ := json.Marshal(flag)
		t.Log(string(value))
	}
}

func TestFlag_StringJoin(t *testing.T) {
	value := strings.Join([]string{"", "abc"}, ",")
	t.Log(value)
}

type toSetFlags struct {
	IsStudent bool
}

func TestFlag_SetBool(t *testing.T) {
	dst := &toSetFlags{}
	flag := &Flag{
		FieldName: "IsStudent",
		Name:      "",
		Aliases:   []string{},
		Default:   false,
		Require:   false,
		Multiple:  false,
		Usage:     "",
	}

	err := bindFlagsToStruct("true", flag, dst)
	if err != nil {
		t.Error(err)
	}

	if !dst.IsStudent {
		t.Errorf("IsStudent support to true but is flase")
	}
}
