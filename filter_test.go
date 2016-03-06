package filter

import (
	"reflect"
	"testing"
)

func even(i int) bool {
	return i%2 == 0
}

func always(i int) bool {
	return true
}

var splitTests = []struct {
	name    string
	list    []int
	p       Predicate
	wantin  []int
	wantout []int
}{
	{
		name:    "trivial",
		list:    []int{},
		p:       always,
		wantin:  []int{},
		wantout: []int{},
	},
	{
		name:    "even",
		list:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		p:       even,
		wantin:  []int{2, 4, 6, 8, 10},
		wantout: []int{1, 3, 5, 7, 9},
	},
}

func TestSplit(t *testing.T) {
	for _, tt := range splitTests {
		gotin, gotout := Split(tt.list, tt.p)
		if !reflect.DeepEqual(gotin, tt.wantin) {
			t.Errorf("%v: gotin %v, wantin %v", tt.name, gotin, tt.wantin)
		}
		if !orderedEqual(gotout, tt.wantout) {
			t.Errorf("%v: gotout %v, wantout %v", tt.name, gotout, tt.wantout)
		}
	}
}

func orderedEqual(a, b []int) bool {
	am := make(map[int]bool)
	for _, aa := range a {
		am[aa] = true
	}
	bm := make(map[int]bool)
	for _, bb := range b {
		bm[bb] = true
	}
	return reflect.DeepEqual(am, bm)
}

var filterTests = []struct {
	name string
	list []int
	p    Predicate
	want []int
}{
	{
		name: "trivial",
		list: []int{},
		p:    always,
		want: []int{},
	},
	{
		name: "even",
		list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		p:    even,
		want: []int{2, 4, 6, 8, 10},
	},
}

func TestFilter(t *testing.T) {
	for _, tt := range filterTests {
		got := Filter(tt.list, tt.p)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v: got %v, want %v", tt.name, got, tt.want)
		}
	}
}
