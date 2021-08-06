package repository

// import (
// 	"log"
// 	"todo/entities"

// 	"github.com/pkg/errors"
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// type GormInsert struct {
// 	db *gorm.DB
// }

// func NewGormInsert() GormInsert {
// 	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }
// 	// db.AutoMigrate(&entities.Task{})

// 	// dsn := "root:secret@tcp(127.0.0.1:3306)/gomysql?charset=utf8mb4&parseTime=True&loc=Local"
// 	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	db.AutoMigrate(&entities.Task{})

// 	return GormInsert{
// 		db: db,
// 	}
// }

// func (insert GormInsert) Insert(v interface{}) error {
// 	return errors.WithMessage(insert.db.Create(v).Error, "gorm insert")
// }

// func (insert GormInsert) TaskDone(id uint) error {
// 	return errors.WithMessage(insert.db.Model(&entities.Task{}).Where("id = ?", id).Update("done", true)).Error, "gorm task done")
// }

// func (insert GormInsert) GetAllTask() error {
// 	return errors.WithMessage(insert.db.Model(&entities.Task{}).Where("id = ?", id).Update("done", true)).Error, "gorm task done")
// }
