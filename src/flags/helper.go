package flags

import (
	"os"
	"strconv"
	"time"
)

func getEnvironmentVariableOrFlag(name string, flag interface{}) error {
	if env, ok := os.LookupEnv(name); ok {
		var err error = nil
		switch v := flag.(type) {
		case *uint:
			var r uint64
			r, err = strconv.ParseUint(env, 10, 0)

			*v = uint(r)
		case *uint64:
			*v, err = strconv.ParseUint(env, 10, 64)
		case *int:
			var r int64
			r, err = strconv.ParseInt(env, 10, 0)

			*v = int(r)
		case *int64:
			*v, err = strconv.ParseInt(env, 10, 64)
		case *float64:
			*v, err = strconv.ParseFloat(env, 64)
		case *bool:
			*v, err = strconv.ParseBool(env)
		case *time.Duration:
			*v, err = time.ParseDuration(env)
		case *string:
			*v = env
		default:
			flag = env
		}

		return err
	}

	return nil
}
