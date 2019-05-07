package flyweight

// Allows sharing the state of a heavy object between many
// instances of a type

// If you hvae to create and store too many objects of a 
// large-memory type that are fundamentally equal

// Typically compined with the factory pattern for instantiation
// and we will share all states in a single common object
// ( LOTS of pointers to already created objects )

import (
	"time"
)

const (
	TEAM_A = iota
	TEAM_B
)

type Player struct {
	Name		 string
	Surname		 string
	PreviousTeam uint64
	Photo		 []byte
}

type HistoricalData struct {
	Year		  uint8
	LeaugeResults []Match
}

type Team struct {
	ID			   uint64
	Name		   string
	Shield		   []byte
	Players		   []Player
	HistoricalData []HistoricalData
}

type Match struct {
	Date		  time.Time
	VisitorID	  uint64
	LocalID		  uint64
	LocalScore 	  byte
	VisitorScore  byte
	LocalShoots	  uint16
	VisitorShoots uint16
}

func getTeamFactory(team int) Team {
	switch team {
	case TEAM_B:
		return Team {
			ID: 2,
			Name: "TEAM_B",
		}
	default:
		return Team {
			ID: 1,
			Name: "TEAM_A",
		}
	}
}

func NewTeamFactory() teamFlyweightFactory {
	return teamFlyweightFactory{
		createdTeams: make(map[int]*Team, 0),
	}
}

type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func (t *teamFlyweightFactory) GetTeam(teamName int) *Team {
	if t.createdTeams[teamName] != nil {
		return t.createdTeams[teamName]
	}

	team := getTeamFactory(teamName)
	t.createdTeams[teamName] = &team

	return t.createdTeams[teamName]
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}