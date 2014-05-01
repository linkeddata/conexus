package conexus

import (
	"flag"
	"os"
)

var (
	debug = flag.Bool("debug", false, "")
	dsn   = flag.String("dsn", "conexus@tcp(localhost:3306)/test", "")
)

func init() {
	if v := os.Getenv("DSN"); len(v) > 0 {
		*dsn = v
	}
}
