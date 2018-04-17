{% include navigation.html %}
{% raw %}
# Data collections (lists/slices and dicts/maps)

## Maps

| Razor                                             | Gotemplate                                            | Note
| ---                                               | ---                                                   | ---
| `@razorDict := dict("test", 1, "test2", 2);`      | `{{- set $ "goDict" (dict "test" 1 "test2" 2) }}`     | Creation
| `@razorDict2 := dict("test3", 3, "test5", 5);`    | `{{- set $ "goDict2" (dict "test3" 3 "test5" 5) }}`   | Creation
| `@set(.razorDict, "test2", 3);`                   | `{{- set .goDict "test2" 3 }}`                        | Update
| `@set(.razorDict, "test3", 4);`                   | `{{- set .goDict "test3" 4 }}`                        | Update 
| `@razorDict := merge(.razorDict, .razorDict2);`   | `{{- set $ "goDict" (merge .goDict .goDict2) }}`      | Merge (First dict has priority)
| `@razorDict.test3;`                               | `{{ $.goDict.test3 }}`                                | Should be `4`
| `@razorDict.test5;`                               | `{{ $.goDict.test5 }}`                                | Should be `5`


### Looping

#### Razor
```go
@range($key, $value := .razorDict)
    @$key, @$value, @get($.razorDict, $key)
@endrange
```

#### Gotemplate
```go
{{- range $key, $value := .goDict }}
    {{ $key }}, {{ $value }}, {{ get $.goDict $key }}
{{- end }}
```

#### Result
```go
    test, 1, 1
    test2, 3, 3
    test3, 4, 4
    test5, 5, 5
```

## Slices

| Razor                                            | Gotemplate                                             | Note
| ---                                              | ---                                                    | ---
| `@razorList := list("test1", "test2", "test3");` | `{{- set $ "goList" (list "test1" "test2" "test3") }}` | Creation
| `@razorList := append(.razorList, "test4");`     | `{{- set $ "goList" (append .goList "test4") }}`       | Append
| `@razorList := prepend(.razorList, "test0");`    | `{{- set $ "goList" (prepend .goList "test0") }}`      | Prepend
| `@razorList;`                                    | `{{ $.goList }}`                                       | Should be `[test0 test1 test2 test3 test4]`

### Looping

#### Razor
```go
@range($index, $value := .razorList)
    @$index, @$value, @extract($.razorList, $index)
@endrange
```

#### Gotemplate
```go
{{- range $index, $value := .goList }}
    {{ $index }}, {{ $value }}, {{ extract $.goList $index }}
{{- end }}
```

#### Result
```go
    0, test0, test0
    1, test1, test1
    2, test2, test2
    3, test3, test3
    4, test4, test4
```


{% endraw %}