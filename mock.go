package main

type ServiceIMock struct {
	GetOneMock          func() (offer.Offer, error)
	GetListMock         func() (map[string]string, error)
	ValidateItemMock    func() (output *ValidateItemOutput, err error)
	CreateMock          func() (offer.Offer, error)
	AcceptMock          func() (offer.Offer, error)
	DeclineMock         func() (offer.Offer, error)
	CancelByUserMock    func() (offer.Offer, error)
	CompletingMock      func() (offer.Offer, error)
	BackToAcceptMock    func() (offer.Offer, error)
	UpdateMock          func() (offer.Offer, error)
	DeleteMock          func() (offer.Offer, error)
	CancelByAdminMock   func() (offer.Offer, error)
	ExpireMock          func() error
	ExpireNotifyMock    func() error
	CompleteByAdminMock func() (err error, f func(error) error)
}

func NewServiceIMock() *ServiceIMock {
	return &ServiceIMock{
		GetOneMock:          func() (offer.Offer, error) { return nil, nil },
		GetListMock:         func() (map[string]string, error) { return nil, nil },
		ValidateItemMock:    func() (output *ValidateItemOutput, err error) { return nil, nil },
		CreateMock:          func() (offer.Offer, error) { return nil, nil },
		AcceptMock:          func() (offer.Offer, error) { return nil, nil },
		DeclineMock:         func() (offer.Offer, error) { return nil, nil },
		CancelByUserMock:    func() (offer.Offer, error) { return nil, nil },
		CompletingMock:      func() (offer.Offer, error) { return nil, nil },
		BackToAcceptMock:    func() (offer.Offer, error) { return nil, nil },
		UpdateMock:          func() (offer.Offer, error) { return nil, nil },
		DeleteMock:          func() (offer.Offer, error) { return nil, nil },
		CancelByAdminMock:   func() (offer.Offer, error) { return nil, nil },
		ExpireMock:          func() error { return nil },
		ExpireNotifyMock:    func() error { return nil },
		CompleteByAdminMock: func() (err error, f func(error) error) { return nil, nil },
	}
}

func (s *ServiceIMock) GetOne(ctx context.Context, ID string) (offer.Offer, error) {
	return s.GetOneMock()
}

func (s *ServiceIMock) GetList(ctx context.Context, input *GetListInput) (map[string]string, error) {
	return s.GetListMock()
}

func (s *ServiceIMock) ValidateItem(ctx context.Context, input *ValidateItemInput) (output *ValidateItemOutput, err error) {
	return s.ValidateItemMock()
}

func (s *ServiceIMock) Create(ctx context.Context, input *CreateInput) (offer.Offer, error) {
	return s.CreateMock()
}

func (s *ServiceIMock) Accept(ctx context.Context, ID string) (offer.Offer, error) {
	return s.AcceptMock()
}

func (s *ServiceIMock) Decline(ctx context.Context, ID string) (offer.Offer, error) {
	return s.DeclineMock()
}

func (s *ServiceIMock) CancelByUser(ctx context.Context, ID string) (offer.Offer, error) {
	return s.CancelByUserMock()
}

func (s *ServiceIMock) Completing(ctx context.Context, ID string) (offer.Offer, error) {
	return s.CompletingMock()
}

func (s *ServiceIMock) BackToAccept(ctx context.Context, ID string) (offer.Offer, error) {
	return s.BackToAcceptMock()
}

func (s *ServiceIMock) Update(ctx context.Context, input *UpdateInput) (offer.Offer, error) {
	return s.UpdateMock()
}

func (s *ServiceIMock) Delete(ctx context.Context, input *DeleteInput) (offer.Offer, error) {
	return s.DeleteMock()
}

func (s *ServiceIMock) CancelByAdmin(ctx context.Context, input *CancelByAdminInput) (offer.Offer, error) {
	return s.CancelByAdminMock()
}

func (s *ServiceIMock) Expire(ctx context.Context, input *ExpireInput) error {
	return s.ExpireMock()
}

func (s *ServiceIMock) ExpireNotify(ctx context.Context, input *ExpireNotifyInput) error {
	return s.ExpireNotifyMock()
}

func (s *ServiceIMock) CompleteByAdmin(ctx *context.Context, input *CompleteByAdminInput) (err error, f func(error) error) {
	return s.CompleteByAdminMock()
}
