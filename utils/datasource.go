package utils

import (
	"../structs"
	"fmt"
	"time"
)

//func GetCookie() string {
//	var cookie []structs.Cookies
//	engine := GetCon()
//	engine.Where("id = 1").Find(&cookie)
//	if len(cookie) > 0 {
//		return cookie[0].Value
//	}
//	return ""
//}
//
//func SaveIndex(name string, chapter string, url string, order int) structs.Index {
//	var index structs.Index
//	engine := GetCon()
//	index.Url = url
//	index.Id = NewKeyId()
//	index.Update = time.Now()
//	index.Total = 0
//	index.Name = name
//	index.Chapter = chapter
//	index.Index = order
//	engine.Insert(&index)
//	fmt.Printf("%+v\n", index)
//	return index
//}
//
//func GetAllIndex() []structs.Anime {
//	var animes []structs.Anime
//	engine := GetCon()
//	engine.Where("flag = 0").OrderBy("`anime` asc").Find(&animes)
//	return animes
//}
//
func SaveOrUpdateAnime(name string, chapter string, year, area, introduction, img string) structs.Anime {
	var Anime structs.Anime
	var Animes []structs.Anime
	engine := GetCon()
	engine.Where("name = ?", name).Find(&Animes)
	if len(Animes) == 0 {
		Anime.Id = NewKeyId()
		Anime.Update = time.Now()
		Anime.Total = 0
		Anime.Name = name
		Anime.Chapter = chapter
		Anime.Year = year
		Anime.Area = area
		Anime.Picture = img
		Anime.Form = "pilipili"
		Anime.Introduction = introduction
		engine.Insert(&Anime)
		Anime.Flag = true
	} else {
		Anime = Animes[0]
		Anime.Chapter = chapter
		Anime.Update = time.Now()
		engine.Update(&Anime)
		Anime.Flag = false
	}
	fmt.Printf("%+v\n", Anime)
	return Anime
}
func SaveDirector(name string, pid string) {
	var director structs.Director
	var directors []structs.Director
	engine := GetCon()
	engine.Where("name = ?", name).Find(&directors)
	if len(directors) == 0 {
		director.Id = NewKeyId()
		director.Pid = pid
		director.Name = name
		director.CreateTime = time.Now()
		engine.Insert(&director)
	} else {
		director = directors[0]
		director.Pid = pid
		director.Name = name
	}
}
func SaveStar(name string, pid string) {
	var star structs.Star
	var stars []structs.Star
	engine := GetCon()
	engine.Where("name = ?", name).Find(&stars)
	if len(stars) == 0 {
		star.Id = NewKeyId()
		star.Pid = pid
		star.Name = name
		star.CreateTime = time.Now()
		engine.Insert(&star)
	} else {
		star = stars[0]
		star.Pid = pid
		star.Name = name
	}
}
func SaveChapter(name string, pid string, url, source string) structs.Chapter {
	var chapter structs.Chapter
	var chapters []structs.Chapter
	engine := GetCon()
	engine.Where("name = ?", name).And("source = ?", source).And("pid = ?", pid).Find(&chapters)
	if len(chapters) == 0 {
		chapter.Id = NewKeyId()
		chapter.Pid = pid
		chapter.Name = name
		chapter.Path = url
		chapter.Source = source
		chapter.JX = ""
		engine.Insert(&chapter)
	} else {
		chapter = chapters[0]
		chapter.Pid = pid
		chapter.Name = name
		chapter.Path = url
		chapter.JX = ""
		chapter.Source = source
	}
	return chapter
}
