package moqi

import (
	"labix.org/v2/mgo"
)

type delegatingQuery struct {
	query *mgo.Query
}

type testingQuery struct {
	*delegatingQuery
	errorToReturn error
	once          bool
}

type mQuery interface {
	One(result interface{}) error
	Select(selector interface{}) mQuery
	Limit(n int) mQuery
}

func (t *testingQuery) One(result interface{}) error {
	if err := t.errorToReturn; err != nil {
		if t.once {
			t.errorToReturn = nil
		}
		return err
	}
	return t.delegatingQuery.One(result)
}

func (d delegatingQuery) One(result interface{}) error {
	return d.query.One(result)
}

func (d *delegatingQuery) Select(selector interface{}) mQuery {
	d.query = d.query.Select(selector)
	return d
}

func (d *delegatingQuery) Limit(n int) mQuery {
	d.query = d.query.Limit(n)
	return d
}
