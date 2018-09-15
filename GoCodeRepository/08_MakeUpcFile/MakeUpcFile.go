// file:        MakeUpcFile.go
// author:      Haulf
// date:        2018.06.22
// brief:       Make upc file.

package main

import (
	"bufio"
	"crypto/md5"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type CDATA struct {
	Text string `xml:",cdata"`
}

type UpcFile struct {
	XMLName               xml.Name `xml:"update-package"`
	CreationDate          string   `xml:"creation-date"`
	DeviceName            string   `xml:"hw"`
	HardwareVersion       string   `xml:"hwv"`
	SourceSoftwareVersion string   `xml:"src_swv"`
	TargetSoftwareVersion string   `xml:"dst_swv"`
	Description           CDATA    `xml:"description"`
	Size                  string   `xml:"size"`
	Priority              string   `xml:"priority"`
	Md5                   string   `xml:"md5"`
	Content               string   `xml:"binary"`
}

func main() {
	upcFile := &UpcFile{}
	var diffFile string
	diffFile, _ = os.Getwd()

	fmt.Println("\n>>>>>>>>>>Make upc file begin\n")

	inputFile, inputError := os.Open("ota.conf")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" + "Have you got access to it?")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			break
		}

		resultSlice := strings.Split(inputString, "=")
		objectValue := strings.Trim(resultSlice[1], "\n")
		switch resultSlice[0] {
		case "hardware":
			upcFile.DeviceName = objectValue
		case "hwversion":
			upcFile.HardwareVersion = objectValue
		case "source_swv1":
			upcFile.SourceSoftwareVersion = objectValue
		case "dest_swv":
			upcFile.TargetSoftwareVersion = objectValue
		case "priority":
			upcFile.Priority = objectValue
		case "diff_file1":
			diffFile = objectValue
		case "description":
			upcFile.Description.Text = objectValue
			fmt.Println(upcFile.Description)
		default:
			fmt.Println(resultSlice[0] + " is noneed item.")
		}
	}

	upcFile.CreationDate = time.Now().Format("2006/01/02 15:04:05")

	if diffFile == "" {
		return
	}

	fmt.Println("Diff file is: " + diffFile)

	fileinfo, err := os.Stat(diffFile)
	if err != nil {
		panic(err)
	}

	upcFile.Size = strconv.FormatInt(fileinfo.Size(), 10)
	fmt.Println("The size of diff file is:" + upcFile.Size)

	file, err := os.Open(diffFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return
	}

	hexString := fmt.Sprintf("%x", h.Sum(nil))
	upcFile.Md5 = hexString

	buf, err := ioutil.ReadFile(diffFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read file error: %s\n", err)
	}
	encodeString := base64.StdEncoding.EncodeToString(buf)
	upcFile.Content = encodeString

	outputData := []byte(xml.Header)

	xmlData, err := xml.MarshalIndent(upcFile, "", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	outputData = append(outputData, xmlData...)

	err = ioutil.WriteFile("UPCFile.xml", outputData, 0644)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Make UPC file success.")
	}

	fmt.Println("\n>>>>>>>>>>Make upc file end\n")
}
