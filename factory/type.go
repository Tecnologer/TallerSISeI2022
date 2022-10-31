package factory

import "encoding/json"

//Type es un tipo de dato que esta basado en un byte, su uso sera crear un enumerador para los
// tipos de servicios que vendran del JSON
type Type byte

const (
	//Unknown es el valor default del tipo de dato Type
	Unknown Type = iota
	//All indica que el tipo puede ser cualquiera
	All
	//Telegram para indicar que es tipo telegram
	Telegram
	//Slack para indicar que es tipo slack
	Slack
)

var (
	//stringToType es un mapa para obtener el valor numerico (byte) de Type cuando en el JSON venga un string
	stringToType = map[string]Type{
		"telegram": Telegram,
		"slack":    Slack,
	}

	//typeToString es un mapa para definir un equivalente de type a string
	typeToString = map[Type]string{
		Telegram: "telegram",
		Slack:    "slack",
	}
)

//messengerFactory es la estructura esperada de los objetos que vienen en el archivo seeder.json
type messengerFactory struct {
	Type  Type   `json:"type"`
	Token string `json:"token"`
}

//UnmarshalJSON se encarga de convertir el valor string del JSON a valor numerico de tipo Type
//
//Es la implementacion de la interface json.Unmarshaler
func (mf *Type) UnmarshalJSON(v []byte) error {
	var tv string
	err := json.Unmarshal(v, &tv)
	if err != nil {
		return err
	}

	*mf = stringToType[tv]

	return nil
}

//MarshalJSON se encarga de convertir el valor numerico de Type a un valor string para JSON
//
//Es la implementacion de la interface json.Marshaler
func (mf Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(mf.String())
}

//String regresa el valor de Type como string
//
//Es la implementacion de la interface fmt.Stringer
func (mf Type) String() string {
	return typeToString[mf]
}
