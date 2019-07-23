package main

import (
	"./structs"
	"./utils"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"model"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const Path = "http://pilipali.cc"

var wg = sync.WaitGroup{}

func main() {
	utils.StartPool()
	//getChapter("https://anime1.me/?cat=333")
	start := 1
	end := 2

	page := make(chan int, 20)
	for i := start; i <= end; i++ {
		wg.Add(1)
		go func(index int) {
			page <- 1
			url := Path + "/vod/show/id/4/page/" + strconv.Itoa(index) + ".html"
			getMenu(url)
			<-page
			wg.Done()
		}(i)
	}

	wg.Wait()
	close(page)
	fmt.Println("chan close")
	//getAllIndex()
	//utils.SaveOrUpdateIndex("onion", "1-13")
}

func testJSOn() {
	s := `[{"type": "123","file":"213123",label:"435345","default":"56456"}]`
	s = strings.Replace(s, ",label:", `,"label":`, -1)
	var arr []structs.UrlData
	_ = json.Unmarshal([]byte(s), &arr)
	log.Printf("Unmarshaled: %+v\n", arr)
	println(s)
}

func getMenu(url string) {
	c := colly.NewCollector()
	println("获取所有目录")
	c.OnHTML(".v_con_box ul li", func(e *colly.HTMLElement) {
		href, _ := e.DOM.Find(".v-txt .v-tit a").Attr("href")
		getInfo(Path + href)
		//name := e.DOM.Find(".v-txt .v-tit a").Text()
		//chapter := e.DOM.Find(".column-2").Text()
		//utils.SaveIndex(name, chapter, href, e.Index)
		//index := utils.SaveOrUpdateIndex(name, chapter)
		//getChapter(url + href, index.Id)
	})

	//c.OnRequest(func(r *colly.Request) {
	//	r.Headers.Set("cookie", utils.GetCookie())
	//	fmt.Println("Visiting", r.URL)
	//})

	c.Visit(url)
}

//提取数据函数
func extractHandle(rs, regStr string, num int) (content []string) {
	reg := regexp.MustCompile(regStr)
	allUrl := reg.FindAllStringSubmatch(rs, num)
	for _, item := range allUrl {
		//content=item[1]
		content = append(content, item[1])
	}
	return
}
func getInfo(url string) {
	c := colly.NewCollector()

	c.OnHTML(".icon", func(e *colly.HTMLElement) {
		name := e.DOM.Find(".txt_intro_con .tit h1").Text()
		chapter := e.DOM.Find(".txt_intro_con .p_txt .em_num").Text()
		other := e.DOM.Find(".txt_intro_con ul li a")
		img, _ := e.DOM.Find(".poster_placeholder img").Attr("src")
		introduction := e.DOM.Find(".infor_intro").Text()

		area := ""
		year := ""
		director := ""
		star := []string{}
		for i := 0; i < other.Length(); i++ {
			if i == 0 {
				year = other.Eq(i).Text()
				continue
			}
			if i == 1 {
				area = other.Eq(i).Text()
				continue
			}
			if i == other.Length()-1 {
				director = other.Eq(i).Text()
				continue
			}
			star = append(star, other.Eq(i).Text())
		}

		str := extractHandle(img, `/([0-9a-z]+\.[a-z]+)`, 1)
		nowTime := int(time.Now().Unix())
		timestr := strconv.Itoa(nowTime)
		path := "./public/upload/anime"
		imgPath := path + "/" + timestr + ".jpg"
		imgUrl := "/upload/anime/" + timestr + ".jpg"
		if len(str) > 0 {
			imgPath = path + "/" + str[0]
			imgUrl = "/upload/anime/" + str[0]
		}
		utils.SaveImg(img, imgPath, path, str)
		anime := utils.SaveOrUpdateAnime(name, chapter, year, area, introduction, imgUrl)
		utils.SaveDirector(director, anime.Id)

		for i := 0; i < len(star); i++ {
			utils.SaveStar(star[i], anime.Id)
		}

		fmt.Println("name:", name)
		fmt.Println("chapter:", chapter)
		fmt.Println("other:", other)
		fmt.Println("img:", img)
		s := e.DOM.Find(".v_con_box ul li a")

		//d := e.DOM.Find(".entry-title a[href]")
		for i := 0; i < s.Length(); i++ {
			href, _ := s.Eq(i).Attr("href")
			name := s.Eq(i).Text()
			getChapterUrl(Path+href, name, anime.Id)
		}
	})

	c.Visit(url)
}
func getChapterUrl(url, name, pid string) {
	c := colly.NewCollector()

	c.OnHTML(".iplays", func(e *colly.HTMLElement) {
		playStr := e.DOM.Find("script")

		for i := 0; i < playStr.Length(); i++ {
			if i == 0 {
				Str := playStr.Eq(i).Text()
				reg := regexp.MustCompile("{(.*)}")
				dramaStr := reg.FindAllStringSubmatch(Str, 1)

				if len(dramaStr) > 0 {
					var b []byte = []byte("{" + dramaStr[0][1] + "}")
					var data model.DramaPlay
					err := json.Unmarshal(b, &data)

					if err != nil {
						fmt.Println("err:", err)
						continue
					}
					chapter := utils.SaveChapter(name, pid, data.Url, data.From)
					fmt.Println("chapter:", chapter)
				}

			}
		}
	})

	c.Visit(url)
}

//func getAllIndex() {
//	index := utils.GetAllIndex()
//	println("获取详情...")
//	for i := 0; i < len(index); i++ {
//		data := index[i]
//		getChapter(path+data.Url, data.Id)
//	}
//}
//
//func getChapter(url string, pid string) {
//	c := colly.NewCollector()
//	// Find and visit all links
//	c.OnHTML("main", func(e *colly.HTMLElement) {
//		s := e.DOM.Find("iframe[src]")
//		d := e.DOM.Find(".entry-title a[href]")
//		for i := 0; i < s.Length(); i++ {
//			src, _ := s.Eq(i).Attr("src")
//			name := d.Eq(i).Text()
//			getChapterUrl(src, name, pid, i)
//		}
//	})
//
//	c.OnHTML(".nav-previous a[href]", func(e *colly.HTMLElement) {
//		e.Request.Visit(e.Attr("href"))
//	})
//
//	c.OnRequest(func(r *colly.Request) {
//		r.Headers.Set("cookie", utils.GetCookie())
//		fmt.Println("Visiting", r.URL)
//	})
//	c.Visit(url)
//}
//
//func getChapterUrl(url string, name string, pid string, num int) {
//	c := colly.NewCollector(colly.Async(true))
//	// Find and visit all links
//	c.OnHTML("body script", func(e *colly.HTMLElement) {
//		data := e.Text
//		start := strings.Index(data, "sources:")
//		end := strings.Index(data, ",controls:true")
//		if start > 0 && end > 0 {
//			s := data[start+8 : end]
//			s = strings.Replace(s, ",label:", `,"label":`, -1)
//			var arr []structs.UrlData
//			_ = json.Unmarshal([]byte(s), &arr)
//			var flag = false
//			for i := 0; i < len(arr); i++ {
//				if arr[i].Default == "true" {
//					utils.SaveChapter(name, pid, arr[i].File, num)
//					flag = true
//				}
//			}
//			if !flag {
//				some := 0
//				file := ""
//				for i := 0; i < len(arr); i++ {
//					s := arr[i].Label
//					if len(s) > 0 {
//						hd, _ := strconv.Atoi(s[0 : len(s)-1])
//						if hd > some {
//							file = arr[i].File
//						}
//						some = hd
//					}
//				}
//				if len(file) > 0 {
//					utils.SaveChapter(name, pid, file, num)
//				}
//			}
//		} else {
//			start := strings.Index(data, `,file:"`)
//			end := strings.Index(data, `",controls:true`)
//			if start > 0 && end > 0 {
//				file := data[start+7 : end]
//				utils.SaveChapter(name, pid, file, num)
//			}
//		}
//	})
//
//	c.OnHTML("video source", func(e *colly.HTMLElement) {
//		file := e.Attr("src")
//		if len(file) > 0 {
//			utils.SaveChapter(name, pid, file, num)
//		}
//	})
//
//	c.OnRequest(func(r *colly.Request) {
//		r.Headers.Set("cookie", utils.GetCookie())
//		fmt.Println("Visiting", r.URL)
//	})
//	c.Visit(url)
//}
