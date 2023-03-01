package model

type SportData struct {
	Data Data `json:"DATA"`
}
type Sport struct {
	SportID         int    `json:"SPORT_ID"`
	EventsCount     int    `json:"EVENTS_COUNT"`
	EventsCountLive int    `json:"EVENTS_COUNT_LIVE"`
	IsPopular       int    `json:"IS_POPULAR"`
	SportName       string `json:"SPORT_NAME"`
}
type ExtendedSports struct {
	SportIDForExtendedInfo int `json:"SPORT_ID_FOR_EXTENDED_INFO"`
	SportSort              int `json:"SPORT_SORT"`
}
type Data struct {
	Sports         []Sport          `json:"SPORTS"`
	ExtendedSports []ExtendedSports `json:"EXTENDED_SPORTS"`
}
