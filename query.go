package moqi

import (
	"time"

	"labix.org/v2/mgo"
)

type delegatingQuery struct {
	query *mgo.Query
}

type mQuery interface {
	All(result interface{}) error
	Apply(change mgo.Change, result interface{}) (info *mgo.ChangeInfo, err error)
	Batch(n int) mQuery
	Count() (n int, err error)
	Distinct(key string, result interface{}) error
	Explain(result interface{}) error
	For(result interface{}, f func() error) error
	Hint(indexKey ...string) mQuery
	Iter() *mgo.Iter
	Limit(n int) mQuery
	LogReplay() mQuery
	MapReduce(job *mgo.MapReduce, result interface{}) (info *mgo.MapReduceInfo, err error)
	One(result interface{}) error
	Prefetch(p float64) mQuery
	Select(selector interface{}) mQuery
	Skip(n int) mQuery
	Snapshot() mQuery
	Sort(fields ...string) mQuery
	Tail(timeout time.Duration) *mgo.Iter
}

func (d delegatingQuery) All(result interface{}) error {
	return d.query.All(result)
}

func (d delegatingQuery) Apply(change mgo.Change, result interface{}) (info *mgo.ChangeInfo, err error) {
	return d.query.Apply(change, result)
}

func (d delegatingQuery) Batch(n int) mQuery {
	d.query = d.query.Batch(n)
	return d
}

func (d delegatingQuery) Count() (n int, err error) {
	return d.query.Count()
}

func (d delegatingQuery) Distinct(key string, result interface{}) error {
	return d.query.Distinct(key, result)
}

func (d delegatingQuery) Explain(result interface{}) error {
	return d.query.Explain(result)
}

func (d delegatingQuery) For(result interface{}, f func() error) error {
	return d.query.For(result, f)
}

func (d delegatingQuery) Hint(indexKey ...string) mQuery {
	d.query = d.query.Hint(indexKey...)
	return d
}

func (d delegatingQuery) Iter() *mgo.Iter {
	return d.query.Iter()
}

func (d delegatingQuery) Limit(n int) mQuery {
	d.query = d.query.Limit(n)
	return d
}

func (d delegatingQuery) LogReplay() mQuery {
	d.query = d.query.LogReplay()
	return d
}

func (d delegatingQuery) MapReduce(job *mgo.MapReduce, result interface{}) (info *mgo.MapReduceInfo, err error) {
	return d.query.MapReduce(job, result)
}

func (d delegatingQuery) One(result interface{}) error {
	return d.query.One(result)
}

func (d delegatingQuery) Prefetch(p float64) mQuery {
	d.query = d.query.Prefetch(p)
	return d
}

func (d delegatingQuery) Select(selector interface{}) mQuery {
	d.query = d.query.Select(selector)
	return d
}

func (d delegatingQuery) Skip(n int) mQuery {
	d.query = d.query.Skip(n)
	return d
}

func (d delegatingQuery) Snapshot() mQuery {
	d.query = d.query.Snapshot()
	return d
}

func (d delegatingQuery) Sort(fields ...string) mQuery {
	d.query = d.query.Sort(fields...)
	return d
}

func (d delegatingQuery) Tail(timeout time.Duration) *mgo.Iter {
	return d.query.Tail(timeout)
}
