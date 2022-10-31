//factory es un paquete encargado de automatizar la carga de servicios
//de mensajeria de manera automatica
package factory

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/tecnologer/SISeI/messenger"
	"github.com/tecnologer/SISeI/slack"
	"github.com/tecnologer/SISeI/telegram"
)

//seederPath nombre del archivo JSON de donde se leeran los tokens para iniciarlizar los servicios
const seederPath = "seeder.json"

//Load carga la informacion de los servicios que se encuentran en el archivo seeder.json
//
//retorna un map (hashtable) donde el key corresponde al nombre y el valor es la instancia del servicio
func Load(t Type) (map[string]messenger.Messenger, error) {
	var services []*messengerFactory

	f, err := os.Open(getCompleteSeederPath())
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&services)
	if err != nil {
		return nil, err
	}

	messengers := map[string]messenger.Messenger{}

	for i, service := range services {
		if t != All && service.Type != t {
			continue
		}

		name := fmt.Sprintf("%s%d", service.Type, i)

		if service := createService(service.Type, service.Token); service != nil {
			messengers[name] = service
		}

		log.Printf("factory: servicio %s cargado correctamente\n", service.Type)
	}

	return messengers, nil
}

//getCompleteSeederPath regresa un string que corresponde a la ruta absoluta del archivo seeder.json
func getCompleteSeederPath() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return seederPath
	}

	return fmt.Sprintf("%s/%s", path.Dir(filename), seederPath)
}

//createService crea una instancia del tipo de servicio especificado
func createService(t Type, token string) messenger.Messenger {
	switch t {
	case Telegram:
		return telegram.New(token)
	case Slack:
		return slack.New(token)
	}

	return nil
}
