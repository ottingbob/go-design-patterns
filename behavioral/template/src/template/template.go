package template

// Provides a user some way to execute code within an algorithm
// So it allows some parts to be implemented by the user while
// others may be handled by the software

// Common when one part of code is interchangeable where others
// need to stay consistent

// Such as an ETL layer. Data will flow in and out as normal
// but the modification may need to be different depending
// on the type of transformation we are doing

import "strings"

type MessageRetriever interface {
	Message() string
}

type Templater interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

type Template struct{}

func (t *Template) first() string {
	return "hello"
}

func (t *Template) third() string {
	return "template"
}

func (t *Template) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}

type AnonymousTemplate struct{}

func (a *AnonymousTemplate) first() string {
	return "hello"
}

func (a *AnonymousTemplate) third() string {
	return "template"
}

func (a *AnonymousTemplate) ExecuteAlgorithm(f func() string) string {
	return strings.Join([]string{a.first(), f(), a.third()}, " ")
}

type adapter struct {
	myFunc func() string
}

func (a *adapter) Message() string {
	if a.myFunc != nil {
		return a.myFunc()
	}

	return ""
}

func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &adapter{myFunc: f}
}