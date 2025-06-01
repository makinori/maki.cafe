package util

import (
	"errors"
	"net/http"
	"net/url"
)

func NotabotEncode(r *http.Request) string {
	v := url.Values{}
	v.Add("p", r.URL.Path)

	ref := r.Referer()
	if ref != "" {
		v.Add("r", r.Referer())
	}

	return v.Encode()
}

type Notabot struct {
	Path string
	Ref  string
}

func NotabotDecode(query string) (Notabot, error) {
	v, err := url.ParseQuery(query)
	if err != nil {
		return Notabot{}, err
	}

	n := Notabot{
		Path: v.Get("p"),
		Ref:  v.Get("r"),
	}

	if n.Path == "" {
		return Notabot{}, errors.New("missing path")
	}

	return n, nil
}
