package lib

import (
	"fmt"
	"strings"
)

func SplitLine(line string, sep string) []string {

	var list []string
	for before, after, _ := strings.Cut(line, sep); !(before == "" && after == ""); {
		list = append(list, before)
		before, after, _ = strings.Cut(after, sep)
	}
	return list
}

func WrapErrors(error1 error, error2 error) error {
	if error1 == nil {
		return error2
	} else {
		wrapped_error := fmt.Errorf("%v, %v", error1, error2)
		return wrapped_error
	}
}
