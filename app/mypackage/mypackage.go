package Stuff

import (
	"errors"
	"strconv"
)

var Name string = "Gaurav"

func IntArrToStrArr(arr []int) []string {
	var strArr []string
	for _, v := range arr {
		strArr = append(strArr, strconv.Itoa(v))
	}
	return strArr
}

type Date struct {
	day   int
	month int
	year  int
}

func (d *Date) SetDay(dat int) error {
	if dat < 1 || dat > 31 {
		return errors.New("invalid day")
	}
	d.day = dat
	return nil
}

func (d *Date) SetMonth(dat int) error {
	if dat < 1 || dat > 12 {
		return errors.New("invalid month")
	}
	d.month = dat
	return nil
}

func (d *Date) SetYear(dat int) error {
	if dat < 1 {
		return errors.New("invalid year")
	}
	d.year = dat
	return nil
}

func (d *Date) GetDay() int {
	return d.day
}

func (d *Date) GetMonth() int {
	return d.month
}

func (d *Date) GetYear() int {
	return d.year
}

func (d *Date) GetDate() string {
	return strconv.Itoa(d.day) + "-" + strconv.Itoa(d.month) + "-" + strconv.Itoa(d.year)
}

