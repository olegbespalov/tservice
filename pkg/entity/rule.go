package entity

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

//Rule response rule
type Rule struct {
	Method string
	Path   string

	Definition Definition

	Error    *Error
	Slowness *Slowness
}

//IsRegExp check if rule regexp
func (r Rule) isRegExp() bool {
	return strings.ContainsAny(r.Path, "()")
}

func (r Rule) match(path string) bool {
	match, err := regexp.MatchString("^"+r.Path+"$", path)
	if err != nil {
		return false
	}

	return match
}

//Fit check if response can be used for the request
func (r Rule) Fit(method, path string) bool {
	return (r.Path == path || (r.isRegExp() && r.match(path))) && (len(r.Method) == 0 || r.Method == method)
}

//ChoseDefinition chose between error an normal definition
func (r Rule) ChoseDefinition() Definition {
	if r.Error != nil && r.Error.Happened() {
		return r.Error.Definition
	}

	return r.Definition
}

//Wait how long we should wait for response
func (r Rule) Wait() time.Duration {
	wait := 0 * time.Nanosecond
	if r.Slowness != nil && r.Slowness.Happened() {
		wait = r.Slowness.Wait()
	}

	return wait
}

func (r Rule) String() string {
	return fmt.Sprintf("%s -  %s\nHeaders: %v", r.Method, r.Path, r.Definition.Headers)
}
