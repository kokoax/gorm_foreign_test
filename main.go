package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Super struct {
	ID       int64 `gorm:"primary_key"`
	Children []Child
	Code     string
	Price    uint
}

type Child struct {
	ID      int64 `gorm:"primary_key"`
	Super   Super `gorm:"foreignkey:SuperID"`
	Super   Super
	SuperID int64
	Title   string
	Value   int64
}

func connectDB() *gorm.DB {
	DBMS := "mysql"
	CONNECT := "root:@tcp(localhost:3306)/gorm_test"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	db.LogMode(true)
	db.AutoMigrate(&Super{}, &Child{})

	return db
}

func create(db *gorm.DB) {
	sc1 := Child{
		Title: "Child Title1",
		Value: 100,
	}
	sc2 := Child{
		Title: "Child Title2",
		Value: 200,
	}
	sc3 := Child{
		Title: "Child Title3",
		Value: 300,
	}

	ss1 := Super{
		Children: []Child{sc1, sc2},
		Code:     "Super Code1",
		Price:    1000,
	}
	ss2 := Super{
		Children: []Child{sc3},
		Code:     "Super Code2",
		Price:    2000,
	}

	db.Create(&ss1)
	db.Create(&ss2)
	db.Create(&sc1)
	db.Create(&sc2)
	db.Create(&sc3)
}

func main() {
	db := connectDB()
	defer db.Close()

	// create(db)

	// super, child := Super{}, Child{}
	// child := Child{}
	// db.First(&child).Related(&child.Super)

	super := Super{}
	db.First(&super, "id = ?", 2).Related(&super.Children)
	fmt.Println("{:?}", super)
	// child := Child{}
	// db.First(&child, "id = ?", 3).Related(&child.Super)
	// fmt.Println("{:?}", child)
}
