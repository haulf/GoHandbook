// file:        InstallApks.go
// author:      Haulf
// date:        2017.09.18
// brief:       Install apks test.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
)

var debugLog *log.Logger

func main() {
	logFileName := "install_apk_log_file"
	logFile, err := os.Create(logFileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("Open file error !")
	}

	debugLog = log.New(logFile, "[ai]", log.Flags()|log.Ldate|log.Lmicroseconds|log.Lshortfile)
	debugLog.Println(">>>>>>>>>>begin")

	fmt.Println("\n\nCopyright (C) aihaofeng 2017\n\n")

	fmt.Println(">> 安装APP开始\n")

	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Get current path failure.")
	}

	apkFiles := listApkFiles(currentPath + "/apks")
	for i := 0; i < len(apkFiles); i++ {
		fmt.Println("开始安装：" + apkFiles[i])
		execInstallCommand(apkFiles[i])
	}

	fmt.Println("\n\n>> 安装APP结束\n\n")

	debugLog.Println(">>>>>>>>>>end")
}

func listApkFiles(apkPath string) (apkFiles []string) {
	files, _ := ioutil.ReadDir(apkPath)
	for _, file := range files {
		if file.IsDir() {
			listApkFiles(apkPath + "/" + file.Name())
		} else {
			apkFileRegexp := regexp.MustCompile(".apk{1}$")
			isApkFile := apkFileRegexp.MatchString(file.Name())
			if isApkFile {
				apkFiles = append(apkFiles, apkPath+"/"+file.Name())
			}
		}
	}
	return
}

func execInstallCommand(apkFile string) {
	executeCommand := exec.Command("CMD", "/C", "adb install -r", apkFile)
	err := executeCommand.Run()
	if err != nil {
		fmt.Println("\n安装过程中出现了错误，请检查要安装的APK的名称中是否有空格或有中文!\n")
		debugLog.Println("安装执行错误，err信息为：", err)
	}
}
