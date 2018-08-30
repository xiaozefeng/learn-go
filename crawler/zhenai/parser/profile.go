package parser

import (
	"imooc.com/learn-go/crawler/engine"
	"imooc.com/learn-go/crawler/model"
	"regexp"
	"strconv"
)

var ageReg = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)

var marriageReg = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)

var incomeReg = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)

var weightReg = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)

var genderReg = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)

var xinzuoReg = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)

var educationReg = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)

var occupationReg = regexp.MustCompile(`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)

var hokouReg = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)

var hourseReg = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)

var carReg = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

var heightReg = regexp.MustCompile(`<td><span class="label">身高：</span><span field="">([0-9]+)CM</span></td>`)

var idUrlRe= regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte,
	url string,
	name string) engine.ParseResult {
	profile := model.Profile{}
	age, err := strconv.Atoi(extractString(contents, ageReg))
	if err == nil {
		profile.Age = age
	}
	profile.Name = name
	profile.Marriage = extractString(contents, marriageReg)
	profile.Income = extractString(contents, incomeReg)
	weight, err := strconv.Atoi(extractString(contents, weightReg))
	if err == nil {
		profile.Weight = weight
	}
	profile.Gender = extractString(contents, genderReg)
	profile.Xinzuo = extractString(contents, xinzuoReg)
	profile.Education = extractString(contents, educationReg)
	profile.Occupation = extractString(contents, occupationReg)
	profile.Hokou = extractString(contents, hokouReg)
	profile.Hourse = extractString(contents, hourseReg)
	profile.Car = extractString(contents, carReg)
	height, err := strconv.Atoi(extractString(contents, heightReg))
	if err == nil {
		profile.Height = height
	}
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}
	// guess matches todo


	return result

}

func extractString(contents []byte,
	re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
