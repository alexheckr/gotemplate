// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package json

import "github.com/coveo/gotemplate/collections"

// List implementation of IGenericList for jsonList
type List = jsonList
type jsonIList = collections.IGenericList
type jsonList []interface{}

func (l jsonList) AsArray() []interface{}              { return []interface{}(l) }
func (l jsonList) Cap() int                            { return cap(l) }
func (l jsonList) Capacity() int                       { return cap(l) }
func (l jsonList) Clone() jsonIList                    { return jsonListHelper.Clone(l) }
func (l jsonList) Contains(values ...interface{}) bool { return jsonListHelper.Contains(l, values...) }
func (l jsonList) Count() int                          { return len(l) }
func (l jsonList) Create(args ...int) jsonIList        { return jsonListHelper.CreateList(args...) }
func (l jsonList) Get(index int) interface{}           { return jsonListHelper.GetIndex(l, index) }
func (l jsonList) Len() int                            { return len(l) }
func (l jsonList) New(args ...interface{}) jsonIList   { return jsonListHelper.NewList(args...) }
func (l jsonList) Reverse() jsonIList                  { return jsonListHelper.Reverse(l) }
func (l jsonList) Strings() []string                   { return jsonListHelper.GetStrings(l) }
func (l jsonList) Unique() jsonIList                   { return jsonListHelper.Unique(l) }

func (l jsonList) Append(values ...interface{}) jsonIList {
	return jsonListHelper.Add(l, false, values...)
}

func (l jsonList) Prepend(values ...interface{}) jsonIList {
	return jsonListHelper.Add(l, true, values...)
}

func (l jsonList) Set(i int, v interface{}) (jsonIList, error) {
	return jsonListHelper.SetIndex(l, i, v)
}

func (l jsonList) Without(values ...interface{}) jsonIList {
	return jsonListHelper.Without(l, values...)
}

// Dictionary implementation of IDictionary for jsonDict
type Dictionary = jsonDict
type jsonIDict = collections.IDictionary
type jsonDict map[string]interface{}

func (d jsonDict) AsMap() map[string]interface{}       { return (map[string]interface{})(d) }
func (d jsonDict) Native() interface{}                 { return collections.ToNativeRepresentation(d) }
func (d jsonDict) Count() int                          { return len(d) }
func (d jsonDict) Len() int                            { return len(d) }
func (d jsonDict) Clone(keys ...interface{}) jsonIDict { return jsonDictHelper.Clone(d, keys) }
func (d jsonDict) CreateList(args ...int) jsonIList    { return jsonHelper.CreateList(args...) }
func (d jsonDict) Flush(keys ...interface{}) jsonIDict { return jsonDictHelper.Flush(d, keys) }
func (d jsonDict) Get(key interface{}) interface{}     { return jsonDictHelper.Get(d, key) }
func (d jsonDict) Has(key interface{}) bool            { return jsonDictHelper.Has(d, key) }
func (d jsonDict) GetKeys() jsonIList                  { return jsonDictHelper.GetKeys(d) }
func (d jsonDict) KeysAsString() []string              { return jsonDictHelper.KeysAsString(d) }
func (d jsonDict) GetValues() jsonIList                { return jsonDictHelper.GetValues(d) }

func (d jsonDict) Default(key, defVal interface{}) interface{} {
	return jsonDictHelper.Default(d, key, defVal)
}

func (d jsonDict) Delete(key interface{}, otherKeys ...interface{}) (jsonIDict, error) {
	return jsonDictHelper.Delete(d, append([]interface{}{key}, otherKeys...))
}

func (d jsonDict) Merge(dict jsonIDict, otherDicts ...jsonIDict) jsonIDict {
	return jsonDictHelper.Merge(d, append([]jsonIDict{dict}, otherDicts...))
}

func (d jsonDict) Omit(key interface{}, otherKeys ...interface{}) jsonIDict {
	return jsonDictHelper.Omit(d, append([]interface{}{key}, otherKeys...))
}

func (d jsonDict) Set(key interface{}, v interface{}) jsonIDict {
	return jsonDictHelper.Set(d, key, v)
}

// Generic helpers to simplify physical implementation
func jsonListConvert(list jsonIList) jsonIList { return jsonList(list.AsArray()) }
func jsonDictConvert(dict jsonIDict) jsonIDict { return jsonDict(dict.AsMap()) }

var jsonHelper = helperBase{ConvertList: jsonListConvert, ConvertDict: jsonDictConvert}
var jsonListHelper = helperList{BaseHelper: jsonHelper}
var jsonDictHelper = helperDict{BaseHelper: jsonHelper}

// DictionaryHelper gives public access to the basic dictionary functions
var DictionaryHelper collections.IDictionaryHelper = jsonDictHelper

// GenericListHelper gives public access to the basic list functions
var GenericListHelper collections.IListHelper = jsonListHelper
