package evm

import "fmt"

type EVM struct {
	Name string
	CPU  string
	RAM  uint64
}

func NewEVM(name, cpu string, ram uint64) *EVM {
	return &EVM{Name: name, CPU: cpu, RAM: ram}
}

func (e *EVM) ShowSpecs() string {
	return fmt.Sprintf("Name: %s, CPU: %s, RAM: %dMB", e.Name, e.CPU, e.RAM)
}

