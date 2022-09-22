// Code generated by ent, DO NOT EDIT.

package atasklog

import (
	"time"
)

const (
	// Label holds the string label denoting the atasklog type in the database.
	Label = "atask_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldQuestID holds the string denoting the quest_id field in the database.
	FieldQuestID = "quest_id"
	// FieldTaskID holds the string denoting the task_id field in the database.
	FieldTaskID = "task_id"
	// FieldMid holds the string denoting the mid field in the database.
	FieldMid = "mid"
	// FieldMeta holds the string denoting the meta field in the database.
	FieldMeta = "meta"
	// FieldMtime holds the string denoting the mtime field in the database.
	FieldMtime = "mtime"
	// FieldCtime holds the string denoting the ctime field in the database.
	FieldCtime = "ctime"
	// Table holds the table name of the atasklog in the database.
	Table = "a_task_logs"
)

// Columns holds all SQL columns for atasklog fields.
var Columns = []string{
	FieldID,
	FieldQuestID,
	FieldTaskID,
	FieldMid,
	FieldMeta,
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
	// DefaultQuestID holds the default value on creation for the "quest_id" field.
	DefaultQuestID uint64
	// DefaultTaskID holds the default value on creation for the "task_id" field.
	DefaultTaskID uint64
	// DefaultMid holds the default value on creation for the "mid" field.
	DefaultMid uint64
	// DefaultMeta holds the default value on creation for the "meta" field.
	DefaultMeta string
	// DefaultMtime holds the default value on creation for the "mtime" field.
	DefaultMtime time.Time
	// UpdateDefaultMtime holds the default value on update for the "mtime" field.
	UpdateDefaultMtime func() time.Time
	// DefaultCtime holds the default value on creation for the "ctime" field.
	DefaultCtime time.Time
)
