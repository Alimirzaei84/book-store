package models

import (
	"github.com/Alimirzaei84/book-store/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


 var db *gorm.DB

 type Book struct {
	 gorm.Model
	 Name string `json:"name"`
	 Author string `json:"author"`
	 Publication string `json:"publication"`
 }


 func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
	 }


 func (b *Book) CreateBook() *Book{
		db.NewRecord(b)
		db.Create(&b)
		return b
	}


 func GetAllBooks() []Book {
	 var books []Book
	 db.Find(&books)
	 return books
 }


 func GetBookById(id int64)(*Book, *gorm.DB) {
	var getBook Book
	db:=db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
 }


 func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}

func UpdateBook(book *Book) *Book {
	db.Save(&book)
	return book
}