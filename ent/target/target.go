// Code generated by ent, DO NOT EDIT.

package target

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the target type in the database.
	Label = "target"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldIntervalSeconds holds the string denoting the interval_seconds field in the database.
	FieldIntervalSeconds = "interval_seconds"
	// FieldTimeoutSeconds holds the string denoting the timeout_seconds field in the database.
	FieldTimeoutSeconds = "timeout_seconds"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldHostname holds the string denoting the hostname field in the database.
	FieldHostname = "hostname"
	// FieldPort holds the string denoting the port field in the database.
	FieldPort = "port"
	// Table holds the table name of the target in the database.
	Table = "targets"
)

// Columns holds all SQL columns for target fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldType,
	FieldIntervalSeconds,
	FieldTimeoutSeconds,
	FieldURL,
	FieldHostname,
	FieldPort,
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
	// IntervalSecondsValidator is a validator for the "interval_seconds" field. It is called by the builders before save.
	IntervalSecondsValidator func(uint32) error
	// TimeoutSecondsValidator is a validator for the "timeout_seconds" field. It is called by the builders before save.
	TimeoutSecondsValidator func(uint32) error
	// PortValidator is a validator for the "port" field. It is called by the builders before save.
	PortValidator func(uint32) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeTYPE_HTTP Type = "TYPE_HTTP"
	TypeTYPE_TCP  Type = "TYPE_TCP"
	TypeTYPE_PING Type = "TYPE_PING"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeTYPE_HTTP, TypeTYPE_TCP, TypeTYPE_PING:
		return nil
	default:
		return fmt.Errorf("target: invalid enum value for type field: %q", _type)
	}
}
