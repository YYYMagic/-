package main

import (
	"encoding/json"
	"fmt"
	"github.com/YYYMagic/term-music/format"
	ui "github.com/gizak/termui/v3"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
)

func GetMusicList(key string) (map[string]string, error) {
	u := url.URL{
		Scheme: "http",
		Host: "www.kuwo.cn",
		Path: "api/www/search/searchMusicBykeyWord",
	}
	q := u.Query()
	q.Set("key", key)
	q.Set("pn", "1")
	q.Set("rn", "10")
	q.Set("httpsStatus", "1")
	q.Set("reqId", "cc337fa0-e856-11ea-8e2d-ab61b365fb50")
	u.RawQuery = q.Encode()

	client := &http.Client {}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63")
	req.Header.Add("Cookie", "_ga=GA1.2.1083049585.1590317697; _gid=GA1.2.2053211683.1598526974; _gat=1; Hm_lvt_cdb524f42f0ce19b169a8071123a4797=1597491567,1598094297,1598096480,1598526974; Hm_lpvt_cdb524f42f0ce19b169a8071123a4797=1598526974; kw_token=HYZQI4KPK3P; kw_token=UOCXAP14GI")
	req.Header.Add("Referer", "http://www.kuwo.cn/search/list?key=%E5%91%A8%E6%9D%B0%E4%BC%A6")
	req.Header.Add("csrf", "HYZQI4KPK3P")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result :=  make(map[string]interface{})
	json.Unmarshal(body, &result)

	list := result["data"].(map[string]interface{})["list"].([]interface{})

	musicList := make(map[string]string)
	for _, v := range list {
		curMusic := v.(map[string]interface{})
		curMusicRID := curMusic["rid"]
		curMusicName := curMusic["name"].(string)
		curMusicRIDStr := fmt.Sprintf("%.0f", curMusicRID)
		curMusicUrl, _ := getMusicUrl(curMusicRIDStr)
		musicList[curMusicName] = curMusicUrl
	}

	return musicList, nil
}

func getMusicUrl(rid string) (string, error) {
	apiMusic := fmt.Sprintf("http://www.kuwo.cn/url?format=mp3&rid=%s&response=url&type=convert_url3&br=128kmp3&from=web&t=1598528574799&httpsStatus=1&reqId=72259df1-e85a-11ea-a367-b5a64c5660e5", rid)
	resp, err := http.Get(apiMusic)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var result map[string]string
	json.Unmarshal(body, &result)

	return result["url"], nil
}

func PrintMusicList(musicList map[string]string) {
	l := format.NewListCtr()
	l.Title = "歌曲列表"
	l.TitleStyle = ui.NewStyle(ui.ColorMagenta)
	l.TextStyle = ui.NewStyle(ui.ColorBlack)
	l.SetRect(0,0,140,12)

	for k, v := range musicList {
		l.Rows = append(l.Rows, k + " " + v)
		l.Meta = append(l.Meta, v)
	}

	ui.Render(l)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			l.ScrollDown()
			l.Next()
		case "k", "<Up>":
			l.ScrollUp()
			l.Pre()
		case "<C-d>":
			l.ScrollHalfPageDown()
		case "<C-u>":
			l.ScrollHalfPageUp()
		case "<C-f>":
			l.ScrollPageDown()
		case "<C-b>":
			l.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
			}
		case "<Home>":
			l.ScrollTop()
		case "G", "<End>":
			l.ScrollBottom()
		case "<Enter>":
			exec.Command(`open`, l.Get().(string)).Start()
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		ui.Render(l)
	}

}