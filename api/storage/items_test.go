package storage

import (
	"testing"

	"github.com/mvsestito/menu-api/api/storage/mock"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func TestGetAllItems(t *testing.T) {
	assert := assert.New(t)

	db := mock.MockDB()
	mock.ResetTables(db)
	mock.AddMockRestaurants(db)

	// should be empty
	items, err := GetAllItems(db, 1, "")
	assert.Nil(err)
	if err != nil {
		return
	}
	assert.Equal(0, len(items))

	// add items
	mock.AddMockItems(db)

	// get all for restaurant 1
	items, err = GetAllItems(db, 1, "")
	assert.Nil(err)
	if err != nil {
		return
	}
	assert.Equal(6, len(items))

	// get rest 2
	items, err = GetAllItems(db, 2, "")
	assert.Nil(err)
	if err != nil {
		return
	}
	assert.Equal(4, len(items))

	// get only sides
	items, err = GetAllItems(db, 1, "side")
	assert.Nil(err)
	if err != nil {
		return
	}
	assert.Equal(2, len(items))
	assert.Equal("french fries", items[0].Name)
	assert.Equal("salad", items[1].Name)
}
