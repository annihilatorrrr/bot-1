// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/gotd/bot/internal/ent"
	"github.com/gotd/bot/internal/ent/lastchannelmessage"
	"github.com/gotd/bot/internal/ent/predicate"
	"github.com/gotd/bot/internal/ent/prnotification"
	"github.com/gotd/bot/internal/ent/telegramchannelstate"
	"github.com/gotd/bot/internal/ent/telegramsession"
	"github.com/gotd/bot/internal/ent/telegramuserstate"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next ent.Querier) ent.Querier {
	return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q ent.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The LastChannelMessageFunc type is an adapter to allow the use of ordinary function as a Querier.
type LastChannelMessageFunc func(context.Context, *ent.LastChannelMessageQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f LastChannelMessageFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.LastChannelMessageQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.LastChannelMessageQuery", q)
}

// The TraverseLastChannelMessage type is an adapter to allow the use of ordinary function as Traverser.
type TraverseLastChannelMessage func(context.Context, *ent.LastChannelMessageQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseLastChannelMessage) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseLastChannelMessage) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.LastChannelMessageQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.LastChannelMessageQuery", q)
}

// The PRNotificationFunc type is an adapter to allow the use of ordinary function as a Querier.
type PRNotificationFunc func(context.Context, *ent.PRNotificationQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f PRNotificationFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.PRNotificationQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.PRNotificationQuery", q)
}

// The TraversePRNotification type is an adapter to allow the use of ordinary function as Traverser.
type TraversePRNotification func(context.Context, *ent.PRNotificationQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePRNotification) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePRNotification) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PRNotificationQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.PRNotificationQuery", q)
}

// The TelegramChannelStateFunc type is an adapter to allow the use of ordinary function as a Querier.
type TelegramChannelStateFunc func(context.Context, *ent.TelegramChannelStateQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f TelegramChannelStateFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.TelegramChannelStateQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.TelegramChannelStateQuery", q)
}

// The TraverseTelegramChannelState type is an adapter to allow the use of ordinary function as Traverser.
type TraverseTelegramChannelState func(context.Context, *ent.TelegramChannelStateQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseTelegramChannelState) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseTelegramChannelState) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TelegramChannelStateQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.TelegramChannelStateQuery", q)
}

// The TelegramSessionFunc type is an adapter to allow the use of ordinary function as a Querier.
type TelegramSessionFunc func(context.Context, *ent.TelegramSessionQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f TelegramSessionFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.TelegramSessionQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.TelegramSessionQuery", q)
}

// The TraverseTelegramSession type is an adapter to allow the use of ordinary function as Traverser.
type TraverseTelegramSession func(context.Context, *ent.TelegramSessionQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseTelegramSession) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseTelegramSession) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TelegramSessionQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.TelegramSessionQuery", q)
}

// The TelegramUserStateFunc type is an adapter to allow the use of ordinary function as a Querier.
type TelegramUserStateFunc func(context.Context, *ent.TelegramUserStateQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f TelegramUserStateFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.TelegramUserStateQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.TelegramUserStateQuery", q)
}

// The TraverseTelegramUserState type is an adapter to allow the use of ordinary function as Traverser.
type TraverseTelegramUserState func(context.Context, *ent.TelegramUserStateQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseTelegramUserState) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseTelegramUserState) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TelegramUserStateQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.TelegramUserStateQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q ent.Query) (Query, error) {
	switch q := q.(type) {
	case *ent.LastChannelMessageQuery:
		return &query[*ent.LastChannelMessageQuery, predicate.LastChannelMessage, lastchannelmessage.OrderOption]{typ: ent.TypeLastChannelMessage, tq: q}, nil
	case *ent.PRNotificationQuery:
		return &query[*ent.PRNotificationQuery, predicate.PRNotification, prnotification.OrderOption]{typ: ent.TypePRNotification, tq: q}, nil
	case *ent.TelegramChannelStateQuery:
		return &query[*ent.TelegramChannelStateQuery, predicate.TelegramChannelState, telegramchannelstate.OrderOption]{typ: ent.TypeTelegramChannelState, tq: q}, nil
	case *ent.TelegramSessionQuery:
		return &query[*ent.TelegramSessionQuery, predicate.TelegramSession, telegramsession.OrderOption]{typ: ent.TypeTelegramSession, tq: q}, nil
	case *ent.TelegramUserStateQuery:
		return &query[*ent.TelegramUserStateQuery, predicate.TelegramUserState, telegramuserstate.OrderOption]{typ: ent.TypeTelegramUserState, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}