package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// 递归遍历当前目录及其子目录
	err := filepath.Walk(".", func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// 检查是否为Maven项目
		if info.IsDir() && hasPomXml(path) {
			fmt.Printf("找到Maven项目: %s\n", path)
			// 清理Maven项目的target文件夹
			cleanMavenTarget(path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("遍历目录时出错: %v\n", err)
	}

	fmt.Println("所有Maven项目的target文件夹已清理完成!")
}

// 检查目录是否包含pom.xml文件
func hasPomXml(dir string) bool {
	_, err := os.Stat(filepath.Join(dir, "pom.xml"))
	return err == nil
}

// 清理Maven项目的target文件夹
func cleanMavenTarget(dir string) {
	cmd := exec.Command("mvn", "clean")
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		fmt.Printf("清理Maven项目 %s 的target文件夹时出错: %v\n", dir, err)
	} else {
		fmt.Printf("已清理target文件夹: %s\n", dir)
	}
}
