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
	return e.Name + " | CPU: " + e.CPU + " | RAM: " + fmt.Sprint(e.RAM) + " GB"
}

