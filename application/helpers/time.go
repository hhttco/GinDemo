package helpers

import "time"

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(j).Format("2006-01-02 15:04:05") + `"`), nil
}

func (j *JsonTime) UnmarshalJSON(data []byte) error {
	local, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	*j = JsonTime(local)
	return err
}

func (j JsonTime) Format(fmtStr string) string {
	return time.Time(j).Format("2006-01-02 15:04:05")
}

func (j JsonTime) FormatWith(fmtstr string) string {
	return time.Time(j).Format(fmtstr)
}

func StringToJsonTimeWithFormat(dataStr, formatLayout string) (JsonTime, error) {
	beginDate, err := time.ParseInLocation(formatLayout, dataStr, time.Local)
	return JsonTime(beginDate), err
}
