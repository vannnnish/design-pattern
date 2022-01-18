package filter

type WetRubbishCriteria struct {
}

func (WetRubbishCriteria) RubbishFilter(rubbishes []Rubbish) []Rubbish {
	res := make([]Rubbish, 0)
	for _, v := range rubbishes {
		if v.IsWet == true {
			res = append(res, v)
		}
	}
	return res
}
