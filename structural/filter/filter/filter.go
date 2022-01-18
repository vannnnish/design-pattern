package filter

type Rubbish struct {
	Name       string
	IsHarm     bool
	IsRecycled bool
	IsDry      bool
	IsWet      bool
}

// 我们过滤的标准接口，即一个抽象的过滤器
type Criteria interface {
	// 定义过滤的标准
	RubbishFilter(rubbish []Rubbish) []Rubbish
}
