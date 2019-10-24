/**
 *    FILENAME      :   main.go
 *    AUTHOR        :   jihonghe
 *    DATE          :   19-10-24
 *    DESCRIPTION   :
 */
package main

import (
    "cloudDisk/handler"
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/file/upload", handler.FileUploadHandler)
    http.HandleFunc("/file/upload/success", handler.FileUploadSuccessHandler)
    http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
    http.HandleFunc("/file/query", handler.FileQueryHandler)
    http.HandleFunc("/file/download", handler.FileDownloadHandler)
    http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
    http.HandleFunc("/file/delete", handler.FileDeleteHandler)
    err := http.ListenAndServe(":8080", nil)

    if err != nil {
        fmt.Printf("Failed to start server, error: %s\n", err.Error())
    }
}
