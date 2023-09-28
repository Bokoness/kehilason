package dto

type CreateUpdateAssignment struct {
	Start     string `validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	End       string `validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	StartHour string `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"start_hour"`
	EndHour   string `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"end_hour"`
}
