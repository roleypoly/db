// Code generated by entc, DO NOT EDIT.

package session

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/roleypoly/db/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldID), id))
		},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
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
		},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
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
		},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldID), id))
		},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldID), id))
		},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldID), id))
		},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldID), id))
		},
	)
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldCreatedAt), v))
		},
	)
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
		},
	)
}

// SessionID applies equality check predicate on the "session_id" field. It's identical to SessionIDEQ.
func SessionID(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldSessionID), v))
		},
	)
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldUserID), v))
		},
	)
}

// ExpiresAt applies equality check predicate on the "expires_at" field. It's identical to ExpiresAtEQ.
func ExpiresAt(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldExpiresAt), v))
		},
	)
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldCreatedAt), v))
		},
	)
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
		},
	)
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Session {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldCreatedAt), v...))
		},
	)
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Session {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
		},
	)
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldCreatedAt), v))
		},
	)
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldCreatedAt), v))
		},
	)
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldCreatedAt), v))
		},
	)
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldCreatedAt), v))
		},
	)
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
		},
	)
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
		},
	)
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Session {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldUpdatedAt), v...))
		},
	)
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Session {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
		},
	)
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldUpdatedAt), v))
		},
	)
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
		},
	)
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldUpdatedAt), v))
		},
	)
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
		},
	)
}

// SessionIDEQ applies the EQ predicate on the "session_id" field.
func SessionIDEQ(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDNEQ applies the NEQ predicate on the "session_id" field.
func SessionIDNEQ(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDIn applies the In predicate on the "session_id" field.
func SessionIDIn(vs ...string) predicate.Session {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldSessionID), v...))
		},
	)
}

// SessionIDNotIn applies the NotIn predicate on the "session_id" field.
func SessionIDNotIn(vs ...string) predicate.Session {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldSessionID), v...))
		},
	)
}

// SessionIDGT applies the GT predicate on the "session_id" field.
func SessionIDGT(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDGTE applies the GTE predicate on the "session_id" field.
func SessionIDGTE(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDLT applies the LT predicate on the "session_id" field.
func SessionIDLT(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDLTE applies the LTE predicate on the "session_id" field.
func SessionIDLTE(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDContains applies the Contains predicate on the "session_id" field.
func SessionIDContains(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDHasPrefix applies the HasPrefix predicate on the "session_id" field.
func SessionIDHasPrefix(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDHasSuffix applies the HasSuffix predicate on the "session_id" field.
func SessionIDHasSuffix(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDEqualFold applies the EqualFold predicate on the "session_id" field.
func SessionIDEqualFold(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDContainsFold applies the ContainsFold predicate on the "session_id" field.
func SessionIDContainsFold(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldSessionID), v))
		},
	)
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldUserID), v))
		},
	)
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldUserID), v))
		},
	)
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Session {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldUserID), v...))
		},
	)
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Session {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldUserID), v...))
		},
	)
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldUserID), v))
		},
	)
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldUserID), v))
		},
	)
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldUserID), v))
		},
	)
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldUserID), v))
		},
	)
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldUserID), v))
		},
	)
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldUserID), v))
		},
	)
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldUserID), v))
		},
	)
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldUserID), v))
		},
	)
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldUserID), v))
		},
	)
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v Source) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldSource), v))
		},
	)
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v Source) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldSource), v))
		},
	)
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...Source) predicate.Session {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldSource), v...))
		},
	)
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...Source) predicate.Session {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldSource), v...))
		},
	)
}

// ExpiresAtEQ applies the EQ predicate on the "expires_at" field.
func ExpiresAtEQ(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldExpiresAt), v))
		},
	)
}

// ExpiresAtNEQ applies the NEQ predicate on the "expires_at" field.
func ExpiresAtNEQ(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldExpiresAt), v))
		},
	)
}

// ExpiresAtIn applies the In predicate on the "expires_at" field.
func ExpiresAtIn(vs ...time.Time) predicate.Session {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldExpiresAt), v...))
		},
	)
}

// ExpiresAtNotIn applies the NotIn predicate on the "expires_at" field.
func ExpiresAtNotIn(vs ...time.Time) predicate.Session {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldExpiresAt), v...))
		},
	)
}

// ExpiresAtGT applies the GT predicate on the "expires_at" field.
func ExpiresAtGT(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldExpiresAt), v))
		},
	)
}

// ExpiresAtGTE applies the GTE predicate on the "expires_at" field.
func ExpiresAtGTE(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldExpiresAt), v))
		},
	)
}

// ExpiresAtLT applies the LT predicate on the "expires_at" field.
func ExpiresAtLT(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldExpiresAt), v))
		},
	)
}

// ExpiresAtLTE applies the LTE predicate on the "expires_at" field.
func ExpiresAtLTE(v time.Time) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldExpiresAt), v))
		},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Session) predicate.Session {
	return predicate.Session(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
