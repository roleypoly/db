// Code generated by entc, DO NOT EDIT.

package challenge

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/roleypoly/db/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// ChallengeID applies equality check predicate on the "challenge_id" field. It's identical to ChallengeIDEQ.
func ChallengeID(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChallengeID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// Human applies equality check predicate on the "human" field. It's identical to HumanEQ.
func Human(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHuman), v))
	})
}

// Magic applies equality check predicate on the "magic" field. It's identical to MagicEQ.
func Magic(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMagic), v))
	})
}

// ExpiresAt applies equality check predicate on the "expires_at" field. It's identical to ExpiresAtEQ.
func ExpiresAt(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiresAt), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Challenge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Challenge(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Challenge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Challenge(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Challenge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Challenge(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Challenge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Challenge(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// ChallengeIDEQ applies the EQ predicate on the "challenge_id" field.
func ChallengeIDEQ(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChallengeID), v))
	})
}

// ChallengeIDNEQ applies the NEQ predicate on the "challenge_id" field.
func ChallengeIDNEQ(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChallengeID), v))
	})
}

// ChallengeIDIn applies the In predicate on the "challenge_id" field.
func ChallengeIDIn(vs ...string) predicate.Challenge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Challenge(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldChallengeID), v...))
	})
}

// ChallengeIDNotIn applies the NotIn predicate on the "challenge_id" field.
func ChallengeIDNotIn(vs ...string) predicate.Challenge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Challenge(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldChallengeID), v...))
	})
}

// ChallengeIDGT applies the GT predicate on the "challenge_id" field.
func ChallengeIDGT(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChallengeID), v))
	})
}

// ChallengeIDGTE applies the GTE predicate on the "challenge_id" field.
func ChallengeIDGTE(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChallengeID), v))
	})
}

// ChallengeIDLT applies the LT predicate on the "challenge_id" field.
func ChallengeIDLT(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChallengeID), v))
	})
}

// ChallengeIDLTE applies the LTE predicate on the "challenge_id" field.
func ChallengeIDLTE(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChallengeID), v))
	})
}

// ChallengeIDContains applies the Contains predicate on the "challenge_id" field.
func ChallengeIDContains(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldChallengeID), v))
	})
}

// ChallengeIDHasPrefix applies the HasPrefix predicate on the "challenge_id" field.
func ChallengeIDHasPrefix(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldChallengeID), v))
	})
}

// ChallengeIDHasSuffix applies the HasSuffix predicate on the "challenge_id" field.
func ChallengeIDHasSuffix(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldChallengeID), v))
	})
}

// ChallengeIDEqualFold applies the EqualFold predicate on the "challenge_id" field.
func ChallengeIDEqualFold(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldChallengeID), v))
	})
}

// ChallengeIDContainsFold applies the ContainsFold predicate on the "challenge_id" field.
func ChallengeIDContainsFold(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldChallengeID), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Challenge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Challenge(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Challenge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Challenge(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserID), v))
	})
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserID), v))
	})
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserID), v))
	})
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserID), v))
	})
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserID), v))
	})
}

// HumanEQ applies the EQ predicate on the "human" field.
func HumanEQ(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHuman), v))
	})
}

// HumanNEQ applies the NEQ predicate on the "human" field.
func HumanNEQ(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHuman), v))
	})
}

// HumanIn applies the In predicate on the "human" field.
func HumanIn(vs ...string) predicate.Challenge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Challenge(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHuman), v...))
	})
}

// HumanNotIn applies the NotIn predicate on the "human" field.
func HumanNotIn(vs ...string) predicate.Challenge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Challenge(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHuman), v...))
	})
}

// HumanGT applies the GT predicate on the "human" field.
func HumanGT(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHuman), v))
	})
}

// HumanGTE applies the GTE predicate on the "human" field.
func HumanGTE(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHuman), v))
	})
}

// HumanLT applies the LT predicate on the "human" field.
func HumanLT(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHuman), v))
	})
}

// HumanLTE applies the LTE predicate on the "human" field.
func HumanLTE(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHuman), v))
	})
}

// HumanContains applies the Contains predicate on the "human" field.
func HumanContains(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHuman), v))
	})
}

// HumanHasPrefix applies the HasPrefix predicate on the "human" field.
func HumanHasPrefix(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHuman), v))
	})
}

// HumanHasSuffix applies the HasSuffix predicate on the "human" field.
func HumanHasSuffix(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHuman), v))
	})
}

// HumanEqualFold applies the EqualFold predicate on the "human" field.
func HumanEqualFold(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHuman), v))
	})
}

// HumanContainsFold applies the ContainsFold predicate on the "human" field.
func HumanContainsFold(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHuman), v))
	})
}

// MagicEQ applies the EQ predicate on the "magic" field.
func MagicEQ(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMagic), v))
	})
}

// MagicNEQ applies the NEQ predicate on the "magic" field.
func MagicNEQ(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMagic), v))
	})
}

// MagicIn applies the In predicate on the "magic" field.
func MagicIn(vs ...string) predicate.Challenge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Challenge(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMagic), v...))
	})
}

// MagicNotIn applies the NotIn predicate on the "magic" field.
func MagicNotIn(vs ...string) predicate.Challenge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Challenge(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMagic), v...))
	})
}

// MagicGT applies the GT predicate on the "magic" field.
func MagicGT(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMagic), v))
	})
}

// MagicGTE applies the GTE predicate on the "magic" field.
func MagicGTE(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMagic), v))
	})
}

// MagicLT applies the LT predicate on the "magic" field.
func MagicLT(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMagic), v))
	})
}

// MagicLTE applies the LTE predicate on the "magic" field.
func MagicLTE(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMagic), v))
	})
}

// MagicContains applies the Contains predicate on the "magic" field.
func MagicContains(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMagic), v))
	})
}

// MagicHasPrefix applies the HasPrefix predicate on the "magic" field.
func MagicHasPrefix(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMagic), v))
	})
}

// MagicHasSuffix applies the HasSuffix predicate on the "magic" field.
func MagicHasSuffix(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMagic), v))
	})
}

// MagicEqualFold applies the EqualFold predicate on the "magic" field.
func MagicEqualFold(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMagic), v))
	})
}

// MagicContainsFold applies the ContainsFold predicate on the "magic" field.
func MagicContainsFold(v string) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMagic), v))
	})
}

// ExpiresAtEQ applies the EQ predicate on the "expires_at" field.
func ExpiresAtEQ(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiresAt), v))
	})
}

// ExpiresAtNEQ applies the NEQ predicate on the "expires_at" field.
func ExpiresAtNEQ(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpiresAt), v))
	})
}

// ExpiresAtIn applies the In predicate on the "expires_at" field.
func ExpiresAtIn(vs ...time.Time) predicate.Challenge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Challenge(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExpiresAt), v...))
	})
}

// ExpiresAtNotIn applies the NotIn predicate on the "expires_at" field.
func ExpiresAtNotIn(vs ...time.Time) predicate.Challenge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Challenge(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExpiresAt), v...))
	})
}

// ExpiresAtGT applies the GT predicate on the "expires_at" field.
func ExpiresAtGT(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpiresAt), v))
	})
}

// ExpiresAtGTE applies the GTE predicate on the "expires_at" field.
func ExpiresAtGTE(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpiresAt), v))
	})
}

// ExpiresAtLT applies the LT predicate on the "expires_at" field.
func ExpiresAtLT(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpiresAt), v))
	})
}

// ExpiresAtLTE applies the LTE predicate on the "expires_at" field.
func ExpiresAtLTE(v time.Time) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpiresAt), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Challenge) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Challenge) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
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
func Not(p predicate.Challenge) predicate.Challenge {
	return predicate.Challenge(func(s *sql.Selector) {
		p(s.Not())
	})
}
