package domain

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestNewVideoIdIsNotAUuid(t *testing.T) {
	video := NewVideo()

	video.ID = "abc"
	video.ResourceID = "a"
	video.FilePath = "path"

	err := video.Validate()
	require.Error(t, err)
}

func TestNewVideoValidation(t *testing.T) {
	video := NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourceID = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}
