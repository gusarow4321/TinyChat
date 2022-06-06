// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/gusarow4321/TinyChat/pkg/ent/migrate"

	"github.com/gusarow4321/TinyChat/pkg/ent/chat"
	"github.com/gusarow4321/TinyChat/pkg/ent/message"
	"github.com/gusarow4321/TinyChat/pkg/ent/user"
	"github.com/gusarow4321/TinyChat/pkg/ent/usermetadata"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Chat is the client for interacting with the Chat builders.
	Chat *ChatClient
	// Message is the client for interacting with the Message builders.
	Message *MessageClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserMetadata is the client for interacting with the UserMetadata builders.
	UserMetadata *UserMetadataClient
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
	c.Chat = NewChatClient(c.config)
	c.Message = NewMessageClient(c.config)
	c.User = NewUserClient(c.config)
	c.UserMetadata = NewUserMetadataClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Chat:         NewChatClient(cfg),
		Message:      NewMessageClient(cfg),
		User:         NewUserClient(cfg),
		UserMetadata: NewUserMetadataClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		ctx:          ctx,
		config:       cfg,
		Chat:         NewChatClient(cfg),
		Message:      NewMessageClient(cfg),
		User:         NewUserClient(cfg),
		UserMetadata: NewUserMetadataClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Chat.
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
	c.Chat.Use(hooks...)
	c.Message.Use(hooks...)
	c.User.Use(hooks...)
	c.UserMetadata.Use(hooks...)
}

// ChatClient is a client for the Chat schema.
type ChatClient struct {
	config
}

// NewChatClient returns a client for the Chat from the given config.
func NewChatClient(c config) *ChatClient {
	return &ChatClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chat.Hooks(f(g(h())))`.
func (c *ChatClient) Use(hooks ...Hook) {
	c.hooks.Chat = append(c.hooks.Chat, hooks...)
}

// Create returns a create builder for Chat.
func (c *ChatClient) Create() *ChatCreate {
	mutation := newChatMutation(c.config, OpCreate)
	return &ChatCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Chat entities.
func (c *ChatClient) CreateBulk(builders ...*ChatCreate) *ChatCreateBulk {
	return &ChatCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Chat.
func (c *ChatClient) Update() *ChatUpdate {
	mutation := newChatMutation(c.config, OpUpdate)
	return &ChatUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChatClient) UpdateOne(ch *Chat) *ChatUpdateOne {
	mutation := newChatMutation(c.config, OpUpdateOne, withChat(ch))
	return &ChatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChatClient) UpdateOneID(id int64) *ChatUpdateOne {
	mutation := newChatMutation(c.config, OpUpdateOne, withChatID(id))
	return &ChatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Chat.
func (c *ChatClient) Delete() *ChatDelete {
	mutation := newChatMutation(c.config, OpDelete)
	return &ChatDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ChatClient) DeleteOne(ch *Chat) *ChatDeleteOne {
	return c.DeleteOneID(ch.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ChatClient) DeleteOneID(id int64) *ChatDeleteOne {
	builder := c.Delete().Where(chat.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChatDeleteOne{builder}
}

// Query returns a query builder for Chat.
func (c *ChatClient) Query() *ChatQuery {
	return &ChatQuery{
		config: c.config,
	}
}

// Get returns a Chat entity by its id.
func (c *ChatClient) Get(ctx context.Context, id int64) (*Chat, error) {
	return c.Query().Where(chat.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChatClient) GetX(ctx context.Context, id int64) *Chat {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Chat.
func (c *ChatClient) QueryOwner(ch *Chat) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ch.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chat.Table, chat.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, chat.OwnerTable, chat.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(ch.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMessages queries the messages edge of a Chat.
func (c *ChatClient) QueryMessages(ch *Chat) *MessageQuery {
	query := &MessageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ch.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(chat.Table, chat.FieldID, id),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, chat.MessagesTable, chat.MessagesColumn),
		)
		fromV = sqlgraph.Neighbors(ch.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ChatClient) Hooks() []Hook {
	return c.hooks.Chat
}

// MessageClient is a client for the Message schema.
type MessageClient struct {
	config
}

// NewMessageClient returns a client for the Message from the given config.
func NewMessageClient(c config) *MessageClient {
	return &MessageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `message.Hooks(f(g(h())))`.
func (c *MessageClient) Use(hooks ...Hook) {
	c.hooks.Message = append(c.hooks.Message, hooks...)
}

// Create returns a create builder for Message.
func (c *MessageClient) Create() *MessageCreate {
	mutation := newMessageMutation(c.config, OpCreate)
	return &MessageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Message entities.
func (c *MessageClient) CreateBulk(builders ...*MessageCreate) *MessageCreateBulk {
	return &MessageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Message.
func (c *MessageClient) Update() *MessageUpdate {
	mutation := newMessageMutation(c.config, OpUpdate)
	return &MessageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MessageClient) UpdateOne(m *Message) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessage(m))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MessageClient) UpdateOneID(id int64) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessageID(id))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Message.
func (c *MessageClient) Delete() *MessageDelete {
	mutation := newMessageMutation(c.config, OpDelete)
	return &MessageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MessageClient) DeleteOne(m *Message) *MessageDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MessageClient) DeleteOneID(id int64) *MessageDeleteOne {
	builder := c.Delete().Where(message.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MessageDeleteOne{builder}
}

// Query returns a query builder for Message.
func (c *MessageClient) Query() *MessageQuery {
	return &MessageQuery{
		config: c.config,
	}
}

// Get returns a Message entity by its id.
func (c *MessageClient) Get(ctx context.Context, id int64) (*Message, error) {
	return c.Query().Where(message.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MessageClient) GetX(ctx context.Context, id int64) *Message {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryChat queries the chat edge of a Message.
func (c *MessageClient) QueryChat(m *Message) *ChatQuery {
	query := &ChatQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, id),
			sqlgraph.To(chat.Table, chat.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, message.ChatTable, message.ChatColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a Message.
func (c *MessageClient) QueryUser(m *Message) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, message.UserTable, message.UserColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MessageClient) Hooks() []Hook {
	return c.hooks.Message
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int64) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int64) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int64) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int64) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryChat queries the chat edge of a User.
func (c *UserClient) QueryChat(u *User) *ChatQuery {
	query := &ChatQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(chat.Table, chat.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.ChatTable, user.ChatColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMetadata queries the metadata edge of a User.
func (c *UserClient) QueryMetadata(u *User) *UserMetadataQuery {
	query := &UserMetadataQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(usermetadata.Table, usermetadata.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.MetadataTable, user.MetadataColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMessages queries the messages edge of a User.
func (c *UserClient) QueryMessages(u *User) *MessageQuery {
	query := &MessageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.MessagesTable, user.MessagesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// UserMetadataClient is a client for the UserMetadata schema.
type UserMetadataClient struct {
	config
}

// NewUserMetadataClient returns a client for the UserMetadata from the given config.
func NewUserMetadataClient(c config) *UserMetadataClient {
	return &UserMetadataClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `usermetadata.Hooks(f(g(h())))`.
func (c *UserMetadataClient) Use(hooks ...Hook) {
	c.hooks.UserMetadata = append(c.hooks.UserMetadata, hooks...)
}

// Create returns a create builder for UserMetadata.
func (c *UserMetadataClient) Create() *UserMetadataCreate {
	mutation := newUserMetadataMutation(c.config, OpCreate)
	return &UserMetadataCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserMetadata entities.
func (c *UserMetadataClient) CreateBulk(builders ...*UserMetadataCreate) *UserMetadataCreateBulk {
	return &UserMetadataCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserMetadata.
func (c *UserMetadataClient) Update() *UserMetadataUpdate {
	mutation := newUserMetadataMutation(c.config, OpUpdate)
	return &UserMetadataUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserMetadataClient) UpdateOne(um *UserMetadata) *UserMetadataUpdateOne {
	mutation := newUserMetadataMutation(c.config, OpUpdateOne, withUserMetadata(um))
	return &UserMetadataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserMetadataClient) UpdateOneID(id int64) *UserMetadataUpdateOne {
	mutation := newUserMetadataMutation(c.config, OpUpdateOne, withUserMetadataID(id))
	return &UserMetadataUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserMetadata.
func (c *UserMetadataClient) Delete() *UserMetadataDelete {
	mutation := newUserMetadataMutation(c.config, OpDelete)
	return &UserMetadataDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserMetadataClient) DeleteOne(um *UserMetadata) *UserMetadataDeleteOne {
	return c.DeleteOneID(um.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserMetadataClient) DeleteOneID(id int64) *UserMetadataDeleteOne {
	builder := c.Delete().Where(usermetadata.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserMetadataDeleteOne{builder}
}

// Query returns a query builder for UserMetadata.
func (c *UserMetadataClient) Query() *UserMetadataQuery {
	return &UserMetadataQuery{
		config: c.config,
	}
}

// Get returns a UserMetadata entity by its id.
func (c *UserMetadataClient) Get(ctx context.Context, id int64) (*UserMetadata, error) {
	return c.Query().Where(usermetadata.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserMetadataClient) GetX(ctx context.Context, id int64) *UserMetadata {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a UserMetadata.
func (c *UserMetadataClient) QueryUser(um *UserMetadata) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := um.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(usermetadata.Table, usermetadata.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, usermetadata.UserTable, usermetadata.UserColumn),
		)
		fromV = sqlgraph.Neighbors(um.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserMetadataClient) Hooks() []Hook {
	return c.hooks.UserMetadata
}
