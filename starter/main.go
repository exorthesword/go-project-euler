
package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"

	"projectEuler"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID: "hi_id",
		TaskQueue: "project-euler",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, projectEuler.Workflow, "Temporal")
	if err != nil {
		log.Fatalln("unable to execute workflow", err)
	}
	
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var res string
	err = we.Get(context.Background(), &res)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}
	log.Println("Workflow result:", res)

}
