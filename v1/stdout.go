package cli

import (
	"fmt"
)

type (
	stdout struct {
		extWord string
	}
)

func newStdout(word ...string) *stdout {

	extWord := ""
	if len(word) > 0 {
		extWord = word[0]
	}

	return &stdout{
		extWord: extWord,
	}
}

func (out *stdout) Scope(word string) Stdout {
	return &stdout{
		extWord: out.extWord + word,
	}
}

func (stdout *stdout) NewLline() Stdout {
	fmt.Println()
	return stdout
}

func (stdout *stdout) Printf(format string, a ...any) Stdout {
	fmt.Printf(format, a...)
	return stdout
}

func (stdout *stdout) Print(a ...any) Stdout {
	fmt.Print(a...)
	return stdout
}

func (stdout *stdout) Printfln(format string, a ...any) Stdout {
	fmt.Printf(format, a...)
	fmt.Println()
	return stdout
}

func (stdout *stdout) Println(a ...any) Stdout {
	fmt.Println(a...)
	return stdout
}
