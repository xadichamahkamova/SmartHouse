package models

type Payload struct {
	Brightness int32 `json:"brightness"`
}

type Command struct {
	CommandID      string  `json:"command_id"`
	DeviceID       string  `json:"device_id"`
	UserID         string  `json:"user_id"`
	CommandType    string  `json:"command_type"`
	CommandPayload Payload `json:"command_payload"`
	Status         string  `json:"status"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	DeletedAt      string  `json:"deleted_at"`
}

type SingleId struct {
	CommandID string `json:"command_id"`
}

type ListRequestOfCommand struct{}

type ListResponseOfCommand struct {
	List []Command `json:"list"`
}

type ResponseOfCommand struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
