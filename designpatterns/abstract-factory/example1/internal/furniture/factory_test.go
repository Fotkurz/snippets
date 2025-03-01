package furniture_test

import (
	"abstractfactoryexample1/internal/furniture"
	"testing"
)

func Test_NewFurnitureFactory(t *testing.T) {
	t.Run("return a modern furnitures", func(t *testing.T) {
		f, err := furniture.NewFurnitureFactory(furniture.ModernFurnitureType)

		if err != nil {
			t.Errorf("expected error to be nil but: %v", err)
		}

		chair := f.CreateChair()
		table := f.CreateTable()
		sofa := f.CreateSofa()

		if chair.TypeName() != furniture.ModernFurnitureType {
			t.Errorf("expected %s but got %s", furniture.ModernFurnitureType, chair.TypeName())
		}
		if table.TypeName() != furniture.ModernFurnitureType {
			t.Errorf("expected %s but got %s", furniture.ModernFurnitureType, table.TypeName())
		}
		if sofa.TypeName() != furniture.ModernFurnitureType {
			t.Errorf("expected %s but got %s", furniture.ModernFurnitureType, sofa.TypeName())
		}
	})

	t.Run("return a victorian furnitures", func(t *testing.T) {
		f, err := furniture.NewFurnitureFactory(furniture.VictorianFurnitureType)

		if err != nil {
			t.Errorf("expected error to be nil but: %v", err)
		}

		chair := f.CreateChair()
		table := f.CreateTable()
		sofa := f.CreateSofa()

		if chair.TypeName() != furniture.VictorianFurnitureType {
			t.Errorf("expected %s but got %s", furniture.VictorianFurnitureType, chair.TypeName())
		}
		if table.TypeName() != furniture.VictorianFurnitureType {
			t.Errorf("expected %s but got %s", furniture.VictorianFurnitureType, table.TypeName())
		}
		if sofa.TypeName() != furniture.VictorianFurnitureType {
			t.Errorf("expected %s but got %s", furniture.VictorianFurnitureType, sofa.TypeName())
		}
	})

	t.Run("return an error", func(t *testing.T) {
		f, err := furniture.NewFurnitureFactory("invalid type")
		if err == nil {
			t.Error("expected an error but got nil")
		}
		if f != nil {
			t.Errorf("expected factory to be nil but got %v", f)
		}

		if err.Error() != "invalid furniture type invalid type" {
			t.Errorf("expected invalid furniture type err but got: '%s'", err)
		}
	})

}
