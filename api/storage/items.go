package storage

import (
	"database/sql"
	"fmt"
	"log"
)

type Item struct {
	ID           int
	RestaurantID int
	Name         string
	ItemType     string
	Desc         string
	Modifiers    []Item
}

func GetAllItems(db *sql.DB, restaurantID int, itemType string) ([]Item, error) {
	var (
		err   error
		rows  *sql.Rows
		items []Item
	)

	q := `
		SELECT i.id, i.restaurant_id, i.name, i.item_type, i.description, i.modifiers
		FROM items i
		JOIN restaurants r on i.restaurant_id = r.id AND r.id = $1
		`

	// protext against sql injection, dont string fmt, use driver built in escaping
	if itemType != "" {
		q += " WHERE item_type = $2"
		rows, err = db.Query(q, restaurantID, itemType)
	} else {
		rows, err = db.Query(q, restaurantID)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no items found")
		} else {
			log.Println("error querying items: ", err)
			return nil, err
		}
	}

	// store queried modifiers in local cache to avoid dupe db queries
	modifiersCache := make(map[string][]Item)

	defer rows.Close()
	for rows.Next() {
		var (
			item     Item
			modTypes []string
		)

		err := rows.Scan(&item.ID, &item.RestaurantID, &item.Name, &item.ItemType, &item.Desc, &modTypes)
		if err != nil {
			log.Println("error scanning items: ", err)
			return nil, err
		}

		// get modifiers
		for _, t := range modTypes {
			var (
				modItems []Item
				ok       bool
				err      error
			)

			// check cache first
			modItems, ok = modifiersCache[t]
			if !ok { // query db
				modItems, err = GetAllItems(db, restaurantID, t)
				if err != nil {
					log.Printf("error getting modifiers: itemtype: %s, err: %s\n", t, err)
					return nil, err
				}

				modifiersCache[t] = modItems
			}

			for _, mod := range modItems {
				item.Modifiers = append(item.Modifiers, mod)
			}
		}
	}

	return items, nil
}

//func GetItem(db *sql.DB, restaurantID int, itemName string) (Item, error) {
//}
