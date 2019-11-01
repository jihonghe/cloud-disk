/**
 *    FILENAME      :   file_meta.go
 *    AUTHOR        :   jihonghe
 *    DATE          :   19-10-24
 *    DESCRIPTION   :
 */
package meta

import (
    mydb "cloudDisk/db"
    "sort"
)

type FileMeta struct {
    FileSha1 string
    FileName string
    FileSize int64
    Location string
    UploadTime string
}

var fileMetas map[string]FileMeta

func init() {
    fileMetas = make(map[string]FileMeta)
}

func UpdateFileMeta(fileMeta FileMeta) {
    fileMetas[fileMeta.FileSha1] = fileMeta
}

// 新增/更新文件元信息到mysql中
func UpdateFileMetaToDB(fileMeta FileMeta) bool {
    return mydb.OnFileUploadFinished(fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize, fileMeta.Location)
}

// 通过sha1值获取文件的元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
    return fileMetas[fileSha1]
}

// 从数据库获取文件元信息
func GetFileMetaFromDB(fileSha1 string) (FileMeta, error) {
    fileInformation, err := mydb.GetFileMeta(fileSha1)
    if err != nil {
        return FileMeta{}, err
    }
    fileMeta := FileMeta{
        FileSha1:   fileInformation.FileHash,
        FileName:   fileInformation.FileName.String,
        FileSize:   fileInformation.FileSize.Int64,
        Location:   fileInformation.FileAddress.String,
        UploadTime: "",
    }
    
    return fileMeta, nil
}

func GetLastFileMetas(limit int) []FileMeta {
    fileMetaArray := make([]FileMeta, len(fileMetas))
    for _, v := range fileMetas {
        fileMetaArray = append(fileMetaArray, v)
    }
    sort.Sort(ByUploadTime(fileMetaArray))

    return fileMetaArray[0:limit]
}

func RemoveFileMeta(fileSha1 string) {
    delete(fileMetas, fileSha1)
}
