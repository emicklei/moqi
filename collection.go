package moqi

import "labix.org/v2/mgo"

type delegatingCollection struct {
	*mgo.Collection
}

type testingCollection struct {
	*delegatingCollection
	errorToReturn error
	once          bool
}

type mCollection interface {
	Insert(doc interface{}) error
	Find(selector interface{}) *mgo.Query
}

func (d delegatingCollection) Insert(doc interface{}) error {
	return d.Collection.Insert(doc)
}

func (d delegatingCollection) Find(selector interface{}) *mgo.Query {
	return d.Collection.Find(selector)
}

func (f *testingCollection) Insert(doc interface{}) error {
	if err := f.errorToReturn; err != nil {
		if f.once {
			f.errorToReturn = nil
		}
		return err
	}
	return f.delegatingCollection.Insert(doc)
}
