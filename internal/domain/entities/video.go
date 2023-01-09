package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type VideoEntity struct {
	Id         string       `json:"encoded_video_folder" valid:"uuid" gorm:"type:uuid;primary_key"`
	ResourceID string       `json:"resource_id" valid:"notnull" gorm:"type:varchar(255)"`
	FilePath   string       `json:"file_path" valid:"notnull" gorm:"type:varchar(255)"`
	Jobs       []*JobEntity `json:"-" valid:"-" gorm:"ForeignKey:VideoID"`
	CreatedAt  time.Time    `json:"-" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideoEntity() *VideoEntity {
	return &VideoEntity{}
}

func (video *VideoEntity) Validate() error {

	_, err := govalidator.ValidateStruct(video)

	if err != nil {
		return err
	}

	return nil
}
