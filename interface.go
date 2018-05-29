package funcy

import (
	"unicode"
	"unicode/utf8"
)

type Interface struct {
	Name   string
	Param  *Param
	Return *Return
}

func (i *Interface) PrivateName() string {
	if i.Name == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(i.Name)
	return string(unicode.ToLower(r)) + i.Name[n:]
}

type Param struct {
	TypeOnly string
	NameOnly string
	Field    string
}

type Return struct {
	Type  string
	Value string
}
