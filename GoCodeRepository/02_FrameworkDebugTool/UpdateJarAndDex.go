// file:       UpdateJarAndDex
// author:     haulf
// date:       2017.09.18
// brief:      Update jar and dex.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var debugLog *log.Logger

func main() {
	logFileName := "HaulfDebugLogFile"
	logFile, err := os.Create(logFileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("Open file error !")
	}

	debugLog = log.New(logFile, "[ahf]", log.Flags()|log.Ldate|log.Lmicroseconds|log.Lshortfile)
	debugLog.Println(">>>>>>>>>>begin\n")

	fmt.Println("\n\nCopyright (C) aihaofeng 2017\n\n")

	if len(os.Args) > 1 && os.Args[1] == "--help" {
		printUsage(os.Stdout)
		os.Exit(0)
	}

	flag.Usage = usage

	classRootPath := flag.String("r", "D:\\Android_Studio_Debug_Space\\workspace\\FrameworkDebug\\fwkdbg711\\build\\classes\\main", "the root path of class file")
	classPath := flag.String("c", "class_path.txt", "the path of class_path.txt")
	debugJarPath := flag.String("j", "debug_jar/service.jar", "the path of debug jar")
	outputTargetJar := flag.String("o", "service.jar", "the name of out target jar")

	flag.Parse()

	debugLog.Println("classRootPath is:", *classRootPath)
	debugLog.Println("classPath is:", *classPath)
	debugLog.Println("debugJarPath is:", *debugJarPath)
	debugLog.Println("outputTargetJar is:", *outputTargetJar)
	debugLog.Println("")

	currentPath, err := os.Getwd()
	if err != nil {
		debugLog.Println("Get current work directory error.")
	}

	tempJarName := "temp.jar"
	copyFile(currentPath+"\\tmp\\"+tempJarName, *debugJarPath)
	packClassToJar(*classRootPath, *classPath, currentPath+"\\tmp\\"+tempJarName)
	execDexOptimize(currentPath+"\\lib\\dx.jar", currentPath+"\\"+*outputTargetJar, currentPath+"\\tmp\\"+tempJarName)

	debugLog.Println(">>>>>>>>>>end")
}

func printUsage(w *os.File) {
	fmt.Fprintf(w, "usage:\n    update_jar_and_dex  \n    -r [CLASS_ROOT_PATH]  \n    -c [CLASS_PATH]  \n    -j [TOOL_PATH/debug_jar/TAREGET_JAR]  \n    -o [TAREGET_JAR]")
}

func usage() {
	printUsage(os.Stderr)
	os.Exit(2)
}

func packClassToJar(classRootPath, classPath, tempJar string) {
	debugClassTotalNames := getDebugClassName(classPath)
	fmt.Println("更新开始\n")
	var debugClasses []string
	for _, debugClassTotalName := range debugClassTotalNames {
		debugLog.Println("#packClassToJar#DebugClassTotalName:", debugClassTotalName)
		singlePathValues := strings.Split(debugClassTotalName, ".")
		relativeClassPath := strings.Join(singlePathValues[:len(singlePathValues)-1], "\\")
		basicClassName := singlePathValues[len(singlePathValues)-1]

		debugClasses = listClassFiles(classRootPath, relativeClassPath, basicClassName)
		for _, debugClass := range debugClasses {
			debugLog.Println(">>> #packClassToJar#debugClass is:", debugClass)
			execUpdateCommand(classRootPath, debugClass, tempJar)
		}
	}
	fmt.Println("\n更新结束\n")
}

func listClassFiles(classRootPath, sourcePath, baseName string) (classFiles []string) {
	debugLog.Println("#listClassFiles#classRootPath is:", classRootPath)
	debugLog.Println("#listClassFiles#sourcePath is:", sourcePath)

	files, _ := ioutil.ReadDir(classRootPath + "\\" + sourcePath)
	for _, file := range files {
		debugLog.Println("#listClassFiles#class file is:", file.Name())
		if file.IsDir() {
			listClassFiles(classRootPath, sourcePath+"\\"+file.Name(), baseName)
		} else {
			classFileRegexp := regexp.MustCompile(baseName + "*")
			isClassFile := classFileRegexp.MatchString(file.Name())
			if isClassFile {
				classFiles = append(classFiles, sourcePath+"\\"+file.Name())
			}
		}
	}
	return
}

func copyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func getDebugClassName(fileName string) (debugClassNames []string) {
	f, err := os.Open(fileName)
	if err != nil {
		debugLog.Println("getDebugClassName#open file failure.")
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		debugClassNames = append(debugClassNames, line)
		if err != nil {
			if err == io.EOF {
				return
			}
			debugLog.Println("getDebugClassName#read string err", err)
		}
	}
	return
}

func execUpdateCommand(classRootPath, classFile, tempJar string) {
	debugLog.Println("#execUpdateCommand#classRootPath is:", classRootPath)
	debugLog.Println("#execUpdateCommand#classFile is:", classFile)
	debugLog.Println("#execUpdateCommand#tempJar is:", tempJar)
	fmt.Println(">> 更新 ", classFile)
	cmd := exec.Command("CMD", "/K", "jar -uvf", tempJar, "-C", classRootPath, classFile)

	debugLog.Println(">>>>>", cmd.Args)
	err := cmd.Run()
	if err != nil {
		fmt.Println("\n更新class文件过程中出现了错误，请仔细检查!\n")
		debugLog.Println("更新class文件过程中出现了错误，err信息为：", err)
	}
}

func execDexOptimize(dexJarPath, outputTargetJar, tempJar string) {
	cmd := exec.Command("CMD", "/C", "java -Xmx2048M -jar ", dexJarPath, "--dex", "--multi-dex", "--output="+outputTargetJar, tempJar)
	debugLog.Println("Dex优化命令：", cmd.Args)
	err := cmd.Run()
	if err != nil {
		debugLog.Println("Dex优化过程中出现了错误，err信息为：", err)
	}
}
