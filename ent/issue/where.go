// Code generated by ent, DO NOT EDIT.

package issue

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tarqeem/ims/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Issue {
	return predicate.Issue(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Issue {
	return predicate.Issue(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Issue {
	return predicate.Issue(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Issue {
	return predicate.Issue(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Issue {
	return predicate.Issue(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Issue {
	return predicate.Issue(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Issue {
	return predicate.Issue(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Issue {
	return predicate.Issue(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Issue {
	return predicate.Issue(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Issue {
	return predicate.Issue(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Issue {
	return predicate.Issue(sql.FieldEQ(FieldDescription, v))
}

// Creator applies equality check predicate on the "Creator" field. It's identical to CreatorEQ.
func Creator(v string) predicate.Issue {
	return predicate.Issue(sql.FieldEQ(FieldCreator, v))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v string) predicate.Issue {
	return predicate.Issue(sql.FieldEQ(FieldDate, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Issue {
	return predicate.Issue(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Issue {
	return predicate.Issue(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Issue {
	return predicate.Issue(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Issue {
	return predicate.Issue(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Issue {
	return predicate.Issue(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Issue {
	return predicate.Issue(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Issue {
	return predicate.Issue(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Issue {
	return predicate.Issue(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Issue {
	return predicate.Issue(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Issue {
	return predicate.Issue(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Issue {
	return predicate.Issue(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Issue {
	return predicate.Issue(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Issue {
	return predicate.Issue(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Issue {
	return predicate.Issue(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Issue {
	return predicate.Issue(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Issue {
	return predicate.Issue(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Issue {
	return predicate.Issue(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Issue {
	return predicate.Issue(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Issue {
	return predicate.Issue(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Issue {
	return predicate.Issue(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Issue {
	return predicate.Issue(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Issue {
	return predicate.Issue(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Issue {
	return predicate.Issue(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Issue {
	return predicate.Issue(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Issue {
	return predicate.Issue(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Issue {
	return predicate.Issue(sql.FieldContainsFold(FieldDescription, v))
}

// CreatorEQ applies the EQ predicate on the "Creator" field.
func CreatorEQ(v string) predicate.Issue {
	return predicate.Issue(sql.FieldEQ(FieldCreator, v))
}

// CreatorNEQ applies the NEQ predicate on the "Creator" field.
func CreatorNEQ(v string) predicate.Issue {
	return predicate.Issue(sql.FieldNEQ(FieldCreator, v))
}

// CreatorIn applies the In predicate on the "Creator" field.
func CreatorIn(vs ...string) predicate.Issue {
	return predicate.Issue(sql.FieldIn(FieldCreator, vs...))
}

// CreatorNotIn applies the NotIn predicate on the "Creator" field.
func CreatorNotIn(vs ...string) predicate.Issue {
	return predicate.Issue(sql.FieldNotIn(FieldCreator, vs...))
}

// CreatorGT applies the GT predicate on the "Creator" field.
func CreatorGT(v string) predicate.Issue {
	return predicate.Issue(sql.FieldGT(FieldCreator, v))
}

// CreatorGTE applies the GTE predicate on the "Creator" field.
func CreatorGTE(v string) predicate.Issue {
	return predicate.Issue(sql.FieldGTE(FieldCreator, v))
}

// CreatorLT applies the LT predicate on the "Creator" field.
func CreatorLT(v string) predicate.Issue {
	return predicate.Issue(sql.FieldLT(FieldCreator, v))
}

// CreatorLTE applies the LTE predicate on the "Creator" field.
func CreatorLTE(v string) predicate.Issue {
	return predicate.Issue(sql.FieldLTE(FieldCreator, v))
}

// CreatorContains applies the Contains predicate on the "Creator" field.
func CreatorContains(v string) predicate.Issue {
	return predicate.Issue(sql.FieldContains(FieldCreator, v))
}

// CreatorHasPrefix applies the HasPrefix predicate on the "Creator" field.
func CreatorHasPrefix(v string) predicate.Issue {
	return predicate.Issue(sql.FieldHasPrefix(FieldCreator, v))
}

// CreatorHasSuffix applies the HasSuffix predicate on the "Creator" field.
func CreatorHasSuffix(v string) predicate.Issue {
	return predicate.Issue(sql.FieldHasSuffix(FieldCreator, v))
}

// CreatorEqualFold applies the EqualFold predicate on the "Creator" field.
func CreatorEqualFold(v string) predicate.Issue {
	return predicate.Issue(sql.FieldEqualFold(FieldCreator, v))
}

// CreatorContainsFold applies the ContainsFold predicate on the "Creator" field.
func CreatorContainsFold(v string) predicate.Issue {
	return predicate.Issue(sql.FieldContainsFold(FieldCreator, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Issue {
	return predicate.Issue(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Issue {
	return predicate.Issue(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Issue {
	return predicate.Issue(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Issue {
	return predicate.Issue(sql.FieldNotIn(FieldStatus, vs...))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v string) predicate.Issue {
	return predicate.Issue(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v string) predicate.Issue {
	return predicate.Issue(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...string) predicate.Issue {
	return predicate.Issue(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...string) predicate.Issue {
	return predicate.Issue(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v string) predicate.Issue {
	return predicate.Issue(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v string) predicate.Issue {
	return predicate.Issue(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v string) predicate.Issue {
	return predicate.Issue(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v string) predicate.Issue {
	return predicate.Issue(sql.FieldLTE(FieldDate, v))
}

// DateContains applies the Contains predicate on the "date" field.
func DateContains(v string) predicate.Issue {
	return predicate.Issue(sql.FieldContains(FieldDate, v))
}

// DateHasPrefix applies the HasPrefix predicate on the "date" field.
func DateHasPrefix(v string) predicate.Issue {
	return predicate.Issue(sql.FieldHasPrefix(FieldDate, v))
}

// DateHasSuffix applies the HasSuffix predicate on the "date" field.
func DateHasSuffix(v string) predicate.Issue {
	return predicate.Issue(sql.FieldHasSuffix(FieldDate, v))
}

// DateEqualFold applies the EqualFold predicate on the "date" field.
func DateEqualFold(v string) predicate.Issue {
	return predicate.Issue(sql.FieldEqualFold(FieldDate, v))
}

// DateContainsFold applies the ContainsFold predicate on the "date" field.
func DateContainsFold(v string) predicate.Issue {
	return predicate.Issue(sql.FieldContainsFold(FieldDate, v))
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		step := newProjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComments applies the HasEdge predicate on the "comments" edge.
func HasComments() predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, CommentsTable, CommentsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentsWith applies the HasEdge predicate on the "comments" edge with a given conditions (other predicates).
func HasCommentsWith(preds ...predicate.Comment) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		step := newCommentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFiles applies the HasEdge predicate on the "files" edge.
func HasFiles() predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, FilesTable, FilesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFilesWith applies the HasEdge predicate on the "files" edge with a given conditions (other predicates).
func HasFilesWith(preds ...predicate.File) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		step := newFilesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Issue) predicate.Issue {
	return predicate.Issue(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Issue) predicate.Issue {
	return predicate.Issue(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Issue) predicate.Issue {
	return predicate.Issue(sql.NotPredicates(p))
}
