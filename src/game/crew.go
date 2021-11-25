package game

import "errors"

const (
	CREW_COMMANDER  = "Commander"
	CREW_PILOT      = "Pilot"
	CREW_MEDICAL    = "Medical"
	CREW_SPECIALIST = "Specialist"
	CREW_ENGINEER   = "Engineer"
)

type Crew struct {
	Commander  CrewMember
	Pilot      CrewMember
	Medical    CrewMember
	Specialist CrewMember
	Engineer   CrewMember
	Manifest   []string
}

type CrewMember struct {
	Dead  bool
	Skill int
}

func (cm CrewMember) TestSkill(difficulty int) bool {
	if cm.Dead {
		return false
	}
	result := Random(difficulty)
	return result <= cm.Skill
}

func (cm CrewMember) Status() string {
	if cm.Dead {
		return "dead"
	}
	return "alive"
}

func (c *Crew) WhoIsAlive() ([]string, error) {
	var newManifest []string
	if !c.Commander.Dead {
		newManifest = append(newManifest, CREW_COMMANDER)
	}
	if !c.Pilot.Dead {
		newManifest = append(newManifest, CREW_PILOT)
	}
	if !c.Medical.Dead {
		newManifest = append(newManifest, CREW_MEDICAL)
	}
	if !c.Specialist.Dead {
		newManifest = append(newManifest, CREW_SPECIALIST)
	}
	if !c.Engineer.Dead {
		newManifest = append(newManifest, CREW_ENGINEER)
	}
	if len(newManifest) == 0 {
		return nil, errors.New("they're dead, Jim")
	}
	return newManifest, nil
}

func (c *Crew) WhoToKill() (crewmember string, err error) {
	c.Manifest, err = c.WhoIsAlive()
	if err != nil {
		return
	}
	i := Random(len(c.Manifest))
	crewmember = c.Manifest[i]
	return
}

func (c *Crew) KillCrewMember(crewmember string) {
	switch crewmember {
	case CREW_COMMANDER:
		c.Commander.Dead = true
	case CREW_PILOT:
		c.Pilot.Dead = true
	case CREW_MEDICAL:
		c.Medical.Dead = true
	case CREW_SPECIALIST:
		c.Specialist.Dead = true
	case CREW_ENGINEER:
		c.Engineer.Dead = true
	}
	c.Manifest, _ = c.WhoIsAlive()
}
