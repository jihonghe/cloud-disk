package handler

import (
    "net/http"
)

// http请求拦截器
func HttpInterceptor(handler http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(
        func(writer http.ResponseWriter, request *http.Request) {
            _ = request.ParseForm()
            userName := request.Form.Get("userName")
            token := request.Form.Get("token")
            
            if len(userName) < 3 || !IsValidToken(token) {
                writer.WriteHeader(http.StatusForbidden)
                return
            }
            // 验证通过后，开始执行请求
            handler(writer, request)
        })
}
