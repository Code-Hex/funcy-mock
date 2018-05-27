package funcy

import "strings"

type Interface struct {
	Name        string
	Param       string
	ReturnType  string
	ReturnValue string
}

func (i *Interface) ReturnField() string {
	if strings.Contains(i.ReturnType, ",") {
		return "(" + i.ReturnType + ")"
	}
	return i.ReturnType
}
