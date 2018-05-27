package main

import (
	"go/types"
	"reflect"
)

type ServiceIMock struct {
	GetOneMock          func() (types.BasicKind, error)
	GetListMock         func() (map[string]string, error)
	ValidateItemMock    func() (output *reflect.Type, err error)
	CreateMock          func() (reflect.Kind, error)
	AcceptMock          func() (reflect.Kind, error)
	DeclineMock         func() (reflect.Kind, error)
	CancelByUserMock    func() (reflect.Kind, error)
	CompletingMock      func() (reflect.Kind, error)
	BackToAcceptMock    func() (reflect.Kind, error)
	UpdateMock          func() (reflect.Kind, error)
	DeleteMock          func() (reflect.Kind, error)
	CancelByAdminMock   func() (Num, error)
	ExpireMock          func() error
	ExpireNotifyMock    func() error
	CompleteByAdminMock func() (err error, f func(error) error)
}

func NewServiceIMock() *ServiceIMock {
	return &ServiceIMock{
		GetOneMock:          func() (types.BasicKind, error) { return 0, nil },
		GetListMock:         func() (map[string]string, error) { return nil, nil },
		ValidateItemMock:    func() (output *reflect.Type, err error) { return nil, nil },
		CreateMock:          func() (reflect.Kind, error) { return 0, nil },
		AcceptMock:          func() (reflect.Kind, error) { return 0, nil },
		DeclineMock:         func() (reflect.Kind, error) { return 0, nil },
		CancelByUserMock:    func() (reflect.Kind, error) { return 0, nil },
		CompletingMock:      func() (reflect.Kind, error) { return 0, nil },
		BackToAcceptMock:    func() (reflect.Kind, error) { return 0, nil },
		UpdateMock:          func() (reflect.Kind, error) { return 0, nil },
		DeleteMock:          func() (reflect.Kind, error) { return 0, nil },
		CancelByAdminMock:   func() (Num, error) { return 0, nil },
		ExpireMock:          func() error { return nil },
		ExpireNotifyMock:    func() error { return nil },
		CompleteByAdminMock: func() (err error, f func(error) error) { return nil, nil },
	}
}

func (s *ServiceIMock) GetOne(ctx reflect.Kind, ID string) (types.BasicKind, error) {
	return s.GetOneMock()
}

func (s *ServiceIMock) GetList(ctx reflect.Kind, input *reflect.Kind) (map[string]string, error) {
	return s.GetListMock()
}

func (s *ServiceIMock) ValidateItem(ctx reflect.Kind, input *reflect.SliceHeader) (output *reflect.Type, err error) {
	return s.ValidateItemMock()
}

func (s *ServiceIMock) Create(ctx reflect.Kind, input *reflect.ChanDir) (reflect.Kind, error) {
	return s.CreateMock()
}

func (s *ServiceIMock) Accept(ctx reflect.Kind, ID string) (reflect.Kind, error) {
	return s.AcceptMock()
}

func (s *ServiceIMock) Decline(ctx reflect.Kind, ID string) (reflect.Kind, error) {
	return s.DeclineMock()
}

func (s *ServiceIMock) CancelByUser(ctx reflect.Kind, ID string) (reflect.Kind, error) {
	return s.CancelByUserMock()
}

func (s *ServiceIMock) Completing(ctx reflect.Kind, ID string) (reflect.Kind, error) {
	return s.CompletingMock()
}

func (s *ServiceIMock) BackToAccept(ctx reflect.Kind, ID string) (reflect.Kind, error) {
	return s.BackToAcceptMock()
}

func (s *ServiceIMock) Update(ctx reflect.Kind, input nil) (reflect.Kind, error) {
	return s.UpdateMock()
}

func (s *ServiceIMock) Delete(ctx reflect.Kind, input nil) (reflect.Kind, error) {
	return s.DeleteMock()
}

func (s *ServiceIMock) CancelByAdmin(ctx reflect.Kind, input nil) (Num, error) {
	return s.CancelByAdminMock()
}

func (s *ServiceIMock) Expire(ctx reflect.Kind, input *reflect.SelectCase) error {
	return s.ExpireMock()
}

func (s *ServiceIMock) ExpireNotify(ctx reflect.Kind, input *reflect.SliceHeader) error {
	return s.ExpireNotifyMock()
}

func (s *ServiceIMock) CompleteByAdmin(ctx *reflect.Kind, input *reflect.Method) (err error, f func(error) error) {
	return s.CompleteByAdminMock()
}
