# go-study

## golang

### 基础

1. 基本命令
```shell
go run main.go # >>> 执行go文件
go build main.go # >>> 生成可执行的二进制文件
go get {url} # >>> 获取网上的代码

```
2. 名词解释

- 包（`package`）:  一个包由位于单个目录下的一个或多个.go源代码文件组成，目录定义包的作用；
  - `main`包：定义一个可以独立执行的函数，而不是一个库，是程序的入口
  - 每一个包都有对应独立的命名空间
  - 通过名字大小写控制变量的私有性（变量名小写字母开头代表小写）
  - 包的第一句注释应当是报的功能概要说明
  - 包的初始化可以使用`func init() {}`来执行，该函数同时不支持被调用或引用，在程序开始执行时按照声明顺序依次调用

3. 程序结构

- 声明(<var|const> name [types] [= value])
  - `var`：定义变量
  - `const`：定义常量
    - 常量表达式的值在编译期计算，而不是在运行期。
    - 每种常量的潜在类型都是基础类型：boolean、string或数字。
    - iota
      - 常量声明可以使用iota常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式
      - 在第一个声明的常量所在的行，iota将会被置为0，然后在每一个有常量声明的行加一。
  - `type`：类型
    - 对于每一个类型T，都有一个对应的类型转换操作T(x)，用于将x转为T类型（译注：如果T是指针类型，可能会需要用小括弧包装T，比如`(*int)(0)`）。只有当两个类型的底层基础类型相同时，才允许这种转型操作，或者是两者都是指向相同底层结构的指针类型，这些转换只改变类型而不会影响值本身。
    - 数值类型之间的转型也是允许的，并且在字符串和一些特定类型的slice之间也是可以转换的，在下一章我们会看到这样的例子。这类转换可能改变值的表现。例如，将一个浮点数转为整数将丢弃小数部分，将一个字符串转为[]byte类型的slice将拷贝一个字符串数据的副本。在任何情况下，运行时不会发生转换失败的错误（错误只会发生在编译阶段）。
  - `func`：函数实体
- 简短声明
  -`:=`：`temp := 2`
- 初始化隐形赋值
    - 字符串
      - `string` >>> ""(空字符串)
      - Byte切片
    - 整型
      - `int` >>> 0
      - int8,int16,int32,int64 >>> 分别对应8、16、32、64bit大小的有符号整数
      - uint8,uint16,uint32,uint64 >>> 分别对应8、16、32、64bit大小的无符号整数
      - 其他运算符号：
 ```go
&      位运算 AND
|      位运算 OR
^      位运算 XOR
&^     位清空（AND NOT）
<<     左移
>>     右移
```
   - 浮点数
     - float32/float64
     - 一个float32类型的浮点数可以提供大约6个十进制数的精度，而float64则可以提供约15个十进制数的精度；通常应该优先使用float64类型，因为float32类型的累计计算误差很容易扩散，并且float32能精确表示的正整数并不是很大
   - 复数
     - complex64/complex128
     - > 我们把形如 z=a+bi（a、b均为实数）的数称为复数。
       > 其中，a 称为实部，b 称为虚部，i 称为虚数单位。
       > 当 z 的虚部 b＝0 时，则 z 为实数；当 z 的虚部 b≠0 时，实部 a＝0 时，常称 z 为纯虚数。
       > 复数域是实数域的代数闭包，即任何复系数多项式在复数域中总有根。
     - 分别对应float32和float64两种浮点数精度。
     - 复数也可以用==和!=进行相等比较。只有两个复数的实部和虚部都相等的时候它们才是相等的
     - 内置的complex函数用于构建复数，内建的real和imag函数分别返回复数的实部和虚部：
   ```go
var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i
fmt.Println(x*y)                 // "(-5+10i)"
fmt.Println(real(x*y))           // "-5"
fmt.Println(imag(x*y))           // "10"
```
  
```go
    s := ""   // >>> 简短声明，只能在函数内部声明
    var s string
    var s = ""
    var s string = ""
```

- 数组
  - 默认情况下，数组的每个元素都被初始化为元素类型对应的零值，对于数字类型来说就是0。
```go
var a [3]int
a[len(a) -1]
```
  - 在数组字面值中，如果在数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始化值的个数来计算。
```go
q := [...]int{1,2,5}
r := [...]int{9: 10}
```
- slice切片
  - Slice（切片）代表变长的序列
  - 序列中每个元素都有相同的类型。一个slice类型一般写作[]T，其中T代表slice中元素的类型；slice的语法和数组很像，只是没有固定长度而已。
  - 组成：
    - 指针：指向第一个slice元素对应的底层数组元素的长度的地址（slice第一个元素不等于数组第一个元素）
    - 长度：slice元素的数量（`len()`获取slice长度）
    - 容量：长度不能超过容量，容量一般时从slice的开始位置到底层数据的结尾位置（`cap()`获取slice容量）
  - 切片操作:**`s[i:j]`，其中0 ≤ i≤ j≤ cap(s)**
    - 新的slice将只有j-i个元素。
    - 如果i位置的索引被省略的话将使用0代替
    - 如果j位置的索引被省略的话将使用len(s)代替
  - 创建切片：
    - `make([]T, len ,[cap])` cap可省略
    - 在底层，make创建了一个匿名的数组变量，然后返回一个slice；只有通过返回的slice才能引用底层匿名的数组变量
  - 复制一个slice只是对底层的数组创建了一个新的slice别名
  - 和数组不同的是，slice之间不能比较，因此我们不能使用==操作符来判断两个slice是否含有全部相等元素。
> slice不支持比较运算符的原因
> 
> - 第一个原因，一个slice的元素是间接引用的，一个slice甚至可以包含自身（当slice声明为[]interface{}时，slice的元素可以是自身）。虽然有很多办法处理这种情形，但是没有一个是简单有效的。 
> 
> - 第二个原因，因为slice的元素是间接引用的，一个固定的slice值（指slice本身的值，不是元素的值）在不同的时刻可能包含不同的元素，因为底层数组的元素可能会被修改。而例如Go语言中map的key只做简单的浅拷贝，它要求key在整个生命周期内保持不变性（例如slice扩容，就会导致其本身的值/地址变化）。而用深度相等判断的话，显然在map的key这种场合不合适。对于像指针或chan之类的引用类型，==相等测试可以判断两个是否是引用相同的对象。一个针对slice的浅相等测试的==操作符可能是有一定用处的，也能临时解决map类型的key问题，但是slice和数组不同的相等测试行为会让人困惑。因此，安全的做法是直接禁止slice之间的比较操作。
```go
o := [...]string{1: "A", 2: 'B', 3: "C", 4: "D", 5: "E"}
o1 = o[1:4]
o1 = o[1:]

//判断数组为空
len(s) == 0

```
- `map`：
  - `Go`的`map`类似于`Java`语言中的`HashMap`,一个map就是一个哈希表的引用，map类型可以写为map[K]V，其中K和V分别对应key和value。
  - 创建map
    - make
    - 字面量
  - 删除元素
    - delete
  - map中的元素并不是一个变量，因此我们不能对map的元素进行取址操作
  - 要想遍历map中全部的key/value对的话，可以使用range风格的for循环实现，和之前的slice遍历语法类似。
  - Map的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序。在实践中，遍历的顺序是随机的，每一次遍历的顺序都不相同。这是故意的，每次都使用随机的遍历顺序可以强制要求程序不会依赖具体的哈希函数实现。如果要按顺序遍历key/value对，我们必须显式地对key进行排序，可以使用sort包的Strings函数对字符串slice进行排序。
  - map类型的零值是nil
  - 和slice一样，map之间也不能进行相等比较
  - 通过key作为索引下标来访问map将产生一个value。如果key在map中是存在的，那么将得到与key对应的value；如果key不存在，那么将得到value对应类型的零值
```go
args := map[string]int {
	"Alice": 25,
	"Tom": 35
}

//相当于
args := make(map[string]int)
args["Alice"] = 25
args["Tom"] = 35

//删除
delete(args, "Alice")

//使用key作为索引下标访问元素
args, ok := args["Bob"]
```
- 结构体
  - 结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。每个值称为结构体的成员。
  - 字面量语法
    - 要求以结构体成员定义的顺序为每个结构体成员指定一个字面值。它要求写代码和读代码的人要记住结构体的每个成员的类型和顺序，不过结构体成员有细微的调整就可能导致上述代码不能编译。因此，上述的语法一般只在定义结构体的包内部使用，或者是在较小的结构体中使用
    - 以成员名字和相应的值来初始化，可以包含部分或全部的成员
  - 结构体比较
    - `==`比较结构体的每个成员
  - 结构体匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针
```go
type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position    string
	Salary  int
	ManagerID int
}
var dilbert Employee
//取地址
position := &dilbert.Position
*position = "Senior" + *position

//结构体字面值
type Point struct {
	X,Y int
}

p := Point{1,2}

anim := git.GIF{LoopCount: nframes}

// 匿名变量
type Point struct {
	X,Y int
}

type Circle struct {
	Center Point
	Radius  int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

var w Wheel
w.Circle.Center.X = 8
w.Circle.Center.Y = 8
w.Circle.Radius = 5


type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

var w Wheel

w.X = 8
w.Y = 9
w.Radius = 5
w.Spokes = 20


```
- JSON
  - 结构体的成员Tag可以是任意的字符串面值，但是通常是一系列用空格分隔的key:"value"键值对序列；因为值中含有双引号字符，因此成员Tag一般用原生字符串面值的形式书写
  - 编码的逆操作是解码，对应将JSON数据解码为Go语言的数据结构，Go语言中一般叫unmarshaling，通过json.Unmarshal函数完成。下面的代码将JSON格式的电影数据解码为一个结构体slice，结构体中只有Title成员。通过定义合适的Go语言数据结构，我们可以选择性地解码JSON中感兴趣的成员。当Unmarshal函数调用返回，slice将被只含有Title信息的值填充，其它JSON成员将被忽略。
- 函数
  - 函数的类型被称为函数的签名。
  - Go语言没有默认参数值，也没有任何方法可以通过参数名指定形参，因此形参和返回值的变量名对于函数调用者而言没有意义。
  - 引用类型实参：切片、map、function、channel等
  - 在Go中，一个函数可以返回多个值。
  - 如果一个函数所有的返回值都有显式的变量名，那么该函数的return语句可以省略操作数
  - 通常，当函数返回non-nil的error时，其他的返回值是未定义的（undefined），这些未定义的返回值应该被忽略
  - Go使用控制流机制（如if和return）处理错误
  - 可变参数
    - 加上`...`
  - Deferred函数
    - 你只需要在调用普通函数或方法前加上关键字defer，就完成了defer所需要的语法。当执行到该条语句时，函数和参数表达式得到计算，但直到包含该defer语句的函数执行完毕时，defer后的函数才会被执行，不论包含defer语句的函数是通过return正常结束，还是由于panic导致的异常结束。你可以在一个函数中执行多条defer语句，它们的执行顺序与声明顺序相反。
  - 匿名函数
    - 当匿名函数需要被递归调用时，我们必须首先声明一个变量（在上面的例子中，我们首先声明了 visitAll），再将匿名函数赋值给这个变量。如果不分成两步，函数字面量无法与visitAll绑定，我们也无法递归调用该匿名函数。
  - Panic异常
  > 不加区分的恢复所有的panic异常，不是可取的做法；因为在panic之后，无法保证包级变量的状态仍然和我们预期一致。比如，对数据结构的一次重要更新没有被完整完成、文件或者网络连接没有被关闭、获得的锁没有被释放。此外，如果写日志时产生的panic被不加区分的恢复，可能会导致漏洞被忽略。
    - 在运行时检查，如数组访问越界、空指针引用等。这些运行时错误会引起painc异常。
    - 不是所有的panic异常都来自运行时，直接调用内置的panic函数也会引发panic异常
    - 一般而言，当panic异常发生时，程序会中断运行，并立即执行在该goroutine中被延迟的函数（defer 机制）。随后，程序崩溃并输出日志信息。日志信息包括panic value和函数调用的堆栈跟踪信息。panic value通常是某种错误信息。对于每个goroutine，日志信息中都会有与之相对的，发生panic时的函数调用堆栈跟踪信息。通常，我们不需要再次运行程序去定位问题，日志信息已经提供了足够的诊断依据。因此，在我们填写问题报告时，一般会将panic异常和日志信息一并记录。
  - Recover异常捕获
    - 如果在deferred函数中调用了内置函数recover，并且定义该defer语句的函数发生了panic异常，recover会使程序从panic中恢复，并返回panic value。
    - deferred函数帮助Parse从panic中恢复。在deferred函数内部，panic value被附加到错误信息中；并用err变量接收错误信息，返回给调用者。我们也可以通过调用runtime.Stack往错误信息中添加完整的堆栈调用信息。
```go
func CountWordsAndImages(url string) (words, images int, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return
    }
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        err = fmt.Errorf("parsing HTML: %s", err)
        return
    }
    words, images = countWordsAndImages(doc)
    return
}

//按照返回值列表的次序，返回所有的返回值，在上面的例子中，每一个return语句等价于：

return words, images, err

// 匿名函数

strings.Map(func (r rune) rune {return r +1}, "HAL-9000")

visitAll := func(items []string) {
// ...
visitAll(m[item]) // compile error: undefined: visitAll
// ...
}
```
- 指针
  - 一个指针的值是另一个变量的地址。一个指针对应变量在内存中的存储位置。并不是每一个值都会有一个内存地址，但是对于每一个变量必然有对应的内存地址。
  - 对于聚合类型每个成员——比如结构体的每个字段、或者是数组的每个元素——也都是对应一个变量，因此可以被取地址。 
  - 任何类型的指针零值都是`nil`

```go
x := 1
p := &x
fmt.println(*p)
```
- `new`函数
  - 表达式`new(T)`将创建一个T类型的匿名变量，初始化为`T`类型的零值，然后**返回变量地址**，**返回的指针类型为`*T`**。
  - new函数类似是一种语法糖，而不是一个新的基础概念
  - 每次调用`new`函数都会返回一个新的变量地址，如果两个类型都是空的，也就是说类型的大小是0，例如`struct{}`和`[0]int`，有可能有相同的地址（依赖具体的语言实现）,请谨慎使用大小为0的类型，因为如果类型的大小为0的话，可能导致Go语言的自动垃圾回收器有不同的行为
```go
p := new(int)
fmt.println(*p)

p := new(int)
q := new(int)
fmt.println(p == q) //false
```

4. 循环
 - `for...range variable {}`

```go
  for a,b := range "hello" {
    fmt.println(a,b)
  }
```

5. `Printf`动词展示

```
%d  十进制整数
%x,%o,%b  十六进制，八进制，二进制
%f,%g,%e  浮点数
%t  bool
%c  char
%s  string
%q  string or char
%v  变量的默认形式
%T  变量的类型
%%  %  
```

6. Go语言的自动垃圾收集器
> 从每个包级的变量和每个当前运行函数的每一个局部变量开始，通过指针或引用的访问路径遍历，是否可以找到该变量。如果不存在这样的访问路径，那么说明该变量是不可达的，也就是说它是否存在并不会影响程序后续的计算结果。




## 实战

1. [gee-web](./gee-web)

## 框架

### Iris

1. Get `Iris`

- 使用命令行
```shell
# install iris v12
go get github.com/kataras/iris/v12@lastest
```
- 修改 `go.mod`

```go
module your_project_name

go 1.14

require (
    github.com/kataras/iris/v12 v12.1.8
)
```

2. 入门

- 创建新的服务
  - `iris.New()`：返回一个可配置的`iris.Application`实例
  - `app.Run`：第一个参数为需要启动的服务或者监听的Listener，第二个及之后的参数是可选的`iris.Configurator`配置参数。**在`iris`中每一个核心的模块（视图引擎、websocket、session等）都有一个内部的`iris.Configurator`**
  - `app.Listen`：传入需要监听的端口号
```go
app := iris.New()
// 这行代码会启动一个服务并监听 localhost:8080或者127.0.0.1:8080
app.Listen(":8080")


/*-------------------------------------*/

// 当然，想要获取完整的`http.Server`的实例的使用也是可以的
import "net/http"

// 这里实现的功能和上面的代码没有差异，但是可以确保你使用完整的`http.Server`实例
server := &http.Server{Addr: ":8080"}
app.Run(iris.Server(server))

/*-------------------------------------*/

// 还有更好的用法就是使用自定义的`net.Listener`

listener, err := net.Listen("tcp4", ":8080")
if err != nil {
	panic(err)
}
app.Run(iris.Listener(listener))
```
- 关闭服务并停止默认行为

```go
package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"time"
)

func main() {
	app := iris.New().
		iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		app.Shutdown(ctx)
	})
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Closed</h1>")
    })
	// iris 配置
	confit := iris.WithConfiguration(iris.Configuration{
		DisableStartupLog: true,
		Charset: "UTF-8",
    })
	app.Run(iris.Addr(":8080"), config)
}

```

3. Router

- Handler Type(请求处理器)
  - 处理过程：响应Http请求 -> 写入响应头和数据到`Context.ResponseWriter()` -> 返回信号 -> 请求处理完成
  - 注意事项：
    - 提前读取`Context.Request().Body`中的数据，因为**在写入`Context.ResponseWriter()`后无法访问`Context.Request().Body`**
    - Handler不因改变传入的Context
    - 服务器出现`panic(异常)`，服务器会认为当前的panic的影响与运行的请求无关。会重启当前的panic，并且记录栈追踪日志到服务器错误日志同时关闭连接。
- 行为
  - `iris`默认接受和注册形如`/api/login`的路径路由，并且尾部不带斜杠。
  - 如果尝试访问`/api/login/`，将会自动永久重定向到`/api/login`
- API
  - 参数：(HTTP方法，请求的路径，多个`iris.Handler`)
```go
app := iris.New()

app.Handle("GET", "/contact", func(ctx iris.Context) {
	ctx.HTML("<h1>Hello World</h1>")
})

app.Get("/", func(ctx iris.Context) {
	ctx.HTML("<h1>hello</h1>")
})
```
- 路由组（`party`）
  - 通过对路由的路径前缀进行分组，共享相同的中间件和模板。
  - 写法1：
```go
app := iris.New()
users := app.Party("/user", handler)
users.Get("/{id:uint64}/info", handler1)
user.Get("/login", handler2)
```
  - 写法2：
```go
app.PartyFunc("/user", func(user, iris.Party) {
	user.Use(AuthMiddleware)
	user.Get("/lgoin", handler2)
})
```
- 路径参数
  - `/user/{id: string}`：`user/*`
  - `/user/{name: path}` ：`/user/**/*`
- 中间件（执行过程类似与`nodejs Express`框架）
  - 中间件仅是一个 Handler 格式的函数 `func(ctx iris.Context)`，当前一个中间件调用 `ctx.Next()` 方法时，此中间件被执行，这可以用作身份验证，即如果请求验证通过，就调用 `ctx.Next()` 来执行该请求剩下链上的处理器，否则触发一个错误响应。
- 处理`http`错误
  - Iris 内建支持 HTTP APIs 的错误详情。 
  - Context.Problem 编写一个 JSON 或者 XML 问题响应，行为完全类似 Context.JSON，但是默认 ProblemOptions.JSON 的缩进是 " "，响应的 Content-type 为 application/problem+json。 
  - 使用 options.RenderXML 和 XML 字段来改变他的行为，用 application/problem+xml 的文本类型替代。

4. MVC

![](https://www.topgoer.com/static/Iris/mvc.png)

控制器结构体内部的模型(在方法函数中设置，并通过视图渲染)。你可以从一个控制器的方法中返回模型，或者在请求的声明周期中设置一个字段，在同一个请求的生命周期中的另一个方法中返回这个字段。

就像你以前使用的流程一样，MVC 程序有自己的 Router，这是 iris/router.Party 类型的，标准的 iris api Controllers 可以被注册到任何 Party 中，包括子域名，Party 的开始和完成处理器与预期的一样工作。

可选的 BeginRequest（ctx） 函数，用于在方法执行之前执行任何初始化，这对调用中间件或许多方法使用相同的数据收集很有用。

可选的 EndRequest（ctx）函数， 可在执行任何方法之后执行任何终结处理。

递归继承，例如 我们的mvc会话控制器示例具有 Session * sessions.Session 作为字段，由会话管理器的 Start 填充为MVC应用程序的动态依赖项：mvcApp.Register(sessions.New(sessions.Config{Cookie:"iris_session_id"}).Start）

通过控制器方法的输入参数访问动态路径参数，不需要绑定。当你使用 Iris 的默认语法从一个控制器中解析处理器，你需要定义方法的后缀为 By，大写字母是新的子路径。