// Code generated by protoc-gen-gogo.
// source: relapse.proto
// DO NOT EDIT!

/*
	Package relapse is a generated protocol buffer package.

	It is generated from these files:
		relapse.proto

	It has these top-level messages:
		Grammar
		PatternDecl
		NameExpr
		Name
		AnyName
		AnyNameExcept
		NameChoice
		Pattern
		Empty
		EmptySet
		TreeNode
		WithSomeTreeNode
		LeafNode
		Concat
		Or
		WithSomeOr
		And
		WithSomeAnd
		ZeroOrMore
		Reference
		Not
		ZAny
*/
package relapse

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"
import expr "github.com/katydid/katydid/expr/ast"

import fmt "fmt"
import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Grammar struct {
	TopPattern   *Pattern       `protobuf:"bytes,1,opt" json:"TopPattern,omitempty"`
	PatternDecls []*PatternDecl `protobuf:"bytes,2,rep" json:"PatternDecls,omitempty"`
	After        *expr.Space    `protobuf:"bytes,3,opt" json:"After,omitempty"`
}

func (m *Grammar) Reset()      { *m = Grammar{} }
func (*Grammar) ProtoMessage() {}

func (m *Grammar) GetTopPattern() *Pattern {
	if m != nil {
		return m.TopPattern
	}
	return nil
}

func (m *Grammar) GetPatternDecls() []*PatternDecl {
	if m != nil {
		return m.PatternDecls
	}
	return nil
}

func (m *Grammar) GetAfter() *expr.Space {
	if m != nil {
		return m.After
	}
	return nil
}

type PatternDecl struct {
	At      *expr.Keyword `protobuf:"bytes,1,opt" json:"At,omitempty"`
	Before  *expr.Space   `protobuf:"bytes,2,opt" json:"Before,omitempty"`
	Name    string        `protobuf:"bytes,3,opt" json:"Name"`
	Eq      *expr.Keyword `protobuf:"bytes,4,opt" json:"Eq,omitempty"`
	Pattern *Pattern      `protobuf:"bytes,5,opt" json:"Pattern,omitempty"`
}

func (m *PatternDecl) Reset()      { *m = PatternDecl{} }
func (*PatternDecl) ProtoMessage() {}

func (m *PatternDecl) GetAt() *expr.Keyword {
	if m != nil {
		return m.At
	}
	return nil
}

func (m *PatternDecl) GetBefore() *expr.Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *PatternDecl) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PatternDecl) GetEq() *expr.Keyword {
	if m != nil {
		return m.Eq
	}
	return nil
}

func (m *PatternDecl) GetPattern() *Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

type NameExpr struct {
	Name          *Name          `protobuf:"bytes,1,opt" json:"Name,omitempty"`
	AnyName       *AnyName       `protobuf:"bytes,2,opt" json:"AnyName,omitempty"`
	AnyNameExcept *AnyNameExcept `protobuf:"bytes,3,opt" json:"AnyNameExcept,omitempty"`
	NameChoice    *NameChoice    `protobuf:"bytes,4,opt" json:"NameChoice,omitempty"`
}

func (m *NameExpr) Reset()      { *m = NameExpr{} }
func (*NameExpr) ProtoMessage() {}

func (m *NameExpr) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *NameExpr) GetAnyName() *AnyName {
	if m != nil {
		return m.AnyName
	}
	return nil
}

func (m *NameExpr) GetAnyNameExcept() *AnyNameExcept {
	if m != nil {
		return m.AnyNameExcept
	}
	return nil
}

func (m *NameExpr) GetNameChoice() *NameChoice {
	if m != nil {
		return m.NameChoice
	}
	return nil
}

type Name struct {
	Before *expr.Space `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	Name   string      `protobuf:"bytes,2,opt" json:"Name"`
}

func (m *Name) Reset()      { *m = Name{} }
func (*Name) ProtoMessage() {}

func (m *Name) GetBefore() *expr.Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *Name) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AnyName struct {
	Underscore *expr.Keyword `protobuf:"bytes,1,opt" json:"Underscore,omitempty"`
}

func (m *AnyName) Reset()      { *m = AnyName{} }
func (*AnyName) ProtoMessage() {}

func (m *AnyName) GetUnderscore() *expr.Keyword {
	if m != nil {
		return m.Underscore
	}
	return nil
}

type AnyNameExcept struct {
	Exclamation *expr.Keyword `protobuf:"bytes,1,opt" json:"Exclamation,omitempty"`
	OpenParen   *expr.Keyword `protobuf:"bytes,2,opt" json:"OpenParen,omitempty"`
	Except      *NameExpr     `protobuf:"bytes,3,opt" json:"Except,omitempty"`
	CloseParen  *expr.Keyword `protobuf:"bytes,4,opt" json:"CloseParen,omitempty"`
}

func (m *AnyNameExcept) Reset()      { *m = AnyNameExcept{} }
func (*AnyNameExcept) ProtoMessage() {}

func (m *AnyNameExcept) GetExclamation() *expr.Keyword {
	if m != nil {
		return m.Exclamation
	}
	return nil
}

func (m *AnyNameExcept) GetOpenParen() *expr.Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *AnyNameExcept) GetExcept() *NameExpr {
	if m != nil {
		return m.Except
	}
	return nil
}

func (m *AnyNameExcept) GetCloseParen() *expr.Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

type NameChoice struct {
	OpenParen  *expr.Keyword `protobuf:"bytes,1,opt" json:"OpenParen,omitempty"`
	Left       *NameExpr     `protobuf:"bytes,2,opt" json:"Left,omitempty"`
	Pipe       *expr.Keyword `protobuf:"bytes,3,opt" json:"Pipe,omitempty"`
	Right      *NameExpr     `protobuf:"bytes,4,opt" json:"Right,omitempty"`
	CloseParen *expr.Keyword `protobuf:"bytes,5,opt" json:"CloseParen,omitempty"`
}

func (m *NameChoice) Reset()      { *m = NameChoice{} }
func (*NameChoice) ProtoMessage() {}

func (m *NameChoice) GetOpenParen() *expr.Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *NameChoice) GetLeft() *NameExpr {
	if m != nil {
		return m.Left
	}
	return nil
}

func (m *NameChoice) GetPipe() *expr.Keyword {
	if m != nil {
		return m.Pipe
	}
	return nil
}

func (m *NameChoice) GetRight() *NameExpr {
	if m != nil {
		return m.Right
	}
	return nil
}

func (m *NameChoice) GetCloseParen() *expr.Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

type Pattern struct {
	Empty            *Empty            `protobuf:"bytes,1,opt" json:"Empty,omitempty"`
	EmptySet         *EmptySet         `protobuf:"bytes,2,opt" json:"EmptySet,omitempty"`
	TreeNode         *TreeNode         `protobuf:"bytes,3,opt" json:"TreeNode,omitempty"`
	LeafNode         *LeafNode         `protobuf:"bytes,4,opt" json:"LeafNode,omitempty"`
	Concat           *Concat           `protobuf:"bytes,5,opt" json:"Concat,omitempty"`
	Or               *Or               `protobuf:"bytes,6,opt" json:"Or,omitempty"`
	And              *And              `protobuf:"bytes,7,opt" json:"And,omitempty"`
	ZeroOrMore       *ZeroOrMore       `protobuf:"bytes,8,opt" json:"ZeroOrMore,omitempty"`
	Reference        *Reference        `protobuf:"bytes,9,opt" json:"Reference,omitempty"`
	Not              *Not              `protobuf:"bytes,10,opt" json:"Not,omitempty"`
	ZAny             *ZAny             `protobuf:"bytes,11,opt" json:"ZAny,omitempty"`
	WithSomeTreeNode *WithSomeTreeNode `protobuf:"bytes,12,opt" json:"WithSomeTreeNode,omitempty"`
}

func (m *Pattern) Reset()      { *m = Pattern{} }
func (*Pattern) ProtoMessage() {}

func (m *Pattern) GetEmpty() *Empty {
	if m != nil {
		return m.Empty
	}
	return nil
}

func (m *Pattern) GetEmptySet() *EmptySet {
	if m != nil {
		return m.EmptySet
	}
	return nil
}

func (m *Pattern) GetTreeNode() *TreeNode {
	if m != nil {
		return m.TreeNode
	}
	return nil
}

func (m *Pattern) GetLeafNode() *LeafNode {
	if m != nil {
		return m.LeafNode
	}
	return nil
}

func (m *Pattern) GetConcat() *Concat {
	if m != nil {
		return m.Concat
	}
	return nil
}

func (m *Pattern) GetOr() *Or {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *Pattern) GetAnd() *And {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *Pattern) GetZeroOrMore() *ZeroOrMore {
	if m != nil {
		return m.ZeroOrMore
	}
	return nil
}

func (m *Pattern) GetReference() *Reference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *Pattern) GetNot() *Not {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *Pattern) GetZAny() *ZAny {
	if m != nil {
		return m.ZAny
	}
	return nil
}

func (m *Pattern) GetWithSomeTreeNode() *WithSomeTreeNode {
	if m != nil {
		return m.WithSomeTreeNode
	}
	return nil
}

type Empty struct {
	Empty *expr.Keyword `protobuf:"bytes,1,opt" json:"Empty,omitempty"`
}

func (m *Empty) Reset()      { *m = Empty{} }
func (*Empty) ProtoMessage() {}

func (m *Empty) GetEmpty() *expr.Keyword {
	if m != nil {
		return m.Empty
	}
	return nil
}

type EmptySet struct {
	EmptySet *expr.Keyword `protobuf:"bytes,1,opt" json:"EmptySet,omitempty"`
}

func (m *EmptySet) Reset()      { *m = EmptySet{} }
func (*EmptySet) ProtoMessage() {}

func (m *EmptySet) GetEmptySet() *expr.Keyword {
	if m != nil {
		return m.EmptySet
	}
	return nil
}

type TreeNode struct {
	Name    *NameExpr     `protobuf:"bytes,1,opt" json:"Name,omitempty"`
	Colon   *expr.Keyword `protobuf:"bytes,2,opt" json:"Colon,omitempty"`
	Pattern *Pattern      `protobuf:"bytes,3,opt" json:"Pattern,omitempty"`
}

func (m *TreeNode) Reset()      { *m = TreeNode{} }
func (*TreeNode) ProtoMessage() {}

func (m *TreeNode) GetName() *NameExpr {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *TreeNode) GetColon() *expr.Keyword {
	if m != nil {
		return m.Colon
	}
	return nil
}

func (m *TreeNode) GetPattern() *Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

type WithSomeTreeNode struct {
	Dot     *expr.Keyword `protobuf:"bytes,1,opt" json:"Dot,omitempty"`
	Pattern *Pattern      `protobuf:"bytes,2,opt" json:"Pattern,omitempty"`
}

func (m *WithSomeTreeNode) Reset()      { *m = WithSomeTreeNode{} }
func (*WithSomeTreeNode) ProtoMessage() {}

func (m *WithSomeTreeNode) GetDot() *expr.Keyword {
	if m != nil {
		return m.Dot
	}
	return nil
}

func (m *WithSomeTreeNode) GetPattern() *Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

type LeafNode struct {
	RightArrow *expr.Keyword `protobuf:"bytes,1,opt" json:"RightArrow,omitempty"`
	Expr       *expr.Expr    `protobuf:"bytes,2,opt" json:"Expr,omitempty"`
}

func (m *LeafNode) Reset()      { *m = LeafNode{} }
func (*LeafNode) ProtoMessage() {}

func (m *LeafNode) GetRightArrow() *expr.Keyword {
	if m != nil {
		return m.RightArrow
	}
	return nil
}

func (m *LeafNode) GetExpr() *expr.Expr {
	if m != nil {
		return m.Expr
	}
	return nil
}

type Concat struct {
	OpenBracket  *expr.Keyword `protobuf:"bytes,2,opt" json:"OpenBracket,omitempty"`
	LeftPattern  *Pattern      `protobuf:"bytes,3,opt" json:"LeftPattern,omitempty"`
	Comma        *expr.Keyword `protobuf:"bytes,4,opt" json:"Comma,omitempty"`
	RightPattern *Pattern      `protobuf:"bytes,5,opt" json:"RightPattern,omitempty"`
	CloseBracket *expr.Keyword `protobuf:"bytes,6,opt" json:"CloseBracket,omitempty"`
}

func (m *Concat) Reset()      { *m = Concat{} }
func (*Concat) ProtoMessage() {}

func (m *Concat) GetOpenBracket() *expr.Keyword {
	if m != nil {
		return m.OpenBracket
	}
	return nil
}

func (m *Concat) GetLeftPattern() *Pattern {
	if m != nil {
		return m.LeftPattern
	}
	return nil
}

func (m *Concat) GetComma() *expr.Keyword {
	if m != nil {
		return m.Comma
	}
	return nil
}

func (m *Concat) GetRightPattern() *Pattern {
	if m != nil {
		return m.RightPattern
	}
	return nil
}

func (m *Concat) GetCloseBracket() *expr.Keyword {
	if m != nil {
		return m.CloseBracket
	}
	return nil
}

type Or struct {
	OpenParen    *expr.Keyword `protobuf:"bytes,1,opt" json:"OpenParen,omitempty"`
	LeftPattern  *Pattern      `protobuf:"bytes,2,opt" json:"LeftPattern,omitempty"`
	Pipe         *expr.Keyword `protobuf:"bytes,3,opt" json:"Pipe,omitempty"`
	RightPattern *Pattern      `protobuf:"bytes,4,opt" json:"RightPattern,omitempty"`
	CloseParen   *expr.Keyword `protobuf:"bytes,5,opt" json:"CloseParen,omitempty"`
}

func (m *Or) Reset()      { *m = Or{} }
func (*Or) ProtoMessage() {}

func (m *Or) GetOpenParen() *expr.Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *Or) GetLeftPattern() *Pattern {
	if m != nil {
		return m.LeftPattern
	}
	return nil
}

func (m *Or) GetPipe() *expr.Keyword {
	if m != nil {
		return m.Pipe
	}
	return nil
}

func (m *Or) GetRightPattern() *Pattern {
	if m != nil {
		return m.RightPattern
	}
	return nil
}

func (m *Or) GetCloseParen() *expr.Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

type WithSomeOr struct {
	OpenCurly    *expr.Keyword `protobuf:"bytes,1,opt" json:"OpenCurly,omitempty"`
	LeftPattern  *Pattern      `protobuf:"bytes,2,opt" json:"LeftPattern,omitempty"`
	Pipe         *expr.Keyword `protobuf:"bytes,3,opt" json:"Pipe,omitempty"`
	RightPattern *Pattern      `protobuf:"bytes,4,opt" json:"RightPattern,omitempty"`
	CloseCurly   *expr.Keyword `protobuf:"bytes,5,opt" json:"CloseCurly,omitempty"`
}

func (m *WithSomeOr) Reset()      { *m = WithSomeOr{} }
func (*WithSomeOr) ProtoMessage() {}

func (m *WithSomeOr) GetOpenCurly() *expr.Keyword {
	if m != nil {
		return m.OpenCurly
	}
	return nil
}

func (m *WithSomeOr) GetLeftPattern() *Pattern {
	if m != nil {
		return m.LeftPattern
	}
	return nil
}

func (m *WithSomeOr) GetPipe() *expr.Keyword {
	if m != nil {
		return m.Pipe
	}
	return nil
}

func (m *WithSomeOr) GetRightPattern() *Pattern {
	if m != nil {
		return m.RightPattern
	}
	return nil
}

func (m *WithSomeOr) GetCloseCurly() *expr.Keyword {
	if m != nil {
		return m.CloseCurly
	}
	return nil
}

type And struct {
	OpenParen    *expr.Keyword `protobuf:"bytes,1,opt" json:"OpenParen,omitempty"`
	LeftPattern  *Pattern      `protobuf:"bytes,2,opt" json:"LeftPattern,omitempty"`
	Ampersand    *expr.Keyword `protobuf:"bytes,3,opt" json:"Ampersand,omitempty"`
	RightPattern *Pattern      `protobuf:"bytes,4,opt" json:"RightPattern,omitempty"`
	CloseParen   *expr.Keyword `protobuf:"bytes,5,opt" json:"CloseParen,omitempty"`
}

func (m *And) Reset()      { *m = And{} }
func (*And) ProtoMessage() {}

func (m *And) GetOpenParen() *expr.Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *And) GetLeftPattern() *Pattern {
	if m != nil {
		return m.LeftPattern
	}
	return nil
}

func (m *And) GetAmpersand() *expr.Keyword {
	if m != nil {
		return m.Ampersand
	}
	return nil
}

func (m *And) GetRightPattern() *Pattern {
	if m != nil {
		return m.RightPattern
	}
	return nil
}

func (m *And) GetCloseParen() *expr.Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

type WithSomeAnd struct {
	OpenCurly    *expr.Keyword `protobuf:"bytes,1,opt" json:"OpenCurly,omitempty"`
	LeftPattern  *Pattern      `protobuf:"bytes,2,opt" json:"LeftPattern,omitempty"`
	Ampersand    *expr.Keyword `protobuf:"bytes,3,opt" json:"Ampersand,omitempty"`
	RightPattern *Pattern      `protobuf:"bytes,4,opt" json:"RightPattern,omitempty"`
	CloseCurly   *expr.Keyword `protobuf:"bytes,5,opt" json:"CloseCurly,omitempty"`
}

func (m *WithSomeAnd) Reset()      { *m = WithSomeAnd{} }
func (*WithSomeAnd) ProtoMessage() {}

func (m *WithSomeAnd) GetOpenCurly() *expr.Keyword {
	if m != nil {
		return m.OpenCurly
	}
	return nil
}

func (m *WithSomeAnd) GetLeftPattern() *Pattern {
	if m != nil {
		return m.LeftPattern
	}
	return nil
}

func (m *WithSomeAnd) GetAmpersand() *expr.Keyword {
	if m != nil {
		return m.Ampersand
	}
	return nil
}

func (m *WithSomeAnd) GetRightPattern() *Pattern {
	if m != nil {
		return m.RightPattern
	}
	return nil
}

func (m *WithSomeAnd) GetCloseCurly() *expr.Keyword {
	if m != nil {
		return m.CloseCurly
	}
	return nil
}

type ZeroOrMore struct {
	OpenParen  *expr.Keyword `protobuf:"bytes,1,opt" json:"OpenParen,omitempty"`
	Pattern    *Pattern      `protobuf:"bytes,2,opt" json:"Pattern,omitempty"`
	CloseParen *expr.Keyword `protobuf:"bytes,3,opt" json:"CloseParen,omitempty"`
	Star       *expr.Keyword `protobuf:"bytes,4,opt" json:"Star,omitempty"`
}

func (m *ZeroOrMore) Reset()      { *m = ZeroOrMore{} }
func (*ZeroOrMore) ProtoMessage() {}

func (m *ZeroOrMore) GetOpenParen() *expr.Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *ZeroOrMore) GetPattern() *Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *ZeroOrMore) GetCloseParen() *expr.Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

func (m *ZeroOrMore) GetStar() *expr.Keyword {
	if m != nil {
		return m.Star
	}
	return nil
}

type Reference struct {
	HashTag *expr.Keyword `protobuf:"bytes,1,opt" json:"HashTag,omitempty"`
	Name    string        `protobuf:"bytes,2,opt" json:"Name"`
}

func (m *Reference) Reset()      { *m = Reference{} }
func (*Reference) ProtoMessage() {}

func (m *Reference) GetHashTag() *expr.Keyword {
	if m != nil {
		return m.HashTag
	}
	return nil
}

func (m *Reference) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Not struct {
	Exclamation *expr.Keyword `protobuf:"bytes,1,opt" json:"Exclamation,omitempty"`
	OpenParen   *expr.Keyword `protobuf:"bytes,2,opt" json:"OpenParen,omitempty"`
	Pattern     *Pattern      `protobuf:"bytes,3,opt" json:"Pattern,omitempty"`
	CloseParen  *expr.Keyword `protobuf:"bytes,4,opt" json:"CloseParen,omitempty"`
}

func (m *Not) Reset()      { *m = Not{} }
func (*Not) ProtoMessage() {}

func (m *Not) GetExclamation() *expr.Keyword {
	if m != nil {
		return m.Exclamation
	}
	return nil
}

func (m *Not) GetOpenParen() *expr.Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *Not) GetPattern() *Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *Not) GetCloseParen() *expr.Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

type ZAny struct {
	Star *expr.Keyword `protobuf:"bytes,1,opt" json:"Star,omitempty"`
}

func (m *ZAny) Reset()      { *m = ZAny{} }
func (*ZAny) ProtoMessage() {}

func (m *ZAny) GetStar() *expr.Keyword {
	if m != nil {
		return m.Star
	}
	return nil
}

func init() {
}
func (this *Grammar) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Grammar)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.TopPattern.Equal(that1.TopPattern) {
		return false
	}
	if len(this.PatternDecls) != len(that1.PatternDecls) {
		return false
	}
	for i := range this.PatternDecls {
		if !this.PatternDecls[i].Equal(that1.PatternDecls[i]) {
			return false
		}
	}
	if !this.After.Equal(that1.After) {
		return false
	}
	return true
}
func (this *PatternDecl) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*PatternDecl)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.At.Equal(that1.At) {
		return false
	}
	if !this.Before.Equal(that1.Before) {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if !this.Eq.Equal(that1.Eq) {
		return false
	}
	if !this.Pattern.Equal(that1.Pattern) {
		return false
	}
	return true
}
func (this *NameExpr) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*NameExpr)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Name.Equal(that1.Name) {
		return false
	}
	if !this.AnyName.Equal(that1.AnyName) {
		return false
	}
	if !this.AnyNameExcept.Equal(that1.AnyNameExcept) {
		return false
	}
	if !this.NameChoice.Equal(that1.NameChoice) {
		return false
	}
	return true
}
func (this *Name) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Name)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Before.Equal(that1.Before) {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *AnyName) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*AnyName)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Underscore.Equal(that1.Underscore) {
		return false
	}
	return true
}
func (this *AnyNameExcept) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*AnyNameExcept)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Exclamation.Equal(that1.Exclamation) {
		return false
	}
	if !this.OpenParen.Equal(that1.OpenParen) {
		return false
	}
	if !this.Except.Equal(that1.Except) {
		return false
	}
	if !this.CloseParen.Equal(that1.CloseParen) {
		return false
	}
	return true
}
func (this *NameChoice) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*NameChoice)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.OpenParen.Equal(that1.OpenParen) {
		return false
	}
	if !this.Left.Equal(that1.Left) {
		return false
	}
	if !this.Pipe.Equal(that1.Pipe) {
		return false
	}
	if !this.Right.Equal(that1.Right) {
		return false
	}
	if !this.CloseParen.Equal(that1.CloseParen) {
		return false
	}
	return true
}
func (this *Pattern) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Pattern)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Empty.Equal(that1.Empty) {
		return false
	}
	if !this.EmptySet.Equal(that1.EmptySet) {
		return false
	}
	if !this.TreeNode.Equal(that1.TreeNode) {
		return false
	}
	if !this.LeafNode.Equal(that1.LeafNode) {
		return false
	}
	if !this.Concat.Equal(that1.Concat) {
		return false
	}
	if !this.Or.Equal(that1.Or) {
		return false
	}
	if !this.And.Equal(that1.And) {
		return false
	}
	if !this.ZeroOrMore.Equal(that1.ZeroOrMore) {
		return false
	}
	if !this.Reference.Equal(that1.Reference) {
		return false
	}
	if !this.Not.Equal(that1.Not) {
		return false
	}
	if !this.ZAny.Equal(that1.ZAny) {
		return false
	}
	if !this.WithSomeTreeNode.Equal(that1.WithSomeTreeNode) {
		return false
	}
	return true
}
func (this *Empty) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Empty)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Empty.Equal(that1.Empty) {
		return false
	}
	return true
}
func (this *EmptySet) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*EmptySet)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.EmptySet.Equal(that1.EmptySet) {
		return false
	}
	return true
}
func (this *TreeNode) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*TreeNode)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Name.Equal(that1.Name) {
		return false
	}
	if !this.Colon.Equal(that1.Colon) {
		return false
	}
	if !this.Pattern.Equal(that1.Pattern) {
		return false
	}
	return true
}
func (this *WithSomeTreeNode) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*WithSomeTreeNode)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Dot.Equal(that1.Dot) {
		return false
	}
	if !this.Pattern.Equal(that1.Pattern) {
		return false
	}
	return true
}
func (this *LeafNode) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*LeafNode)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.RightArrow.Equal(that1.RightArrow) {
		return false
	}
	if !this.Expr.Equal(that1.Expr) {
		return false
	}
	return true
}
func (this *Concat) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Concat)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.OpenBracket.Equal(that1.OpenBracket) {
		return false
	}
	if !this.LeftPattern.Equal(that1.LeftPattern) {
		return false
	}
	if !this.Comma.Equal(that1.Comma) {
		return false
	}
	if !this.RightPattern.Equal(that1.RightPattern) {
		return false
	}
	if !this.CloseBracket.Equal(that1.CloseBracket) {
		return false
	}
	return true
}
func (this *Or) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Or)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.OpenParen.Equal(that1.OpenParen) {
		return false
	}
	if !this.LeftPattern.Equal(that1.LeftPattern) {
		return false
	}
	if !this.Pipe.Equal(that1.Pipe) {
		return false
	}
	if !this.RightPattern.Equal(that1.RightPattern) {
		return false
	}
	if !this.CloseParen.Equal(that1.CloseParen) {
		return false
	}
	return true
}
func (this *WithSomeOr) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*WithSomeOr)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.OpenCurly.Equal(that1.OpenCurly) {
		return false
	}
	if !this.LeftPattern.Equal(that1.LeftPattern) {
		return false
	}
	if !this.Pipe.Equal(that1.Pipe) {
		return false
	}
	if !this.RightPattern.Equal(that1.RightPattern) {
		return false
	}
	if !this.CloseCurly.Equal(that1.CloseCurly) {
		return false
	}
	return true
}
func (this *And) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*And)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.OpenParen.Equal(that1.OpenParen) {
		return false
	}
	if !this.LeftPattern.Equal(that1.LeftPattern) {
		return false
	}
	if !this.Ampersand.Equal(that1.Ampersand) {
		return false
	}
	if !this.RightPattern.Equal(that1.RightPattern) {
		return false
	}
	if !this.CloseParen.Equal(that1.CloseParen) {
		return false
	}
	return true
}
func (this *WithSomeAnd) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*WithSomeAnd)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.OpenCurly.Equal(that1.OpenCurly) {
		return false
	}
	if !this.LeftPattern.Equal(that1.LeftPattern) {
		return false
	}
	if !this.Ampersand.Equal(that1.Ampersand) {
		return false
	}
	if !this.RightPattern.Equal(that1.RightPattern) {
		return false
	}
	if !this.CloseCurly.Equal(that1.CloseCurly) {
		return false
	}
	return true
}
func (this *ZeroOrMore) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ZeroOrMore)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.OpenParen.Equal(that1.OpenParen) {
		return false
	}
	if !this.Pattern.Equal(that1.Pattern) {
		return false
	}
	if !this.CloseParen.Equal(that1.CloseParen) {
		return false
	}
	if !this.Star.Equal(that1.Star) {
		return false
	}
	return true
}
func (this *Reference) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Reference)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.HashTag.Equal(that1.HashTag) {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *Not) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Not)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Exclamation.Equal(that1.Exclamation) {
		return false
	}
	if !this.OpenParen.Equal(that1.OpenParen) {
		return false
	}
	if !this.Pattern.Equal(that1.Pattern) {
		return false
	}
	if !this.CloseParen.Equal(that1.CloseParen) {
		return false
	}
	return true
}
func (this *ZAny) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ZAny)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Star.Equal(that1.Star) {
		return false
	}
	return true
}
func (this *Grammar) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.Grammar{` +
		`TopPattern:` + fmt.Sprintf("%#v", this.TopPattern),
		`PatternDecls:` + fmt.Sprintf("%#v", this.PatternDecls),
		`After:` + fmt.Sprintf("%#v", this.After) + `}`}, ", ")
	return s
}
func (this *PatternDecl) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.PatternDecl{` +
		`At:` + fmt.Sprintf("%#v", this.At),
		`Before:` + fmt.Sprintf("%#v", this.Before),
		`Name:` + fmt.Sprintf("%#v", this.Name),
		`Eq:` + fmt.Sprintf("%#v", this.Eq),
		`Pattern:` + fmt.Sprintf("%#v", this.Pattern) + `}`}, ", ")
	return s
}
func (this *NameExpr) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.NameExpr{` +
		`Name:` + fmt.Sprintf("%#v", this.Name),
		`AnyName:` + fmt.Sprintf("%#v", this.AnyName),
		`AnyNameExcept:` + fmt.Sprintf("%#v", this.AnyNameExcept),
		`NameChoice:` + fmt.Sprintf("%#v", this.NameChoice) + `}`}, ", ")
	return s
}
func (this *Name) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.Name{` +
		`Before:` + fmt.Sprintf("%#v", this.Before),
		`Name:` + fmt.Sprintf("%#v", this.Name) + `}`}, ", ")
	return s
}
func (this *AnyName) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.AnyName{` +
		`Underscore:` + fmt.Sprintf("%#v", this.Underscore) + `}`}, ", ")
	return s
}
func (this *AnyNameExcept) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.AnyNameExcept{` +
		`Exclamation:` + fmt.Sprintf("%#v", this.Exclamation),
		`OpenParen:` + fmt.Sprintf("%#v", this.OpenParen),
		`Except:` + fmt.Sprintf("%#v", this.Except),
		`CloseParen:` + fmt.Sprintf("%#v", this.CloseParen) + `}`}, ", ")
	return s
}
func (this *NameChoice) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.NameChoice{` +
		`OpenParen:` + fmt.Sprintf("%#v", this.OpenParen),
		`Left:` + fmt.Sprintf("%#v", this.Left),
		`Pipe:` + fmt.Sprintf("%#v", this.Pipe),
		`Right:` + fmt.Sprintf("%#v", this.Right),
		`CloseParen:` + fmt.Sprintf("%#v", this.CloseParen) + `}`}, ", ")
	return s
}
func (this *Pattern) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.Pattern{` +
		`Empty:` + fmt.Sprintf("%#v", this.Empty),
		`EmptySet:` + fmt.Sprintf("%#v", this.EmptySet),
		`TreeNode:` + fmt.Sprintf("%#v", this.TreeNode),
		`LeafNode:` + fmt.Sprintf("%#v", this.LeafNode),
		`Concat:` + fmt.Sprintf("%#v", this.Concat),
		`Or:` + fmt.Sprintf("%#v", this.Or),
		`And:` + fmt.Sprintf("%#v", this.And),
		`ZeroOrMore:` + fmt.Sprintf("%#v", this.ZeroOrMore),
		`Reference:` + fmt.Sprintf("%#v", this.Reference),
		`Not:` + fmt.Sprintf("%#v", this.Not),
		`ZAny:` + fmt.Sprintf("%#v", this.ZAny),
		`WithSomeTreeNode:` + fmt.Sprintf("%#v", this.WithSomeTreeNode) + `}`}, ", ")
	return s
}
func (this *Empty) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.Empty{` +
		`Empty:` + fmt.Sprintf("%#v", this.Empty) + `}`}, ", ")
	return s
}
func (this *EmptySet) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.EmptySet{` +
		`EmptySet:` + fmt.Sprintf("%#v", this.EmptySet) + `}`}, ", ")
	return s
}
func (this *TreeNode) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.TreeNode{` +
		`Name:` + fmt.Sprintf("%#v", this.Name),
		`Colon:` + fmt.Sprintf("%#v", this.Colon),
		`Pattern:` + fmt.Sprintf("%#v", this.Pattern) + `}`}, ", ")
	return s
}
func (this *WithSomeTreeNode) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.WithSomeTreeNode{` +
		`Dot:` + fmt.Sprintf("%#v", this.Dot),
		`Pattern:` + fmt.Sprintf("%#v", this.Pattern) + `}`}, ", ")
	return s
}
func (this *LeafNode) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.LeafNode{` +
		`RightArrow:` + fmt.Sprintf("%#v", this.RightArrow),
		`Expr:` + fmt.Sprintf("%#v", this.Expr) + `}`}, ", ")
	return s
}
func (this *Concat) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.Concat{` +
		`OpenBracket:` + fmt.Sprintf("%#v", this.OpenBracket),
		`LeftPattern:` + fmt.Sprintf("%#v", this.LeftPattern),
		`Comma:` + fmt.Sprintf("%#v", this.Comma),
		`RightPattern:` + fmt.Sprintf("%#v", this.RightPattern),
		`CloseBracket:` + fmt.Sprintf("%#v", this.CloseBracket) + `}`}, ", ")
	return s
}
func (this *Or) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.Or{` +
		`OpenParen:` + fmt.Sprintf("%#v", this.OpenParen),
		`LeftPattern:` + fmt.Sprintf("%#v", this.LeftPattern),
		`Pipe:` + fmt.Sprintf("%#v", this.Pipe),
		`RightPattern:` + fmt.Sprintf("%#v", this.RightPattern),
		`CloseParen:` + fmt.Sprintf("%#v", this.CloseParen) + `}`}, ", ")
	return s
}
func (this *WithSomeOr) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.WithSomeOr{` +
		`OpenCurly:` + fmt.Sprintf("%#v", this.OpenCurly),
		`LeftPattern:` + fmt.Sprintf("%#v", this.LeftPattern),
		`Pipe:` + fmt.Sprintf("%#v", this.Pipe),
		`RightPattern:` + fmt.Sprintf("%#v", this.RightPattern),
		`CloseCurly:` + fmt.Sprintf("%#v", this.CloseCurly) + `}`}, ", ")
	return s
}
func (this *And) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.And{` +
		`OpenParen:` + fmt.Sprintf("%#v", this.OpenParen),
		`LeftPattern:` + fmt.Sprintf("%#v", this.LeftPattern),
		`Ampersand:` + fmt.Sprintf("%#v", this.Ampersand),
		`RightPattern:` + fmt.Sprintf("%#v", this.RightPattern),
		`CloseParen:` + fmt.Sprintf("%#v", this.CloseParen) + `}`}, ", ")
	return s
}
func (this *WithSomeAnd) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.WithSomeAnd{` +
		`OpenCurly:` + fmt.Sprintf("%#v", this.OpenCurly),
		`LeftPattern:` + fmt.Sprintf("%#v", this.LeftPattern),
		`Ampersand:` + fmt.Sprintf("%#v", this.Ampersand),
		`RightPattern:` + fmt.Sprintf("%#v", this.RightPattern),
		`CloseCurly:` + fmt.Sprintf("%#v", this.CloseCurly) + `}`}, ", ")
	return s
}
func (this *ZeroOrMore) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.ZeroOrMore{` +
		`OpenParen:` + fmt.Sprintf("%#v", this.OpenParen),
		`Pattern:` + fmt.Sprintf("%#v", this.Pattern),
		`CloseParen:` + fmt.Sprintf("%#v", this.CloseParen),
		`Star:` + fmt.Sprintf("%#v", this.Star) + `}`}, ", ")
	return s
}
func (this *Reference) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.Reference{` +
		`HashTag:` + fmt.Sprintf("%#v", this.HashTag),
		`Name:` + fmt.Sprintf("%#v", this.Name) + `}`}, ", ")
	return s
}
func (this *Not) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.Not{` +
		`Exclamation:` + fmt.Sprintf("%#v", this.Exclamation),
		`OpenParen:` + fmt.Sprintf("%#v", this.OpenParen),
		`Pattern:` + fmt.Sprintf("%#v", this.Pattern),
		`CloseParen:` + fmt.Sprintf("%#v", this.CloseParen) + `}`}, ", ")
	return s
}
func (this *ZAny) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&relapse.ZAny{` +
		`Star:` + fmt.Sprintf("%#v", this.Star) + `}`}, ", ")
	return s
}
func valueToGoStringRelapse(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringRelapse(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (this *NameExpr) GetValue() interface{} {
	if this.Name != nil {
		return this.Name
	}
	if this.AnyName != nil {
		return this.AnyName
	}
	if this.AnyNameExcept != nil {
		return this.AnyNameExcept
	}
	if this.NameChoice != nil {
		return this.NameChoice
	}
	return nil
}

func (this *NameExpr) SetValue(value interface{}) bool {
	switch vt := value.(type) {
	case *Name:
		this.Name = vt
	case *AnyName:
		this.AnyName = vt
	case *AnyNameExcept:
		this.AnyNameExcept = vt
	case *NameChoice:
		this.NameChoice = vt
	default:
		return false
	}
	return true
}
func (this *Pattern) GetValue() interface{} {
	if this.Empty != nil {
		return this.Empty
	}
	if this.EmptySet != nil {
		return this.EmptySet
	}
	if this.TreeNode != nil {
		return this.TreeNode
	}
	if this.LeafNode != nil {
		return this.LeafNode
	}
	if this.Concat != nil {
		return this.Concat
	}
	if this.Or != nil {
		return this.Or
	}
	if this.And != nil {
		return this.And
	}
	if this.ZeroOrMore != nil {
		return this.ZeroOrMore
	}
	if this.Reference != nil {
		return this.Reference
	}
	if this.Not != nil {
		return this.Not
	}
	if this.ZAny != nil {
		return this.ZAny
	}
	if this.WithSomeTreeNode != nil {
		return this.WithSomeTreeNode
	}
	return nil
}

func (this *Pattern) SetValue(value interface{}) bool {
	switch vt := value.(type) {
	case *Empty:
		this.Empty = vt
	case *EmptySet:
		this.EmptySet = vt
	case *TreeNode:
		this.TreeNode = vt
	case *LeafNode:
		this.LeafNode = vt
	case *Concat:
		this.Concat = vt
	case *Or:
		this.Or = vt
	case *And:
		this.And = vt
	case *ZeroOrMore:
		this.ZeroOrMore = vt
	case *Reference:
		this.Reference = vt
	case *Not:
		this.Not = vt
	case *ZAny:
		this.ZAny = vt
	case *WithSomeTreeNode:
		this.WithSomeTreeNode = vt
	default:
		return false
	}
	return true
}
