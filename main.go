package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"genshin_text_client/define"
	"github.com/flopp/go-findfont"
	"os"
	"strings"
)

func init() {
	fontPaths := findfont.List()
	for _, fontPath := range fontPaths {
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(fontPath, "simhei.ttf") {
			err := os.Setenv("FYNE_FONT", fontPath)
			if err != nil {
				return
			}
			break
		}
	}
}

var (
	MainApp fyne.App
	Window  fyne.Window
	Entry   *define.Ent // 多行文本框 用来展示页面
	MyRow   = 0         // 我的坐标
	MyCol   = 0
	VecRow  = 0 // 移动向量
	VecCol  = 0
)

func WindowsBaseSize(w fyne.Window) {
	w.CenterOnScreen() // 屏幕居中启动
	w.Resize(fyne.NewSize(720, 480))
	w.SetFixedSize(true) // 不允许修改窗口大小
}

// 负责将全局的ent new出来
func newEntry(tapped func()) *define.Ent {
	Entry = &define.Ent{
		Entry: widget.NewMultiLineEntry(),
	}
	Entry.ExtendBaseWidget(Entry)
	Entry.OnCursorChanged = tapped
	return Entry
}

// 填充entry 文本区域 （初始化兼刷新）
func fillText() {
	text := ""
	for i := 0; i < define.MapRow; i++ {
		for j := 0; j < define.MapCOL; j++ {
			text += define.Ref[define.Map[i][j]]
		}
		text += "\n"
	}
	Entry.SetText(text)
}

func main() {
	defer func() {
		err := os.Unsetenv("FYNE_FONT")
		if err != nil {
			fmt.Println("取消环境变量错误", err)
		}
	}()

	MainApp = app.New()
	Window = MainApp.NewWindow("Client")
	Window.SetOnClosed(func() {
		MainApp.Quit() // 主窗口退出时应用结束
	})
	WindowsBaseSize(Window)

	// 需要先将输入法切换到英文模式 才能捕获字母跟数字
	// 因为canvas是interface 无法传递指针 这里需要使用一个协程处理按键事件
	go HandleEnterKey(Window.Canvas())

	//r := canvas.NewRectangle(color.Gray{0x66})
	//r.Resize(fyne.NewSize(200, 200))
	//r.Refresh()

	Entry = newEntry(func() {
		// 每次窗口其中元素变化就会执行的方法
		//fmt.Println("my Entry cursor change")
	})
	Entry.Disable()

	fillText()
	Entry.Resize(fyne.NewSize(400, 400))

	//p := container.NewVBox()
	//p.Resize(fyne.NewSize(300, 300))
	//p.Add(widget.NewLabel("xxx"))
	//p.Add(Entry)

	Window.SetContent(Entry)

	Window.ShowAndRun()
}
