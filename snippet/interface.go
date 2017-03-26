package main

import (
    "fmt"
    "reflect"
    "strconv"
)

func main() {

    //interface类型
    //interface类型定义了一组方法，如果某个对象实现了某个接口的"所有方法"，则此对象就实现了此接口
    //interface可以被任意的对象实现,一个对象可以实现任意多个interface

    //任意的类型都实现了空interface(我们这样定义：interface{})，也就是包含0个method的interface。

    //interface的值
    /*
       mike := student{Human{"mike", 25}, "110"}
       paul := student{Human{"paul", 26}, "120"}
       lucy := employee{Human{"lucy", 18}, "001"}
       lily := employee{Human{"lily", 20}, "002"}

       //定义common类型的接口变量co
       var co common

       //co能够存储mike
       co = mike
       co.sayHi()
       co.sing()

       //co能够存储paul
       co = paul
       co.sayHi()
       co.sing()

       //co能够存储lucy
       co = lucy
       co.sayHi()
       co.sing()

       //co能够存储lily
       co = lily
       co.sayHi()
       co.sing()
    */

    //空interface
    //空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface。
    //空interface在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。
    /*
       var nullInterface interface{}
       var i int = 5
       var str string
       str = "Hello world"

       Jim := student{Human{"Jim", 27}, "101"}

       nullInterface = i
       nullInterface = str
       nullInterface = Jim

       //一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数，
       //如果一个函数返回interface{},那么也就可以返回任意类型的值。
       userInterfaceParam(nullInterface)
       fmt.Println("...")
    */

    //interface函数参数
    //任何实现了String方法的类型都能作为参数被fmt.Println调用
    //实现了error接口的对象（即实现了Error() string的对象），
    //使用fmt输出时，会调用Error()方法，因此不必再定义String()方法了

    //interface变量存储的类型
    //知道interface的变量里面可以存储任意类型的数值(该类型实现了interface)。
    //怎么反向知道这个变量里面实际保存了的是哪个类型的对象

    //Comma-ok断言

    //Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，
    //这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。

    //如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。

    //示例
    type Element interface{}
    type List []Element
    list := make(List, 3)
    list[0] = 1
    list[1] = "HelloWorld"
    list[2] = Human{"yang", 27}
    for index, element := range list {
        switch value := element.(type) {
        case int:
            fmt.Printf("list[%d] ,value is %d\n", index, value)
        case string:
            fmt.Printf("list[%d] ,value is %s\n", index, value)
        case Human:
            fmt.Printf("list[%d] ,value is %s\n ", index, value)
        default:
            fmt.Printf("list[%d] ,value is \n", index)
        }
    }

    //嵌入interface
    //如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的method。

    //反射
    //1:反射成reflect对象-->2:对reflect对象进行操作,比如获取它的值，或修改它的值
    //1:反射成reflect对象
    //t := reflect.TypeOf(i)    //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
    //v := reflect.ValueOf(i)   //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值

    //2:对reflect对象进行操作,引入reflect包
    //tag := t.Elem().Field(0).Tag  //获取定义在struct里面的标签
    //name := v.Elem().Field(0).String()  //获取存储在第一个字段里面的值

    //示例
    //获取值和类型
    var x float64 = 3.4
    v := reflect.ValueOf(x)
    fmt.Println("type:", v.Type())
    fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
    fmt.Println("value:", v.Float())

    //修改值 要使用引用
    var f float32 = 2.9
    ff := reflect.ValueOf(&f)
    ff.Elem().SetFloat(3.8)
    fmt.Println(f)

    //这们会出错
    //ff := reflect.ValueOf(f)
    //ff.SetFloat(3.8)
}

type Human struct {
    name string
    age  int
}

type student struct {
    Human
    schoolNumber string
}
type employee struct {
    Human
    employeeNumber string
}

func (h Human) sayHi() {
    fmt.Println("Hi!")
}

func (h Human) sing() {
    fmt.Println("la la la ~~")
}

func (s student) readBook() {
    fmt.Println(" reading book")
}

func (e employee) work() {
    fmt.Println("I'm working")
}

//Human、student、employee都实现了这个接口
type common interface {
    sayHi()
    sing()
}

//student实现了这个接口
type stuInterface interface {
    sayHi()
    sing()
    readBook()
}

//employee实现了这个接口
type empInterface interface {
    sayHi()
    sing()
    work()
}

//接收和返回interface类型,如果interface{}为空，那么它可以接收和返回任意类型的参数和值
func userInterfaceParam(i interface{}) interface{} {
    return i
}

func (h Human) String() string {
    return "(name: " + h.name + " - age: " + strconv.Itoa(h.age) + " years)"
}
