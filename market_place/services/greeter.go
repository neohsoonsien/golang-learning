package services

import "fmt"

type GreeterService struct {
	// Could include dependencies like repositories, configs, etc.
}

func NewGreeterService() *GreeterService {
	return &GreeterService{}
}

// business logics
func (s *GreeterService) GenerateGreeting(name string) (string, error) {
	// validation and business rules, etc.
	if name == "" {
		return "", fmt.Errorf("name cannot be empty")
	}

	return "Hello " + name, nil
}
