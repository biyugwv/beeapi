package controllers
import (
    "beeapi/models"
    "beeapi/img"
    "github.com/astaxie/beego"
    "image"
    "beeapi/lib"
    "time"
    "strconv"
)

// Operations about Users
type AdminController struct {
        beego.Controller
}



// @Title Get
// @Description find admin 
// @Param       name                path    string  true            "the admin you want to get"
// @Success 200 {admin} models.Admin
// @Failure 403 :name is empty
// @router /:name [get]
func (u *AdminController) Get() {
    name := u.GetString(":name")
    ssid := u.GetString("ssid")
    if myssid := lib.RedisGet(ssid) ; myssid == "" && name!="login" {
        u.Data["json"] = map[string]interface{}{"err":"SSID错误或缺失","code":"200","ssid":ssid};
    }else{

        if(name  == "verifyimg"){
            u.verifyimg()

        }else if(name == "login"){
            u.login()
        }else{
             if name != "" {
                list, err :=models.AdminList(name)
                if err != nil {
                        u.Data["json"] = err.Error()
                } else {
                        //u.Data["json"] = list
                        u.Data["json"] =map[string]interface{}{"ssid":myssid,"code":"200","list":list};
                }
            }else{
                u.Data["json"] = map[string]interface{}{"err":"参数错误","code":"200"};
            }

        }
    }
    u.ServeJSON()
}



func (u *AdminController) verifyimg() {
      u.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", u.Ctx.Request.Header.Get("Origin"))
    i :=  img.VerifyImg{200,60,20,"",&image.RGBA{}}
    base64 := i.Create()

    u.Data["json"] = map[string]interface{}{"img":base64,"code":"200"}

}

func (u *AdminController) login() {
    uname := u.GetString("name")
    pwd := u.GetString("pwd")
    ok,msg := models.LoginNew(uname,pwd)
    if ok {
        ts := strconv.FormatInt(time.Now().Unix(),10)
        createssid := uname+"_"+ ts
        ok = lib.RedisPut(lib.Md5(createssid),createssid,7200)
        if(ok){
            u.Data["json"] = map[string]interface{}{"msg":"登录成功","ssid":lib.Md5(createssid),"_ssid":createssid,"code":"200"}
        }else{
            u.Data["json"] = map[string]interface{}{"msg":"登录失败","code":"200"}
        }
    }else{
            u.Data["json"] = map[string]interface{}{"msg":msg,"code":"200"}
    }
}

