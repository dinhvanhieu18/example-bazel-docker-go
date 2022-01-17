package util

import (
	"fmt"
)

func MergeString(inputs []string, pattern string) string {
	res := inputs[0]
	for _, item := range inputs[1:] {
		res = fmt.Sprintf("%s%s%s", res, pattern, item)
	}

	return res
}
