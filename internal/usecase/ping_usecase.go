package usecase

type PingUseCase interface {
	GetPingResponse() string
}

type pingUseCase struct{}

func NewPingUseCase() PingUseCase {
	return &pingUseCase{}
}

func (u *pingUseCase) GetPingResponse() string {
	return "pong"
}