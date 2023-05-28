package cli

import (
	"reflect"
	"strings"
)

type (
	flag struct {
		FieldName string
		Name      string
		Aliases   []string
		Default   any
		Require   bool
		Usage     string
	}

	flagTagField   string
	flagTagHandler func(*flag, string) error
)

var (
	flagTagName string = "flag"

	ftf_Name    flagTagField = "name"
	ftf_Aliases flagTagField = "aliases"
	ftf_Require flagTagField = "require"
	ftf_Usage   flagTagField = "usage"

	flagHandlers = map[flagTagField]flagTagHandler{
		ftf_Name:    flagTag_Name,
		ftf_Aliases: flagTag_Aliases,
		ftf_Require: flagTag_Require,
		ftf_Usage:   flagTag_Usage,
	}
)

func flagTag_Name(flag *flag, value string) error {
	flag.Name = value
	return nil
}

func flagTag_Aliases(flag *flag, value string) error {
	flag.Aliases = strings.Split(value, "|")
	return nil
}

func flagTag_Require(flag *flag, value string) error {
	flag.Require = true
	return nil
}

func flagTag_Usage(flag *flag, value string) error {
	flag.Usage = value
	return nil
}

func anaylseFlags(a any) []*flag {

	dstType := reflect.TypeOf(a)
	defaultValue := reflect.ValueOf(a)

	if dstType.Kind() == reflect.Pointer {
		dstType = dstType.Elem()
		defaultValue = defaultValue.Elem()
	}

	fieldCout := dstType.NumField()

	flags := []*flag{}

	for i := 0; i < fieldCout; i++ {
		field := dstType.Field(i)
		tag := field.Tag.Get(flagTagName)

		flag := &flag{}

		tagFields := strings.Split(tag, ",")
		for _, tagField := range tagFields {
			words := strings.Split(tagField, ":")
			field := words[0]
			value := ""
			if len(words) >= 2 {
				value = words[1]
			}

			handler, exist := flagHandlers[flagTagField(field)]
			if exist {
				handler(flag, value)
			}
		}

		flag.FieldName = field.Name
		flag.Default = defaultValue.Field(i)

		flags = append(flags, flag)
	}

	return flags
}
