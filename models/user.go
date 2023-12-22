package models

type User struct {
	ID   int    `bun:"id,pk,autoincrement"`
	Name string `bun:"name"`
}
