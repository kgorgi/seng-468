package e2e

import (
	"strconv"
	"testing"
	"time"

	"extremeWorkload.com/daytrader/lib"
	userClient "extremeWorkload.com/daytrader/lib/user"
)

func TestTriggerBuy(t *testing.T) {
	userid := "thewolf"
	var addAmount uint64 = 1000234
	var buyAmount uint64 = 5000
	var triggerPrice uint64 = 500
	stockSymbol := "DOG"

	status, body, _ := userClient.CancelSetBuyRequest(userid, stockSymbol)
	status, body, _ = userClient.CancelSetSellRequest(userid, stockSymbol)

	status, body, _ = userClient.AddRequest(userid, lib.CentsToDollars(addAmount))
	if status != lib.StatusOk {
		t.Error("add failed\n" + strconv.Itoa(status) + body)
	}

	summaryBefore, err := userClient.GetSummary(userid)
	if err != nil {
		t.Error("Display Summary failed")
	}

	status, body, _ = userClient.SetBuyAmountRequest(userid, stockSymbol, lib.CentsToDollars(buyAmount))
	if status != lib.StatusOk {
		t.Error("Set Buy Amount failed\n" + strconv.Itoa(status) + body)
	}

	status, body, _ = userClient.SetBuyTriggerRequest(userid, stockSymbol, lib.CentsToDollars(triggerPrice))
	if status != lib.StatusOk {
		t.Error("Set Buy Trigger failed\n" + strconv.Itoa(status) + body)
	}

	summaryAfter, err := userClient.GetSummary(userid)
	if err != nil {
		t.Error("Display Summary failed")
	}

	if len(summaryAfter.Triggers) != len(summaryBefore.Triggers)+1 {
		t.Error("Trigger was not saved")
	}

	time.Sleep(65 * time.Second)

	summaryAfter, err = userClient.GetSummary(userid)
	if err != nil {
		t.Error("Display Summary failed")
	}

	if len(summaryAfter.Triggers) != len(summaryBefore.Triggers) {
		t.Error("Trigger was not cleared")
	}

	expectedStocksBought := (buyAmount / quoteValue)
	expectedStockCount := summaryBefore.Investments[0].Amount + expectedStocksBought
	if len(summaryAfter.Investments) > 0 && summaryAfter.Investments[0].Amount != expectedStockCount {
		t.Error("Trigger was not properly executed")
	}

	expectedBalance := summaryBefore.Cents - (expectedStocksBought * quoteValue)
	if summaryAfter.Cents != expectedBalance {
		t.Error("Money was not properly subtracted")
	}

}
