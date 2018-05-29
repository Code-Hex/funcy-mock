package tmp

import "context"

type Num int

type Person interface {
	Say(context.Context) (string, error)
	Age() Num
}

/*
// ServiceI interface holds interfaces for offer services
type ServiceI interface {
	// Offer interfaces
	GetOne(ctx reflect.Kind, ID string) (types.BasicKind, error)
	GetList(ctx reflect.Kind, input *reflect.Kind) (map[string]string, error)
	ValidateItem(ctx reflect.Kind, input *reflect.SliceHeader) (output *reflect.Type, err error)
	Create(ctx reflect.Kind, input *reflect.ChanDir) (reflect.Kind, error)
	Accept(ctx reflect.Kind, ID string) (reflect.Kind, error)
	Decline(ctx reflect.Kind, ID string) (reflect.Kind, error)
	CancelByUser(reflect.Kind, string) (reflect.Kind, error)
	Completing(ctx reflect.Kind, ID string) (reflect.Kind, error)
	BackToAccept(ctx reflect.Kind, ID string) (reflect.Kind, error)
	// Admin interfaces
	Update(ctx reflect.Kind, input chan string) (reflect.Kind, error)
	Delete(ctx reflect.Kind, input chan<- int) (reflect.Kind, error)
	CancelByAdmin(ctx reflect.Kind, input <-chan struct{}) (Num, error)
	Expire(ctx reflect.Kind, input *reflect.SelectCase) error
	ExpireNotify(ctx reflect.Kind, input *reflect.SliceHeader) error
	CompleteByAdmin(ctx *reflect.Kind, input *reflect.Method) (err error, f func(error) error)
}
*/
