package main

import (

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"projectEuler"
)

func main() {

	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "project-euler", worker.Options{})

	w.RegisterWorkflow(projectEuler.Workflow)
	w.RegisterActivity(projectEuler.Activity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}

}
