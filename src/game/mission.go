package game

const (
	MISSION_TYPE_COLONY    = "colony"
	MISSION_TYPE_SATELLITE = "satellite"
)

type Mission struct {
	MissionType       string
	MissionFailure    bool
	Budget            int
	SuccessfulMission bool
	MissionCrewAlive  bool
}
