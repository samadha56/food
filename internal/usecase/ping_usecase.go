package usecase

// PingUseCase interface defines a method to get a ping response.
type PingUseCase interface {
	GetPingResponse() string
}

// pingUseCase struct is an empty struct that implements the PingUseCase interface.
type pingUseCase struct{}

// NewPingUseCase function returns a new instance of pingUseCase.
func NewPingUseCase() PingUseCase {
	return &pingUseCase{}
}

// GetPingResponse method returns the string "pong".
func (u *pingUseCase) GetPingResponse() string {
	return "pong"
}
