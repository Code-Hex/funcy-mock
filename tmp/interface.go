package main

import (
	"context"

	"github.com/kouzoh/mercari-offer-jp/app/model/offer"
)

// ServiceI interface holds interfaces for offer services
type ServiceI interface {
	// Offer interfaces
	GetOne(ctx context.Context, ID string) (offer.Offer, error)
	GetList(ctx context.Context, input *GetListInput) (map[string]string, error)
	ValidateItem(ctx context.Context, input *ValidateItemInput) (output *ValidateItemOutput, err error)
	Create(ctx context.Context, input *CreateInput) (offer.Offer, error)
	Accept(ctx context.Context, ID string) (offer.Offer, error)
	Decline(ctx context.Context, ID string) (offer.Offer, error)
	CancelByUser(ctx context.Context, ID string) (offer.Offer, error)
	Completing(ctx context.Context, ID string) (offer.Offer, error)
	BackToAccept(ctx context.Context, ID string) (offer.Offer, error)
	// Admin interfaces
	Update(ctx context.Context, input *UpdateInput) (offer.Offer, error)
	Delete(ctx context.Context, input *DeleteInput) (offer.Offer, error)
	CancelByAdmin(ctx context.Context, input *CancelByAdminInput) (offer.Offer, error)
	Expire(ctx context.Context, input *ExpireInput) error
	ExpireNotify(ctx context.Context, input *ExpireNotifyInput) error
	CompleteByAdmin(ctx *context.Context, input *CompleteByAdminInput) (err error, f func(error) error)
}

func main() {
}
