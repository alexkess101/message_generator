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

func main() {
	now := time.Now()
	message := sales.OrderPlaced{
		Metadata:          meta.Meta{Timestamp: now},
		Timestamp:         now,
		AccountID:         0,
		OrderID:           0,
		PricingScheduleID: 0,
		TaxableLocationID: 0,
		TaxRateID:         0,
		SubscriptionID:    0,
		UserID:            0,
		IPAddress:         "",
		UserAgent:         "",
		Reason:            "test",
		Terms:             0,
		AgreementID:       "",
		Items:             nil,
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

	utils.PrintSuccessMessage("sales:order-placed")
}
