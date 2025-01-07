package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	url_endpoint := "postgres://user:pass@host.com:5432/path?k=v#f"

	urlComponents, err := url.Parse(url_endpoint)

	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme: ", urlComponents.Scheme)
	fmt.Println("User: ", urlComponents.User)
	fmt.Println("UserName: ", urlComponents.User.Username())
	password, _ := urlComponents.User.Password()
	fmt.Println("Password: ", password)
	fmt.Println("Host: ", urlComponents.Host)
	fmt.Println("Port: ", urlComponents.Port())

	host, port, _ := net.SplitHostPort(url_endpoint)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println("Path", urlComponents.Path)
	fmt.Println("Fragment", urlComponents.Fragment)

	fmt.Println("Query", urlComponents.Query())
	m, _ := url.ParseQuery(urlComponents.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}
