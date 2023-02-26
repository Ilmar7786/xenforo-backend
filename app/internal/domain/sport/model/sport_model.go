package model

type (
	SportData struct {
		DATA struct {
			Sports        []Sport         `json:"SPORTS"`
			ExtendedSport []ExtendedSport `json:"EXTENDED_SPORTS"`
		} `json:"DATA"`
	}

	Sport struct {
		SportID         int    `json:"SPORT_ID"`
		EventsCount     int    `json:"EVENTS_COUNT"`
		EventsCountLive int    `json:"EVENTS_COUNT_LIVE"`
		IsPopular       int    `json:"IS_POPULAR"`
		SportName       string `json:"SPORT_NAME"`
	}

	ExtendedSport struct {
		SportIdForExtendedInfo int `json:"SPORT_ID_FOR_EXTENDED_INFO"`
		SportSort              int `json:"SPORT_SORT"`
	}
)
