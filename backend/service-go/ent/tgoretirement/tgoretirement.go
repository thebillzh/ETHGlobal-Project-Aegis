// Code generated by ent, DO NOT EDIT.

package tgoretirement

import (
	"time"
)

const (
	// Label holds the string label denoting the tgoretirement type in the database.
	Label = "tgo_retirement"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreationTx holds the string denoting the creation_tx field in the database.
	FieldCreationTx = "creation_tx"
	// FieldCreatorAddress holds the string denoting the creator_address field in the database.
	FieldCreatorAddress = "creator_address"
	// FieldBeneficiaryAddress holds the string denoting the beneficiary_address field in the database.
	FieldBeneficiaryAddress = "beneficiary_address"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldTokenAddress holds the string denoting the token_address field in the database.
	FieldTokenAddress = "token_address"
	// FieldTokenName holds the string denoting the token_name field in the database.
	FieldTokenName = "token_name"
	// FieldTokenType holds the string denoting the token_type field in the database.
	FieldTokenType = "token_type"
	// FieldRetirementMessage holds the string denoting the retirement_message field in the database.
	FieldRetirementMessage = "retirement_message"
	// FieldRetirementTime holds the string denoting the retirement_time field in the database.
	FieldRetirementTime = "retirement_time"
	// FieldMtime holds the string denoting the mtime field in the database.
	FieldMtime = "mtime"
	// FieldCtime holds the string denoting the ctime field in the database.
	FieldCtime = "ctime"
	// Table holds the table name of the tgoretirement in the database.
	Table = "t_go_retirements"
)

// Columns holds all SQL columns for tgoretirement fields.
var Columns = []string{
	FieldID,
	FieldCreationTx,
	FieldCreatorAddress,
	FieldBeneficiaryAddress,
	FieldAmount,
	FieldTokenAddress,
	FieldTokenName,
	FieldTokenType,
	FieldRetirementMessage,
	FieldRetirementTime,
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
	// DefaultCreationTx holds the default value on creation for the "creation_tx" field.
	DefaultCreationTx string
	// DefaultCreatorAddress holds the default value on creation for the "creator_address" field.
	DefaultCreatorAddress string
	// DefaultBeneficiaryAddress holds the default value on creation for the "beneficiary_address" field.
	DefaultBeneficiaryAddress string
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount float64
	// DefaultTokenAddress holds the default value on creation for the "token_address" field.
	DefaultTokenAddress string
	// DefaultTokenName holds the default value on creation for the "token_name" field.
	DefaultTokenName string
	// DefaultTokenType holds the default value on creation for the "token_type" field.
	DefaultTokenType string
	// DefaultRetirementMessage holds the default value on creation for the "retirement_message" field.
	DefaultRetirementMessage string
	// DefaultMtime holds the default value on creation for the "mtime" field.
	DefaultMtime time.Time
	// UpdateDefaultMtime holds the default value on update for the "mtime" field.
	UpdateDefaultMtime func() time.Time
	// DefaultCtime holds the default value on creation for the "ctime" field.
	DefaultCtime time.Time
)
