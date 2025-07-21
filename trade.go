package payment_center_sdk

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/damon-dai/payment-center-sdk/proto"
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

// HnapayNewSelfAccount 新生个人开户
func (t *Trade) HnapayNewSelfAccount(request *proto.HnapayNewSelfAccountReq) (*proto.HnapayNewSelfAccountResp, error) {
	conn, err := GetConnect(t.Url)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewTradeClient(conn.Value())
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.HnapayNewSelfAccount(ctx, request)
}

// HnapayBindBankCard 新生个人用户绑卡
func (t *Trade) HnapayBindBankCard(request *proto.HnapayBindBankCardReq) (*proto.HnapayBindBankCardResp, error) {
	conn, err := GetConnect(t.Url)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewTradeClient(conn.Value())
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.HnapayBindBankCard(ctx, request)
}

// HnapayBindBankCardConfirm 新生个人用户绑卡确认
func (t *Trade) HnapayBindBankCardConfirm(request *proto.HnapayBindBankCardConfirmReq) (*proto.HnapayBindBankCardConfirmResp, error) {
	conn, err := GetConnect(t.Url)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewTradeClient(conn.Value())
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.HnapayBindBankCardConfirm(ctx, request)
}

// HnapayUnbindBankCard 新生个人用户解绑银行卡
func (t *Trade) HnapayUnbindBankCard(request *proto.HnapayUnbindBankCardReq) (*proto.HnapayUnbindBankCardResp, error) {
	conn, err := GetConnect(t.Url)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewTradeClient(conn.Value())
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.HnapayUnbindBankCard(ctx, request)
}

// HnapayQuickPayOrder 新生快捷支付下单：银行卡充值
func (t *Trade) HnapayQuickPayOrder(request *proto.HnapayQuickPayOrderReq) (*proto.HnapayQuickPayOrderResp, error) {
	conn, err := GetConnect(t.Url)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewTradeClient(conn.Value())
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.HnapayQuickPayOrder(ctx, request)
}

// HnapayQuickPayConfirm 新生快捷支付确认
func (t *Trade) HnapayQuickPayConfirm(request *proto.HnapayQuickPayConfirmReq) (*proto.HnapayQuickPayConfirmResp, error) {
	conn, err := GetConnect(t.Url)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewTradeClient(conn.Value())
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.HnapayQuickPayConfirm(ctx, request)
}

// HnapayTransfer 新生转账
func (t *Trade) HnapayTransfer(request *proto.HnapayTransferReq) (*proto.HnapayTransferResp, error) {
	conn, err := GetConnect(t.Url)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewTradeClient(conn.Value())
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.HnapayTransfer(ctx, request)
}

// HnapayWithdraw 新生提现
func (t *Trade) HnapayWithdraw(request *proto.HnapayWithdrawReq) (*proto.HnapayWithdrawResp, error) {
	conn, err := GetConnect(t.Url)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewTradeClient(conn.Value())
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.HnapayWithdraw(ctx, request)
}

// HnapayCheckUserAccountInfo 新生查询用户账户信息
func (t *Trade) HnapayCheckUserAccountInfo(request *proto.HnapayCheckUserAccountInfoReq) (*proto.HnapayCheckUserAccountInfoResp, error) {
	conn, err := GetConnect(t.Url)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewTradeClient(conn.Value())
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.HnapayCheckUserAccountInfo(ctx, request)
}

// GetNcountAccountInfo 获取系统账户信息
func (t *Trade) GetNcountAccountInfo(request *proto.GetNcountAccountInfoReq) (*proto.GetNcountAccountInfoResp, error) {
	conn, err := GetConnect(t.Url)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto.NewTradeClient(conn.Value())
	ctx := GetMetadataCtx(t.RequestId, t.BizCode, t.Token)
	return client.GetNcountAccountInfo(ctx, request)
}
