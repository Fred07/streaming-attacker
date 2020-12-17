package main

import (
	"flag"
	"fmt"
	"net/url"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

var attackCmd string
var attackNum int
var missionTag string
var streamingServerURL string
var streamingServerPort int
var streamingServerAPP string
var mediaFile string

func init() {
	// Flags
	flag.StringVar(&attackCmd, "c", "ffmpeg", "ffmpeg 路徑")
	flag.IntVar(&attackNum, "n", 1, "模擬的 client 數量")
	flag.StringVar(&missionTag, "t", "A", "任務標籤，會標示在每個 stream 前面")
	flag.StringVar(&streamingServerURL, "s", "", "目標 URL")
	flag.IntVar(&streamingServerPort, "p", 1935, "Port")
	flag.StringVar(&streamingServerAPP, "a", "app", "Streaming app")
	flag.StringVar(&mediaFile, "f", "./audio/example_1.mp4", "欲串流的檔案路徑")
}

func main() {
	flag.Parse()

	// 確認指令存在，並取得完整路徑
	var err error
	attackCmd, err = exec.LookPath(attackCmd)
	if err != nil {
		fmt.Println("找不到指定的 c 參數指令")
		return
	}

	// Check required args
	if streamingServerURL == "" {
		fmt.Println("s 參數不得為空")
		return
	}

	// Display args
	displayInfo()

	var wg sync.WaitGroup
	for i := 1; i <= attackNum; i++ {
		var streamKey = genStreamKey(uint8(i))
		wg.Add(1)
		go attack(streamKey, &wg)
		time.Sleep(time.Second * 1)
	}

	wg.Wait()
}

func attack(streamKey string, wg *sync.WaitGroup) {
	defer wg.Done()
	// fmt.Println(getTargetURL(streamKey))
	cmd := exec.Command(attackCmd, "-re", "-i", mediaFile, "-c", "copy", "-f", "flv", getTargetURL(streamKey))
	fmt.Println("streaming: " + streamKey + " is running")
	cmd.Run()
}

func getTargetURL(streamKey string) string {
	// return fmt.Sprintf("rtmp://%s:%d/%s/%s", streamingServerURL, streamingServerPort, streamingServerAPP, streamKey)
	targetURL := url.URL{
		Scheme: "rtmp",
		Host:   fmt.Sprintf("%s:%d", streamingServerURL, streamingServerPort),
		Path:   fmt.Sprintf("%s/%s", streamingServerAPP, streamKey),
	}

	return targetURL.String()
}

func genStreamKey(initN uint8) string {
	return missionTag + strconv.Itoa(int(initN))
}

func displayInfo() {
	fmt.Println("Cmd path: " + attackCmd)
	fmt.Println("Runner Num: " + strconv.Itoa(attackNum))
	fmt.Println("Task tag: " + missionTag)
	fmt.Println("Target Server: " + streamingServerURL)
	fmt.Println("Target Port: " + strconv.Itoa(streamingServerPort))
	fmt.Println("Stream App: " + streamingServerAPP)
	fmt.Println("Target file: " + mediaFile)
	fmt.Println(" ")
}
