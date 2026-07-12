package controller

import (
	"testing"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/setting/operation_setting"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestEpayFractionalAmount(t *testing.T) {
	originalPrice := operation_setting.Price
	originalMinTopUp := operation_setting.MinTopUp
	originalQuotaDisplayType := operation_setting.GetGeneralSetting().QuotaDisplayType
	originalDiscounts := operation_setting.GetPaymentSetting().AmountDiscount

	t.Cleanup(func() {
		operation_setting.Price = originalPrice
		operation_setting.MinTopUp = originalMinTopUp
		operation_setting.GetGeneralSetting().QuotaDisplayType = originalQuotaDisplayType
		operation_setting.GetPaymentSetting().AmountDiscount = originalDiscounts
	})

	operation_setting.Price = 6.8
	operation_setting.MinTopUp = 0.147
	operation_setting.GetGeneralSetting().QuotaDisplayType = operation_setting.QuotaDisplayTypeUSD
	operation_setting.GetPaymentSetting().AmountDiscount = map[int]float64{}

	require.True(t, decimal.RequireFromString("0.1").LessThan(getMinTopup()))
	require.Equal(t, "1.00", getPayMoney(decimal.RequireFromString("0.147"), "default").StringFixed(2))
	require.Equal(t, "0.68", getPayMoney(decimal.RequireFromString("0.1"), "default").StringFixed(2))
	require.Equal(t, int64(50000), getEpayQuota(decimal.RequireFromString("0.1")))
	require.Equal(t, 0.1, getEpayDisplayAmount(decimal.RequireFromString("0.1")).InexactFloat64())
	require.Equal(t, int64(common.QuotaPerUnit*0.1), getEpayQuota(decimal.RequireFromString("0.1")))
}
