package worldcommonmodel

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type unitTypeValue string

const (
	unitTypeStatic unitTypeValue = "static"
	unitTypePortal unitTypeValue = "portal"
)

type UnitType struct {
	value unitTypeValue
}

// Interface Implementation Check
var _ domain.ValueObject[UnitType] = (*UnitType)(nil)

func NewUnitType(value string) (UnitType, error) {
	switch value {
	case "static":
		return UnitType{
			value: unitTypeStatic,
		}, nil
	case "portal":
		return UnitType{
			value: unitTypePortal,
		}, nil
	default:
		return UnitType{}, fmt.Errorf("invalid UnitType: %s", value)
	}
}

func NewStaticUnitType() UnitType {
	unitType, _ := NewUnitType(string(unitTypeStatic))
	return unitType
}

func NewPortalUnitType() UnitType {
	unitType, _ := NewUnitType(string(unitTypePortal))
	return unitType
}

func (unitType UnitType) IsEqual(otherUnitType UnitType) bool {
	return unitType.value == otherUnitType.value
}

func (unitType UnitType) String() string {
	return string(unitType.value)
}

func (unitType UnitType) IsStatic() bool {
	return unitType.value == unitTypeStatic
}

func (unitType UnitType) IsPortal() bool {
	return unitType.value == unitTypePortal
}
