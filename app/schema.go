package main

type User struct {
	UserId uint `gorm:"primaryKey;"`
	Name   string
}

//type BookShelf struct {
//	BookShelfId uint `gorm:"primaryKey;autoIncrement:true;"`
//	OwnerId uint
//	Owner User `gorm:"foreignKey:OwnerId;"`
//	BookShelfName string
//}

type Book struct {
	BookId uint `gorm:"primaryKey;autoIncrement:true;"`
	//BookShelfId uint
	//BookShelf BookShelf `gorm:"foreignKey:BookShelfId"`
	OwnerId uint
	Owner   User `gorm:"foreignKey:OwnerId;"`
	ISBN13  uint
	Title   string
}

//type Borrowing struct {
//	BorrowingId uint `gorm:"primaryKey;autoIncrement:true;"`
//	BorrowerId uint
//	Borrower User `gorm:"foreignKey:BorrowerId;"`
//	LenderId uint
//	Lender User `gorm:"foreignKey:LenderId;"`
//}
