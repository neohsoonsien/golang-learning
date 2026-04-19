package mock

type Student struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func (s *Student) GetName() (*string, error) {
	if s != nil {
		return &s.Name, nil
	}
	return nil, nil
}
