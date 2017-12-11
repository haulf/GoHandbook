// author:     aihaofeng
// date:       2017.09.07
// version:    1.0
// go version: 1.9
// brief:      Install apks test.

package main

import (
	"fmt"
	"io/ioutil"
	//"log"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	//log.SetFlags(log.Flags() | log.Ldate | log.Lmicroseconds)
	//log.Println(">>>>>>>>>>begin")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Copyright (C) aihaofeng 2017")
	fmt.Println("")
	fmt.Println("")

	fmt.Println("开始安装应用程序....")
	fmt.Println("")
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Get current path failure.")
	}

	apkFiles := listApkFiles(currentPath + "/apks")
	for i := 0; i < len(apkFiles); i++ {
		fmt.Println("安装：" + apkFiles[i])
		execInstallCommand(apkFiles[i])
	}

	// log.Println(">>>>>>>>>>end")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("安装程序结束")
	fmt.Println("")
	fmt.Println("")
}

func listApkFiles(folder string) (apkFiles []string) {
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		if file.IsDir() {
			listApkFiles(folder + "/" + file.Name())
		} else {
			apkFileRegexp := regexp.MustCompile(".apk{1}$")
			isApkFile := apkFileRegexp.MatchString(file.Name())
			if isApkFile {
				apkFiles = append(apkFiles, folder+"/"+file.Name())
			}
		}
	}
	return
}

func execInstallCommand(apkFile string) {
	cmd := exec.Command("CMD", "/C", "adb install -r", apkFile)
	err := cmd.Run()
	if err != nil {
		fmt.Println("安装执行错误: ", err)
	}
}
