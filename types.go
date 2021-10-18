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

func c2goStr(p *C.char) string {
	return C.GoString(p)
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
	ret.TradingDay = c2goStr(&p.TradingDay[0])
	ret.LoginTime = c2goStr(&p.LoginTime[0])
	ret.BrokerID = c2goStr(&p.BrokerID[0])
	ret.UserID = c2goStr(&p.UserID[0])
	ret.SystemName = c2goStr(&p.SystemName[0])
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.MaxOrderRef = c2goStr(&p.MaxOrderRef[0])
	ret.SHFETime = c2goStr(&p.SHFETime[0])
	ret.DCETime = c2goStr(&p.DCETime[0])
	ret.CZCETime = c2goStr(&p.CZCETime[0])
	ret.FFEXTime = c2goStr(&p.FFEXTime[0])
	ret.INETime = c2goStr(&p.INETime[0])
	return ret
}

type CThostFtdcRspInfoField struct {
}

type CThostFtdcUserLogoutField struct {
}
type CThostFtdcMulticastInstrumentField struct{}
type CThostFtdcSpecificInstrumentField struct{}
type CThostFtdcDepthMarketDataField struct{}
type CThostFtdcForQuoteRspField struct{}
type CThostFtdcFensUserInfoField struct {
}

type CThostFtdcQryMulticastInstrumentField struct {
}
