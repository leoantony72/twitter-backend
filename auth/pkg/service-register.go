package pkg

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hashicorp/consul/api"
)

// type service struct{

// }

func RegisterService() error {
	client, err := api.NewClient(&api.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}

	ID := getId()
	NAME := getName()
	PORT := getPort()
	ADDRESS := getHostname()

	register := &api.AgentServiceRegistration{
		ID:      ID,
		Name:    NAME,
		Port:    PORT,
		Address: ADDRESS,
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%v/check", ADDRESS, PORT),
			Interval: "10s",
			Timeout:  "30s",
		},
	}

	regiErr := client.Agent().ServiceRegister(register)
	if regiErr != nil {
		log.Panic(regiErr)
		log.Printf("Failed to register service: %s:%v ", ADDRESS, PORT)
		return regiErr
	} else {
		log.Printf("successfully register service: %s:%v", ADDRESS, PORT)
	}
	return nil
}

func getPort() int {
	port := GetEnv("PORT")
	intPort, _ := strconv.Atoi(port)
	return intPort

}

func getName() string {
	name := GetEnv("SERVICE")
	return name
}
func getId() string {
	Id := GetEnv("ID")
	return Id
}

func getHostname() (hostname string) {
	hostname, _ = os.Hostname()
	return hostname
}
