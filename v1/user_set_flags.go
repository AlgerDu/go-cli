package cli

type UserSetFlags map[string][]string

func (usfs UserSetFlags) Set(key string, value string) {
	values, exist := usfs[key]
	if !exist {
		values = []string{}
	}

	usfs[key] = append(values, value)
}

func (usfs UserSetFlags) Range(handler func(string, string)) {
	for key, values := range usfs {
		for _, value := range values {
			handler(key, value)
		}
	}
}
