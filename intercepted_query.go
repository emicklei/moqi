package moqi

type testingQuery struct {
	*delegatingQuery
	errorToReturn error
	once          bool
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
