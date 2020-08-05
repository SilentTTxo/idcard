package idcard

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// timeLayout 时间format
const timeLayout = "20060102"

// adultAge 成年岁数
const adultAge = 18

// v1 第一代身份证解析器
type v1 struct {
	code string
}

func (v *v1) check() bool {
	// 第一代身份证没有校验位
	return true
}

// GetBirthday Doc: 添加注释
func (v *v1) GetBirthday() (time.Time, bool) {
	// 取 7 - 12 位
	dateStr := "19" + v.code[6:12]
	birthday, err := time.ParseInLocation(timeLayout, dateStr, time.Local)
	if err != nil {
		return time.Time{}, false
	}

	return birthday, true
}

// GetGender Doc: 添加注释
func (v *v1) GetGender() (Gender, bool) {
	sex, _ := strconv.Atoi(string(v.code[14]))
	return Gender(uint8(sex) % 2), true
}

// GetVersion Doc: 添加注释
func (v *v1) GetVersion() Version {
	return V1
}

// v2 第二代身份证解析器
type v2 struct {
	code string
}

// v2WeightMap 权重码表
var v2WeightMap = map[int]int{
	0:  1,
	1:  0,
	2:  10,
	3:  9,
	4:  8,
	5:  7,
	6:  6,
	7:  5,
	8:  4,
	9:  3,
	10: 2,
}

// GetBirthday Doc: 添加注释
func (v *v2) GetBirthday() (time.Time, bool) {
	// 取 7 - 14 位
	dateStr := v.code[6:14]
	birthday, err := time.ParseInLocation(timeLayout, dateStr, time.Local)
	if err != nil {
		return time.Time{}, false
	}

	return birthday, true
}

// GetGender Doc: 添加注释
func (v *v2) GetGender() (Gender, bool) {
	// 取第17位
	sex, _ := strconv.Atoi(string(v.code[16]))
	return Gender(uint8(sex) % 2), true
}

// GetVersion Doc: 添加注释
func (v *v2) GetVersion() Version {
	return V2
}

func (v *v2) check() bool {
	// 第二代身份证的酷炫校验
	var idStr = strings.ToUpper(string(v.code))

	var sum int
	var signChar = ""
	for index, c := range idStr {
		var i = 18 - index
		if i != 1 {
			if v, err := strconv.Atoi(string(c)); err == nil {
				//计算加权因子
				var weight = int(math.Pow(2, float64(i-1))) % 11
				sum += v * weight
			} else {
				return false
			}
		} else {
			signChar = string(c)
		}
	}
	var a1 = v2WeightMap[sum%11]
	var a1Str = fmt.Sprintf("%d", a1)
	if a1 == 10 {
		a1Str = "X"
	}
	return a1Str == signChar
}

// common 公共的身份证结构
type common struct {
	VersionDecoder

	code string
}

// IsAdult Doc: 添加注释
func (v *common) IsAdult() (isAdult bool, ok bool) {
	age, ok := v.GetAge()
	if !ok {
		return false, false
	}
	return age >= adultAge, true
}

func (v *common) checkWithOption(op *CheckOption) bool {
	if op.Birthday {
		if _, ok := v.GetBirthday(); !ok {
			return false
		}
	}

	if op.Province && !op.City {
		if _, ok := v.GetProvince(); !ok {
			return false
		}
	}

	if op.City {
		if _, ok := v.GetCity(); !ok {
			return false
		}
	}

	return true
}

// GetCity Doc: 添加注释
func (v *common) GetCity() (city string, ok bool) {
	if city, ok = cityMap[v.code[0:6]]; !ok {
		return "", false
	}
	return city, true
}

// GetProvince Doc: 添加注释
func (v *common) GetProvince() (province string, ok bool) {
	if province, ok = provenceMap[v.code[0:2]]; !ok {
		return "", false
	}
	return province, true
}

// GetAge Doc: 添加注释
func (v *common) GetAge() (int, bool) {
	birthday, ok := v.GetBirthday()
	if !ok {
		return 0, false
	}
	now := time.Now()
	offset := int8(0)
	if now.Sub(time.Date(now.Year(), 0, 0, 0, 0, 0, 0, time.Local))-birthday.Sub(time.Date(birthday.Year(), 0, 0, 0, 0, 0, 0, time.Local)) < 0 {
		offset = -1
	}

	age := now.Year() - birthday.Year() + int(offset)
	if age < 0 {
		return 0, false
	}
	return age, true
}

// GetCode Doc: 添加注释
func (v *common) GetCode() (code string) {
	return v.code
}
