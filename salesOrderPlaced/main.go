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
	orderItemID := utils.GetRandomUint64ID()
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
		Items: []sales.OrderItem{
			{
				ItemID:     orderItemID,
				PlanID:     78608848,
				Lookups:    5000,
				UnitAmount: 5400,
				Amount:     5400,
			},
		},
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
	utils.PrintBold("sales:order-placed")
	fmt.Print(" copied to clipboard\n")
	utils.PrintBold("OrderItemID: ")
	fmt.Printf("%v\n", orderItemID)
}
