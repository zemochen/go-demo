package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// 递归遍历当前目录及其子目录
	err := processDirectory("D:\\Work\\bat_test\\git")
	if err != nil {
		fmt.Printf("遍历目录时出错: %v\n", err)
	}

	fmt.Println("所有Maven项目的target文件夹已清理完成!")
}

// 递归处理目录
func processDirectory(dir string) error {
	// 读取目录内容
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	// 遍历目录内容
	for _, entry := range entries {
		// 检查是否为Maven项目
		if entry.IsDir() && hasPomXml(filepath.Join(dir, entry.Name())) {
			fmt.Printf("找到Maven项目: %s\n", filepath.Join(dir, entry.Name()))
			// 清理Maven项目的target文件夹
			err := cleanMavenTarget(filepath.Join(dir, entry.Name()))
			if err != nil {
				return err
			}
		} else if entry.IsDir() {
			// 递归处理子目录
			err := processDirectory(filepath.Join(dir, entry.Name()))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// 检查目录是否包含pom.xml文件
func hasPomXml(dir string) bool {
	_, err := os.Stat(filepath.Join(dir, "pom.xml"))
	return err == nil
}

// 清理Maven项目的target文件夹
func cleanMavenTarget(dir string) error {
	const mvnPath = "C:\\Program Files\\JetBrains\\IntelliJ IDEA Community Edition 2023.3.4\\plugins\\maven\\lib\\maven3\\bin"
	cmd := exec.Command(mvnPath+" mvn", "clean")
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("清理Maven项目 %s 的target文件夹时出错: %w", dir, err)
	}
	fmt.Printf("已清理target文件夹: %s\n", dir)
	return nil
}
