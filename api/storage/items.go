package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/lib/pq"
)

var (
	// on any state changing exec against db cache should be flushed
	// ensure data consistency
	modifiersCache             = make(map[string][]Item)
	cacheLock      *sync.Mutex = &sync.Mutex{}
)

type Item struct {
	ID           int    `json:"id"`
	RestaurantID int    `json:"restaurant_id"`
	Name         string `json:"name"`
	ItemType     string `json:"item_type"`
	Desc         string `json:"description"`
	Modifiers    []Item `json:"modifiers"`
}

func GetAllItems(restaurantID int, itemType string) ([]Item, error) {
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

	// protect against sql injection, dont string fmt, use driver built in escaping
	if itemType != "" {
		q += " WHERE item_type = $2"
		rows, err = DB.Query(q, restaurantID, itemType)
	} else {
		rows, err = DB.Query(q, restaurantID)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no items found")
		} else {
			log.Println("error querying items: ", err)
			return nil, err
		}
	}

	defer rows.Close()
	for rows.Next() {
		var (
			item     Item
			modTypes []string
		)

		err := rows.Scan(&item.ID, &item.RestaurantID, &item.Name, &item.ItemType, &item.Desc, pq.Array(&modTypes))
		if err != nil {
			log.Println("error scanning items: ", err)
			return nil, err
		}

		// get modifiers
		for _, t := range modTypes {
			var err error

			// check cache first
			cacheLock.Lock()
			modItems, ok := modifiersCache[t]
			cacheLock.Unlock()
			if !ok { // query db
				modItems, err = GetAllItems(restaurantID, t)
				if err != nil {
					log.Printf("error getting modifiers: itemtype: %s, err: %s\n", t, err)
					return nil, err
				}

				cacheLock.Lock()
				modifiersCache[t] = modItems
				cacheLock.Unlock()
			}

			for _, mod := range modItems {
				item.Modifiers = append(item.Modifiers, mod)
			}
		}

		items = append(items, item)
	}

	return items, nil
}

//func GetItem(restaurantID int, itemName string) (Item, error) {
//}
