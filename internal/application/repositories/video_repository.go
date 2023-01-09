package repositories

import (
	"fmt"

	"github.com/lucassimon/golang-encoder-video-mpegdash-service/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Insert(video *entities.VideoEntity) (*entities.VideoEntity, error)
	Find(id string) (*entities.VideoEntity, error)
}

type VideoRepositoryDb struct {
	Db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepositoryDb {
	return &VideoRepositoryDb{Db: db}
}

func (repo VideoRepositoryDb) Insert(video *entities.VideoEntity) (*entities.VideoEntity, error) {

	if video.Id == "" {
		video.Id = uuid.NewV4().String()
	}

	err := repo.Db.Create(video).Error

	if err != nil {
		return nil, err
	}

	return video, nil

}

func (repo VideoRepositoryDb) Find(id string) (*entities.VideoEntity, error) {

	var video entities.VideoEntity
	repo.Db.Preload("Jobs").First(&video, "id = ?", id)

	if video.Id == "" {
		return nil, fmt.Errorf("video does not exist")
	}

	return &video, nil

}
