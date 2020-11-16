//ALL FUNCTIONS
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/sqweek/dialog"
	"github.com/webview/webview"
	"gopkg.in/ini.v1"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("LOG Output: %s\n", err)
	}
}
func serverOlustur() {
	http.ListenAndServe(":5555", nil)
}

func getVideoInfo(w webview.WebView, URL string) {
	getJSON := fmt.Sprintf("youtube-dl -j %s", URL)
	cmd := exec.Command("bash", "-c", getJSON)
	output, err := cmd.CombinedOutput()

	if err == nil {
		jsonData := []byte(output)

		var videoData VideoInfo

		err = json.Unmarshal(jsonData, &videoData)
		if err != nil {
			log.Fatalf("error when unmarshaling json\n")
		}
		log.Printf("Looking for video: %s\n", URL)
		sendEval := fmt.Sprintf("showVideoInfo(`%s`,`%s`,`%s`,`%s`,`%s`,`%s`)", videoData.Title,
			videoData.Uploader, viewCountParser(videoData.Views), videoData.Thumnails[0].URL,
			videoData.Description, durationParser(videoData.Duration))
		w.Dispatch(func() {
			w.Eval(sendEval)
		})
		comboboxCode := "$('#videoFormat').html(`<option sounddata='null' value='null'>Choose Video Format</option>`)"
		downloadMP3Code := "$('#downloadMP3Button').attr('data',`" + videoData.Title + "" + URL + "`)"
		var audioFileURL string
		for _, v := range videoData.VideoFormats {
			if v.Resolution == "tiny" {
				audioFileURL = v.URL
				break
			}
		}
		setAudio := "$('#playAudioButton').attr('data','" + audioFileURL + "')"
		setVideo := "$('#watchVideoButton').attr('data','" + videoData.VideoID + "')"
		w.Dispatch(func() {
			w.Eval(setAudio)
			w.Eval(setVideo)
			w.Eval(comboboxCode)
			w.Eval(downloadMP3Code)
		})
		for _, v := range videoData.VideoFormats {
			if v.Resolution != "tiny" {
				res := fmt.Sprintf("%sx%s", strconv.Itoa(v.Width), strconv.Itoa(v.Height))
				addOption := fmt.Sprintf("$('#videoFormat').append(`<option value='%s%s%s%s%s'>%s - %s (%s)</option>`)",
					v.Acodec, v.ID, videoData.Title, v.Ext, v.Resolution, strings.ToUpper(v.Ext), res, fileSizeParser(v.FileSize))
				w.Dispatch(func() {
					w.Eval(addOption)
				})
			}
		}
	} else {
		w.Dispatch(func() {
			w.Eval("showWarning('Video not found')")
			w.Eval("$('#loading').fadeOut()")
		})
		log.Println("Video not found!")
	}
}
func downloadVideo(w webview.WebView, Data string, URL string) {
	splitData := strings.Split(Data, "")
	log.Printf("data lenght: %d", len(splitData))
	if len(splitData) != 5 {
		log.Fatalln("err: lenght of data is not 5")
	}
	soundData := splitData[0]
	ID := splitData[1]
	Title := splitData[2]
	Title = strings.Replace(Title, `"`, "", -1)
	Title = strings.Replace(Title, `/`, "", -1)
	Title = strings.Replace(Title, `'`, "", -1)
	log.Printf("incoming data: %s\n", Data)
	extension := splitData[3]
	downloadName := Title + "." + extension + " is downloading..."
	downloading(w, downloadName)
	resolution := splitData[4]
	log.Printf("Compound Sound: %s\n", soundData)
	file := fmt.Sprintf("%s/%s.%s", VDL, Title, extension)
	log.Printf("Checking file existing: %s\n", file)
	if soundData == "none" {

		//COMMAND-LINES
		downloadLocation := "cd " + VDL
		downloadSound := fmt.Sprintf("youtube-dl --no-overwrites --extract-audio --audio-format aac --output downtube.aac %s", URL)
		downloadVideo := fmt.Sprintf("youtube-dl --no-overwrites --output downtube.%s -f %s %s", extension, ID, URL)
		mergeProcess := `ffmpeg -i downtube.` + extension + ` -i downtube.aac -c:v copy -c:a copy "` + Title + `-` + resolution + `.` + extension + `" -y`
		rmOldFiles := "rm downtube.aac && rm downtube." + extension

		log.Println("###DOWNLOADING SOUND FILE (1/4)")
		command := fmt.Sprintf("%s && %s", downloadLocation, downloadSound)
		cmd := exec.Command("bash", "-c", command)
		_, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("error when downloading sound file")
		}
		log.Println("###DOWNLOADING VIDEO FILE (2/4)")
		command = fmt.Sprintf("%s && %s", downloadLocation, downloadVideo)
		cmd = exec.Command("bash", "-c", command)
		_, err = cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("error when downloading video file")
		}
		log.Println("###MERGING VIDEO AND SOUND FILES (3/4)")
		command = fmt.Sprintf("%s && %s", downloadLocation, mergeProcess)
		cmd = exec.Command("bash", "-c", command)
		_, err = cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("error when merging video and sound files")
		}
		log.Println("###REMOVING OLD VIDEO AND SOUND FILES (4/4)")
		command = fmt.Sprintf("%s && %s", downloadLocation, rmOldFiles)
		cmd = exec.Command("bash", "-c", command)
		_, err = cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("error when removing old files")
		}
		log.Println("###VIDEO DOWNLOADED SUCCESSFULLY!")
		downloaded := Title + "." + extension + " has downloaded"
		w.Dispatch(func() {
			w.Eval("$('#downloading').fadeOut()")
		})
		success(w, downloaded)

	} else {
		log.Println("###DOWNLOADING VIDEO (1/1)")
		command := `cd ` + VDL + ` && youtube-dl --no-overwrites --output ` + Title + `-` + resolution + `.` + extension + ` -f ` + ID + ` ` + URL
		cmd := exec.Command("bash", "-c", command)
		_, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("error when downloading video: %s", err)
		}
		log.Println("###VIDEO DOWNLOADED SUCCESSFULLY!")
		downloaded := Title + "." + extension + " has downloaded"
		w.Dispatch(func() {
			w.Eval("$('#downloading').fadeOut()")
		})
		success(w, downloaded)
	}
}
func downloadMP3(w webview.WebView, Data string) {
	splitData := strings.Split(Data, "")
	Title := splitData[0]
	URL := splitData[1]
	Title = strings.Replace(Title, `"`, "", -1)
	Title = strings.Replace(Title, `/`, "", -1)
	Title = strings.Replace(Title, `'`, "", -1)
	log.Println("###DOWNLOADING AUDIO")
	downloadName := Title + ".mp3 is downloading..."
	downloading(w, downloadName)
	downloadSound := `youtube-dl --no-overwrites --extract-audio --audio-format mp3 --output "` + Title + `.mp3" ` + URL
	downloadLocation := "cd " + MDL
	command := downloadLocation + " && " + downloadSound
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("error when downloading audio file: %v", output)
	}
	log.Println("###AUDIO DOWNLOADED SUCCESSFULLY!")
	downloaded := Title + ".mp3 has downloaded."
	w.Dispatch(func() {
		w.Eval("$('#downloading').fadeOut()")
	})
	success(w, downloaded)
}

func success(w webview.WebView, message string) {
	w.Dispatch(func() {
		w.Eval("showSuccess('" + message + "')")
	})
}

func warning(w webview.WebView, message string) {
	w.Dispatch(func() {
		w.Eval("showWarning('" + message + "')")
	})
}
func downloading(w webview.WebView, message string) {
	w.Dispatch(func() {
		w.Eval("downloading('" + message + "')")
	})
}
func setVideoLocation() string {
	directory, err := dialog.Directory().Title("Video Folder").Browse()
	checkErr(err)
	if directory != "" {
		fmt.Printf("Choosen folder for video: %s\n", directory)
		VDL = directory
		settings, err := ini.Load("settings.ini")
		checkErr(err)
		settings.Section("directories").Key("video").SetValue(directory)
		settings.SaveTo("settings.ini")
	} else {

	}

	return directory
}
func setMP3Location() string {
	directory, err := dialog.Directory().Title("MP3 Folder").Browse()
	checkErr(err)
	if directory != "" {
		fmt.Printf("Choosen folder for MP3: %s\n", directory)
		MDL = directory
		settings, err := ini.Load("settings.ini")
		checkErr(err)
		settings.Section("directories").Key("mp3").SetValue(directory)
		settings.SaveTo("settings.ini")
	} else {

	}
	return directory
}
func getSettings(w webview.WebView) {
	videoLocation := `$("#videoLocation").val("` + VDL + `")`
	w.Eval(videoLocation)
	mp3Location := `$("#mp3Location").val("` + MDL + `")`
	w.Eval(mp3Location)
}
