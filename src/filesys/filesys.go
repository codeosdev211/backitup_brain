package filesys

import (
    "io/ioutil"
    "fmt"
    "os"
)

var basePath string = "/home/codeos/Desktop"

func WriteFile(owner *string, name *string, data *string) error {
    filePath := getPath(&owner, &name)
    fileData := []byte(*data)
    return ioutil.WriteFile(filePath, fileData, 0644)
}

func getPath(owner **string, name **string) string {
    return fmt.Sprintf("%v/%v/%v",basePath, **owner, **name)
}

func ReadFile(filePath *string) ([]byte, error) {
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        return nil, err
    }
    return ioutil.ReadFile(*filePath)
}

func CreateOwnerDir(owner *string) error {
    dirPath := fmt.Sprintf("%v/%v", basePath, *owner)
    return os.Mkdir(dirPath, 0755)
}
