package model

import (
	"fmt"
	"../env"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
    "../configuration"
)

var (
	_db *gorm.DB
)

func beginTx(db *gorm.DB) func() {
    tx := db.Begin()
    tx.LogMode(true)
    return func() {
        if err := recover(); err != nil {
            tx.Rollback()
            panic(err)
        } else {
            tx.Commit()
        }
    }
}

func init() {
    config := configuration.Get().Database
    switch (config.Kind) {
    case configuration.DbKind_SQLite:
        db, err := gorm.Open("sqlite3", config.SQLite.DatabaseFilePath)
        if err != nil {
            panic(err)
        }
        _db = &db
        initializeDatabase(_db)
    case configuration.DbKind_PostgreSQL:
        cs := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
            env.ResolveEnv(config.PostgreSQL.Host),
            env.ResolveEnv(config.PostgreSQL.Port),
            env.ResolveEnv(config.PostgreSQL.User),
            env.ResolveEnv(config.PostgreSQL.Database),
            env.ResolveEnv(config.PostgreSQL.Password),
        )
        fmt.Println("★", cs)
        db ,err := gorm.Open("postgres", cs)
        if err != nil {
            panic(err)
        }
        _db = &db
        initializeDatabase(_db)
    }
}

func initializeDatabase(db *gorm.DB) {
    db.SingularTable(true) // 自動生成されるテーブル名を単数形にする.

    db.CreateTable(&Property{})
    db.CreateTable(&Room{})
//    db.AutoMigrate(&Family{}, &FamilyCredential{}) // 使いこなせないので当面自動マイグレーションはoff.

    db.LogMode(true)
    defer beginTx(db)()
}

type NotFound interface {
	// nodef
}
type notFoundImpl struct {
	// nodef
}

func NewNotFound() NotFound {
	return notFoundImpl{}
}

type InvalidValue interface {
	GetDescription() string
}

type invalidValue struct {
	description string
}

func (e *invalidValue) GetDescription() string {
	return e.description
}

func NewInvalidValue(description string) InvalidValue {
	return &invalidValue{description: description}
}

type Duplicate interface {
    GetDescription() string
}

type duplicateImpl struct {
    description string
}

func (e *duplicateImpl) GetDescription() string {
    return e.description
}

func NewDuplicate(description string) Duplicate {
    return &duplicateImpl { description: description }
}

func mustInsert(db *gorm.DB, entity interface{}) {
	if err := db.Create(entity).Error; err != nil {
		panic(err)
	}
}
