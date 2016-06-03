package models

const (
	// Up direction
	Up = "up"
	// Same or no direction
	Same = "same"
	// Down direction
	Down = "down"
	// Red level or serious problems
	Red = "red"
	// Orange level or needs a lot of improvement
	Orange = "orange"
	// Yellow level or some issues
	Yellow = "yellow"
	// Green level or absolutely fine
	Green = "green"
)

type (
	// State is value of a single check
	State struct {
		Level     string `json:"level"`
		Direction string `json:"direction"`
	}

	// Check data object is a stored series of checks
	Check struct {
		ID        string           `json:"id"`
		Name      string           `json:"name"`
		Team      string           `json:"team"`
		Health    map[string]State `json:"health"`
		Timestamp int64            `json:"time"`
	}
)
