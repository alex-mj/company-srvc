package domain

//go:generate mockgen -source=user.go -destination=mock/user_mock.go -package=domain

type User struct {
	//ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Access   Access
	IP       string `json:"ip"`
}

type Access struct {
	Token string `json:"token"`
	AccessMatrix
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
	GetAccessMatrix(token string) (Access, error)
}

type TokenCreator interface {
	GetToken(Name, Password string) (string, error)
}

// for IPChekService
type AccessMatrixModifier interface {
	ModifyByIP(AccessMatrix) (AccessMatrix, error)
}
