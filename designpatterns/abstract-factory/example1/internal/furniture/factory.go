package furniture

import "fmt"

type FurnitureType string

const (
	ModernFurnitureType    FurnitureType = "modern"
	VictorianFurnitureType FurnitureType = "victorian"
)

type FurnitureFactory interface {
	CreateChair() Chair
	CreateTable() Table
	CreateSofa() Sofa
}

func NewFurnitureFactory(style FurnitureType) (FurnitureFactory, error) {
	switch style {
	case ModernFurnitureType:
		return NewModernFactory(), nil
	case VictorianFurnitureType:
		return NewVictorianFurnitureFactory(), nil
	default:
		return nil, fmt.Errorf("invalid furniture type %s", style)
	}
}
