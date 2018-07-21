package mock

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"
)

func MockDB() *sql.DB {
	dbstr := "host=%s dbname=test port=5432 user=postgres sslmode=disable"
	cmd := exec.Command("uname")
	if stdout, _ := cmd.Output(); string(stdout[:len(stdout)-1]) == "Darwin" { // running on local mac
		dbstr = fmt.Sprintf(dbstr, "localhost")
	} else { // inside docker bridge network container
		dbstr = fmt.Sprintf(dbstr, "db")
	}

	db, err := sql.Open("postgres", dbstr)
	if err != nil {
		log.Fatal("error opening database connection: ", err)
	}

	return db
}

// ResetTables truncates all tables in test db
func ResetTables(db *sql.DB) {
	for _, table := range []string{"restaurants", "items"} {
		_, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", table))
		if err != nil {
			log.Fatal("error resetting table: ", table, " err: ", err)
		}
	}
}

// AddMockRestaurants adds mock restaurant data to the db
func AddMockRestaurants(db *sql.DB) {
	var err error

	// add restaurants
	q := `INSERT INTO restaurants (id, name, street, city, state, zip) VALUES
			(1, 'restone', '1 food way', 'chicago', 'illinois', '12345'),
			(2, 'resttwo', '2 food way', 'boston', 'massachusetts', '02116')`

	_, err = db.Exec(q)
	if err != nil {
		log.Fatal("error adding mock restaurants to database: ", err)
	}
}

// AddMockItems adds mock item data to the db
func AddMockItems(db *sql.DB) {
	var err error

	// add toppings
	q := `INSERT INTO items (id, restaurant_id, name, item_type, description) VALUES
			(1, 1, 'cheese', 'topping', 'cheddar cheese'),
			(2, 1, 'ketchup', 'topping', 'regular ketchup'),
			(3, 2, 'mustard', 'topping', 'romaine lettuce'),
			(4, 2, 'olive oil', 'topping', 'olive oil')`

	_, err = db.Exec(q)
	if err != nil {
		log.Fatal("error adding mock items to database: ", err)
	}

	// add sides
	q = `INSERT INTO items (id, restaurant_id, name, item_type, description, modifiers) VALUES
			(5, 1, 'french fries', 'side', 'fried potatoes', '{topping}'),
			(6, 1, 'salad', 'side', 'house salad', '{topping}'),
			(7, 2, 'biscuit', 'side', 'fluffy wheat bread', '{topping}')`

	_, err = db.Exec(q)
	if err != nil {
		log.Fatal("error adding mock items to database: ", err)
	}

	// add entrees
	q = `INSERT INTO items (id, restaurant_id, name, item_type, description, modifiers) VALUES
			(8, 1, 'hamburger', 'entree', 'beef patty with wheat bun', '{topping, side}'),
			(9, 1, 'fried chicken', 'entree', 'chicken battered and fried', '{topping, side}'),
			(10, 2, 'salmon', 'entree', 'atlantic salmon', '{topping, side}')`

	_, err = db.Exec(q)
	if err != nil {
		log.Fatal("error adding mock items to database: ", err)
	}
}
