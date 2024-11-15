package models

type Configuration struct {
	Brightness int32  `json:"brightness"`
	Color      string `json:"color"`
}

type Device struct {
	DeviceID           string        `json:"device_id"`
	UserID             string        `json:"user_id"`
	DeviceType         string        `json:"device_type"`
	DeviceName         string        `json:"device_name"`
	DeviceStatus       string        `json:"device_status"`
	Configuration      Configuration `json:"configuration"`
	Location           string        `json:"location"`
	FirmwareVersion    string        `json:"firmware_version"`
	ConnectivityStatus string        `json:"connectivity_status"`
	CreatedAt          string        `json:"created_at"`
	UpdatedAt          string        `json:"updated_at"`
	DeletedAt          string        `json:"deleted_at"`
}

type SingleID struct {
	DeviceID string `json:"device_id"`
}

type ListRequestOfDevice struct{}

type ListResponseOfDevice struct {
	List []Device `json:"list"`
}

type ResponseOfDevice struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
