// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	gits "github.com/jenkins-x/jx/pkg/gits"
)

func AnyPtrToGitsGitFileContent() *gits.GitFileContent {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*gits.GitFileContent))(nil)).Elem()))
	var nullValue *gits.GitFileContent
	return nullValue
}

func EqPtrToGitsGitFileContent(value *gits.GitFileContent) *gits.GitFileContent {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *gits.GitFileContent
	return nullValue
}
