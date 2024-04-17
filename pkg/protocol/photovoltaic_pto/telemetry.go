package photovoltaic_pto

type Telemetry struct {
	Keep0  uint32 `json:"keep_0,omitempty"`
	Type   byte   `json:"type,omitempty"`
	Object byte   `json:"object,omitempty"`
}
