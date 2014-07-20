package moqi

import "labix.org/v2/mgo"

type delegatingCollection struct {
	*mgo.Collection
}

// mCollection is the interface that describes the mgo.Collection exported functions
type mCollection interface {
	//Bulk() *mgo.Bulk
	Count() (n int, err error)
	Create(info *mgo.CollectionInfo) error
	DropCollection() error
	DropIndex(key ...string) error
	EnsureIndex(index mgo.Index) error
	EnsureIndexKey(key ...string) error
	Find(selector interface{}) *mgo.Query
	FindId(id interface{}) *mgo.Query
	Indexes() (indexes []mgo.Index, err error)
	Insert(docs ...interface{}) error
	Pipe(pipeline interface{}) *mgo.Pipe
	Remove(selector interface{}) error
	RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error)
	RemoveId(id interface{}) error
	Update(selector interface{}, update interface{}) error
	UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	UpdateId(id interface{}, update interface{}) error
	Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	UpsertId(id interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	With(s *mgo.Session) *mgo.Collection
}

//func (d delegatingCollection) Bulk() *mgo.Bulk {
//	return d.Collection.Bulk()
//}

func (d delegatingCollection) Count() (n int, err error) {
	return d.Collection.Count()
}
func (d delegatingCollection) Create(info *mgo.CollectionInfo) error {
	return d.Collection.Create(info)
}
func (d delegatingCollection) DropCollection() error {
	return d.Collection.DropCollection()
}
func (d delegatingCollection) DropIndex(key ...string) error {
	return d.Collection.DropIndex(key...)
}
func (d delegatingCollection) EnsureIndex(index mgo.Index) error {
	return d.Collection.EnsureIndex(index)
}
func (d delegatingCollection) EnsureIndexKey(key ...string) error {
	return d.Collection.EnsureIndexKey(key...)
}
func (d delegatingCollection) Find(selector interface{}) *mgo.Query {
	return d.Collection.Find(selector)
}
func (d delegatingCollection) FindId(id interface{}) *mgo.Query {
	return d.Collection.FindId(id)
}
func (d delegatingCollection) Indexes() (indexes []mgo.Index, err error) {
	return d.Collection.Indexes()
}
func (d delegatingCollection) Insert(docs ...interface{}) error {
	return d.Collection.Insert(docs)
}
func (d delegatingCollection) Pipe(pipeline interface{}) *mgo.Pipe {
	return d.Collection.Pipe(pipeline)
}
func (d delegatingCollection) Remove(selector interface{}) error {
	return d.Collection.Remove(selector)
}
func (d delegatingCollection) RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error) {
	return d.Collection.RemoveAll(selector)
}
func (d delegatingCollection) RemoveId(id interface{}) error {
	return d.Collection.RemoveId(id)
}
func (d delegatingCollection) Update(selector interface{}, update interface{}) error {
	return d.Collection.Update(selector, update)
}
func (d delegatingCollection) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	return d.Collection.UpdateAll(selector, update)
}
func (d delegatingCollection) UpdateId(id interface{}, update interface{}) error {
	return d.Collection.UpdateId(id, update)
}
func (d delegatingCollection) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	return d.Collection.Upsert(selector, update)
}
func (d delegatingCollection) UpsertId(id interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	return d.Collection.UpsertId(id, update)
}
func (d delegatingCollection) With(s *mgo.Session) *mgo.Collection {
	return d.Collection.With(s)
}
