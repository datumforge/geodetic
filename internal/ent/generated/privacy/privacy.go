// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"

	"github.com/datumforge/geodetic/internal/ent/generated"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns a formatted wrapped Allow decision.
func Allowf(format string, a ...any) error {
	return privacy.Allowf(format, a...)
}

// Denyf returns a formatted wrapped Deny decision.
func Denyf(format string, a ...any) error {
	return privacy.Denyf(format, a...)
}

// Skipf returns a formatted wrapped Skip decision.
func Skipf(format string, a ...any) error {
	return privacy.Skipf(format, a...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
	// MutationRuleFunc type is an adapter which allows the use of
	// ordinary functions as mutation rules.
	MutationRuleFunc = privacy.MutationRuleFunc

	// QueryMutationRule is an interface which groups query and mutation rules.
	QueryMutationRule = privacy.QueryMutationRule
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, generated.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	return f(ctx, q)
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return privacy.AlwaysAllowRule()
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return privacy.AlwaysDenyRule()
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return privacy.ContextQueryMutationRule(eval)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op generated.Op) MutationRule {
	return privacy.OnMutationOperation(rule, op)
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op generated.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m generated.Mutation) error {
		return Denyf("generated/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The DatabaseQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DatabaseQueryRuleFunc func(context.Context, *generated.DatabaseQuery) error

// EvalQuery return f(ctx, q).
func (f DatabaseQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.DatabaseQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.DatabaseQuery", q)
}

// The DatabaseMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DatabaseMutationRuleFunc func(context.Context, *generated.DatabaseMutation) error

// EvalMutation calls f(ctx, m).
func (f DatabaseMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.DatabaseMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.DatabaseMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q generated.Query) (Filter, error) {
	switch q := q.(type) {
	case *generated.DatabaseQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("generated/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m generated.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *generated.DatabaseMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("generated/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
