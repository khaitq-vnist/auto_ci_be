package usecase

type ICreateIntegrationUseCase interface {
}
type CreateIntegrationUseCase struct {
	encryptUseCase IEncryptUseCase
}

func NewCreateIntegrationUseCase(encryptUseCase IEncryptUseCase) ICreateIntegrationUseCase {
	return &CreateIntegrationUseCase{
		encryptUseCase: encryptUseCase,
	}
}
