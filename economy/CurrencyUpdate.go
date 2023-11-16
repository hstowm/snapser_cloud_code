package economy

import (
	"errors"
	"golang.org/x/net/context"
	"snapser/cloudecode/internal_connector"
	proto2 "snapser/cloudecode/snapserpb/inventory"
)

func GetUserCurrencies(userID string, c context.Context, pgs *internal_connector.SnapserServiceConnector) (map[string]int32, error) {
	currencies, err := pgs.InventoryClient.GetUserCurrencies(c, &proto2.GetUserCurrenciesRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	return currencies.Currencies, nil
}

func UpdateCurrencies(amountUpdate map[string]int32, userID string, c context.Context, pgs *internal_connector.SnapserServiceConnector) (map[string]int32, error) {

	currentCurrencies, err := GetUserCurrencies(userID, c, pgs)
	if err != nil {
		return nil, err
	}
	for currencyKey, value := range amountUpdate {
		if value+currentCurrencies[currencyKey] < 0 {
			return currentCurrencies, errors.New("not enough money")
		}
	}
	for currencyKey, value := range amountUpdate {
		if value == 0 {
			continue
		}
		_, err := pgs.InventoryClient.UpdateUserVirtualCurrency(c, &proto2.UpdateUserVirtualCurrencyRequest{
			Amount:       value,
			CurrencyName: currencyKey,
			UserId:       userID,
		})
		if err != nil {
			return currentCurrencies, errors.New("not enough money")
		}
		currentCurrencies[currencyKey] += value

	}
	return currentCurrencies, nil
}
