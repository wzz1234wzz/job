package study

import (
	"fmt"
	"reflect"
)

//定义一个Enum类型
type Enum int
const (
	Zero Enum = 0
)

type Student struct {
	Name string
	Age  int
	float32
	bool
	next *Student
}

func add(a, b int) int { //普通函数
	return a + b
}

type MyMath struct {
	Pi float64
}

//普通函数
func (myMath MyMath) Sum(a, b int) (int,int) {
	return a + b,a*b
}

func (myMath MyMath) Dec(a, b int) int {
	return a - b
}

func ReflectMain() {
	fmt.Println("----------------------常量-------------------------")
	typeOfZero := reflect.TypeOf(Zero) // 获取Zero常量的反射类型对象
	fmt.Printf("Zero：%v\n",typeOfZero)
	fmt.Println("Zero的类型名：",typeOfZero.Name())
	//fmt.Println("Zero的字段数：",typeOfZero.NumField())
	fmt.Println("Zero的类 型：",typeOfZero.Kind())

	fmt.Println("-----------------------结构体------------------------")
	var stu Student
	typeOfStu := reflect.TypeOf(stu) // 使用 reflect.TypeOf() 函数可以获得任意值的类型对象（reflect.Type）
	fmt.Printf("typeOfStu：%v\n",typeOfStu)
	fmt.Println("Student的类型名：",typeOfStu.Name())
	fmt.Println("Student的字段数：",typeOfStu.NumField())
	fmt.Println("Student的类 型：",typeOfStu.Kind())

	fmt.Println("----------------------指针-------------------------")
	var stu1 = &Student{Name:"kitty", Age: 20} //定义一个Student类型的指针变量
	typeOfStu1 := reflect.TypeOf(stu1) //获取结构体实例的反射类型对象
	fmt.Printf("name: '%v', kind: '%v'\n", typeOfStu1.Name(), typeOfStu1.Kind()) //显示反射类型对象的名称和种类
	typeOfStu1 = typeOfStu1.Elem() //取类型的元素。Go 程序中对指针获取反射对象时，可以通过 reflect.Elem() 方法获取这个指针指向的元素类型。这个获取过程被称为取元素，等效于对指针类型变量做了一个*操作
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfStu1.Name(), typeOfStu1.Kind()) //显示反射类型对象的名称和种类

	fmt.Println("----------------------结构体元素-------------------------")
	type cat struct { // 声明一个空结构体
		Name string
		Type int `json:"type" id:"100"` // 带有结构体tag的字段
	}
	ins := cat{Name: "mimi", Type: 1} // 创建cat的实例
	typeOfCat := reflect.TypeOf(ins) // 获取结构体实例的反射类型对象
	for i := 0; i < typeOfCat.NumField(); i++ { // 遍历结构体所有成员
		fieldType := typeOfCat.Field(i) // 获取每个成员的结构体字段类型
		fmt.Printf("fieldType: %v\n", fieldType) // 输出成员名和tag
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag) // 输出成员名和tag
		fmt.Printf("PkgPath: %v  Type: '%v'\n", fieldType.PkgPath, fieldType.Type) // 输出成员名和tag
	}
	if catType, ok := typeOfCat.FieldByName("Type"); ok { // 通过字段名, 找到字段类型信息
		fmt.Printf("catType：%v\n",catType)
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))// 从tag中取出需要的tag
	}

	fmt.Println("----------------------取值-------------------------")
	var a int = 1024 //声明整型变量a并赋初值
	valueOfA := reflect.ValueOf(a) //获取变量a的反射值对象
	fmt.Printf("valueOfA:%v,%T\n",valueOfA,valueOfA)
	fmt.Printf("valueOfA.Interface():%v,%T\n",valueOfA.Interface(),valueOfA.Interface())
	fmt.Printf("valueOfAint:%T,%T\n",valueOfA.Int(),int(valueOfA.Int()))
	var getA int = valueOfA.Interface().(int) //获取interface{}类型的值，通过类型断言转换
	var getB int = int(valueOfA.Int()) //获取64位的值，强制类型转换为int类型
	var getC float64 = float64(valueOfA.Int()) //获取64位的值，强制类型转换为float64类型
	fmt.Println(getA, getB,getC)

	fmt.Println("----------------------访问结构体成员的值-------------------------")
	rValue := reflect.ValueOf(Student{ //值包装结构体
		next: &Student{},
	})
	fmt.Println("NumField:", rValue.NumField())//获取字段数量
	//获取索引为2的字段(float32字段) 注:经过测试发现Field(i)的参数索引是从0开始的，并且是按照定义的结构体的顺序来的，而不是按照字段名字的ASCii码值来的
	floatField := rValue.Field(2)
	fmt.Println("Field:", floatField.Type()) //输出字段类型
	fmt.Println("FieldByName(\"Age\").Type:", rValue.FieldByName("Age").Type()) //根据名字查找字段
	fmt.Println("FieldByIndex([]int{4, 1}).Type()", rValue.FieldByIndex([]int{4, 1}).Type()) //根据索引查找值中next字段的int字段的值

	fmt.Println("----------------------判断反射值得空和有效性-------------------------")
	var a1 *int //*int的空指针,IsNil() 常被用于判断指针是否为空；IsValid() 常被用于判定返回值是否有效。
	fmt.Println("var a *int:", reflect.ValueOf(a1).IsNil())
	fmt.Println("nil:", reflect.ValueOf(nil).IsValid()) //nil值
	fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid()) //*int类型的空指针
	s := struct {}{} //实例化一个结构体
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(s).FieldByName("").IsValid()) //尝试从结构体中查找一个不存在的字段
	fmt.Println("不存在的方法:", reflect.ValueOf(s).MethodByName("").IsValid()) //尝试从结构体中查找一个不存在的方法
	m := map[int]int{} //实例化一个map
	fmt.Println("不存在的键:", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid()) //尝试从map中查找一个不存在的键

	fmt.Println("----------------------通过反射修改变量的值-------------------------")
	var a2 int = 1024 //声明整形变量a并赋初值
	rValue2 := reflect.ValueOf(&a2) //获取变量a的反射值对象
	fmt.Printf("reflect.ValueOf(&a2):%v\n",reflect.ValueOf(&a2)) //地址
	rValue2 = rValue2.Elem() //取出a地址的元素(a的值)
	fmt.Printf("rValue2.Elem():%v\n",rValue2) //打印a的值
	rValue2.SetInt(1) //尝试将a修改为1
	fmt.Println(rValue2.Int()) //打印a的值
	fmt.Println("a2=",a2) //打印a的值
	type dog struct {
		LegCount int
	}
	valueOfDog := reflect.ValueOf(&dog{})//获取dog实例的反射值对象
	valueOfDog = valueOfDog.Elem()
	fmt.Printf("valueOfDog:%v\n",valueOfDog) //{0}
	vLegCount := valueOfDog.FieldByName("LegCount") //获取legCount字段的值
	fmt.Printf("vLegCount:%v, %T\n",vLegCount,vLegCount) //0
	vLegCount.SetInt(4) //尝试设置legCount的值(这里会发生崩溃)
	fmt.Println(vLegCount.Int())

	fmt.Println("----------------------通过类型信息创建实例-------------------------")
	var a3 int
	typeOfA := reflect.TypeOf(a3) //取变量a的反射类型对象
	aIns := reflect.New(typeOfA) //根据反射类型对象创建类型实例
	fmt.Println(aIns.Type(), aIns.Kind()) //输出Value的类型和种类

	fmt.Println("----------------------通过反射调用函数-------------------------")
	funcValue := reflect.ValueOf(add) //将函数包装为反射值对象
	fmt.Printf("funcValue:%v, %T\n",funcValue,funcValue) //0
	paramList := []reflect.Value{reflect.ValueOf(2), reflect.ValueOf(3)} //构造函数参数，传入两个整形值
	fmt.Printf("paramList:%v, %T\n",paramList,paramList) //
	retList := funcValue.Call(paramList) //反射调用函数
	fmt.Println(retList[0].Int())

	fmt.Println("----------------------通过反射调用方法-------------------------")
	var myMath = MyMath{Pi:3.14159}
	rValue3 := reflect.ValueOf(myMath) //获取myMath的值对象
	numOfMethod := rValue3.NumMethod()  //获取到该结构体有多少个方法
	fmt.Printf("numOfMethod:%v, %T\n",numOfMethod,numOfMethod) //
	paramList2 := []reflect.Value{reflect.ValueOf(30), reflect.ValueOf(20)} //构造函数参数，传入两个整形值
	//调用结构体的第一个方法Method(0) 注意:在反射值对象中方法索引的顺序并不是结构体方法定义的先后顺序
	//而是根据方法的ASCII码值来从小到大排序，所以Dec排在第一个，也就是Method(0)
	result := rValue3.Method(1).Call(paramList2)
	fmt.Printf("result:%v, %T\n",result,result) //
	fmt.Println(result[0].Int())
	fmt.Println(result[1].Int())
}
