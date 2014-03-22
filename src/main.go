package main

import (
	"fmt"
	"net/url"
	"netRadio/config"
	"netRadio/http"
	"os"
	"time"
)

func main() {
	configFile, configErr := os.Open(".\\config.properties")
	if nil != configErr {
		fmt.Println("Failed to configure program -> ", configErr)
		fmt.Scanln()
		return
	}

	shangHaiZone, err := time.LoadLocation("Asia/Shanghai")
	if nil != err {
		fmt.Println("Failed to fetch system time : ", err)
		fmt.Scanln()
		return
	}
	currentTime := time.Now()
	shangHaiTime := currentTime.In(shangHaiZone)
	fmt.Println("ShangHai: ", shangHaiTime.Format("01-02-2006 15:04:05"))
	fmt.Println("Pick a date:")
	fmt.Println("-n : n days ago | 0 : today | n : n days in advance")

	var days int
	fmt.Scanln(&days)
	targetTime := shangHaiTime.AddDate(0, 0, days)

	weekday := targetTime.Weekday()

	choiceDate := targetTime.Format("01/02/2006")
	if time.Sunday == weekday || time.Saturday == weekday {
		fmt.Println(choiceDate+" is ", weekday, ", not a Feiyu day.")
		fmt.Scanln()
		return
	}

	date := targetTime.Format("20060102")

	properties := *(config.Properties(configFile))

	baseFolder := properties["saveLocation"]

	localMp3Path := baseFolder + "\\" + date + ".mp3"
	_, fileInfoErr := os.Stat(localMp3Path)
	if nil == fileInfoErr {
		fmt.Println(localMp3Path, " is exist!")
		fmt.Scanln()
		return
	} else {
		if !os.IsNotExist(fileInfoErr) {
			fmt.Println("Failed to create mp3 file for wrong file path -> ", fileInfoErr)
			fmt.Scanln()
			return
		}
	}

	baseUrl := properties["baseUrl"] + "/" + date[0:4] + "/ezm" + date[2:] + ".mp3"
	_, urlInfoErr := url.Parse(baseUrl)
	if nil != urlInfoErr {
		fmt.Println("Failed to fetch mp3 file for wrong http path -> ", urlInfoErr)
		fmt.Scanln()
		return
	}

	fmt.Println("Start to download FeiYu of " + choiceDate + "...")
	downloadErr := http.Download(baseUrl, localMp3Path)
	if nil != downloadErr {
		fmt.Println(downloadErr)
	} else {
		fmt.Println("Done!")
	}
	fmt.Scanln()
}
