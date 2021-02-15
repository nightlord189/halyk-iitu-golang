package main

import (
	"fmt"
	"iitu/db/gorm/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Close db
func close(db *gorm.DB) {
	DB, err := db.DB()
	if err != nil {
		panic(err)
	}
	DB.Close()
	fmt.Println("connection closed")
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(model.User{}, model.UserState{})
	fmt.Println("migrated")
}

func main() {
	//Connect
	dsn := "host=localhost user=postgres password=123456 dbname=iitu2 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	defer close(db)
	fmt.Println("connection established")
	//Automigrate
	migrate(db)

	//Delete
	db.Where("id > ?", 1).Delete(model.User{})

	//Create
	user := model.User{
		ID:        15,
		FirstName: "Joe",
		LastName:  "Biden",
		Login:     "user1",
	}
	db.Create(&user)

	//Read
	var readUser model.User
	db.Where("id = ?", 15).First(&readUser)
	fmt.Printf("read: %v\n", readUser)

	//read many
	user2 := model.User{
		ID:        25,
		FirstName: "Donald",
		LastName:  "Trump",
		Login:     "user123",
	}
	db.Create(&user2)
	var users []model.User
	db.Where("id > ?", 0).Find(&users)
	fmt.Printf("read array: %v\n", users)

	//Update
	user.Birth = "1942"
	user.Login = "user83737"
	db.Save(&user)

	db.Where("id = ?", 15).First(&readUser)
	fmt.Printf("read: %v\n", readUser)

	//Batch update by fields
	fieldsToUpdate := map[string]interface{}{"firstname": "Ivan", "lastname": "Ivanov"}
	db.Model(&user).Where("id > ?", 1).Updates(fieldsToUpdate)

	db.Where("id = ?", 15).First(&readUser)
	fmt.Printf("read: %v\n", readUser)

	//Exec
	//db.Exec("truncate table \"user\"")
}
