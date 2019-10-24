/**
 *    FILENAME      :   file_meta.go
 *    AUTHOR        :   jihonghe
 *    DATE          :   19-10-24
 *    DESCRIPTION   :
 */
package meta

import (
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

// 通过sha1值获取文件的元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
    return fileMetas[fileSha1]
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
