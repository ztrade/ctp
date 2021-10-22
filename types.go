package ctp

/*
 #include <string.h>
 #include "include/ThostFtdcUserApiStruct.h"
 #include "ctp.h"
 int go2cStr(_GoString_ str, char* p, int l){
    size_t n = _GoStringLen(str);
    if (n == 0){
      return 1;
    }
    if (n+1 > l){
       return 2;
    }
    const char *src = _GoStringPtr(str);
    memcpy(p, src, n);
    p[n] = 0;
    return 0;
 }
*/
import "C"
import "errors"

func c2goStr(p *C.char, n int) string {
	return C.GoStringN(p, C.int(n))
}

func go2cStr(str string, p *C.char, l int) (err error) {
	ret := C.go2cStr(str, p, C.int(l))
	if ret == 1 {
		err = errors.New("str empty")
	} else if ret == 2 {
		err = errors.New("str too long")
	}
	return
}

func go2cStrArray(strs []string) (ret []*C.char) {
	for _, v := range strs {
		ret = append(ret, C.CString(v))
	}
	return ret
}

type CThostFtdcReqUserLoginField struct {
	///交易日
	TradingDay string
	///经纪公司代码
	BrokerID string
	///用户代码
	UserID string
	///密码
	Password string
	///用户端产品信息
	UserProductInfo string
	///接口端产品信息
	InterfaceProductInfo string
	///协议信息
	ProtocolInfo string
	///Mac地址
	MacAddress string
	///动态密码
	OneTimePassword string
	///保留的无效字段
	Reserve1 string
	///登录备注
	LoginRemark string
	///终端IP端口
	ClientIPPort int
	///终端IP地址
	ClientIPAddress string
}

func (l *CThostFtdcReqUserLoginField) CValue() *C.CThostFtdcReqUserLoginField {
	ptr := &C.CThostFtdcReqUserLoginField{}

	go2cStr(l.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(l.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(l.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(l.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(l.UserProductInfo, &ptr.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(l.InterfaceProductInfo, &ptr.InterfaceProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(l.ProtocolInfo, &ptr.ProtocolInfo[0], C.sizeof_TThostFtdcProtocolInfoType)
	go2cStr(l.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(l.OneTimePassword, &ptr.OneTimePassword[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(l.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(l.LoginRemark, &ptr.LoginRemark[0], C.sizeof_TThostFtdcLoginRemarkType)
	ptr.ClientIPPort = C.int(l.ClientIPPort)
	go2cStr(l.ClientIPAddress, &ptr.ClientIPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcRspUserLoginField struct {
	TradingDay string
	///登录成功时间
	LoginTime string
	///经纪公司代码
	BrokerID string
	///用户代码
	UserID string
	///交易系统名称
	SystemName string
	///前置编号
	FrontID int
	///会话编号
	SessionID int
	///最大报单引用
	MaxOrderRef string
	///上期所时间
	SHFETime string
	///大商所时间
	DCETime string
	///郑商所时间
	CZCETime string
	///中金所时间
	FFEXTime string
	///能源中心时间
	INETime string
}

func NewCThostFtdcRspUserLoginField(p C.CThostFtdcRspUserLoginField) *CThostFtdcRspUserLoginField {
	ret := &CThostFtdcRspUserLoginField{}
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.LoginTime = c2goStr(&p.LoginTime[0], C.sizeof_TThostFtdcTimeType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.SystemName = c2goStr(&p.SystemName[0], C.sizeof_TThostFtdcSystemNameType)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.MaxOrderRef = c2goStr(&p.MaxOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.SHFETime = c2goStr(&p.SHFETime[0], C.sizeof_TThostFtdcTimeType)
	ret.DCETime = c2goStr(&p.DCETime[0], C.sizeof_TThostFtdcTimeType)
	ret.CZCETime = c2goStr(&p.CZCETime[0], C.sizeof_TThostFtdcTimeType)
	ret.FFEXTime = c2goStr(&p.FFEXTime[0], C.sizeof_TThostFtdcTimeType)
	ret.INETime = c2goStr(&p.INETime[0], C.sizeof_TThostFtdcTimeType)
	return ret
}

type CThostFtdcRspInfoField struct {
	ErrorID int
	///错误信息
	ErrorMsg string
}

func NewCThostFtdcRspInfoField(p *C.CThostFtdcRspInfoField) *CThostFtdcRspInfoField {
	ret := &CThostFtdcRspInfoField{
		ErrorID:  int(p.ErrorID),
		ErrorMsg: c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType),
	}
	return ret
}

type CThostFtdcUserLogoutField struct {
	///经纪公司代码
	BrokerID string
	///用户代码
	UserID string
}

func NewCThostFtdcUserLogoutField(p *C.CThostFtdcUserLogoutField) *CThostFtdcUserLogoutField {
	ret := &CThostFtdcUserLogoutField{
		BrokerID: c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType),
		UserID:   c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType),
	}
	return ret
}

func (l *CThostFtdcUserLogoutField) CValue() *C.CThostFtdcUserLogoutField {
	ptr := &C.CThostFtdcUserLogoutField{}
	go2cStr(l.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(l.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ptr
}

type CThostFtdcMulticastInstrumentField struct {
	///主题号
	TopicID int
	///保留的无效字段
	Reserve1 string
	///合约编号
	InstrumentNo int
	///基准价
	CodePrice float64
	///合约数量乘数
	VolumeMultiple int
	///最小变动价位
	PriceTick float64
	///合约代码
	InstrumentID string
}

func NewCThostFtdcMulticastInstrumentField(p *C.CThostFtdcMulticastInstrumentField) *CThostFtdcMulticastInstrumentField {
	ret := &CThostFtdcMulticastInstrumentField{
		TopicID:        int(p.TopicID),
		Reserve1:       c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType),
		InstrumentNo:   int(p.InstrumentNo),
		CodePrice:      float64(p.CodePrice),
		VolumeMultiple: int(p.VolumeMultiple),
		PriceTick:      float64(p.PriceTick),
		InstrumentID:   c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDsType),
	}
	return ret
}

type CThostFtdcSpecificInstrumentField struct {
	///保留的无效字段
	Reserve1 string
	///合约代码
	InstrumentID string
}

func NewCThostFtdcSpecificInstrumentField(p *C.CThostFtdcSpecificInstrumentField) *CThostFtdcSpecificInstrumentField {
	ret := &CThostFtdcSpecificInstrumentField{
		Reserve1:     c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType),
		InstrumentID: c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType),
	}
	return ret
}

type CThostFtdcDepthMarketDataField struct {
	TradingDay string
	///保留的无效字段
	Reserve1 string
	///交易所代码
	ExchangeID string
	///保留的无效字段
	Reserve2 string
	///最新价
	LastPrice float64
	///上次结算价
	PreSettlementPrice float64
	///昨收盘
	PreClosePrice float64
	///昨持仓量
	PreOpenInterest float64
	///今开盘
	OpenPrice float64
	///最高价
	HighestPrice float64
	///最低价
	LowestPrice float64
	///数量
	Volume int
	///成交金额
	Turnover float64
	///持仓量
	OpenInterest float64
	///今收盘
	ClosePrice float64
	///本次结算价
	SettlementPrice float64
	///涨停板价
	UpperLimitPrice float64
	///跌停板价
	LowerLimitPrice float64
	///昨虚实度
	PreDelta float64
	///今虚实度
	CurrDelta float64
	///最后修改时间
	UpdateTime string
	///最后修改毫秒
	UpdateMillisec int
	///申买价一
	BidPrice1 float64
	///申买量一
	BidVolume1 int
	///申卖价一
	AskPrice1 float64
	///申卖量一
	AskVolume1 int
	///申买价二
	BidPrice2 float64
	///申买量二
	BidVolume2 int
	///申卖价二
	AskPrice2 float64
	///申卖量二
	AskVolume2 int
	///申买价三
	BidPrice3 float64
	///申买量三
	BidVolume3 int
	///申卖价三
	AskPrice3 float64
	///申卖量三
	AskVolume3 int
	///申买价四
	BidPrice4 float64
	///申买量四
	BidVolume4 int
	///申卖价四
	AskPrice4 float64
	///申卖量四
	AskVolume4 int
	///申买价五
	BidPrice5 float64
	///申买量五
	BidVolume5 int
	///申卖价五
	AskPrice5 float64
	///申卖量五
	AskVolume5 int
	///当日均价
	AveragePrice float64
	///业务日期
	ActionDay string
	///合约代码
	InstrumentID string
	///合约在交易所的代码
	ExchangeInstID string
	///上带价
	BandingUpperPrice float64
	///下带价
	BandingLowerPrice float64
}

func NewCThostFtdcDepthMarketDataField(p *C.CThostFtdcDepthMarketDataField) *CThostFtdcDepthMarketDataField {
	if p == nil {
		return nil
	}
	ret := &CThostFtdcDepthMarketDataField{
		TradingDay: c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType),

		Reserve1: c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType),
		///交易所代码
		ExchangeID: c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType),
		///保留的无效字段
		Reserve2: c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType),
		///最新价
		LastPrice: float64(p.LastPrice),
		///上次结算价
		PreSettlementPrice: float64(p.PreSettlementPrice),
		///昨收盘
		PreClosePrice: float64(p.PreClosePrice),
		///昨持仓量
		PreOpenInterest: float64(p.PreOpenInterest),
		///今开盘
		OpenPrice: float64(p.OpenPrice),
		///最高价
		HighestPrice: float64(p.HighestPrice),
		///最低价
		LowestPrice: float64(p.LowestPrice),
		///数量
		Volume:            int(p.Volume),
		Turnover:          float64(p.Turnover),
		OpenInterest:      float64(p.OpenInterest),
		ClosePrice:        float64(p.ClosePrice),
		SettlementPrice:   float64(p.SettlementPrice),
		UpperLimitPrice:   float64(p.UpperLimitPrice),
		LowerLimitPrice:   float64(p.LowerLimitPrice),
		PreDelta:          float64(p.PreDelta),
		CurrDelta:         float64(p.CurrDelta),
		UpdateTime:        c2goStr(&p.UpdateTime[0], C.sizeof_TThostFtdcTimeType),
		UpdateMillisec:    int(p.UpdateMillisec),
		BidPrice1:         float64(p.BidPrice1),
		BidVolume1:        int(p.BidVolume1),
		AskPrice1:         float64(p.AskPrice1),
		AskVolume1:        int(p.AskVolume1),
		BidPrice2:         float64(p.BidPrice2),
		BidVolume2:        int(p.BidVolume2),
		AskPrice2:         float64(p.AskPrice2),
		AskVolume2:        int(p.AskVolume2),
		BidPrice3:         float64(p.BidPrice3),
		BidVolume3:        int(p.BidVolume3),
		AskPrice3:         float64(p.AskPrice3),
		AskVolume3:        int(p.AskVolume3),
		BidPrice4:         float64(p.BidPrice4),
		BidVolume4:        int(p.BidVolume4),
		AskPrice4:         float64(p.AskPrice4),
		AskVolume4:        int(p.AskVolume4),
		BidPrice5:         float64(p.BidPrice5),
		BidVolume5:        int(p.BidVolume5),
		AskPrice5:         float64(p.AskPrice5),
		AskVolume5:        int(p.AskVolume5),
		AveragePrice:      float64(p.AveragePrice),
		ActionDay:         c2goStr(&p.ActionDay[0], C.sizeof_TThostFtdcDateType),
		InstrumentID:      c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType),
		ExchangeInstID:    c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType),
		BandingUpperPrice: float64(p.BandingUpperPrice),
		BandingLowerPrice: float64(p.BandingLowerPrice),
	}
	if ret.BidVolume2 == 0 {
		ret.BidPrice2 = 0
	}
	if ret.BidVolume3 == 0 {
		ret.BidPrice3 = 0
	}
	if ret.BidVolume4 == 0 {
		ret.BidPrice4 = 0
	}
	if ret.BidVolume5 == 0 {
		ret.BidPrice5 = 0
	}
	if ret.AskVolume2 == 0 {
		ret.AskPrice2 = 0
	}
	if ret.AskVolume3 == 0 {
		ret.AskPrice3 = 0
	}
	if ret.AskVolume4 == 0 {
		ret.AskPrice4 = 0
	}
	if ret.AskVolume5 == 0 {
		ret.AskPrice5 = 0
	}

	return ret
}

type CThostFtdcForQuoteRspField struct {
	///交易日
	TradingDay string
	///保留的无效字段
	Reserve1 string
	///询价编号
	ForQuoteSysID string
	///询价时间
	ForQuoteTime string
	///业务日期
	ActionDay string
	///交易所代码
	ExchangeID string
	///合约代码
	InstrumentID string
}

func NewCThostFtdcForQuoteRspField(p *C.CThostFtdcForQuoteRspField) *CThostFtdcForQuoteRspField {
	ret := &CThostFtdcForQuoteRspField{
		TradingDay:    c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType),
		Reserve1:      c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType),
		ForQuoteSysID: c2goStr(&p.ForQuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType),
		ForQuoteTime:  c2goStr(&p.ForQuoteTime[0], C.sizeof_TThostFtdcTimeType),
		ActionDay:     c2goStr(&p.ActionDay[0], C.sizeof_TThostFtdcDateType),
		ExchangeID:    c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType),
		InstrumentID:  c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType),
	}
	return ret
}

type CThostFtdcFensUserInfoField struct {
	///经纪公司代码
	BrokerID string
	///用户代码
	UserID string
	///登录模式
	LoginMode byte
}

func (l *CThostFtdcFensUserInfoField) CValue() *C.CThostFtdcFensUserInfoField {
	ptr := &C.CThostFtdcFensUserInfoField{}
	go2cStr(l.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(l.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.LoginMode = C.char(l.LoginMode)
	return ptr
}

type CThostFtdcQryMulticastInstrumentField struct {
	///主题号
	TopicID int
	///保留的无效字段
	Reserve1 string
	///合约代码
	InstrumentID string
}

func (l *CThostFtdcQryMulticastInstrumentField) CValue() *C.CThostFtdcQryMulticastInstrumentField {
	ptr := &C.CThostFtdcQryMulticastInstrumentField{}
	ptr.TopicID = C.int(l.TopicID)
	go2cStr(l.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(l.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}
