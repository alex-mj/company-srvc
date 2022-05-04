package domain

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Access   AccessMatrix
	IP       string `json:"ip"`
}

type AccessMatrix struct {
	Create bool `json:"create"`
	Read   bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}

type UserHandler interface {
	TokenCreator
	AccessMatrixGeter
	AccessMatrixModifier
}

// for UserService
type AccessMatrixGeter interface {
	GetAccessMatrix(token, IP string) (AccessMatrix, error)
}

type TokenCreator interface {
	GetToken(Name, Password string) (string, error)
}

// for IPChekService
type AccessMatrixModifier interface {
	ModifyByIP(AccessMatrix) (AccessMatrix, error)
}
