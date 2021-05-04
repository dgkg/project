package model

type Relation struct {
	ID             int `json:"id"`
	Dateslocations struct {
		DunedinNewZealand []string `json:"dunedin-new_zealand"`
		GeorgiaUsa        []string `json:"georgia-usa"`
		LosAngelesUsa     []string `json:"los_angeles-usa"`
		NagoyaJapan       []string `json:"nagoya-japan"`
		NorthCarolinaUsa  []string `json:"north_carolina-usa"`
		OsakaJapan        []string `json:"osaka-japan"`
		PenroseNewZealand []string `json:"penrose-new_zealand"`
		SaitamaJapan      []string `json:"saitama-japan"`
	} `json:"datesLocations"`
}
