package utils

import (
    "os"
    "fmt"
)

func IsDir(fileName string) (bool, error) {
    fileStat, err := os.Stat(fileName)
    if err != nil {
        return false, err
    }
    return fileStat.IsDir(), nil
}

func ListDir(fileName string) ([]os.FileInfo, error) {
    dirFile, err := os.Open(fileName)
    if err != nil {
        return nil, err
    }
    return dirFile.Readdir(0)
}

func CreateDir(fileName string) error {
    fileStat, err := os.Stat(fileName)
    if err == nil {
        if !fileStat.IsDir() {
            return fmt.Errorf("file(%s) already exist, but not directory", fileName)
        }
        return nil
    } else {
        if !os.IsNotExist(err) {
            return fmt.Errorf("get file(%s) stat failed, errmsg(%v)", fileName, err)
        }
        return os.MkdirAll(fileName, 0755)
    }
    return nil
}
