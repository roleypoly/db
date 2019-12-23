// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/roleypoly/db/ent/guild"
	"github.com/roleypoly/db/ent/predicate"
	"github.com/roleypoly/db/ent/schema"
)

// GuildUpdate is the builder for updating Guild entities.
type GuildUpdate struct {
	config

	updated_at *time.Time

	message      *string
	categories   *[]schema.Category
	entitlements *[]string
	predicates   []predicate.Guild
}

// Where adds a new predicate for the builder.
func (gu *GuildUpdate) Where(ps ...predicate.Guild) *GuildUpdate {
	gu.predicates = append(gu.predicates, ps...)
	return gu
}

// SetUpdatedAt sets the updated_at field.
func (gu *GuildUpdate) SetUpdatedAt(t time.Time) *GuildUpdate {
	gu.updated_at = &t
	return gu
}

// SetMessage sets the message field.
func (gu *GuildUpdate) SetMessage(s string) *GuildUpdate {
	gu.message = &s
	return gu
}

// SetCategories sets the categories field.
func (gu *GuildUpdate) SetCategories(s []schema.Category) *GuildUpdate {
	gu.categories = &s
	return gu
}

// SetEntitlements sets the entitlements field.
func (gu *GuildUpdate) SetEntitlements(s []string) *GuildUpdate {
	gu.entitlements = &s
	return gu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (gu *GuildUpdate) Save(ctx context.Context) (int, error) {
	if gu.updated_at == nil {
		v := guild.UpdateDefaultUpdatedAt()
		gu.updated_at = &v
	}
	return gu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GuildUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GuildUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GuildUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GuildUpdate) sqlSave(ctx context.Context) (n int, err error) {
	var (
		builder  = sql.Dialect(gu.driver.Dialect())
		selector = builder.Select(guild.FieldID).From(builder.Table(guild.Table))
	)
	for _, p := range gu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = gu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("ent: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := gu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		updater = builder.Update(guild.Table)
	)
	updater = updater.Where(sql.InInts(guild.FieldID, ids...))
	if value := gu.updated_at; value != nil {
		updater.Set(guild.FieldUpdatedAt, *value)
	}
	if value := gu.message; value != nil {
		updater.Set(guild.FieldMessage, *value)
	}
	if value := gu.categories; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return 0, err
		}
		updater.Set(guild.FieldCategories, buf)
	}
	if value := gu.entitlements; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return 0, err
		}
		updater.Set(guild.FieldEntitlements, buf)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// GuildUpdateOne is the builder for updating a single Guild entity.
type GuildUpdateOne struct {
	config
	id int

	updated_at *time.Time

	message      *string
	categories   *[]schema.Category
	entitlements *[]string
}

// SetUpdatedAt sets the updated_at field.
func (guo *GuildUpdateOne) SetUpdatedAt(t time.Time) *GuildUpdateOne {
	guo.updated_at = &t
	return guo
}

// SetMessage sets the message field.
func (guo *GuildUpdateOne) SetMessage(s string) *GuildUpdateOne {
	guo.message = &s
	return guo
}

// SetCategories sets the categories field.
func (guo *GuildUpdateOne) SetCategories(s []schema.Category) *GuildUpdateOne {
	guo.categories = &s
	return guo
}

// SetEntitlements sets the entitlements field.
func (guo *GuildUpdateOne) SetEntitlements(s []string) *GuildUpdateOne {
	guo.entitlements = &s
	return guo
}

// Save executes the query and returns the updated entity.
func (guo *GuildUpdateOne) Save(ctx context.Context) (*Guild, error) {
	if guo.updated_at == nil {
		v := guild.UpdateDefaultUpdatedAt()
		guo.updated_at = &v
	}
	return guo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GuildUpdateOne) SaveX(ctx context.Context) *Guild {
	gu, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return gu
}

// Exec executes the query on the entity.
func (guo *GuildUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GuildUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GuildUpdateOne) sqlSave(ctx context.Context) (gu *Guild, err error) {
	var (
		builder  = sql.Dialect(guo.driver.Dialect())
		selector = builder.Select(guild.Columns...).From(builder.Table(guild.Table))
	)
	guild.ID(guo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = guo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		gu = &Guild{config: guo.config}
		if err := gu.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Guild: %v", err)
		}
		id = gu.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("Guild with id: %v", guo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Guild with the same id: %v", guo.id)
	}

	tx, err := guo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		updater = builder.Update(guild.Table)
	)
	updater = updater.Where(sql.InInts(guild.FieldID, ids...))
	if value := guo.updated_at; value != nil {
		updater.Set(guild.FieldUpdatedAt, *value)
		gu.UpdatedAt = *value
	}
	if value := guo.message; value != nil {
		updater.Set(guild.FieldMessage, *value)
		gu.Message = *value
	}
	if value := guo.categories; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return nil, err
		}
		updater.Set(guild.FieldCategories, buf)
		gu.Categories = *value
	}
	if value := guo.entitlements; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return nil, err
		}
		updater.Set(guild.FieldEntitlements, buf)
		gu.Entitlements = *value
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return gu, nil
}
