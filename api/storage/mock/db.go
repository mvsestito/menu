package mock

import (
	"database/sql"
	"log"
)

// AddMockRestaurants adds mock restaurant data to the db
func AddMockRestaurants(db *sql.DB) {
	var err error

	// add restaurants
	q := `INSERT INTO restaurants (id, name, street, city, state, zip)
			VALUES(1, 'restone', '1 food way', 'chicago', 'illinois', '12345')
			VALUES(2, 'resttwo', '2 food way', 'boston', 'massachusetts', '02116')`

	_, err = db.Exec(q)
	if err != nil {
		log.Fatal("error adding mock restaurants to database: ", err)
	}
}

// AddMockItems adds mock item data to the db
func AddMockItems(db *sql.DB) {
	var err error

	// add toppings
	q := `INSERT INTO items (id, restaurant_id, name, item_type, description)
			VALUES(1, 1, 'cheese', 'topping', 'cheddar cheese')
			VALUES(2, 1, 'ketchup', 'topping', 'regular ketchup')
			VALUES(3, 2, 'mustard', 'topping', 'romaine lettuce')
			VALUES(4, 2, 'olive oil', 'topping', 'olive oil')`

	_, err = db.Exec(q)
	if err != nil {
		log.Fatal("error adding mock items to database: ", err)
	}

	// add sides
	q = `INSERT INTO items (id, restaurant_id, name, item_type, description, modifiers)
			VALUES(5, 1, 'french fries', 'side', 'fried potatoes', '{topping}')
			VALUES(6, 1, 'salad', 'side', 'house salad', '{topping}')
			VALUES(7, 2, 'biscuit', 'side', 'fluffy wheat bread', '{topping}')`

	_, err = db.Exec(q)
	if err != nil {
		log.Fatal("error adding mock items to database: ", err)
	}

	// add entrees
	q = `INSERT INTO items (id, restaurant_id, name, item_type, description, modifiers)
			VALUES(8, 1, 'hamburger', 'entree', 'beef patty with wheat bun', '{topping, side}'),
			VALUES(9, 1, 'fried chicken', 'entree', 'chicken battered and fried', '{topping, side}'),
			VALUES(10, 2, 'salmon', 'entree', 'atlantic salmon', '{topping, side}')`

	_, err = db.Exec(q)
	if err != nil {
		log.Fatal("error adding mock items to database: ", err)
	}
}
