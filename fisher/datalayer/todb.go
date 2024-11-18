package datalayer

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"main/protobufs"
	"os"
)

var db *pgx.Conn

//type Pool struct {
//	FishID int
//	PoolID string
//	Value  float64
//}
//
//type Attraction struct {
//	FishID int
//	BaitID int
//	Value  float64
//}

type CatchValue struct {
	FishID int
	BaitID int
	Value  int
}

type Fish struct {
	ID          int
	Name        string
	Taxonomy    string
	Low_weight  float64
	High_weight float64
	Peak        float64
	Spread      float64
}

var catchValues []CatchValue

func InitDB() error {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return err
	}
	db = conn
	return nil
}

func SaveFish(fish *protobufs.Fish, userID string) {
	row, _ := db.Query(context.Background(),
		"INSERT INTO fishing_hell.fisher.caught(fishid, weight, whencaught, frompool, baitused, caughtby, owner) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		fish.Id, fish.Weight, fish.TimeCaught, fish.Pool, fish.BaitUsed, userID, userID)
	defer row.Close()
}

func GetCatchValues(poolID string, baitID int32) ([]CatchValue, error) {
	var catchValues []CatchValue

	// Parameterized query to prevent SQL injection
	query := `
        SELECT 
            catch.fishid, 
            attraction.baitid, 
            ROUND(
                CASE 
                    WHEN attraction.baitid = $1 
                    THEN catch.value * attraction.value 
                    ELSE catch.value 
                END
            ) AS catch_value
        FROM fisher.catch_value catch
        RIGHT JOIN fisher.attraction attraction 
        ON catch.fishid = attraction.fishid
        WHERE catch.poolid = $2;
    `

	rows, err := db.Query(context.Background(), query, baitID, poolID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var catchValue CatchValue
		err := rows.Scan(&catchValue.FishID, &catchValue.BaitID, &catchValue.Value)
		if err != nil {
			return nil, err
		}
		catchValues = append(catchValues, catchValue)
	}

	// No matches, return an empty slice (not nil)
	if len(catchValues) == 0 {
		log.Printf("No catch values found for poolID '%s' and baitID %d", poolID, baitID)
	}

	return catchValues, nil
}

func GetFish(fishID int) Fish {
	var fish Fish
	// Run the query
	row := db.QueryRow(context.Background(), "SELECT * FROM fishing_hell.fisher.fish WHERE id = $1", fishID)
	// Scan the row into all fields of the Fish struct, including ID
	err := row.Scan(&fish.ID, &fish.Name, &fish.Taxonomy, &fish.Low_weight, &fish.High_weight, &fish.Peak, &fish.Spread)
	if err != nil {
		log.Fatal(err)
	}
	return fish
}

//func GetPools() ([]Pool, error) {
//	var pools []Pool
//	rows, err := db.Query(context.Background(), "select fishID, poolID, value from fisher.catch_value")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var pool Pool
//		err := rows.Scan(&pool.FishID, &pool.PoolID, &pool.Value)
//		if err != nil {
//			return nil, err
//		}
//		pools = append(pools, pool)
//	}
//
//	if rows.Err() != nil {
//		return nil, rows.Err()
//	}
//
//	return pools, nil
//}
//
//func GetAttractions() ([]Attraction, error) {
//	var attractions []Attraction
//	rows, err := db.Query(context.Background(), "select fishID, baitID, value from fisher.attraction")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var attraction Attraction
//		err := rows.Scan(&attraction.FishID, &attraction.BaitID, &attraction.Value)
//		if err != nil {
//			return nil, err
//		}
//		attractions = append(attractions, attraction)
//	}
//
//	if rows.Err() != nil {
//		return nil, rows.Err()
//	}
//
//	return attractions, nil
//}

func CloseDB() {
	err := db.Close(context.Background())
	if err != nil {
		return
	}
}
