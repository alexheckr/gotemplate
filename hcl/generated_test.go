// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package hcl

import (
	"reflect"
	"strings"
	"testing"
)

var strFixture = hclList(hclListHelper.NewStringList(strings.Split("Hello World, I'm Foo Bar!", " ")...).AsArray())

func Test_list_Append(t *testing.T) {
	tests := []struct {
		name   string
		l      hclIList
		values []interface{}
		want   hclIList
	}{
		{"Empty", hclList{}, []interface{}{1, 2, 3}, hclList{1, 2, 3}},
		{"List of int", hclList{1, 2, 3}, []interface{}{4}, hclList{1, 2, 3, 4}},
		{"List of string", strFixture, []interface{}{"That's all folks!"}, hclList{"Hello", "World,", "I'm", "Foo", "Bar!", "That's all folks!"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Append(tt.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclList.Append():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_AsArray(t *testing.T) {
	tests := []struct {
		name string
		l    hclList
		want []interface{}
	}{
		{"Empty List", hclList{}, []interface{}{}},
		{"List of int", hclList{1, 2, 3}, []interface{}{1, 2, 3}},
		{"List of string", strFixture, []interface{}{"Hello", "World,", "I'm", "Foo", "Bar!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.AsArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclList.AsList():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_HclList_Strings(t *testing.T) {
	tests := []struct {
		name string
		l    hclList
		want []string
	}{
		{"Empty List", hclList{}, []string{}},
		{"List of int", hclList{1, 2, 3}, []string{"1", "2", "3"}},
		{"List of string", strFixture, []string{"Hello", "World,", "I'm", "Foo", "Bar!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Strings(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclList.Strings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_Capacity(t *testing.T) {
	tests := []struct {
		name string
		l    hclIList
		want int
	}{
		{"Empty List with 100 spaces", hclListHelper.CreateList(0, 100), 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Capacity(); got != tt.want {
				t.Errorf("hclList.Capacity() = %v, want %v", got, tt.want)
			}
			if tt.l.Capacity() != tt.l.Cap() {
				t.Errorf("Cap and Capacity return different values")
			}
		})
	}
}

func Test_list_Clone(t *testing.T) {
	tests := []struct {
		name string
		l    hclList
		want hclIList
	}{
		{"Empty List", hclList{}, hclList{}},
		{"List of int", hclList{1, 2, 3}, hclList{1, 2, 3}},
		{"List of string", strFixture, hclList{"Hello", "World,", "I'm", "Foo", "Bar!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclList.Clone():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_Get(t *testing.T) {
	tests := []struct {
		name  string
		l     hclList
		index int
		want  interface{}
	}{
		{"Empty List", hclList{}, 0, nil},
		{"Negative index", hclList{}, -1, nil},
		{"List of int", hclList{1, 2, 3}, 0, 1},
		{"List of string", strFixture, 1, "World,"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Get(tt.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclList.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_Len(t *testing.T) {
	tests := []struct {
		name string
		l    hclList
		want int
	}{
		{"Empty List", hclList{}, 0},
		{"List of int", hclList{1, 2, 3}, 3},
		{"List of string", strFixture, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Len(); got != tt.want {
				t.Errorf("hclList.Len() = %v, want %v", got, tt.want)
			}
			if tt.l.Len() != tt.l.Count() {
				t.Errorf("Len and Count return different values")
			}
		})
	}
}

func Test_NewList(t *testing.T) {
	type args struct {
		size     int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want hclIList
	}{
		{"Empty", args{0, 0}, hclList{}},
		{"With nil elements", args{10, 0}, make(hclList, 10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hclListHelper.CreateList(tt.args.size, tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclList.CreateList():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_Reverse(t *testing.T) {
	tests := []struct {
		name string
		l    hclList
		want hclIList
	}{
		{"Empty List", hclList{}, hclList{}},
		{"List of int", hclList{1, 2, 3}, hclList{3, 2, 1}},
		{"List of string", strFixture, hclList{"Bar!", "Foo", "I'm", "World,", "Hello"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.l.Clone()
			if got := l.Reverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclList.Reverse():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_list_Set(t *testing.T) {
	type args struct {
		i int
		v interface{}
	}
	tests := []struct {
		name    string
		l       hclIList
		args    args
		want    hclIList
		wantErr bool
	}{
		{"Empty", hclList{}, args{2, 1}, hclList{nil, nil, 1}, false},
		{"List of int", hclList{1, 2, 3}, args{0, 10}, hclList{10, 2, 3}, false},
		{"List of string", strFixture, args{2, "You're"}, hclList{"Hello", "World,", "You're", "Foo", "Bar!"}, false},
		{"Negative", hclList{}, args{-1, "negative value"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.Clone().Set(tt.args.i, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("hclList.Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclList.Set():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

var mapFixture = map[string]interface{}{
	"int":     123,
	"float":   1.23,
	"string":  "Foo bar",
	"list":    []interface{}{1, "two"},
	"listInt": []int{1, 2, 3},
	"map": map[string]interface{}{
		"sub1": 1,
		"sub2": "two",
	},
	"mapInt": map[int]interface{}{
		1: 1,
		2: "two",
	},
}

var dictFixture = hclDict(hclDictHelper.AsDictionary(mapFixture).AsMap())

func dumpKeys(t *testing.T, d1, d2 hclIDict) {
	for key := range d1.AsMap() {
		v1, v2 := d1.Get(key), d2.Get(key)
		if reflect.DeepEqual(v1, v2) {
			continue
		}
		t.Logf("'%[1]v' = %[2]v (%[2]T) vs %[3]v (%[3]T)", key, v1, v2)
	}
}

func Test_dict_AsMap(t *testing.T) {
	tests := []struct {
		name string
		d    hclDict
		want map[string]interface{}
	}{
		{"Nil", nil, nil},
		{"Empty", hclDict{}, map[string]interface{}{}},
		{"Map", dictFixture, map[string]interface{}(dictFixture)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.AsMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclDict.AsMap():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_dict_Clone(t *testing.T) {
	tests := []struct {
		name string
		d    hclDict
		keys []interface{}
		want hclIDict
	}{
		{"Nil", nil, nil, hclDict{}},
		{"Empty", hclDict{}, nil, hclDict{}},
		{"Map", dictFixture, nil, dictFixture},
		{"Map with Fields", dictFixture, []interface{}{"int", "list"}, hclDict(dictFixture).Omit("float", "string", "listInt", "map", "mapInt")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d.Clone(tt.keys...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclDict.Clone():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
				dumpKeys(t, got, tt.want)
			}

			// Ensure that the copy is distinct from the original
			got.Set("NewFields", "Test")
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Should be different:\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
			if !got.Has("NewFields") || !reflect.DeepEqual(got.Get("NewFields"), "Test") {
				t.Errorf("Element has not been added")
			}
			if got.Len() != tt.want.Count()+1 {
				t.Errorf("Len and Count don't return the same value")
			}
		})
	}
}

func Test_HclDict_CreateList(t *testing.T) {
	tests := []struct {
		name         string
		d            hclDict
		args         []int
		want         hclIList
		wantLen      int
		wantCapacity int
	}{
		{"Nil", nil, nil, hclList{}, 0, 0},
		{"Empty", hclDict{}, nil, hclList{}, 0, 0},
		{"Map", dictFixture, nil, hclList{}, 0, 0},
		{"Map with size", dictFixture, []int{3}, hclList{nil, nil, nil}, 3, 3},
		{"Map with capacity", dictFixture, []int{0, 10}, hclList{}, 0, 10},
		{"Map with size&capacity", dictFixture, []int{3, 10}, hclList{nil, nil, nil}, 3, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d.CreateList(tt.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclDict.CreateList() = %v, want %v", got, tt.want)
			}
			if got.Len() != tt.wantLen || got.Cap() != tt.wantCapacity {
				t.Errorf("hclDict.CreateList() size: %d, %d vs %d, %d", got.Len(), got.Cap(), tt.wantLen, tt.wantCapacity)
			}
		})
	}
}

func Test_dict_Default(t *testing.T) {
	type args struct {
		key    interface{}
		defVal interface{}
	}
	tests := []struct {
		name string
		d    hclDict
		args args
		want interface{}
	}{
		{"Empty", nil, args{"Foo", "Bar"}, "Bar"},
		{"Map int", dictFixture, args{"int", 1}, 123},
		{"Map float", dictFixture, args{"float", 1}, 1.23},
		{"Map Non existant", dictFixture, args{"Foo", "Bar"}, "Bar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Default(tt.args.key, tt.args.defVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclDict.Default() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dict_Delete(t *testing.T) {
	type args struct {
		key  interface{}
		keys []interface{}
	}
	tests := []struct {
		name    string
		d       hclDict
		args    args
		want    hclIDict
		wantErr bool
	}{
		{"Empty", nil, args{}, hclDict{}, true},
		{"Map", dictFixture, args{}, dictFixture, true},
		{"Non existant key", dictFixture, args{"Test", nil}, dictFixture, true},
		{"Map with keys", dictFixture, args{"int", []interface{}{"list"}}, dictFixture.Clone("float", "string", "listInt", "map", "mapInt"), false},
		{"Map with keys + non existant", dictFixture, args{"int", []interface{}{"list", "Test"}}, dictFixture.Clone("float", "string", "listInt", "map", "mapInt"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d.Clone()
			got, err := d.Delete(tt.args.key, tt.args.keys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("hclDict.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclDict.Delete():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
				dumpKeys(t, got, tt.want)
			}
		})
	}
}

func Test_dict_Flush(t *testing.T) {
	tests := []struct {
		name string
		d    hclDict
		keys []interface{}
		want hclIDict
	}{
		{"Empty", nil, nil, hclDict{}},
		{"Map", dictFixture, nil, hclDict{}},
		{"Non existant key", dictFixture, []interface{}{"Test"}, dictFixture},
		{"Map with keys", dictFixture, []interface{}{"int", "list"}, dictFixture.Clone("float", "string", "listInt", "map", "mapInt")},
		{"Map with keys + non existant", dictFixture, []interface{}{"int", "list", "Test"}, dictFixture.Clone("float", "string", "listInt", "map", "mapInt")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d.Clone()
			got := d.Flush(tt.keys...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclDict.Flush():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
				dumpKeys(t, got, tt.want)
			}
			if !reflect.DeepEqual(d, got) {
				t.Errorf("Should be equal after: %v, want %v", d, got)
				dumpKeys(t, d, got)
			}
		})
	}
}

func Test_dict_Keys(t *testing.T) {
	tests := []struct {
		name string
		d    hclDict
		want hclIList
	}{
		{"Empty", nil, hclList{}},
		{"Map", dictFixture, hclList{"float", "int", "list", "listInt", "map", "mapInt", "string"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.GetKeys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclDict.GetKeys():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_dict_KeysAsString(t *testing.T) {
	tests := []struct {
		name string
		d    hclDict
		want []string
	}{
		{"Empty", nil, []string{}},
		{"Map", dictFixture, []string{"float", "int", "list", "listInt", "map", "mapInt", "string"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.KeysAsString(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclDict.KeysAsString():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_dict_Merge(t *testing.T) {
	adding1 := hclDict{
		"int":        1000,
		"Add1Int":    1,
		"Add1String": "string",
	}
	adding2 := hclDict{
		"Add2Int":    1,
		"Add2String": "string",
		"map": map[string]interface{}{
			"sub1":   2,
			"newVal": "NewValue",
		},
	}
	type args struct {
		hclDict hclIDict
		dicts   []hclIDict
	}
	tests := []struct {
		name string
		d    hclDict
		args args
		want hclIDict
	}{
		{"Empty", nil, args{nil, []hclIDict{}}, hclDict{}},
		{"Add map to empty", nil, args{dictFixture, []hclIDict{}}, dictFixture},
		{"Add map to same map", dictFixture, args{dictFixture, []hclIDict{}}, dictFixture},
		{"Add empty to map", dictFixture, args{nil, []hclIDict{}}, dictFixture},
		{"Add new1 to map", dictFixture, args{adding1, []hclIDict{}}, dictFixture.Clone().Merge(adding1)},
		{"Add new2 to map", dictFixture, args{adding2, []hclIDict{}}, dictFixture.Clone().Merge(adding2)},
		{"Add new1 & new2 to map", dictFixture, args{adding1, []hclIDict{adding2}}, dictFixture.Clone().Merge(adding1, adding2)},
		{"Add new1 & new2 to map", dictFixture, args{adding1, []hclIDict{adding2}}, dictFixture.Clone().Merge(adding1).Merge(adding2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d.Clone()
			got := d.Merge(tt.args.hclDict, tt.args.dicts...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclDict.Merge():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
				dumpKeys(t, got, tt.want)
			}
		})
	}
}

func Test_dict_Values(t *testing.T) {
	tests := []struct {
		name string
		d    hclDict
		want hclIList
	}{
		{"Empty", nil, hclList{}},
		{"Map", dictFixture, hclList{1.23, 123, hclList{1, "two"}, hclList{1, 2, 3}, hclDict{"sub1": 1, "sub2": "two"}, hclDict{"1": 1, "2": "two"}, "Foo bar"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.GetValues(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hclDict.GetValues():\n got %[1]v (%[1]T)\nwant %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}
