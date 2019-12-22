package im
 
import (
		"github.com/gorilla/websocket"
		"fmt"
		"encoding/json"
		"time"
		"strconv"
	)

const  (
    pingMaxTime = 12
)
type Client struct {
    Conn *websocket.Conn    // 用户websocket连接
    SSID string             // 登陆后的验证Token
    PingTime int64          // 心跳监测 客户端发起  5秒收不到自动断开
    PongTime int64          // 心跳监测 服务端发起每隔1秒左右发送一次
    Userid string           // 登录用户名
    ConnectTime int64
    lastMsgTime int64
    stat int
}

type Message struct {
    EventType int8  `json:"type"`       //  -2 表示pong； -1 表ping ； 0表示用户发布消息；1表示用户进入；2表示用户退出；
    Name string     `json:"name"`       // 用户名称
    Message string  `json:"message"`    // 消息内容
}
type msgGet struct {                 // 接收到的消息
    msgtype  int8 `json:"msgtype"` 
    msg   string `json:"msg"` 
    SSID  string `json:"SSID"` 
}


var (
    Clients map[string] Client
    Join  chan Client
    Leave chan Client
    Msg   chan Message
)

func readMsgLoop(){
    for {
        for _,client := range Clients {
            // 监听connection发送消息
            readMsg(client)
           
        }
        time.Sleep( 1 * time.Second )
    }
}

func channelLoop(){  // 三个channel为空 读取操作会造成阻塞
    for{
        
        select {
            case msg := <- Msg:
                for _,client := range Clients {
                    
                    data, err := json.Marshal(msg)
                    if err != nil {
                        fmt.Println("json.Marshal 错误")
                        return
                    }
                    // 转换成字符串类型便于查看
                    if err := client.Conn.WriteMessage(websocket.TextMessage, data) ; err != nil {
                        fmt.Println(err)
                    }
                    fmt.Println(string(data))
                }
            case client := <- Join:
                Clients[client.SSID] = client
                msg := Message{1,"system",fmt.Sprintf("%s加入了房间",client.SSID)}
                Msg<-msg
                
            case client := <- Leave:
                msg := Message{2,"system",fmt.Sprintf("%s离开了房间",client.SSID)}
                Msg<-msg
                delete(Clients, client.SSID)
                client.Conn.Close()
        }
        
    }
}


//  
func readMsg(client Client){
    now  := time.Now().Unix()
    // 心跳判断
    if client.PingTime == 0 {
        client.PingTime = now
    } 
    if now - client.PingTime>= pingMaxTime{  // 超时没有收到客户端的心跳监测消息 将处理掉该链接
        Leave <- client
        return 
    }
    
    var (
        data []byte
        err error
    )
    if _,data,err = client.Conn.ReadMessage() ; err!=nil {
        Leave <- client
        return 
    }
    msgget := &msgGet{}
    json.Unmarshal(data,msgget)
    fmt.Println(msgget)
    fmt.Println(string(data))
    // 心跳数据接收
    if msgget.msgtype == -1 {
        client.PingTime = now
        // 向客户端发送监测数据
        pong(client)
    }else if  msgget.msgtype == 0 {
        msg := Message{0,client.SSID,msgget.msg}
        Msg<-msg
    }else {
        
        fmt.Println("invilid data：%s",string(data))
    }

    // 向客户端发送消息 
    
}


func pong(client Client){
    now  := time.Now().Unix()
    client.PongTime = now
    msg := Message{-2,"system",strconv.FormatInt(now,10)}
    data, err := json.Marshal(msg)
    if err != nil {
        fmt.Println("pong err - 1")
        return
    }
    if err := client.Conn.WriteMessage(websocket.TextMessage, data) ;err != nil {
        
    }
}


func Run(){
    // 初始化
    Clients = make(map [string] Client)      
    Join = make(chan Client ,100)
    Leave = make(chan Client ,100)
    Msg = make(chan Message ,1000)

    go channelLoop()
    go readMsgLoop()
}
