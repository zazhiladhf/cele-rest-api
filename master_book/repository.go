package master_book

import (
	"project-rest-api/entities"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	GetMasterBook() ([]entities.MasterBook, error)
	StoreMasterBook(masterBook entities.MasterBook) (entities.MasterBook, error)
	Update(ID string, input MasterBookInput) (entities.MasterBook, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetMasterBook() ([]entities.MasterBook, error) {
	var mb []entities.MasterBook

	if err := r.db.Find(&mb).Error; err != nil {
		return nil, err
	}

	return mb, nil
}

func (r *repository) StoreMasterBook(masterBook entities.MasterBook) (entities.MasterBook, error) {
	var masterAuthor entities.MasterAuthor
	err := r.db.Table("master_authors").Where("id = ?", masterBook.AuthorID).First(&masterAuthor).Error

	if err != nil {
		return masterBook, err
	}

	// if errors.Is(err, gorm.ErrRecordNotFound) == true {
	// 	return masterBook, gorm.ErrRecordNotFound
	// }

	err = r.db.Create(&masterBook).Error

	if err != nil {
		return masterBook, err
	}

	return masterBook, nil
}

func (r *repository) Update(ID string, input MasterBookInput) (entities.MasterBook, error) {
	var masterBook entities.MasterBook

	err := r.db.Where("id = ?", ID).Find(&masterBook).Error
	if err != nil {
		return masterBook, err
	}

	masterBook.Name = input.Name
	masterBook.Amount = input.Amount
	masterBook.Price = input.Price
	masterBook.AuthorID = input.AuthorID
	masterBook.UpdatedAt = time.Now()

	err = r.db.Updates(&masterBook).Error
	if err != nil {
		return masterBook, err
	}

	return masterBook, nil
}
