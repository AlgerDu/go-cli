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
		ff_Aliases: {
			UserSet: ffa_Aliases_UserSet,
			Defaut:  ffa_Aliases_Default,
		},
		ff_Require: {
			UserSet: ffa_Require_UserSet,
			Defaut:  ffa_Require_Default,
		},
		ff_Usage: {
			UserSet: ffa_Usage_UserSet,
			Defaut:  ffa_Usage_Default,
		},
	}
)

func ffa_Name_UserSet(flag *flag, value string) error {
	flag.Name = value
	return nil
}

func ffa_Name_Default(flag *flag, value string) error {
	flag.Name = Ext_StringTo(value)
	return nil
}

func ffa_Aliases_UserSet(flag *flag, value string) error {
	flag.Aliases = strings.Split(value, "|")
	return nil
}

func ffa_Aliases_Default(flag *flag, value string) error {
	return nil
}

func ffa_Require_UserSet(flag *flag, value string) error {
	flag.Require = true
	return nil
}

func ffa_Require_Default(flag *flag, value string) error {
	flag.Require = false
	return nil
}

func ffa_Usage_UserSet(flag *flag, value string) error {
	flag.Usage = value
	return nil
}

func ffa_Usage_Default(flag *flag, value string) error {
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
		userSets := map[string]string{}

		sets := strings.Split(tag, ",")
		for _, set := range sets {
			words := strings.Split(set, ":")
			field := words[0]
			value := ""
			if len(words) >= 2 {
				value = words[1]
			}

			userSets[field] = value
		}

		for name, handlers := range ff_Handlers {

			value, userSet := userSets[string(name)]
			if userSet {
				handlers.UserSet(flag, value)
			} else {
				handlers.Defaut(flag, field.Name)
			}
		}

		flag.FieldName = field.Name
		flag.Default = defaultValue.Field(i)

		flags = append(flags, flag)
	}

	return flags
}
