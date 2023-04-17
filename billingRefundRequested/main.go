package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"time"

	"root/utils"

	"bitbucket.org/smartystreets/message-registry/v2/billing"
	"bitbucket.org/smartystreets/message-registry/v2/meta"
)

func main() {
	now := time.Now()
	message := billing.RefundRequested{
		Metadata:      meta.Meta{Timestamp: now},
		AccountID:     0,
		TransactionID: 0,
		Amount:        0,
		Reason:        "test",
	}

	args := utils.GetArguments()
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

	utils.PrintSuccessMessage("billing:refund-requested")
}
