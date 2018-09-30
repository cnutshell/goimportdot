package core

import (
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
)

func ParseGoImport(gofile string) (ss StrSet, err error) {
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, gofile, nil, parser.ImportsOnly)
	if err != nil {
		return
	}
	ss = NewStrSet()
	for _, s := range f.Imports {
		ss.Put(strings.Trim(s.Path.Value, `"`))
	}
	return
}
func PkgOfFile(gofile string) (pkg string) {
	unixPath := filepath.ToSlash(filepath.Dir(gofile))
	return strings.SplitN(unixPath, "/src/", 2)[1]
}

var placeholder = struct{}{}

type StrSet map[string]struct{}

func NewStrSet(strs ...string) StrSet {
	ss := StrSet(make(map[string]struct{}))
	for _, str := range strs {
		ss.Put(str)
	}
	return ss
}
func (this StrSet) Put(str string)                { this[str] = placeholder }
func (this StrSet) Del(str string)                { delete(this, str) }
func (this StrSet) Contains(str string) (ok bool) { _, ok = this[str]; return ok }
func (this StrSet) Merge(that StrSet) {
	for str := range that {
		this[str] = placeholder
	}
}
func (this StrSet) Array() []string {
	ret := make([]string, 0, len(this))
	for str := range this {
		ret = append(ret, str)
	}
	return ret
}
