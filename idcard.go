package idcard

import (
	"errors"
	"regexp"
	"time"
)

//revive:disable:exported

var (
	// v2Reg 二代身份证校验正则
	v2Reg, _ = regexp.Compile("^" +
		"\\d{6}" + // 6位地区码
		"\\d{4}" + // 年YYYY
		"((0[1-9])|(10|11|12))" + // 月MM
		"(([0-2][1-9])|10|20|30|31)" + // 日DD
		"\\d{3}" + // 3位顺序码
		"[0-9Xx]" + // 校验码
		"$")
	// v1Reg 一代身份证校验正则
	v1Reg, _ = regexp.Compile("^" +
		"\\d{6}" + // 6位地区码
		"\\d{2}" + // 年19YY
		"((0[1-9])|(10|11|12))" + // 月MM
		"(([0-2][1-9])|10|20|30|31)" + // 日DD
		"\\d{3}" + // 3位顺序码
		"$")

	// ErrInvalidIDCard 不合法的身份证
	ErrInvalidIDCard = errors.New("invalid idcard")
)

// Version 身份证版本
type Version int8

const (
	// V1 第一代身份证
	V1 Version = 1
	// V2 第二代身份证
	V2 Version = 2
)

// 性别常量定义
type Gender uint8

const (
	// Man 男性
	Man Gender = 1
	// Woman 女性
	Woman Gender = 0
)

//revive:enable:exported

// Common 与身份证版本无关的接口
type Common interface {
	GetCity() (city string, ok bool)         // 获取市、县
	GetProvince() (province string, ok bool) // 获取省、直辖市
	GetCode() (code string)                  // 获取身份证号码
	IsAdult() (isAdult bool, ok bool)        // 是否成年
	GetAge() (int, bool)                     // 获取年龄
	checkWithOption(op *CheckOption) bool    // 自定义校验身份证，可指定校验内容
}

// VersionDecoder 身份证解析器
type VersionDecoder interface {
	GetBirthday() (time.Time, bool) // 获取生日
	GetGender() (Gender, bool)      // 获取性别，0：女 1：男
	GetVersion() Version            // 获取身份证版本
	check() bool                    // 校验是否为合法身份证,仅校验身份证合法性
}

// IDCard 身份证接口
type IDCard interface {
	Common
	VersionDecoder
}

// CheckOption 校验选项
type CheckOption struct {
	Birthday bool // 生日日期是否合法
	Province bool // 校验省、直辖市
	City     bool // 校验市、县
}

// NewIDCard 生成身份证实例
func NewIDCard(code string) (IDCard, error) {
	card, err := newIDCard(code)
	if err != nil {
		return nil, err
	}

	return card, nil
}

func newIDCard(code string) (IDCard, error) {
	var ic IDCard
	// 优先第二代身份证
	if v2Reg.Match([]byte(code)) {
		ic = &common{
			VersionDecoder: &v2{
				code: code,
			},
			code: code,
		}
	} else if v1Reg.Match([]byte(code)) {
		ic = &common{
			VersionDecoder: &v1{
				code: code,
			},
			code: code,
		}
	}

	if ic == nil || !ic.check() {
		return nil, ErrInvalidIDCard
	}

	return ic, nil
}

// CheckIDCard 检查身份证号是否符合规则
func CheckIDCard(code string) bool {
	_, err := NewIDCard(code)
	if err != nil {
		return false
	}

	return true
}

// CheckIDCardWithOption 检查身份证号是否符合规则,并且检查额外属性
func CheckIDCardWithOption(code string, option *CheckOption) bool {
	card, err := newIDCard(code)
	if err != nil {
		return false
	}

	return card.checkWithOption(option)
}
