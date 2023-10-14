package main

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v4"
)

const (
	dbDSN = "host=localhost port=54321 dbname=auth user=auth password=pass sslmode=disable"
)

func main() {
	ctx := context.Background()

	con, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer con.Close(ctx)

	pass := gofakeit.Password(true, true, true, true, false, 6)

	res, err := con.Exec(ctx, "INSERT INTO users (name, email, role, password, password_confirm) VALUES ($1, $2, $3, $4, $5)",
		gofakeit.Name(), gofakeit.Email(), 0, pass, pass)
	if err != nil {
		log.Printf("failed to insert user: %v", err)
	}

	log.Printf("inserted %d rows", res.RowsAffected())

	rows, err := con.Query(ctx, "SELECT id, name, email, role, password FROM users")
	if err != nil {
		log.Fatalf("failed to select users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, email, password string
		var role int

		err = rows.Scan(&id, &name, &email, &role, &password)
		if err != nil {
			log.Fatalf("failed to scan user: %v", err)
		}

		log.Printf("id: %d, name: %s, email: %s, role: %v, password: %v\n", id, name, email, role, password)
	}
}
