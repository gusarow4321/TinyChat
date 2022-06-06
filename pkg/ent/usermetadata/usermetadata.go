// Code generated by entc, DO NOT EDIT.

package usermetadata

const (
	// Label holds the string label denoting the usermetadata type in the database.
	Label = "user_metadata"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the userid field in the database.
	FieldUserID = "user_id"
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the usermetadata in the database.
	Table = "user_metadata"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_metadata"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for usermetadata fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldColor,
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
