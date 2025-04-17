package evm

import "fmt"

type EVM struct {
	Name string
	CPU  string
	RAM  uint64
}

func (s *EVM) GetName() string {
	return s.Name
}

func (s *EVM) SetName(name string) {
	s.Name = name
}

func (s *EVM) AddRAM(addram uint64) {
	s.RAM += addram
}

func (s *EVM) ShowSpecs() {
	fmt.Printf("%s\n%s\n%d\n", s.Name, s.CPU, s.RAM)
}

