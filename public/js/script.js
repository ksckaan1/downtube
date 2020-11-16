//import * as $ from './jquery';
//VARIABLES
let isAudioPlaying = false
let isVideoPlaying = false
$("#videoUrlInput").focus()
$("#videoUrlButton").click(function () {
    let isVideo = $("#videoUrlInput").val().includes("https://www.youtube.com/watch?v=")
    if (isVideo) {
        $("#loading").fadeIn()
        setTimeout(() => {
            let url = $("#videoUrlInput").val()
            getVideoInfo(url)
            $("#downloadButton").attr("data", url)
        }, 500);
    }else {
        showWarning("Wrong URL!")
    }

})
$("#videoUrlInput").on('input', function () {
    if ($("#videoUrlInput").val() == "") {
        $("#clear").fadeOut()
    } else {
        $("#clear").fadeIn()
    }
})

$("#clear").click(function () {
    $("#videoUrlInput").val("")
    $("#clear").fadeOut()
    $("#videoUrlInput").focus()
})

function showVideoInfo(videoName, uploader, views, thumbnail, videoDescription,duration) {
    $("#loading").fadeOut()
    $("#videoShow").fadeIn()
    $("#videoName").text(videoName);
    $("#uploader").text(uploader);
    $("#views").text(views);
    $("#videoThumbnail").attr("src", thumbnail)
    $("#videoDescription").text(videoDescription)
    $("#duration").text(duration)
}
$("#closeVideoShow").click(function () {
    isVideoPlaying=false
    isAudioPlaying=false
    window.open($("#watchVideoButton").attr("data"), "videoPlayer");
    $("#videoShow").fadeOut()
    $("#videoDescription").fadeIn()
    $("#audioPlayZone").fadeOut()
    $("#audioPlayer").attr("src", "")
})

function showSuccess(text) {
    $("#success").text(text)
    $("#success").fadeIn()

    setTimeout(() => {
        $("#success").fadeOut()
    }, 3000);
}
function showWarning(text) {
    $("#warning").text(text)
    $("#warning").fadeIn()

    setTimeout(() => {
        $("#warning").fadeOut()
    }, 3000);
}
function downloading(text) {
    $("#downloading").text(text)
    $("#downloading").fadeIn()
}
$("#downloadButton").click(function () {
    if ($("#videoFormat").val() == "null") {
        showWarning("Please Select Video Format!")
    } else {
        console.log($("#downloadButton").attr("data"))
        downloadVideo($("#videoFormat").val(), $("#downloadButton").attr("data"))
        
    }

})
$("#downloadMP3Button").click(function () {
    downloadMP3($("#downloadMP3Button").attr("data"))
})
$("#chooseVideoLocation").click(function () {
    setVideoLocation().then(function (videoLocation) {
        if (videoLocation != "") {
            $("#videoLocation").val(videoLocation)
            showSuccess("The video download location has been set.")
        } else {
            showWarning("The video download location has not changed.")
        }

    })
})
$("#chooseMP3Location").click(function () {
    setMP3Location().then(function (MP3Location) {
        if (MP3Location != "") {
            $("#mp3Location").val(MP3Location)
            showSuccess("The MP3 download location has been set.")
        } else {
            showWarning("The MP3 download location has not changed.")
        }

    })
})


$("#playAudioButton").click(function () {
    
    if (isAudioPlaying != true) {
        isAudioPlaying = true
        $("#audioPlayZone").fadeIn()
        $("#audioPlayer").attr("src", $("#playAudioButton").attr("data"))
    } else {
        isAudioPlaying = false
        $("#audioPlayZone").fadeOut()
        $("#audioPlayer").attr("src", "")
    }
})

$("#watchVideoButton").click(function () {
    if (isAudioPlaying != true) {
        isVideoPlaying = true
        let videoSrc = "http://www.youtube.com/embed/" + $("#watchVideoButton").attr("data")
        window.open(videoSrc, "videoPlayer");
        $("#videoPlayer").fadeIn()
        $("#videoDescription").fadeOut()
    } else {
        isVideoPlaying = false
        window.open($("#watchVideoButton").attr("data"), "videoPlayer");
        $("#videoPlayer").fadeOut()
        $("#videoDescription").fadeIn()
    }
})
