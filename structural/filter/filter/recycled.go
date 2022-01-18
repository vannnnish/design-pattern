package filter

// 可回收垃圾
type RecycledRubbishCriteria struct{}

func (RecycledRubbishCriteria) RubbishFilter(rubbishs []Rubbish) []Rubbish {
	res := make([]Rubbish, 0)
	for _, v := range rubbishs {
		if v.IsRecycled == true {
			res = append(res, v)
		}
	}
	return res
}
