package portalunitmodel

type PortalUnitInfo struct {
	TargetPos *struct {
		X int `json:"x"`
		Z int `json:"z"`
	} `json:"targetPosition"`
}
