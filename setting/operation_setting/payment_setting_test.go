package operation_setting

import (
	"math"
	"testing"
)

func TestGetTopupAmountOptionsUsesDisplayCurrency(t *testing.T) {
	originalGeneralSetting := generalSetting
	originalPaymentSetting := paymentSetting
	originalUSDExchangeRate := USDExchangeRate
	defer func() {
		generalSetting = originalGeneralSetting
		paymentSetting = originalPaymentSetting
		USDExchangeRate = originalUSDExchangeRate
	}()

	paymentSetting.AmountOptions = []int{10, 20}
	paymentSetting.CNYAmountOptions = []int{10, 20}
	paymentSetting.CNYMinTopUp = 10
	USDExchangeRate = 6.8

	generalSetting.QuotaDisplayType = QuotaDisplayTypeUSD
	if got := GetTopupAmountOptions(); len(got) != 2 || got[0] != 10 || got[1] != 20 {
		t.Fatalf("USD preset amounts = %v, want [10 20]", got)
	}

	generalSetting.QuotaDisplayType = QuotaDisplayTypeCNY
	got := GetTopupAmountOptions()
	if len(got) != 2 || math.Abs(got[0]-10.0/6.8) > 1e-12 || math.Abs(got[1]-20.0/6.8) > 1e-12 {
		t.Fatalf("CNY preset amounts = %v, want converted internal USD amounts", got)
	}
	if got := GetTopupMinAmount(); math.Abs(got-10.0/6.8) > 1e-12 {
		t.Fatalf("CNY minimum amount = %v, want converted internal USD amount", got)
	}
	if got := GetTopupDisplayMinAmount(); got != 10 {
		t.Fatalf("CNY display minimum amount = %v, want 10", got)
	}
}
