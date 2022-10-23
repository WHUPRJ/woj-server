package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"
	"time"
)

var _ global.Repo = (*Repo)(nil)

type Repo struct {
	db  *gorm.DB
	log *zap.Logger
}

func (r *Repo) Get() interface{} {
	return r.db
}

func (r *Repo) Close() error {
	db, err := r.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (r *Repo) Setup(g *global.Global) {
	r.log = g.Log

	r.log.Info("Connecting to database...")

	logger := zapgorm2.New(r.log)
	logger.IgnoreRecordNotFoundError = true

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		g.Conf.Database.User,
		g.Conf.Database.Password,
		g.Conf.Database.Database,
		g.Conf.Database.Host,
		g.Conf.Database.Port)

	var err error
	r.db, err = gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
				TablePrefix:   g.Conf.Database.Prefix,
			},
			PrepareStmt: true,
			Logger:      logger,
		})

	if err != nil {
		r.log.Fatal("Failed to connect to database", zap.Error(err))
		return
	}

	db, err := r.checkAlive(3)
	if err != nil {
		r.log.Fatal("Database is not alive", zap.Error(err))
		return
	}

	db.SetMaxOpenConns(g.Conf.Database.MaxOpenConns)
	db.SetMaxIdleConns(g.Conf.Database.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(g.Conf.Database.ConnMaxLifetime) * time.Minute)

	r.migrateDatabase()
}

func (r *Repo) migrateDatabase() {
	r.log.Info("Auto Migrating database...")

	_ = r.db.AutoMigrate(&model.User{})
	_ = r.db.AutoMigrate(&model.Problem{})
	_ = r.db.AutoMigrate(&model.ProblemVersion{})
	_ = r.db.AutoMigrate(&model.Submission{})
	_ = r.db.AutoMigrate(&model.Status{})
}

// checkAlive deprecated
func (r *Repo) checkAlive(retry int) (*sql.DB, error) {
	if retry <= 0 {
		return nil, errors.New("all retries are used up. failed to connect to database")
	}

	db, err := r.db.DB()
	if err != nil {
		r.log.Warn("failed to get sql.DB instance", zap.Error(err))
		time.Sleep(5 * time.Second)
		return r.checkAlive(retry - 1)
	}

	err = db.Ping()
	if err != nil {
		r.log.Warn("failed to ping database", zap.Error(err))
		time.Sleep(5 * time.Second)
		return r.checkAlive(retry - 1)
	}

	r.log.Info("database connect established")
	return db, nil
}
