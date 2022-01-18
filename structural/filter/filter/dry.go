package filter

// 具体的过滤类
type DryRubbishCriteria struct{}

func (DryRubbishCriteria) RubbishFilter(rubbishes []Rubbish) []Rubbish {
	res := make([]Rubbish, 0)
	for _, v := range rubbishes {
		if v.IsDry == true {
			res = append(res, v)
		}
	}
	return res
}
