package lib

import (
		"github.com/astaxie/beego"
        "github.com/astaxie/beego/cache"
        _ "github.com/astaxie/beego/cache/redis"
        "github.com/astaxie/beego/orm"
        "fmt"
        "time"
        "io"
        "crypto/md5"
        "crypto/aes"
        "encoding/hex"
        "crypto/cipher"
        "strings"
        "bytes"
)

func RedisGet(key string) string {
		redisurl := beego.AppConfig.String("redisurl")
		redispwd := beego.AppConfig.String("redispwd")
		redisconf :=  `{"conn":"` +redisurl+ `","password":"` +redispwd+ `"}`
        bm, err := cache.NewCache("redis", redisconf)
        if err != nil {
                fmt.Println(err)
        }
        v := bm.Get(key)
        if v == nil {
            return  ""
        }

        return  string(v.([]byte))
}

func RedisPut(key string,val string,timeout int64) bool{
        redisurl := beego.AppConfig.String("redisurl")
		redispwd := beego.AppConfig.String("redispwd")
		redisconf :=  `{"conn":"` +redisurl+ `","password":"` +redispwd+ `"}`
        bm, err := cache.NewCache("redis", redisconf)
        if err != nil {
                fmt.Println(err)
                return  false
        }
        timeoutDuration := time.Duration(timeout)  * time.Second
        if err = bm.Put(key,val,timeoutDuration) ;err==nil {
                return  true
        }
        return  false
}

func Sql(name string) orm.QuerySeter {
        o := orm.NewOrm() 
        tablename := "wfm_" + name
        qs := o.QueryTable(tablename)
        return qs
}

func Md5(str string) string {
        m := md5.New()
        _, err := io.WriteString(m, str)
        if err != nil {
                return  ""
        }
        arr := m.Sum(nil)
        return fmt.Sprintf("%x", arr)
}

func AESEncodeStr(src, key string) string {
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        fmt.Println("key error1", err)
    }
    if src == "" {
        fmt.Println("plain content empty")
    }
    ivspec := []byte("0000000000000000")
    ecb := cipher.NewCBCEncrypter(block, ivspec)
    content := []byte(src)
    content = PKCS5Padding(content, block.BlockSize())
    crypted := make([]byte, len(content))
    ecb.CryptBlocks(crypted, content)
    return hex.EncodeToString(crypted)
}

func AESDecodeStr(crypt, key string) string {
    crypted, err := hex.DecodeString(strings.ToLower(crypt))
    if err != nil || len(crypted) == 0 {
        fmt.Println("plain content empty")
    }
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        fmt.Println("key error1", err)
    }
    ivspec := []byte("0000000000000000")
    ecb := cipher.NewCBCDecrypter(block, ivspec)
    decrypted := make([]byte, len(crypted))
    ecb.CryptBlocks(decrypted, crypted)

    return string(PKCS5Trimming(decrypted))
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
    padding := encrypt[len(encrypt)-1]
    return encrypt[:len(encrypt)-int(padding)]
}
