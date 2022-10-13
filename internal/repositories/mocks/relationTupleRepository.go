package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/Permify/permify/pkg/errors"
	base "github.com/Permify/permify/pkg/pb/base/v1"
	"github.com/Permify/permify/pkg/tuple"
)

// RelationTupleRepository is an autogenerated mock type for the RelationTupleRepository type
type RelationTupleRepository struct {
	mock.Mock
}

func (_m *RelationTupleRepository) Migrate() errors.Error {
	return nil
}

// ReverseQueryTuples -
func (_m *RelationTupleRepository) ReverseQueryTuples(ctx context.Context, entity string, relation string, subjectEntity string, subjectIDs []string, subjectRelation string) (tuple.ITupleIterator, errors.Error) {
	ret := _m.Called(entity, relation, subjectEntity, subjectIDs, subjectRelation)

	var r0 *tuple.TupleIterator
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, []string, string) *tuple.TupleIterator); ok {
		r0 = rf(ctx, entity, relation, subjectEntity, subjectIDs, subjectRelation)
	} else {
		r0 = ret.Get(0).(*tuple.TupleIterator)
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, []string, string) errors.Error); ok {
		r1 = rf(ctx, entity, relation, subjectEntity, subjectIDs, subjectRelation)
	} else {
		if e, ok := ret.Get(1).(errors.Error); ok {
			r1 = e
		} else {
			r1 = nil
		}
	}

	return r0, r1
}

// QueryTuples -
func (_m *RelationTupleRepository) QueryTuples(ctx context.Context, entityType string, entityID string, relation string) (tuple.ITupleIterator, errors.Error) {
	ret := _m.Called(entityType, entityID, relation)

	var r0 *tuple.TupleIterator
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *tuple.TupleIterator); ok {
		r0 = rf(ctx, entityType, entityID, relation)
	} else {
		r0 = ret.Get(0).(*tuple.TupleIterator)
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) errors.Error); ok {
		r1 = rf(ctx, entityType, entityID, relation)
	} else {
		if e, ok := ret.Get(1).(errors.Error); ok {
			r1 = e
		} else {
			r1 = nil
		}
	}

	return r0, r1
}

// Read -
func (_m *RelationTupleRepository) Read(ctx context.Context, filter *base.TupleFilter) (tuple.ITupleCollection, errors.Error) {
	ret := _m.Called(filter)

	var r0 *tuple.TupleCollection
	if rf, ok := ret.Get(0).(func(context.Context, *base.TupleFilter) *tuple.TupleCollection); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(*tuple.TupleCollection)
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, *base.TupleFilter) errors.Error); ok {
		r1 = rf(ctx, filter)
	} else {
		if e, ok := ret.Get(1).(errors.Error); ok {
			r1 = e
		} else {
			r1 = nil
		}
	}

	return r0, r1
}

// Write -
func (_m *RelationTupleRepository) Write(ctx context.Context, iterator tuple.ITupleIterator) errors.Error {
	ret := _m.Called(iterator)

	var r0 errors.Error
	if rf, ok := ret.Get(0).(func(context.Context, tuple.ITupleIterator) errors.Error); ok {
		r0 = rf(ctx, iterator)
	} else {
		if e, ok := ret.Get(1).(errors.Error); ok {
			r0 = e
		} else {
			r0 = nil
		}
	}

	return r0
}

// Delete -
func (_m *RelationTupleRepository) Delete(ctx context.Context, iterator tuple.ITupleIterator) errors.Error {
	ret := _m.Called(iterator)

	var r0 errors.Error
	if rf, ok := ret.Get(0).(func(context.Context, tuple.ITupleIterator) errors.Error); ok {
		r0 = rf(ctx, iterator)
	} else {
		if e, ok := ret.Get(1).(errors.Error); ok {
			r0 = e
		} else {
			r0 = nil
		}
	}

	return r0
}
