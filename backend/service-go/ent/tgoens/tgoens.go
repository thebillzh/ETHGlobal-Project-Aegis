// Code generated by ent, DO NOT EDIT.

package tgoens

import (
	"time"
)

const (
	// Label holds the string label denoting the tgoens type in the database.
	Label = "tgo_ens"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldWalletPub holds the string denoting the wallet_pub field in the database.
	FieldWalletPub = "wallet_pub"
	// FieldEns holds the string denoting the ens field in the database.
	FieldEns = "ens"
	// FieldMtime holds the string denoting the mtime field in the database.
	FieldMtime = "mtime"
	// FieldCtime holds the string denoting the ctime field in the database.
	FieldCtime = "ctime"
	// Table holds the table name of the tgoens in the database.
	Table = "t_go_ens"
)

// Columns holds all SQL columns for tgoens fields.
var Columns = []string{
	FieldID,
	FieldWalletPub,
	FieldEns,
	FieldMtime,
	FieldCtime,
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
	// DefaultEns holds the default value on creation for the "ens" field.
	DefaultEns string
	// DefaultMtime holds the default value on creation for the "mtime" field.
	DefaultMtime time.Time
	// UpdateDefaultMtime holds the default value on update for the "mtime" field.
	UpdateDefaultMtime func() time.Time
	// DefaultCtime holds the default value on creation for the "ctime" field.
	DefaultCtime time.Time
)
