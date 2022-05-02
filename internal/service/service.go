package service

// repository level
type AuthorizationService interface {
	// проверить токен
	// создать токен
}

// repository level
type CompanyStorage interface {
	// сохранить компанию в хранилище
	// обновить компанию в хранилище
	// запросить лист компаний из хранилища
}

// maybe Repository
type CompanyService struct {
	AuthorizationService
	CompanyStorage
}

func NewService(AuthService *AuthorizationService, cStorage *CompanyStorage) *CompanyService {
	return &CompanyService{
		AuthorizationService: AuthService,
		CompanyStorage:       cStorage,
	}
}
