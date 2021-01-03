package loader

import "fmt"

// Week structure for NFL
//// time,x,y,s,a,dis,o,dir,event,nflId,displayName,jerseyNumber,position,frameId,team,gameId,playId,playDirection,route
type Week struct {
	Time          string `csv:"time"`
	X             string `csv:"x"`
	Y             string `csv:"y"`
	S             string `csv:"s"`
	A             string `csv:"a"`
	Dis           string `csv:"dis"`
	O             string `csv:"o"`
	Dir           string `csv:"dir"`
	Event         string `csv:"event"`
	NflID         string `csv:"nflId"`
	DisplayName   string `csv:"displayName"`
	JerseyNumber  string `csv:"jerseyNumber"`
	Position      string `csv:"position"`
	FrameID       string `csv:"frameId"`
	Team          string `csv:"team"`
	GameID        string `csv:"gameId"`
	PlayID        string `csv:"playId"`
	PlayDirection string `csv:"playDirection"`
	Route         string `csv:"route"`
}

// JWeek structure for NFL
//// time,x,y,s,a,dis,o,dir,event,nflId,displayName,jerseyNumber,position,frameId,team,gameId,playId,playDirection,route
type JWeek struct {
	Time          string `json:"time"`
	X             string `json:"x"`
	Y             string `json:"y"`
	S             string `json:"s"`
	A             string `json:"a"`
	Dis           string `json:"dis"`
	O             string `json:"o"`
	Dir           string `json:"dir"`
	Event         string `json:"event"`
	NflID         string `json:"nflId"`
	DisplayName   string `json:"displayName"`
	JerseyNumber  string `json:"jerseyNumber"`
	Position      string `json:"position"`
	FrameID       string `json:"frameId"`
	Team          string `json:"team"`
	GameID        string `json:"gameId"`
	PlayID        string `json:"playId"`
	PlayDirection string `json:"playDirection"`
	Route         string `json:"route"`
}

// Print just a string printer
func (week Week) Print() string {
	return fmt.Sprintf("%+v", week)
}

func newWeek(data []string) *Week {
	w := new(Week)
	w.Time = data[0]
	w.X = data[1]
	w.Y = data[2]
	w.S = data[3]
	w.A = data[4]
	w.Dis = data[5]
	w.O = data[6]
	w.Dir = data[7]
	w.Event = data[8]
	w.NflID = data[9]
	w.DisplayName = data[10]
	w.JerseyNumber = data[11]
	w.Position = data[12]
	w.FrameID = data[13]
	w.Team = data[14]
	w.GameID = data[15]
	w.PlayID = data[16]
	w.PlayDirection = data[17]
	w.Route = data[18]

	return w
}

func newJWeek(data []string) *JWeek {
	w := new(JWeek)
	w.Time = data[0]
	w.X = data[1]
	w.Y = data[2]
	w.S = data[3]
	w.A = data[4]
	w.Dis = data[5]
	w.O = data[6]
	w.Dir = data[7]
	w.Event = data[8]
	w.NflID = data[9]
	w.DisplayName = data[10]
	w.JerseyNumber = data[11]
	w.Position = data[12]
	w.FrameID = data[13]
	w.Team = data[14]
	w.GameID = data[15]
	w.PlayID = data[16]
	w.PlayDirection = data[17]
	w.Route = data[18]

	return w
}

// Print just a string printer
func (week JWeek) Print() string {
	return fmt.Sprintf("%+v", week)
}
