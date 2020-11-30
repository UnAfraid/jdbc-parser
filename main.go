package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	var field string
	flag.StringVar(&field, "field", "", "the field to read")
	flag.Parse()

	switch field {
	case "proto", "protocol", "host", "hostname", "port", "user", "username", "pass", "password", "path", "query":
	default:
		fmt.Printf("unknown field: %s\n", field)
		os.Exit(-1)
		return
	}

	if len(flag.Args()) < 1 {
		fmt.Printf("jdbc url expected: %s\n", field)
		os.Exit(-1)
	}

	baseUrl, err := url.Parse(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	u, err := url.Parse(baseUrl.Opaque)
	if err != nil {
		panic(err)
	}
	u.RawQuery = baseUrl.RawQuery

	switch field {
	case "proto", "protocol":
		fmt.Println(u.Scheme)
	case "host", "hostname":
		fmt.Println(u.Hostname())
	case "port":
		fmt.Println(u.Port())
	case "user", "username":
		fmt.Println(u.User.Username())
	case "pass", "password":
		password, _ := u.User.Password()
		fmt.Println(password)
	case "path":
		fmt.Println(strings.TrimLeft(u.Path, "/"))
	case "query":
		fmt.Println(u.RawQuery)
	}
	os.Exit(0)
}
