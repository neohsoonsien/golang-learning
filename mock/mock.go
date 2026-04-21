package mock

type Student struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
func (s *Student) SetName(name string) error {
	s.Name = name

	return nil
}

func (s *Student) GetName() (*string, error) {
	if s != nil {
		return &s.Name, nil
	}
	return nil, nil
}

func (s *Student) SetId(id string) error {
	s.Id = id

	return nil
}

func (s *Student) GetId() (*string, error) {
	if s != nil {
		return &s.Id, nil
	}
	return nil, nil
}
