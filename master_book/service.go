package master_book

import (
	"project-rest-api/entities"
	"time"
)

type Service interface {
	GetBooks() ([]entities.MasterBook, error)
	CreateMasterBook(input MasterBookInput) (entities.MasterBook, error)
	UpdateMasterBook(ID string, input MasterBookInput) (entities.MasterBook, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetBooks() ([]entities.MasterBook, error) {
	books, err := s.repository.GetMasterBook()

	if err != nil {
		return books, err
	}

	return books, err
}

func (s *service) CreateMasterBook(input MasterBookInput) (entities.MasterBook, error) {
	var masterBook entities.MasterBook
	masterBook.Name = input.Name
	masterBook.Amount = input.Amount
	masterBook.Price = input.Price
	masterBook.AuthorID = input.AuthorID
	masterBook.CreatedAt = time.Now()

	newBook, err := s.repository.StoreMasterBook(masterBook)
	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (s *service) UpdateMasterBook(ID string, input MasterBookInput) (entities.MasterBook, error) {
	masterBook, err := s.repository.Update(ID, input)
	if err != nil {
		return masterBook, err
	}

	return masterBook, nil
}
