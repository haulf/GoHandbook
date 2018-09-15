// file:       UpdateJarAndDex.go
// author:     haulf
// date:       20180712
// brief:      Update jar and dex.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
	// "time"
)

func main() {
	fmt.Println("\n\nCopyright (C) aihaofeng 2017 2018\n\n")

	if len(os.Args) > 1 && os.Args[1] == "--help" {
		printUsage(os.Stdout)
		os.Exit(0)
	}

	flag.Usage = usage

	classRootPath := flag.String("r", "D:\\Android_Studio_Debug_Space\\workspace\\FrameworkDebug\\fwkdbg711\\build\\classes\\main", "the root path of class file")
	classPath := flag.String("c", "ClassPath.txt", "the path of ClassPath.txt")
	debugJarPath := flag.String("j", "debug_jar/service.jar", "the path of debug jar")
	outputTargetJar := flag.String("o", "service.jar", "the name of out target jar")

	flag.Parse()

	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Get current work directory error.")
	}

	tempJarName := "temp.jar"
	copyFile(currentPath+"\\tmp\\"+tempJarName, *debugJarPath)
	packClassToJar(*classRootPath, *classPath, currentPath+"\\tmp\\"+tempJarName)
	execDexOptimize(currentPath+"\\lib\\d8.jar", currentPath+"\\"+*outputTargetJar, currentPath+"\\tmp\\"+tempJarName)
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

	fmt.Println("更新开始，请稍侯...\n")

	var debugClasses []string
	for _, debugClassTotalName := range debugClassTotalNames {
		singlePathValues := strings.Split(debugClassTotalName, ".")
		relativeClassPath := strings.Join(singlePathValues[:len(singlePathValues)-1], "\\")
		basicClassName := singlePathValues[len(singlePathValues)-1]

		debugClasses = listClassFiles(classRootPath, relativeClassPath, basicClassName)
		// fmt.Println(time.Now().Format("2006-01-02 15:04:05.0000000"))

		classArgs := []string{"uf", tempJar}

		for _, debugClass := range debugClasses {
			fmt.Println(">> 更新  " + debugClass)
			classArgs = append(classArgs, "-C")
			classArgs = append(classArgs, classRootPath)
			classArgs = append(classArgs, debugClass)
		}

		execUpdateCommand(tempJar, classArgs)
		// fmt.Println(time.Now().Format("2006-01-02 15:04:05.0000000"))
	}
	fmt.Println("\n更新结束\n")
}

func listClassFiles(classRootPath, sourcePath, baseName string) (classFiles []string) {
	files, _ := ioutil.ReadDir(classRootPath + "\\" + sourcePath)
	for _, file := range files {
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
		fmt.Println("Open class file failure.")
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
			fmt.Println("Get debug class name error.")
		}
	}
	return
}

func execUpdateCommand(tempJar string, classArgs []string) {
	cmd := exec.Command("jar", classArgs...)
	err := cmd.Run()
	if err != nil {
		fmt.Println("\n更新class文件过程中出现了错误，请仔细检查!\n")
	}
}

func execDexOptimize(dexJarPath, outputTargetJar, tempJar string) {
	fmt.Println("优化开始，请稍侯...\n")
	cmd := exec.Command("CMD", "/K", "java", "-jar", dexJarPath, "--release", "--min-api", "26", "--output", outputTargetJar, tempJar)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Dex优化过程中出现错误。\n")
	}
	fmt.Println("优化结束\n")
}
