package singleton

type Singleton struct {
	name string
}

func GetInstance(content string) *Singleton {
	var instance *Singleton
	if instance == nil {
		instance = &Singleton{name: content}
	}
	return instance
}

func (s *Singleton) GetName() string {
	return s.name
}
