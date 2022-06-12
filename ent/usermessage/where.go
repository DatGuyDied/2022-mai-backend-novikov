// Code generated by entc, DO NOT EDIT.

package usermessage

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DatGuyDied/2022-mai-backend-novikov/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
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
func IDGT(id int) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// From applies equality check predicate on the "from" field. It's identical to FromEQ.
func From(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFrom), v))
	})
}

// To applies equality check predicate on the "to" field. It's identical to ToEQ.
func To(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTo), v))
	})
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldText), v))
	})
}

// CtreatedAt applies equality check predicate on the "ctreated_at" field. It's identical to CtreatedAtEQ.
func CtreatedAt(v time.Time) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCtreatedAt), v))
	})
}

// FromEQ applies the EQ predicate on the "from" field.
func FromEQ(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFrom), v))
	})
}

// FromNEQ applies the NEQ predicate on the "from" field.
func FromNEQ(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFrom), v))
	})
}

// FromIn applies the In predicate on the "from" field.
func FromIn(vs ...string) predicate.UserMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMessage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFrom), v...))
	})
}

// FromNotIn applies the NotIn predicate on the "from" field.
func FromNotIn(vs ...string) predicate.UserMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMessage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFrom), v...))
	})
}

// FromGT applies the GT predicate on the "from" field.
func FromGT(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFrom), v))
	})
}

// FromGTE applies the GTE predicate on the "from" field.
func FromGTE(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFrom), v))
	})
}

// FromLT applies the LT predicate on the "from" field.
func FromLT(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFrom), v))
	})
}

// FromLTE applies the LTE predicate on the "from" field.
func FromLTE(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFrom), v))
	})
}

// FromContains applies the Contains predicate on the "from" field.
func FromContains(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFrom), v))
	})
}

// FromHasPrefix applies the HasPrefix predicate on the "from" field.
func FromHasPrefix(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFrom), v))
	})
}

// FromHasSuffix applies the HasSuffix predicate on the "from" field.
func FromHasSuffix(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFrom), v))
	})
}

// FromEqualFold applies the EqualFold predicate on the "from" field.
func FromEqualFold(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFrom), v))
	})
}

// FromContainsFold applies the ContainsFold predicate on the "from" field.
func FromContainsFold(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFrom), v))
	})
}

// ToEQ applies the EQ predicate on the "to" field.
func ToEQ(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTo), v))
	})
}

// ToNEQ applies the NEQ predicate on the "to" field.
func ToNEQ(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTo), v))
	})
}

// ToIn applies the In predicate on the "to" field.
func ToIn(vs ...string) predicate.UserMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMessage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTo), v...))
	})
}

// ToNotIn applies the NotIn predicate on the "to" field.
func ToNotIn(vs ...string) predicate.UserMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMessage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTo), v...))
	})
}

// ToGT applies the GT predicate on the "to" field.
func ToGT(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTo), v))
	})
}

// ToGTE applies the GTE predicate on the "to" field.
func ToGTE(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTo), v))
	})
}

// ToLT applies the LT predicate on the "to" field.
func ToLT(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTo), v))
	})
}

// ToLTE applies the LTE predicate on the "to" field.
func ToLTE(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTo), v))
	})
}

// ToContains applies the Contains predicate on the "to" field.
func ToContains(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTo), v))
	})
}

// ToHasPrefix applies the HasPrefix predicate on the "to" field.
func ToHasPrefix(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTo), v))
	})
}

// ToHasSuffix applies the HasSuffix predicate on the "to" field.
func ToHasSuffix(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTo), v))
	})
}

// ToEqualFold applies the EqualFold predicate on the "to" field.
func ToEqualFold(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTo), v))
	})
}

// ToContainsFold applies the ContainsFold predicate on the "to" field.
func ToContainsFold(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTo), v))
	})
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldText), v))
	})
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldText), v))
	})
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.UserMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMessage(func(s *sql.Selector) {
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
func TextNotIn(vs ...string) predicate.UserMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMessage(func(s *sql.Selector) {
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
func TextGT(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldText), v))
	})
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldText), v))
	})
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldText), v))
	})
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldText), v))
	})
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldText), v))
	})
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldText), v))
	})
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldText), v))
	})
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldText), v))
	})
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldText), v))
	})
}

// CtreatedAtEQ applies the EQ predicate on the "ctreated_at" field.
func CtreatedAtEQ(v time.Time) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCtreatedAt), v))
	})
}

// CtreatedAtNEQ applies the NEQ predicate on the "ctreated_at" field.
func CtreatedAtNEQ(v time.Time) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCtreatedAt), v))
	})
}

// CtreatedAtIn applies the In predicate on the "ctreated_at" field.
func CtreatedAtIn(vs ...time.Time) predicate.UserMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMessage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCtreatedAt), v...))
	})
}

// CtreatedAtNotIn applies the NotIn predicate on the "ctreated_at" field.
func CtreatedAtNotIn(vs ...time.Time) predicate.UserMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMessage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCtreatedAt), v...))
	})
}

// CtreatedAtGT applies the GT predicate on the "ctreated_at" field.
func CtreatedAtGT(v time.Time) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCtreatedAt), v))
	})
}

// CtreatedAtGTE applies the GTE predicate on the "ctreated_at" field.
func CtreatedAtGTE(v time.Time) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCtreatedAt), v))
	})
}

// CtreatedAtLT applies the LT predicate on the "ctreated_at" field.
func CtreatedAtLT(v time.Time) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCtreatedAt), v))
	})
}

// CtreatedAtLTE applies the LTE predicate on the "ctreated_at" field.
func CtreatedAtLTE(v time.Time) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCtreatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserMessage) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserMessage) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
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
func Not(p predicate.UserMessage) predicate.UserMessage {
	return predicate.UserMessage(func(s *sql.Selector) {
		p(s.Not())
	})
}
