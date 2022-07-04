package main

import (
	"ginblog/model"
	"ginblog/routes"
	"os"
)

func main() {
	model.InitDb()
	routes.InitRouter()

	var symlinkPath string
	symlinkPath="./log"
	if _, err := os.Lstat(symlinkPath); err == nil {
		if err := os.Remove(symlinkPath); err != nil {
			//return fmt.Errorf("failed to unlink: %+v", err)
		}
	} else if os.IsNotExist(err) {
	//	return fmt.Errorf("failed to check symlink: %+v", err)
	}
}
