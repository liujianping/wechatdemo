# wechat demo

## 开发

````
package main 

import "fmt"
import "github.com/liujianping/wechat"
import "github.com/liujianping/wechat/entry"

type Echo struct{
	Name string
	wechat.Callback
}
func NewEcho(name string) *Echo{
	return &Echo{name, wechat.Callback{}}
}
func (e *Echo) MsgText(txt *entry.TextRequest, back chan interface{}){
	wechat.Info("Echo: MsgText ", txt)
}
func (e *Echo) MsgImage(img *entry.ImageRequest, back chan interface{}){
	wechat.Info("Echo: MsgImage ", img)
}
func (e *Echo) MsgVoice(voice *entry.VoiceRequest, back chan interface{}){
	wechat.Info("Echo: MsgVoice ", voice)
}
func (e *Echo) MsgVideo(video *entry.VideoRequest, back chan interface{}){
	wechat.Info("Echo: MsgVideo ", video)
}
func (e *Echo) MsgLink(link *entry.LinkRequest, back chan interface{}){
	wechat.Info("Echo: MsgLink ", link)
}
func (e *Echo) Location(location *entry.LocationRequest, back chan interface{}){
	wechat.Info("Echo: Location ", location)
}

func (e *Echo) EventSubscribe(appoid string, oid string, back chan interface{}){
	wechat.Info("Echo: EventSubscribe ", oid)
	var subscriber entry.Subscriber
	if err := e.Api.GetSubscriber(oid, &subscriber); err != nil {
		wechat.Error("Echo: get subscriber failed ", err)
	}

	response := entry.NewTextResponse(appoid, oid, fmt.Sprintf("%s 欢迎您的关注!", subscriber.Nickname))

	back <- response
}
func (e *Echo) EventUnsubscribe(appoid string, oid string, back chan interface{}){
	wechat.Info("Echo: EventUnsubscribe ", oid)	

}
func (e *Echo) EventMenu(appoid string, oid string, key string, back chan interface{}){
	wechat.Info("Echo: EventMenu ", oid, key)	
}

func main() {
	wechat.SetLogger("console", "")
	app := wechat.NewWeChatApp()
	app.SetConfig("ini", "demo.ini")	
	app.SetCallback(NewEcho("demo"))

	//! 添加菜单
	menu := entry.NewMenu()
	btn1 := entry.NewViewButton("新浪","http://sina.com")
	btn2 := entry.NewClickButton("点击","EVENT_MENU_CLICK")
	btn3 := entry.NewButton("更多")
	btn3.Append(entry.NewViewButton("腾讯","http://qq.com"))
	btn3.Append(entry.NewViewButton("百度","http://baidu.com"))
	btn3.Append(entry.NewViewButton("google","http://google.com"))
	menu.Add(btn1)
	menu.Add(btn2)
	menu.Add(btn3)

	app.SetMenu(menu)


	app.Run()
}

````

## 编译

	go get github.com/liujianping/wechat
	go get github.com/liujianping/wechatdemo
	cd src/github.com/liujianping/wechatdemo
	go build

## 运行

	修改demo.ini配置之前，首先注册一个测试号：

	https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login

	设置好测试服务号配置后，相应设置demo.ini配置

	服务程序不能在本机运行，微信服务必须是具有域名服务的HTTP服务。
