package models

import (
   //    "errors"
//        "strconv"
//       "time"
      "github.com/astaxie/beego/orm"
      "beeapi/lib"
)



type Admin struct {
        Id       int32   `orm:"column(id);pk"`
        Username string
        Password string
}

func LoginNew(username ,password string) (ok bool,msg string){
	var admin Admin
	qs := lib.Sql("admin")
	err := qs.Filter("username__iexact",username).One(&admin)
	if err!=nil || admin.Username!= username {
		return false,"用户名不存在"
	}
	err = qs.Filter("username__iexact",username).Filter("password__iexact",lib.Md5(lib.Md5(password))).One(&admin)
	if err!=nil || admin.Username!= username {
		return false,"密码错误：" +"   "+ password +"   "+ lib.Md5(lib.Md5(password))
	}
	return true,""
}

func AdminList(name string)([]Admin,error){
    var admins  []Admin
   
    qs :=lib.Sql("admin")
    _,err := qs.Filter("username__iexact",name).Limit(10).All(&admins)
    if(err!=nil){

        return admins,err
    }else{
		for k,_ := range admins{
			admins[k].Password = ""
		}
        return admins,nil
    }
    
}

func init() {
        orm.RegisterModelWithPrefix("wfm_",new(Admin))
}
