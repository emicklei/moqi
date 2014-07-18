package moqi

import "labix.org/v2/mgo"

func Enable(flag bool) {
	ErrorSimulator.enabled = flag
}

func AddExperiment(label string, message string, once bool) {
	ErrorSimulator.AddExperiment(label, message, once)
}

func Query(label string, query *mgo.Query) mQuery {
	if ErrorSimulator.enabled {
		return ErrorSimulator.NewQuery(label, query)
	}
	return &delegatingQuery{query}
}
