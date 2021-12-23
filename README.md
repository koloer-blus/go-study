# go-study

## 基础

1. 基本命令
```shell
go run main.go # >>> 执行go文件
go build main.go # >>> 生成可执行的二进制文件
go get {url} # >>> 获取网上的代码

```
2. 名词解释

- 包（`package`）:  一个包由位于单个目录下的一个或多个.go源代码文件组成，目录定义包的作用；
  - `main`包：定义一个可以独立执行的函数，而不是一个库，是程序的入口

3. 变量赋值

- 定义变量
  - `var`：定义变量
  - `const`：定义常量
- 初始化隐形赋值
  - `string` >>> ""(空字符串)
  - `int` >>> 0
  ```go
  s := ""   // >>> 简短声明，只能在函数内部声明
  var s string
  var s = ""
  var s string = ""
  ```
- `map`：`Go`的`map`类似于`Java`语言中的`HashMap`

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

## 实战

1. [gee-web](./gee-web)