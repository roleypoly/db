// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/roleypoly/db/ent/challenge"
)

// ChallengeCreate is the builder for creating a Challenge entity.
type ChallengeCreate struct {
	config
	created_at   *time.Time
	updated_at   *time.Time
	Challenge_id *string
	user_id      *string
	source       *challenge.Source
	expires_at   *time.Time
}

// SetCreatedAt sets the created_at field.
func (cc *ChallengeCreate) SetCreatedAt(t time.Time) *ChallengeCreate {
	cc.created_at = &t
	return cc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (cc *ChallengeCreate) SetNillableCreatedAt(t *time.Time) *ChallengeCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the updated_at field.
func (cc *ChallengeCreate) SetUpdatedAt(t time.Time) *ChallengeCreate {
	cc.updated_at = &t
	return cc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (cc *ChallengeCreate) SetNillableUpdatedAt(t *time.Time) *ChallengeCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetChallengeID sets the Challenge_id field.
func (cc *ChallengeCreate) SetChallengeID(s string) *ChallengeCreate {
	cc.Challenge_id = &s
	return cc
}

// SetUserID sets the user_id field.
func (cc *ChallengeCreate) SetUserID(s string) *ChallengeCreate {
	cc.user_id = &s
	return cc
}

// SetSource sets the source field.
func (cc *ChallengeCreate) SetSource(c challenge.Source) *ChallengeCreate {
	cc.source = &c
	return cc
}

// SetExpiresAt sets the expires_at field.
func (cc *ChallengeCreate) SetExpiresAt(t time.Time) *ChallengeCreate {
	cc.expires_at = &t
	return cc
}

// SetNillableExpiresAt sets the expires_at field if the given value is not nil.
func (cc *ChallengeCreate) SetNillableExpiresAt(t *time.Time) *ChallengeCreate {
	if t != nil {
		cc.SetExpiresAt(*t)
	}
	return cc
}

// Save creates the Challenge in the database.
func (cc *ChallengeCreate) Save(ctx context.Context) (*Challenge, error) {
	if cc.created_at == nil {
		v := challenge.DefaultCreatedAt()
		cc.created_at = &v
	}
	if cc.updated_at == nil {
		v := challenge.DefaultUpdatedAt()
		cc.updated_at = &v
	}
	if cc.Challenge_id == nil {
		return nil, errors.New("ent: missing required field \"Challenge_id\"")
	}
	if cc.user_id == nil {
		return nil, errors.New("ent: missing required field \"user_id\"")
	}
	if cc.source == nil {
		return nil, errors.New("ent: missing required field \"source\"")
	}
	if err := challenge.SourceValidator(*cc.source); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"source\": %v", err)
	}
	if cc.expires_at == nil {
		v := challenge.DefaultExpiresAt()
		cc.expires_at = &v
	}
	return cc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChallengeCreate) SaveX(ctx context.Context) *Challenge {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *ChallengeCreate) sqlSave(ctx context.Context) (*Challenge, error) {
	var (
		builder = sql.Dialect(cc.driver.Dialect())
		c       = &Challenge{config: cc.config}
	)
	tx, err := cc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(challenge.Table).Default()
	if value := cc.created_at; value != nil {
		insert.Set(challenge.FieldCreatedAt, *value)
		c.CreatedAt = *value
	}
	if value := cc.updated_at; value != nil {
		insert.Set(challenge.FieldUpdatedAt, *value)
		c.UpdatedAt = *value
	}
	if value := cc.Challenge_id; value != nil {
		insert.Set(challenge.FieldChallengeID, *value)
		c.ChallengeID = *value
	}
	if value := cc.user_id; value != nil {
		insert.Set(challenge.FieldUserID, *value)
		c.UserID = *value
	}
	if value := cc.source; value != nil {
		insert.Set(challenge.FieldSource, *value)
		c.Source = *value
	}
	if value := cc.expires_at; value != nil {
		insert.Set(challenge.FieldExpiresAt, *value)
		c.ExpiresAt = *value
	}

	id, err := insertLastID(ctx, tx, insert.Returning(challenge.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	c.ID = int(id)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return c, nil
}
