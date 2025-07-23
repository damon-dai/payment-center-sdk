/**
 * @Author: Damon
 * @Date: 2025/7/23 11:19
 * @Description:
 **/

package payment_center_sdk

import (
	"github.com/damon-dai/payment-center-sdk/proto/trade"
	"reflect"
	"testing"
)

const (
	//url = "127.0.0.1:9528"
	url       = "tarsproxy.xingdong.sh.cn:9081"
	bizCode   = "B1003"
	requestId = ""
)

func TestTrade_GetNcountAccountInfo(t1 *testing.T) {
	type fields struct {
		Url       string
		BizCode   string
		Token     string
		RequestId string
	}
	type args struct {
		request *trade.GetNcountAccountInfoReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *trade.GetNcountAccountInfoResp
		wantErr bool
	}{
		{
			name: "埃克",
			fields: fields{
				Url:     url,
				BizCode: bizCode,
				Token:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjE1MDE2MSwiZXhwIjoxNzUzNDkzMTUyfQ.OUh7HtxI4bUxXZH7w75XntdSR0QyXlM7--VSo-7n6vw",
			},
			args: args{
				&trade.GetNcountAccountInfoReq{
					BizCode: bizCode,
				},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trade{
				Url:       tt.fields.Url,
				BizCode:   tt.fields.BizCode,
				Token:     tt.fields.Token,
				RequestId: tt.fields.RequestId,
			}
			got, err := t.GetNcountAccountInfo(tt.args.request)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetNcountAccountInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetNcountAccountInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrade_HnapayBindBankCard(t1 *testing.T) {
	type fields struct {
		Url       string
		BizCode   string
		Token     string
		RequestId string
	}
	type args struct {
		request *trade.HnapayBindBankCardReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *trade.HnapayBindBankCardResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trade{
				Url:       tt.fields.Url,
				BizCode:   tt.fields.BizCode,
				Token:     tt.fields.Token,
				RequestId: tt.fields.RequestId,
			}
			got, err := t.HnapayBindBankCard(tt.args.request)
			if (err != nil) != tt.wantErr {
				t1.Errorf("HnapayBindBankCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("HnapayBindBankCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrade_HnapayBindBankCardConfirm(t1 *testing.T) {
	type fields struct {
		Url       string
		BizCode   string
		Token     string
		RequestId string
	}
	type args struct {
		request *trade.HnapayBindBankCardConfirmReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *trade.HnapayBindBankCardConfirmResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trade{
				Url:       tt.fields.Url,
				BizCode:   tt.fields.BizCode,
				Token:     tt.fields.Token,
				RequestId: tt.fields.RequestId,
			}
			got, err := t.HnapayBindBankCardConfirm(tt.args.request)
			if (err != nil) != tt.wantErr {
				t1.Errorf("HnapayBindBankCardConfirm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("HnapayBindBankCardConfirm() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrade_HnapayCheckUserAccountInfo(t1 *testing.T) {
	type fields struct {
		Url       string
		BizCode   string
		Token     string
		RequestId string
	}
	type args struct {
		request *trade.HnapayCheckUserAccountInfoReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *trade.HnapayCheckUserAccountInfoResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trade{
				Url:       tt.fields.Url,
				BizCode:   tt.fields.BizCode,
				Token:     tt.fields.Token,
				RequestId: tt.fields.RequestId,
			}
			got, err := t.HnapayCheckUserAccountInfo(tt.args.request)
			if (err != nil) != tt.wantErr {
				t1.Errorf("HnapayCheckUserAccountInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("HnapayCheckUserAccountInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrade_HnapayNewSelfAccount(t1 *testing.T) {
	type fields struct {
		Url       string
		BizCode   string
		Token     string
		RequestId string
	}
	type args struct {
		request *trade.HnapayNewSelfAccountReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *trade.HnapayNewSelfAccountResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trade{
				Url:       tt.fields.Url,
				BizCode:   tt.fields.BizCode,
				Token:     tt.fields.Token,
				RequestId: tt.fields.RequestId,
			}
			got, err := t.HnapayNewSelfAccount(tt.args.request)
			if (err != nil) != tt.wantErr {
				t1.Errorf("HnapayNewSelfAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("HnapayNewSelfAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrade_HnapayQuickPayConfirm(t1 *testing.T) {
	type fields struct {
		Url       string
		BizCode   string
		Token     string
		RequestId string
	}
	type args struct {
		request *trade.HnapayQuickPayConfirmReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *trade.HnapayQuickPayConfirmResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trade{
				Url:       tt.fields.Url,
				BizCode:   tt.fields.BizCode,
				Token:     tt.fields.Token,
				RequestId: tt.fields.RequestId,
			}
			got, err := t.HnapayQuickPayConfirm(tt.args.request)
			if (err != nil) != tt.wantErr {
				t1.Errorf("HnapayQuickPayConfirm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("HnapayQuickPayConfirm() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrade_HnapayQuickPayOrder(t1 *testing.T) {
	type fields struct {
		Url       string
		BizCode   string
		Token     string
		RequestId string
	}
	type args struct {
		request *trade.HnapayQuickPayOrderReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *trade.HnapayQuickPayOrderResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trade{
				Url:       tt.fields.Url,
				BizCode:   tt.fields.BizCode,
				Token:     tt.fields.Token,
				RequestId: tt.fields.RequestId,
			}
			got, err := t.HnapayQuickPayOrder(tt.args.request)
			if (err != nil) != tt.wantErr {
				t1.Errorf("HnapayQuickPayOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("HnapayQuickPayOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrade_HnapayTransfer(t1 *testing.T) {
	type fields struct {
		Url       string
		BizCode   string
		Token     string
		RequestId string
	}
	type args struct {
		request *trade.HnapayTransferReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *trade.HnapayTransferResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trade{
				Url:       tt.fields.Url,
				BizCode:   tt.fields.BizCode,
				Token:     tt.fields.Token,
				RequestId: tt.fields.RequestId,
			}
			got, err := t.HnapayTransfer(tt.args.request)
			if (err != nil) != tt.wantErr {
				t1.Errorf("HnapayTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("HnapayTransfer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrade_HnapayUnbindBankCard(t1 *testing.T) {
	type fields struct {
		Url       string
		BizCode   string
		Token     string
		RequestId string
	}
	type args struct {
		request *trade.HnapayUnbindBankCardReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *trade.HnapayUnbindBankCardResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trade{
				Url:       tt.fields.Url,
				BizCode:   tt.fields.BizCode,
				Token:     tt.fields.Token,
				RequestId: tt.fields.RequestId,
			}
			got, err := t.HnapayUnbindBankCard(tt.args.request)
			if (err != nil) != tt.wantErr {
				t1.Errorf("HnapayUnbindBankCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("HnapayUnbindBankCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrade_HnapayWithdraw(t1 *testing.T) {
	type fields struct {
		Url       string
		BizCode   string
		Token     string
		RequestId string
	}
	type args struct {
		request *trade.HnapayWithdrawReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *trade.HnapayWithdrawResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trade{
				Url:       tt.fields.Url,
				BizCode:   tt.fields.BizCode,
				Token:     tt.fields.Token,
				RequestId: tt.fields.RequestId,
			}
			got, err := t.HnapayWithdraw(tt.args.request)
			if (err != nil) != tt.wantErr {
				t1.Errorf("HnapayWithdraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("HnapayWithdraw() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrade_SetPaymentSecret(t1 *testing.T) {
	type fields struct {
		Url       string
		BizCode   string
		Token     string
		RequestId string
	}
	type args struct {
		request *trade.SetPaymentSecretReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *trade.SetPaymentSecretResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trade{
				Url:       tt.fields.Url,
				BizCode:   tt.fields.BizCode,
				Token:     tt.fields.Token,
				RequestId: tt.fields.RequestId,
			}
			got, err := t.SetPaymentSecret(tt.args.request)
			if (err != nil) != tt.wantErr {
				t1.Errorf("SetPaymentSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("SetPaymentSecret() got = %v, want %v", got, tt.want)
			}
		})
	}
}
