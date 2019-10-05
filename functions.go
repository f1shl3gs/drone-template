package template

import (
	"os"
	"time"
)

// like `1569406130, tz, format`
func date(i int64, s ...string) string {
	t := time.Unix(i, 0)
	if len(s) == 0 {
		return t.Format(time.RFC3339)
	}
	loc, err := time.LoadLocation(s[0])
	if err != nil {
		panic(err)
	}
	if len(s) == 1 {
		return t.In(loc).Format(time.RFC3339)
	}
	return t.In(loc).Format(s[1])
}

func now(s ...string) string {
	n := time.Now()
	return date(n.Unix(), s...)
}

func env(s string) string {
	return os.Getenv(s)
}
