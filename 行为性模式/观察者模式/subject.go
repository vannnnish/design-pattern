package main

type Subject struct {
	obs     []Observer // 群聊
	context string     // 更新群聊信息
}


func NewSubject() *Subject {
	return &Subject{
		obs: make([]Observer, 0),
	}
}

func (s *Subject) Attach(o Observer) {
	s.obs = append(s.obs, o)
}

func (s *Subject) notify() {
	for _, o := range s.obs {
		o.Update(s)
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}
