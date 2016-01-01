// Copyright 2015 Sevki <s@sevki.org>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ast defines build data structures.
package ast // import "sevki.org/build/ast"

import (
	"fmt"
	"io"

	"reflect"
)

var (
	targets map[string]reflect.Type
)

func init() {
	targets = make(map[string]reflect.Type)

}

// BuildFile defines a set of targets.
type BuildFile struct {
	Targets []Target
	Imports []string
	Vars    map[string]interface{}
}

// START1 OMIT
// Target is for implementing different build targets.
type Target interface {
	Build() error
	GetName() string
	GetDependencies() []string
	GetSource() string
	Reader() io.Reader
	Hash() []byte
}

// END1 OMIT

type Path string

func Register(name string, t interface{}) error {
	ty := reflect.TypeOf(t)
	if _, build := reflect.PtrTo(reflect.TypeOf(t)).MethodByName("Build"); !build {
		return fmt.Errorf("%s doesn't implement Build.", reflect.TypeOf(t))
	}
	targets[name] = ty

	return nil
}
func Get(name string) reflect.Type {
	if t, ok := targets[name]; ok {
		return t
	} else {
		return nil
	}
}
func GetFieldByTag(tn, tag string, p reflect.Type) *reflect.StructField {
	for i := 0; i < p.NumField(); i++ {
		f := p.Field(i)
		if f.Tag.Get(tn) == tag {
			return &f
		}
	}
	return nil
}
