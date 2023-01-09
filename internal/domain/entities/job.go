package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type JobEntity struct {
	Id               string       `json:"job_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OutputBucketPath string       `json:"output_bucket_path" valid:"notnull"`
	Status           string       `json:"status" valid:"notnull"`
	Video            *VideoEntity `json:"video" valid:"-"`
	VideoID          string       `json:"-" valid:"-" gorm:"column:video_id;type:uuid;notnull"`
	Error            string       `valid:"-"`
	CreatedAt        time.Time    `json:"created_at" valid:"-"`
	UpdatedAt        time.Time    `json:"updated_at" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewJobEntity(output string, status string, video *VideoEntity) (*JobEntity, error) {
	job := JobEntity{
		OutputBucketPath: output,
		Status:           status,
		Video:            video,
	}

	job.prepare()

	err := job.Validate()

	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (job *JobEntity) prepare() {

	job.Id = uuid.NewV4().String()
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()

}

func (job *JobEntity) Validate() error {
	_, err := govalidator.ValidateStruct(job)

	if err != nil {
		return err
	}

	return nil
}
