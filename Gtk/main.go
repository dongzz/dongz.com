package main

import (
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"unsafe"
)

type ChessWidget struct {
	window *gtk.Window
}

type ChessInfo struct {
	w, h int //宽高
	x, y int //坐标
}

type Chessboard struct {
	ChessWidget //匿名字段
	ChessInfo
}

//方法
func (obj *Chessboard) CreateWindow() {
	//加载glade文件
	builder := gtk.NewBuilder()
	builder.AddFromFile("/home/dongzhi/go/src/dongz.com/Gtk/chess.glade")

	//func NewWindowFromObject(obj *glib.GObject) *Window{
	//	return &Window{Bin{Container{Widget{C.toGWidget(obj.Object)}}}}
	//}

	obj.window = gtk.NewWindowFromObject(builder.GetObject("window1"))
	obj.window.SetAppPaintable(true) //允许绘图
	obj.window.SetPosition(gtk.WIN_POS_CENTER)
	obj.w, obj.h = 800, 480
	obj.window.SetSizeRequest(800, 480)

	obj.window.SetDecorated(false) //去边框

	//设置事件，让窗口捕获鼠标点击事件
	obj.window.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))

}

//鼠标点击时间
func MouseClickEvent(ctx *glib.CallbackContext) {}

//方法：事件，信号处理
func (obj *Chessboard) HandleSingal() {
	//鼠标点击事件
	obj.window.Connect("button-press-event", MouseClickEvent, obj)

	//鼠标移动事件
	obj.window.Connect("motion-notify-event", func(ctx *glib.CallbackContext) {
		args := ctx.Args(0)

		motion := *(**gdk.EventMotion)(unsafe.Pointer(&args))

		obj.window.Move(int(motion.XRoot)-obj.x, int(motion.YRoot)-obj.y)
	})

}

func main() {
	gtk.Init(&os.Args)

	//创建结构体变量
	var obj Chessboard

	obj.CreateWindow()
	obj.HandleSingal()

	obj.window.ShowAll()

	gtk.Main()

}
