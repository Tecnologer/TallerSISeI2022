package factory

import "encoding/json"

type Type byte

const (
	Unknown Type = iota
	All
	Telegram
	Slack
)

var (
	stringToType = map[string]Type{
		"telegram": Telegram,
		"slack":    Slack,
	}

	typeToString = map[Type]string{
		Telegram: "telegram",
		Slack:    "slack",
	}
)

type messengerFactory struct {
	Type  Type   `json:"type"`
	Token string `json:"token"`
}

func (mf *Type) UnmarshalJSON(v []byte) error {
	var tv string
	err := json.Unmarshal(v, &tv)
	if err != nil {
		return err
	}

	*mf = stringToType[tv]

	return nil
}

func (mf Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(mf.String())
}

func (mf Type) String() string {
	return typeToString[mf]
}
