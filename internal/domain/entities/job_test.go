package entities

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
	video := NewVideoEntity()
	video.Id = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, err := NewJobEntity("path", "Converted", video)
	require.NotNil(t, job)
	require.Nil(t, err)
}
