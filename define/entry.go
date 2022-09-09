package define

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type Ent struct { // 继承文本框结构 重写监听鼠标按键方法
	*widget.Entry
	focus bool // 获得焦点
}

func (e *Ent) Tapped(pe *fyne.PointEvent) {
	fmt.Println("左键单击")
}

func (e *Ent) DoubleTapped(pe *fyne.PointEvent) {
	fmt.Println("左键双击")
}

func (e *Ent) TappedSecondary(pe *fyne.PointEvent) {
	fmt.Println("右键单击")
}

func (e *Ent) MouseIn(event *desktop.MouseEvent) {
	fmt.Println("mouse in")
	e.focus = true
}

func (e *Ent) MouseOut() {
	fmt.Println("mouse out")
	e.focus = false
}

func (e *Ent) MouseMoved(event *desktop.MouseEvent) {
}

// func (e *Ent) RefreshText(key *fyne.KeyEvent) {}
