package db

import (
    mydb "cloudDisk/db/mysql"
    "database/sql"
    "fmt"
)

// 保存上传文件的meta信息
func OnFileUploadFinished(fileHash string, fileName string, fileSize int64, fileAddress string) bool {
    conn := mydb.DBConn()
    stmt, err := conn.Prepare(
        "insert ignore into " +
            "file_meta(file_sha1, file_name, file_size, file_addr, status) " +
            "values(?, ?, ?, ?, 1)")
    if err != nil {
        fmt.Println("Failed to prepare statement, error: " + err.Error())
        return false
    }
    defer stmt.Close()
    
    result, err := stmt.Exec(fileHash, fileName, fileSize, fileAddress)
    if err != nil {
        fmt.Println("Failed to Execute sql, error: " + err.Error())
        return false
    }
    if affectedRows, err := result.RowsAffected(); err == nil {
        if affectedRows <= 0 {
            fmt.Printf("File with hash: %s has been uploaded before\n", fileHash)
        }
        return true
    }
    return false
}

type FileInformation struct {
    FileHash string
    FileName sql.NullString
    FileSize sql.NullInt64
    FileAddress sql.NullString
}

// 从数据库获取文件元信息
func GetFileMeta(fileHash string) (*FileInformation, error) {
    stmt, err := mydb.DBConn().Prepare(
        "SELECT file_sha1, file_addr, file_name, file_size FROM file_meta WHERE file_sha1=? AND status=1")
    if err != nil {
        fmt.Println(err.Error())
        return nil, err
    }
    defer stmt.Close()
    
    fileInformation := FileInformation{}
    err = stmt.QueryRow(fileHash).Scan(&fileInformation.FileHash, &fileInformation.FileAddress,
        &fileInformation.FileName, &fileInformation.FileSize)
    if err != nil {
        fmt.Println(err.Error())
        return nil, err
    }
    
    return &fileInformation, nil
}
