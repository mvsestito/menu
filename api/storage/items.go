package storage

import "database/sql"

type Item struct {
	ID           int
	RestaurantID int
	Name         string
	ItemType     string
	Desc         string
	Modifiers    []Item
}

func GetAllItems(db *sql.DB, restaurantID int, itemType string) ([]Item, error) {
	return items, nil
}

//func GetItem(db *sql.DB, restaurantID int, itemName string) (Item, error) {
//}
