// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPassowrd holds the string denoting the passowrd field in the database.
	FieldPassowrd = "passowrd"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldOrganization holds the string denoting the organization field in the database.
	FieldOrganization = "organization"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeProjects holds the string denoting the projects edge name in mutations.
	EdgeProjects = "projects"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ProjectsTable is the table that holds the projects relation/edge.
	ProjectsTable = "projects"
	// ProjectsInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectsInverseTable = "projects"
	// ProjectsColumn is the table column denoting the projects relation/edge.
	ProjectsColumn = "user_projects"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPassowrd,
	FieldEmail,
	FieldPhone,
	FieldCreatedAt,
	FieldOrganization,
	FieldTitle,
	FieldType,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// PassowrdValidator is a validator for the "passowrd" field. It is called by the builders before save.
	PassowrdValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// OrganizationValidator is a validator for the "Organization" field. It is called by the builders before save.
	OrganizationValidator func(string) error
	// TitleValidator is a validator for the "Title" field. It is called by the builders before save.
	TitleValidator func(string) error
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeCoordinator Type = "coordinator"
	TypeMember      Type = "member"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeCoordinator, TypeMember:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPassowrd orders the results by the passowrd field.
func ByPassowrd(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassowrd, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByOrganization orders the results by the Organization field.
func ByOrganization(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrganization, opts...).ToFunc()
}

// ByTitle orders the results by the Title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByProjectsCount orders the results by projects count.
func ByProjectsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProjectsStep(), opts...)
	}
}

// ByProjects orders the results by projects terms.
func ByProjects(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProjectsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProjectsTable, ProjectsColumn),
	)
}