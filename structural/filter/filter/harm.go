package filter

// 有害垃圾
type HarmfulRubbishCriteria struct{}

func (HarmfulRubbishCriteria) RubbishFilter(rubbishs []Rubbish) []Rubbish {
	res := make([]Rubbish, 0)
	for _, v := range rubbishs {
		if v.IsHarm == true {
			res = append(res, v)
		}
	}
	return res
}
