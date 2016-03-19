package config

import (
	"reflect"
	"testing"
)

func assert(t *testing.T, name string, a, b interface{}) {
	if a != b {
		t.Errorf("%s should be %v, but %v", name, a, b)
	}
}

func assertP(t *testing.T, name string, a, b interface{}) {
	v := reflect.ValueOf(b)
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			t.Errorf("%s should be %v, but %v", name, a, b)
			return
		}
		v = v.Elem()
	}
	if a != v.Interface() {
		t.Errorf("%s should be %v, but %v", name, a, b)
	}
}

type PrimaryT struct {
	B   bool    `default:"true"`
	N   int     `default:"12345"`
	F   float32 `default:"12.234"`
	S   string  `default:"hello"`
	Bp  *bool   `default:"false"`
	Bpn *bool
}

func (p *PrimaryT) checkFields(t *testing.T) {
	assert(t, "p.B", true, p.B)
	assert(t, "p.N", 12345, p.N)
	assert(t, "p.F", float32(12.234), p.F)
	assert(t, "p.S", "hello", p.S)
	assertP(t, "p.Bp", false, p.Bp)
	assert(t, "p.Bpn", (*bool)(nil), p.Bpn)
}

func TestPrimaryType(t *testing.T) {
	var p PrimaryT
	if err := DefaultConfig(&p); err != nil {
		t.Errorf("set defaults fail: %v", err)
	} else {
		p.checkFields(t)
	}
}

type EmbedT struct {
	PrimaryT
}

func TestEmbedType(t *testing.T) {
	var e EmbedT
	if err := DefaultConfig(&e); err != nil {
		t.Errorf("set defaults fail: %v", err)
	} else {
		e.checkFields(t)
	}
}

type EmbedPT struct {
	*PrimaryT
}

func TestEmbedPType(t *testing.T) {
	var e EmbedPT
	if err := DefaultConfig(&e); err != nil {
		t.Errorf("set defaults fail: %v", err)
	} else {
		e.checkFields(t)
	}
}

type EmbedPNoneT struct {
	*PrimaryT `default:"-"`
}

func TestEmbedPNoneType(t *testing.T) {
	var e EmbedPNoneT
	if err := DefaultConfig(&e); err != nil {
		t.Errorf("set defaults fail: %v", err)
	} else if e.PrimaryT != nil {
		t.Errorf("none default fail: %v", e.PrimaryT)
	}
}

type StructFieldT struct {
	P PrimaryT
}

func TestStructFieldType(t *testing.T) {
	var s StructFieldT
	if err := DefaultConfig(&s); err != nil {
		t.Errorf("set defaults fail: %v", err)
	} else {
		s.P.checkFields(t)
	}
}

type StructFieldPT struct {
	P *PrimaryT
}

func TestStructFieldPType(t *testing.T) {
	var s StructFieldPT
	if err := DefaultConfig(&s); err != nil {
		t.Errorf("set defaults fail: %v", err)
	} else {
		s.P.checkFields(t)
	}
}