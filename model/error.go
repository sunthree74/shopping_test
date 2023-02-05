package model

import (
	"fmt"
	"strings"
)

type ErrNotFound struct {
	Datasource string
	Query      string
	Err        error
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("%s data not found: %s: %v", e.Datasource, e.Query, e.Err)
}

func (e *ErrNotFound) Unwrap() error {
	return e.Err
}

type ErrInitialization struct {
	Object          string
	ReqDependencies []string
	Err             error
}

func (e *ErrInitialization) Error() string {
	if e.ReqDependencies != nil {
		if len(e.ReqDependencies) > 1 {
			return fmt.Sprintf("cannot initialize %s: dependencies of %s is required", e.Object, strings.Join(e.ReqDependencies, ","))
		}

		return fmt.Sprintf("cannot initialize %s: dependency of %s is required", e.Object, e.ReqDependencies[0])
	}

	if e.Err != nil {
		return fmt.Sprintf("cannot initialize %s: %v", e.Object, e.Err)
	}

	return fmt.Sprintf("cannot initialize %s", e.Object)
}

func (e *ErrInitialization) Unwrap() error {
	return e.Err
}
