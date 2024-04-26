// Code generated by ent, DO NOT EDIT.

package file

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the file type in the database.
	Label = "file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFilePath holds the string denoting the file_path field in the database.
	FieldFilePath = "file_path"
	// FieldFileName holds the string denoting the file_name field in the database.
	FieldFileName = "file_name"
	// FieldFileSize holds the string denoting the file_size field in the database.
	FieldFileSize = "file_size"
	// EdgeIssue holds the string denoting the issue edge name in mutations.
	EdgeIssue = "issue"
	// Table holds the table name of the file in the database.
	Table = "files"
	// IssueTable is the table that holds the issue relation/edge. The primary key declared below.
	IssueTable = "issue_files"
	// IssueInverseTable is the table name for the Issue entity.
	// It exists in this package in order to avoid circular dependency with the "issue" package.
	IssueInverseTable = "issues"
)

// Columns holds all SQL columns for file fields.
var Columns = []string{
	FieldID,
	FieldFilePath,
	FieldFileName,
	FieldFileSize,
}

var (
	// IssuePrimaryKey and IssueColumn2 are the table columns denoting the
	// primary key for the issue relation (M2M).
	IssuePrimaryKey = []string{"issue_id", "file_id"}
)

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
	// DefaultFileSize holds the default value on creation for the "file_size" field.
	DefaultFileSize int64
)

// OrderOption defines the ordering options for the File queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFilePath orders the results by the file_path field.
func ByFilePath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFilePath, opts...).ToFunc()
}

// ByFileName orders the results by the file_name field.
func ByFileName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileName, opts...).ToFunc()
}

// ByFileSize orders the results by the file_size field.
func ByFileSize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileSize, opts...).ToFunc()
}

// ByIssueCount orders the results by issue count.
func ByIssueCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIssueStep(), opts...)
	}
}

// ByIssue orders the results by issue terms.
func ByIssue(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIssueStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newIssueStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IssueInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, IssueTable, IssuePrimaryKey...),
	)
}