package service

import (
	"edm-back/entity"
	"edm-back/repository"
)

type TypeDefinitionService interface {
	Save(typeDefinition entity.TypeDefinition) entity.TypeDefinition
	FindAll() []entity.TypeDefinition
}

type typeDefinitionService struct {
	typeDefinitionRepository repository.TypeDefinitionRepository
}

func New(repository repository.TypeDefinitionRepository) TypeDefinitionService{
	return &typeDefinitionService{
		typeDefinitionRepository: repository,
	}
}

func (service *typeDefinitionService) Save(typeDefinition entity.TypeDefinition) entity.TypeDefinition {
	service.typeDefinitionRepository.Save(typeDefinition)
	return typeDefinition
}

func (service *typeDefinitionService) FindAll() []entity.TypeDefinition {
	return service.typeDefinitionRepository.FindAll()
}