package model

type (
	LiveEventData struct {
		Data          []Data `json:"DATA"`
		Meta          Meta   `json:"META"`
		LastChangeKey string `json:"LAST_CHANGE_KEY"`
	}

	Data struct {
		Name                 string  `json:"NAME"`
		Header               string  `json:"HEADER"`
		NamePart1            string  `json:"NAME_PART_1"`
		NamePart2            string  `json:"NAME_PART_2"`
		TournamentTemplateID string  `json:"TOURNAMENT_TEMPLATE_ID"`
		CountryID            int     `json:"COUNTRY_ID"`
		CountryName          string  `json:"COUNTRY_NAME"`
		TournamentStageID    string  `json:"TOURNAMENT_STAGE_ID"`
		TournamentType       string  `json:"TOURNAMENT_TYPE"`
		TournamentID         string  `json:"TOURNAMENT_ID"`
		SourceType           int     `json:"SOURCE_TYPE"`
		HasLiveTable         int     `json:"HAS_LIVE_TABLE"`
		StandingInfo         int     `json:"STANDING_INFO"`
		TemplateID           string  `json:"TEMPLATE_ID"`
		TournamentStageType  int     `json:"TOURNAMENT_STAGE_TYPE"`
		ShortName            string  `json:"SHORT_NAME"`
		URL                  string  `json:"URL"`
		TournamentImage      string  `json:"TOURNAMENT_IMAGE"`
		Sort                 string  `json:"SORT"`
		StagesCount          int     `json:"STAGES_COUNT"`
		Zkl                  string  `json:"ZKL"`
		Zku                  string  `json:"ZKU"`
		TournamentSeasonID   string  `json:"TOURNAMENT_SEASON_ID"`
		CategoryName         string  `json:"CATEGORY_NAME"`
		Events               []Event `json:"EVENTS"`
	}

	Event struct {
		EventID                   string          `json:"EVENT_ID"`
		StartTime                 int             `json:"START_TIME"`
		StartUtime                int             `json:"START_UTIME"`
		StageType                 string          `json:"STAGE_TYPE"`
		MergeStageType            string          `json:"MERGE_STAGE_TYPE"`
		Stage                     string          `json:"STAGE"`
		Sort                      string          `json:"SORT"`
		Round                     string          `json:"ROUND"`
		VisibleRunRate            int             `json:"VISIBLE_RUN_RATE"`
		LiveMark                  string          `json:"LIVE_MARK"`
		HasLineps                 int             `json:"HAS_LINEPS"`
		StageStartTime            int             `json:"STAGE_START_TIME"`
		GameTime                  any             `json:"GAME_TIME"`
		PlayingOnSets             any             `json:"PLAYING_ON_SETS"`
		RecentOvers               any             `json:"RECENT_OVERS"`
		ShortnameHome             string          `json:"SHORTNAME_HOME"`
		HomeParticipantIds        []string        `json:"HOME_PARTICIPANT_IDS"`
		HomeParticipantTypes      []int           `json:"HOME_PARTICIPANT_TYPES"`
		HomeName                  string          `json:"HOME_NAME"`
		HomeParticipantNameOne    string          `json:"HOME_PARTICIPANT_NAME_ONE"`
		HomeEventParticipantID    string          `json:"HOME_EVENT_PARTICIPANT_ID"`
		HomeGoalVar               int             `json:"HOME_GOAL_VAR"`
		HomeScoreCurrent          string          `json:"HOME_SCORE_CURRENT"`
		HomeScorePart1            string          `json:"HOME_SCORE_PART_1"`
		HomeScorePart2            string          `json:"HOME_SCORE_PART_2"`
		HomeImages                []string        `json:"HOME_IMAGES"`
		Imm                       string          `json:"IMM"`
		Imw                       string          `json:"IMW"`
		Imp                       string          `json:"IMP"`
		Ime                       string          `json:"IME"`
		ShortnameAway             string          `json:"SHORTNAME_AWAY"`
		AwayParticipantIds        []string        `json:"AWAY_PARTICIPANT_IDS"`
		AwayParticipantTypes      []int           `json:"AWAY_PARTICIPANT_TYPES"`
		AwayName                  string          `json:"AWAY_NAME"`
		AwayParticipantNameOne    string          `json:"AWAY_PARTICIPANT_NAME_ONE"`
		AwayEventParticipantID    string          `json:"AWAY_EVENT_PARTICIPANT_ID"`
		AwayGoalVar               int             `json:"AWAY_GOAL_VAR"`
		AwayScoreCurrent          string          `json:"AWAY_SCORE_CURRENT"`
		AwayScorePart1            string          `json:"AWAY_SCORE_PART_1"`
		AwayScorePart2            string          `json:"AWAY_SCORE_PART_2"`
		AwayImages                []string        `json:"AWAY_IMAGES"`
		TvLiveStreaming           TvLiveStreaming `json:"TV_LIVE_STREAMING"`
		HasLiveCentre             int             `json:"HAS_LIVE_CENTRE"`
		An                        string          `json:"AN"`
		BookmakersWithLiveInOffer []string        `json:"BOOKMAKERS_WITH_LIVE_IN_OFFER"`
		LiveInOfferBookmakerID    int             `json:"LIVE_IN_OFFER_BOOKMAKER_ID"`
		LiveInOfferStatus         int             `json:"LIVE_IN_OFFER_STATUS"`
	}

	TvLiveStreaming struct {
		Num2 []Num2 `json:"2"`
	}

	Num2 struct {
		Bu string `json:"BU"`
		Iu string `json:"IU"`
		Bn string `json:"BN"`
		Bi int    `json:"BI"`
		Bt string `json:"BT"`
	}

	Meta struct {
		Bookmakers         []Bookmakers       `json:"BOOKMAKERS"`
		DatacoreTranslates DatacoreTranslates `json:"DATACORE_TRANSLATES"`
	}

	Bookmakers struct {
		BookmakerID          int    `json:"BOOKMAKER_ID"`
		BookmakerBettingType int    `json:"BOOKMAKER_BETTING_TYPE"`
		BookmakerName        string `json:"BOOKMAKER_NAME"`
	}

	DatacoreTranslates struct {
	}
)
