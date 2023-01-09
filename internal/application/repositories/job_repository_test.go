package repositories

import (
	"testing"
	"time"

	"github.com/lucassimon/golang-encoder-video-mpegdash-service/internal/adapters/database"
	"github.com/lucassimon/golang-encoder-video-mpegdash-service/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	// defer db.Close()

	video := entities.NewVideoEntity()
	video.Id = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := entities.NewJobEntity("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob := JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	j, err := repoJob.Find(job.Id)
	require.NotEmpty(t, j.Id)
	require.Nil(t, err)
	require.Equal(t, j.Id, job.Id)
	require.Equal(t, j.VideoID, video.Id)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()

	video := entities.NewVideoEntity()
	video.Id = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := entities.NewJobEntity("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob := JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	job.Status = "Complete"

	repoJob.Update(job)

	j, err := repoJob.Find(job.Id)
	require.NotEmpty(t, j.Id)
	require.Nil(t, err)
	require.Equal(t, j.Status, job.Status)
}
