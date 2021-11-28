package lunar

import (
	"time"
)

/*
@Time : 2021/11/28 14:21
@Author : onns
@File : /lunar.go
*/

const (
	MinYear = 1900
	MaxYear = 2049
)

var Info = []int{
	19416,
	19168, 42352, 21717, 53856, 55632, 91476, 22176, 39632,
	21970, 19168, 42422, 42192, 53840, 119381, 46400, 54944,
	44450, 38320, 84343, 18800, 42160, 46261, 27216, 27968,
	109396, 11104, 38256, 21234, 18800, 25958, 54432, 59984,
	92821, 23248, 11104, 100067, 37600, 116951, 51536, 54432,
	120998, 46416, 22176, 107956, 9680, 37584, 53938, 43344,
	46423, 27808, 46416, 86869, 19872, 42416, 83315, 21168,
	43432, 59728, 27296, 44710, 43856, 19296, 43748, 42352,
	21088, 62051, 55632, 23383, 22176, 38608, 19925, 19152,
	42192, 54484, 53840, 54616, 46400, 46752, 103846, 38320,
	18864, 43380, 42160, 45690, 27216, 27968, 44870, 43872,
	38256, 19189, 18800, 25776, 29859, 59984, 27480, 23232,
	43872, 38613, 37600, 51552, 55636, 54432, 55888, 30034,
	22176, 43959, 9680, 37584, 51893, 43344, 46240, 47780,
	44368, 21977, 19360, 42416, 86390, 21168, 43312, 31060,
	27296, 44368, 23378, 19296, 42726, 42208, 53856, 60005,
	54576, 23200, 30371, 38608, 19195, 19152, 42192, 118966,
	53840, 54560, 56645, 46496, 22224, 21938, 18864, 42359,
	42160, 43600, 111189, 27936, 44448, 84835, 37744, 18936,
	18800, 25776, 92326, 59984, 27296, 108228, 43744, 37600,
	53987, 51552, 54615, 54432, 55888, 23893, 22176, 42704,
	21972, 21200, 43448, 43344, 46240, 46758, 44368, 21920,
	43940, 42416, 21168, 45683, 26928, 29495, 27296, 44368,
	84821, 19296, 42352, 21732, 53600, 59752, 54560, 55968,
	92838, 22224, 19168, 43476, 41680, 53584, 62034, 54560,
}

type Time struct {
	year, month, day     int
	hour, minute, second int
	leap                 bool
	location *time.Location
}

type Year struct {
	month []*Month
	leap  bool // 是否含有闰月
}

type Month struct {
	month int
	day   int
	leap  bool // 是否是闰月
}

func Parse(t time.Time) (lunar *Time) {
	lunar = &Time{
		location: t.Location(),
	}
	offset := int(t.Sub(getStartTime(t.Location())).Hours() / 24)
	for i := MinYear; i < MaxYear; i++ {
		daysOfYear := parseYear(i).Count()
		if offset-daysOfYear < 1 {
			lunar.year = i
			break
		} else {
			offset -= daysOfYear
		}
	}
	year := parseYear(lunar.year)
	for _, month := range year.month {
		if offset < month.day {
			lunar.month = month.month
			lunar.day = offset
			lunar.leap = month.leap
			break
		}
		offset -= month.day
	}
	return
}

// AddDate TODO 目前只做了年份的
func (t *Time) AddDate(years int, months int, days int) (res *Time) {
	res = &Time{
		year:   t.year + years,
		month:  t.month + months,
		day:    t.day + days,
		hour:   t.hour,
		minute: t.minute,
		second: t.second,
		location: t.location,
		leap:   false,
	}
	return
}

func (t *Time) ToSolar() (res time.Time) {
	offset := 0
	for i := MinYear; i < t.year; i++ {
		offset += parseYear(i).Count()
	}
	year := parseYear(t.year)
	for i := 0; year.month[i].month < t.month; i++ {
		offset += year.month[i].day
	}
	offset += t.day

	return getStartTime(t.location).AddDate(0, 0, offset)
}

func parseYear(year int) (res *Year) {
	res = &Year{
		month: make([]*Month, 0),
	}
	leapMonth := GetLeapMonth(year)
	if leapMonth > 0 {
		res.leap = true
	}
	for month := 1; month <= 12; month++ {
		res.month = append(res.month, &Month{
			month: month,
			day:   29 + GetMonthType(year, month),
			leap:  false,
		})
		if res.leap && month == leapMonth {
			res.month = append(res.month, &Month{
				month: month,
				day:   29 + GetLeapType(year),
				leap:  true,
			})
		}
	}
	return
}

func (y *Year) Count() (res int) {
	for _, m := range y.month {
		res += m.day
	}
	return
}

func GetLeapMonth(year int) (res int) {
	return Info[year-MinYear] & 0xf
}

func GetLeapType(year int) (res int) {
	return (Info[year-MinYear] & 0x10000) / 0x10000
}

func GetMonthType(year int, month int) (res int) {
	return (Info[year-MinYear] & 0x0FFFF) & (1 << (16 - month)) / (1 << (16 - month))
}

func getStartTime(location *time.Location) (t time.Time) {
	return time.Date(1900, 1, 30, 0, 0, 0, 0, location)
}
