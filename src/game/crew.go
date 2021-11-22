package game

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
	result := Random(difficulty)
	//fmt.Println("Skill Test Result: ", result)
	//fmt.Println("Needed to be at or below: ", cm.Skill)
	return result <= cm.Skill
}

func (cm CrewMember) Status() string {
	if cm.Dead {
		return "dead"
	}
	return "alive"
}
