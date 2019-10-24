/**
 *    FILENAME      :   handler.go
 *    AUTHOR        :   jihonghe
 *    DATE          :   19-10-24
 *    DESCRIPTION   :
 */
package handler

import (
    "cloudDisk/meta"
    "cloudDisk/utils"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "strconv"
    "time"
)

func FileUploadHandler(writer http.ResponseWriter, request *http.Request) {
    if request.Method == "GET" {
        // 返回上传文件的页面
        data, err := ioutil.ReadFile("../static/view/index.html")
        if err != nil {
            _, _ =io.WriteString(writer, "Internal server error")
            return
        }
        _, _ = io.WriteString(writer, string(data))
    } else if request.Method == "POST" {
        // 接收文件流并存储到本地目录中
        file, head, err := request.FormFile("file")
        if err != nil {
            fmt.Printf("Failed to get data, error: %s\n", err.Error())
            return
        }
        defer file.Close()

        fileMeta := meta.FileMeta{
            FileName: head.Filename,
            Location: "/tmp/" + head.Filename,
            UploadTime: time.Now().Format("2006-01-02 15:04:05"),
        }

        newFile, err := os.Create(fileMeta.Location)
        if err != nil {
            fmt.Printf("Failed to create file, error: %s\n", err.Error())
            return
        }
        defer newFile.Close()

        fileMeta.FileSize, err = io.Copy(newFile, file)
        if err != nil {
            fmt.Printf("Failed to save data into file, error: %s\n", err.Error())
            return
        }

        _, _ = newFile.Seek(0, 0)
        fileMeta.FileSha1 = utils.FileSha1(newFile)
        meta.UpdateFileMeta(fileMeta)

        http.Redirect(writer, request, "/file/upload/success", http.StatusFound)
    }
}

func UploadFileSuccessHandler(writer http.ResponseWriter, request *http.Request) {
    _, _ = io.WriteString(writer, "Upload file successfully.")
}

func GetFileMetaHandler(writer http.ResponseWriter, request *http.Request) {
    _ = request.ParseForm()
    fileHash := request.Form["fileHash"][0]
    fileMeta := meta.GetFileMeta(fileHash)
    data, err := json.Marshal(fileMeta)
    if err != nil {
        writer.WriteHeader(http.StatusInternalServerError)
        return
    }
    _, _ = writer.Write(data)
}

func FileQueryHandler(writer http.ResponseWriter, request *http.Request) {
    _ = request.ParseForm()
    limitCount, _ := strconv.Atoi(request.Form.Get("limit"))
    fileMetas := meta.GetLastFileMetas(limitCount)
    data, err := json.Marshal(fileMetas)
    if err != nil {
        writer.WriteHeader(http.StatusInternalServerError)
        return
    }
    _, _ = writer.Write(data)
}

func FileDownloadHandler(writer http.ResponseWriter, request *http.Request) {
    _ = request.ParseForm()
    fileSha1 := request.Form.Get("fileHash")
    fileMeta := meta.GetFileMeta(fileSha1)
    file, err := os.Open(fileMeta.Location)
    if err != nil {
        writer.WriteHeader(http.StatusInternalServerError)
        return
    }
    defer file.Close()

    fileData, err := ioutil.ReadAll(file)
    if err != nil {
        writer.WriteHeader(http.StatusInternalServerError)
        return
    }
    writer.Header().Set("Content-Type", "application/octect-stream")
    writer.Header().Set("content-disposition", "attachment;filename=\"" + fileMeta.FileName + "\"")
    _, _ = writer.Write(fileData)
}
