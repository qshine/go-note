package main

import (
	"fmt"
	"reflect"
)

// reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息。
func reflectType(x interface{}) {
	res := reflect.TypeOf(x)
	fmt.Printf("type:%v  kind:%v\n", res.Name(), res.Kind())
}

func test_0() {
	var a float32 = 3.14
	reflectType(a)

	var b string = "hello world"
	reflectType(b)

	//type:float32  kind:float32
	//type:string  kind:string

}

// Name和Kind
func test_1() {
	var a *float32
	reflectType(a)

	type person struct {
		name string
		age  int
	}
	b := person{
		name: "xxx",
		age:  20,
	}
	reflectType(b)

	//type:  kind:ptr
	//type:person  kind:struct

}

// 通过反射获取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x) // 返回的是reflect.Value类型
	k := v.Kind()           // 通过Kind获取底层数据类型
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func test_2() {
	var a float32 = 3.14
	var b int64 = 100
	reflectValue(a) // type is float32, value is 3.140000
	reflectValue(b) // type is int64, value is 100
	// 将int类型的原始值转换为reflect.Value类型
	c := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", c) // type c :reflect.Value
}

// 通过反射设置变量的值
func test_3() {
	a := 100
	fmt.Println(&a)
	v := reflect.ValueOf(&a)
	fmt.Println(v)
	// 获取值
	fmt.Println(v.Elem())
	// 修改值
	v.Elem().SetInt(999)
	fmt.Println(a)

	//0xc0000b2008
	//0xc0000b2008
	//100
	//999
}

// isNil / isValid
func test_4() {
	var a *int
	fmt.Println(reflect.ValueOf(a).IsNil())     // true
	fmt.Println(reflect.ValueOf(nil).IsValid()) // false

	b := struct{}{}
	// 查找是否有name字段
	fmt.Println(reflect.ValueOf(b).FieldByName("name").IsValid()) // false
	// 是否有Say方法
	fmt.Println(reflect.ValueOf(b).MethodByName("Say").IsValid()) // false

	// map查找一个键
	c := map[string]interface{}{
		"name": "xxx",
	}
	fmt.Println(reflect.ValueOf(c).MapIndex(reflect.ValueOf("name")).IsValid()) // true
	fmt.Println(reflect.ValueOf(c).MapIndex(reflect.ValueOf("age")).IsValid())  // false
}

// 结构体反射
func test_5() {
	p1 := person{
		Name: "xxx",
		Age:  20,
	}
	r1 := reflect.TypeOf(p1)
	fmt.Println(r1)                   // main.person
	fmt.Println(r1.Name(), r1.Kind()) // person struct

	fmt.Println("====")

	// 遍历所有字段信息
	for i := 0; i < r1.NumField(); i++ {
		field := r1.Field(i)
		fmt.Printf("field_name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取结构体字段信息
	if scoreField, ok := r1.FieldByName("Name"); ok {
		// name:Name index:[0] type:string json tag:name
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}

func iterStructInfo(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println(t) // main.person
	v := reflect.ValueOf(x)
	fmt.Println(v) // {root 99}

	//Name -> root
	//Age -> 99
	for i := 0; i < v.NumField(); i++ {
		k := v.Field(i)
		switch k.Kind() {
		case reflect.String:
			fmt.Printf("%s -> %v\n", t.Field(i).Name, k.String())
		case reflect.Int:
			fmt.Printf("%s -> %v\n", t.Field(i).Name, k.Int())
		}
	}
}

func iterStructMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumField())
	for i := 0; i < t.NumField(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}

	/*
		2
		method name:Learn
		method:func()
		root is learning ...
		method name:Sleep
		method:func()
		root is sleep !
	*/
}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p person) Learn() {
	fmt.Printf("%s is learning ...\n", p.Name)
}

func (p person) Sleep() {
	fmt.Printf("%s is sleep !\n", p.Name)
}

// 使用反射遍历结构体
func test_6() {
	p := person{
		Name: "root",
		Age:  99,
	}
	iterStructInfo(p)
	fmt.Println("****")
	iterStructMethod(p)
}

func main() {
	//test_0()
	//test_1()
	//test_2()
	//test_3()
	//test_4()
	//test_5()
	test_6()
}
