package proxysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"

	"time"
)

func ProxySQLDB() {
	//db, err := sql.Open("mysql", "/mysql")
	//if err != nil {
	//panic(err)
	//}

	start := time.Now()
	db, err := sql.Open("mysql", "stnduser:stnduser@tcp(127.0.0.1:16033)/testdbrep")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Welcome to week1 - Using ProxySQL for routing to DBs")
	// Users table exists on both shards (DB servers)
	// Range-based sharding is done to divide data
	// All records with username starting with greater than 'm' below to second shard

	// Other sharding options -
	// - Hash-based - Using a hash function
	// - Geo-based - Using location info
	// - Directory-based - Uses a look-up table to find the appropriate shard

	username := "johndoe" // example username
	var (
		userID int
		email  string
		//createdAt string // assuming datetime can be scanned as a string for simplicity
	)

	query := "SELECT user_id, email FROM Users WHERE username = ?"
	row := db.QueryRow(query, username)

	// Scan for the values from the row
	if err := row.Scan(&userID, &email); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No results found.")
			return
		}
		log.Fatal(err)
	}

	fmt.Printf("User ID: %d, Email: %s", userID, email)

	duration := time.Since(start)
	fmt.Println("Time spent", duration)
}
