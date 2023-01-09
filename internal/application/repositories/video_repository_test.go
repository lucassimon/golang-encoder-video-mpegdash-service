package repositories

import (
	"testing"
	"time"

	"github.com/lucassimon/golang-encoder-video-mpegdash-service/internal/adapters/database"
	"github.com/lucassimon/golang-encoder-video-mpegdash-service/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()

	video := entities.NewVideoEntity()
	video.Id = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := VideoRepositoryDb{Db: db}
	repo.Insert(video)

	v, err := repo.Find(video.Id)

	require.NotEmpty(t, v.Id)
	require.Nil(t, err)
	require.Equal(t, v.Id, video.Id)
}
