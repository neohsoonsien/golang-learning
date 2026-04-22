// this package is to be called from other packages
package mock

type MockStudent struct {
	SetNameFunc func() error
	GetNameFunc func() (*string, error)
}

func (s *MockStudent) SetName(name string) (error) {
	if s.SetNameFunc != nil {
		return s.SetNameFunc()
	}
	return nil
}

func (s *MockStudent) GetName() (*string, error) {
	if s.GetNameFunc != nil {
		return s.GetNameFunc()
	}
	return nil, nil
}
