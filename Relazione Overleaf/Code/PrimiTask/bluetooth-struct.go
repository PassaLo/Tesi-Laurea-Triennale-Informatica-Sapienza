type BluetoothConnectionsStruct struct {
	Trip  *int         `json:"trip,omitempty"`
	Mode  *string      `json:"mode"`
	Scans []ScanStruct `json:"scans"`
}