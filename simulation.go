package moqi

import (
	"errors"

	"labix.org/v2/mgo"
)

var ErrorSimulator = &ErrorSimulation{false, map[string]SimulatedError{}}

type ErrorSimulation struct {
	enabled     bool
	experiments map[string]SimulatedError
}

func (e ErrorSimulation) Enable(flag bool) {
	e.enabled = flag
}

func (e ErrorSimulation) NewQuery(label string, query *mgo.Query) mQuery {
	experiment, ok := e.experiments[label]
	delegator := &delegatingQuery{query}
	if !ok {
		return delegator
	}
	return &testingQuery{
		delegatingQuery: delegator,
		errorToReturn:   errors.New(experiment.ErrorMessage),
		once:            experiment.Once,
	}
}

func (e ErrorSimulation) NewCollection(label string, collection *mgo.Collection) mCollection {
	experiment, ok := e.experiments[label]
	delegator := &delegatingCollection{collection}
	if !ok {
		return delegator
	}
	return &testingCollection{
		delegatingCollection: delegator,
		errorToReturn:        errors.New(experiment.ErrorMessage),
		once:                 experiment.Once,
	}
}

func (e *ErrorSimulation) AddExperiment(label string, message string, once bool) {
	e.experiments[label] = SimulatedError{
		Label:        label,
		ErrorMessage: message,
		Once:         once,
	}
}

func (e *ErrorSimulation) Experiments() []SimulatedError {
	result := []SimulatedError{}
	for _, each := range e.experiments {
		result = append(result, each)
	}
	return result
}

type SimulatedError struct {
	Label        string
	ErrorMessage string
	Once         bool
}
