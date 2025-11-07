package payment_center_sdk

import (
	"fmt"
	"github.com/damon-dai/payment-center-sdk/proto/trade"
	"github.com/google/uuid"
)

type Trade struct {
	Url       string
	BizCode   string
	Token     string
	RequestId string
}

func Init(url, bizCode, requestId string) Trade {
	if requestId == "" {
		requestId = uuid.New().String()
	}

	requestId = fmt.Sprintf("%s_sdk_%s", bizCode, requestId)
	return Trade{
		Url:       url,
		BizCode:   bizCode,
		RequestId: requestId,
	}
}

// MallBookNewSelfAccount 个人开户
func (t *Trade) MallBookNewSelfAccount(request *trade.MallBookNewSelfAccountReq) (*trade.MallBookNewSelfAccountResp, error) {
	conn, err := NewGrpcClient(t.Url)
	if err != nil {
		return nil, err
	}

	client := trade.NewTradeClient(conn)
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.MallBookNewSelfAccount(ctx, request)
}

// MallBookBindBankCard 个人用户绑卡
func (t *Trade) MallBookBindBankCard(request *trade.MallBookBindBankCardReq) (*trade.MallBookBindBankCardResp, error) {
	conn, err := NewGrpcClient(t.Url)
	if err != nil {
		return nil, err
	}

	client := trade.NewTradeClient(conn)
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.MallBookBindBankCard(ctx, request)
}

// MallBookUnbindBankCard 个人用户解绑银行卡
func (t *Trade) MallBookUnbindBankCard(request *trade.MallBookUnbindBankCardReq) (*trade.MallBookUnbindBankCardResp, error) {
	conn, err := NewGrpcClient(t.Url)
	if err != nil {
		return nil, err
	}

	client := trade.NewTradeClient(conn)
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.MallBookUnbindBankCard(ctx, request)
}

// MallBookDeposit 充值
func (t *Trade) MallBookDeposit(request *trade.MallBookDepositReq) (*trade.MallBookDepositResp, error) {
	conn, err := NewGrpcClient(t.Url)
	if err != nil {
		return nil, err
	}

	client := trade.NewTradeClient(conn)
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.MallBookDeposit(ctx, request)
}

// MallBookTransfer 转账
func (t *Trade) MallBookTransfer(request *trade.MallBookTransferReq) (*trade.MallBookTransferResp, error) {
	conn, err := NewGrpcClient(t.Url)
	if err != nil {
		return nil, err
	}

	client := trade.NewTradeClient(conn)
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.MallBookTransfer(ctx, request)
}

// MallBookWithdraw 提现
func (t *Trade) MallBookWithdraw(request *trade.MallBookWithdrawReq) (*trade.MallBookWithdrawResp, error) {
	conn, err := NewGrpcClient(t.Url)
	if err != nil {
		return nil, err
	}

	client := trade.NewTradeClient(conn)
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.MallBookWithdraw(ctx, request)
}

// MallBookCheckUserAccountInfo 查询用户账户信息
func (t *Trade) MallBookCheckUserAccountInfo(request *trade.MallBookCheckUserAccountInfoReq) (*trade.MallBookCheckUserAccountInfoResp, error) {
	conn, err := NewGrpcClient(t.Url)
	if err != nil {
		return nil, err
	}

	client := trade.NewTradeClient(conn)
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.MallBookCheckUserAccountInfo(ctx, request)
}

// GetNcountAccountInfo 获取系统账户信息
func (t *Trade) GetNcountAccountInfo(request *trade.GetNcountAccountInfoReq) (*trade.GetNcountAccountInfoResp, error) {
	conn, err := NewGrpcClient(t.Url)
	if err != nil {
		return nil, err
	}

	client := trade.NewTradeClient(conn)
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.GetNcountAccountInfo(ctx, request)
}

func (t *Trade) SetPaymentSecret(request *trade.SetPaymentSecretReq) (*trade.SetPaymentSecretResp, error) {
	conn, err := NewGrpcClient(t.Url)
	if err != nil {
		return nil, err
	}

	client := trade.NewTradeClient(conn)
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.SetPaymentSecret(ctx, request)
}
