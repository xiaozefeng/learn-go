# Learning Golang
## 通用规则
```
变量、常量、方法名、结构体 首字母大写时代表可以被其他包访问，小写表示当前包私有
在go语言中，变量申明了必须使用，否则编译报错， 如果有时候要忽略方法的返回值使用 _ 接收就不会报错了
在go语言中，变量名在前，类型在后
在go语言中，只有值传递没有引用传递，需要改变传递的参数的值时，可以传递一个指针
rune 相当于go 的char
离线文档: godoc -http 4000
go tour: go get golang.org/x/tour/gotour 
使用 gotour
```
## Variables
### 变量类型
```
bool string
(u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
byte, rune
float32 , float64, complex64,  complex128
```
### 变量初始化
```go
// 使用默认初始化值
// 申明类型
var a, b int = 1,2
// 不申明类型
var a, b, c, d = 1, 2, true, "hello"
// 直接赋值 (:=)只有在初始化的时候可以使用
a, b :=  1, 2
// 包内变量
var (
	aa = 3
	ss = "hello"
	flag = true
)
```
### 类型转换
```go
// 没有隐式类型转换,类型之间需要强制类型转换
// float64 -> int 
var a int = 1
var b float64 = float64(a)
// int -> string
s := strconv.Itoa(1)
// string -> int
i, _:= strconv.Atoi(s)
// byte -> string or inversion
b := []byte("hello")
s := string(b)
```
### 常量
```go
// 申明常量
const pi = 3.14
// 数值类型的常量可以直接当值任意数据类型使用，不用强制转换
// 特殊的常量 -> 枚举
const(
    golang = iota
    java
    python
    javascript
)
// iota 安装一定的规则递增, 默认初始值为1, 每次递增1
// 硬盘容量枚举
// b, kb, mb, gb, tb, pb
const (
	b = 1<< (10*iota)
	kb
	mb
	gb
	tb
	pb
)
```

## Branch 流程
```go
// if , switch
// switch 不用添加break
// if 另一种语法
if file, err := os.Open(filename); err !=nil {
	fmt.Println(err)
}else {
	fmt.Println(file)
}
```

## Loop 循环
```go
// i 循环
for i:=0 ; i<10; i++ {}
var arr = [...]int{1,2,3,4,5}
// 遍历数组 index: 下标， value: 数据元素
for index, value range  arr{}
// 只要index
for index range arr {}
// 只要value
for _, value range arr{}
// 死循环
for {}
// 相当于while 
for 条件 {}
```

## Pointer 指针
```go
var a = 1
// 赋值 a的内存地址给 p, p是一个指针
p := &a
// 给a 重新赋值
*p = 2
// 打印一下a 的值
fmt.Println(a)
```

## Function 函数
```go
// 普通函数
func abs(){}
// 有返回值
func abs(number int) int 
// 多返回值
func swap(a, b int) (int ,int ){
    return b, a
}
// 函数可以作为参数或者返回值
func apply(op func(int, int) (int, int), a,b int) (int ,int ){
    return op(a, b)
}
// 函数式编程
// 函数式一等公民: 参数，变量，返回值都可以是函数
```

## Array 数组
// 在golang 中 array 是值类型
```go
// 申明有5个空间的数据
var arr1 [5] int
arr2 := [3]int {1,3,5}
arr3 := [...]int {2,4,6,8,10}
// 多维数据
var grid  [4][5]int
```

## Slice(切片)
```go
// slice 底层是一个数组, 所以slice 只是对数据的一个view
// slice 有三个重要的属性
// data -> 底层数组
// length -> 数组的长度
// capacity -> 容量
var arr [...]int {0,1,2,3,4,5,6,7}
// slice
s1 := arr[2:6]
s2 := arr[2:]
s3 := arr[:6]
s4 := arr[:]
// ReSlice
s5 := s1[3:5]
// 当s5超过了切片时，超过了s1的范围时，不会报错，因为底层的数组
// ReSlice时是可以向后兼容的，但是不向前面兼容


// slice 常用操作
// 添加元素
s1 = append(s1, 10)
s1 = append(s1, 20)
// 使用append，当添加的数据超过slice的capacity时，会重新分配一个数组

// copy slice
copy(s1,s2)

// delete element from slice
s1 = append(s1[:2], [3:])

// pop from front
front := s1[0]
s1 = s1[1:]
// pop from end
end := s1[len(s1)-1]
s1 = s1[:len(s1)-1]

```

## Map 字典
```go
// 定义
m := map[string]string{
    "name":"jack",
    "age": "18",
}

m1 := make(map[string]int) // empty map

var m2 map[string]int // zero value is nil

// 遍历
for k, v range m {}
for k range m {}
for _,v range m{}
// 根据key 获取 value
v := m[k]
// 判断可以是否存在
if v, ok := m[k] ; ok{
	fmt.Println(v)
}else {
	fmt.Println("key doesn't exist")
}

// delete element
delete(m, key)

```

## struct 结构体
```go
// 定义 Node 为结构体名称
type Node struct{
	// 属性
	value int
	left, right *Node
}
// 初始化
var node Node
// 空对象
node := &Node{}
node := new(Node) // 和上面是等值的
// 赋值并初始化
node := &Node{value: 3}

// 定义method
func (node Node) print(){
	fmt.Println(node.value)
}
node.print()  //使用
// 需要注意的是，这里的node 是值传递, 如果要传指针 
func (node *Node) setValue(value int){
	node.value = value
}
// 如果可以用指针接收就尽量使用指针接受

```

### interface 接口
```go
// 定义接口
type Writer interface{
	write() bool
}
//  duck typing 是动态类型的一种风格
// 当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子
// 在这种风格中，一个对象有效的语义，不是由继承自特定的类或实现特定的接口，而是由"当前方法和属性的集合"决定
type FileWriter struct{
	content string
}
// FileWriter 的方法
func (w FileWriter) write() bool{
	return true
}
// 可以说FileWriter 实现了Writer接口
// 接口组合
type Writer interface{
	write()
}
type Reader interface{
	read()
}
// 组合
type WriterAndReader interface {
	Writer
	Reader
}
// 内置接口
// Stringer
// Reader/Writer
```

### 资源管理与错误处理
```go
// defer  确保调用在函数结束时发生 (相当于finally)
// 使用场景 :  Open/Close , Lock/Unlock , PrintHeader/PrintFooter
// 多个defer 是先进后出的
for i := 0; i < 3; i++ {
		defer fmt.Print(i," ")
}
// 输出 2 1 0

// 错误处理
// error 接口
type error interface {
	Error() string
}
// 正确处理error
if err != nil {
    if pathError, ok:= err.(*os.PathError); !ok{
        panic(err)
    }else {
        fmt.Printf("%s, %s, %s", pathError.Op, pathError.Path, pathError.Err)
    }
    return
}
```

### test 测试
```go
// golang 是表格驱动测试的
// 编写测试
// 有一个 add.go 编写的测试文件叫 add_test.go
// 需要测试的方法为 add(), 测试的方法为 TestAdd(t *testing.T)
// example
func TestAdd(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{1, 2, 3},
		{2, 2, 4},
		{12, 21, 33},
	}

	for _, tt := range tests {
		if actual := add(tt.a, tt.b); actual != tt.c {
			t.Errorf("add(%d, %d); got %d; expected %d ",
				tt.a, tt.b, actual, tt.c)
		}
	}
}
// 命令行执行测试
go test .
// 使用test.B 性能测试 
func BenchmarkAdd(b *testing.B) {
	a, c := 1, 2
	ans := 3
	actual := add(a, c)
	for i := 0; i < b.N; i++ {
		if ans != actual {
			b.Errorf("got %d for input %d expected %d", actual, ans, actual)
		}
	}
}
// 命令行工具来查看性能损耗
// 生成cpu.out文件
go test -bench . -cpuprofile cpu.out
// 进入pprof交互式环境
go tool pprof cpu.out
// 查看web, 需要安装 graphviz https://www.graphviz.org/
web

// 编写example
// Example
func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	// Output:
	// 1
	// 2
	// false
	// 3
	// true

}
```



