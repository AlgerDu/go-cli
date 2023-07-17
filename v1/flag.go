package cli

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type (
	Flag struct {
		FieldName string
		Name      string
		Aliases   []string
		Default   any
		Require   bool
		Multiple  bool
		Usage     string
	}

	GlobalFlag struct {
		*Flag
		Action PipelineAction
	}

	flagFieldAnaylser func(*Flag, string) error

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

	ErrNoField             = errors.New("no field")
	ErrNotSupportFieldType = errors.New("not support field type")
)

func anaylseFlags(parentPath string, flagDefaultValue any) []*Flag {

	dstType := reflect.TypeOf(flagDefaultValue)
	defaultValue := reflect.ValueOf(flagDefaultValue)

	if dstType.Kind() == reflect.Pointer {
		dstType = dstType.Elem()
		defaultValue = defaultValue.Elem()
	}

	fieldCout := dstType.NumField()
	flags := []*Flag{}

	for i := 0; i < fieldCout; i++ {
		field := dstType.Field(i)

		fieldPath := field.Name
		if len(parentPath) > 0 {
			fieldPath = fmt.Sprintf("%s.%s", parentPath, fieldPath)
		}

		rv := defaultValue.Field(i)
		var fieldValue any
		if rv.CanInterface() {
			fieldValue = rv.Interface()
		} else {
			continue
		}

		fieldType := field.Type
		if fieldType.Kind() == reflect.Pointer {
			fieldType = fieldType.Elem()
		}

		if fieldType.Kind() == reflect.Struct {
			flags = append(flags, anaylseFlags(fieldPath, fieldValue)...)
			continue
		}

		tag := field.Tag.Get(flagTagName)
		flag := &Flag{}
		sets := map[string]string{}
		tagSets := strings.Split(tag, ",")

		for _, set := range tagSets {
			words := strings.Split(set, ":")
			field := words[0]
			value := ""
			if len(words) >= 2 {
				value = words[1]
			}

			sets[field] = value
		}

		for name, handlers := range ff_Handlers {
			value, userSet := sets[string(name)]
			if userSet {
				handlers.UserSet(flag, value)
			} else {
				handlers.Defaut(flag, field.Name)
			}
		}

		flag.FieldName = fieldPath
		flag.Default = defaultValue.Field(i).Interface()

		if Ext_TypeIsArray(field.Type) {
			flag.Multiple = true
		}

		flags = append(flags, flag)
	}

	return flags
}

func bindFlagsToStruct(value string, flag *Flag, dst any) error {

	dstValue := reflect.ValueOf(dst)
	if dstValue.Kind() == reflect.Pointer {
		dstValue = dstValue.Elem()
	}

	names := strings.Split(flag.FieldName, ".")
	fieldValue := dstValue

	for _, name := range names {
		fieldValue = fieldValue.FieldByName(name)

		if fieldValue == (reflect.Value{}) {
			return ErrNoField
		}

		if fieldValue.Kind() == reflect.Pointer && fieldValue.IsNil() {
			fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
			fieldValue = fieldValue.Elem()
		}
	}

	var v any
	var err error

	switch fieldValue.Kind() {
	case reflect.Bool:
		v, err = strconv.ParseBool(value)
	default:
		return ErrNotSupportFieldType
	}

	if err != nil {
		return err
	}

	fieldValue.Set(reflect.ValueOf(v))
	return nil
}

func ffa_Name_UserSet(flag *Flag, value string) error {
	flag.Name = value
	return nil
}

func ffa_Name_Default(flag *Flag, value string) error {
	flag.Name = Ext_StringTo(value)
	return nil
}

func ffa_Aliases_UserSet(flag *Flag, value string) error {
	flag.Aliases = strings.Split(value, "|")
	return nil
}

func ffa_Aliases_Default(flag *Flag, value string) error {
	return nil
}

func ffa_Require_UserSet(flag *Flag, value string) error {
	flag.Require = true
	return nil
}

func ffa_Require_Default(flag *Flag, value string) error {
	flag.Require = false
	return nil
}

func ffa_Usage_UserSet(flag *Flag, value string) error {
	flag.Usage = value
	return nil
}

func ffa_Usage_Default(flag *Flag, value string) error {
	flag.Usage = Ext_StringTo(value)
	return nil
}
