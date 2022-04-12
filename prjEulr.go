package projectEuler

import (

	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"

)

func Workflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions {
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("prjEulr wf started", "name", name)

	var res string
	err := workflow.ExecuteActivity(ctx, Activity, name).Get(ctx, &res)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}
	logger.Info("prjEulr workflow completed.", "result", res)

	return res, nil
}

func Activity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)
	return "ahoy " + name + ".", nil
}
