package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

type User struct {
	Id       int       `json:"id"`
	Username string    `json:"username"`
	Profile  Profile   `json:"profile"`
	Articles []Article `json:"articles"`
}

type Profile struct {
	FullName string   `json:"full_name"`
	BirthDay string   `json:"birthday"`
	Phones   []string `json:"phones"`
}

type Article struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	PublishedAt string `json:"published_at"`
}

// Users with no phone number
func noPhoneNumber(users []*User) {
	var result []*User
	for _, v := range users {
		if len(v.Profile.Phones) == 0 {
			result = append(result, v)
		}
	}
	data, _ := json.MarshalIndent(result, "", " ")
	fmt.Println("Users with no phone number")
	fmt.Println(string(data))
}

// Users with article
func haveArticles(users []*User) {
	var result []*User
	for _, v := range users {
		if len(v.Articles) > 0 {
			result = append(result, v)
		}
	}
	data, _ := json.MarshalIndent(result, "", " ")
	fmt.Println("Users with article")
	fmt.Println(string(data))
}

// Users with "annis" on their name
func haveAnnis(users []*User) {
	var result []*User
	for _, v := range users {
		matched, _ := regexp.MatchString(`annis`, strings.ToLower(v.Profile.FullName))
		if matched {
			result = append(result, v)
		}
	}
	data, _ := json.MarshalIndent(result, "", " ")
	fmt.Println("Users with name annis")
	fmt.Println(string(data))
}

// Users with article on the year 2000
func articleOnTheYearTwoThousandTwenty(users []*User) {
	var result []*User
	for _, v := range users {
		for _, article := range v.Articles {
			if article.PublishedAt[0:4] == "2020" {
				result = append(result, v)
				break
			}
		}
	}
	data, _ := json.MarshalIndent(result, "", " ")
	fmt.Println("Users with article on the year 2000")
	fmt.Println(string(data))
}

// Users who are born in 1986
func bornEightySix(users []*User) {
	var result []*User
	for _, v := range users {
		if v.Profile.BirthDay[0:4] == "1986" {
			result = append(result, v)
		}
	}
	data, _ := json.MarshalIndent(result, "", " ")
	fmt.Println("Users who are born in 1986")
	fmt.Println(string(data))
}

// Articles that contain tips on the title
func titleContainTips(users []*User) {
	//var result []*Article
	fmt.Println("Articles that contain tips on the title")
	for _, v := range users {
		for _, article := range v.Articles {
			if strings.Contains(strings.ToLower(article.Title), "tips") {
				fmt.Println(article)
				//result = append(result, &article)
			}
		}
	}
	//data, _ := json.MarshalIndent(result, "", " ")
	//fmt.Println(string(data))
}

// Articles that are published before August 2019
func articleBeforeAugustTwentyNineteen(users []*User) {
	var layout = "2006-01-02T15:04:05"
	//var result []*Article
	fmt.Println("Articles that are published before August 2019")
	for _, v := range users {
		for _, article := range v.Articles {
			t, _ := time.Parse(layout, article.PublishedAt)
			if int(t.Month()) < 8 && t.Year() <= 2019 {
				fmt.Println(article)
				//result = append(result, &article)
			}
		}
	}
	//data, _ := json.MarshalIndent(result, "", " ")
	//fmt.Println(string(data))
}

func main() {
	jsonFile, err := os.Open("profile_list.json")

	if err != nil {
		panic(err)
	}

	fmt.Println("Success")

	defer jsonFile.Close()

	file, _ := ioutil.ReadAll(jsonFile)

	var users []*User

	err = json.Unmarshal(file, &users)
	if err != nil {
		panic(err)
	}

	noPhoneNumber(users)
	haveArticles(users)
	haveAnnis(users)
	articleOnTheYearTwoThousandTwenty(users)
	bornEightySix(users)
	titleContainTips(users)
	articleBeforeAugustTwentyNineteen(users)
}
