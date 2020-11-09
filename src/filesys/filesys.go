package filesys

import (
    "io/ioutil"
    "fmt"
    "encoding/base64"
    "os"
)

var basePath string = "/home/codeos/Documents"

func WriteFile(path *string, data *string) error {
    bytes, _ := base64.StdEncoding.DecodeString(*data)
    return ioutil.WriteFile(*path, bytes, 0644)
}

func CreatePath(owner *string, name *string) string {
    return fmt.Sprintf("%v/%v/%v", basePath, *owner, *name)
}


func ReadFile(filePath *string) (string, error) {
    if _, err := os.Stat(*filePath); os.IsNotExist(err) {
        return "", err
    }
    file, err := ioutil.ReadFile(*filePath)
    if err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(file), nil
}

func CreateDir(owner *string) error {
    dirPath := fmt.Sprintf("%v/%v", basePath, *owner)
    return os.Mkdir(dirPath, 0755)
}
