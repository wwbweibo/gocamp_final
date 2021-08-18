// Code generated by entc, DO NOT EDIT.

package beer

import (
	"time"
)

const (
	// Label holds the string label denoting the beer type in the database.
	Label = "beer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldImages holds the string denoting the images field in the database.
	FieldImages = "images"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the beer in the database.
	Table = "beers"
)

// Columns holds all SQL columns for beer fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldCount,
	FieldPrice,
	FieldImages,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)
