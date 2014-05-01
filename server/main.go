package main

import (
	"flag"
	"github.com/linkeddata/conexus"
	"log"
	"net/http"
	"net/http/fcgi"
)

var (
	bind    = flag.String("bind", "", "bind address (empty: fcgi)")
	handler = new(conexus.Handler)
)

func init() {
	flag.Parse()
}

func main() {
	var err error

	if bind == nil || len(*bind) == 0 {
		err = fcgi.Serve(nil, handler)
	} else {
		err = http.ListenAndServe(*bind, handler)
	}

	if err != nil {
		log.Fatalln(err)
	}
}
