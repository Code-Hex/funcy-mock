package main

type ServiceIMock struct {
	GetOneMock          func() (reflect.Kind, error)
	GetListMock         func() (map[string]string, error)
	ValidateItemMock    func() (output *ValidateItemOutput, err error)
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
		GetOneMock:          func() (reflect.Kind, error) { return nil, nil },
		GetListMock:         func() (map[string]string, error) { return nil, nil },
		ValidateItemMock:    func() (output *ValidateItemOutput, err error) { return nil, nil },
		CreateMock:          func() (reflect.Kind, error) { return nil, nil },
		AcceptMock:          func() (reflect.Kind, error) { return nil, nil },
		DeclineMock:         func() (reflect.Kind, error) { return nil, nil },
		CancelByUserMock:    func() (reflect.Kind, error) { return nil, nil },
		CompletingMock:      func() (reflect.Kind, error) { return nil, nil },
		BackToAcceptMock:    func() (reflect.Kind, error) { return nil, nil },
		UpdateMock:          func() (reflect.Kind, error) { return nil, nil },
		DeleteMock:          func() (reflect.Kind, error) { return nil, nil },
		CancelByAdminMock:   func() (Num, error) { return 0, nil },
		ExpireMock:          func() error { return nil },
		ExpireNotifyMock:    func() error { return nil },
		CompleteByAdminMock: func() (err error, f func(error) error) { return nil, nil },
	}
}

func (s *ServiceIMock) GetOne(ctx reflect.Kind, ID string) (reflect.Kind, error) {
	return s.GetOneMock()
}

func (s *ServiceIMock) GetList(ctx reflect.Kind, input *GetListInput) (map[string]string, error) {
	return s.GetListMock()
}

func (s *ServiceIMock) ValidateItem(ctx reflect.Kind, input *ValidateItemInput) (output *ValidateItemOutput, err error) {
	return s.ValidateItemMock()
}

func (s *ServiceIMock) Create(ctx reflect.Kind, input *CreateInput) (reflect.Kind, error) {
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

func (s *ServiceIMock) Update(ctx reflect.Kind, input *UpdateInput) (reflect.Kind, error) {
	return s.UpdateMock()
}

func (s *ServiceIMock) Delete(ctx reflect.Kind, input *DeleteInput) (reflect.Kind, error) {
	return s.DeleteMock()
}

func (s *ServiceIMock) CancelByAdmin(ctx reflect.Kind, input *CancelByAdminInput) (Num, error) {
	return s.CancelByAdminMock()
}

func (s *ServiceIMock) Expire(ctx reflect.Kind, input *ExpireInput) error {
	return s.ExpireMock()
}

func (s *ServiceIMock) ExpireNotify(ctx reflect.Kind, input *ExpireNotifyInput) error {
	return s.ExpireNotifyMock()
}

func (s *ServiceIMock) CompleteByAdmin(ctx *reflect.Kind, input *CompleteByAdminInput) (err error, f func(error) error) {
	return s.CompleteByAdminMock()
}
