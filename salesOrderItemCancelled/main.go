package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"time"

	"root/utils"

	"bitbucket.org/smartystreets/message-registry/v2/meta"
	"bitbucket.org/smartystreets/message-registry/v2/sales"
)

var Version string

func main() {
	now := time.Now()
	message := sales.OrderItemCancelled{
		Metadata:               meta.Meta{Timestamp: now},
		AccountID:              0,
		OrderID:                0,
		OrderItemID:            0,
		PlanID:                 0,
		ReplacementOrderID:     0,
		ReplacementOrderItemID: 0,
		ReplacementPlanID:      0,
		Reason:                 "test",
		CompensationPlan:       0,
	}

	args := utils.GetArguments(Version)
	messageValue := reflect.ValueOf(&message).Elem()

	for i := 0; i < len(args); i += 2 {
		utils.ApplyValueToMessage(args[i], args[i+1], messageValue)
	}

	orderJSON, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	err = utils.WriteClipboard(string(orderJSON))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Print("Success!! ")
	utils.PrintBold("sales:order-item-cancelled")
	fmt.Print(" copied to clipboard\n")
}
