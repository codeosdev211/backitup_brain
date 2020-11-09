package filesys

import (
    "io/ioutil"
    "fmt"
    "encoding/base64"
    "os"
)

var basePath string = "/home/codeos/Desktop"

func WriteFile(path *string, data *string) error {
    bytes, _ := base64.StdEncoding.DecodeString(*data)
    return ioutil.WriteFile(*path, bytes, 0644)
}

func CreatePath(owner *string, name *string) string {
    return fmt.Sprintf("%v/%v/%v", basePath, *owner, *name)
}


func ReadFile(filePath *string) ([]byte, error) {
    if _, err := os.Stat(*filePath); os.IsNotExist(err) {
        return nil, err
    }
    return ioutil.ReadFile(*filePath)
}
func CreateDir(owner *string) error {
    dirPath := fmt.Sprintf("%v/%v", basePath, *owner)
    return os.Mkdir(dirPath, 0755)
}
