package main

import (
	"log"
	"net/http"
	"os/user"

	_ "./statik" //Oluşturulmuş statik.go dosyasının konumu
	"github.com/rakyll/statik/fs"
	"github.com/webview/webview"
	"gopkg.in/ini.v1"
)

func init() {

	//Onload settings.ini
	settings, err := ini.Load("settings.ini")
	if err != nil {
		log.Fatalln("error when reading ini file")
	}
	tempVDL := settings.Section("directories").Key("video").String()
	tempMDL := settings.Section("directories").Key("mp3").String()

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if tempVDL == "default" {
		VDL = "/home/" + user.Username + "/Videos"
	} else {
		VDL = tempVDL
	}
	if tempMDL == "default" {
		MDL = "/home/" + user.Username + "/Music"
	} else {
		MDL = tempMDL
	}
}

func main() {
	statikFS, _ := fs.New()
	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	go serverOlustur()
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("DownTube")
	w.SetSize(800, 600, webview.HintFixed)
	w.Navigate("http://localhost:5555/")
	err := w.Bind("getVideoInfo", func(URL string) {
		go getVideoInfo(w, URL)
	})
	checkErr(err)
	err = w.Bind("downloadVideo", func(Data string, URL string) {
		go downloadVideo(w, Data, URL)
	})
	checkErr(err)
	err = w.Bind("downloadMP3", func(Data string) {
		go downloadMP3(w, Data)

	})
	checkErr(err)
	err = w.Bind("setVideoLocation", func() string {
		return setVideoLocation()
	})
	checkErr(err)
	err = w.Bind("setMP3Location", func() string {
		return setMP3Location()
	})
	checkErr(err)
	err = w.Bind("getSettings", func() {
		getSettings(w)
	})
	checkErr(err)
	w.Run()
}
