package idcard

import (
	"fmt"
	"testing"
	"time"
)

type IDCardEG struct {
	Code     string
	Sex      int8
	Birthday time.Time
	Province string
	City     string
}

func (c IDCardEG) GetAge() int {
	now := time.Now()
	offset := int8(0)
	if now.Sub(time.Date(now.Year(), 0, 0, 0, 0, 0, 0, time.Local))-c.Birthday.Sub(time.Date(c.Birthday.Year(), 0, 0, 0, 0, 0, 0, time.Local)) < 0 {
		offset = -1
	}

	age := now.Year() - c.Birthday.Year() + int(offset)
	return age
}

var (
	idcardV2Man = IDCardEG{
		Code:     "130421197410056037",
		Sex:      1,
		Birthday: time.Date(1974, 10, 5, 0, 0, 0, 0, time.Local),
		Province: "河北",
		City:     "邯郸县",
	}

	idcardV2Woman = IDCardEG{
		Code:     "220381199308294161",
		Sex:      0,
		Birthday: time.Date(1993, 8, 29, 0, 0, 0, 0, time.Local),
		Province: "吉林",
		City:     "公主岭市",
	}

	idcardV1Woman = IDCardEG{
		Code:     "220381930829416",
		Sex:      0,
		Birthday: time.Date(1993, 8, 29, 0, 0, 0, 0, time.Local),
		Province: "吉林",
		City:     "公主岭市",
	}

	idcardV1Man = IDCardEG{
		Code:     "220381930829417",
		Sex:      1,
		Birthday: time.Date(1993, 8, 29, 0, 0, 0, 0, time.Local),
		Province: "吉林",
		City:     "公主岭市",
	}

	idcardV1WomanWrong = IDCardEG{
		Code:     "2203819308294171",
		Sex:      1,
		Birthday: time.Date(1993, 8, 29, 0, 0, 0, 0, time.Local),
		Province: "吉林",
		City:     "公主岭市",
	}

	idcardV2ManWrong = IDCardEG{
		Code:     "130421197410053037",
		Sex:      1,
		Birthday: time.Date(1974, 10, 5, 0, 0, 0, 0, time.Local),
		Province: "河北",
		City:     "邯郸县",
	}

	idcardV2ManBirthdayWrong = IDCardEG{
		Code:     "110101199003400754",
		Sex:      1,
		Birthday: time.Date(1974, 10, 5, 0, 0, 0, 0, time.Local),
		Province: "北京",
		City:     "东城区",
	}
	idcardV1ManBirthdayWrong = IDCardEG{
		Code:     "220381930340416",
		Sex:      0,
		Birthday: time.Date(1993, 8, 29, 0, 0, 0, 0, time.Local),
		Province: "吉林",
		City:     "公主岭市",
	}
	idcardV1ProvinceWrong = IDCardEG{
		Code: "960381930340416",
	}
	idcardV2ProvinceWrong = IDCardEG{
		Code: "980101199003073433",
	}
	idcardV1CityWrong = IDCardEG{
		Code: "110398930340416",
	}
	idcardV2CityWrong = IDCardEG{
		Code: "110198199003079197",
	}
)

func TestNewIDCard(t *testing.T) {
	// 初始化身份证对象
	card, err := NewIDCard(idcardV1Man.Code)
	if err == ErrInvalidIDCard || card == nil {
		t.Fatalf("idcardV1Man: %s should be ok", idcardV1Man.Code)
	}
	card, err = NewIDCard(idcardV2Man.Code)
	if err == ErrInvalidIDCard || card == nil {
		t.Fatalf("idcardV2Man: %s should be ok", idcardV2Man.Code)
	}
	card, err = NewIDCard(idcardV1WomanWrong.Code)
	if err != ErrInvalidIDCard || card != nil {
		t.Fatalf("idcardV1WomanWrong: %s should be wrong", idcardV1WomanWrong.Code)
	}
	card, err = NewIDCard(idcardV2ManWrong.Code)
	if err != ErrInvalidIDCard || card != nil {
		t.Fatalf("idcardV2ManWrong: %s should be wrong", idcardV2ManWrong.Code)
	}

	// 校验属性

	// gender
	v1Man, err := NewIDCard(idcardV1Man.Code)
	if err == ErrInvalidIDCard || v1Man == nil {
		t.Fatalf("idcardV1Man: %s should be ok", idcardV1Man.Code)
	}
	gender, ok := v1Man.GetGender()
	if !ok || gender != Man {
		t.Fatalf("idcardV1Man: %s sex should be ok", idcardV1Man.Code)
	}

	v2Man, err := NewIDCard(idcardV2Man.Code)
	if err == ErrInvalidIDCard || v2Man == nil {
		t.Fatalf("idcardV2Man: %s should be ok", idcardV2Man.Code)
	}
	gender, ok = v2Man.GetGender()
	if !ok || gender != Man {
		t.Fatalf("idcardV2Man: %s sex should be ok", idcardV2Man.Code)
	}

	v1Woman, err := NewIDCard(idcardV1Woman.Code)
	if err == ErrInvalidIDCard || v1Woman == nil {
		t.Fatalf("idcardV1Woman: %s should be ok", idcardV1Woman.Code)
	}
	gender, ok = v1Woman.GetGender()
	if !ok || gender != Woman {
		t.Fatalf("idcardV1Woman: %s sex should be ok", idcardV1Woman.Code)
	}

	v2Woman, err := NewIDCard(idcardV2Woman.Code)
	if err == ErrInvalidIDCard || v2Woman == nil {
		t.Fatalf("idcardV2Woman: %s should be ok", idcardV2Woman.Code)
	}
	gender, ok = v2Woman.GetGender()
	if !ok || gender != Woman {
		t.Fatalf("idcardV2Woman: %s sex should be ok", idcardV2Woman.Code)
	}

	// 生日
	v1Man, err = NewIDCard(idcardV1Man.Code)
	if err == ErrInvalidIDCard || v1Man == nil {
		t.Fatalf("idcardV1Man: %s should be ok", idcardV1Man.Code)
	}
	birthday, ok := v1Man.GetBirthday()
	if !ok {
		t.Fatalf("idcardV1Man: %s birthday should be ok", idcardV1Man.Code)
	}
	if !birthday.Equal(idcardV1Man.Birthday) {
		t.Fatalf("idcardV1Man: %s birthday not equal", idcardV1Man.Code)
	}

	v2Man, err = NewIDCard(idcardV2Man.Code)
	if err == ErrInvalidIDCard || v2Man == nil {
		t.Fatalf("idcardV2Man: %s should be ok", idcardV2Man.Code)
	}
	birthday, ok = v2Man.GetBirthday()
	if !ok {
		t.Fatalf("idcardV2Man: %s birthday should be ok", idcardV2Man.Code)
	}
	if !birthday.Equal(idcardV2Man.Birthday) {
		t.Fatalf("idcardV2Man: %s birthday not equal", idcardV2Man.Code)
	}

	// 省份
	v1Man, err = NewIDCard(idcardV1Man.Code)
	if err == ErrInvalidIDCard || v1Man == nil {
		t.Fatalf("idcardV1Man: %s should be ok", idcardV1Man.Code)
	}
	province, ok := v1Man.GetProvince()
	if !ok {
		t.Fatalf("idcardV1Man: %s Province should be ok", idcardV1Man.Code)
	}
	if province != idcardV1Man.Province {
		t.Fatalf("idcardV1Man: %s Province not equal", idcardV1Man.Code)
	}

	// 城市
	v1Man, err = NewIDCard(idcardV1Man.Code)
	if err == ErrInvalidIDCard || v1Man == nil {
		t.Fatalf("idcardV1Man: %s should be ok", idcardV1Man.Code)
	}
	city, ok := v1Man.GetCity()
	if !ok {
		t.Fatalf("idcardV1Man: %s City should be ok", idcardV1Man.Code)
	}
	if city != idcardV1Man.City {
		t.Fatalf("idcardV1Man: %s City not equal", idcardV1Man.Code)
	}

	// 版本
	v1Man, err = NewIDCard(idcardV1Man.Code)
	if err == ErrInvalidIDCard || v1Man == nil {
		t.Fatalf("idcardV1Man: %s should be ok", idcardV1Man.Code)
	}
	if v1Man.GetVersion() != V1 {
		t.Fatalf("idcardV1Man: %s version should be v1", idcardV1Man.Code)
	}

	v2Man, err = NewIDCard(idcardV2Man.Code)
	if err == ErrInvalidIDCard || v2Man == nil {
		t.Fatalf("idcardV2Man: %s should be ok", idcardV2Man.Code)
	}
	if v2Man.GetVersion() != V2 {
		t.Fatalf("idcardV2Man: %s version should be V2", idcardV2Man.Code)
	}

	// 成年判定 && 年龄
	v1Man, err = NewIDCard(idcardV1Man.Code)
	if err == ErrInvalidIDCard || v1Man == nil {
		t.Fatalf("idcardV1Man: %s should be ok", idcardV1Man.Code)
	}
	age, ok := v1Man.GetAge()
	if !ok {
		t.Fatalf("idcardV1Man: %s age should be ok", idcardV1Man.Code)
	}
	birthday, ok = v1Man.GetBirthday()
	if !ok {
		t.Fatalf("idcardV1Man: %s birthday should be ok", idcardV1Man.Code)
	}
	realAge := getAge(birthday)
	if age != realAge {
		t.Fatalf("idcardV1Man: %s age should be %v not %v", idcardV1Man.Code, realAge, age)
	}
	isAdult, ok := v1Man.IsAdult()
	if !ok {
		t.Fatalf("idcardV1Man: %s isAdult should be ok", idcardV1Man.Code)
	}
	if isAdult != (realAge >= 18) {
		t.Fatalf("idcardV1Man: %s isAdult check false", idcardV1Man.Code)
	}

	v2Man, err = NewIDCard(idcardV2Man.Code)
	if err == ErrInvalidIDCard || v2Man == nil {
		t.Fatalf("idcardV2Man: %s should be ok", idcardV2Man.Code)
	}
	age, ok = v2Man.GetAge()
	if !ok {
		t.Fatalf("idcardV2Man: %s age should be ok", idcardV2Man.Code)
	}
	birthday, ok = v2Man.GetBirthday()
	if !ok {
		t.Fatalf("idcardV2Man: %s birthday should be ok", idcardV2Man.Code)
	}
	realAge = getAge(birthday)
	if age != realAge {
		t.Fatalf("idcardV2Man: %s age should be %v not %v", idcardV2Man.Code, realAge, age)
	}
	isAdult, ok = v2Man.IsAdult()
	if !ok {
		t.Fatalf("idcardV2Man: %s isAdult should be ok", idcardV2Man.Code)
	}
	if isAdult != (realAge >= 18) {
		t.Fatalf("idcardV2Man: %s isAdult check false", idcardV2Man.Code)
	}
}

func TestCheckIDCard(t *testing.T) {
	if !CheckIDCard(idcardV1Man.Code) {
		t.Fatalf("idcardV1Man: %s should be ok", idcardV1Man.Code)
	}
	if !CheckIDCard(idcardV2Man.Code) {
		t.Fatalf("idcardV2Man: %s should be ok", idcardV2Man.Code)
	}
	if CheckIDCard(idcardV1WomanWrong.Code) {
		t.Fatalf("idcardV1WomanWrong: %s should be wrong", idcardV1WomanWrong.Code)
	}
	if CheckIDCard(idcardV2ManWrong.Code) {
		t.Fatalf("idcardV2ManWrong: %s should be wrong", idcardV2ManWrong.Code)
	}
}

func TestCheckIDCardWithOption(t *testing.T) {
	op := &CheckOption{
		Birthday: false,
		Province: false,
		City:     false,
	}
	op1 := &CheckOption{
		Birthday: true,
		Province: false,
		City:     false,
	}
	op2 := &CheckOption{
		Birthday: true,
		Province: true,
		City:     false,
	}
	op3 := &CheckOption{
		Birthday: true,
		Province: true,
		City:     true,
	}

	// op
	if !CheckIDCardWithOption(idcardV1Man.Code, op) {
		t.Fatalf("idcardV1Man: %s should be ok", idcardV1Man.Code)
	}
	if !CheckIDCardWithOption(idcardV2Man.Code, op) {
		t.Fatalf("idcardV2Man: %s should be ok", idcardV2Man.Code)
	}
	if CheckIDCardWithOption(idcardV1WomanWrong.Code, op) {
		t.Fatalf("idcardV1WomanWrong: %s should be wrong", idcardV1WomanWrong.Code)
	}
	if CheckIDCardWithOption(idcardV2ManWrong.Code, op) {
		t.Fatalf("idcardV2ManWrong: %s should be wrong", idcardV2ManWrong.Code)
	}

	// op1
	if !CheckIDCardWithOption(idcardV1Man.Code, op1) {
		t.Fatalf("idcardV1Man: %s should be ok", idcardV1Man.Code)
	}
	if !CheckIDCardWithOption(idcardV2Man.Code, op1) {
		t.Fatalf("idcardV2Man: %s should be ok", idcardV2Man.Code)
	}
	if CheckIDCardWithOption(idcardV1ManBirthdayWrong.Code, op1) {
		t.Fatalf("idcardV1ManBirthdayWrong: %s should be wrong", idcardV1WomanWrong.Code)
	}
	if CheckIDCardWithOption(idcardV2ManBirthdayWrong.Code, op1) {
		t.Fatalf("idcardV2ManBirthdayWrong: %s should be wrong", idcardV2ManWrong.Code)
	}

	// op2
	if !CheckIDCardWithOption(idcardV1Man.Code, op2) {
		t.Fatalf("idcardV1Man: %s should be ok", idcardV1Man.Code)
	}
	if !CheckIDCardWithOption(idcardV2Man.Code, op2) {
		t.Fatalf("idcardV2Man: %s should be ok", idcardV2Man.Code)
	}
	if CheckIDCardWithOption(idcardV1ProvinceWrong.Code, op2) {
		t.Fatalf("idcardV1ProvinceWrong: %s should be wrong", idcardV1WomanWrong.Code)
	}
	if CheckIDCardWithOption(idcardV2ProvinceWrong.Code, op2) {
		t.Fatalf("idcardV2ProvinceWrong: %s should be wrong", idcardV2ManWrong.Code)
	}

	// op3
	if !CheckIDCardWithOption(idcardV1Man.Code, op3) {
		t.Fatalf("idcardV1Man: %s should be ok", idcardV1Man.Code)
	}
	if !CheckIDCardWithOption(idcardV2Man.Code, op3) {
		t.Fatalf("idcardV2Man: %s should be ok", idcardV2Man.Code)
	}
	if CheckIDCardWithOption(idcardV1CityWrong.Code, op3) {
		t.Fatalf("idcardV1CityWrong: %s should be wrong", idcardV1WomanWrong.Code)
	}
	if CheckIDCardWithOption(idcardV2CityWrong.Code, op3) {
		t.Fatalf("idcardV2CityWrong: %s should be wrong", idcardV2ManWrong.Code)
	}
}

func ExampleCheckIDCard() {
	ok := CheckIDCard(idcardV2Man.Code)
	fmt.Printf("code: %s ok:%v", idcardV2Man.Code, ok)
}

func ExampleCheckIDCardWithOption() {
	ok := CheckIDCardWithOption(idcardV1Man.Code, &CheckOption{
		Birthday: true,
		Province: true,
		City:     true,
	})
	fmt.Printf("code: %s ok:%v", idcardV1Man.Code, ok)
}

func ExampleNewIDCard() {
	card, err := NewIDCard(idcardV1Man.Code)
	if err != nil {
		fmt.Printf("code: %s err:%s", idcardV1Man.Code, err)
		return
	}

	if city, ok := card.GetCity(); !ok {
		fmt.Printf("code: %s city:%s", idcardV1Man.Code, city)
	}
	if sex, ok := card.GetGender(); !ok {
		fmt.Printf("code: %s sex:%d", idcardV1Man.Code, sex)
	}
	if age, ok := card.GetAge(); !ok {
		fmt.Printf("code: %s age:%d", idcardV1Man.Code, age)
	}
	if birthday, ok := card.GetBirthday(); !ok {
		fmt.Printf("code: %s birthday:%s", idcardV1Man.Code, birthday.Format("2006-01-02"))
	}
	if province, ok := card.GetProvince(); !ok {
		fmt.Printf("code: %s province:%s", idcardV1Man.Code, province)
	}
	version := card.GetVersion()
	fmt.Printf("code: %s version:%d", idcardV1Man.Code, version)
}

func getAge(birthday time.Time) int {
	now := time.Now()
	age := now.Year() - birthday.Year()
	if now.Month() < birthday.Month() {
		age -= 1
	} else if now.Day() < birthday.Day() {
		age -= 1
	}

	return age
}
