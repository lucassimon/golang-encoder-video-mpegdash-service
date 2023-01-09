package database

import (
	"log"
	"os"
	"time"

	"github.com/lucassimon/golang-encoder-video-mpegdash-service/internal/domain/entities"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "test"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.DsnTest = ":memory:"
	dbInstance.AutoMigrateDb = true
	dbInstance.Debug = true

	connection, err := dbInstance.Connect()

	if err != nil {
		log.Fatalf("Test db error: %v", err)
	}

	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	gormConfig := &gorm.Config{}

	// TODO: Discover how it works with loguru
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	if d.Debug {
		gormConfig = &gorm.Config{Logger: newLogger}
	}

	var err error

	if d.Env == "production" {
		d.Db, err = gorm.Open(postgres.Open(d.Dsn), gormConfig)
	} else {
		d.Db, err = gorm.Open(sqlite.Open(d.DsnTest), gormConfig)
	}

	if err != nil {
		return nil, err
	}

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&entities.VideoEntity{}, &entities.JobEntity{})
		// d.Db.Model(domain.Job{}).AddForeignKey("video_id", "videos (id)", "CASCADE", "CASCADE")
	}

	return d.Db, nil
}

// func (d *Database) Close() (*gorm.DB, error) {

// }
