package service

type (
	Services interface {
	}
	services struct {
	}
)

func NewServices() Services {
	return &services{}
}
