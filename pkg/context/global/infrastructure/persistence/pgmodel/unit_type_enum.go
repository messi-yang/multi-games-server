package pgmodel

type UnitTypeEnum string

const (
	UnitTypeEnumStatic UnitTypeEnum = "static"
	UnitTypeEnumFence  UnitTypeEnum = "fence"
	UnitTypeEnumPortal UnitTypeEnum = "portal"
	UnitTypeEnumLink   UnitTypeEnum = "link"
	UnitTypeEnumEmbed  UnitTypeEnum = "embed"
)
