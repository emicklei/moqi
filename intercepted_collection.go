package moqi

type testingCollection struct {
	*delegatingCollection
	errorToReturn error
	once          bool
}

func (f *testingCollection) Insert(docs ...interface{}) error {
	if err := f.errorToReturn; err != nil {
		if f.once {
			f.errorToReturn = nil
		}
		return err
	}
	return f.delegatingCollection.Insert(docs)
}
