package console

import (
	"fmt"
	"net/url"
	"strings"
)

type genericUrlBuilder struct {
	host  string
	path  string
	query []string
	hash  string
}

func (u genericUrlBuilder) Path(path string) genericUrlBuilder {
	n := u
	n.path = n.path + path
	return n
}

func (u genericUrlBuilder) Query(key string, value string) genericUrlBuilder {
	n := u
	n.query = append(n.query, fmt.Sprintf("%s=%s", key, url.QueryEscape(value)))
	return n
}

func (u genericUrlBuilder) String() string {
	b := strings.Builder{}
	b.WriteString("https://")
	b.WriteString(u.host)
	b.WriteString(u.path)

	for i, each := range u.query {
		if i == 0 {
			b.WriteRune('?')
		} else {
			b.WriteRune('&')
		}
		b.WriteString(each)
	}

	if len(u.hash) > 0 {
		b.WriteRune('#')
		b.WriteString(u.hash)
	}

	return b.String()
}

func (u genericUrlBuilder) Hash(s string) genericUrlBuilder {
	n := u
	n.hash = s
	return n
}

func (u genericUrlBuilder) Hashf(s string, args ...any) genericUrlBuilder {
	return u.Hash(fmt.Sprintf(s, args...))
}

func (u genericUrlBuilder) Pathf(path string, args ...any) genericUrlBuilder {
	return u.Path(fmt.Sprintf(path, args...))
}
