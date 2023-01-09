package repositories

import (
	"fmt"

	"github.com/lucassimon/golang-encoder-video-mpegdash-service/internal/domain/entities"
	"gorm.io/gorm"
)

type JobRepository interface {
	Insert(job *entities.JobEntity) (*entities.JobEntity, error)
	Find(id string) (*entities.JobEntity, error)
	Update(job *entities.JobEntity) (*entities.JobEntity, error)
}

type JobRepositoryDb struct {
	Db *gorm.DB
}

func (repo JobRepositoryDb) Insert(job *entities.JobEntity) (*entities.JobEntity, error) {

	err := repo.Db.Create(job).Error

	if err != nil {
		return nil, err
	}

	return job, nil

}

func (repo JobRepositoryDb) Find(id string) (*entities.JobEntity, error) {

	var job entities.JobEntity
	repo.Db.Preload("Video").First(&job, "id = ?", id)

	if job.Id == "" {
		return nil, fmt.Errorf("job does not exist")
	}

	return &job, nil
}

func (repo JobRepositoryDb) Update(job *entities.JobEntity) (*entities.JobEntity, error) {
	err := repo.Db.Save(&job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}
