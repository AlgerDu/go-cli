package cli

import "reflect"

func Ext_StringTo(value string) string {

	newValue := make([]byte, 0, len(value)*2)

	upperCharIndex := -1
	hasUpper := false

	for i, currChar := range value {
		if currChar >= 65 && currChar <= 90 {
			if !hasUpper {
				hasUpper = true
				upperCharIndex = i
			}
		} else {
			if hasUpper {
				hasUpper = false

				for ; upperCharIndex < i-1; upperCharIndex++ {
					char := value[upperCharIndex] + 32
					newValue = append(newValue, byte(char))
				}

				if upperCharIndex != 0 {
					newValue = append(newValue, byte('-'))
				}

				newValue = append(newValue, byte(value[i-1])+32)
			}

			newValue = append(newValue, byte(currChar))
		}
	}

	if hasUpper {
		if (upperCharIndex) != 0 {
			newValue = append(newValue, byte('-'))
		}

		for k := upperCharIndex; k < len(value); k++ {
			char := value[k] + 32
			newValue = append(newValue, byte(char))
		}
	}

	return string(newValue)
}

func Ext_TypeIsArray(t reflect.Type) bool {

	typeKind := t.Kind()

	if typeKind == reflect.Pointer {
		typeKind = t.Elem().Kind()
	}

	if typeKind == reflect.Array || typeKind == reflect.Slice {
		return true
	}

	return false
}
