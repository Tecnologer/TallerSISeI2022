package factory

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/tecnologer/SISeI/messenger"
	"github.com/tecnologer/SISeI/slack"
	"github.com/tecnologer/SISeI/telegram"
)

const seederPath = "seeder.json"

func Load(t Type) ([]messenger.Messenger, error) {
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

	messengers := make([]messenger.Messenger, 0)
	for _, service := range services {
		if t != All && service.Type != t {
			continue
		}

		switch service.Type {
		case Telegram:
			messengers = append(messengers, telegram.New(service.Token))
		case Slack:
			messengers = append(messengers, slack.New(service.Token))
		}
	}

	return messengers, nil
}

func getCompleteSeederPath() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return seederPath
	}

	return fmt.Sprintf("%s/%s", path.Dir(filename), seederPath)
}
