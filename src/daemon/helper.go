package daemon

import (
	"strconv"

	"github.com/nikita-petko/s3-forcer/flags"
)

func getChars() []interface{} {
	chars := []byte(*flags.CharCombinations)
	result := make([]interface{}, len(chars))

	for i, c := range chars {
		result[i] = c
	}

	return result
}

func joinTogether(args [][]interface{}) []string {
	var result []string
	for _, arg := range args {
		var str string
		for _, item := range arg {
			switch v := item.(type) {
			case string:
				str += v
			case int:
				str += strconv.Itoa(v)
			case float64:
				str += strconv.FormatFloat(v, 'f', -1, 64)
			case byte:
				str += string([]byte{v})
			default:
				continue
			}
		}
		result = append(result, str)
	}
	return result
}
