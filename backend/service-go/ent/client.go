// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"aegis/ent/migrate"

	"aegis/ent/tgocache"
	"aegis/ent/tgoens"
	"aegis/ent/tgonft"
	"aegis/ent/tgoretirement"
	"aegis/ent/tuser"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// TGoCache is the client for interacting with the TGoCache builders.
	TGoCache *TGoCacheClient
	// TGoEns is the client for interacting with the TGoEns builders.
	TGoEns *TGoEnsClient
	// TGoNFT is the client for interacting with the TGoNFT builders.
	TGoNFT *TGoNFTClient
	// TGoRetirement is the client for interacting with the TGoRetirement builders.
	TGoRetirement *TGoRetirementClient
	// TUser is the client for interacting with the TUser builders.
	TUser *TUserClient
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
	c.TGoCache = NewTGoCacheClient(c.config)
	c.TGoEns = NewTGoEnsClient(c.config)
	c.TGoNFT = NewTGoNFTClient(c.config)
	c.TGoRetirement = NewTGoRetirementClient(c.config)
	c.TUser = NewTUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		TGoCache:      NewTGoCacheClient(cfg),
		TGoEns:        NewTGoEnsClient(cfg),
		TGoNFT:        NewTGoNFTClient(cfg),
		TGoRetirement: NewTGoRetirementClient(cfg),
		TUser:         NewTUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		TGoCache:      NewTGoCacheClient(cfg),
		TGoEns:        NewTGoEnsClient(cfg),
		TGoNFT:        NewTGoNFTClient(cfg),
		TGoRetirement: NewTGoRetirementClient(cfg),
		TUser:         NewTUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		TGoCache.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
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
	c.TGoCache.Use(hooks...)
	c.TGoEns.Use(hooks...)
	c.TGoNFT.Use(hooks...)
	c.TGoRetirement.Use(hooks...)
	c.TUser.Use(hooks...)
}

// TGoCacheClient is a client for the TGoCache schema.
type TGoCacheClient struct {
	config
}

// NewTGoCacheClient returns a client for the TGoCache from the given config.
func NewTGoCacheClient(c config) *TGoCacheClient {
	return &TGoCacheClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tgocache.Hooks(f(g(h())))`.
func (c *TGoCacheClient) Use(hooks ...Hook) {
	c.hooks.TGoCache = append(c.hooks.TGoCache, hooks...)
}

// Create returns a builder for creating a TGoCache entity.
func (c *TGoCacheClient) Create() *TGoCacheCreate {
	mutation := newTGoCacheMutation(c.config, OpCreate)
	return &TGoCacheCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TGoCache entities.
func (c *TGoCacheClient) CreateBulk(builders ...*TGoCacheCreate) *TGoCacheCreateBulk {
	return &TGoCacheCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TGoCache.
func (c *TGoCacheClient) Update() *TGoCacheUpdate {
	mutation := newTGoCacheMutation(c.config, OpUpdate)
	return &TGoCacheUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TGoCacheClient) UpdateOne(tc *TGoCache) *TGoCacheUpdateOne {
	mutation := newTGoCacheMutation(c.config, OpUpdateOne, withTGoCache(tc))
	return &TGoCacheUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TGoCacheClient) UpdateOneID(id uint64) *TGoCacheUpdateOne {
	mutation := newTGoCacheMutation(c.config, OpUpdateOne, withTGoCacheID(id))
	return &TGoCacheUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TGoCache.
func (c *TGoCacheClient) Delete() *TGoCacheDelete {
	mutation := newTGoCacheMutation(c.config, OpDelete)
	return &TGoCacheDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TGoCacheClient) DeleteOne(tc *TGoCache) *TGoCacheDeleteOne {
	return c.DeleteOneID(tc.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TGoCacheClient) DeleteOneID(id uint64) *TGoCacheDeleteOne {
	builder := c.Delete().Where(tgocache.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TGoCacheDeleteOne{builder}
}

// Query returns a query builder for TGoCache.
func (c *TGoCacheClient) Query() *TGoCacheQuery {
	return &TGoCacheQuery{
		config: c.config,
	}
}

// Get returns a TGoCache entity by its id.
func (c *TGoCacheClient) Get(ctx context.Context, id uint64) (*TGoCache, error) {
	return c.Query().Where(tgocache.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TGoCacheClient) GetX(ctx context.Context, id uint64) *TGoCache {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TGoCacheClient) Hooks() []Hook {
	return c.hooks.TGoCache
}

// TGoEnsClient is a client for the TGoEns schema.
type TGoEnsClient struct {
	config
}

// NewTGoEnsClient returns a client for the TGoEns from the given config.
func NewTGoEnsClient(c config) *TGoEnsClient {
	return &TGoEnsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tgoens.Hooks(f(g(h())))`.
func (c *TGoEnsClient) Use(hooks ...Hook) {
	c.hooks.TGoEns = append(c.hooks.TGoEns, hooks...)
}

// Create returns a builder for creating a TGoEns entity.
func (c *TGoEnsClient) Create() *TGoEnsCreate {
	mutation := newTGoEnsMutation(c.config, OpCreate)
	return &TGoEnsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TGoEns entities.
func (c *TGoEnsClient) CreateBulk(builders ...*TGoEnsCreate) *TGoEnsCreateBulk {
	return &TGoEnsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TGoEns.
func (c *TGoEnsClient) Update() *TGoEnsUpdate {
	mutation := newTGoEnsMutation(c.config, OpUpdate)
	return &TGoEnsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TGoEnsClient) UpdateOne(te *TGoEns) *TGoEnsUpdateOne {
	mutation := newTGoEnsMutation(c.config, OpUpdateOne, withTGoEns(te))
	return &TGoEnsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TGoEnsClient) UpdateOneID(id uint64) *TGoEnsUpdateOne {
	mutation := newTGoEnsMutation(c.config, OpUpdateOne, withTGoEnsID(id))
	return &TGoEnsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TGoEns.
func (c *TGoEnsClient) Delete() *TGoEnsDelete {
	mutation := newTGoEnsMutation(c.config, OpDelete)
	return &TGoEnsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TGoEnsClient) DeleteOne(te *TGoEns) *TGoEnsDeleteOne {
	return c.DeleteOneID(te.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TGoEnsClient) DeleteOneID(id uint64) *TGoEnsDeleteOne {
	builder := c.Delete().Where(tgoens.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TGoEnsDeleteOne{builder}
}

// Query returns a query builder for TGoEns.
func (c *TGoEnsClient) Query() *TGoEnsQuery {
	return &TGoEnsQuery{
		config: c.config,
	}
}

// Get returns a TGoEns entity by its id.
func (c *TGoEnsClient) Get(ctx context.Context, id uint64) (*TGoEns, error) {
	return c.Query().Where(tgoens.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TGoEnsClient) GetX(ctx context.Context, id uint64) *TGoEns {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TGoEnsClient) Hooks() []Hook {
	return c.hooks.TGoEns
}

// TGoNFTClient is a client for the TGoNFT schema.
type TGoNFTClient struct {
	config
}

// NewTGoNFTClient returns a client for the TGoNFT from the given config.
func NewTGoNFTClient(c config) *TGoNFTClient {
	return &TGoNFTClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tgonft.Hooks(f(g(h())))`.
func (c *TGoNFTClient) Use(hooks ...Hook) {
	c.hooks.TGoNFT = append(c.hooks.TGoNFT, hooks...)
}

// Create returns a builder for creating a TGoNFT entity.
func (c *TGoNFTClient) Create() *TGoNFTCreate {
	mutation := newTGoNFTMutation(c.config, OpCreate)
	return &TGoNFTCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TGoNFT entities.
func (c *TGoNFTClient) CreateBulk(builders ...*TGoNFTCreate) *TGoNFTCreateBulk {
	return &TGoNFTCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TGoNFT.
func (c *TGoNFTClient) Update() *TGoNFTUpdate {
	mutation := newTGoNFTMutation(c.config, OpUpdate)
	return &TGoNFTUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TGoNFTClient) UpdateOne(tn *TGoNFT) *TGoNFTUpdateOne {
	mutation := newTGoNFTMutation(c.config, OpUpdateOne, withTGoNFT(tn))
	return &TGoNFTUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TGoNFTClient) UpdateOneID(id uint64) *TGoNFTUpdateOne {
	mutation := newTGoNFTMutation(c.config, OpUpdateOne, withTGoNFTID(id))
	return &TGoNFTUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TGoNFT.
func (c *TGoNFTClient) Delete() *TGoNFTDelete {
	mutation := newTGoNFTMutation(c.config, OpDelete)
	return &TGoNFTDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TGoNFTClient) DeleteOne(tn *TGoNFT) *TGoNFTDeleteOne {
	return c.DeleteOneID(tn.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TGoNFTClient) DeleteOneID(id uint64) *TGoNFTDeleteOne {
	builder := c.Delete().Where(tgonft.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TGoNFTDeleteOne{builder}
}

// Query returns a query builder for TGoNFT.
func (c *TGoNFTClient) Query() *TGoNFTQuery {
	return &TGoNFTQuery{
		config: c.config,
	}
}

// Get returns a TGoNFT entity by its id.
func (c *TGoNFTClient) Get(ctx context.Context, id uint64) (*TGoNFT, error) {
	return c.Query().Where(tgonft.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TGoNFTClient) GetX(ctx context.Context, id uint64) *TGoNFT {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TGoNFTClient) Hooks() []Hook {
	return c.hooks.TGoNFT
}

// TGoRetirementClient is a client for the TGoRetirement schema.
type TGoRetirementClient struct {
	config
}

// NewTGoRetirementClient returns a client for the TGoRetirement from the given config.
func NewTGoRetirementClient(c config) *TGoRetirementClient {
	return &TGoRetirementClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tgoretirement.Hooks(f(g(h())))`.
func (c *TGoRetirementClient) Use(hooks ...Hook) {
	c.hooks.TGoRetirement = append(c.hooks.TGoRetirement, hooks...)
}

// Create returns a builder for creating a TGoRetirement entity.
func (c *TGoRetirementClient) Create() *TGoRetirementCreate {
	mutation := newTGoRetirementMutation(c.config, OpCreate)
	return &TGoRetirementCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TGoRetirement entities.
func (c *TGoRetirementClient) CreateBulk(builders ...*TGoRetirementCreate) *TGoRetirementCreateBulk {
	return &TGoRetirementCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TGoRetirement.
func (c *TGoRetirementClient) Update() *TGoRetirementUpdate {
	mutation := newTGoRetirementMutation(c.config, OpUpdate)
	return &TGoRetirementUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TGoRetirementClient) UpdateOne(tr *TGoRetirement) *TGoRetirementUpdateOne {
	mutation := newTGoRetirementMutation(c.config, OpUpdateOne, withTGoRetirement(tr))
	return &TGoRetirementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TGoRetirementClient) UpdateOneID(id uint64) *TGoRetirementUpdateOne {
	mutation := newTGoRetirementMutation(c.config, OpUpdateOne, withTGoRetirementID(id))
	return &TGoRetirementUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TGoRetirement.
func (c *TGoRetirementClient) Delete() *TGoRetirementDelete {
	mutation := newTGoRetirementMutation(c.config, OpDelete)
	return &TGoRetirementDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TGoRetirementClient) DeleteOne(tr *TGoRetirement) *TGoRetirementDeleteOne {
	return c.DeleteOneID(tr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TGoRetirementClient) DeleteOneID(id uint64) *TGoRetirementDeleteOne {
	builder := c.Delete().Where(tgoretirement.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TGoRetirementDeleteOne{builder}
}

// Query returns a query builder for TGoRetirement.
func (c *TGoRetirementClient) Query() *TGoRetirementQuery {
	return &TGoRetirementQuery{
		config: c.config,
	}
}

// Get returns a TGoRetirement entity by its id.
func (c *TGoRetirementClient) Get(ctx context.Context, id uint64) (*TGoRetirement, error) {
	return c.Query().Where(tgoretirement.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TGoRetirementClient) GetX(ctx context.Context, id uint64) *TGoRetirement {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TGoRetirementClient) Hooks() []Hook {
	return c.hooks.TGoRetirement
}

// TUserClient is a client for the TUser schema.
type TUserClient struct {
	config
}

// NewTUserClient returns a client for the TUser from the given config.
func NewTUserClient(c config) *TUserClient {
	return &TUserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tuser.Hooks(f(g(h())))`.
func (c *TUserClient) Use(hooks ...Hook) {
	c.hooks.TUser = append(c.hooks.TUser, hooks...)
}

// Create returns a builder for creating a TUser entity.
func (c *TUserClient) Create() *TUserCreate {
	mutation := newTUserMutation(c.config, OpCreate)
	return &TUserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TUser entities.
func (c *TUserClient) CreateBulk(builders ...*TUserCreate) *TUserCreateBulk {
	return &TUserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TUser.
func (c *TUserClient) Update() *TUserUpdate {
	mutation := newTUserMutation(c.config, OpUpdate)
	return &TUserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TUserClient) UpdateOne(t *TUser) *TUserUpdateOne {
	mutation := newTUserMutation(c.config, OpUpdateOne, withTUser(t))
	return &TUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TUserClient) UpdateOneID(id uint64) *TUserUpdateOne {
	mutation := newTUserMutation(c.config, OpUpdateOne, withTUserID(id))
	return &TUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TUser.
func (c *TUserClient) Delete() *TUserDelete {
	mutation := newTUserMutation(c.config, OpDelete)
	return &TUserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TUserClient) DeleteOne(t *TUser) *TUserDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TUserClient) DeleteOneID(id uint64) *TUserDeleteOne {
	builder := c.Delete().Where(tuser.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TUserDeleteOne{builder}
}

// Query returns a query builder for TUser.
func (c *TUserClient) Query() *TUserQuery {
	return &TUserQuery{
		config: c.config,
	}
}

// Get returns a TUser entity by its id.
func (c *TUserClient) Get(ctx context.Context, id uint64) (*TUser, error) {
	return c.Query().Where(tuser.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TUserClient) GetX(ctx context.Context, id uint64) *TUser {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TUserClient) Hooks() []Hook {
	return c.hooks.TUser
}
