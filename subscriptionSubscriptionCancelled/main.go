package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"time"

	"root/utils"

	"bitbucket.org/smartystreets/message-registry/v2/meta"
	"bitbucket.org/smartystreets/message-registry/v2/subscription"
)

var Version string

func main() {
	now := time.Now()
	message := subscription.SubscriptionCancelled{
		Metadata:       meta.Meta{Timestamp: now},
		SubscriptionID: 0,
		OrderID:        0,
		OrderItemID:    0,
		LookupsIssued:  0,
		LookupsUsed:    0,
		SecondsIssued:  0,
		SecondsUsed:    0,
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
	utils.PrintBold("subscription:subscription-cancelled")
	fmt.Print(" copied to clipboard\n")
}
