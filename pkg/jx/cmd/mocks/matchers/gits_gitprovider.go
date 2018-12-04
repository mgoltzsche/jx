// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	gits "github.com/jenkins-x/jx/pkg/gits"
)

func AnyGitsGitProvider() gits.GitProvider {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(gits.GitProvider))(nil)).Elem()))
	var nullValue gits.GitProvider
	return nullValue
}

func EqGitsGitProvider(value gits.GitProvider) gits.GitProvider {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue gits.GitProvider
	return nullValue
}
