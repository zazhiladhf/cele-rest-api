package master_author

import (
	"project-rest-api/entities"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	GetMasterAuthor() ([]entities.MasterAuthor, error)
	StoreMasterAuthor(masterAuthor entities.MasterAuthor) (entities.MasterAuthor, error)
	// FindByID(ID int) (entities.MasterAuthor, error)
	Update(ID string, input MasterAuthorInput) (entities.MasterAuthor, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetMasterAuthor() ([]entities.MasterAuthor, error) {
	var masterAuthor []entities.MasterAuthor

	// SELECT * FROM master_authors
	if err := r.db.Find(&masterAuthor).Error; err != nil {
		return nil, err
	}

	return masterAuthor, nil
}

func (r *repository) StoreMasterAuthor(masterAuthor entities.MasterAuthor) (entities.MasterAuthor, error) {
	err := r.db.Create(&masterAuthor).Error

	if err != nil {
		return masterAuthor, err
	}

	return masterAuthor, nil
}

// func (r *repository) FindByID(ID int) (entities.MasterAuthor, error) {
// 	var masterAuthor entities.MasterAuthor

// 	err := r.db.Preload("MasterBook").Where("id = ?", ID).Find(&masterAuthor).Error
// 	if err != nil {
// 		return masterAuthor, err
// 	}

// 	return masterAuthor, nil
// }

func (r *repository) Update(ID string, input MasterAuthorInput) (entities.MasterAuthor, error) {
	var masterAuthor entities.MasterAuthor

	err := r.db.Preload("MasterBook").Where("id = ?", ID).Find(&masterAuthor).Error
	if err != nil {
		return masterAuthor, err
	}

	masterAuthor.Name = input.Name
	masterAuthor.UpdatedAt = time.Now()

	err = r.db.Updates(&masterAuthor).Error
	if err != nil {
		return masterAuthor, err
	}

	return masterAuthor, nil
}
