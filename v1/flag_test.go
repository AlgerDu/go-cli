package cli

import (
	"encoding/json"
	"testing"
)

type exampleFlag struct {
	HasChild    bool `flag:"name:has-child2,usage:hasChild"`
	TestDefault bool
}

func TestAnaylseFlag(t *testing.T) {

	flags := anaylseFlags(&exampleFlag{
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

	flags := anaylseFlags(&arrayFlag{})

	if len(flags) <= 0 {
		t.Error("flags length <= 0")
	}

	if flags[0].Multiple == false {
		t.Error("flags 0 is not multiple")
	}
}
