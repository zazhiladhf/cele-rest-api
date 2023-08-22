package master_author

import (
	"project-rest-api/entities"
	"time"
)

type Service interface {
	GetAuthors() ([]entities.MasterAuthor, error)
	CreateMasterAuthor(input MasterAuthorInput) (entities.MasterAuthor, error)
	UpdateMasterAuthor(ID string, input MasterAuthorInput) (entities.MasterAuthor, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAuthors() ([]entities.MasterAuthor, error) {
	authors, err := s.repository.GetMasterAuthor()

	if err != nil {
		return []entities.MasterAuthor{}, err
	}

	return authors, err
}

func (s *service) CreateMasterAuthor(input MasterAuthorInput) (entities.MasterAuthor, error) {
	var masterAuthor entities.MasterAuthor
	masterAuthor.Name = input.Name
	masterAuthor.CreatedAt = time.Now()

	newAuthor, err := s.repository.StoreMasterAuthor(masterAuthor)
	if err != nil {
		return newAuthor, err
	}

	return newAuthor, nil
}

func (s *service) UpdateMasterAuthor(ID string, input MasterAuthorInput) (entities.MasterAuthor, error) {
	masterAuthor, err := s.repository.Update(ID, input)
	if err != nil {
		return masterAuthor, err
	}

	return masterAuthor, nil
}
