package polymorphism

import "fmt"

type T struct {
	readOnly bool
}

type action interface {
	isOnlyRead() bool
	updateRole()
}

func (t *T) isOnlyRead() bool {
	return t.readOnly
}

func (t *T) updateRole() {
	t.readOnly = !t.readOnly
}

func NewT() {
	temp := &T{
		readOnly: false,
	}
	var _temp action = temp
	_temp.updateRole()
	fmt.Println(_temp.isOnlyRead())
}
