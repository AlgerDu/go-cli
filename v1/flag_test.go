package cli

import (
	"testing"
)

type exampleFlag struct {
	HasChild    bool `flag:"name:has-child2,usage:hasChild"`
	TestDefault bool
}

func TestAnaylseFlag(t *testing.T) {

	flags := anaylseFlags(&exampleFlag{
		HasChild: false,
	})

	if len(flags) <= 0 {
		t.Error("flags length <= 0")
	}

	t.Log(flags[0])
	t.Log(flags[1])
}
