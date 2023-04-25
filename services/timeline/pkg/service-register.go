package pkg

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hashicorp/consul/api"
)

func RegisterService() error {
	config := api.DefaultConfig()
	config.Address = "http://service_register:8500"
	client, err := api.NewClient(config)

	if err != nil {
		fmt.Println(err)
		return err
	}

	PORT, ID, NAME := getConstaints()
	ADDRESS := getHostname()
	register := &api.AgentServiceRegistration{
		ID:      ID,
		Name:    NAME,
		Port:    PORT,
		Address: ADDRESS,
		Tags:    []string{"Twitter", "Timeline"},
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%v/check", ADDRESS, PORT),
			Interval: "10s",
			Timeout:  "10s",
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

func getConstaints() (int, string, string) {
	Port := GetEnv("PORT")
	Id := GetEnv("ID")
	Name := GetEnv("SERVICE")

	intPort, _ := strconv.Atoi(Port)

	return intPort, Id, Name
}

func getHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}
