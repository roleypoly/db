// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/roleypoly/db/ent/migrate"

	"github.com/roleypoly/db/ent/challenge"
	"github.com/roleypoly/db/ent/guild"
	"github.com/roleypoly/db/ent/session"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Challenge is the client for interacting with the Challenge builders.
	Challenge *ChallengeClient
	// Guild is the client for interacting with the Guild builders.
	Guild *GuildClient
	// Session is the client for interacting with the Session builders.
	Session *SessionClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Challenge = NewChallengeClient(c.config)
	c.Guild = NewGuildClient(c.config)
	c.Session = NewSessionClient(c.config)
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:    cfg,
		Challenge: NewChallengeClient(cfg),
		Guild:     NewGuildClient(cfg),
		Session:   NewSessionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Challenge.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Challenge.Use(hooks...)
	c.Guild.Use(hooks...)
	c.Session.Use(hooks...)
}

// ChallengeClient is a client for the Challenge schema.
type ChallengeClient struct {
	config
}

// NewChallengeClient returns a client for the Challenge from the given config.
func NewChallengeClient(c config) *ChallengeClient {
	return &ChallengeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `challenge.Hooks(f(g(h())))`.
func (c *ChallengeClient) Use(hooks ...Hook) {
	c.hooks.Challenge = append(c.hooks.Challenge, hooks...)
}

// Create returns a create builder for Challenge.
func (c *ChallengeClient) Create() *ChallengeCreate {
	mutation := newChallengeMutation(c.config, OpCreate)
	return &ChallengeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Challenge.
func (c *ChallengeClient) Update() *ChallengeUpdate {
	mutation := newChallengeMutation(c.config, OpUpdate)
	return &ChallengeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChallengeClient) UpdateOne(ch *Challenge) *ChallengeUpdateOne {
	return c.UpdateOneID(ch.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *ChallengeClient) UpdateOneID(id int) *ChallengeUpdateOne {
	mutation := newChallengeMutation(c.config, OpUpdateOne)
	mutation.id = &id
	return &ChallengeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Challenge.
func (c *ChallengeClient) Delete() *ChallengeDelete {
	mutation := newChallengeMutation(c.config, OpDelete)
	return &ChallengeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ChallengeClient) DeleteOne(ch *Challenge) *ChallengeDeleteOne {
	return c.DeleteOneID(ch.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ChallengeClient) DeleteOneID(id int) *ChallengeDeleteOne {
	builder := c.Delete().Where(challenge.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChallengeDeleteOne{builder}
}

// Create returns a query builder for Challenge.
func (c *ChallengeClient) Query() *ChallengeQuery {
	return &ChallengeQuery{config: c.config}
}

// Get returns a Challenge entity by its id.
func (c *ChallengeClient) Get(ctx context.Context, id int) (*Challenge, error) {
	return c.Query().Where(challenge.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChallengeClient) GetX(ctx context.Context, id int) *Challenge {
	ch, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ch
}

// Hooks returns the client hooks.
func (c *ChallengeClient) Hooks() []Hook {
	return c.hooks.Challenge
}

// GuildClient is a client for the Guild schema.
type GuildClient struct {
	config
}

// NewGuildClient returns a client for the Guild from the given config.
func NewGuildClient(c config) *GuildClient {
	return &GuildClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `guild.Hooks(f(g(h())))`.
func (c *GuildClient) Use(hooks ...Hook) {
	c.hooks.Guild = append(c.hooks.Guild, hooks...)
}

// Create returns a create builder for Guild.
func (c *GuildClient) Create() *GuildCreate {
	mutation := newGuildMutation(c.config, OpCreate)
	return &GuildCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Guild.
func (c *GuildClient) Update() *GuildUpdate {
	mutation := newGuildMutation(c.config, OpUpdate)
	return &GuildUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GuildClient) UpdateOne(gu *Guild) *GuildUpdateOne {
	return c.UpdateOneID(gu.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *GuildClient) UpdateOneID(id int) *GuildUpdateOne {
	mutation := newGuildMutation(c.config, OpUpdateOne)
	mutation.id = &id
	return &GuildUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Guild.
func (c *GuildClient) Delete() *GuildDelete {
	mutation := newGuildMutation(c.config, OpDelete)
	return &GuildDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GuildClient) DeleteOne(gu *Guild) *GuildDeleteOne {
	return c.DeleteOneID(gu.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GuildClient) DeleteOneID(id int) *GuildDeleteOne {
	builder := c.Delete().Where(guild.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GuildDeleteOne{builder}
}

// Create returns a query builder for Guild.
func (c *GuildClient) Query() *GuildQuery {
	return &GuildQuery{config: c.config}
}

// Get returns a Guild entity by its id.
func (c *GuildClient) Get(ctx context.Context, id int) (*Guild, error) {
	return c.Query().Where(guild.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GuildClient) GetX(ctx context.Context, id int) *Guild {
	gu, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return gu
}

// Hooks returns the client hooks.
func (c *GuildClient) Hooks() []Hook {
	return c.hooks.Guild
}

// SessionClient is a client for the Session schema.
type SessionClient struct {
	config
}

// NewSessionClient returns a client for the Session from the given config.
func NewSessionClient(c config) *SessionClient {
	return &SessionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `session.Hooks(f(g(h())))`.
func (c *SessionClient) Use(hooks ...Hook) {
	c.hooks.Session = append(c.hooks.Session, hooks...)
}

// Create returns a create builder for Session.
func (c *SessionClient) Create() *SessionCreate {
	mutation := newSessionMutation(c.config, OpCreate)
	return &SessionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Session.
func (c *SessionClient) Update() *SessionUpdate {
	mutation := newSessionMutation(c.config, OpUpdate)
	return &SessionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SessionClient) UpdateOne(s *Session) *SessionUpdateOne {
	return c.UpdateOneID(s.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *SessionClient) UpdateOneID(id int) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne)
	mutation.id = &id
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Session.
func (c *SessionClient) Delete() *SessionDelete {
	mutation := newSessionMutation(c.config, OpDelete)
	return &SessionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SessionClient) DeleteOne(s *Session) *SessionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SessionClient) DeleteOneID(id int) *SessionDeleteOne {
	builder := c.Delete().Where(session.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SessionDeleteOne{builder}
}

// Create returns a query builder for Session.
func (c *SessionClient) Query() *SessionQuery {
	return &SessionQuery{config: c.config}
}

// Get returns a Session entity by its id.
func (c *SessionClient) Get(ctx context.Context, id int) (*Session, error) {
	return c.Query().Where(session.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SessionClient) GetX(ctx context.Context, id int) *Session {
	s, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return s
}

// Hooks returns the client hooks.
func (c *SessionClient) Hooks() []Hook {
	return c.hooks.Session
}
