package format

import (
	"github.com/gizak/termui/v3/widgets"
)

type listCtr struct {
	*widgets.List
	Meta []string
	cur int
}

func NewListCtr() *listCtr {
	list := widgets.NewList()
	return &listCtr{list, []string{}, list.SelectedRow}
}

func (l *listCtr) Pre() {
	if l.cur > 0 {
		l.cur--
	}
}

func (l *listCtr) Next() {
	if l.cur < len(l.Meta) - 1 {
		l.cur++
	}
}

func (l *listCtr) Get() interface{} {
	return l.Meta[l.cur]
}
