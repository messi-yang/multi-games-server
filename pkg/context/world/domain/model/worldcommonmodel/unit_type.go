package worldcommonmodel

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type unitTypeValue string

const (
	unitTypeStatic unitTypeValue = "static"
	unitTypeFence  unitTypeValue = "fence"
	unitTypePortal unitTypeValue = "portal"
	unitTypeLink   unitTypeValue = "link"
	unitTypeEmbed  unitTypeValue = "embed"
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
	case "fence":
		return UnitType{
			value: unitTypeFence,
		}, nil
	case "portal":
		return UnitType{
			value: unitTypePortal,
		}, nil
	case "link":
		return UnitType{
			value: unitTypeLink,
		}, nil
	case "embed":
		return UnitType{
			value: unitTypeEmbed,
		}, nil
	default:
		return UnitType{}, fmt.Errorf("invalid UnitType: %s", value)
	}
}

func NewStaticUnitType() UnitType {
	unitType, _ := NewUnitType(string(unitTypeStatic))
	return unitType
}

func NewFenceUnitType() UnitType {
	unitType, _ := NewUnitType(string(unitTypeFence))
	return unitType
}

func NewPortalUnitType() UnitType {
	unitType, _ := NewUnitType(string(unitTypePortal))
	return unitType
}

func NewLinkUnitType() UnitType {
	unitType, _ := NewUnitType(string(unitTypeLink))
	return unitType
}

func NewEmbedUnitType() UnitType {
	unitType, _ := NewUnitType(string(unitTypeEmbed))
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

func (unitType UnitType) IsFence() bool {
	return unitType.value == unitTypeFence
}

func (unitType UnitType) IsPortal() bool {
	return unitType.value == unitTypePortal
}

func (unitType UnitType) IsLink() bool {
	return unitType.value == unitTypeLink
}

func (unitType UnitType) IsEmbed() bool {
	return unitType.value == unitTypeEmbed
}
