# go-copier : Deep copy any struct

## desc ##
 copier 为了在 golang 中进行不同定义，但结构体（变量名）完全相同的两个 struct

## install ##

 go get github.com/q294043308/go-copier

 ## usage ##

```
type S11 struct {
    Des string
}

type S1 struct {
	ID    int64
    Info  *S11
}

type S22 struct {
    Des string
}

type S2 struct {
	ID    int64
    Info  *S22
}

s1 := new(S1)
s2 := &S2{
    ID: 1,
    Info: &S22 {
        Des: "hello world",
    }
}

copier.Copy(s1, s2)

// s1 <==> s2
```

## support ##

 copier 支持基本类型拷贝
 copier 支持自定义类型拷贝
 copier 支持 prt/array/selice/map 类型拷贝
 copier 支持上述所有嵌套式深拷贝

## License ##

This package is licensed under MIT license. See LICENSE for details.