package services

import (
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/lucassimon/golang-encoder-video-mpegdash-service/internal/adapters/database"
	"github.com/lucassimon/golang-encoder-video-mpegdash-service/internal/application/repositories"
	"github.com/lucassimon/golang-encoder-video-mpegdash-service/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func prepare() (*entities.VideoEntity, repositories.VideoRepositoryDb) {
	db := database.NewDbTest()

	video := entities.NewVideoEntity()
	video.Id = uuid.NewV4().String()
	video.FilePath = "convite.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}

	return video, repo
}

func TestVideoService_Download(t *testing.T) {

	video, repo := prepare()

	videoService := NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	err := videoService.Download("medias-storage")
	require.Nil(t, err)
}

func TestVideoService_Fragment(t *testing.T) {

	video, repo := prepare()

	videoService := NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	err := videoService.Download("medias-storage")
	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)

	err = videoService.Encode()
	require.Nil(t, err)

	err = videoService.Finish()
	require.Nil(t, err)
}
