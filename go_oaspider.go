package main

import (
	"fmt"
	"time"
	"net/http"
	"io"
	"strconv"
	"strings"
	"golang.org/x/text/encoding"
    "bufio"
	"golang.org/x/net/html/charset"
	"github.com/PuerkitoBio/goquery"
)
//NORMALTIMEFORMAT 时间格式
const (
	NORMALTIMEFORMAT = "2006-01-02 15:04:05"
	HOME = ""
	LOGINURL = ""
	URL = ""
	LIION = ""
	USERINFO = ""
)

/*
oa 爬虫
*/
func oaspider()  {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	myCookie := getCookie()
	login(myCookie)
	req.Header.Set("Cookie", myCookie)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("获取信息失败")
		return
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("HTML解析失败")
		return
	}
	t := weekDayMonday()
	var hours float32 = 40
	doc.Find("tr.TableData").Each(func(i int, s *goquery.Selection) {
		var s1 string
		var s2 string
		s.Find("td").Each(func(i int, s *goquery.Selection){
			if i == 1 {
				s1 = s.Text()
			}
			if i == 2 {
				s2 = s.Text()
			}
		})
		t2, err := time.ParseInLocation(NORMALTIMEFORMAT, s2, time.Local)
		if err != nil {
			t2 = time.Now()
			s2 = t2.Format(NORMALTIMEFORMAT)
		}
		if t.Before(t2){
			hour := timeComp(s1, s2)
			hours = hours - hour
			fmt.Printf("time: %s ==> %s, hour:%6.2f\n", s1, s2, hour)
		}
	})
	hourss := timeTostring(hours)
	fmt.Printf("本周还需工作%s时间\n", hourss)
}

func login(cookie string)  {

	client := &http.Client{Timeout: 15 * time.Second}
	payload := strings.NewReader(USERINFO)
	req, _ := http.NewRequest("POST", LOGINURL, payload)
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	resp, err := client.Do(req)
	if err != nil {
        panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("登录成功")
		// e := determineEncoding(resp.Body)
		// reader := transform.NewReader(resp.Body, e.NewDecoder())
		// bodyBytes, err := ioutil.ReadAll(reader)
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Printf("%s\n", bodyBytes)
	}
}

/*
获取Cookie
*/
func getCookie() string {
	resp,err := http.Get(HOME)
	if err != nil {
		panic(err)
	}
	s := resp.Header.Get("Set-Cookie")
	ss := strings.Split(s, ";")
	cookie := LIION + ss[0]
	fmt.Println(cookie)
	return cookie
}
/*
* 将时间float 转化为 string（xx小时xx分钟） 
*/
func timeTostring(f float32) string {
	//将float32转化为string
	s := strconv.FormatFloat(float64(f), 'f', 1, 64)
	//进行字符串操作，切割字符串
	ss := strings.Split(s, ".")
	m,err := strconv.ParseFloat(ss[1], 1)
	var t float64 = 6
	var mi string
	if err == nil {
		mi = strconv.FormatFloat(m * t, 'f', 0, 64)
	}
	return ss[0] + "小时" + mi + "分钟"
}
/**
* 计算s2 - s1的时间差
*/
func timeComp(s1, s2 string) float32 {
	var hour float32
	t1, err := time.ParseInLocation(NORMALTIMEFORMAT, s1, time.Local)
	t2, err := time.ParseInLocation(NORMALTIMEFORMAT, s2, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() 
		hour = float32(diff) / 3600 - 1
	}
	return hour
}
/*
* weekDayMonday 获取本周周一的时间
*/
func weekDayMonday() time.Time {
    now := time.Now()
    offset := int(time.Monday - now.Weekday())
    if offset > 0 {
        offset = -6
    }
    weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	
	return weekStart
}

/*
* 中文乱码处理
*/
func determineEncoding(r io.Reader) encoding.Encoding {
    bytes, err := bufio.NewReader(r).Peek(1024)		
    if err != nil {
        panic(err)
    }
    e, _, _ := charset.DetermineEncoding(bytes, "")
    return e
}