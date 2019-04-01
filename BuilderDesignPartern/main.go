package main

import "fmt"

type ElectronicProduct struct {
	Structure string
	Monitor   int
	Camera    int
}

type BuildProcess interface {
	SetStructure() BuildProcess
	SetCamera() BuildProcess
	SetMonitor() BuildProcess
	GetGadget() ElectronicProduct
}
type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

func (m *ManufacturingDirector) Construct() ElectronicProduct {
	m.builder.SetStructure().SetMonitor().SetCamera()
	return m.builder.GetGadget()
}

func (m *ManufacturingDirector) PrintProduct() {
	gadget := m.builder.GetGadget()
	fmt.Printf("Structure: %s \n", gadget.Structure)
	fmt.Printf("Monitor: %d \n", gadget.Monitor)
	fmt.Printf("Camera: %d \n", gadget.Camera)
	fmt.Printf("===============\n")
}

type Laptop struct {
	electronicProduct ElectronicProduct
}

func (l *Laptop) SetStructure() BuildProcess {
	l.electronicProduct.Structure = "Laptop"
	return l
}

func (l *Laptop) SetMonitor() BuildProcess {
	l.electronicProduct.Monitor = 1
	return l
}

func (l *Laptop) SetCamera() BuildProcess {
	l.electronicProduct.Camera = 1
	return l
}

func (l *Laptop) GetGadget() ElectronicProduct {
	return l.electronicProduct
}

type Smartphone struct {
	electronicProduct ElectronicProduct
}

func (s *Smartphone) SetStructure() BuildProcess {
	s.electronicProduct.Structure = "Smartphone"
	return s
}

func (s *Smartphone) SetMonitor() BuildProcess {
	s.electronicProduct.Monitor = 1
	return s
}

func (s *Smartphone) SetCamera() BuildProcess {
	s.electronicProduct.Camera = 2
	return s
}

func (s *Smartphone) GetGadget() ElectronicProduct {
	return s.electronicProduct
}
func main() {
	manufacturingDirector := ManufacturingDirector{}

	laptop := &Laptop{}
	manufacturingDirector.SetBuilder(laptop)
	manufacturingDirector.Construct()
	manufacturingDirector.PrintProduct()

	smartphone := &Smartphone{}
	manufacturingDirector.SetBuilder(smartphone)
	manufacturingDirector.Construct()
	manufacturingDirector.PrintProduct()
}
