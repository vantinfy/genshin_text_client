package define

import (
	"fyne.io/fyne/v2"
)

var (
	MainApp fyne.App    // 全局主程序
	Window  fyne.Window // 全局主窗口
	Entry   *Ent        // 多行文本框 用来展示页面
	MyRow   = 0         // 我的坐标 todo 后面如果要实现地图动态加载 我的坐标应该是固定在屏幕中间 其它元素刷新（woc感觉好难）
	MyCol   = 0
	VecRow  = 0 // 移动向量
	VecCol  = 0
)
