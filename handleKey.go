package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"genshin_text_client/define"
	"genshin_text_client/utils"
	"image/color"
)

func HandleEnterKey(cv fyne.Canvas) {
	cv.SetOnTypedKey(func(key *fyne.KeyEvent) {
		switch string(key.Name) {
		// 右键跟元素视野就不管了
		case "Escape":
			fmt.Println("Esc 打开主页/返回")
			// todo 新开主页窗口 或者返回（关闭非主窗口）
			cv.SetContent(widget.NewLabel("esc"))
		case "M":
			fmt.Println("M 打开地图")
		case "W":
			VecRow = -1
			VecCol = 0
			Move(VecRow, VecCol)
			//fmt.Println("W 向前移动")
		case "S":
			VecRow = 1
			VecCol = 0
			Move(VecRow, VecCol)
			//fmt.Println("S 向后移动")
		case "A":
			VecRow = 0
			VecCol = -1
			Move(VecRow, VecCol)
			//fmt.Println("A 向左移动")
		case "D":
			VecRow = 0
			VecCol = 1
			Move(VecRow, VecCol)
			//fmt.Println("D 向右移动")
		case "C":
			fmt.Println("C 角色界面")
		case "B":
			fmt.Println("B 背包界面")
		case "F":
			fmt.Println("F 拾取")
		case "J":
			fmt.Println("J 任务列表")
		case "Y":
			fmt.Println("Y 部分场景用到 比如同意联机/副本")
		case "O":
			fmt.Println("O 好友页面")
		case "P":
			fmt.Println("P 部分场景用到 如中断挑战")
		case "E":
			fmt.Println("E 元素战技")
		case "Q":
			fmt.Println("Q 元素爆发")
		case "L":
			fmt.Println("L 队伍配置")
		case "V":
			fmt.Println("V 追踪任务")
		case "Z":
			fmt.Println("Z 使用道具")
		case "X":
			fmt.Println("Z 取消攀爬")
		case "LeftShift":
			fmt.Println("LeftShift 冲刺")
		case "LeftAlt":
			fmt.Println("LeftAlt 展示鼠标")
		case "LeftControl":
			fmt.Println("LeftControl 行走/跑步切换")
		case "1":
			fmt.Println("1 切换角色1")
		case "2":
			fmt.Println("2 切换角色2")
		case "3":
			fmt.Println("3 切换角色3")
		case "4":
			fmt.Println("4 切换角色4")
		case "5":
			fmt.Println("5 切换角色5 如果存在的话")
		case "F1":
			fmt.Println("F1 冒险书")
		case "F2":
			fmt.Println("F2 联机")
		case "F3":
			fmt.Println("F3 祈愿")
			F3()
		case "F4":
			fmt.Println("F4 纪行")
		case "F5":
			fmt.Println("F5 活动")
		default:
			fmt.Println("enter key", key)
		}
	})
}

func Move(vecR, vecC int) {
	// 目的地坐标
	targetRow := MyRow + vecR
	targetCol := MyCol + vecC
	if targetRow < 0 || targetRow >= define.MapRow || targetCol < 0 || targetCol >= define.MapCOL {
		//	地图边界 todo 现有地图（蒙德璃月稻妻须弥四国）的动态加载 感觉好像很难...
		music := utils.MusicEntry{
			Id:         "1",
			Name:       "song.mp3",
			Artist:     "无所谓",
			Source:     "resource/sounds/map_lock.mp3",
			Type:       "mp3",
			Filestream: nil,
		}
		// 需要先调用os.open读取文件流后才能播放
		music.Open()
		go func() {
			// play中使用了select阻塞 这里用协程播放音乐
			music.Play()
		}()
	} else if define.Map[targetRow][targetCol] == 0 {
		define.Map[targetRow][targetCol] += 1
		define.Map[MyRow][MyCol] = 0
		// 更新角色位置
		MyRow += vecR
		MyCol += vecC

		// 刷新文本区域
		fillText()
	}
}

func F3() {
	wish := MainApp.NewWindow("祈愿")
	WindowsBaseSize(wish)
	wish.SetOnClosed(func() {
		Window.Show()
	})

	crystal := canvas.NewText("剩余原石: 1000", color.Black)
	// todo 纯文本界面
	draw1 := widget.NewButton("祈愿1次", func() {
		fmt.Println("祈愿1次")
	})
	draw10 := widget.NewButton("祈愿10次", func() {
		fmt.Println("祈愿10次")
	})

	// 顶部展示数据
	head := container.NewCenter(container.NewHBox(crystal))
	// 展示中间图片
	img := canvas.NewImageFromFile("resource/images/pray.png")
	imgContainer := container.NewGridWithColumns(3, widget.NewLabel(""), img, widget.NewLabel(""))
	imgContainer.Resize(fyne.NewSize(720, 480))
	// 抽卡按钮
	btns := container.NewBorder(layout.NewSpacer(), layout.NewSpacer(), draw1, draw10)
	btnCenter := container.NewCenter(btns)
	// 抽卡结果 新窗口展示...
	//entry := widget.NewMultiLineEntry()
	//entry.SetText("抽卡结果")
	//entry.Disable()

	// 使用这个就可以设置其中元素高度了 但是注意每个元素都堆在左上角 需要各元素都调用move方法偏移
	//contain := container.NewWithoutLayout(head, img, btnCenter)
	// 组合
	//contain := container.NewVBox(head, imgContainer, btnCenter)
	contain := container.NewGridWithRows(3, head, imgContainer, btnCenter)
	wish.SetContent(contain)

	wish.Show()
	Window.Hide()
}
