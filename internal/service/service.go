package service

type SimpleOperationService interface {
	Plus(first int64, second int64) int64
	Minus(first int64, second int64) int64
	Multify(first int64, second int64) int64
}

type service struct{}

func New() *service {
	return &service{}
}

func (svc *service) Plus(first int64, second int64) int64 {
	return first + second
}

func (svc *service) Minus(first int64, second int64) int64 {
	return first - second
}

func (svc *service) Multify(first int64, second int64) int64 {
	return first * second
}
