package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// User Пользователь
type User struct {
	ID        int    `json:"id"`
	LastName  string `json:"LastName"`
	FirstName string `json:"FirstName"`
	Birth     string `json:"Birth"`
	Login     string `json:"Login"`
	Password  string `json:"Password"`
}

func insert(db *sql.DB, user User) error {
	_, err := db.Exec(
		"insert into \"user\" (id, lastname, firstname) values ($1, $2, $3)",
		user.ID, user.LastName, user.FirstName,
	)
	return err
}

func update(db *sql.DB, user User) error {
	_, err := db.Exec(
		"update \"user\" set lastname = $2, firstname = $3 where id = $1",
		user.ID, user.LastName, user.FirstName,
	)
	return err
}

func read(db *sql.DB, id int) (User, error) {
	rows, err := db.Query("select id, lastname, firstname from \"user\" u where u.id = $1", id)
	if err != nil {
		panic(err)
	}
	var user User
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.LastName, &user.FirstName)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return user, nil
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123456 dbname=iitu1 sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection Established")
	}
	defer db.Close()
	db.Exec("delete from \"user\"")
	if err != nil {
		panic(err)
	}
	newUser := User{
		ID:        15,
		LastName:  "Biden",
		FirstName: "Joe",
		Birth:     "1942",
	}
	err = insert(db, newUser)
	if err != nil {
		panic(err)
	}
	userRead, err := read(db, 15)
	if err != nil {
		panic(err)
	}
	fmt.Println(userRead)
	userRead.LastName = "Ivanov"
	update(db, userRead)
	userRead, _ = read(db, 15)
	fmt.Println(userRead)
}
