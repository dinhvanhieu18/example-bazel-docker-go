package other

import (
	"fmt"
	"log"
)

func MergeString(inputs []string, pattern string) string {
	res := inputs[0]
	log.Println("abc")
	for _, item := range inputs[1:] {
		res = fmt.Sprintf("%s%s%s", res, pattern, item)
	}

	return res
}
