package logger

import "fmt"

type readwrite string

type Log struct {
	name       string
	content    string
	createtime string
}

type write interface {
	echo()
	out()
}

type read interface {
	scan()
	input()
}

func (rw *readwrite) echo() {
	fmt.Println("read write:echo()")
}

func (rw *readwrite) out() {
	fmt.Println("read write:out()")
}

func (rw *readwrite) scan() {
	fmt.Println("read write: scan()")
}

func (rw *readwrite) input() {
	fmt.Println("read write: input()")
}

func Init() {
	var readwrite readwrite
	readwrite.echo()

	log := &Log{
		name:       "test-log",
		content:    "test-info",
		createtime: "2021-12-10",
	}
	var _io *iowrite
	var _net *netwrite
	log.writeLog(_io)
	log.writeLog(_net)
}

type NetLog struct {
	Log
}
type IOLog struct {
	Log
}

type iowrite string
type netwrite string

func (iow *iowrite) echo() {
	fmt.Println("iowrite: echo()")
}

func (iow *iowrite) out() {
	fmt.Println("iowrite: out()")
}

func (netw *netwrite) echo() {
	fmt.Println("netwrite: echo()")
}

func (netw *netwrite) out() {
	fmt.Println("netwrite: out()")
}

func (l *Log) writeLog(w write) {
	fmt.Println(l.name + "---" + l.content)
}
