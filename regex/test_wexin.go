package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	mm := make(map[string][]string)
	var que url.Values = url.Values(mm)
	que.Add("appid", "wxb3789cd0954891d9")
	que.Add("redirect_uri", "http://www.8844028.com")
	que.Add("response_type", "code")
	que.Add("scope", "snsapi_login")
	que.Add("state", "value")
	tmp1 := que.Encode()
	ra := "https://open.weixin.qq.com/connect/qrconnect?" + tmp1
	resp, err := http.Get(ra)
	if err != nil {
		fmt.Println("http get error.")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}
	src := string(body)
	//<div class="wrp_code"><img class="qrcode lightBorder" src="/connect/qrcode/031dNCgfe-789o2w" /></div>
	fmt.Println("ret:>>>>>>>>", src)

	re1, _ := regexp.Compile(`\<div class="wrp_code"\>\<img class="qrcode lightBorder" src="([[:graph:]]*)" /\>\</div\>`)

	tmp := re1.FindAllStringSubmatch(src, -1)
	fmt.Println(tmp)
	fmt.Println("============head:", *resp)

	address := "https://open.weixin.qq.com" + tmp[0][1]
	fmt.Println(address)
	resp1, _ := http.Get(address)
	defer resp1.Body.Close()
	body1, err1 := ioutil.ReadAll(resp1.Body)
	if err1 != nil {
		fmt.Println("err1:", err1)
		return
	}
	//fmt.Println(body1)
	sp := strings.Split(tmp[0][1], "/")
	uuid := sp[len(sp)-1]
	ioutil.WriteFile("tmp.jpg", body1, os.ModePerm)
	for {
		n := strconv.FormatUint(uint64(time.Now().UnixNano()), 10)
		append := "?uuid=" + uuid + "&_=" + n
		address := "https://long.open.weixin.qq.com/connect/l/qrconnect" + append
		fmt.Println("connect to===========:", address)
		resp, err := http.Get(address)
		if err != nil {
			continue
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("http read error")
			return
		}
		src := string(body)
		fmt.Println(">>>>", src)
	}
}
