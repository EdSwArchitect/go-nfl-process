package loader

// Games structure for NFL
//// gameId,gameDate,gameTimeEastern,homeTeamAbbr,visitorTeamAbbr,week
type Games struct {
	GameID          string `csv:"gameId"`
	GameDate        string `csv:"gameDate"`
	GameTimeEastern string `csv:"gameTimeEastern"`
	HomeTeamAbbr    string `csv:"homeTeamAbbr"`
	VisitorTeamAbbr string `csv:"visitorTeamAbbr"`
	Week            string `csv:"week"`
}

// JGames structure for NFL
type JGames struct {
	GameID          string `json:"gameId"`
	GameDate        string `json:"gameDate"`
	GameTimeEastern string `json:"gameTimeEastern"`
	HomeTeamAbbr    string `json:"homeTeamAbbr"`
	VisitorTeamAbbr string `json:"visitorTeamAbbr"`
	Week            string `json:"week"`
}

// NewGames thingy
func NewGames(data []string) *Games {
	w := new(Games)
	w.GameID = data[0]
	w.GameDate = data[1]
	w.GameTimeEastern = data[2]
	w.HomeTeamAbbr = data[3]
	w.VisitorTeamAbbr = data[4]
	w.Week = data[5]
	return w
}

// NewJGames thingy
func NewJGames(data []string) *JGames {
	w := new(JGames)
	w.GameID = data[0]
	w.GameDate = data[1]
	w.GameTimeEastern = data[2]
	w.HomeTeamAbbr = data[3]
	w.VisitorTeamAbbr = data[4]
	w.Week = data[5]

	return w
}
