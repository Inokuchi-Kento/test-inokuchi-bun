package main

import (
	"context"
	"database/sql"
	"log"
	"test-bun/models"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"

	_ "github.com/lib/pq"
)

func main() {
	// bunでpostgresに接続
	sqldb, err := sql.Open("postgres", "host=localhost port=5433 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db := bun.NewDB(sqldb, pgdialect.New())

	// // テーブル作成
	// if err := CreateTable(db); err != nil {
	// 	log.Fatal(err)
	// }
	SelectUserByID(db)
	// _ = InsertUser(db)

}

func CreateTable(db *bun.DB) error {
	// テーブル作成
	_, err := db.NewCreateTable().
		Model((*models.User)(nil)).
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// select user
func SelectUserByID(db *bun.DB) error {
	user := new(models.User)
	err := db.NewSelect().Model(user).Where("id = ?", 0).Scan(context.Background())
	if err != nil {
		return err
	}
	log.Println(user)
	return nil
}

// insert
func InsertUser(db *bun.DB) error {
	user := &models.User{
		Name: "User1",
	}

	// Insert user into the database.
	_, err := db.NewInsert().Model(user).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
