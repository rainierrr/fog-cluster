package fogNode

import (
	"log"

	"gorm.io/gorm"
)

type FogNodeRepository struct {
}

type createParams struct {
	Name  string `json:"name"`
	Token string `json:"token"`
	Tag   string `json:"tag"`
	Ip    string `json:"ip"`
}

func (r *FogNodeRepository) FindAll(db *gorm.DB) ([]FogNode, error) {
	var fogNodes []FogNode
	if err := db.Find(&fogNodes).Error; err != nil {
		log.Println(fogNodes)
		return fogNodes, err
	}

	return fogNodes, nil
}

func (r *FogNodeRepository) Create(db *gorm.DB, createParams *createParams) error {
	fogNode := FogNode{
		Name:  createParams.Name,
		Token: createParams.Token,
		Tag:   createParams.Tag,
		Ip:    createParams.Ip,
	}
	if err := db.Create(&fogNode); err != nil {
		return err.Error
	}
	return nil
}
