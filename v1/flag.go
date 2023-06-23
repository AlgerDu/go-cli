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
		Multiple  bool
		Usage     string
	}

	flagFieldAnaylser func(*flag, string) error

	flagFieldAnaylsers struct {
		UserSet flagFieldAnaylser
		Defaut  flagFieldAnaylser
	}

	flagField string
)

var (
	flagTagName string = "flag"

	ff_Name    flagField = "name"
	ff_Aliases flagField = "aliases"
	ff_Require flagField = "require"
	ff_Usage   flagField = "usage"

	ff_Handlers = map[flagField]flagFieldAnaylsers{
		ff_Name: {
			UserSet: ffa_Name_UserSet,
			Defaut:  ffa_Name_Default,
		},
	}
)

func ffa_Name_UserSet(flag *flag, value string) error {
	flag.Name = value
	return nil
}

func ffa_Name_Default(flag *flag, value string) error {
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

			handler, exist := flagHandlers[flagField(field)]
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
