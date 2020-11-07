package filesys

import (
    "io/ioutil"
    "fmt"
    "os"
)

var basePath string = "/home/codeos/Desktop"

func WriteFile(path *string, data *string) error {
    return ioutil.WriteFile(*path, []byte(*data), 0644)
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
