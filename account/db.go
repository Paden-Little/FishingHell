package main

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
	"main/protobufs"
	"os"
)

var db *pgx.Conn

func CreateAccount(username string, email string, password string) error {
	row, err := db.Query(context.Background(), "INSERT INTO fishing_hell.account.user(username, email, password, cash) VALUES ($1, $2, $3, $4)", username, email, password, 100)
	defer row.Close()
	return err
}

func Login(account *protobufs.Account) {
	// Query the user account details
	row := db.QueryRow(context.Background(),
		"SELECT id, cash, username, email, password FROM fishing_hell.account.user WHERE username = $1 AND password = $2",
		account.Username, account.Password)

	// Scan the results into the account struct
	err := row.Scan(&account.Id, &account.Cash, &account.Username, &account.Email, &account.Password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Printf("None found")
			account.Username = ""
			account.Password = ""
			return
		}
		log.Fatal(err)
	}

	// Query poles owned by the user
	rows, err := db.Query(context.Background(),
		"SELECT p.id AS pole_id, p.name AS pole_name, p.ticks, p.price FROM fishing_hell.shop.poleOwned po JOIN fishing_hell.shop.pole p ON po.poleID = p.id WHERE po.userID = $1",
		account.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Populate the poles into the account's poles list
	for rows.Next() {
		var pole protobufs.Pole
		err := rows.Scan(&pole.Id, &pole.Name, &pole.Ticks, &pole.Price)
		if err != nil {
			log.Fatal(err)
		}
		account.Poles = append(account.Poles, &pole)
	}

	// Handle any errors from row iteration
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Query baits owned by the user
	rows, err = db.Query(context.Background(),
		"SELECT b.id AS bait_id, b.name, b.value, b.uses, bo.owned FROM fishing_hell.shop.baitOwned bo JOIN fishing_hell.shop.bait b ON bo.baitID = b.id WHERE bo.userID = $1",
		account.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Populate the baits into the account's baits list
	for rows.Next() {
		var bait protobufs.Bait
		err := rows.Scan(&bait.Id, &bait.Name, &bait.Value, &bait.Uses, &bait.Owned)
		if err != nil {
			log.Fatal(err)
		}
		account.Baits = append(account.Baits, &bait)
	}

	// Query fish owned by the user
	rows, err = db.Query(context.Background(),
		`SELECT c.id AS fish_id, f.name, c.weight, c.fromPool, c.frozen 
         FROM fishing_hell.fisher.caught c 
         JOIN fishing_hell.fisher.fish f ON c.fishID = f.id 
         WHERE c.owner = $1`, account.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Populate the fish into the account's fish list
	for rows.Next() {
		var fish protobufs.FishCaught
		err := rows.Scan(&fish.Id, &fish.Name, &fish.Weight, &fish.FromPool, &fish.Frozen)
		if err != nil {
			log.Fatal(err)
		}
		account.Fish = append(account.Fish, &fish)
	}

	// Handle any errors from row iteration
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

//
//func GetAccountInformation(id int) {
//
//}
//
//func DecrementBait(userID int, bait int) {
//
//}

func InitDB() error {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return err
	}
	db = conn
	return nil
}

func GetAccountDetails(username string, password string) {
	account := protobufs.Account{}
	row, err := db.Query(context.Background(), "")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	row.Scan(&account)
}
