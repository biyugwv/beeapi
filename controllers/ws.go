package controllers
import (
    //"beeapi/models"
    //"beeapi/img"
    "github.com/astaxie/beego"
    "github.com/gorilla/websocket"
    //"image"
    //"beeapi/lib"
    //"time"
    //"strconv"
    "fmt"
    "net/http"
    "beeapi/im"
    "time"
)




// Operations about Users
type WsController struct {
        beego.Controller
}



// @Title Join
// @Description  a member join  this chat room
// @Success 200  success
// @router /join [get]
func (u *WsController) Join(){
    
    SSID := u.GetString("SSID")
    if len(SSID) == 0 {
        fmt.Println("need SSID")
        return
    }
     // 检验http头中upgrader属性，若为websocket，则将http协议升级为websocket协议
    conn, err := (&websocket.Upgrader{
		// 允许跨域
		CheckOrigin : func(r *http.Request) bool{
			return true
		},

	}).Upgrade(u.Ctx.ResponseWriter, u.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
        fmt.Println("HandshakeError")
        return
    } else if err != nil {
        fmt.Println(err)
        return
    }
   var client im.Client
    client.SSID = SSID
    client.Conn = conn
    client.ConnectTime = time.Now().Unix()
    // 如果用户列表中没有该用户
    if _,ok := im.Clients[SSID] ; !ok  {
        im.Join <- client
    }else{
        fmt.Println("该用户已在线")
    }
    
    

}

// @Title Test
// @Description  It's a test uri
// @Success 200 {string} test success
// @router /test [get]
func (u *WsController) Test(){
    u.Data["json"] = map[string]interface{}{"msg":"test","code":"200"}
    u.ServeJSON()
}

