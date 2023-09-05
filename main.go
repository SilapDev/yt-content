package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	UNAVAILABLE = "Video unavailable"
)

func main() {

	startTime := time.Now()

	res := IsAgeRestricted("ieumolVsiWI")

	fmt.Println(res)

	fmt.Println(time.Since(startTime).String())

}

func IsAgeRestricted(videoId string) bool {

	res, err := SendRequestToUrl(videoId)

	if err != nil {
		log.Printf("error in SendRequestToUrl(): %s", err)
		return true
	}

	isRestricted := strings.Contains(res, UNAVAILABLE)

	return isRestricted

}

func SendRequestToUrl(videoId string) (string, error) {

	youtubeUrl := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoId)

	method := "GET"

	req, err := http.NewRequest(method, youtubeUrl, nil)

	if err != nil {
		log.Printf("error when create request: %s", err.Error())
		return "", err
	}

	req.Header.Add("authority", "www.youtube.com")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("cookie", "YSC=agopq3oR14c; VISITOR_PRIVACY_METADATA=CgJERRIA; CONSENT=PENDING+307; SOCS=CAESEwgDEgk1NjE1Mzc5MjgaAnJ1IAEaBgiA39mnBg; __Secure-YEC=CgtZM2hwN21RVVBRTSid3NunBjIGCgJERRIA; PREF=tz=Asia.Tashkent&f2=8000000")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Printf("error when send request %s", err.Error())
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Printf("error when read response body %s ", err)
		return "", err
	}

	return string(body), nil

}
