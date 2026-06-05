package main

func main() {
	skill1 := Skill{
		ID:     1,
		Name:   "A",
		Desc:   "ATTACK ATTACK ATTACK",
		Params: []int{1, 2, 3},
	}
	skill2 := Skill{
		ID:     2,
		Name:   "A",
		Desc:   "ATTACK ATTBCK DEFEND",
		Params: []int{1, 2, 4},
	}

	char1 := Character{
		Name:   "A",
		Desc:   "a person",
		Skills: []Skill{},
	}
	char2 := Character{
		Name:   "B",
		Desc:   "b person",
		Skills: []Skill{skill1},
	}
	char3 := Character{
		Name:   "C",
		Desc:   "b person",
		Skills: []Skill{skill2},
	}

	CompareStructs(char1, char2)
	CompareStructs(char2, char3)
}
