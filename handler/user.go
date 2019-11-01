package handler

import (
    dblayer "cloudDisk/db"
    "cloudDisk/utils"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

const (
    passwordSalt = "helloworld"
)

// 处理用户注册请求
func UserSignupHandler(writer http.ResponseWriter, request *http.Request) {
    if request.Method == http.MethodGet {
        data, err := ioutil.ReadFile("../static/view/signup.html")
        if err != nil {
            writer.WriteHeader(http.StatusInternalServerError)
            return
        }
        _, _ = writer.Write(data)
        return
    } else if request.Method == http.MethodPost {
        _ = request.ParseForm()
        userName := string(request.Form.Get("userName"))
        email := string(request.Form.Get("email"))
        password := string(request.Form.Get("password"))
        fmt.Printf("userName: %s, email: %s, password: %s\n", userName, email, password)
        // 基础校验，待改进
        if len(userName) < 3 || email == "" || len(password) < 5{
            _, _ = writer.Write([]byte("用户名或密码长度太短"))
            return
        }
    
        encryptPassword := utils.Sha1([]byte(password + passwordSalt))
        println(encryptPassword)
        success := dblayer.UserSignup(userName, email, encryptPassword)
        if success {
            _, _ = writer.Write([]byte("SUCCESS"))
        } else {
            _, _ = writer.Write([]byte("FAILED"))
        }
    }
}

// 处理用户登录请求
func UserSignInHandler(writer http.ResponseWriter, request *http.Request) {
    if request.Method == http.MethodGet {
        data, err := ioutil.ReadFile("../static/view/signin.html")
        if err != nil {
            writer.WriteHeader(http.StatusInternalServerError)
            return
        }
        _, _ = writer.Write(data)
        return
    } else if request.Method == http.MethodPost {
        // 校验用户名和密码
        _ = request.ParseForm()
        userName := request.Form.Get("userName")
        encryptPassword := utils.Sha1([]byte(request.Form.Get("password") + passwordSalt))
        accountValid := dblayer.UserSignin(userName, encryptPassword)
        if !accountValid {
            _, _ = writer.Write([]byte("FAILED"))
            return
        }
        // 生成访问凭证（token）
        token := GetToken(userName)
        success := dblayer.UpdateToken(userName, token)
        if !success {
            _, _ = writer.Write([]byte("FAILED"))
            return
        }
        resp := utils.RespMsg{
            Code: 0,
            Msg: "OK",
            Data: struct{
                Location string
                UserName string
                Token string
            }{
                Location: "http://" + request.Host + "/static/view/home.html",
                UserName: userName,
                Token: token,
            },
        }
        _, _ = writer.Write(resp.JSONBytes())
    } else {
        _, _ = writer.Write([]byte("Unsupported request method."))
        return
    }
}

// 查询用户信息
func UserInformationHandler(writer http.ResponseWriter, request *http.Request) {
    // 解析请求参数
    _ = request.ParseForm()
    userName := request.Form.Get("userName")
    //token := request.Form.Get("token")
    // 验证Token
    //isValidToken := IsValidToken(token)
    //if !isValidToken {
    //    writer.WriteHeader(http.StatusForbidden)
    //    return
    //}
    // 查询用户信息
    user, err := dblayer.GetUserInformation(userName)
    if err != nil {
        writer.WriteHeader(http.StatusForbidden)
        return
    }
    // 封装用户数据
    resp := utils.RespMsg{
        Code: 0,
        Msg: "OK",
        Data: user,
    }
    _, _ = writer.Write(resp.JSONBytes())
}

// 生成token
func GetToken(userName string) string {
    // 取40位字符：md5(userName + timestamp + tokenSalt) + timestamp[:8]
    timeStamp := fmt.Sprintf("%x", time.Now().Unix())
    tokenPrefix := utils.MD5([]byte(userName + timeStamp + "_tokenSalt"))
    
    return tokenPrefix + timeStamp[:8]
}

// 验证token是否有效
func IsValidToken(token string) bool {
    if len(token) != 40 {
        return false
    }
    // 判断token的时效性，是否过期
    // 从数据库表user_token查询userName对应的token
    // 对比两个token是否一致
    return true
}

