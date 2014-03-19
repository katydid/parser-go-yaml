//  Copyright 2013 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main_test

import (
	"testing"

	protoparser "code.google.com/p/gogoprotobuf/parser"
	"code.google.com/p/gogoprotobuf/proto"
	"github.com/awalterschulze/katydid/exp/asm/ast"
	"github.com/awalterschulze/katydid/exp/asm/compiler"
	"github.com/awalterschulze/katydid/exp/asm/lexer"
	"github.com/awalterschulze/katydid/exp/asm/parser"
)

func test(t *testing.T, protoFilename string, m proto.Message, katydidStr string, positive bool) {
	fileDescriptorSet, err := protoparser.ParseFile(protoFilename, ".", "../../../../../../")
	if err != nil {
		panic(err)
	}
	p := parser.NewParser()
	r, err := p.Parse(lexer.NewLexer([]byte(katydidStr)))
	if err != nil {
		t.Fatalf(err.Error())
	}
	rules := r.(*ast.Rules)
	e, err := compiler.Compile(rules, fileDescriptorSet)
	if err != nil {
		panic(err)
	}
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	if match, err := e.Eval(data); err != nil {
		t.Errorf("Error: %v", err)
	} else if match != positive {
		t.Errorf("Expected a %v match from \n%v \non \n%v", positive, katydidStr, m)
	}
}
