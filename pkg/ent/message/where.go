// Code generated by entc, DO NOT EDIT.

package message

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gusarow4321/TinyChat/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ChatID applies equality check predicate on the "chatID" field. It's identical to ChatIDEQ.
func ChatID(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChatID), v))
	})
}

// UserID applies equality check predicate on the "userID" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldText), v))
	})
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// ChatIDEQ applies the EQ predicate on the "chatID" field.
func ChatIDEQ(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChatID), v))
	})
}

// ChatIDNEQ applies the NEQ predicate on the "chatID" field.
func ChatIDNEQ(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChatID), v))
	})
}

// ChatIDIn applies the In predicate on the "chatID" field.
func ChatIDIn(vs ...int64) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldChatID), v...))
	})
}

// ChatIDNotIn applies the NotIn predicate on the "chatID" field.
func ChatIDNotIn(vs ...int64) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldChatID), v...))
	})
}

// UserIDEQ applies the EQ predicate on the "userID" field.
func UserIDEQ(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "userID" field.
func UserIDNEQ(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "userID" field.
func UserIDIn(vs ...int64) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "userID" field.
func UserIDNotIn(vs ...int64) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldText), v))
	})
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldText), v))
	})
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldText), v...))
	})
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldText), v...))
	})
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldText), v))
	})
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldText), v))
	})
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldText), v))
	})
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldText), v))
	})
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldText), v))
	})
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldText), v))
	})
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldText), v))
	})
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldText), v))
	})
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldText), v))
	})
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTimestamp), v))
	})
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTimestamp), v...))
	})
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTimestamp), v...))
	})
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTimestamp), v))
	})
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTimestamp), v))
	})
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTimestamp), v))
	})
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTimestamp), v))
	})
}

// HasChat applies the HasEdge predicate on the "chat" edge.
func HasChat() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChatTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ChatTable, ChatColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChatWith applies the HasEdge predicate on the "chat" edge with a given conditions (other predicates).
func HasChatWith(preds ...predicate.Chat) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChatInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ChatTable, ChatColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		p(s.Not())
	})
}
