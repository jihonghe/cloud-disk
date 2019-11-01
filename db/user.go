package db

import (
    mydb "cloudDisk/db/mysql"
    "fmt"
)

// 用户注册
func UserSignup(userName string, email string, password string) bool {
    stmt, err := mydb.DBConn().Prepare("INSERT IGNORE INTO user(user_name, email, password) values(?, ?, ?)")
    if err != nil {
        fmt.Println("Failed to insert new user, error: " + err.Error())
        return false
    }
    defer stmt.Close()
    
    result, err := stmt.Exec(userName, email, password)
    if err != nil {
        println("执行插入失败")
        return false
    }
    rowsAffected, err := result.RowsAffected()
    fmt.Printf("rows: %d, err: %v\n", rowsAffected, err)
    if err == nil && rowsAffected > 0 {
        return true
    }
    println("插入失败")
    return false
}

// 用户登录：主要校验用户名和密码是否正确
func UserSignin(userName string, encryptPassword string) bool {
    stmt, err := mydb.DBConn().Prepare("SELECT * FROM user WHERE user_name=? LIMIT 1")
    if err != nil {
        fmt.Println(err.Error())
        return false
    }
    defer stmt.Close()
    rows, err := stmt.Query(userName)
    if err != nil {
        fmt.Println(err.Error())
        return false
    } else if rows == nil {
        fmt.Println("userName not found: " + userName)
        return false
    }
    parseRows := mydb.ParseRows(rows)
    if len(parseRows) > 0 && string(parseRows[0]["password"].([]byte)) == encryptPassword {
        return true
    }
    return false
}

// 更新用户登录的token
func UpdateToken(userName string, token string) bool {
    stmt, err := mydb.DBConn().Prepare("REPLACE INTO user_token(user_name, user_token) values(?, ?)")
    if err != nil {
        fmt.Println(err.Error())
        return false
    }
    defer stmt.Close()
    _, err = stmt.Exec(userName, token)
    if err != nil {
        fmt.Println(err.Error())
        return false
    }
    return true
}

type User struct {
    UserName string
    Email string
    Phone string
    SignupTime string
    LastActive string
    Status int
}

func GetUserInformation(userName string) (User, error) {
    user := User{}
    
    stmt, err := mydb.DBConn().Prepare("SELECT user_name, signup_time FROM user WHERE user_name=? LIMIT 1")
    if err != nil {
        fmt.Println(err.Error())
        return user, err
    }
    defer stmt.Close()
    
    // 执行查询操作
    err = stmt.QueryRow(userName).Scan(&user.UserName, &user.SignupTime)
    if err != nil {
        return user, err
    }
    return user, err
}
