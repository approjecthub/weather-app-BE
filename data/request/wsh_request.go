package request

type DeleteWeatherSearchHistoryRequest struct {
	WshIds []uint `validate:"required" json:"wshIds"`
}
