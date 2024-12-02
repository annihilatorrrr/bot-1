// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gotd/bot/internal/ent/predicate"
	"github.com/gotd/bot/internal/ent/prnotification"
)

// PRNotificationUpdate is the builder for updating PRNotification entities.
type PRNotificationUpdate struct {
	config
	hooks    []Hook
	mutation *PRNotificationMutation
}

// Where appends a list predicates to the PRNotificationUpdate builder.
func (pnu *PRNotificationUpdate) Where(ps ...predicate.PRNotification) *PRNotificationUpdate {
	pnu.mutation.Where(ps...)
	return pnu
}

// SetRepoID sets the "repo_id" field.
func (pnu *PRNotificationUpdate) SetRepoID(i int64) *PRNotificationUpdate {
	pnu.mutation.ResetRepoID()
	pnu.mutation.SetRepoID(i)
	return pnu
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (pnu *PRNotificationUpdate) SetNillableRepoID(i *int64) *PRNotificationUpdate {
	if i != nil {
		pnu.SetRepoID(*i)
	}
	return pnu
}

// AddRepoID adds i to the "repo_id" field.
func (pnu *PRNotificationUpdate) AddRepoID(i int64) *PRNotificationUpdate {
	pnu.mutation.AddRepoID(i)
	return pnu
}

// SetPullRequestID sets the "pull_request_id" field.
func (pnu *PRNotificationUpdate) SetPullRequestID(i int) *PRNotificationUpdate {
	pnu.mutation.ResetPullRequestID()
	pnu.mutation.SetPullRequestID(i)
	return pnu
}

// SetNillablePullRequestID sets the "pull_request_id" field if the given value is not nil.
func (pnu *PRNotificationUpdate) SetNillablePullRequestID(i *int) *PRNotificationUpdate {
	if i != nil {
		pnu.SetPullRequestID(*i)
	}
	return pnu
}

// AddPullRequestID adds i to the "pull_request_id" field.
func (pnu *PRNotificationUpdate) AddPullRequestID(i int) *PRNotificationUpdate {
	pnu.mutation.AddPullRequestID(i)
	return pnu
}

// SetPullRequestTitle sets the "pull_request_title" field.
func (pnu *PRNotificationUpdate) SetPullRequestTitle(s string) *PRNotificationUpdate {
	pnu.mutation.SetPullRequestTitle(s)
	return pnu
}

// SetNillablePullRequestTitle sets the "pull_request_title" field if the given value is not nil.
func (pnu *PRNotificationUpdate) SetNillablePullRequestTitle(s *string) *PRNotificationUpdate {
	if s != nil {
		pnu.SetPullRequestTitle(*s)
	}
	return pnu
}

// SetPullRequestBody sets the "pull_request_body" field.
func (pnu *PRNotificationUpdate) SetPullRequestBody(s string) *PRNotificationUpdate {
	pnu.mutation.SetPullRequestBody(s)
	return pnu
}

// SetNillablePullRequestBody sets the "pull_request_body" field if the given value is not nil.
func (pnu *PRNotificationUpdate) SetNillablePullRequestBody(s *string) *PRNotificationUpdate {
	if s != nil {
		pnu.SetPullRequestBody(*s)
	}
	return pnu
}

// SetPullRequestAuthorLogin sets the "pull_request_author_login" field.
func (pnu *PRNotificationUpdate) SetPullRequestAuthorLogin(s string) *PRNotificationUpdate {
	pnu.mutation.SetPullRequestAuthorLogin(s)
	return pnu
}

// SetNillablePullRequestAuthorLogin sets the "pull_request_author_login" field if the given value is not nil.
func (pnu *PRNotificationUpdate) SetNillablePullRequestAuthorLogin(s *string) *PRNotificationUpdate {
	if s != nil {
		pnu.SetPullRequestAuthorLogin(*s)
	}
	return pnu
}

// SetMessageID sets the "message_id" field.
func (pnu *PRNotificationUpdate) SetMessageID(i int) *PRNotificationUpdate {
	pnu.mutation.ResetMessageID()
	pnu.mutation.SetMessageID(i)
	return pnu
}

// SetNillableMessageID sets the "message_id" field if the given value is not nil.
func (pnu *PRNotificationUpdate) SetNillableMessageID(i *int) *PRNotificationUpdate {
	if i != nil {
		pnu.SetMessageID(*i)
	}
	return pnu
}

// AddMessageID adds i to the "message_id" field.
func (pnu *PRNotificationUpdate) AddMessageID(i int) *PRNotificationUpdate {
	pnu.mutation.AddMessageID(i)
	return pnu
}

// Mutation returns the PRNotificationMutation object of the builder.
func (pnu *PRNotificationUpdate) Mutation() *PRNotificationMutation {
	return pnu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pnu *PRNotificationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pnu.sqlSave, pnu.mutation, pnu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pnu *PRNotificationUpdate) SaveX(ctx context.Context) int {
	affected, err := pnu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pnu *PRNotificationUpdate) Exec(ctx context.Context) error {
	_, err := pnu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pnu *PRNotificationUpdate) ExecX(ctx context.Context) {
	if err := pnu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pnu *PRNotificationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(prnotification.Table, prnotification.Columns, sqlgraph.NewFieldSpec(prnotification.FieldID, field.TypeInt))
	if ps := pnu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pnu.mutation.RepoID(); ok {
		_spec.SetField(prnotification.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := pnu.mutation.AddedRepoID(); ok {
		_spec.AddField(prnotification.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := pnu.mutation.PullRequestID(); ok {
		_spec.SetField(prnotification.FieldPullRequestID, field.TypeInt, value)
	}
	if value, ok := pnu.mutation.AddedPullRequestID(); ok {
		_spec.AddField(prnotification.FieldPullRequestID, field.TypeInt, value)
	}
	if value, ok := pnu.mutation.PullRequestTitle(); ok {
		_spec.SetField(prnotification.FieldPullRequestTitle, field.TypeString, value)
	}
	if value, ok := pnu.mutation.PullRequestBody(); ok {
		_spec.SetField(prnotification.FieldPullRequestBody, field.TypeString, value)
	}
	if value, ok := pnu.mutation.PullRequestAuthorLogin(); ok {
		_spec.SetField(prnotification.FieldPullRequestAuthorLogin, field.TypeString, value)
	}
	if value, ok := pnu.mutation.MessageID(); ok {
		_spec.SetField(prnotification.FieldMessageID, field.TypeInt, value)
	}
	if value, ok := pnu.mutation.AddedMessageID(); ok {
		_spec.AddField(prnotification.FieldMessageID, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pnu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prnotification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pnu.mutation.done = true
	return n, nil
}

// PRNotificationUpdateOne is the builder for updating a single PRNotification entity.
type PRNotificationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PRNotificationMutation
}

// SetRepoID sets the "repo_id" field.
func (pnuo *PRNotificationUpdateOne) SetRepoID(i int64) *PRNotificationUpdateOne {
	pnuo.mutation.ResetRepoID()
	pnuo.mutation.SetRepoID(i)
	return pnuo
}

// SetNillableRepoID sets the "repo_id" field if the given value is not nil.
func (pnuo *PRNotificationUpdateOne) SetNillableRepoID(i *int64) *PRNotificationUpdateOne {
	if i != nil {
		pnuo.SetRepoID(*i)
	}
	return pnuo
}

// AddRepoID adds i to the "repo_id" field.
func (pnuo *PRNotificationUpdateOne) AddRepoID(i int64) *PRNotificationUpdateOne {
	pnuo.mutation.AddRepoID(i)
	return pnuo
}

// SetPullRequestID sets the "pull_request_id" field.
func (pnuo *PRNotificationUpdateOne) SetPullRequestID(i int) *PRNotificationUpdateOne {
	pnuo.mutation.ResetPullRequestID()
	pnuo.mutation.SetPullRequestID(i)
	return pnuo
}

// SetNillablePullRequestID sets the "pull_request_id" field if the given value is not nil.
func (pnuo *PRNotificationUpdateOne) SetNillablePullRequestID(i *int) *PRNotificationUpdateOne {
	if i != nil {
		pnuo.SetPullRequestID(*i)
	}
	return pnuo
}

// AddPullRequestID adds i to the "pull_request_id" field.
func (pnuo *PRNotificationUpdateOne) AddPullRequestID(i int) *PRNotificationUpdateOne {
	pnuo.mutation.AddPullRequestID(i)
	return pnuo
}

// SetPullRequestTitle sets the "pull_request_title" field.
func (pnuo *PRNotificationUpdateOne) SetPullRequestTitle(s string) *PRNotificationUpdateOne {
	pnuo.mutation.SetPullRequestTitle(s)
	return pnuo
}

// SetNillablePullRequestTitle sets the "pull_request_title" field if the given value is not nil.
func (pnuo *PRNotificationUpdateOne) SetNillablePullRequestTitle(s *string) *PRNotificationUpdateOne {
	if s != nil {
		pnuo.SetPullRequestTitle(*s)
	}
	return pnuo
}

// SetPullRequestBody sets the "pull_request_body" field.
func (pnuo *PRNotificationUpdateOne) SetPullRequestBody(s string) *PRNotificationUpdateOne {
	pnuo.mutation.SetPullRequestBody(s)
	return pnuo
}

// SetNillablePullRequestBody sets the "pull_request_body" field if the given value is not nil.
func (pnuo *PRNotificationUpdateOne) SetNillablePullRequestBody(s *string) *PRNotificationUpdateOne {
	if s != nil {
		pnuo.SetPullRequestBody(*s)
	}
	return pnuo
}

// SetPullRequestAuthorLogin sets the "pull_request_author_login" field.
func (pnuo *PRNotificationUpdateOne) SetPullRequestAuthorLogin(s string) *PRNotificationUpdateOne {
	pnuo.mutation.SetPullRequestAuthorLogin(s)
	return pnuo
}

// SetNillablePullRequestAuthorLogin sets the "pull_request_author_login" field if the given value is not nil.
func (pnuo *PRNotificationUpdateOne) SetNillablePullRequestAuthorLogin(s *string) *PRNotificationUpdateOne {
	if s != nil {
		pnuo.SetPullRequestAuthorLogin(*s)
	}
	return pnuo
}

// SetMessageID sets the "message_id" field.
func (pnuo *PRNotificationUpdateOne) SetMessageID(i int) *PRNotificationUpdateOne {
	pnuo.mutation.ResetMessageID()
	pnuo.mutation.SetMessageID(i)
	return pnuo
}

// SetNillableMessageID sets the "message_id" field if the given value is not nil.
func (pnuo *PRNotificationUpdateOne) SetNillableMessageID(i *int) *PRNotificationUpdateOne {
	if i != nil {
		pnuo.SetMessageID(*i)
	}
	return pnuo
}

// AddMessageID adds i to the "message_id" field.
func (pnuo *PRNotificationUpdateOne) AddMessageID(i int) *PRNotificationUpdateOne {
	pnuo.mutation.AddMessageID(i)
	return pnuo
}

// Mutation returns the PRNotificationMutation object of the builder.
func (pnuo *PRNotificationUpdateOne) Mutation() *PRNotificationMutation {
	return pnuo.mutation
}

// Where appends a list predicates to the PRNotificationUpdate builder.
func (pnuo *PRNotificationUpdateOne) Where(ps ...predicate.PRNotification) *PRNotificationUpdateOne {
	pnuo.mutation.Where(ps...)
	return pnuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pnuo *PRNotificationUpdateOne) Select(field string, fields ...string) *PRNotificationUpdateOne {
	pnuo.fields = append([]string{field}, fields...)
	return pnuo
}

// Save executes the query and returns the updated PRNotification entity.
func (pnuo *PRNotificationUpdateOne) Save(ctx context.Context) (*PRNotification, error) {
	return withHooks(ctx, pnuo.sqlSave, pnuo.mutation, pnuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pnuo *PRNotificationUpdateOne) SaveX(ctx context.Context) *PRNotification {
	node, err := pnuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pnuo *PRNotificationUpdateOne) Exec(ctx context.Context) error {
	_, err := pnuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pnuo *PRNotificationUpdateOne) ExecX(ctx context.Context) {
	if err := pnuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pnuo *PRNotificationUpdateOne) sqlSave(ctx context.Context) (_node *PRNotification, err error) {
	_spec := sqlgraph.NewUpdateSpec(prnotification.Table, prnotification.Columns, sqlgraph.NewFieldSpec(prnotification.FieldID, field.TypeInt))
	id, ok := pnuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PRNotification.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pnuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, prnotification.FieldID)
		for _, f := range fields {
			if !prnotification.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != prnotification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pnuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pnuo.mutation.RepoID(); ok {
		_spec.SetField(prnotification.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := pnuo.mutation.AddedRepoID(); ok {
		_spec.AddField(prnotification.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := pnuo.mutation.PullRequestID(); ok {
		_spec.SetField(prnotification.FieldPullRequestID, field.TypeInt, value)
	}
	if value, ok := pnuo.mutation.AddedPullRequestID(); ok {
		_spec.AddField(prnotification.FieldPullRequestID, field.TypeInt, value)
	}
	if value, ok := pnuo.mutation.PullRequestTitle(); ok {
		_spec.SetField(prnotification.FieldPullRequestTitle, field.TypeString, value)
	}
	if value, ok := pnuo.mutation.PullRequestBody(); ok {
		_spec.SetField(prnotification.FieldPullRequestBody, field.TypeString, value)
	}
	if value, ok := pnuo.mutation.PullRequestAuthorLogin(); ok {
		_spec.SetField(prnotification.FieldPullRequestAuthorLogin, field.TypeString, value)
	}
	if value, ok := pnuo.mutation.MessageID(); ok {
		_spec.SetField(prnotification.FieldMessageID, field.TypeInt, value)
	}
	if value, ok := pnuo.mutation.AddedMessageID(); ok {
		_spec.AddField(prnotification.FieldMessageID, field.TypeInt, value)
	}
	_node = &PRNotification{config: pnuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pnuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prnotification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pnuo.mutation.done = true
	return _node, nil
}
