package copier

type S1111 int64
type S2222 int64

// S1 be copied
type S1 struct {
	ObjID    int64
	Age      int32
	TmpAdd   string
	TmpAdd1  string
	Name     *string
	ObjInfo  *S11
	Extre    []*S11
	Extre1   map[string]*S11
	Extre2   map[string]string
	Extre3   map[string]S1111
	BaseInfo []byte
}

// S2 copier
type S2 struct {
	ObjID    int64
	Age      int32
	TmpAdd   string
	TmpAdd2  string
	Name     *string
	ObjInfo  *S22
	Extre    []*S22
	Extre1   map[string]*S22
	Extre2   map[string]string
	Extre3   map[string]S2222
	BaseInfo []byte
}

// S11 be copied
type S11 struct {
	Info   string
	Childs []*S111
	Child1 []S111
}

// S22 copier
type S22 struct {
	Info   string
	Childs []*S222
	Child1 []S222
}

// S111 be copied
type S111 struct {
	Info string
	Val  S1111
	Val1 *S1111
}

// S222 copier
type S222 struct {
	Info string
	Val  S2222
	Val1 *S2222
}

func main() {
	s1 := new(S1)
	hel := "hel"
	var val S2222 = 3
	s2 := &S2{
		ObjID: 1,
		Name:  &hel,
		Age:   1,
		ObjInfo: &S22{
			Info: "hello world",
			Childs: []*S222{
				&S222{
					Info: "dxx",
					Val:  1,
					Val1: &val,
				},
			},
			Child1: []S222{
				S222{
					Info: "dxx",
					Val:  1,
					Val1: &val,
				},
			}},
		Extre: []*S22{&S22{
			Info: "hello world",
			Childs: []*S222{
				&S222{
					Info: "dxx",
					Val:  1,
					Val1: &val,
				},
			},
			Child1: []S222{
				S222{
					Info: "dxx",
					Val:  1,
					Val1: &val,
				},
			}}, &S22{
			Info: "hello world1",
			Childs: []*S222{
				&S222{
					Info: "dxx",
					Val:  1,
					Val1: &val,
				},
			},
			Child1: []S222{
				S222{
					Info: "dxx",
					Val:  1,
					Val1: &val,
				},
			}}, nil},
		Extre1: map[string]*S22{"hello": &S22{
			Info: "hello world",
			Childs: []*S222{
				&S222{
					Info: "dxx",
					Val:  1,
					Val1: &val,
				},
			},
			Child1: []S222{
				S222{
					Info: "dxx",
					Val:  1,
					Val1: &val,
				},
			}}, "hello1": &S22{
			Info: "hello world",
			Childs: []*S222{
				&S222{
					Info: "dxx",
					Val:  1,
					Val1: &val,
				},
			},
			Child1: []S222{
				S222{
					Info: "dxx",
					Val:  1,
					Val1: &val,
				},
			}}, "hello2": nil},
		Extre2:   map[string]string{"hello": "world"},
		Extre3:   map[string]S2222{"hi": 123},
		BaseInfo: []byte{1, 1, 1, 1, 1},
		TmpAdd:   "not",
		TmpAdd2:  "not not "}
	Copy(s1, s2)

	if s1.ObjID != s2.ObjID {
		panic("err")
	}

	if s1.TmpAdd != s2.TmpAdd {
		panic("err")
	}

	if s1.TmpAdd1 == s2.TmpAdd2 {
		panic("err")
	}

	if *(s1.Name) != *(s2.Name) {
		panic("err")
	}

	if s1.Age != s2.Age {
		panic("err")
	}

	if len(s1.Extre) != len(s2.Extre) {
		panic("err")
	}

	for i := 0; i < len(s1.Extre); i++ {
		if s2.Extre[i] == nil {
			continue
		}

		if s1.Extre[i].Info != s2.Extre[i].Info {
			panic("err")
		}

		if len(s1.Extre[i].Childs) != len(s2.Extre[i].Childs) {
			panic("err")
		}

		for j := 0; j < len(s1.Extre[i].Childs); j++ {
			if s2.Extre[i].Childs[j] == nil {
				continue
			}

			if s1.Extre[i].Childs[j].Info != s2.Extre[i].Childs[j].Info {
				panic("err")
			}

			if int64(s1.Extre[i].Childs[j].Val) != int64(s2.Extre[i].Childs[j].Val) {
				panic("err")
			}

			if int64(*(s1.Extre[i].Childs[j].Val1)) != int64(*(s2.Extre[i].Childs[j].Val1)) {
				panic("err")
			}
		}

		for j := 0; j < len(s1.Extre[i].Child1); j++ {
			if s1.Extre[i].Child1[j].Info != s2.Extre[i].Child1[j].Info {
				panic("err")
			}

			if int64(s1.Extre[i].Child1[j].Val) != int64(s2.Extre[i].Child1[j].Val) {
				panic("err")
			}

			if int64(*(s1.Extre[i].Child1[j].Val1)) != int64(*(s2.Extre[i].Child1[j].Val1)) {
				panic("err")
			}
		}
	}

	if len(s1.Extre1) != len(s2.Extre1) {
		panic("err")
	}

	for k := range s1.Extre1 {
		if s2.Extre1[k] == nil {
			continue
		}

		if s1.Extre1[k].Info != s2.Extre1[k].Info {
			panic("err")
		}
	}

	for k := range s1.Extre2 {
		if s1.Extre2[k] != s2.Extre2[k] {
			panic("err")
		}
	}

	for i := 0; i < len(s1.BaseInfo); i++ {
		if s1.BaseInfo[i] != s2.BaseInfo[i] {
			panic("err")
		}
	}

	println("copy finish")
}
