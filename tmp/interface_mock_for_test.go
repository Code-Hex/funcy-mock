package tmp

import "context"

type PersonMock struct {
	say func(context.Context) (string, error)
	age func() Num
}

func NewPersonMock() *PersonMock {
	return &PersonMock{
		say: func(context.Context) (string, error) { return "", nil },
		age: func() Num { return 0 },
	}
}

func (p *PersonMock) Say(c context.Context) (string, error) {
	return p.say(c)
}

func (p *PersonMock) SetSay(f func(context.Context) (string, error)) {
	if f == nil {
		panic("You should specify the mock function")
	}
	p.say = f
}

func (p *PersonMock) Age() Num {
	return p.age()
}

func (p *PersonMock) SetAge(f func() Num) {
	if f == nil {
		panic("You should specify the mock function")
	}
	p.age = f
}
