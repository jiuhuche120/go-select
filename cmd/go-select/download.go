package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"runtime"
	"strings"
)

const (
	GoDev = "https://go.dev/dl/"
)

func AvailableVersion() ([]string, error) {
	var versions [] string
	var end string
	if runtime.GOOS == "windows" {
		end = runtime.GOOS + "-" + runtime.GOARCH + ".zip"
	} else {
		end = runtime.GOOS + "-" + runtime.GOARCH + ".tar.gz"
	}
	resp, err := get(GoDev)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	strs := strings.Split(string(data), "\n")
	for _, str := range strs {
		s := regexp.MustCompile("go.*" + end).FindString(str)
		if s != "" {
			version := strings.Split(s, "\">")[0]
			versions = append(versions, version)
		}
	}
	return versions, nil
}

func get(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
