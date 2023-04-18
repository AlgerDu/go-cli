package cli

import "strings"

func anaylseArgs(args []string) ([]string, map[string]string) {

	args = args[1:]

	paths := []string{}
	flags := map[string]string{}

	for i := 0; i < len(args); i++ {

		word := args[i]

		if strings.HasPrefix(word, "-") {
			nextWord := args[i+1]

			if strings.HasPrefix(nextWord, "-") {
				flags[word] = ""
			} else {
				flags[word] = nextWord
				i++
			}
		} else {
			paths = append(paths, word)
		}

	}

	return paths, flags
}
