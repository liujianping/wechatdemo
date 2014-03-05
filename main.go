package main 

import "github.com/liujianping/wechat"
import "github.com/liujianping/wechat/entry"

type Echo struct{
	wechat.Callback
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
func (e *Echo) EventSubscribe(oid string, back chan interface{}){
	wechat.Info("Echo: EventSubscribe ", oid)
}
func (e *Echo) EventUnsubscribe(oid string, back chan interface{}){
	wechat.Info("Echo: EventUnsubscribe ", oid)	
}
func (e *Echo) EventMenu(oid string, key string, back chan interface{}){
	wechat.Info("Echo: EventMenu ", oid, key)	
}

func main() {
	wechat.SetLogger("console", "")
	app := wechat.NewWeChatApp()
	app.SetConfig("ini", "demo.ini")	
	app.SetCallback(new(Echo))
	app.Run()	
}