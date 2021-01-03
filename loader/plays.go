package loader

// Plays structure
type Plays struct {
	AmeID                  string `csv:"ameId"`
	PlayID                 string `csv:"playId"`
	PlayDescription        string `csv:"playDescription"`
	Quarter                string `csv:"quarter"`
	Down                   string `csv:"down"`
	YardsToGo              string `csv:"yardsToGo"`
	PossessionTeam         string `csv:"possessionTeam"`
	PlayType               string `csv:"playType"`
	YardlineSide           string `csv:"yardlineSide"`
	YardlineNumber         string `csv:"yardlineNumber"`
	OffenseFormation       string `csv:"offenseFormation"`
	PersonnelO             string `csv:"personnelO"`
	DefendersInTheBox      string `csv:"defendersInTheBox"`
	NumberOfPassRushers    string `csv:"numberOfPassRushers"`
	PersonnelD             string `csv:"personnelD"`
	TypeDropback           string `csv:"typeDropback"`
	PreSnapVisitorScore    string `csv:"preSnapVisitorScore"`
	PreSnapHomeScore       string `csv:"preSnapHomeScore"`
	GameClock              string `csv:"gameClock"`
	AbsoluteYardlineNumber string `csv:"absoluteYardlineNumber"`
	PenaltyCodes           string `csv:"penaltyCodes"`
	PenaltyJerseyNumbers   string `csv:"penaltyJerseyNumbers"`
	PassResult             string `csv:"passResult"`
	OffensePlayResult      string `csv:"offensePlayResult"`
	PlayResult             string `csv:"playResult"`
	Epa                    string `csv:"epa"`
	IsDefensivePI          string `csv:"isDefensivePI"`
}

// NewPlays make new object from array
func NewPlays(data []string) *Plays {
	p := new(Plays)

	p.AmeID = data[0]
	p.PlayID = data[1]
	p.PlayDescription = data[2]
	p.Quarter = data[3]
	p.Down = data[4]
	p.YardsToGo = data[5]
	p.PossessionTeam = data[6]
	p.PlayType = data[7]
	p.YardlineSide = data[8]
	p.YardlineNumber = data[9]
	p.OffenseFormation = data[10]
	p.PersonnelO = data[11]
	p.DefendersInTheBox = data[12]
	p.NumberOfPassRushers = data[13]
	p.PersonnelD = data[14]
	p.TypeDropback = data[15]
	p.PreSnapVisitorScore = data[16]
	p.PreSnapHomeScore = data[17]
	p.GameClock = data[18]
	p.AbsoluteYardlineNumber = data[19]
	p.PenaltyCodes = data[20]
	p.PenaltyJerseyNumbers = data[21]
	p.PassResult = data[22]
	p.OffensePlayResult = data[23]
	p.PlayResult = data[24]
	p.Epa = data[25]
	p.IsDefensivePI = data[26]

	return p
}

// JPlays JSON version of Plays
type JPlays struct {
	AmeID                  string `json:"ameId"`
	PlayID                 string `json:"playId"`
	PlayDescription        string `json:"playDescription"`
	Quarter                string `json:"quarter"`
	Down                   string `json:"down"`
	YardsToGo              string `json:"yardsToGo"`
	PossessionTeam         string `json:"possessionTeam"`
	PlayType               string `json:"playType"`
	YardlineSide           string `json:"yardlineSide"`
	YardlineNumber         string `json:"yardlineNumber"`
	OffenseFormation       string `json:"offenseFormation"`
	PersonnelO             string `json:"personnelO"`
	DefendersInTheBox      string `json:"defendersInTheBox"`
	NumberOfPassRushers    string `json:"numberOfPassRushers"`
	PersonnelD             string `json:"personnelD"`
	TypeDropback           string `json:"typeDropback"`
	PreSnapVisitorScore    string `json:"preSnapVisitorScore"`
	PreSnapHomeScore       string `json:"preSnapHomeScore"`
	GameClock              string `json:"gameClock"`
	AbsoluteYardlineNumber string `json:"absoluteYardlineNumber"`
	PenaltyCodes           string `json:"penaltyCodes"`
	PenaltyJerseyNumbers   string `json:"penaltyJerseyNumbers"`
	PassResult             string `json:"passResult"`
	OffensePlayResult      string `json:"offensePlayResult"`
	PlayResult             string `json:"playResult"`
	Epa                    string `json:"epa"`
	IsDefensivePI          string `json:"isDefensivePI"`
}

// NewJPlays make new object
func NewJPlays(data []string) *JPlays {
	p := new(JPlays)

	p.AmeID = data[0]
	p.PlayID = data[1]
	p.PlayDescription = data[2]
	p.Quarter = data[3]
	p.Down = data[4]
	p.YardsToGo = data[5]
	p.PossessionTeam = data[6]
	p.PlayType = data[7]
	p.YardlineSide = data[8]
	p.YardlineNumber = data[9]
	p.OffenseFormation = data[10]
	p.PersonnelO = data[11]
	p.DefendersInTheBox = data[12]
	p.NumberOfPassRushers = data[13]
	p.PersonnelD = data[14]
	p.TypeDropback = data[15]
	p.PreSnapVisitorScore = data[16]
	p.PreSnapHomeScore = data[17]
	p.GameClock = data[18]
	p.AbsoluteYardlineNumber = data[19]
	p.PenaltyCodes = data[20]
	p.PenaltyJerseyNumbers = data[21]
	p.PassResult = data[22]
	p.OffensePlayResult = data[23]
	p.PlayResult = data[24]
	p.Epa = data[25]
	p.IsDefensivePI = data[26]

	return p
}
