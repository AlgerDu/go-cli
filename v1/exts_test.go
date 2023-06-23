package cli

import "testing"

func TestString1(t *testing.T) {
	src := "NewField"
	dst := "new-field"

	rst := Ext_StringTo(src)

	if rst != dst {
		t.Errorf("%s != %s", rst, dst)
	}
}

func TestString2(t *testing.T) {
	src := "NewID"
	dst := "new-id"

	rst := Ext_StringTo(src)

	if rst != dst {
		t.Errorf("%s != %s", rst, dst)
	}
}

func TestString3(t *testing.T) {
	src := "IDTest"
	dst := "id-test"

	rst := Ext_StringTo(src)

	if rst != dst {
		t.Errorf("%s != %s", rst, dst)
	}
}
