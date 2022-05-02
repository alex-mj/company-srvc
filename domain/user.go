package domain

type User struct {
	ID       int    `json:"id"`
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

type UserService interface {
	AccessMatrixGeter
	AccessMatrixModifier
}

// for UserService
type AccessMatrixGeter interface {
	GetAccessMatrix(userID int) (Access, error)
}

// for IPChekService
type AccessMatrixModifier interface {
	ModifyByIP(AccessMatrix) (AccessMatrix, error)
}
