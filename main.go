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

	// テーブル作成
	// if err := CreateTable(db); err != nil {
	// 	log.Fatal(err)
	// }

	// InsertUser(db)
	// BulkInsertUser(db)
	// SelectUserByID(db)
	SelectAllUser(db)
	// DropTable(db)

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

// drop table
func DropTable(db *bun.DB) error {
	_, err := db.NewDropTable().
		Model((*models.User)(nil)).
		IfExists().
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// select all
func SelectAllUser(db *bun.DB) {
	var users []*models.User
	err := db.NewSelect().Model(&users).Scan(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// usersの中身を表示
	for _, user := range users {
		log.Println(user)
	}
}

// select user by ID
func SelectUserByID(db *bun.DB) {
	user := new(models.User)
	err := db.NewSelect().Model(user).Where("id = ?", 1).Scan(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user)
}

// insert
func InsertUser(db *bun.DB) {
	user := &models.User{
		ID:   1,
		Name: "User1",
	}

	_, err := db.NewInsert().Model(user).Exec(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

// bulk insert
func BulkInsertUser(db *bun.DB) {
	users := []*models.User{
		{
			ID:   1001,
			Name: "User1",
		},
		{
			ID:   1002,
			Name: "User2",
		},
	}

	_, err := db.NewInsert().Model(&users).Exec(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
