package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/tecnologer/SISeI/factory"
	"github.com/tecnologer/SISeI/messenger"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	var option int

	messenger := messenger.NewMessenger()
	seededServices, err := factory.Load(factory.All)
	if err != nil {
		log.Fatal(err)
	}

	nameServicesList := bytes.NewBufferString("")
	i := 0
	for name, service := range seededServices {
		err := messenger.AddService(name, service)
		if err != nil {
			log.Printf("error al agregar servicio %s. Err: %v", name, err)
			continue
		}

		nameServicesList.WriteString(name)

		if i+1 < len(seededServices) {
			nameServicesList.WriteString(", ")
		}

		i++

		log.Printf("el servicio %s fue creado correctamente\n", name)

	}

	for {
		option = selectMenuOption()

		switch option {
		case 1:
			var service, dst, msg string

			fmt.Printf("\nselecciona un servicio (%s): ", nameServicesList.String())
			fmt.Scanf("%s", &service)

			fmt.Print("\nespecifica destinatario: ")
			fmt.Scanf("%s", &dst)

			fmt.Print("\nmensaje a enviar: ")
			msg, _ = reader.ReadString('\n')

			err := messenger.SendMessage(service, dst, msg)
			if err != nil {
				log.Println(err)
			}
		case 2:
			fmt.Print("\nmensaje a enviar: ")
			msg, _ := reader.ReadString('\n')

			messenger.Broadcast(msg)
		case 3:
			go func() {
				messages, err := messenger.RegisterForAllMessages()
				if err != nil {
					log.Println(err)
					return
				}

				fmt.Println("el servicio para recibir mensajes ha sido iniciado")

				for message := range messages {
					if message.From().IsBot() {
						continue
					}

					fmt.Printf("Nuevo mensaje: %s -> %s\n", message.From().GetName(), message.GetMsg())

					msg := fmt.Sprintf("`Fuente: %s, Autor: %s`\n\n%s", message.GetService(), message.From().GetName(), message.GetMsg())
					messenger.BroadcastToOtherServices(message.GetService(), msg)
				}

				fmt.Println("el servicio para recibir mensajes se ha detenido")
			}()
		case 4:
			messenger.UnregisterForAllMessages()
		case 5:
			messenger.UnregisterForAllMessages()
			os.Exit(0)
		}
	}
}

func printMenu() {
	fmt.Printf(`

	***************************************************
	* 1. Send Message                                 *
	* 2. Send broadcast (message to all services)     *
	* 3. Register for incoming messages               *
	* 4. Unregister for incoming messages             *
	* 5. Exit                                         *
	***************************************************

	`)
}

func selectMenuOption() (option int) {
	printMenu()

	fmt.Print("Select an option: ")
	fmt.Scanf("%d", &option)

	if option < 1 || option > 5 {
		fmt.Printf("the option %d is not valid. Try again.\n", option)
		option = selectMenuOption()
	}

	return
}
