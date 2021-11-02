package ctp

/*
#include "types_gen.h"

*/
import "C"

type CThostFtdcDisseminationField struct {
	SequenceSeries int16
	SequenceNo     int
}

func NewCThostFtdcDisseminationField(p *C.CThostFtdcDisseminationField) *CThostFtdcDisseminationField {
	ret := new(CThostFtdcDisseminationField)
	ret.SequenceSeries = int16(p.SequenceSeries)
	ret.SequenceNo = int(p.SequenceNo)
	return ret
}
func (s *CThostFtdcDisseminationField) CValue() *C.CThostFtdcDisseminationField {
	ptr := C.newCThostFtdcDisseminationField()
	ptr.SequenceSeries = C.short(s.SequenceSeries)
	ptr.SequenceNo = C.int(s.SequenceNo)
	return ptr
}

type CThostFtdcReqUserLoginField struct {
	TradingDay           string
	BrokerID             string
	UserID               string
	Password             string
	UserProductInfo      string
	InterfaceProductInfo string
	ProtocolInfo         string
	MacAddress           string
	OneTimePassword      string
	Reserve1             string
	LoginRemark          string
	ClientIPPort         int
	ClientIPAddress      string
}

func NewCThostFtdcReqUserLoginField(p *C.CThostFtdcReqUserLoginField) *CThostFtdcReqUserLoginField {
	ret := new(CThostFtdcReqUserLoginField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.UserProductInfo = c2goStr(&p.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.InterfaceProductInfo = c2goStr(&p.InterfaceProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.ProtocolInfo = c2goStr(&p.ProtocolInfo[0], C.sizeof_TThostFtdcProtocolInfoType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.OneTimePassword = c2goStr(&p.OneTimePassword[0], C.sizeof_TThostFtdcPasswordType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.LoginRemark = c2goStr(&p.LoginRemark[0], C.sizeof_TThostFtdcLoginRemarkType)
	ret.ClientIPPort = int(p.ClientIPPort)
	ret.ClientIPAddress = c2goStr(&p.ClientIPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcReqUserLoginField) CValue() *C.CThostFtdcReqUserLoginField {
	ptr := C.newCThostFtdcReqUserLoginField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.UserProductInfo, &ptr.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.InterfaceProductInfo, &ptr.InterfaceProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.ProtocolInfo, &ptr.ProtocolInfo[0], C.sizeof_TThostFtdcProtocolInfoType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.OneTimePassword, &ptr.OneTimePassword[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.LoginRemark, &ptr.LoginRemark[0], C.sizeof_TThostFtdcLoginRemarkType)
	ptr.ClientIPPort = C.int(s.ClientIPPort)
	go2cStr(s.ClientIPAddress, &ptr.ClientIPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcRspUserLoginField struct {
	TradingDay  string
	LoginTime   string
	BrokerID    string
	UserID      string
	SystemName  string
	FrontID     int
	SessionID   int
	MaxOrderRef string
	SHFETime    string
	DCETime     string
	CZCETime    string
	FFEXTime    string
	INETime     string
}

func NewCThostFtdcRspUserLoginField(p *C.CThostFtdcRspUserLoginField) *CThostFtdcRspUserLoginField {
	ret := new(CThostFtdcRspUserLoginField)
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
func (s *CThostFtdcRspUserLoginField) CValue() *C.CThostFtdcRspUserLoginField {
	ptr := C.newCThostFtdcRspUserLoginField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.LoginTime, &ptr.LoginTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.SystemName, &ptr.SystemName[0], C.sizeof_TThostFtdcSystemNameType)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.MaxOrderRef, &ptr.MaxOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.SHFETime, &ptr.SHFETime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.DCETime, &ptr.DCETime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.CZCETime, &ptr.CZCETime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.FFEXTime, &ptr.FFEXTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.INETime, &ptr.INETime[0], C.sizeof_TThostFtdcTimeType)
	return ptr
}

type CThostFtdcUserLogoutField struct {
	BrokerID string
	UserID   string
}

func NewCThostFtdcUserLogoutField(p *C.CThostFtdcUserLogoutField) *CThostFtdcUserLogoutField {
	ret := new(CThostFtdcUserLogoutField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ret
}
func (s *CThostFtdcUserLogoutField) CValue() *C.CThostFtdcUserLogoutField {
	ptr := C.newCThostFtdcUserLogoutField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ptr
}

type CThostFtdcForceUserLogoutField struct {
	BrokerID string
	UserID   string
}

func NewCThostFtdcForceUserLogoutField(p *C.CThostFtdcForceUserLogoutField) *CThostFtdcForceUserLogoutField {
	ret := new(CThostFtdcForceUserLogoutField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ret
}
func (s *CThostFtdcForceUserLogoutField) CValue() *C.CThostFtdcForceUserLogoutField {
	ptr := C.newCThostFtdcForceUserLogoutField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ptr
}

type CThostFtdcReqAuthenticateField struct {
	BrokerID        string
	UserID          string
	UserProductInfo string
	AuthCode        string
	AppID           string
}

func NewCThostFtdcReqAuthenticateField(p *C.CThostFtdcReqAuthenticateField) *CThostFtdcReqAuthenticateField {
	ret := new(CThostFtdcReqAuthenticateField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.UserProductInfo = c2goStr(&p.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.AuthCode = c2goStr(&p.AuthCode[0], C.sizeof_TThostFtdcAuthCodeType)
	ret.AppID = c2goStr(&p.AppID[0], C.sizeof_TThostFtdcAppIDType)
	return ret
}
func (s *CThostFtdcReqAuthenticateField) CValue() *C.CThostFtdcReqAuthenticateField {
	ptr := C.newCThostFtdcReqAuthenticateField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.UserProductInfo, &ptr.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.AuthCode, &ptr.AuthCode[0], C.sizeof_TThostFtdcAuthCodeType)
	go2cStr(s.AppID, &ptr.AppID[0], C.sizeof_TThostFtdcAppIDType)
	return ptr
}

type CThostFtdcRspAuthenticateField struct {
	BrokerID        string
	UserID          string
	UserProductInfo string
	AppID           string
	AppType         byte
}

func NewCThostFtdcRspAuthenticateField(p *C.CThostFtdcRspAuthenticateField) *CThostFtdcRspAuthenticateField {
	ret := new(CThostFtdcRspAuthenticateField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.UserProductInfo = c2goStr(&p.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.AppID = c2goStr(&p.AppID[0], C.sizeof_TThostFtdcAppIDType)
	ret.AppType = byte(p.AppType)
	return ret
}
func (s *CThostFtdcRspAuthenticateField) CValue() *C.CThostFtdcRspAuthenticateField {
	ptr := C.newCThostFtdcRspAuthenticateField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.UserProductInfo, &ptr.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.AppID, &ptr.AppID[0], C.sizeof_TThostFtdcAppIDType)
	ptr.AppType = C.char(s.AppType)
	return ptr
}

type CThostFtdcAuthenticationInfoField struct {
	BrokerID        string
	UserID          string
	UserProductInfo string
	AuthInfo        string
	IsResult        int
	AppID           string
	AppType         byte
	Reserve1        string
	ClientIPAddress string
}

func NewCThostFtdcAuthenticationInfoField(p *C.CThostFtdcAuthenticationInfoField) *CThostFtdcAuthenticationInfoField {
	ret := new(CThostFtdcAuthenticationInfoField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.UserProductInfo = c2goStr(&p.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.AuthInfo = c2goStr(&p.AuthInfo[0], C.sizeof_TThostFtdcAuthInfoType)
	ret.IsResult = int(p.IsResult)
	ret.AppID = c2goStr(&p.AppID[0], C.sizeof_TThostFtdcAppIDType)
	ret.AppType = byte(p.AppType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.ClientIPAddress = c2goStr(&p.ClientIPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcAuthenticationInfoField) CValue() *C.CThostFtdcAuthenticationInfoField {
	ptr := C.newCThostFtdcAuthenticationInfoField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.UserProductInfo, &ptr.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.AuthInfo, &ptr.AuthInfo[0], C.sizeof_TThostFtdcAuthInfoType)
	ptr.IsResult = C.int(s.IsResult)
	go2cStr(s.AppID, &ptr.AppID[0], C.sizeof_TThostFtdcAppIDType)
	ptr.AppType = C.char(s.AppType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.ClientIPAddress, &ptr.ClientIPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcRspUserLogin2Field struct {
	TradingDay   string
	LoginTime    string
	BrokerID     string
	UserID       string
	SystemName   string
	FrontID      int
	SessionID    int
	MaxOrderRef  string
	SHFETime     string
	DCETime      string
	CZCETime     string
	FFEXTime     string
	INETime      string
	RandomString string
}

func NewCThostFtdcRspUserLogin2Field(p *C.CThostFtdcRspUserLogin2Field) *CThostFtdcRspUserLogin2Field {
	ret := new(CThostFtdcRspUserLogin2Field)
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
	ret.RandomString = c2goStr(&p.RandomString[0], C.sizeof_TThostFtdcRandomStringType)
	return ret
}
func (s *CThostFtdcRspUserLogin2Field) CValue() *C.CThostFtdcRspUserLogin2Field {
	ptr := C.newCThostFtdcRspUserLogin2Field()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.LoginTime, &ptr.LoginTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.SystemName, &ptr.SystemName[0], C.sizeof_TThostFtdcSystemNameType)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.MaxOrderRef, &ptr.MaxOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.SHFETime, &ptr.SHFETime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.DCETime, &ptr.DCETime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.CZCETime, &ptr.CZCETime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.FFEXTime, &ptr.FFEXTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.INETime, &ptr.INETime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.RandomString, &ptr.RandomString[0], C.sizeof_TThostFtdcRandomStringType)
	return ptr
}

type CThostFtdcTransferHeaderField struct {
	Version     string
	TradeCode   string
	TradeDate   string
	TradeTime   string
	TradeSerial string
	FutureID    string
	BankID      string
	BankBrchID  string
	OperNo      string
	DeviceID    string
	RecordNum   string
	SessionID   int
	RequestID   int
}

func NewCThostFtdcTransferHeaderField(p *C.CThostFtdcTransferHeaderField) *CThostFtdcTransferHeaderField {
	ret := new(CThostFtdcTransferHeaderField)
	ret.Version = c2goStr(&p.Version[0], C.sizeof_TThostFtdcVersionType)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.TradeSerial = c2goStr(&p.TradeSerial[0], C.sizeof_TThostFtdcTradeSerialType)
	ret.FutureID = c2goStr(&p.FutureID[0], C.sizeof_TThostFtdcFutureIDType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBrchID = c2goStr(&p.BankBrchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.RecordNum = c2goStr(&p.RecordNum[0], C.sizeof_TThostFtdcRecordNumType)
	ret.SessionID = int(p.SessionID)
	ret.RequestID = int(p.RequestID)
	return ret
}
func (s *CThostFtdcTransferHeaderField) CValue() *C.CThostFtdcTransferHeaderField {
	ptr := C.newCThostFtdcTransferHeaderField()
	go2cStr(s.Version, &ptr.Version[0], C.sizeof_TThostFtdcVersionType)
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.TradeSerial, &ptr.TradeSerial[0], C.sizeof_TThostFtdcTradeSerialType)
	go2cStr(s.FutureID, &ptr.FutureID[0], C.sizeof_TThostFtdcFutureIDType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBrchID, &ptr.BankBrchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	go2cStr(s.RecordNum, &ptr.RecordNum[0], C.sizeof_TThostFtdcRecordNumType)
	ptr.SessionID = C.int(s.SessionID)
	ptr.RequestID = C.int(s.RequestID)
	return ptr
}

type CThostFtdcTransferBankToFutureReqField struct {
	FutureAccount string
	FuturePwdFlag byte
	FutureAccPwd  string
	TradeAmt      float64
	CustFee       float64
	CurrencyCode  string
}

func NewCThostFtdcTransferBankToFutureReqField(p *C.CThostFtdcTransferBankToFutureReqField) *CThostFtdcTransferBankToFutureReqField {
	ret := new(CThostFtdcTransferBankToFutureReqField)
	ret.FutureAccount = c2goStr(&p.FutureAccount[0], C.sizeof_TThostFtdcAccountIDType)
	ret.FuturePwdFlag = byte(p.FuturePwdFlag)
	ret.FutureAccPwd = c2goStr(&p.FutureAccPwd[0], C.sizeof_TThostFtdcFutureAccPwdType)
	ret.TradeAmt = goFloat64(p.TradeAmt)
	ret.CustFee = goFloat64(p.CustFee)
	ret.CurrencyCode = c2goStr(&p.CurrencyCode[0], C.sizeof_TThostFtdcCurrencyCodeType)
	return ret
}
func (s *CThostFtdcTransferBankToFutureReqField) CValue() *C.CThostFtdcTransferBankToFutureReqField {
	ptr := C.newCThostFtdcTransferBankToFutureReqField()
	go2cStr(s.FutureAccount, &ptr.FutureAccount[0], C.sizeof_TThostFtdcAccountIDType)
	ptr.FuturePwdFlag = C.char(s.FuturePwdFlag)
	go2cStr(s.FutureAccPwd, &ptr.FutureAccPwd[0], C.sizeof_TThostFtdcFutureAccPwdType)
	ptr.TradeAmt = C.double(s.TradeAmt)
	ptr.CustFee = C.double(s.CustFee)
	go2cStr(s.CurrencyCode, &ptr.CurrencyCode[0], C.sizeof_TThostFtdcCurrencyCodeType)
	return ptr
}

type CThostFtdcTransferBankToFutureRspField struct {
	RetCode       string
	RetInfo       string
	FutureAccount string
	TradeAmt      float64
	CustFee       float64
	CurrencyCode  string
}

func NewCThostFtdcTransferBankToFutureRspField(p *C.CThostFtdcTransferBankToFutureRspField) *CThostFtdcTransferBankToFutureRspField {
	ret := new(CThostFtdcTransferBankToFutureRspField)
	ret.RetCode = c2goStr(&p.RetCode[0], C.sizeof_TThostFtdcRetCodeType)
	ret.RetInfo = c2goStr(&p.RetInfo[0], C.sizeof_TThostFtdcRetInfoType)
	ret.FutureAccount = c2goStr(&p.FutureAccount[0], C.sizeof_TThostFtdcAccountIDType)
	ret.TradeAmt = goFloat64(p.TradeAmt)
	ret.CustFee = goFloat64(p.CustFee)
	ret.CurrencyCode = c2goStr(&p.CurrencyCode[0], C.sizeof_TThostFtdcCurrencyCodeType)
	return ret
}
func (s *CThostFtdcTransferBankToFutureRspField) CValue() *C.CThostFtdcTransferBankToFutureRspField {
	ptr := C.newCThostFtdcTransferBankToFutureRspField()
	go2cStr(s.RetCode, &ptr.RetCode[0], C.sizeof_TThostFtdcRetCodeType)
	go2cStr(s.RetInfo, &ptr.RetInfo[0], C.sizeof_TThostFtdcRetInfoType)
	go2cStr(s.FutureAccount, &ptr.FutureAccount[0], C.sizeof_TThostFtdcAccountIDType)
	ptr.TradeAmt = C.double(s.TradeAmt)
	ptr.CustFee = C.double(s.CustFee)
	go2cStr(s.CurrencyCode, &ptr.CurrencyCode[0], C.sizeof_TThostFtdcCurrencyCodeType)
	return ptr
}

type CThostFtdcTransferFutureToBankReqField struct {
	FutureAccount string
	FuturePwdFlag byte
	FutureAccPwd  string
	TradeAmt      float64
	CustFee       float64
	CurrencyCode  string
}

func NewCThostFtdcTransferFutureToBankReqField(p *C.CThostFtdcTransferFutureToBankReqField) *CThostFtdcTransferFutureToBankReqField {
	ret := new(CThostFtdcTransferFutureToBankReqField)
	ret.FutureAccount = c2goStr(&p.FutureAccount[0], C.sizeof_TThostFtdcAccountIDType)
	ret.FuturePwdFlag = byte(p.FuturePwdFlag)
	ret.FutureAccPwd = c2goStr(&p.FutureAccPwd[0], C.sizeof_TThostFtdcFutureAccPwdType)
	ret.TradeAmt = goFloat64(p.TradeAmt)
	ret.CustFee = goFloat64(p.CustFee)
	ret.CurrencyCode = c2goStr(&p.CurrencyCode[0], C.sizeof_TThostFtdcCurrencyCodeType)
	return ret
}
func (s *CThostFtdcTransferFutureToBankReqField) CValue() *C.CThostFtdcTransferFutureToBankReqField {
	ptr := C.newCThostFtdcTransferFutureToBankReqField()
	go2cStr(s.FutureAccount, &ptr.FutureAccount[0], C.sizeof_TThostFtdcAccountIDType)
	ptr.FuturePwdFlag = C.char(s.FuturePwdFlag)
	go2cStr(s.FutureAccPwd, &ptr.FutureAccPwd[0], C.sizeof_TThostFtdcFutureAccPwdType)
	ptr.TradeAmt = C.double(s.TradeAmt)
	ptr.CustFee = C.double(s.CustFee)
	go2cStr(s.CurrencyCode, &ptr.CurrencyCode[0], C.sizeof_TThostFtdcCurrencyCodeType)
	return ptr
}

type CThostFtdcTransferFutureToBankRspField struct {
	RetCode       string
	RetInfo       string
	FutureAccount string
	TradeAmt      float64
	CustFee       float64
	CurrencyCode  string
}

func NewCThostFtdcTransferFutureToBankRspField(p *C.CThostFtdcTransferFutureToBankRspField) *CThostFtdcTransferFutureToBankRspField {
	ret := new(CThostFtdcTransferFutureToBankRspField)
	ret.RetCode = c2goStr(&p.RetCode[0], C.sizeof_TThostFtdcRetCodeType)
	ret.RetInfo = c2goStr(&p.RetInfo[0], C.sizeof_TThostFtdcRetInfoType)
	ret.FutureAccount = c2goStr(&p.FutureAccount[0], C.sizeof_TThostFtdcAccountIDType)
	ret.TradeAmt = goFloat64(p.TradeAmt)
	ret.CustFee = goFloat64(p.CustFee)
	ret.CurrencyCode = c2goStr(&p.CurrencyCode[0], C.sizeof_TThostFtdcCurrencyCodeType)
	return ret
}
func (s *CThostFtdcTransferFutureToBankRspField) CValue() *C.CThostFtdcTransferFutureToBankRspField {
	ptr := C.newCThostFtdcTransferFutureToBankRspField()
	go2cStr(s.RetCode, &ptr.RetCode[0], C.sizeof_TThostFtdcRetCodeType)
	go2cStr(s.RetInfo, &ptr.RetInfo[0], C.sizeof_TThostFtdcRetInfoType)
	go2cStr(s.FutureAccount, &ptr.FutureAccount[0], C.sizeof_TThostFtdcAccountIDType)
	ptr.TradeAmt = C.double(s.TradeAmt)
	ptr.CustFee = C.double(s.CustFee)
	go2cStr(s.CurrencyCode, &ptr.CurrencyCode[0], C.sizeof_TThostFtdcCurrencyCodeType)
	return ptr
}

type CThostFtdcTransferQryBankReqField struct {
	FutureAccount string
	FuturePwdFlag byte
	FutureAccPwd  string
	CurrencyCode  string
}

func NewCThostFtdcTransferQryBankReqField(p *C.CThostFtdcTransferQryBankReqField) *CThostFtdcTransferQryBankReqField {
	ret := new(CThostFtdcTransferQryBankReqField)
	ret.FutureAccount = c2goStr(&p.FutureAccount[0], C.sizeof_TThostFtdcAccountIDType)
	ret.FuturePwdFlag = byte(p.FuturePwdFlag)
	ret.FutureAccPwd = c2goStr(&p.FutureAccPwd[0], C.sizeof_TThostFtdcFutureAccPwdType)
	ret.CurrencyCode = c2goStr(&p.CurrencyCode[0], C.sizeof_TThostFtdcCurrencyCodeType)
	return ret
}
func (s *CThostFtdcTransferQryBankReqField) CValue() *C.CThostFtdcTransferQryBankReqField {
	ptr := C.newCThostFtdcTransferQryBankReqField()
	go2cStr(s.FutureAccount, &ptr.FutureAccount[0], C.sizeof_TThostFtdcAccountIDType)
	ptr.FuturePwdFlag = C.char(s.FuturePwdFlag)
	go2cStr(s.FutureAccPwd, &ptr.FutureAccPwd[0], C.sizeof_TThostFtdcFutureAccPwdType)
	go2cStr(s.CurrencyCode, &ptr.CurrencyCode[0], C.sizeof_TThostFtdcCurrencyCodeType)
	return ptr
}

type CThostFtdcTransferQryBankRspField struct {
	RetCode       string
	RetInfo       string
	FutureAccount string
	TradeAmt      float64
	UseAmt        float64
	FetchAmt      float64
	CurrencyCode  string
}

func NewCThostFtdcTransferQryBankRspField(p *C.CThostFtdcTransferQryBankRspField) *CThostFtdcTransferQryBankRspField {
	ret := new(CThostFtdcTransferQryBankRspField)
	ret.RetCode = c2goStr(&p.RetCode[0], C.sizeof_TThostFtdcRetCodeType)
	ret.RetInfo = c2goStr(&p.RetInfo[0], C.sizeof_TThostFtdcRetInfoType)
	ret.FutureAccount = c2goStr(&p.FutureAccount[0], C.sizeof_TThostFtdcAccountIDType)
	ret.TradeAmt = goFloat64(p.TradeAmt)
	ret.UseAmt = goFloat64(p.UseAmt)
	ret.FetchAmt = goFloat64(p.FetchAmt)
	ret.CurrencyCode = c2goStr(&p.CurrencyCode[0], C.sizeof_TThostFtdcCurrencyCodeType)
	return ret
}
func (s *CThostFtdcTransferQryBankRspField) CValue() *C.CThostFtdcTransferQryBankRspField {
	ptr := C.newCThostFtdcTransferQryBankRspField()
	go2cStr(s.RetCode, &ptr.RetCode[0], C.sizeof_TThostFtdcRetCodeType)
	go2cStr(s.RetInfo, &ptr.RetInfo[0], C.sizeof_TThostFtdcRetInfoType)
	go2cStr(s.FutureAccount, &ptr.FutureAccount[0], C.sizeof_TThostFtdcAccountIDType)
	ptr.TradeAmt = C.double(s.TradeAmt)
	ptr.UseAmt = C.double(s.UseAmt)
	ptr.FetchAmt = C.double(s.FetchAmt)
	go2cStr(s.CurrencyCode, &ptr.CurrencyCode[0], C.sizeof_TThostFtdcCurrencyCodeType)
	return ptr
}

type CThostFtdcTransferQryDetailReqField struct {
	FutureAccount string
}

func NewCThostFtdcTransferQryDetailReqField(p *C.CThostFtdcTransferQryDetailReqField) *CThostFtdcTransferQryDetailReqField {
	ret := new(CThostFtdcTransferQryDetailReqField)
	ret.FutureAccount = c2goStr(&p.FutureAccount[0], C.sizeof_TThostFtdcAccountIDType)
	return ret
}
func (s *CThostFtdcTransferQryDetailReqField) CValue() *C.CThostFtdcTransferQryDetailReqField {
	ptr := C.newCThostFtdcTransferQryDetailReqField()
	go2cStr(s.FutureAccount, &ptr.FutureAccount[0], C.sizeof_TThostFtdcAccountIDType)
	return ptr
}

type CThostFtdcTransferQryDetailRspField struct {
	TradeDate     string
	TradeTime     string
	TradeCode     string
	FutureSerial  int
	FutureID      string
	FutureAccount string
	BankSerial    int
	BankID        string
	BankBrchID    string
	BankAccount   string
	CertCode      string
	CurrencyCode  string
	TxAmount      float64
	Flag          byte
}

func NewCThostFtdcTransferQryDetailRspField(p *C.CThostFtdcTransferQryDetailRspField) *CThostFtdcTransferQryDetailRspField {
	ret := new(CThostFtdcTransferQryDetailRspField)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.FutureSerial = int(p.FutureSerial)
	ret.FutureID = c2goStr(&p.FutureID[0], C.sizeof_TThostFtdcFutureIDType)
	ret.FutureAccount = c2goStr(&p.FutureAccount[0], C.sizeof_TThostFtdcFutureAccountType)
	ret.BankSerial = int(p.BankSerial)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBrchID = c2goStr(&p.BankBrchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.CertCode = c2goStr(&p.CertCode[0], C.sizeof_TThostFtdcCertCodeType)
	ret.CurrencyCode = c2goStr(&p.CurrencyCode[0], C.sizeof_TThostFtdcCurrencyCodeType)
	ret.TxAmount = goFloat64(p.TxAmount)
	ret.Flag = byte(p.Flag)
	return ret
}
func (s *CThostFtdcTransferQryDetailRspField) CValue() *C.CThostFtdcTransferQryDetailRspField {
	ptr := C.newCThostFtdcTransferQryDetailRspField()
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ptr.FutureSerial = C.int(s.FutureSerial)
	go2cStr(s.FutureID, &ptr.FutureID[0], C.sizeof_TThostFtdcFutureIDType)
	go2cStr(s.FutureAccount, &ptr.FutureAccount[0], C.sizeof_TThostFtdcFutureAccountType)
	ptr.BankSerial = C.int(s.BankSerial)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBrchID, &ptr.BankBrchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.CertCode, &ptr.CertCode[0], C.sizeof_TThostFtdcCertCodeType)
	go2cStr(s.CurrencyCode, &ptr.CurrencyCode[0], C.sizeof_TThostFtdcCurrencyCodeType)
	ptr.TxAmount = C.double(s.TxAmount)
	ptr.Flag = C.char(s.Flag)
	return ptr
}

type CThostFtdcRspInfoField struct {
	ErrorID  int
	ErrorMsg string
}

func NewCThostFtdcRspInfoField(p *C.CThostFtdcRspInfoField) *CThostFtdcRspInfoField {
	ret := new(CThostFtdcRspInfoField)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ret
}
func (s *CThostFtdcRspInfoField) CValue() *C.CThostFtdcRspInfoField {
	ptr := C.newCThostFtdcRspInfoField()
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ptr
}

type CThostFtdcExchangeField struct {
	ExchangeID       string
	ExchangeName     string
	ExchangeProperty byte
}

func NewCThostFtdcExchangeField(p *C.CThostFtdcExchangeField) *CThostFtdcExchangeField {
	ret := new(CThostFtdcExchangeField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ExchangeName = c2goStr(&p.ExchangeName[0], C.sizeof_TThostFtdcExchangeNameType)
	ret.ExchangeProperty = byte(p.ExchangeProperty)
	return ret
}
func (s *CThostFtdcExchangeField) CValue() *C.CThostFtdcExchangeField {
	ptr := C.newCThostFtdcExchangeField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ExchangeName, &ptr.ExchangeName[0], C.sizeof_TThostFtdcExchangeNameType)
	ptr.ExchangeProperty = C.char(s.ExchangeProperty)
	return ptr
}

type CThostFtdcProductField struct {
	Reserve1             string
	ProductName          string
	ExchangeID           string
	ProductClass         byte
	VolumeMultiple       int
	PriceTick            float64
	MaxMarketOrderVolume int
	MinMarketOrderVolume int
	MaxLimitOrderVolume  int
	MinLimitOrderVolume  int
	PositionType         byte
	PositionDateType     byte
	CloseDealType        byte
	TradeCurrencyID      string
	MortgageFundUseRange byte
	Reserve2             string
	UnderlyingMultiple   float64
	ProductID            string
	ExchangeProductID    string
}

func NewCThostFtdcProductField(p *C.CThostFtdcProductField) *CThostFtdcProductField {
	ret := new(CThostFtdcProductField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ProductName = c2goStr(&p.ProductName[0], C.sizeof_TThostFtdcProductNameType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ProductClass = byte(p.ProductClass)
	ret.VolumeMultiple = int(p.VolumeMultiple)
	ret.PriceTick = goFloat64(p.PriceTick)
	ret.MaxMarketOrderVolume = int(p.MaxMarketOrderVolume)
	ret.MinMarketOrderVolume = int(p.MinMarketOrderVolume)
	ret.MaxLimitOrderVolume = int(p.MaxLimitOrderVolume)
	ret.MinLimitOrderVolume = int(p.MinLimitOrderVolume)
	ret.PositionType = byte(p.PositionType)
	ret.PositionDateType = byte(p.PositionDateType)
	ret.CloseDealType = byte(p.CloseDealType)
	ret.TradeCurrencyID = c2goStr(&p.TradeCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.MortgageFundUseRange = byte(p.MortgageFundUseRange)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.UnderlyingMultiple = goFloat64(p.UnderlyingMultiple)
	ret.ProductID = c2goStr(&p.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ExchangeProductID = c2goStr(&p.ExchangeProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcProductField) CValue() *C.CThostFtdcProductField {
	ptr := C.newCThostFtdcProductField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ProductName, &ptr.ProductName[0], C.sizeof_TThostFtdcProductNameType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ptr.ProductClass = C.char(s.ProductClass)
	ptr.VolumeMultiple = C.int(s.VolumeMultiple)
	ptr.PriceTick = C.double(s.PriceTick)
	ptr.MaxMarketOrderVolume = C.int(s.MaxMarketOrderVolume)
	ptr.MinMarketOrderVolume = C.int(s.MinMarketOrderVolume)
	ptr.MaxLimitOrderVolume = C.int(s.MaxLimitOrderVolume)
	ptr.MinLimitOrderVolume = C.int(s.MinLimitOrderVolume)
	ptr.PositionType = C.char(s.PositionType)
	ptr.PositionDateType = C.char(s.PositionDateType)
	ptr.CloseDealType = C.char(s.CloseDealType)
	go2cStr(s.TradeCurrencyID, &ptr.TradeCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.MortgageFundUseRange = C.char(s.MortgageFundUseRange)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.UnderlyingMultiple = C.double(s.UnderlyingMultiple)
	go2cStr(s.ProductID, &ptr.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ExchangeProductID, &ptr.ExchangeProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcInstrumentField struct {
	Reserve1               string
	ExchangeID             string
	InstrumentName         string
	Reserve2               string
	Reserve3               string
	ProductClass           byte
	DeliveryYear           int
	DeliveryMonth          int
	MaxMarketOrderVolume   int
	MinMarketOrderVolume   int
	MaxLimitOrderVolume    int
	MinLimitOrderVolume    int
	VolumeMultiple         int
	PriceTick              float64
	CreateDate             string
	OpenDate               string
	ExpireDate             string
	StartDelivDate         string
	EndDelivDate           string
	InstLifePhase          byte
	IsTrading              int
	PositionType           byte
	PositionDateType       byte
	LongMarginRatio        float64
	ShortMarginRatio       float64
	MaxMarginSideAlgorithm byte
	Reserve4               string
	StrikePrice            float64
	OptionsType            byte
	UnderlyingMultiple     float64
	CombinationType        byte
	InstrumentID           string
	ExchangeInstID         string
	ProductID              string
	UnderlyingInstrID      string
}

func NewCThostFtdcInstrumentField(p *C.CThostFtdcInstrumentField) *CThostFtdcInstrumentField {
	ret := new(CThostFtdcInstrumentField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InstrumentName = c2goStr(&p.InstrumentName[0], C.sizeof_TThostFtdcInstrumentNameType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.Reserve3 = c2goStr(&p.reserve3[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ProductClass = byte(p.ProductClass)
	ret.DeliveryYear = int(p.DeliveryYear)
	ret.DeliveryMonth = int(p.DeliveryMonth)
	ret.MaxMarketOrderVolume = int(p.MaxMarketOrderVolume)
	ret.MinMarketOrderVolume = int(p.MinMarketOrderVolume)
	ret.MaxLimitOrderVolume = int(p.MaxLimitOrderVolume)
	ret.MinLimitOrderVolume = int(p.MinLimitOrderVolume)
	ret.VolumeMultiple = int(p.VolumeMultiple)
	ret.PriceTick = goFloat64(p.PriceTick)
	ret.CreateDate = c2goStr(&p.CreateDate[0], C.sizeof_TThostFtdcDateType)
	ret.OpenDate = c2goStr(&p.OpenDate[0], C.sizeof_TThostFtdcDateType)
	ret.ExpireDate = c2goStr(&p.ExpireDate[0], C.sizeof_TThostFtdcDateType)
	ret.StartDelivDate = c2goStr(&p.StartDelivDate[0], C.sizeof_TThostFtdcDateType)
	ret.EndDelivDate = c2goStr(&p.EndDelivDate[0], C.sizeof_TThostFtdcDateType)
	ret.InstLifePhase = byte(p.InstLifePhase)
	ret.IsTrading = int(p.IsTrading)
	ret.PositionType = byte(p.PositionType)
	ret.PositionDateType = byte(p.PositionDateType)
	ret.LongMarginRatio = goFloat64(p.LongMarginRatio)
	ret.ShortMarginRatio = goFloat64(p.ShortMarginRatio)
	ret.MaxMarginSideAlgorithm = byte(p.MaxMarginSideAlgorithm)
	ret.Reserve4 = c2goStr(&p.reserve4[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.StrikePrice = goFloat64(p.StrikePrice)
	ret.OptionsType = byte(p.OptionsType)
	ret.UnderlyingMultiple = goFloat64(p.UnderlyingMultiple)
	ret.CombinationType = byte(p.CombinationType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.ProductID = c2goStr(&p.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.UnderlyingInstrID = c2goStr(&p.UnderlyingInstrID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcInstrumentField) CValue() *C.CThostFtdcInstrumentField {
	ptr := C.newCThostFtdcInstrumentField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InstrumentName, &ptr.InstrumentName[0], C.sizeof_TThostFtdcInstrumentNameType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.Reserve3, &ptr.reserve3[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.ProductClass = C.char(s.ProductClass)
	ptr.DeliveryYear = C.int(s.DeliveryYear)
	ptr.DeliveryMonth = C.int(s.DeliveryMonth)
	ptr.MaxMarketOrderVolume = C.int(s.MaxMarketOrderVolume)
	ptr.MinMarketOrderVolume = C.int(s.MinMarketOrderVolume)
	ptr.MaxLimitOrderVolume = C.int(s.MaxLimitOrderVolume)
	ptr.MinLimitOrderVolume = C.int(s.MinLimitOrderVolume)
	ptr.VolumeMultiple = C.int(s.VolumeMultiple)
	ptr.PriceTick = C.double(s.PriceTick)
	go2cStr(s.CreateDate, &ptr.CreateDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.OpenDate, &ptr.OpenDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ExpireDate, &ptr.ExpireDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.StartDelivDate, &ptr.StartDelivDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.EndDelivDate, &ptr.EndDelivDate[0], C.sizeof_TThostFtdcDateType)
	ptr.InstLifePhase = C.char(s.InstLifePhase)
	ptr.IsTrading = C.int(s.IsTrading)
	ptr.PositionType = C.char(s.PositionType)
	ptr.PositionDateType = C.char(s.PositionDateType)
	ptr.LongMarginRatio = C.double(s.LongMarginRatio)
	ptr.ShortMarginRatio = C.double(s.ShortMarginRatio)
	ptr.MaxMarginSideAlgorithm = C.char(s.MaxMarginSideAlgorithm)
	go2cStr(s.Reserve4, &ptr.reserve4[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.StrikePrice = C.double(s.StrikePrice)
	ptr.OptionsType = C.char(s.OptionsType)
	ptr.UnderlyingMultiple = C.double(s.UnderlyingMultiple)
	ptr.CombinationType = C.char(s.CombinationType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	go2cStr(s.ProductID, &ptr.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.UnderlyingInstrID, &ptr.UnderlyingInstrID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcBrokerField struct {
	BrokerID   string
	BrokerAbbr string
	BrokerName string
	IsActive   int
}

func NewCThostFtdcBrokerField(p *C.CThostFtdcBrokerField) *CThostFtdcBrokerField {
	ret := new(CThostFtdcBrokerField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerAbbr = c2goStr(&p.BrokerAbbr[0], C.sizeof_TThostFtdcBrokerAbbrType)
	ret.BrokerName = c2goStr(&p.BrokerName[0], C.sizeof_TThostFtdcBrokerNameType)
	ret.IsActive = int(p.IsActive)
	return ret
}
func (s *CThostFtdcBrokerField) CValue() *C.CThostFtdcBrokerField {
	ptr := C.newCThostFtdcBrokerField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerAbbr, &ptr.BrokerAbbr[0], C.sizeof_TThostFtdcBrokerAbbrType)
	go2cStr(s.BrokerName, &ptr.BrokerName[0], C.sizeof_TThostFtdcBrokerNameType)
	ptr.IsActive = C.int(s.IsActive)
	return ptr
}

type CThostFtdcTraderField struct {
	ExchangeID    string
	TraderID      string
	ParticipantID string
	Password      string
	InstallCount  int
	BrokerID      string
}

func NewCThostFtdcTraderField(p *C.CThostFtdcTraderField) *CThostFtdcTraderField {
	ret := new(CThostFtdcTraderField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.InstallCount = int(p.InstallCount)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	return ret
}
func (s *CThostFtdcTraderField) CValue() *C.CThostFtdcTraderField {
	ptr := C.newCThostFtdcTraderField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	ptr.InstallCount = C.int(s.InstallCount)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	return ptr
}

type CThostFtdcInvestorField struct {
	InvestorID         string
	BrokerID           string
	InvestorGroupID    string
	InvestorName       string
	IdentifiedCardType byte
	IdentifiedCardNo   string
	IsActive           int
	Telephone          string
	Address            string
	OpenDate           string
	Mobile             string
	CommModelID        string
	MarginModelID      string
}

func NewCThostFtdcInvestorField(p *C.CThostFtdcInvestorField) *CThostFtdcInvestorField {
	ret := new(CThostFtdcInvestorField)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorGroupID = c2goStr(&p.InvestorGroupID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.InvestorName = c2goStr(&p.InvestorName[0], C.sizeof_TThostFtdcPartyNameType)
	ret.IdentifiedCardType = byte(p.IdentifiedCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.IsActive = int(p.IsActive)
	ret.Telephone = c2goStr(&p.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	ret.Address = c2goStr(&p.Address[0], C.sizeof_TThostFtdcAddressType)
	ret.OpenDate = c2goStr(&p.OpenDate[0], C.sizeof_TThostFtdcDateType)
	ret.Mobile = c2goStr(&p.Mobile[0], C.sizeof_TThostFtdcMobileType)
	ret.CommModelID = c2goStr(&p.CommModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.MarginModelID = c2goStr(&p.MarginModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ret
}
func (s *CThostFtdcInvestorField) CValue() *C.CThostFtdcInvestorField {
	ptr := C.newCThostFtdcInvestorField()
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorGroupID, &ptr.InvestorGroupID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.InvestorName, &ptr.InvestorName[0], C.sizeof_TThostFtdcPartyNameType)
	ptr.IdentifiedCardType = C.char(s.IdentifiedCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.IsActive = C.int(s.IsActive)
	go2cStr(s.Telephone, &ptr.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	go2cStr(s.Address, &ptr.Address[0], C.sizeof_TThostFtdcAddressType)
	go2cStr(s.OpenDate, &ptr.OpenDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.Mobile, &ptr.Mobile[0], C.sizeof_TThostFtdcMobileType)
	go2cStr(s.CommModelID, &ptr.CommModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.MarginModelID, &ptr.MarginModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ptr
}

type CThostFtdcTradingCodeField struct {
	InvestorID   string
	BrokerID     string
	ExchangeID   string
	ClientID     string
	IsActive     int
	ClientIDType byte
	BranchID     string
	BizType      byte
	InvestUnitID string
}

func NewCThostFtdcTradingCodeField(p *C.CThostFtdcTradingCodeField) *CThostFtdcTradingCodeField {
	ret := new(CThostFtdcTradingCodeField)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.IsActive = int(p.IsActive)
	ret.ClientIDType = byte(p.ClientIDType)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.BizType = byte(p.BizType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ret
}
func (s *CThostFtdcTradingCodeField) CValue() *C.CThostFtdcTradingCodeField {
	ptr := C.newCThostFtdcTradingCodeField()
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ptr.IsActive = C.int(s.IsActive)
	ptr.ClientIDType = C.char(s.ClientIDType)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ptr.BizType = C.char(s.BizType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ptr
}

type CThostFtdcPartBrokerField struct {
	BrokerID      string
	ExchangeID    string
	ParticipantID string
	IsActive      int
}

func NewCThostFtdcPartBrokerField(p *C.CThostFtdcPartBrokerField) *CThostFtdcPartBrokerField {
	ret := new(CThostFtdcPartBrokerField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.IsActive = int(p.IsActive)
	return ret
}
func (s *CThostFtdcPartBrokerField) CValue() *C.CThostFtdcPartBrokerField {
	ptr := C.newCThostFtdcPartBrokerField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ptr.IsActive = C.int(s.IsActive)
	return ptr
}

type CThostFtdcSuperUserField struct {
	UserID   string
	UserName string
	Password string
	IsActive int
}

func NewCThostFtdcSuperUserField(p *C.CThostFtdcSuperUserField) *CThostFtdcSuperUserField {
	ret := new(CThostFtdcSuperUserField)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.UserName = c2goStr(&p.UserName[0], C.sizeof_TThostFtdcUserNameType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.IsActive = int(p.IsActive)
	return ret
}
func (s *CThostFtdcSuperUserField) CValue() *C.CThostFtdcSuperUserField {
	ptr := C.newCThostFtdcSuperUserField()
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.UserName, &ptr.UserName[0], C.sizeof_TThostFtdcUserNameType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	ptr.IsActive = C.int(s.IsActive)
	return ptr
}

type CThostFtdcSuperUserFunctionField struct {
	UserID       string
	FunctionCode byte
}

func NewCThostFtdcSuperUserFunctionField(p *C.CThostFtdcSuperUserFunctionField) *CThostFtdcSuperUserFunctionField {
	ret := new(CThostFtdcSuperUserFunctionField)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.FunctionCode = byte(p.FunctionCode)
	return ret
}
func (s *CThostFtdcSuperUserFunctionField) CValue() *C.CThostFtdcSuperUserFunctionField {
	ptr := C.newCThostFtdcSuperUserFunctionField()
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.FunctionCode = C.char(s.FunctionCode)
	return ptr
}

type CThostFtdcInvestorGroupField struct {
	BrokerID          string
	InvestorGroupID   string
	InvestorGroupName string
}

func NewCThostFtdcInvestorGroupField(p *C.CThostFtdcInvestorGroupField) *CThostFtdcInvestorGroupField {
	ret := new(CThostFtdcInvestorGroupField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorGroupID = c2goStr(&p.InvestorGroupID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.InvestorGroupName = c2goStr(&p.InvestorGroupName[0], C.sizeof_TThostFtdcInvestorGroupNameType)
	return ret
}
func (s *CThostFtdcInvestorGroupField) CValue() *C.CThostFtdcInvestorGroupField {
	ptr := C.newCThostFtdcInvestorGroupField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorGroupID, &ptr.InvestorGroupID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.InvestorGroupName, &ptr.InvestorGroupName[0], C.sizeof_TThostFtdcInvestorGroupNameType)
	return ptr
}

type CThostFtdcTradingAccountField struct {
	BrokerID                       string
	AccountID                      string
	PreMortgage                    float64
	PreCredit                      float64
	PreDeposit                     float64
	PreBalance                     float64
	PreMargin                      float64
	InterestBase                   float64
	Interest                       float64
	Deposit                        float64
	Withdraw                       float64
	FrozenMargin                   float64
	FrozenCash                     float64
	FrozenCommission               float64
	CurrMargin                     float64
	CashIn                         float64
	Commission                     float64
	CloseProfit                    float64
	PositionProfit                 float64
	Balance                        float64
	Available                      float64
	WithdrawQuota                  float64
	Reserve                        float64
	TradingDay                     string
	SettlementID                   int
	Credit                         float64
	Mortgage                       float64
	ExchangeMargin                 float64
	DeliveryMargin                 float64
	ExchangeDeliveryMargin         float64
	ReserveBalance                 float64
	CurrencyID                     string
	PreFundMortgageIn              float64
	PreFundMortgageOut             float64
	FundMortgageIn                 float64
	FundMortgageOut                float64
	FundMortgageAvailable          float64
	MortgageableFund               float64
	SpecProductMargin              float64
	SpecProductFrozenMargin        float64
	SpecProductCommission          float64
	SpecProductFrozenCommission    float64
	SpecProductPositionProfit      float64
	SpecProductCloseProfit         float64
	SpecProductPositionProfitByAlg float64
	SpecProductExchangeMargin      float64
	BizType                        byte
	FrozenSwap                     float64
	RemainSwap                     float64
}

func NewCThostFtdcTradingAccountField(p *C.CThostFtdcTradingAccountField) *CThostFtdcTradingAccountField {
	ret := new(CThostFtdcTradingAccountField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.PreMortgage = goFloat64(p.PreMortgage)
	ret.PreCredit = goFloat64(p.PreCredit)
	ret.PreDeposit = goFloat64(p.PreDeposit)
	ret.PreBalance = goFloat64(p.PreBalance)
	ret.PreMargin = goFloat64(p.PreMargin)
	ret.InterestBase = goFloat64(p.InterestBase)
	ret.Interest = goFloat64(p.Interest)
	ret.Deposit = goFloat64(p.Deposit)
	ret.Withdraw = goFloat64(p.Withdraw)
	ret.FrozenMargin = goFloat64(p.FrozenMargin)
	ret.FrozenCash = goFloat64(p.FrozenCash)
	ret.FrozenCommission = goFloat64(p.FrozenCommission)
	ret.CurrMargin = goFloat64(p.CurrMargin)
	ret.CashIn = goFloat64(p.CashIn)
	ret.Commission = goFloat64(p.Commission)
	ret.CloseProfit = goFloat64(p.CloseProfit)
	ret.PositionProfit = goFloat64(p.PositionProfit)
	ret.Balance = goFloat64(p.Balance)
	ret.Available = goFloat64(p.Available)
	ret.WithdrawQuota = goFloat64(p.WithdrawQuota)
	ret.Reserve = goFloat64(p.Reserve)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.Credit = goFloat64(p.Credit)
	ret.Mortgage = goFloat64(p.Mortgage)
	ret.ExchangeMargin = goFloat64(p.ExchangeMargin)
	ret.DeliveryMargin = goFloat64(p.DeliveryMargin)
	ret.ExchangeDeliveryMargin = goFloat64(p.ExchangeDeliveryMargin)
	ret.ReserveBalance = goFloat64(p.ReserveBalance)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.PreFundMortgageIn = goFloat64(p.PreFundMortgageIn)
	ret.PreFundMortgageOut = goFloat64(p.PreFundMortgageOut)
	ret.FundMortgageIn = goFloat64(p.FundMortgageIn)
	ret.FundMortgageOut = goFloat64(p.FundMortgageOut)
	ret.FundMortgageAvailable = goFloat64(p.FundMortgageAvailable)
	ret.MortgageableFund = goFloat64(p.MortgageableFund)
	ret.SpecProductMargin = goFloat64(p.SpecProductMargin)
	ret.SpecProductFrozenMargin = goFloat64(p.SpecProductFrozenMargin)
	ret.SpecProductCommission = goFloat64(p.SpecProductCommission)
	ret.SpecProductFrozenCommission = goFloat64(p.SpecProductFrozenCommission)
	ret.SpecProductPositionProfit = goFloat64(p.SpecProductPositionProfit)
	ret.SpecProductCloseProfit = goFloat64(p.SpecProductCloseProfit)
	ret.SpecProductPositionProfitByAlg = goFloat64(p.SpecProductPositionProfitByAlg)
	ret.SpecProductExchangeMargin = goFloat64(p.SpecProductExchangeMargin)
	ret.BizType = byte(p.BizType)
	ret.FrozenSwap = goFloat64(p.FrozenSwap)
	ret.RemainSwap = goFloat64(p.RemainSwap)
	return ret
}
func (s *CThostFtdcTradingAccountField) CValue() *C.CThostFtdcTradingAccountField {
	ptr := C.newCThostFtdcTradingAccountField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ptr.PreMortgage = C.double(s.PreMortgage)
	ptr.PreCredit = C.double(s.PreCredit)
	ptr.PreDeposit = C.double(s.PreDeposit)
	ptr.PreBalance = C.double(s.PreBalance)
	ptr.PreMargin = C.double(s.PreMargin)
	ptr.InterestBase = C.double(s.InterestBase)
	ptr.Interest = C.double(s.Interest)
	ptr.Deposit = C.double(s.Deposit)
	ptr.Withdraw = C.double(s.Withdraw)
	ptr.FrozenMargin = C.double(s.FrozenMargin)
	ptr.FrozenCash = C.double(s.FrozenCash)
	ptr.FrozenCommission = C.double(s.FrozenCommission)
	ptr.CurrMargin = C.double(s.CurrMargin)
	ptr.CashIn = C.double(s.CashIn)
	ptr.Commission = C.double(s.Commission)
	ptr.CloseProfit = C.double(s.CloseProfit)
	ptr.PositionProfit = C.double(s.PositionProfit)
	ptr.Balance = C.double(s.Balance)
	ptr.Available = C.double(s.Available)
	ptr.WithdrawQuota = C.double(s.WithdrawQuota)
	ptr.Reserve = C.double(s.Reserve)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	ptr.Credit = C.double(s.Credit)
	ptr.Mortgage = C.double(s.Mortgage)
	ptr.ExchangeMargin = C.double(s.ExchangeMargin)
	ptr.DeliveryMargin = C.double(s.DeliveryMargin)
	ptr.ExchangeDeliveryMargin = C.double(s.ExchangeDeliveryMargin)
	ptr.ReserveBalance = C.double(s.ReserveBalance)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.PreFundMortgageIn = C.double(s.PreFundMortgageIn)
	ptr.PreFundMortgageOut = C.double(s.PreFundMortgageOut)
	ptr.FundMortgageIn = C.double(s.FundMortgageIn)
	ptr.FundMortgageOut = C.double(s.FundMortgageOut)
	ptr.FundMortgageAvailable = C.double(s.FundMortgageAvailable)
	ptr.MortgageableFund = C.double(s.MortgageableFund)
	ptr.SpecProductMargin = C.double(s.SpecProductMargin)
	ptr.SpecProductFrozenMargin = C.double(s.SpecProductFrozenMargin)
	ptr.SpecProductCommission = C.double(s.SpecProductCommission)
	ptr.SpecProductFrozenCommission = C.double(s.SpecProductFrozenCommission)
	ptr.SpecProductPositionProfit = C.double(s.SpecProductPositionProfit)
	ptr.SpecProductCloseProfit = C.double(s.SpecProductCloseProfit)
	ptr.SpecProductPositionProfitByAlg = C.double(s.SpecProductPositionProfitByAlg)
	ptr.SpecProductExchangeMargin = C.double(s.SpecProductExchangeMargin)
	ptr.BizType = C.char(s.BizType)
	ptr.FrozenSwap = C.double(s.FrozenSwap)
	ptr.RemainSwap = C.double(s.RemainSwap)
	return ptr
}

type CThostFtdcInvestorPositionField struct {
	Reserve1           string
	BrokerID           string
	InvestorID         string
	PosiDirection      byte
	HedgeFlag          byte
	PositionDate       byte
	YdPosition         int
	Position           int
	LongFrozen         int
	ShortFrozen        int
	LongFrozenAmount   float64
	ShortFrozenAmount  float64
	OpenVolume         int
	CloseVolume        int
	OpenAmount         float64
	CloseAmount        float64
	PositionCost       float64
	PreMargin          float64
	UseMargin          float64
	FrozenMargin       float64
	FrozenCash         float64
	FrozenCommission   float64
	CashIn             float64
	Commission         float64
	CloseProfit        float64
	PositionProfit     float64
	PreSettlementPrice float64
	SettlementPrice    float64
	TradingDay         string
	SettlementID       int
	OpenCost           float64
	ExchangeMargin     float64
	CombPosition       int
	CombLongFrozen     int
	CombShortFrozen    int
	CloseProfitByDate  float64
	CloseProfitByTrade float64
	TodayPosition      int
	MarginRateByMoney  float64
	MarginRateByVolume float64
	StrikeFrozen       int
	StrikeFrozenAmount float64
	AbandonFrozen      int
	ExchangeID         string
	YdStrikeFrozen     int
	InvestUnitID       string
	PositionCostOffset float64
	TasPosition        int
	TasPositionCost    float64
	InstrumentID       string
}

func NewCThostFtdcInvestorPositionField(p *C.CThostFtdcInvestorPositionField) *CThostFtdcInvestorPositionField {
	ret := new(CThostFtdcInvestorPositionField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.PosiDirection = byte(p.PosiDirection)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.PositionDate = byte(p.PositionDate)
	ret.YdPosition = int(p.YdPosition)
	ret.Position = int(p.Position)
	ret.LongFrozen = int(p.LongFrozen)
	ret.ShortFrozen = int(p.ShortFrozen)
	ret.LongFrozenAmount = goFloat64(p.LongFrozenAmount)
	ret.ShortFrozenAmount = goFloat64(p.ShortFrozenAmount)
	ret.OpenVolume = int(p.OpenVolume)
	ret.CloseVolume = int(p.CloseVolume)
	ret.OpenAmount = goFloat64(p.OpenAmount)
	ret.CloseAmount = goFloat64(p.CloseAmount)
	ret.PositionCost = goFloat64(p.PositionCost)
	ret.PreMargin = goFloat64(p.PreMargin)
	ret.UseMargin = goFloat64(p.UseMargin)
	ret.FrozenMargin = goFloat64(p.FrozenMargin)
	ret.FrozenCash = goFloat64(p.FrozenCash)
	ret.FrozenCommission = goFloat64(p.FrozenCommission)
	ret.CashIn = goFloat64(p.CashIn)
	ret.Commission = goFloat64(p.Commission)
	ret.CloseProfit = goFloat64(p.CloseProfit)
	ret.PositionProfit = goFloat64(p.PositionProfit)
	ret.PreSettlementPrice = goFloat64(p.PreSettlementPrice)
	ret.SettlementPrice = goFloat64(p.SettlementPrice)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.OpenCost = goFloat64(p.OpenCost)
	ret.ExchangeMargin = goFloat64(p.ExchangeMargin)
	ret.CombPosition = int(p.CombPosition)
	ret.CombLongFrozen = int(p.CombLongFrozen)
	ret.CombShortFrozen = int(p.CombShortFrozen)
	ret.CloseProfitByDate = goFloat64(p.CloseProfitByDate)
	ret.CloseProfitByTrade = goFloat64(p.CloseProfitByTrade)
	ret.TodayPosition = int(p.TodayPosition)
	ret.MarginRateByMoney = goFloat64(p.MarginRateByMoney)
	ret.MarginRateByVolume = goFloat64(p.MarginRateByVolume)
	ret.StrikeFrozen = int(p.StrikeFrozen)
	ret.StrikeFrozenAmount = goFloat64(p.StrikeFrozenAmount)
	ret.AbandonFrozen = int(p.AbandonFrozen)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.YdStrikeFrozen = int(p.YdStrikeFrozen)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.PositionCostOffset = goFloat64(p.PositionCostOffset)
	ret.TasPosition = int(p.TasPosition)
	ret.TasPositionCost = goFloat64(p.TasPositionCost)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcInvestorPositionField) CValue() *C.CThostFtdcInvestorPositionField {
	ptr := C.newCThostFtdcInvestorPositionField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.PosiDirection = C.char(s.PosiDirection)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.PositionDate = C.char(s.PositionDate)
	ptr.YdPosition = C.int(s.YdPosition)
	ptr.Position = C.int(s.Position)
	ptr.LongFrozen = C.int(s.LongFrozen)
	ptr.ShortFrozen = C.int(s.ShortFrozen)
	ptr.LongFrozenAmount = C.double(s.LongFrozenAmount)
	ptr.ShortFrozenAmount = C.double(s.ShortFrozenAmount)
	ptr.OpenVolume = C.int(s.OpenVolume)
	ptr.CloseVolume = C.int(s.CloseVolume)
	ptr.OpenAmount = C.double(s.OpenAmount)
	ptr.CloseAmount = C.double(s.CloseAmount)
	ptr.PositionCost = C.double(s.PositionCost)
	ptr.PreMargin = C.double(s.PreMargin)
	ptr.UseMargin = C.double(s.UseMargin)
	ptr.FrozenMargin = C.double(s.FrozenMargin)
	ptr.FrozenCash = C.double(s.FrozenCash)
	ptr.FrozenCommission = C.double(s.FrozenCommission)
	ptr.CashIn = C.double(s.CashIn)
	ptr.Commission = C.double(s.Commission)
	ptr.CloseProfit = C.double(s.CloseProfit)
	ptr.PositionProfit = C.double(s.PositionProfit)
	ptr.PreSettlementPrice = C.double(s.PreSettlementPrice)
	ptr.SettlementPrice = C.double(s.SettlementPrice)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	ptr.OpenCost = C.double(s.OpenCost)
	ptr.ExchangeMargin = C.double(s.ExchangeMargin)
	ptr.CombPosition = C.int(s.CombPosition)
	ptr.CombLongFrozen = C.int(s.CombLongFrozen)
	ptr.CombShortFrozen = C.int(s.CombShortFrozen)
	ptr.CloseProfitByDate = C.double(s.CloseProfitByDate)
	ptr.CloseProfitByTrade = C.double(s.CloseProfitByTrade)
	ptr.TodayPosition = C.int(s.TodayPosition)
	ptr.MarginRateByMoney = C.double(s.MarginRateByMoney)
	ptr.MarginRateByVolume = C.double(s.MarginRateByVolume)
	ptr.StrikeFrozen = C.int(s.StrikeFrozen)
	ptr.StrikeFrozenAmount = C.double(s.StrikeFrozenAmount)
	ptr.AbandonFrozen = C.int(s.AbandonFrozen)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ptr.YdStrikeFrozen = C.int(s.YdStrikeFrozen)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ptr.PositionCostOffset = C.double(s.PositionCostOffset)
	ptr.TasPosition = C.int(s.TasPosition)
	ptr.TasPositionCost = C.double(s.TasPositionCost)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcInstrumentMarginRateField struct {
	Reserve1                 string
	InvestorRange            byte
	BrokerID                 string
	InvestorID               string
	HedgeFlag                byte
	LongMarginRatioByMoney   float64
	LongMarginRatioByVolume  float64
	ShortMarginRatioByMoney  float64
	ShortMarginRatioByVolume float64
	IsRelative               int
	ExchangeID               string
	InvestUnitID             string
	InstrumentID             string
}

func NewCThostFtdcInstrumentMarginRateField(p *C.CThostFtdcInstrumentMarginRateField) *CThostFtdcInstrumentMarginRateField {
	ret := new(CThostFtdcInstrumentMarginRateField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.LongMarginRatioByMoney = goFloat64(p.LongMarginRatioByMoney)
	ret.LongMarginRatioByVolume = goFloat64(p.LongMarginRatioByVolume)
	ret.ShortMarginRatioByMoney = goFloat64(p.ShortMarginRatioByMoney)
	ret.ShortMarginRatioByVolume = goFloat64(p.ShortMarginRatioByVolume)
	ret.IsRelative = int(p.IsRelative)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcInstrumentMarginRateField) CValue() *C.CThostFtdcInstrumentMarginRateField {
	ptr := C.newCThostFtdcInstrumentMarginRateField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.LongMarginRatioByMoney = C.double(s.LongMarginRatioByMoney)
	ptr.LongMarginRatioByVolume = C.double(s.LongMarginRatioByVolume)
	ptr.ShortMarginRatioByMoney = C.double(s.ShortMarginRatioByMoney)
	ptr.ShortMarginRatioByVolume = C.double(s.ShortMarginRatioByVolume)
	ptr.IsRelative = C.int(s.IsRelative)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcInstrumentCommissionRateField struct {
	Reserve1                string
	InvestorRange           byte
	BrokerID                string
	InvestorID              string
	OpenRatioByMoney        float64
	OpenRatioByVolume       float64
	CloseRatioByMoney       float64
	CloseRatioByVolume      float64
	CloseTodayRatioByMoney  float64
	CloseTodayRatioByVolume float64
	ExchangeID              string
	BizType                 byte
	InvestUnitID            string
	InstrumentID            string
}

func NewCThostFtdcInstrumentCommissionRateField(p *C.CThostFtdcInstrumentCommissionRateField) *CThostFtdcInstrumentCommissionRateField {
	ret := new(CThostFtdcInstrumentCommissionRateField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.OpenRatioByMoney = goFloat64(p.OpenRatioByMoney)
	ret.OpenRatioByVolume = goFloat64(p.OpenRatioByVolume)
	ret.CloseRatioByMoney = goFloat64(p.CloseRatioByMoney)
	ret.CloseRatioByVolume = goFloat64(p.CloseRatioByVolume)
	ret.CloseTodayRatioByMoney = goFloat64(p.CloseTodayRatioByMoney)
	ret.CloseTodayRatioByVolume = goFloat64(p.CloseTodayRatioByVolume)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.BizType = byte(p.BizType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcInstrumentCommissionRateField) CValue() *C.CThostFtdcInstrumentCommissionRateField {
	ptr := C.newCThostFtdcInstrumentCommissionRateField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.OpenRatioByMoney = C.double(s.OpenRatioByMoney)
	ptr.OpenRatioByVolume = C.double(s.OpenRatioByVolume)
	ptr.CloseRatioByMoney = C.double(s.CloseRatioByMoney)
	ptr.CloseRatioByVolume = C.double(s.CloseRatioByVolume)
	ptr.CloseTodayRatioByMoney = C.double(s.CloseTodayRatioByMoney)
	ptr.CloseTodayRatioByVolume = C.double(s.CloseTodayRatioByVolume)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ptr.BizType = C.char(s.BizType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcDepthMarketDataField struct {
	TradingDay         string
	Reserve1           string
	ExchangeID         string
	Reserve2           string
	LastPrice          float64
	PreSettlementPrice float64
	PreClosePrice      float64
	PreOpenInterest    float64
	OpenPrice          float64
	HighestPrice       float64
	LowestPrice        float64
	Volume             int
	Turnover           float64
	OpenInterest       float64
	ClosePrice         float64
	SettlementPrice    float64
	UpperLimitPrice    float64
	LowerLimitPrice    float64
	PreDelta           float64
	CurrDelta          float64
	UpdateTime         string
	UpdateMillisec     int
	BidPrice1          float64
	BidVolume1         int
	AskPrice1          float64
	AskVolume1         int
	BidPrice2          float64
	BidVolume2         int
	AskPrice2          float64
	AskVolume2         int
	BidPrice3          float64
	BidVolume3         int
	AskPrice3          float64
	AskVolume3         int
	BidPrice4          float64
	BidVolume4         int
	AskPrice4          float64
	AskVolume4         int
	BidPrice5          float64
	BidVolume5         int
	AskPrice5          float64
	AskVolume5         int
	AveragePrice       float64
	ActionDay          string
	InstrumentID       string
	ExchangeInstID     string
	BandingUpperPrice  float64
	BandingLowerPrice  float64
}

func NewCThostFtdcDepthMarketDataField(p *C.CThostFtdcDepthMarketDataField) *CThostFtdcDepthMarketDataField {
	ret := new(CThostFtdcDepthMarketDataField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.LastPrice = goFloat64(p.LastPrice)
	ret.PreSettlementPrice = goFloat64(p.PreSettlementPrice)
	ret.PreClosePrice = goFloat64(p.PreClosePrice)
	ret.PreOpenInterest = goFloat64(p.PreOpenInterest)
	ret.OpenPrice = goFloat64(p.OpenPrice)
	ret.HighestPrice = goFloat64(p.HighestPrice)
	ret.LowestPrice = goFloat64(p.LowestPrice)
	ret.Volume = int(p.Volume)
	ret.Turnover = goFloat64(p.Turnover)
	ret.OpenInterest = goFloat64(p.OpenInterest)
	ret.ClosePrice = goFloat64(p.ClosePrice)
	ret.SettlementPrice = goFloat64(p.SettlementPrice)
	ret.UpperLimitPrice = goFloat64(p.UpperLimitPrice)
	ret.LowerLimitPrice = goFloat64(p.LowerLimitPrice)
	ret.PreDelta = goFloat64(p.PreDelta)
	ret.CurrDelta = goFloat64(p.CurrDelta)
	ret.UpdateTime = c2goStr(&p.UpdateTime[0], C.sizeof_TThostFtdcTimeType)
	ret.UpdateMillisec = int(p.UpdateMillisec)
	ret.BidPrice1 = goFloat64(p.BidPrice1)
	ret.BidVolume1 = int(p.BidVolume1)
	ret.AskPrice1 = goFloat64(p.AskPrice1)
	ret.AskVolume1 = int(p.AskVolume1)
	ret.BidPrice2 = goFloat64(p.BidPrice2)
	ret.BidVolume2 = int(p.BidVolume2)
	ret.AskPrice2 = goFloat64(p.AskPrice2)
	ret.AskVolume2 = int(p.AskVolume2)
	ret.BidPrice3 = goFloat64(p.BidPrice3)
	ret.BidVolume3 = int(p.BidVolume3)
	ret.AskPrice3 = goFloat64(p.AskPrice3)
	ret.AskVolume3 = int(p.AskVolume3)
	ret.BidPrice4 = goFloat64(p.BidPrice4)
	ret.BidVolume4 = int(p.BidVolume4)
	ret.AskPrice4 = goFloat64(p.AskPrice4)
	ret.AskVolume4 = int(p.AskVolume4)
	ret.BidPrice5 = goFloat64(p.BidPrice5)
	ret.BidVolume5 = int(p.BidVolume5)
	ret.AskPrice5 = goFloat64(p.AskPrice5)
	ret.AskVolume5 = int(p.AskVolume5)
	ret.AveragePrice = goFloat64(p.AveragePrice)
	ret.ActionDay = c2goStr(&p.ActionDay[0], C.sizeof_TThostFtdcDateType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.BandingUpperPrice = goFloat64(p.BandingUpperPrice)
	ret.BandingLowerPrice = goFloat64(p.BandingLowerPrice)
	return ret
}
func (s *CThostFtdcDepthMarketDataField) CValue() *C.CThostFtdcDepthMarketDataField {
	ptr := C.newCThostFtdcDepthMarketDataField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ptr.LastPrice = C.double(s.LastPrice)
	ptr.PreSettlementPrice = C.double(s.PreSettlementPrice)
	ptr.PreClosePrice = C.double(s.PreClosePrice)
	ptr.PreOpenInterest = C.double(s.PreOpenInterest)
	ptr.OpenPrice = C.double(s.OpenPrice)
	ptr.HighestPrice = C.double(s.HighestPrice)
	ptr.LowestPrice = C.double(s.LowestPrice)
	ptr.Volume = C.int(s.Volume)
	ptr.Turnover = C.double(s.Turnover)
	ptr.OpenInterest = C.double(s.OpenInterest)
	ptr.ClosePrice = C.double(s.ClosePrice)
	ptr.SettlementPrice = C.double(s.SettlementPrice)
	ptr.UpperLimitPrice = C.double(s.UpperLimitPrice)
	ptr.LowerLimitPrice = C.double(s.LowerLimitPrice)
	ptr.PreDelta = C.double(s.PreDelta)
	ptr.CurrDelta = C.double(s.CurrDelta)
	go2cStr(s.UpdateTime, &ptr.UpdateTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.UpdateMillisec = C.int(s.UpdateMillisec)
	ptr.BidPrice1 = C.double(s.BidPrice1)
	ptr.BidVolume1 = C.int(s.BidVolume1)
	ptr.AskPrice1 = C.double(s.AskPrice1)
	ptr.AskVolume1 = C.int(s.AskVolume1)
	ptr.BidPrice2 = C.double(s.BidPrice2)
	ptr.BidVolume2 = C.int(s.BidVolume2)
	ptr.AskPrice2 = C.double(s.AskPrice2)
	ptr.AskVolume2 = C.int(s.AskVolume2)
	ptr.BidPrice3 = C.double(s.BidPrice3)
	ptr.BidVolume3 = C.int(s.BidVolume3)
	ptr.AskPrice3 = C.double(s.AskPrice3)
	ptr.AskVolume3 = C.int(s.AskVolume3)
	ptr.BidPrice4 = C.double(s.BidPrice4)
	ptr.BidVolume4 = C.int(s.BidVolume4)
	ptr.AskPrice4 = C.double(s.AskPrice4)
	ptr.AskVolume4 = C.int(s.AskVolume4)
	ptr.BidPrice5 = C.double(s.BidPrice5)
	ptr.BidVolume5 = C.int(s.BidVolume5)
	ptr.AskPrice5 = C.double(s.AskPrice5)
	ptr.AskVolume5 = C.int(s.AskVolume5)
	ptr.AveragePrice = C.double(s.AveragePrice)
	go2cStr(s.ActionDay, &ptr.ActionDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ptr.BandingUpperPrice = C.double(s.BandingUpperPrice)
	ptr.BandingLowerPrice = C.double(s.BandingLowerPrice)
	return ptr
}

type CThostFtdcInstrumentTradingRightField struct {
	Reserve1      string
	InvestorRange byte
	BrokerID      string
	InvestorID    string
	TradingRight  byte
	InstrumentID  string
}

func NewCThostFtdcInstrumentTradingRightField(p *C.CThostFtdcInstrumentTradingRightField) *CThostFtdcInstrumentTradingRightField {
	ret := new(CThostFtdcInstrumentTradingRightField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.TradingRight = byte(p.TradingRight)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcInstrumentTradingRightField) CValue() *C.CThostFtdcInstrumentTradingRightField {
	ptr := C.newCThostFtdcInstrumentTradingRightField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.TradingRight = C.char(s.TradingRight)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcBrokerUserField struct {
	BrokerID    string
	UserID      string
	UserName    string
	UserType    byte
	IsActive    int
	IsUsingOTP  int
	IsAuthForce int
}

func NewCThostFtdcBrokerUserField(p *C.CThostFtdcBrokerUserField) *CThostFtdcBrokerUserField {
	ret := new(CThostFtdcBrokerUserField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.UserName = c2goStr(&p.UserName[0], C.sizeof_TThostFtdcUserNameType)
	ret.UserType = byte(p.UserType)
	ret.IsActive = int(p.IsActive)
	ret.IsUsingOTP = int(p.IsUsingOTP)
	ret.IsAuthForce = int(p.IsAuthForce)
	return ret
}
func (s *CThostFtdcBrokerUserField) CValue() *C.CThostFtdcBrokerUserField {
	ptr := C.newCThostFtdcBrokerUserField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.UserName, &ptr.UserName[0], C.sizeof_TThostFtdcUserNameType)
	ptr.UserType = C.char(s.UserType)
	ptr.IsActive = C.int(s.IsActive)
	ptr.IsUsingOTP = C.int(s.IsUsingOTP)
	ptr.IsAuthForce = C.int(s.IsAuthForce)
	return ptr
}

type CThostFtdcBrokerUserPasswordField struct {
	BrokerID       string
	UserID         string
	Password       string
	LastUpdateTime string
	LastLoginTime  string
	ExpireDate     string
	WeakExpireDate string
}

func NewCThostFtdcBrokerUserPasswordField(p *C.CThostFtdcBrokerUserPasswordField) *CThostFtdcBrokerUserPasswordField {
	ret := new(CThostFtdcBrokerUserPasswordField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.LastUpdateTime = c2goStr(&p.LastUpdateTime[0], C.sizeof_TThostFtdcDateTimeType)
	ret.LastLoginTime = c2goStr(&p.LastLoginTime[0], C.sizeof_TThostFtdcDateTimeType)
	ret.ExpireDate = c2goStr(&p.ExpireDate[0], C.sizeof_TThostFtdcDateType)
	ret.WeakExpireDate = c2goStr(&p.WeakExpireDate[0], C.sizeof_TThostFtdcDateType)
	return ret
}
func (s *CThostFtdcBrokerUserPasswordField) CValue() *C.CThostFtdcBrokerUserPasswordField {
	ptr := C.newCThostFtdcBrokerUserPasswordField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.LastUpdateTime, &ptr.LastUpdateTime[0], C.sizeof_TThostFtdcDateTimeType)
	go2cStr(s.LastLoginTime, &ptr.LastLoginTime[0], C.sizeof_TThostFtdcDateTimeType)
	go2cStr(s.ExpireDate, &ptr.ExpireDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.WeakExpireDate, &ptr.WeakExpireDate[0], C.sizeof_TThostFtdcDateType)
	return ptr
}

type CThostFtdcBrokerUserFunctionField struct {
	BrokerID           string
	UserID             string
	BrokerFunctionCode byte
}

func NewCThostFtdcBrokerUserFunctionField(p *C.CThostFtdcBrokerUserFunctionField) *CThostFtdcBrokerUserFunctionField {
	ret := new(CThostFtdcBrokerUserFunctionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.BrokerFunctionCode = byte(p.BrokerFunctionCode)
	return ret
}
func (s *CThostFtdcBrokerUserFunctionField) CValue() *C.CThostFtdcBrokerUserFunctionField {
	ptr := C.newCThostFtdcBrokerUserFunctionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.BrokerFunctionCode = C.char(s.BrokerFunctionCode)
	return ptr
}

type CThostFtdcTraderOfferField struct {
	ExchangeID               string
	TraderID                 string
	ParticipantID            string
	Password                 string
	InstallID                int
	OrderLocalID             string
	TraderConnectStatus      byte
	ConnectRequestDate       string
	ConnectRequestTime       string
	LastReportDate           string
	LastReportTime           string
	ConnectDate              string
	ConnectTime              string
	StartDate                string
	StartTime                string
	TradingDay               string
	BrokerID                 string
	MaxTradeID               string
	MaxOrderMessageReference string
}

func NewCThostFtdcTraderOfferField(p *C.CThostFtdcTraderOfferField) *CThostFtdcTraderOfferField {
	ret := new(CThostFtdcTraderOfferField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.InstallID = int(p.InstallID)
	ret.OrderLocalID = c2goStr(&p.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.TraderConnectStatus = byte(p.TraderConnectStatus)
	ret.ConnectRequestDate = c2goStr(&p.ConnectRequestDate[0], C.sizeof_TThostFtdcDateType)
	ret.ConnectRequestTime = c2goStr(&p.ConnectRequestTime[0], C.sizeof_TThostFtdcTimeType)
	ret.LastReportDate = c2goStr(&p.LastReportDate[0], C.sizeof_TThostFtdcDateType)
	ret.LastReportTime = c2goStr(&p.LastReportTime[0], C.sizeof_TThostFtdcTimeType)
	ret.ConnectDate = c2goStr(&p.ConnectDate[0], C.sizeof_TThostFtdcDateType)
	ret.ConnectTime = c2goStr(&p.ConnectTime[0], C.sizeof_TThostFtdcTimeType)
	ret.StartDate = c2goStr(&p.StartDate[0], C.sizeof_TThostFtdcDateType)
	ret.StartTime = c2goStr(&p.StartTime[0], C.sizeof_TThostFtdcTimeType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.MaxTradeID = c2goStr(&p.MaxTradeID[0], C.sizeof_TThostFtdcTradeIDType)
	ret.MaxOrderMessageReference = c2goStr(&p.MaxOrderMessageReference[0], C.sizeof_TThostFtdcReturnCodeType)
	return ret
}
func (s *CThostFtdcTraderOfferField) CValue() *C.CThostFtdcTraderOfferField {
	ptr := C.newCThostFtdcTraderOfferField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.OrderLocalID, &ptr.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ptr.TraderConnectStatus = C.char(s.TraderConnectStatus)
	go2cStr(s.ConnectRequestDate, &ptr.ConnectRequestDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ConnectRequestTime, &ptr.ConnectRequestTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.LastReportDate, &ptr.LastReportDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.LastReportTime, &ptr.LastReportTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.ConnectDate, &ptr.ConnectDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ConnectTime, &ptr.ConnectTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.StartDate, &ptr.StartDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.StartTime, &ptr.StartTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.MaxTradeID, &ptr.MaxTradeID[0], C.sizeof_TThostFtdcTradeIDType)
	go2cStr(s.MaxOrderMessageReference, &ptr.MaxOrderMessageReference[0], C.sizeof_TThostFtdcReturnCodeType)
	return ptr
}

type CThostFtdcSettlementInfoField struct {
	TradingDay   string
	SettlementID int
	BrokerID     string
	InvestorID   string
	SequenceNo   int
	Content      string
	AccountID    string
	CurrencyID   string
}

func NewCThostFtdcSettlementInfoField(p *C.CThostFtdcSettlementInfoField) *CThostFtdcSettlementInfoField {
	ret := new(CThostFtdcSettlementInfoField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.SequenceNo = int(p.SequenceNo)
	ret.Content = c2goStr(&p.Content[0], C.sizeof_TThostFtdcContentType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ret
}
func (s *CThostFtdcSettlementInfoField) CValue() *C.CThostFtdcSettlementInfoField {
	ptr := C.newCThostFtdcSettlementInfoField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.SequenceNo = C.int(s.SequenceNo)
	go2cStr(s.Content, &ptr.Content[0], C.sizeof_TThostFtdcContentType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ptr
}

type CThostFtdcInstrumentMarginRateAdjustField struct {
	Reserve1                 string
	InvestorRange            byte
	BrokerID                 string
	InvestorID               string
	HedgeFlag                byte
	LongMarginRatioByMoney   float64
	LongMarginRatioByVolume  float64
	ShortMarginRatioByMoney  float64
	ShortMarginRatioByVolume float64
	IsRelative               int
	InstrumentID             string
}

func NewCThostFtdcInstrumentMarginRateAdjustField(p *C.CThostFtdcInstrumentMarginRateAdjustField) *CThostFtdcInstrumentMarginRateAdjustField {
	ret := new(CThostFtdcInstrumentMarginRateAdjustField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.LongMarginRatioByMoney = goFloat64(p.LongMarginRatioByMoney)
	ret.LongMarginRatioByVolume = goFloat64(p.LongMarginRatioByVolume)
	ret.ShortMarginRatioByMoney = goFloat64(p.ShortMarginRatioByMoney)
	ret.ShortMarginRatioByVolume = goFloat64(p.ShortMarginRatioByVolume)
	ret.IsRelative = int(p.IsRelative)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcInstrumentMarginRateAdjustField) CValue() *C.CThostFtdcInstrumentMarginRateAdjustField {
	ptr := C.newCThostFtdcInstrumentMarginRateAdjustField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.LongMarginRatioByMoney = C.double(s.LongMarginRatioByMoney)
	ptr.LongMarginRatioByVolume = C.double(s.LongMarginRatioByVolume)
	ptr.ShortMarginRatioByMoney = C.double(s.ShortMarginRatioByMoney)
	ptr.ShortMarginRatioByVolume = C.double(s.ShortMarginRatioByVolume)
	ptr.IsRelative = C.int(s.IsRelative)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcExchangeMarginRateField struct {
	BrokerID                 string
	Reserve1                 string
	HedgeFlag                byte
	LongMarginRatioByMoney   float64
	LongMarginRatioByVolume  float64
	ShortMarginRatioByMoney  float64
	ShortMarginRatioByVolume float64
	ExchangeID               string
	InstrumentID             string
}

func NewCThostFtdcExchangeMarginRateField(p *C.CThostFtdcExchangeMarginRateField) *CThostFtdcExchangeMarginRateField {
	ret := new(CThostFtdcExchangeMarginRateField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.LongMarginRatioByMoney = goFloat64(p.LongMarginRatioByMoney)
	ret.LongMarginRatioByVolume = goFloat64(p.LongMarginRatioByVolume)
	ret.ShortMarginRatioByMoney = goFloat64(p.ShortMarginRatioByMoney)
	ret.ShortMarginRatioByVolume = goFloat64(p.ShortMarginRatioByVolume)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcExchangeMarginRateField) CValue() *C.CThostFtdcExchangeMarginRateField {
	ptr := C.newCThostFtdcExchangeMarginRateField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.LongMarginRatioByMoney = C.double(s.LongMarginRatioByMoney)
	ptr.LongMarginRatioByVolume = C.double(s.LongMarginRatioByVolume)
	ptr.ShortMarginRatioByMoney = C.double(s.ShortMarginRatioByMoney)
	ptr.ShortMarginRatioByVolume = C.double(s.ShortMarginRatioByVolume)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcExchangeMarginRateAdjustField struct {
	BrokerID                     string
	Reserve1                     string
	HedgeFlag                    byte
	LongMarginRatioByMoney       float64
	LongMarginRatioByVolume      float64
	ShortMarginRatioByMoney      float64
	ShortMarginRatioByVolume     float64
	ExchLongMarginRatioByMoney   float64
	ExchLongMarginRatioByVolume  float64
	ExchShortMarginRatioByMoney  float64
	ExchShortMarginRatioByVolume float64
	NoLongMarginRatioByMoney     float64
	NoLongMarginRatioByVolume    float64
	NoShortMarginRatioByMoney    float64
	NoShortMarginRatioByVolume   float64
	InstrumentID                 string
}

func NewCThostFtdcExchangeMarginRateAdjustField(p *C.CThostFtdcExchangeMarginRateAdjustField) *CThostFtdcExchangeMarginRateAdjustField {
	ret := new(CThostFtdcExchangeMarginRateAdjustField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.LongMarginRatioByMoney = goFloat64(p.LongMarginRatioByMoney)
	ret.LongMarginRatioByVolume = goFloat64(p.LongMarginRatioByVolume)
	ret.ShortMarginRatioByMoney = goFloat64(p.ShortMarginRatioByMoney)
	ret.ShortMarginRatioByVolume = goFloat64(p.ShortMarginRatioByVolume)
	ret.ExchLongMarginRatioByMoney = goFloat64(p.ExchLongMarginRatioByMoney)
	ret.ExchLongMarginRatioByVolume = goFloat64(p.ExchLongMarginRatioByVolume)
	ret.ExchShortMarginRatioByMoney = goFloat64(p.ExchShortMarginRatioByMoney)
	ret.ExchShortMarginRatioByVolume = goFloat64(p.ExchShortMarginRatioByVolume)
	ret.NoLongMarginRatioByMoney = goFloat64(p.NoLongMarginRatioByMoney)
	ret.NoLongMarginRatioByVolume = goFloat64(p.NoLongMarginRatioByVolume)
	ret.NoShortMarginRatioByMoney = goFloat64(p.NoShortMarginRatioByMoney)
	ret.NoShortMarginRatioByVolume = goFloat64(p.NoShortMarginRatioByVolume)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcExchangeMarginRateAdjustField) CValue() *C.CThostFtdcExchangeMarginRateAdjustField {
	ptr := C.newCThostFtdcExchangeMarginRateAdjustField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.LongMarginRatioByMoney = C.double(s.LongMarginRatioByMoney)
	ptr.LongMarginRatioByVolume = C.double(s.LongMarginRatioByVolume)
	ptr.ShortMarginRatioByMoney = C.double(s.ShortMarginRatioByMoney)
	ptr.ShortMarginRatioByVolume = C.double(s.ShortMarginRatioByVolume)
	ptr.ExchLongMarginRatioByMoney = C.double(s.ExchLongMarginRatioByMoney)
	ptr.ExchLongMarginRatioByVolume = C.double(s.ExchLongMarginRatioByVolume)
	ptr.ExchShortMarginRatioByMoney = C.double(s.ExchShortMarginRatioByMoney)
	ptr.ExchShortMarginRatioByVolume = C.double(s.ExchShortMarginRatioByVolume)
	ptr.NoLongMarginRatioByMoney = C.double(s.NoLongMarginRatioByMoney)
	ptr.NoLongMarginRatioByVolume = C.double(s.NoLongMarginRatioByVolume)
	ptr.NoShortMarginRatioByMoney = C.double(s.NoShortMarginRatioByMoney)
	ptr.NoShortMarginRatioByVolume = C.double(s.NoShortMarginRatioByVolume)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcExchangeRateField struct {
	BrokerID         string
	FromCurrencyID   string
	FromCurrencyUnit float64
	ToCurrencyID     string
	ExchangeRate     float64
}

func NewCThostFtdcExchangeRateField(p *C.CThostFtdcExchangeRateField) *CThostFtdcExchangeRateField {
	ret := new(CThostFtdcExchangeRateField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.FromCurrencyID = c2goStr(&p.FromCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.FromCurrencyUnit = goFloat64(p.FromCurrencyUnit)
	ret.ToCurrencyID = c2goStr(&p.ToCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.ExchangeRate = goFloat64(p.ExchangeRate)
	return ret
}
func (s *CThostFtdcExchangeRateField) CValue() *C.CThostFtdcExchangeRateField {
	ptr := C.newCThostFtdcExchangeRateField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.FromCurrencyID, &ptr.FromCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.FromCurrencyUnit = C.double(s.FromCurrencyUnit)
	go2cStr(s.ToCurrencyID, &ptr.ToCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.ExchangeRate = C.double(s.ExchangeRate)
	return ptr
}

type CThostFtdcSettlementRefField struct {
	TradingDay   string
	SettlementID int
}

func NewCThostFtdcSettlementRefField(p *C.CThostFtdcSettlementRefField) *CThostFtdcSettlementRefField {
	ret := new(CThostFtdcSettlementRefField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	return ret
}
func (s *CThostFtdcSettlementRefField) CValue() *C.CThostFtdcSettlementRefField {
	ptr := C.newCThostFtdcSettlementRefField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	return ptr
}

type CThostFtdcCurrentTimeField struct {
	CurrDate     string
	CurrTime     string
	CurrMillisec int
	ActionDay    string
}

func NewCThostFtdcCurrentTimeField(p *C.CThostFtdcCurrentTimeField) *CThostFtdcCurrentTimeField {
	ret := new(CThostFtdcCurrentTimeField)
	ret.CurrDate = c2goStr(&p.CurrDate[0], C.sizeof_TThostFtdcDateType)
	ret.CurrTime = c2goStr(&p.CurrTime[0], C.sizeof_TThostFtdcTimeType)
	ret.CurrMillisec = int(p.CurrMillisec)
	ret.ActionDay = c2goStr(&p.ActionDay[0], C.sizeof_TThostFtdcDateType)
	return ret
}
func (s *CThostFtdcCurrentTimeField) CValue() *C.CThostFtdcCurrentTimeField {
	ptr := C.newCThostFtdcCurrentTimeField()
	go2cStr(s.CurrDate, &ptr.CurrDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.CurrTime, &ptr.CurrTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.CurrMillisec = C.int(s.CurrMillisec)
	go2cStr(s.ActionDay, &ptr.ActionDay[0], C.sizeof_TThostFtdcDateType)
	return ptr
}

type CThostFtdcCommPhaseField struct {
	TradingDay  string
	CommPhaseNo int16
	SystemID    string
}

func NewCThostFtdcCommPhaseField(p *C.CThostFtdcCommPhaseField) *CThostFtdcCommPhaseField {
	ret := new(CThostFtdcCommPhaseField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.CommPhaseNo = int16(p.CommPhaseNo)
	ret.SystemID = c2goStr(&p.SystemID[0], C.sizeof_TThostFtdcSystemIDType)
	return ret
}
func (s *CThostFtdcCommPhaseField) CValue() *C.CThostFtdcCommPhaseField {
	ptr := C.newCThostFtdcCommPhaseField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.CommPhaseNo = C.short(s.CommPhaseNo)
	go2cStr(s.SystemID, &ptr.SystemID[0], C.sizeof_TThostFtdcSystemIDType)
	return ptr
}

type CThostFtdcLoginInfoField struct {
	FrontID              int
	SessionID            int
	BrokerID             string
	UserID               string
	LoginDate            string
	LoginTime            string
	Reserve1             string
	UserProductInfo      string
	InterfaceProductInfo string
	ProtocolInfo         string
	SystemName           string
	PasswordDeprecated   string
	MaxOrderRef          string
	SHFETime             string
	DCETime              string
	CZCETime             string
	FFEXTime             string
	MacAddress           string
	OneTimePassword      string
	INETime              string
	IsQryControl         int
	LoginRemark          string
	Password             string
	IPAddress            string
}

func NewCThostFtdcLoginInfoField(p *C.CThostFtdcLoginInfoField) *CThostFtdcLoginInfoField {
	ret := new(CThostFtdcLoginInfoField)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.LoginDate = c2goStr(&p.LoginDate[0], C.sizeof_TThostFtdcDateType)
	ret.LoginTime = c2goStr(&p.LoginTime[0], C.sizeof_TThostFtdcTimeType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.UserProductInfo = c2goStr(&p.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.InterfaceProductInfo = c2goStr(&p.InterfaceProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.ProtocolInfo = c2goStr(&p.ProtocolInfo[0], C.sizeof_TThostFtdcProtocolInfoType)
	ret.SystemName = c2goStr(&p.SystemName[0], C.sizeof_TThostFtdcSystemNameType)
	ret.PasswordDeprecated = c2goStr(&p.PasswordDeprecated[0], C.sizeof_TThostFtdcPasswordType)
	ret.MaxOrderRef = c2goStr(&p.MaxOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.SHFETime = c2goStr(&p.SHFETime[0], C.sizeof_TThostFtdcTimeType)
	ret.DCETime = c2goStr(&p.DCETime[0], C.sizeof_TThostFtdcTimeType)
	ret.CZCETime = c2goStr(&p.CZCETime[0], C.sizeof_TThostFtdcTimeType)
	ret.FFEXTime = c2goStr(&p.FFEXTime[0], C.sizeof_TThostFtdcTimeType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.OneTimePassword = c2goStr(&p.OneTimePassword[0], C.sizeof_TThostFtdcPasswordType)
	ret.INETime = c2goStr(&p.INETime[0], C.sizeof_TThostFtdcTimeType)
	ret.IsQryControl = int(p.IsQryControl)
	ret.LoginRemark = c2goStr(&p.LoginRemark[0], C.sizeof_TThostFtdcLoginRemarkType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcLoginInfoField) CValue() *C.CThostFtdcLoginInfoField {
	ptr := C.newCThostFtdcLoginInfoField()
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.LoginDate, &ptr.LoginDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.LoginTime, &ptr.LoginTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.UserProductInfo, &ptr.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.InterfaceProductInfo, &ptr.InterfaceProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.ProtocolInfo, &ptr.ProtocolInfo[0], C.sizeof_TThostFtdcProtocolInfoType)
	go2cStr(s.SystemName, &ptr.SystemName[0], C.sizeof_TThostFtdcSystemNameType)
	go2cStr(s.PasswordDeprecated, &ptr.PasswordDeprecated[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.MaxOrderRef, &ptr.MaxOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.SHFETime, &ptr.SHFETime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.DCETime, &ptr.DCETime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.CZCETime, &ptr.CZCETime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.FFEXTime, &ptr.FFEXTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.OneTimePassword, &ptr.OneTimePassword[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.INETime, &ptr.INETime[0], C.sizeof_TThostFtdcTimeType)
	ptr.IsQryControl = C.int(s.IsQryControl)
	go2cStr(s.LoginRemark, &ptr.LoginRemark[0], C.sizeof_TThostFtdcLoginRemarkType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcLogoutAllField struct {
	FrontID    int
	SessionID  int
	SystemName string
}

func NewCThostFtdcLogoutAllField(p *C.CThostFtdcLogoutAllField) *CThostFtdcLogoutAllField {
	ret := new(CThostFtdcLogoutAllField)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.SystemName = c2goStr(&p.SystemName[0], C.sizeof_TThostFtdcSystemNameType)
	return ret
}
func (s *CThostFtdcLogoutAllField) CValue() *C.CThostFtdcLogoutAllField {
	ptr := C.newCThostFtdcLogoutAllField()
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.SystemName, &ptr.SystemName[0], C.sizeof_TThostFtdcSystemNameType)
	return ptr
}

type CThostFtdcFrontStatusField struct {
	FrontID        int
	LastReportDate string
	LastReportTime string
	IsActive       int
}

func NewCThostFtdcFrontStatusField(p *C.CThostFtdcFrontStatusField) *CThostFtdcFrontStatusField {
	ret := new(CThostFtdcFrontStatusField)
	ret.FrontID = int(p.FrontID)
	ret.LastReportDate = c2goStr(&p.LastReportDate[0], C.sizeof_TThostFtdcDateType)
	ret.LastReportTime = c2goStr(&p.LastReportTime[0], C.sizeof_TThostFtdcTimeType)
	ret.IsActive = int(p.IsActive)
	return ret
}
func (s *CThostFtdcFrontStatusField) CValue() *C.CThostFtdcFrontStatusField {
	ptr := C.newCThostFtdcFrontStatusField()
	ptr.FrontID = C.int(s.FrontID)
	go2cStr(s.LastReportDate, &ptr.LastReportDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.LastReportTime, &ptr.LastReportTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.IsActive = C.int(s.IsActive)
	return ptr
}

type CThostFtdcUserPasswordUpdateField struct {
	BrokerID    string
	UserID      string
	OldPassword string
	NewPassword string
}

func NewCThostFtdcUserPasswordUpdateField(p *C.CThostFtdcUserPasswordUpdateField) *CThostFtdcUserPasswordUpdateField {
	ret := new(CThostFtdcUserPasswordUpdateField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.OldPassword = c2goStr(&p.OldPassword[0], C.sizeof_TThostFtdcPasswordType)
	ret.NewPassword = c2goStr(&p.NewPassword[0], C.sizeof_TThostFtdcPasswordType)
	return ret
}
func (s *CThostFtdcUserPasswordUpdateField) CValue() *C.CThostFtdcUserPasswordUpdateField {
	ptr := C.newCThostFtdcUserPasswordUpdateField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.OldPassword, &ptr.OldPassword[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.NewPassword, &ptr.NewPassword[0], C.sizeof_TThostFtdcPasswordType)
	return ptr
}

type CThostFtdcInputOrderField struct {
	BrokerID            string
	InvestorID          string
	Reserve1            string
	OrderRef            string
	UserID              string
	OrderPriceType      byte
	Direction           byte
	CombOffsetFlag      string
	CombHedgeFlag       string
	LimitPrice          float64
	VolumeTotalOriginal int
	TimeCondition       byte
	GTDDate             string
	VolumeCondition     byte
	MinVolume           int
	ContingentCondition byte
	StopPrice           float64
	ForceCloseReason    byte
	IsAutoSuspend       int
	BusinessUnit        string
	RequestID           int
	UserForceClose      int
	IsSwapOrder         int
	ExchangeID          string
	InvestUnitID        string
	AccountID           string
	CurrencyID          string
	ClientID            string
	Reserve2            string
	MacAddress          string
	InstrumentID        string
	IPAddress           string
}

func NewCThostFtdcInputOrderField(p *C.CThostFtdcInputOrderField) *CThostFtdcInputOrderField {
	ret := new(CThostFtdcInputOrderField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.OrderRef = c2goStr(&p.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.OrderPriceType = byte(p.OrderPriceType)
	ret.Direction = byte(p.Direction)
	ret.CombOffsetFlag = c2goStr(&p.CombOffsetFlag[0], C.sizeof_TThostFtdcCombOffsetFlagType)
	ret.CombHedgeFlag = c2goStr(&p.CombHedgeFlag[0], C.sizeof_TThostFtdcCombHedgeFlagType)
	ret.LimitPrice = goFloat64(p.LimitPrice)
	ret.VolumeTotalOriginal = int(p.VolumeTotalOriginal)
	ret.TimeCondition = byte(p.TimeCondition)
	ret.GTDDate = c2goStr(&p.GTDDate[0], C.sizeof_TThostFtdcDateType)
	ret.VolumeCondition = byte(p.VolumeCondition)
	ret.MinVolume = int(p.MinVolume)
	ret.ContingentCondition = byte(p.ContingentCondition)
	ret.StopPrice = goFloat64(p.StopPrice)
	ret.ForceCloseReason = byte(p.ForceCloseReason)
	ret.IsAutoSuspend = int(p.IsAutoSuspend)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.RequestID = int(p.RequestID)
	ret.UserForceClose = int(p.UserForceClose)
	ret.IsSwapOrder = int(p.IsSwapOrder)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcInputOrderField) CValue() *C.CThostFtdcInputOrderField {
	ptr := C.newCThostFtdcInputOrderField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.OrderRef, &ptr.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.OrderPriceType = C.char(s.OrderPriceType)
	ptr.Direction = C.char(s.Direction)
	go2cStr(s.CombOffsetFlag, &ptr.CombOffsetFlag[0], C.sizeof_TThostFtdcCombOffsetFlagType)
	go2cStr(s.CombHedgeFlag, &ptr.CombHedgeFlag[0], C.sizeof_TThostFtdcCombHedgeFlagType)
	ptr.LimitPrice = C.double(s.LimitPrice)
	ptr.VolumeTotalOriginal = C.int(s.VolumeTotalOriginal)
	ptr.TimeCondition = C.char(s.TimeCondition)
	go2cStr(s.GTDDate, &ptr.GTDDate[0], C.sizeof_TThostFtdcDateType)
	ptr.VolumeCondition = C.char(s.VolumeCondition)
	ptr.MinVolume = C.int(s.MinVolume)
	ptr.ContingentCondition = C.char(s.ContingentCondition)
	ptr.StopPrice = C.double(s.StopPrice)
	ptr.ForceCloseReason = C.char(s.ForceCloseReason)
	ptr.IsAutoSuspend = C.int(s.IsAutoSuspend)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.UserForceClose = C.int(s.UserForceClose)
	ptr.IsSwapOrder = C.int(s.IsSwapOrder)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcOrderField struct {
	BrokerID             string
	InvestorID           string
	Reserve1             string
	OrderRef             string
	UserID               string
	OrderPriceType       byte
	Direction            byte
	CombOffsetFlag       string
	CombHedgeFlag        string
	LimitPrice           float64
	VolumeTotalOriginal  int
	TimeCondition        byte
	GTDDate              string
	VolumeCondition      byte
	MinVolume            int
	ContingentCondition  byte
	StopPrice            float64
	ForceCloseReason     byte
	IsAutoSuspend        int
	BusinessUnit         string
	RequestID            int
	OrderLocalID         string
	ExchangeID           string
	ParticipantID        string
	ClientID             string
	Reserve2             string
	TraderID             string
	InstallID            int
	OrderSubmitStatus    byte
	NotifySequence       int
	TradingDay           string
	SettlementID         int
	OrderSysID           string
	OrderSource          byte
	OrderStatus          byte
	OrderType            byte
	VolumeTraded         int
	VolumeTotal          int
	InsertDate           string
	InsertTime           string
	ActiveTime           string
	SuspendTime          string
	UpdateTime           string
	CancelTime           string
	ActiveTraderID       string
	ClearingPartID       string
	SequenceNo           int
	FrontID              int
	SessionID            int
	UserProductInfo      string
	StatusMsg            string
	UserForceClose       int
	ActiveUserID         string
	BrokerOrderSeq       int
	RelativeOrderSysID   string
	ZCETotalTradedVolume int
	IsSwapOrder          int
	BranchID             string
	InvestUnitID         string
	AccountID            string
	CurrencyID           string
	Reserve3             string
	MacAddress           string
	InstrumentID         string
	ExchangeInstID       string
	IPAddress            string
}

func NewCThostFtdcOrderField(p *C.CThostFtdcOrderField) *CThostFtdcOrderField {
	ret := new(CThostFtdcOrderField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.OrderRef = c2goStr(&p.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.OrderPriceType = byte(p.OrderPriceType)
	ret.Direction = byte(p.Direction)
	ret.CombOffsetFlag = c2goStr(&p.CombOffsetFlag[0], C.sizeof_TThostFtdcCombOffsetFlagType)
	ret.CombHedgeFlag = c2goStr(&p.CombHedgeFlag[0], C.sizeof_TThostFtdcCombHedgeFlagType)
	ret.LimitPrice = goFloat64(p.LimitPrice)
	ret.VolumeTotalOriginal = int(p.VolumeTotalOriginal)
	ret.TimeCondition = byte(p.TimeCondition)
	ret.GTDDate = c2goStr(&p.GTDDate[0], C.sizeof_TThostFtdcDateType)
	ret.VolumeCondition = byte(p.VolumeCondition)
	ret.MinVolume = int(p.MinVolume)
	ret.ContingentCondition = byte(p.ContingentCondition)
	ret.StopPrice = goFloat64(p.StopPrice)
	ret.ForceCloseReason = byte(p.ForceCloseReason)
	ret.IsAutoSuspend = int(p.IsAutoSuspend)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.RequestID = int(p.RequestID)
	ret.OrderLocalID = c2goStr(&p.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.OrderSubmitStatus = byte(p.OrderSubmitStatus)
	ret.NotifySequence = int(p.NotifySequence)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.OrderSysID = c2goStr(&p.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.OrderSource = byte(p.OrderSource)
	ret.OrderStatus = byte(p.OrderStatus)
	ret.OrderType = byte(p.OrderType)
	ret.VolumeTraded = int(p.VolumeTraded)
	ret.VolumeTotal = int(p.VolumeTotal)
	ret.InsertDate = c2goStr(&p.InsertDate[0], C.sizeof_TThostFtdcDateType)
	ret.InsertTime = c2goStr(&p.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	ret.ActiveTime = c2goStr(&p.ActiveTime[0], C.sizeof_TThostFtdcTimeType)
	ret.SuspendTime = c2goStr(&p.SuspendTime[0], C.sizeof_TThostFtdcTimeType)
	ret.UpdateTime = c2goStr(&p.UpdateTime[0], C.sizeof_TThostFtdcTimeType)
	ret.CancelTime = c2goStr(&p.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	ret.ActiveTraderID = c2goStr(&p.ActiveTraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.ClearingPartID = c2goStr(&p.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.SequenceNo = int(p.SequenceNo)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.UserProductInfo = c2goStr(&p.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.StatusMsg = c2goStr(&p.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.UserForceClose = int(p.UserForceClose)
	ret.ActiveUserID = c2goStr(&p.ActiveUserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.BrokerOrderSeq = int(p.BrokerOrderSeq)
	ret.RelativeOrderSysID = c2goStr(&p.RelativeOrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ZCETotalTradedVolume = int(p.ZCETotalTradedVolume)
	ret.IsSwapOrder = int(p.IsSwapOrder)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.Reserve3 = c2goStr(&p.reserve3[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcOrderField) CValue() *C.CThostFtdcOrderField {
	ptr := C.newCThostFtdcOrderField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.OrderRef, &ptr.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.OrderPriceType = C.char(s.OrderPriceType)
	ptr.Direction = C.char(s.Direction)
	go2cStr(s.CombOffsetFlag, &ptr.CombOffsetFlag[0], C.sizeof_TThostFtdcCombOffsetFlagType)
	go2cStr(s.CombHedgeFlag, &ptr.CombHedgeFlag[0], C.sizeof_TThostFtdcCombHedgeFlagType)
	ptr.LimitPrice = C.double(s.LimitPrice)
	ptr.VolumeTotalOriginal = C.int(s.VolumeTotalOriginal)
	ptr.TimeCondition = C.char(s.TimeCondition)
	go2cStr(s.GTDDate, &ptr.GTDDate[0], C.sizeof_TThostFtdcDateType)
	ptr.VolumeCondition = C.char(s.VolumeCondition)
	ptr.MinVolume = C.int(s.MinVolume)
	ptr.ContingentCondition = C.char(s.ContingentCondition)
	ptr.StopPrice = C.double(s.StopPrice)
	ptr.ForceCloseReason = C.char(s.ForceCloseReason)
	ptr.IsAutoSuspend = C.int(s.IsAutoSuspend)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.RequestID = C.int(s.RequestID)
	go2cStr(s.OrderLocalID, &ptr.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.OrderSubmitStatus = C.char(s.OrderSubmitStatus)
	ptr.NotifySequence = C.int(s.NotifySequence)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	go2cStr(s.OrderSysID, &ptr.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ptr.OrderSource = C.char(s.OrderSource)
	ptr.OrderStatus = C.char(s.OrderStatus)
	ptr.OrderType = C.char(s.OrderType)
	ptr.VolumeTraded = C.int(s.VolumeTraded)
	ptr.VolumeTotal = C.int(s.VolumeTotal)
	go2cStr(s.InsertDate, &ptr.InsertDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.InsertTime, &ptr.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.ActiveTime, &ptr.ActiveTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.SuspendTime, &ptr.SuspendTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.UpdateTime, &ptr.UpdateTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.CancelTime, &ptr.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.ActiveTraderID, &ptr.ActiveTraderID[0], C.sizeof_TThostFtdcTraderIDType)
	go2cStr(s.ClearingPartID, &ptr.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ptr.SequenceNo = C.int(s.SequenceNo)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.UserProductInfo, &ptr.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.StatusMsg, &ptr.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ptr.UserForceClose = C.int(s.UserForceClose)
	go2cStr(s.ActiveUserID, &ptr.ActiveUserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.BrokerOrderSeq = C.int(s.BrokerOrderSeq)
	go2cStr(s.RelativeOrderSysID, &ptr.RelativeOrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ptr.ZCETotalTradedVolume = C.int(s.ZCETotalTradedVolume)
	ptr.IsSwapOrder = C.int(s.IsSwapOrder)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.Reserve3, &ptr.reserve3[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcExchangeOrderField struct {
	OrderPriceType      byte
	Direction           byte
	CombOffsetFlag      string
	CombHedgeFlag       string
	LimitPrice          float64
	VolumeTotalOriginal int
	TimeCondition       byte
	GTDDate             string
	VolumeCondition     byte
	MinVolume           int
	ContingentCondition byte
	StopPrice           float64
	ForceCloseReason    byte
	IsAutoSuspend       int
	BusinessUnit        string
	RequestID           int
	OrderLocalID        string
	ExchangeID          string
	ParticipantID       string
	ClientID            string
	Reserve1            string
	TraderID            string
	InstallID           int
	OrderSubmitStatus   byte
	NotifySequence      int
	TradingDay          string
	SettlementID        int
	OrderSysID          string
	OrderSource         byte
	OrderStatus         byte
	OrderType           byte
	VolumeTraded        int
	VolumeTotal         int
	InsertDate          string
	InsertTime          string
	ActiveTime          string
	SuspendTime         string
	UpdateTime          string
	CancelTime          string
	ActiveTraderID      string
	ClearingPartID      string
	SequenceNo          int
	BranchID            string
	Reserve2            string
	MacAddress          string
	ExchangeInstID      string
	IPAddress           string
}

func NewCThostFtdcExchangeOrderField(p *C.CThostFtdcExchangeOrderField) *CThostFtdcExchangeOrderField {
	ret := new(CThostFtdcExchangeOrderField)
	ret.OrderPriceType = byte(p.OrderPriceType)
	ret.Direction = byte(p.Direction)
	ret.CombOffsetFlag = c2goStr(&p.CombOffsetFlag[0], C.sizeof_TThostFtdcCombOffsetFlagType)
	ret.CombHedgeFlag = c2goStr(&p.CombHedgeFlag[0], C.sizeof_TThostFtdcCombHedgeFlagType)
	ret.LimitPrice = goFloat64(p.LimitPrice)
	ret.VolumeTotalOriginal = int(p.VolumeTotalOriginal)
	ret.TimeCondition = byte(p.TimeCondition)
	ret.GTDDate = c2goStr(&p.GTDDate[0], C.sizeof_TThostFtdcDateType)
	ret.VolumeCondition = byte(p.VolumeCondition)
	ret.MinVolume = int(p.MinVolume)
	ret.ContingentCondition = byte(p.ContingentCondition)
	ret.StopPrice = goFloat64(p.StopPrice)
	ret.ForceCloseReason = byte(p.ForceCloseReason)
	ret.IsAutoSuspend = int(p.IsAutoSuspend)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.RequestID = int(p.RequestID)
	ret.OrderLocalID = c2goStr(&p.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.OrderSubmitStatus = byte(p.OrderSubmitStatus)
	ret.NotifySequence = int(p.NotifySequence)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.OrderSysID = c2goStr(&p.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.OrderSource = byte(p.OrderSource)
	ret.OrderStatus = byte(p.OrderStatus)
	ret.OrderType = byte(p.OrderType)
	ret.VolumeTraded = int(p.VolumeTraded)
	ret.VolumeTotal = int(p.VolumeTotal)
	ret.InsertDate = c2goStr(&p.InsertDate[0], C.sizeof_TThostFtdcDateType)
	ret.InsertTime = c2goStr(&p.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	ret.ActiveTime = c2goStr(&p.ActiveTime[0], C.sizeof_TThostFtdcTimeType)
	ret.SuspendTime = c2goStr(&p.SuspendTime[0], C.sizeof_TThostFtdcTimeType)
	ret.UpdateTime = c2goStr(&p.UpdateTime[0], C.sizeof_TThostFtdcTimeType)
	ret.CancelTime = c2goStr(&p.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	ret.ActiveTraderID = c2goStr(&p.ActiveTraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.ClearingPartID = c2goStr(&p.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.SequenceNo = int(p.SequenceNo)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcExchangeOrderField) CValue() *C.CThostFtdcExchangeOrderField {
	ptr := C.newCThostFtdcExchangeOrderField()
	ptr.OrderPriceType = C.char(s.OrderPriceType)
	ptr.Direction = C.char(s.Direction)
	go2cStr(s.CombOffsetFlag, &ptr.CombOffsetFlag[0], C.sizeof_TThostFtdcCombOffsetFlagType)
	go2cStr(s.CombHedgeFlag, &ptr.CombHedgeFlag[0], C.sizeof_TThostFtdcCombHedgeFlagType)
	ptr.LimitPrice = C.double(s.LimitPrice)
	ptr.VolumeTotalOriginal = C.int(s.VolumeTotalOriginal)
	ptr.TimeCondition = C.char(s.TimeCondition)
	go2cStr(s.GTDDate, &ptr.GTDDate[0], C.sizeof_TThostFtdcDateType)
	ptr.VolumeCondition = C.char(s.VolumeCondition)
	ptr.MinVolume = C.int(s.MinVolume)
	ptr.ContingentCondition = C.char(s.ContingentCondition)
	ptr.StopPrice = C.double(s.StopPrice)
	ptr.ForceCloseReason = C.char(s.ForceCloseReason)
	ptr.IsAutoSuspend = C.int(s.IsAutoSuspend)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.RequestID = C.int(s.RequestID)
	go2cStr(s.OrderLocalID, &ptr.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.OrderSubmitStatus = C.char(s.OrderSubmitStatus)
	ptr.NotifySequence = C.int(s.NotifySequence)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	go2cStr(s.OrderSysID, &ptr.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ptr.OrderSource = C.char(s.OrderSource)
	ptr.OrderStatus = C.char(s.OrderStatus)
	ptr.OrderType = C.char(s.OrderType)
	ptr.VolumeTraded = C.int(s.VolumeTraded)
	ptr.VolumeTotal = C.int(s.VolumeTotal)
	go2cStr(s.InsertDate, &ptr.InsertDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.InsertTime, &ptr.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.ActiveTime, &ptr.ActiveTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.SuspendTime, &ptr.SuspendTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.UpdateTime, &ptr.UpdateTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.CancelTime, &ptr.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.ActiveTraderID, &ptr.ActiveTraderID[0], C.sizeof_TThostFtdcTraderIDType)
	go2cStr(s.ClearingPartID, &ptr.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ptr.SequenceNo = C.int(s.SequenceNo)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcExchangeOrderInsertErrorField struct {
	ExchangeID    string
	ParticipantID string
	TraderID      string
	InstallID     int
	OrderLocalID  string
	ErrorID       int
	ErrorMsg      string
}

func NewCThostFtdcExchangeOrderInsertErrorField(p *C.CThostFtdcExchangeOrderInsertErrorField) *CThostFtdcExchangeOrderInsertErrorField {
	ret := new(CThostFtdcExchangeOrderInsertErrorField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.OrderLocalID = c2goStr(&p.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ret
}
func (s *CThostFtdcExchangeOrderInsertErrorField) CValue() *C.CThostFtdcExchangeOrderInsertErrorField {
	ptr := C.newCThostFtdcExchangeOrderInsertErrorField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.OrderLocalID, &ptr.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ptr
}

type CThostFtdcInputOrderActionField struct {
	BrokerID       string
	InvestorID     string
	OrderActionRef int
	OrderRef       string
	RequestID      int
	FrontID        int
	SessionID      int
	ExchangeID     string
	OrderSysID     string
	ActionFlag     byte
	LimitPrice     float64
	VolumeChange   int
	UserID         string
	Reserve1       string
	InvestUnitID   string
	Reserve2       string
	MacAddress     string
	InstrumentID   string
	IPAddress      string
}

func NewCThostFtdcInputOrderActionField(p *C.CThostFtdcInputOrderActionField) *CThostFtdcInputOrderActionField {
	ret := new(CThostFtdcInputOrderActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.OrderActionRef = int(p.OrderActionRef)
	ret.OrderRef = c2goStr(&p.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.RequestID = int(p.RequestID)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.OrderSysID = c2goStr(&p.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ActionFlag = byte(p.ActionFlag)
	ret.LimitPrice = goFloat64(p.LimitPrice)
	ret.VolumeChange = int(p.VolumeChange)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcInputOrderActionField) CValue() *C.CThostFtdcInputOrderActionField {
	ptr := C.newCThostFtdcInputOrderActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.OrderActionRef = C.int(s.OrderActionRef)
	go2cStr(s.OrderRef, &ptr.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.OrderSysID, &ptr.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ptr.ActionFlag = C.char(s.ActionFlag)
	ptr.LimitPrice = C.double(s.LimitPrice)
	ptr.VolumeChange = C.int(s.VolumeChange)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcOrderActionField struct {
	BrokerID          string
	InvestorID        string
	OrderActionRef    int
	OrderRef          string
	RequestID         int
	FrontID           int
	SessionID         int
	ExchangeID        string
	OrderSysID        string
	ActionFlag        byte
	LimitPrice        float64
	VolumeChange      int
	ActionDate        string
	ActionTime        string
	TraderID          string
	InstallID         int
	OrderLocalID      string
	ActionLocalID     string
	ParticipantID     string
	ClientID          string
	BusinessUnit      string
	OrderActionStatus byte
	UserID            string
	StatusMsg         string
	Reserve1          string
	BranchID          string
	InvestUnitID      string
	Reserve2          string
	MacAddress        string
	InstrumentID      string
	IPAddress         string
}

func NewCThostFtdcOrderActionField(p *C.CThostFtdcOrderActionField) *CThostFtdcOrderActionField {
	ret := new(CThostFtdcOrderActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.OrderActionRef = int(p.OrderActionRef)
	ret.OrderRef = c2goStr(&p.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.RequestID = int(p.RequestID)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.OrderSysID = c2goStr(&p.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ActionFlag = byte(p.ActionFlag)
	ret.LimitPrice = goFloat64(p.LimitPrice)
	ret.VolumeChange = int(p.VolumeChange)
	ret.ActionDate = c2goStr(&p.ActionDate[0], C.sizeof_TThostFtdcDateType)
	ret.ActionTime = c2goStr(&p.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.OrderLocalID = c2goStr(&p.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ActionLocalID = c2goStr(&p.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.OrderActionStatus = byte(p.OrderActionStatus)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.StatusMsg = c2goStr(&p.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcOrderActionField) CValue() *C.CThostFtdcOrderActionField {
	ptr := C.newCThostFtdcOrderActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.OrderActionRef = C.int(s.OrderActionRef)
	go2cStr(s.OrderRef, &ptr.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.OrderSysID, &ptr.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ptr.ActionFlag = C.char(s.ActionFlag)
	ptr.LimitPrice = C.double(s.LimitPrice)
	ptr.VolumeChange = C.int(s.VolumeChange)
	go2cStr(s.ActionDate, &ptr.ActionDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ActionTime, &ptr.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.OrderLocalID, &ptr.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ActionLocalID, &ptr.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.OrderActionStatus = C.char(s.OrderActionStatus)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.StatusMsg, &ptr.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcExchangeOrderActionField struct {
	ExchangeID        string
	OrderSysID        string
	ActionFlag        byte
	LimitPrice        float64
	VolumeChange      int
	ActionDate        string
	ActionTime        string
	TraderID          string
	InstallID         int
	OrderLocalID      string
	ActionLocalID     string
	ParticipantID     string
	ClientID          string
	BusinessUnit      string
	OrderActionStatus byte
	UserID            string
	BranchID          string
	Reserve1          string
	MacAddress        string
	IPAddress         string
}

func NewCThostFtdcExchangeOrderActionField(p *C.CThostFtdcExchangeOrderActionField) *CThostFtdcExchangeOrderActionField {
	ret := new(CThostFtdcExchangeOrderActionField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.OrderSysID = c2goStr(&p.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ActionFlag = byte(p.ActionFlag)
	ret.LimitPrice = goFloat64(p.LimitPrice)
	ret.VolumeChange = int(p.VolumeChange)
	ret.ActionDate = c2goStr(&p.ActionDate[0], C.sizeof_TThostFtdcDateType)
	ret.ActionTime = c2goStr(&p.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.OrderLocalID = c2goStr(&p.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ActionLocalID = c2goStr(&p.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.OrderActionStatus = byte(p.OrderActionStatus)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcExchangeOrderActionField) CValue() *C.CThostFtdcExchangeOrderActionField {
	ptr := C.newCThostFtdcExchangeOrderActionField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.OrderSysID, &ptr.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ptr.ActionFlag = C.char(s.ActionFlag)
	ptr.LimitPrice = C.double(s.LimitPrice)
	ptr.VolumeChange = C.int(s.VolumeChange)
	go2cStr(s.ActionDate, &ptr.ActionDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ActionTime, &ptr.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.OrderLocalID, &ptr.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ActionLocalID, &ptr.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.OrderActionStatus = C.char(s.OrderActionStatus)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcExchangeOrderActionErrorField struct {
	ExchangeID    string
	OrderSysID    string
	TraderID      string
	InstallID     int
	OrderLocalID  string
	ActionLocalID string
	ErrorID       int
	ErrorMsg      string
}

func NewCThostFtdcExchangeOrderActionErrorField(p *C.CThostFtdcExchangeOrderActionErrorField) *CThostFtdcExchangeOrderActionErrorField {
	ret := new(CThostFtdcExchangeOrderActionErrorField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.OrderSysID = c2goStr(&p.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.OrderLocalID = c2goStr(&p.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ActionLocalID = c2goStr(&p.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ret
}
func (s *CThostFtdcExchangeOrderActionErrorField) CValue() *C.CThostFtdcExchangeOrderActionErrorField {
	ptr := C.newCThostFtdcExchangeOrderActionErrorField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.OrderSysID, &ptr.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.OrderLocalID, &ptr.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ActionLocalID, &ptr.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ptr
}

type CThostFtdcExchangeTradeField struct {
	ExchangeID     string
	TradeID        string
	Direction      byte
	OrderSysID     string
	ParticipantID  string
	ClientID       string
	TradingRole    byte
	Reserve1       string
	OffsetFlag     byte
	HedgeFlag      byte
	Price          float64
	Volume         int
	TradeDate      string
	TradeTime      string
	TradeType      byte
	PriceSource    byte
	TraderID       string
	OrderLocalID   string
	ClearingPartID string
	BusinessUnit   string
	SequenceNo     int
	TradeSource    byte
	ExchangeInstID string
}

func NewCThostFtdcExchangeTradeField(p *C.CThostFtdcExchangeTradeField) *CThostFtdcExchangeTradeField {
	ret := new(CThostFtdcExchangeTradeField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.TradeID = c2goStr(&p.TradeID[0], C.sizeof_TThostFtdcTradeIDType)
	ret.Direction = byte(p.Direction)
	ret.OrderSysID = c2goStr(&p.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.TradingRole = byte(p.TradingRole)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.OffsetFlag = byte(p.OffsetFlag)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.Price = goFloat64(p.Price)
	ret.Volume = int(p.Volume)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTimeType)
	ret.TradeType = byte(p.TradeType)
	ret.PriceSource = byte(p.PriceSource)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.OrderLocalID = c2goStr(&p.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ClearingPartID = c2goStr(&p.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.SequenceNo = int(p.SequenceNo)
	ret.TradeSource = byte(p.TradeSource)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ret
}
func (s *CThostFtdcExchangeTradeField) CValue() *C.CThostFtdcExchangeTradeField {
	ptr := C.newCThostFtdcExchangeTradeField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.TradeID, &ptr.TradeID[0], C.sizeof_TThostFtdcTradeIDType)
	ptr.Direction = C.char(s.Direction)
	go2cStr(s.OrderSysID, &ptr.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ptr.TradingRole = C.char(s.TradingRole)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ptr.OffsetFlag = C.char(s.OffsetFlag)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.Price = C.double(s.Price)
	ptr.Volume = C.int(s.Volume)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.TradeType = C.char(s.TradeType)
	ptr.PriceSource = C.char(s.PriceSource)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	go2cStr(s.OrderLocalID, &ptr.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ClearingPartID, &ptr.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.SequenceNo = C.int(s.SequenceNo)
	ptr.TradeSource = C.char(s.TradeSource)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ptr
}

type CThostFtdcTradeField struct {
	BrokerID       string
	InvestorID     string
	Reserve1       string
	OrderRef       string
	UserID         string
	ExchangeID     string
	TradeID        string
	Direction      byte
	OrderSysID     string
	ParticipantID  string
	ClientID       string
	TradingRole    byte
	Reserve2       string
	OffsetFlag     byte
	HedgeFlag      byte
	Price          float64
	Volume         int
	TradeDate      string
	TradeTime      string
	TradeType      byte
	PriceSource    byte
	TraderID       string
	OrderLocalID   string
	ClearingPartID string
	BusinessUnit   string
	SequenceNo     int
	TradingDay     string
	SettlementID   int
	BrokerOrderSeq int
	TradeSource    byte
	InvestUnitID   string
	InstrumentID   string
	ExchangeInstID string
}

func NewCThostFtdcTradeField(p *C.CThostFtdcTradeField) *CThostFtdcTradeField {
	ret := new(CThostFtdcTradeField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.OrderRef = c2goStr(&p.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.TradeID = c2goStr(&p.TradeID[0], C.sizeof_TThostFtdcTradeIDType)
	ret.Direction = byte(p.Direction)
	ret.OrderSysID = c2goStr(&p.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.TradingRole = byte(p.TradingRole)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.OffsetFlag = byte(p.OffsetFlag)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.Price = goFloat64(p.Price)
	ret.Volume = int(p.Volume)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTimeType)
	ret.TradeType = byte(p.TradeType)
	ret.PriceSource = byte(p.PriceSource)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.OrderLocalID = c2goStr(&p.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ClearingPartID = c2goStr(&p.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.SequenceNo = int(p.SequenceNo)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.BrokerOrderSeq = int(p.BrokerOrderSeq)
	ret.TradeSource = byte(p.TradeSource)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ret
}
func (s *CThostFtdcTradeField) CValue() *C.CThostFtdcTradeField {
	ptr := C.newCThostFtdcTradeField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.OrderRef, &ptr.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.TradeID, &ptr.TradeID[0], C.sizeof_TThostFtdcTradeIDType)
	ptr.Direction = C.char(s.Direction)
	go2cStr(s.OrderSysID, &ptr.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ptr.TradingRole = C.char(s.TradingRole)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ptr.OffsetFlag = C.char(s.OffsetFlag)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.Price = C.double(s.Price)
	ptr.Volume = C.int(s.Volume)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.TradeType = C.char(s.TradeType)
	ptr.PriceSource = C.char(s.PriceSource)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	go2cStr(s.OrderLocalID, &ptr.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ClearingPartID, &ptr.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.SequenceNo = C.int(s.SequenceNo)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	ptr.BrokerOrderSeq = C.int(s.BrokerOrderSeq)
	ptr.TradeSource = C.char(s.TradeSource)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ptr
}

type CThostFtdcUserSessionField struct {
	FrontID              int
	SessionID            int
	BrokerID             string
	UserID               string
	LoginDate            string
	LoginTime            string
	Reserve1             string
	UserProductInfo      string
	InterfaceProductInfo string
	ProtocolInfo         string
	MacAddress           string
	LoginRemark          string
	IPAddress            string
}

func NewCThostFtdcUserSessionField(p *C.CThostFtdcUserSessionField) *CThostFtdcUserSessionField {
	ret := new(CThostFtdcUserSessionField)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.LoginDate = c2goStr(&p.LoginDate[0], C.sizeof_TThostFtdcDateType)
	ret.LoginTime = c2goStr(&p.LoginTime[0], C.sizeof_TThostFtdcTimeType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.UserProductInfo = c2goStr(&p.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.InterfaceProductInfo = c2goStr(&p.InterfaceProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.ProtocolInfo = c2goStr(&p.ProtocolInfo[0], C.sizeof_TThostFtdcProtocolInfoType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.LoginRemark = c2goStr(&p.LoginRemark[0], C.sizeof_TThostFtdcLoginRemarkType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcUserSessionField) CValue() *C.CThostFtdcUserSessionField {
	ptr := C.newCThostFtdcUserSessionField()
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.LoginDate, &ptr.LoginDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.LoginTime, &ptr.LoginTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.UserProductInfo, &ptr.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.InterfaceProductInfo, &ptr.InterfaceProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.ProtocolInfo, &ptr.ProtocolInfo[0], C.sizeof_TThostFtdcProtocolInfoType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.LoginRemark, &ptr.LoginRemark[0], C.sizeof_TThostFtdcLoginRemarkType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryMaxOrderVolumeField struct {
	BrokerID     string
	InvestorID   string
	Reserve1     string
	Direction    byte
	OffsetFlag   byte
	HedgeFlag    byte
	MaxVolume    int
	ExchangeID   string
	InvestUnitID string
	InstrumentID string
}

func NewCThostFtdcQryMaxOrderVolumeField(p *C.CThostFtdcQryMaxOrderVolumeField) *CThostFtdcQryMaxOrderVolumeField {
	ret := new(CThostFtdcQryMaxOrderVolumeField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.Direction = byte(p.Direction)
	ret.OffsetFlag = byte(p.OffsetFlag)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.MaxVolume = int(p.MaxVolume)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryMaxOrderVolumeField) CValue() *C.CThostFtdcQryMaxOrderVolumeField {
	ptr := C.newCThostFtdcQryMaxOrderVolumeField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.Direction = C.char(s.Direction)
	ptr.OffsetFlag = C.char(s.OffsetFlag)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.MaxVolume = C.int(s.MaxVolume)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcSettlementInfoConfirmField struct {
	BrokerID     string
	InvestorID   string
	ConfirmDate  string
	ConfirmTime  string
	SettlementID int
	AccountID    string
	CurrencyID   string
}

func NewCThostFtdcSettlementInfoConfirmField(p *C.CThostFtdcSettlementInfoConfirmField) *CThostFtdcSettlementInfoConfirmField {
	ret := new(CThostFtdcSettlementInfoConfirmField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.ConfirmDate = c2goStr(&p.ConfirmDate[0], C.sizeof_TThostFtdcDateType)
	ret.ConfirmTime = c2goStr(&p.ConfirmTime[0], C.sizeof_TThostFtdcTimeType)
	ret.SettlementID = int(p.SettlementID)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ret
}
func (s *CThostFtdcSettlementInfoConfirmField) CValue() *C.CThostFtdcSettlementInfoConfirmField {
	ptr := C.newCThostFtdcSettlementInfoConfirmField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.ConfirmDate, &ptr.ConfirmDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ConfirmTime, &ptr.ConfirmTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.SettlementID = C.int(s.SettlementID)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ptr
}

type CThostFtdcSyncDepositField struct {
	DepositSeqNo    string
	BrokerID        string
	InvestorID      string
	Deposit         float64
	IsForce         int
	CurrencyID      string
	IsFromSopt      int
	TradingPassword string
}

func NewCThostFtdcSyncDepositField(p *C.CThostFtdcSyncDepositField) *CThostFtdcSyncDepositField {
	ret := new(CThostFtdcSyncDepositField)
	ret.DepositSeqNo = c2goStr(&p.DepositSeqNo[0], C.sizeof_TThostFtdcDepositSeqNoType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Deposit = goFloat64(p.Deposit)
	ret.IsForce = int(p.IsForce)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.IsFromSopt = int(p.IsFromSopt)
	ret.TradingPassword = c2goStr(&p.TradingPassword[0], C.sizeof_TThostFtdcPasswordType)
	return ret
}
func (s *CThostFtdcSyncDepositField) CValue() *C.CThostFtdcSyncDepositField {
	ptr := C.newCThostFtdcSyncDepositField()
	go2cStr(s.DepositSeqNo, &ptr.DepositSeqNo[0], C.sizeof_TThostFtdcDepositSeqNoType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.Deposit = C.double(s.Deposit)
	ptr.IsForce = C.int(s.IsForce)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.IsFromSopt = C.int(s.IsFromSopt)
	go2cStr(s.TradingPassword, &ptr.TradingPassword[0], C.sizeof_TThostFtdcPasswordType)
	return ptr
}

type CThostFtdcSyncFundMortgageField struct {
	MortgageSeqNo  string
	BrokerID       string
	InvestorID     string
	FromCurrencyID string
	MortgageAmount float64
	ToCurrencyID   string
}

func NewCThostFtdcSyncFundMortgageField(p *C.CThostFtdcSyncFundMortgageField) *CThostFtdcSyncFundMortgageField {
	ret := new(CThostFtdcSyncFundMortgageField)
	ret.MortgageSeqNo = c2goStr(&p.MortgageSeqNo[0], C.sizeof_TThostFtdcDepositSeqNoType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.FromCurrencyID = c2goStr(&p.FromCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.MortgageAmount = goFloat64(p.MortgageAmount)
	ret.ToCurrencyID = c2goStr(&p.ToCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ret
}
func (s *CThostFtdcSyncFundMortgageField) CValue() *C.CThostFtdcSyncFundMortgageField {
	ptr := C.newCThostFtdcSyncFundMortgageField()
	go2cStr(s.MortgageSeqNo, &ptr.MortgageSeqNo[0], C.sizeof_TThostFtdcDepositSeqNoType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.FromCurrencyID, &ptr.FromCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.MortgageAmount = C.double(s.MortgageAmount)
	go2cStr(s.ToCurrencyID, &ptr.ToCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ptr
}

type CThostFtdcBrokerSyncField struct {
	BrokerID string
}

func NewCThostFtdcBrokerSyncField(p *C.CThostFtdcBrokerSyncField) *CThostFtdcBrokerSyncField {
	ret := new(CThostFtdcBrokerSyncField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	return ret
}
func (s *CThostFtdcBrokerSyncField) CValue() *C.CThostFtdcBrokerSyncField {
	ptr := C.newCThostFtdcBrokerSyncField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	return ptr
}

type CThostFtdcSyncingInvestorField struct {
	InvestorID         string
	BrokerID           string
	InvestorGroupID    string
	InvestorName       string
	IdentifiedCardType byte
	IdentifiedCardNo   string
	IsActive           int
	Telephone          string
	Address            string
	OpenDate           string
	Mobile             string
	CommModelID        string
	MarginModelID      string
}

func NewCThostFtdcSyncingInvestorField(p *C.CThostFtdcSyncingInvestorField) *CThostFtdcSyncingInvestorField {
	ret := new(CThostFtdcSyncingInvestorField)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorGroupID = c2goStr(&p.InvestorGroupID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.InvestorName = c2goStr(&p.InvestorName[0], C.sizeof_TThostFtdcPartyNameType)
	ret.IdentifiedCardType = byte(p.IdentifiedCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.IsActive = int(p.IsActive)
	ret.Telephone = c2goStr(&p.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	ret.Address = c2goStr(&p.Address[0], C.sizeof_TThostFtdcAddressType)
	ret.OpenDate = c2goStr(&p.OpenDate[0], C.sizeof_TThostFtdcDateType)
	ret.Mobile = c2goStr(&p.Mobile[0], C.sizeof_TThostFtdcMobileType)
	ret.CommModelID = c2goStr(&p.CommModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.MarginModelID = c2goStr(&p.MarginModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ret
}
func (s *CThostFtdcSyncingInvestorField) CValue() *C.CThostFtdcSyncingInvestorField {
	ptr := C.newCThostFtdcSyncingInvestorField()
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorGroupID, &ptr.InvestorGroupID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.InvestorName, &ptr.InvestorName[0], C.sizeof_TThostFtdcPartyNameType)
	ptr.IdentifiedCardType = C.char(s.IdentifiedCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.IsActive = C.int(s.IsActive)
	go2cStr(s.Telephone, &ptr.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	go2cStr(s.Address, &ptr.Address[0], C.sizeof_TThostFtdcAddressType)
	go2cStr(s.OpenDate, &ptr.OpenDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.Mobile, &ptr.Mobile[0], C.sizeof_TThostFtdcMobileType)
	go2cStr(s.CommModelID, &ptr.CommModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.MarginModelID, &ptr.MarginModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ptr
}

type CThostFtdcSyncingTradingCodeField struct {
	InvestorID   string
	BrokerID     string
	ExchangeID   string
	ClientID     string
	IsActive     int
	ClientIDType byte
}

func NewCThostFtdcSyncingTradingCodeField(p *C.CThostFtdcSyncingTradingCodeField) *CThostFtdcSyncingTradingCodeField {
	ret := new(CThostFtdcSyncingTradingCodeField)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.IsActive = int(p.IsActive)
	ret.ClientIDType = byte(p.ClientIDType)
	return ret
}
func (s *CThostFtdcSyncingTradingCodeField) CValue() *C.CThostFtdcSyncingTradingCodeField {
	ptr := C.newCThostFtdcSyncingTradingCodeField()
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ptr.IsActive = C.int(s.IsActive)
	ptr.ClientIDType = C.char(s.ClientIDType)
	return ptr
}

type CThostFtdcSyncingInvestorGroupField struct {
	BrokerID          string
	InvestorGroupID   string
	InvestorGroupName string
}

func NewCThostFtdcSyncingInvestorGroupField(p *C.CThostFtdcSyncingInvestorGroupField) *CThostFtdcSyncingInvestorGroupField {
	ret := new(CThostFtdcSyncingInvestorGroupField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorGroupID = c2goStr(&p.InvestorGroupID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.InvestorGroupName = c2goStr(&p.InvestorGroupName[0], C.sizeof_TThostFtdcInvestorGroupNameType)
	return ret
}
func (s *CThostFtdcSyncingInvestorGroupField) CValue() *C.CThostFtdcSyncingInvestorGroupField {
	ptr := C.newCThostFtdcSyncingInvestorGroupField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorGroupID, &ptr.InvestorGroupID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.InvestorGroupName, &ptr.InvestorGroupName[0], C.sizeof_TThostFtdcInvestorGroupNameType)
	return ptr
}

type CThostFtdcSyncingTradingAccountField struct {
	BrokerID                       string
	AccountID                      string
	PreMortgage                    float64
	PreCredit                      float64
	PreDeposit                     float64
	PreBalance                     float64
	PreMargin                      float64
	InterestBase                   float64
	Interest                       float64
	Deposit                        float64
	Withdraw                       float64
	FrozenMargin                   float64
	FrozenCash                     float64
	FrozenCommission               float64
	CurrMargin                     float64
	CashIn                         float64
	Commission                     float64
	CloseProfit                    float64
	PositionProfit                 float64
	Balance                        float64
	Available                      float64
	WithdrawQuota                  float64
	Reserve                        float64
	TradingDay                     string
	SettlementID                   int
	Credit                         float64
	Mortgage                       float64
	ExchangeMargin                 float64
	DeliveryMargin                 float64
	ExchangeDeliveryMargin         float64
	ReserveBalance                 float64
	CurrencyID                     string
	PreFundMortgageIn              float64
	PreFundMortgageOut             float64
	FundMortgageIn                 float64
	FundMortgageOut                float64
	FundMortgageAvailable          float64
	MortgageableFund               float64
	SpecProductMargin              float64
	SpecProductFrozenMargin        float64
	SpecProductCommission          float64
	SpecProductFrozenCommission    float64
	SpecProductPositionProfit      float64
	SpecProductCloseProfit         float64
	SpecProductPositionProfitByAlg float64
	SpecProductExchangeMargin      float64
	FrozenSwap                     float64
	RemainSwap                     float64
}

func NewCThostFtdcSyncingTradingAccountField(p *C.CThostFtdcSyncingTradingAccountField) *CThostFtdcSyncingTradingAccountField {
	ret := new(CThostFtdcSyncingTradingAccountField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.PreMortgage = goFloat64(p.PreMortgage)
	ret.PreCredit = goFloat64(p.PreCredit)
	ret.PreDeposit = goFloat64(p.PreDeposit)
	ret.PreBalance = goFloat64(p.PreBalance)
	ret.PreMargin = goFloat64(p.PreMargin)
	ret.InterestBase = goFloat64(p.InterestBase)
	ret.Interest = goFloat64(p.Interest)
	ret.Deposit = goFloat64(p.Deposit)
	ret.Withdraw = goFloat64(p.Withdraw)
	ret.FrozenMargin = goFloat64(p.FrozenMargin)
	ret.FrozenCash = goFloat64(p.FrozenCash)
	ret.FrozenCommission = goFloat64(p.FrozenCommission)
	ret.CurrMargin = goFloat64(p.CurrMargin)
	ret.CashIn = goFloat64(p.CashIn)
	ret.Commission = goFloat64(p.Commission)
	ret.CloseProfit = goFloat64(p.CloseProfit)
	ret.PositionProfit = goFloat64(p.PositionProfit)
	ret.Balance = goFloat64(p.Balance)
	ret.Available = goFloat64(p.Available)
	ret.WithdrawQuota = goFloat64(p.WithdrawQuota)
	ret.Reserve = goFloat64(p.Reserve)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.Credit = goFloat64(p.Credit)
	ret.Mortgage = goFloat64(p.Mortgage)
	ret.ExchangeMargin = goFloat64(p.ExchangeMargin)
	ret.DeliveryMargin = goFloat64(p.DeliveryMargin)
	ret.ExchangeDeliveryMargin = goFloat64(p.ExchangeDeliveryMargin)
	ret.ReserveBalance = goFloat64(p.ReserveBalance)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.PreFundMortgageIn = goFloat64(p.PreFundMortgageIn)
	ret.PreFundMortgageOut = goFloat64(p.PreFundMortgageOut)
	ret.FundMortgageIn = goFloat64(p.FundMortgageIn)
	ret.FundMortgageOut = goFloat64(p.FundMortgageOut)
	ret.FundMortgageAvailable = goFloat64(p.FundMortgageAvailable)
	ret.MortgageableFund = goFloat64(p.MortgageableFund)
	ret.SpecProductMargin = goFloat64(p.SpecProductMargin)
	ret.SpecProductFrozenMargin = goFloat64(p.SpecProductFrozenMargin)
	ret.SpecProductCommission = goFloat64(p.SpecProductCommission)
	ret.SpecProductFrozenCommission = goFloat64(p.SpecProductFrozenCommission)
	ret.SpecProductPositionProfit = goFloat64(p.SpecProductPositionProfit)
	ret.SpecProductCloseProfit = goFloat64(p.SpecProductCloseProfit)
	ret.SpecProductPositionProfitByAlg = goFloat64(p.SpecProductPositionProfitByAlg)
	ret.SpecProductExchangeMargin = goFloat64(p.SpecProductExchangeMargin)
	ret.FrozenSwap = goFloat64(p.FrozenSwap)
	ret.RemainSwap = goFloat64(p.RemainSwap)
	return ret
}
func (s *CThostFtdcSyncingTradingAccountField) CValue() *C.CThostFtdcSyncingTradingAccountField {
	ptr := C.newCThostFtdcSyncingTradingAccountField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ptr.PreMortgage = C.double(s.PreMortgage)
	ptr.PreCredit = C.double(s.PreCredit)
	ptr.PreDeposit = C.double(s.PreDeposit)
	ptr.PreBalance = C.double(s.PreBalance)
	ptr.PreMargin = C.double(s.PreMargin)
	ptr.InterestBase = C.double(s.InterestBase)
	ptr.Interest = C.double(s.Interest)
	ptr.Deposit = C.double(s.Deposit)
	ptr.Withdraw = C.double(s.Withdraw)
	ptr.FrozenMargin = C.double(s.FrozenMargin)
	ptr.FrozenCash = C.double(s.FrozenCash)
	ptr.FrozenCommission = C.double(s.FrozenCommission)
	ptr.CurrMargin = C.double(s.CurrMargin)
	ptr.CashIn = C.double(s.CashIn)
	ptr.Commission = C.double(s.Commission)
	ptr.CloseProfit = C.double(s.CloseProfit)
	ptr.PositionProfit = C.double(s.PositionProfit)
	ptr.Balance = C.double(s.Balance)
	ptr.Available = C.double(s.Available)
	ptr.WithdrawQuota = C.double(s.WithdrawQuota)
	ptr.Reserve = C.double(s.Reserve)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	ptr.Credit = C.double(s.Credit)
	ptr.Mortgage = C.double(s.Mortgage)
	ptr.ExchangeMargin = C.double(s.ExchangeMargin)
	ptr.DeliveryMargin = C.double(s.DeliveryMargin)
	ptr.ExchangeDeliveryMargin = C.double(s.ExchangeDeliveryMargin)
	ptr.ReserveBalance = C.double(s.ReserveBalance)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.PreFundMortgageIn = C.double(s.PreFundMortgageIn)
	ptr.PreFundMortgageOut = C.double(s.PreFundMortgageOut)
	ptr.FundMortgageIn = C.double(s.FundMortgageIn)
	ptr.FundMortgageOut = C.double(s.FundMortgageOut)
	ptr.FundMortgageAvailable = C.double(s.FundMortgageAvailable)
	ptr.MortgageableFund = C.double(s.MortgageableFund)
	ptr.SpecProductMargin = C.double(s.SpecProductMargin)
	ptr.SpecProductFrozenMargin = C.double(s.SpecProductFrozenMargin)
	ptr.SpecProductCommission = C.double(s.SpecProductCommission)
	ptr.SpecProductFrozenCommission = C.double(s.SpecProductFrozenCommission)
	ptr.SpecProductPositionProfit = C.double(s.SpecProductPositionProfit)
	ptr.SpecProductCloseProfit = C.double(s.SpecProductCloseProfit)
	ptr.SpecProductPositionProfitByAlg = C.double(s.SpecProductPositionProfitByAlg)
	ptr.SpecProductExchangeMargin = C.double(s.SpecProductExchangeMargin)
	ptr.FrozenSwap = C.double(s.FrozenSwap)
	ptr.RemainSwap = C.double(s.RemainSwap)
	return ptr
}

type CThostFtdcSyncingInvestorPositionField struct {
	Reserve1           string
	BrokerID           string
	InvestorID         string
	PosiDirection      byte
	HedgeFlag          byte
	PositionDate       byte
	YdPosition         int
	Position           int
	LongFrozen         int
	ShortFrozen        int
	LongFrozenAmount   float64
	ShortFrozenAmount  float64
	OpenVolume         int
	CloseVolume        int
	OpenAmount         float64
	CloseAmount        float64
	PositionCost       float64
	PreMargin          float64
	UseMargin          float64
	FrozenMargin       float64
	FrozenCash         float64
	FrozenCommission   float64
	CashIn             float64
	Commission         float64
	CloseProfit        float64
	PositionProfit     float64
	PreSettlementPrice float64
	SettlementPrice    float64
	TradingDay         string
	SettlementID       int
	OpenCost           float64
	ExchangeMargin     float64
	CombPosition       int
	CombLongFrozen     int
	CombShortFrozen    int
	CloseProfitByDate  float64
	CloseProfitByTrade float64
	TodayPosition      int
	MarginRateByMoney  float64
	MarginRateByVolume float64
	StrikeFrozen       int
	StrikeFrozenAmount float64
	AbandonFrozen      int
	ExchangeID         string
	YdStrikeFrozen     int
	InvestUnitID       string
	PositionCostOffset float64
	TasPosition        int
	TasPositionCost    float64
	InstrumentID       string
}

func NewCThostFtdcSyncingInvestorPositionField(p *C.CThostFtdcSyncingInvestorPositionField) *CThostFtdcSyncingInvestorPositionField {
	ret := new(CThostFtdcSyncingInvestorPositionField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.PosiDirection = byte(p.PosiDirection)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.PositionDate = byte(p.PositionDate)
	ret.YdPosition = int(p.YdPosition)
	ret.Position = int(p.Position)
	ret.LongFrozen = int(p.LongFrozen)
	ret.ShortFrozen = int(p.ShortFrozen)
	ret.LongFrozenAmount = goFloat64(p.LongFrozenAmount)
	ret.ShortFrozenAmount = goFloat64(p.ShortFrozenAmount)
	ret.OpenVolume = int(p.OpenVolume)
	ret.CloseVolume = int(p.CloseVolume)
	ret.OpenAmount = goFloat64(p.OpenAmount)
	ret.CloseAmount = goFloat64(p.CloseAmount)
	ret.PositionCost = goFloat64(p.PositionCost)
	ret.PreMargin = goFloat64(p.PreMargin)
	ret.UseMargin = goFloat64(p.UseMargin)
	ret.FrozenMargin = goFloat64(p.FrozenMargin)
	ret.FrozenCash = goFloat64(p.FrozenCash)
	ret.FrozenCommission = goFloat64(p.FrozenCommission)
	ret.CashIn = goFloat64(p.CashIn)
	ret.Commission = goFloat64(p.Commission)
	ret.CloseProfit = goFloat64(p.CloseProfit)
	ret.PositionProfit = goFloat64(p.PositionProfit)
	ret.PreSettlementPrice = goFloat64(p.PreSettlementPrice)
	ret.SettlementPrice = goFloat64(p.SettlementPrice)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.OpenCost = goFloat64(p.OpenCost)
	ret.ExchangeMargin = goFloat64(p.ExchangeMargin)
	ret.CombPosition = int(p.CombPosition)
	ret.CombLongFrozen = int(p.CombLongFrozen)
	ret.CombShortFrozen = int(p.CombShortFrozen)
	ret.CloseProfitByDate = goFloat64(p.CloseProfitByDate)
	ret.CloseProfitByTrade = goFloat64(p.CloseProfitByTrade)
	ret.TodayPosition = int(p.TodayPosition)
	ret.MarginRateByMoney = goFloat64(p.MarginRateByMoney)
	ret.MarginRateByVolume = goFloat64(p.MarginRateByVolume)
	ret.StrikeFrozen = int(p.StrikeFrozen)
	ret.StrikeFrozenAmount = goFloat64(p.StrikeFrozenAmount)
	ret.AbandonFrozen = int(p.AbandonFrozen)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.YdStrikeFrozen = int(p.YdStrikeFrozen)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.PositionCostOffset = goFloat64(p.PositionCostOffset)
	ret.TasPosition = int(p.TasPosition)
	ret.TasPositionCost = goFloat64(p.TasPositionCost)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcSyncingInvestorPositionField) CValue() *C.CThostFtdcSyncingInvestorPositionField {
	ptr := C.newCThostFtdcSyncingInvestorPositionField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.PosiDirection = C.char(s.PosiDirection)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.PositionDate = C.char(s.PositionDate)
	ptr.YdPosition = C.int(s.YdPosition)
	ptr.Position = C.int(s.Position)
	ptr.LongFrozen = C.int(s.LongFrozen)
	ptr.ShortFrozen = C.int(s.ShortFrozen)
	ptr.LongFrozenAmount = C.double(s.LongFrozenAmount)
	ptr.ShortFrozenAmount = C.double(s.ShortFrozenAmount)
	ptr.OpenVolume = C.int(s.OpenVolume)
	ptr.CloseVolume = C.int(s.CloseVolume)
	ptr.OpenAmount = C.double(s.OpenAmount)
	ptr.CloseAmount = C.double(s.CloseAmount)
	ptr.PositionCost = C.double(s.PositionCost)
	ptr.PreMargin = C.double(s.PreMargin)
	ptr.UseMargin = C.double(s.UseMargin)
	ptr.FrozenMargin = C.double(s.FrozenMargin)
	ptr.FrozenCash = C.double(s.FrozenCash)
	ptr.FrozenCommission = C.double(s.FrozenCommission)
	ptr.CashIn = C.double(s.CashIn)
	ptr.Commission = C.double(s.Commission)
	ptr.CloseProfit = C.double(s.CloseProfit)
	ptr.PositionProfit = C.double(s.PositionProfit)
	ptr.PreSettlementPrice = C.double(s.PreSettlementPrice)
	ptr.SettlementPrice = C.double(s.SettlementPrice)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	ptr.OpenCost = C.double(s.OpenCost)
	ptr.ExchangeMargin = C.double(s.ExchangeMargin)
	ptr.CombPosition = C.int(s.CombPosition)
	ptr.CombLongFrozen = C.int(s.CombLongFrozen)
	ptr.CombShortFrozen = C.int(s.CombShortFrozen)
	ptr.CloseProfitByDate = C.double(s.CloseProfitByDate)
	ptr.CloseProfitByTrade = C.double(s.CloseProfitByTrade)
	ptr.TodayPosition = C.int(s.TodayPosition)
	ptr.MarginRateByMoney = C.double(s.MarginRateByMoney)
	ptr.MarginRateByVolume = C.double(s.MarginRateByVolume)
	ptr.StrikeFrozen = C.int(s.StrikeFrozen)
	ptr.StrikeFrozenAmount = C.double(s.StrikeFrozenAmount)
	ptr.AbandonFrozen = C.int(s.AbandonFrozen)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ptr.YdStrikeFrozen = C.int(s.YdStrikeFrozen)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ptr.PositionCostOffset = C.double(s.PositionCostOffset)
	ptr.TasPosition = C.int(s.TasPosition)
	ptr.TasPositionCost = C.double(s.TasPositionCost)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcSyncingInstrumentMarginRateField struct {
	Reserve1                 string
	InvestorRange            byte
	BrokerID                 string
	InvestorID               string
	HedgeFlag                byte
	LongMarginRatioByMoney   float64
	LongMarginRatioByVolume  float64
	ShortMarginRatioByMoney  float64
	ShortMarginRatioByVolume float64
	IsRelative               int
	InstrumentID             string
}

func NewCThostFtdcSyncingInstrumentMarginRateField(p *C.CThostFtdcSyncingInstrumentMarginRateField) *CThostFtdcSyncingInstrumentMarginRateField {
	ret := new(CThostFtdcSyncingInstrumentMarginRateField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.LongMarginRatioByMoney = goFloat64(p.LongMarginRatioByMoney)
	ret.LongMarginRatioByVolume = goFloat64(p.LongMarginRatioByVolume)
	ret.ShortMarginRatioByMoney = goFloat64(p.ShortMarginRatioByMoney)
	ret.ShortMarginRatioByVolume = goFloat64(p.ShortMarginRatioByVolume)
	ret.IsRelative = int(p.IsRelative)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcSyncingInstrumentMarginRateField) CValue() *C.CThostFtdcSyncingInstrumentMarginRateField {
	ptr := C.newCThostFtdcSyncingInstrumentMarginRateField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.LongMarginRatioByMoney = C.double(s.LongMarginRatioByMoney)
	ptr.LongMarginRatioByVolume = C.double(s.LongMarginRatioByVolume)
	ptr.ShortMarginRatioByMoney = C.double(s.ShortMarginRatioByMoney)
	ptr.ShortMarginRatioByVolume = C.double(s.ShortMarginRatioByVolume)
	ptr.IsRelative = C.int(s.IsRelative)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcSyncingInstrumentCommissionRateField struct {
	Reserve1                string
	InvestorRange           byte
	BrokerID                string
	InvestorID              string
	OpenRatioByMoney        float64
	OpenRatioByVolume       float64
	CloseRatioByMoney       float64
	CloseRatioByVolume      float64
	CloseTodayRatioByMoney  float64
	CloseTodayRatioByVolume float64
	InstrumentID            string
}

func NewCThostFtdcSyncingInstrumentCommissionRateField(p *C.CThostFtdcSyncingInstrumentCommissionRateField) *CThostFtdcSyncingInstrumentCommissionRateField {
	ret := new(CThostFtdcSyncingInstrumentCommissionRateField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.OpenRatioByMoney = goFloat64(p.OpenRatioByMoney)
	ret.OpenRatioByVolume = goFloat64(p.OpenRatioByVolume)
	ret.CloseRatioByMoney = goFloat64(p.CloseRatioByMoney)
	ret.CloseRatioByVolume = goFloat64(p.CloseRatioByVolume)
	ret.CloseTodayRatioByMoney = goFloat64(p.CloseTodayRatioByMoney)
	ret.CloseTodayRatioByVolume = goFloat64(p.CloseTodayRatioByVolume)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcSyncingInstrumentCommissionRateField) CValue() *C.CThostFtdcSyncingInstrumentCommissionRateField {
	ptr := C.newCThostFtdcSyncingInstrumentCommissionRateField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.OpenRatioByMoney = C.double(s.OpenRatioByMoney)
	ptr.OpenRatioByVolume = C.double(s.OpenRatioByVolume)
	ptr.CloseRatioByMoney = C.double(s.CloseRatioByMoney)
	ptr.CloseRatioByVolume = C.double(s.CloseRatioByVolume)
	ptr.CloseTodayRatioByMoney = C.double(s.CloseTodayRatioByMoney)
	ptr.CloseTodayRatioByVolume = C.double(s.CloseTodayRatioByVolume)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcSyncingInstrumentTradingRightField struct {
	Reserve1      string
	InvestorRange byte
	BrokerID      string
	InvestorID    string
	TradingRight  byte
	InstrumentID  string
}

func NewCThostFtdcSyncingInstrumentTradingRightField(p *C.CThostFtdcSyncingInstrumentTradingRightField) *CThostFtdcSyncingInstrumentTradingRightField {
	ret := new(CThostFtdcSyncingInstrumentTradingRightField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.TradingRight = byte(p.TradingRight)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcSyncingInstrumentTradingRightField) CValue() *C.CThostFtdcSyncingInstrumentTradingRightField {
	ptr := C.newCThostFtdcSyncingInstrumentTradingRightField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.TradingRight = C.char(s.TradingRight)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryOrderField struct {
	BrokerID        string
	InvestorID      string
	Reserve1        string
	ExchangeID      string
	OrderSysID      string
	InsertTimeStart string
	InsertTimeEnd   string
	InvestUnitID    string
	InstrumentID    string
}

func NewCThostFtdcQryOrderField(p *C.CThostFtdcQryOrderField) *CThostFtdcQryOrderField {
	ret := new(CThostFtdcQryOrderField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.OrderSysID = c2goStr(&p.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.InsertTimeStart = c2goStr(&p.InsertTimeStart[0], C.sizeof_TThostFtdcTimeType)
	ret.InsertTimeEnd = c2goStr(&p.InsertTimeEnd[0], C.sizeof_TThostFtdcTimeType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryOrderField) CValue() *C.CThostFtdcQryOrderField {
	ptr := C.newCThostFtdcQryOrderField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.OrderSysID, &ptr.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.InsertTimeStart, &ptr.InsertTimeStart[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.InsertTimeEnd, &ptr.InsertTimeEnd[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryTradeField struct {
	BrokerID       string
	InvestorID     string
	Reserve1       string
	ExchangeID     string
	TradeID        string
	TradeTimeStart string
	TradeTimeEnd   string
	InvestUnitID   string
	InstrumentID   string
}

func NewCThostFtdcQryTradeField(p *C.CThostFtdcQryTradeField) *CThostFtdcQryTradeField {
	ret := new(CThostFtdcQryTradeField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.TradeID = c2goStr(&p.TradeID[0], C.sizeof_TThostFtdcTradeIDType)
	ret.TradeTimeStart = c2goStr(&p.TradeTimeStart[0], C.sizeof_TThostFtdcTimeType)
	ret.TradeTimeEnd = c2goStr(&p.TradeTimeEnd[0], C.sizeof_TThostFtdcTimeType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryTradeField) CValue() *C.CThostFtdcQryTradeField {
	ptr := C.newCThostFtdcQryTradeField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.TradeID, &ptr.TradeID[0], C.sizeof_TThostFtdcTradeIDType)
	go2cStr(s.TradeTimeStart, &ptr.TradeTimeStart[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.TradeTimeEnd, &ptr.TradeTimeEnd[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryInvestorPositionField struct {
	BrokerID     string
	InvestorID   string
	Reserve1     string
	ExchangeID   string
	InvestUnitID string
	InstrumentID string
}

func NewCThostFtdcQryInvestorPositionField(p *C.CThostFtdcQryInvestorPositionField) *CThostFtdcQryInvestorPositionField {
	ret := new(CThostFtdcQryInvestorPositionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryInvestorPositionField) CValue() *C.CThostFtdcQryInvestorPositionField {
	ptr := C.newCThostFtdcQryInvestorPositionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryTradingAccountField struct {
	BrokerID   string
	InvestorID string
	CurrencyID string
	BizType    byte
	AccountID  string
}

func NewCThostFtdcQryTradingAccountField(p *C.CThostFtdcQryTradingAccountField) *CThostFtdcQryTradingAccountField {
	ret := new(CThostFtdcQryTradingAccountField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.BizType = byte(p.BizType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	return ret
}
func (s *CThostFtdcQryTradingAccountField) CValue() *C.CThostFtdcQryTradingAccountField {
	ptr := C.newCThostFtdcQryTradingAccountField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.BizType = C.char(s.BizType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	return ptr
}

type CThostFtdcQryInvestorField struct {
	BrokerID   string
	InvestorID string
}

func NewCThostFtdcQryInvestorField(p *C.CThostFtdcQryInvestorField) *CThostFtdcQryInvestorField {
	ret := new(CThostFtdcQryInvestorField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ret
}
func (s *CThostFtdcQryInvestorField) CValue() *C.CThostFtdcQryInvestorField {
	ptr := C.newCThostFtdcQryInvestorField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ptr
}

type CThostFtdcQryTradingCodeField struct {
	BrokerID     string
	InvestorID   string
	ExchangeID   string
	ClientID     string
	ClientIDType byte
	InvestUnitID string
}

func NewCThostFtdcQryTradingCodeField(p *C.CThostFtdcQryTradingCodeField) *CThostFtdcQryTradingCodeField {
	ret := new(CThostFtdcQryTradingCodeField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.ClientIDType = byte(p.ClientIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ret
}
func (s *CThostFtdcQryTradingCodeField) CValue() *C.CThostFtdcQryTradingCodeField {
	ptr := C.newCThostFtdcQryTradingCodeField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ptr.ClientIDType = C.char(s.ClientIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ptr
}

type CThostFtdcQryInvestorGroupField struct {
	BrokerID string
}

func NewCThostFtdcQryInvestorGroupField(p *C.CThostFtdcQryInvestorGroupField) *CThostFtdcQryInvestorGroupField {
	ret := new(CThostFtdcQryInvestorGroupField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	return ret
}
func (s *CThostFtdcQryInvestorGroupField) CValue() *C.CThostFtdcQryInvestorGroupField {
	ptr := C.newCThostFtdcQryInvestorGroupField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	return ptr
}

type CThostFtdcQryInstrumentMarginRateField struct {
	BrokerID     string
	InvestorID   string
	Reserve1     string
	HedgeFlag    byte
	ExchangeID   string
	InvestUnitID string
	InstrumentID string
}

func NewCThostFtdcQryInstrumentMarginRateField(p *C.CThostFtdcQryInstrumentMarginRateField) *CThostFtdcQryInstrumentMarginRateField {
	ret := new(CThostFtdcQryInstrumentMarginRateField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryInstrumentMarginRateField) CValue() *C.CThostFtdcQryInstrumentMarginRateField {
	ptr := C.newCThostFtdcQryInstrumentMarginRateField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryInstrumentCommissionRateField struct {
	BrokerID     string
	InvestorID   string
	Reserve1     string
	ExchangeID   string
	InvestUnitID string
	InstrumentID string
}

func NewCThostFtdcQryInstrumentCommissionRateField(p *C.CThostFtdcQryInstrumentCommissionRateField) *CThostFtdcQryInstrumentCommissionRateField {
	ret := new(CThostFtdcQryInstrumentCommissionRateField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryInstrumentCommissionRateField) CValue() *C.CThostFtdcQryInstrumentCommissionRateField {
	ptr := C.newCThostFtdcQryInstrumentCommissionRateField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryInstrumentTradingRightField struct {
	BrokerID     string
	InvestorID   string
	Reserve1     string
	InstrumentID string
}

func NewCThostFtdcQryInstrumentTradingRightField(p *C.CThostFtdcQryInstrumentTradingRightField) *CThostFtdcQryInstrumentTradingRightField {
	ret := new(CThostFtdcQryInstrumentTradingRightField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryInstrumentTradingRightField) CValue() *C.CThostFtdcQryInstrumentTradingRightField {
	ptr := C.newCThostFtdcQryInstrumentTradingRightField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryBrokerField struct {
	BrokerID string
}

func NewCThostFtdcQryBrokerField(p *C.CThostFtdcQryBrokerField) *CThostFtdcQryBrokerField {
	ret := new(CThostFtdcQryBrokerField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	return ret
}
func (s *CThostFtdcQryBrokerField) CValue() *C.CThostFtdcQryBrokerField {
	ptr := C.newCThostFtdcQryBrokerField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	return ptr
}

type CThostFtdcQryTraderField struct {
	ExchangeID    string
	ParticipantID string
	TraderID      string
}

func NewCThostFtdcQryTraderField(p *C.CThostFtdcQryTraderField) *CThostFtdcQryTraderField {
	ret := new(CThostFtdcQryTraderField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	return ret
}
func (s *CThostFtdcQryTraderField) CValue() *C.CThostFtdcQryTraderField {
	ptr := C.newCThostFtdcQryTraderField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	return ptr
}

type CThostFtdcQrySuperUserFunctionField struct {
	UserID string
}

func NewCThostFtdcQrySuperUserFunctionField(p *C.CThostFtdcQrySuperUserFunctionField) *CThostFtdcQrySuperUserFunctionField {
	ret := new(CThostFtdcQrySuperUserFunctionField)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ret
}
func (s *CThostFtdcQrySuperUserFunctionField) CValue() *C.CThostFtdcQrySuperUserFunctionField {
	ptr := C.newCThostFtdcQrySuperUserFunctionField()
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ptr
}

type CThostFtdcQryUserSessionField struct {
	FrontID   int
	SessionID int
	BrokerID  string
	UserID    string
}

func NewCThostFtdcQryUserSessionField(p *C.CThostFtdcQryUserSessionField) *CThostFtdcQryUserSessionField {
	ret := new(CThostFtdcQryUserSessionField)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ret
}
func (s *CThostFtdcQryUserSessionField) CValue() *C.CThostFtdcQryUserSessionField {
	ptr := C.newCThostFtdcQryUserSessionField()
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ptr
}

type CThostFtdcQryPartBrokerField struct {
	ExchangeID    string
	BrokerID      string
	ParticipantID string
}

func NewCThostFtdcQryPartBrokerField(p *C.CThostFtdcQryPartBrokerField) *CThostFtdcQryPartBrokerField {
	ret := new(CThostFtdcQryPartBrokerField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	return ret
}
func (s *CThostFtdcQryPartBrokerField) CValue() *C.CThostFtdcQryPartBrokerField {
	ptr := C.newCThostFtdcQryPartBrokerField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	return ptr
}

type CThostFtdcQryFrontStatusField struct {
	FrontID int
}

func NewCThostFtdcQryFrontStatusField(p *C.CThostFtdcQryFrontStatusField) *CThostFtdcQryFrontStatusField {
	ret := new(CThostFtdcQryFrontStatusField)
	ret.FrontID = int(p.FrontID)
	return ret
}
func (s *CThostFtdcQryFrontStatusField) CValue() *C.CThostFtdcQryFrontStatusField {
	ptr := C.newCThostFtdcQryFrontStatusField()
	ptr.FrontID = C.int(s.FrontID)
	return ptr
}

type CThostFtdcQryExchangeOrderField struct {
	ParticipantID  string
	ClientID       string
	Reserve1       string
	ExchangeID     string
	TraderID       string
	ExchangeInstID string
}

func NewCThostFtdcQryExchangeOrderField(p *C.CThostFtdcQryExchangeOrderField) *CThostFtdcQryExchangeOrderField {
	ret := new(CThostFtdcQryExchangeOrderField)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ret
}
func (s *CThostFtdcQryExchangeOrderField) CValue() *C.CThostFtdcQryExchangeOrderField {
	ptr := C.newCThostFtdcQryExchangeOrderField()
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ptr
}

type CThostFtdcQryOrderActionField struct {
	BrokerID   string
	InvestorID string
	ExchangeID string
}

func NewCThostFtdcQryOrderActionField(p *C.CThostFtdcQryOrderActionField) *CThostFtdcQryOrderActionField {
	ret := new(CThostFtdcQryOrderActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ret
}
func (s *CThostFtdcQryOrderActionField) CValue() *C.CThostFtdcQryOrderActionField {
	ptr := C.newCThostFtdcQryOrderActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ptr
}

type CThostFtdcQryExchangeOrderActionField struct {
	ParticipantID string
	ClientID      string
	ExchangeID    string
	TraderID      string
}

func NewCThostFtdcQryExchangeOrderActionField(p *C.CThostFtdcQryExchangeOrderActionField) *CThostFtdcQryExchangeOrderActionField {
	ret := new(CThostFtdcQryExchangeOrderActionField)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	return ret
}
func (s *CThostFtdcQryExchangeOrderActionField) CValue() *C.CThostFtdcQryExchangeOrderActionField {
	ptr := C.newCThostFtdcQryExchangeOrderActionField()
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	return ptr
}

type CThostFtdcQrySuperUserField struct {
	UserID string
}

func NewCThostFtdcQrySuperUserField(p *C.CThostFtdcQrySuperUserField) *CThostFtdcQrySuperUserField {
	ret := new(CThostFtdcQrySuperUserField)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ret
}
func (s *CThostFtdcQrySuperUserField) CValue() *C.CThostFtdcQrySuperUserField {
	ptr := C.newCThostFtdcQrySuperUserField()
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ptr
}

type CThostFtdcQryExchangeField struct {
	ExchangeID string
}

func NewCThostFtdcQryExchangeField(p *C.CThostFtdcQryExchangeField) *CThostFtdcQryExchangeField {
	ret := new(CThostFtdcQryExchangeField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ret
}
func (s *CThostFtdcQryExchangeField) CValue() *C.CThostFtdcQryExchangeField {
	ptr := C.newCThostFtdcQryExchangeField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ptr
}

type CThostFtdcQryProductField struct {
	Reserve1     string
	ProductClass byte
	ExchangeID   string
	ProductID    string
}

func NewCThostFtdcQryProductField(p *C.CThostFtdcQryProductField) *CThostFtdcQryProductField {
	ret := new(CThostFtdcQryProductField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ProductClass = byte(p.ProductClass)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ProductID = c2goStr(&p.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryProductField) CValue() *C.CThostFtdcQryProductField {
	ptr := C.newCThostFtdcQryProductField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.ProductClass = C.char(s.ProductClass)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ProductID, &ptr.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryInstrumentField struct {
	Reserve1       string
	ExchangeID     string
	Reserve2       string
	Reserve3       string
	InstrumentID   string
	ExchangeInstID string
	ProductID      string
}

func NewCThostFtdcQryInstrumentField(p *C.CThostFtdcQryInstrumentField) *CThostFtdcQryInstrumentField {
	ret := new(CThostFtdcQryInstrumentField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.Reserve3 = c2goStr(&p.reserve3[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.ProductID = c2goStr(&p.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryInstrumentField) CValue() *C.CThostFtdcQryInstrumentField {
	ptr := C.newCThostFtdcQryInstrumentField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.Reserve3, &ptr.reserve3[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	go2cStr(s.ProductID, &ptr.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryDepthMarketDataField struct {
	Reserve1     string
	ExchangeID   string
	InstrumentID string
}

func NewCThostFtdcQryDepthMarketDataField(p *C.CThostFtdcQryDepthMarketDataField) *CThostFtdcQryDepthMarketDataField {
	ret := new(CThostFtdcQryDepthMarketDataField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryDepthMarketDataField) CValue() *C.CThostFtdcQryDepthMarketDataField {
	ptr := C.newCThostFtdcQryDepthMarketDataField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryBrokerUserField struct {
	BrokerID string
	UserID   string
}

func NewCThostFtdcQryBrokerUserField(p *C.CThostFtdcQryBrokerUserField) *CThostFtdcQryBrokerUserField {
	ret := new(CThostFtdcQryBrokerUserField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ret
}
func (s *CThostFtdcQryBrokerUserField) CValue() *C.CThostFtdcQryBrokerUserField {
	ptr := C.newCThostFtdcQryBrokerUserField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ptr
}

type CThostFtdcQryBrokerUserFunctionField struct {
	BrokerID string
	UserID   string
}

func NewCThostFtdcQryBrokerUserFunctionField(p *C.CThostFtdcQryBrokerUserFunctionField) *CThostFtdcQryBrokerUserFunctionField {
	ret := new(CThostFtdcQryBrokerUserFunctionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ret
}
func (s *CThostFtdcQryBrokerUserFunctionField) CValue() *C.CThostFtdcQryBrokerUserFunctionField {
	ptr := C.newCThostFtdcQryBrokerUserFunctionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ptr
}

type CThostFtdcQryTraderOfferField struct {
	ExchangeID    string
	ParticipantID string
	TraderID      string
}

func NewCThostFtdcQryTraderOfferField(p *C.CThostFtdcQryTraderOfferField) *CThostFtdcQryTraderOfferField {
	ret := new(CThostFtdcQryTraderOfferField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	return ret
}
func (s *CThostFtdcQryTraderOfferField) CValue() *C.CThostFtdcQryTraderOfferField {
	ptr := C.newCThostFtdcQryTraderOfferField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	return ptr
}

type CThostFtdcQrySyncDepositField struct {
	BrokerID     string
	DepositSeqNo string
}

func NewCThostFtdcQrySyncDepositField(p *C.CThostFtdcQrySyncDepositField) *CThostFtdcQrySyncDepositField {
	ret := new(CThostFtdcQrySyncDepositField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.DepositSeqNo = c2goStr(&p.DepositSeqNo[0], C.sizeof_TThostFtdcDepositSeqNoType)
	return ret
}
func (s *CThostFtdcQrySyncDepositField) CValue() *C.CThostFtdcQrySyncDepositField {
	ptr := C.newCThostFtdcQrySyncDepositField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.DepositSeqNo, &ptr.DepositSeqNo[0], C.sizeof_TThostFtdcDepositSeqNoType)
	return ptr
}

type CThostFtdcQrySettlementInfoField struct {
	BrokerID   string
	InvestorID string
	TradingDay string
	AccountID  string
	CurrencyID string
}

func NewCThostFtdcQrySettlementInfoField(p *C.CThostFtdcQrySettlementInfoField) *CThostFtdcQrySettlementInfoField {
	ret := new(CThostFtdcQrySettlementInfoField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ret
}
func (s *CThostFtdcQrySettlementInfoField) CValue() *C.CThostFtdcQrySettlementInfoField {
	ptr := C.newCThostFtdcQrySettlementInfoField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ptr
}

type CThostFtdcQryExchangeMarginRateField struct {
	BrokerID     string
	Reserve1     string
	HedgeFlag    byte
	ExchangeID   string
	InstrumentID string
}

func NewCThostFtdcQryExchangeMarginRateField(p *C.CThostFtdcQryExchangeMarginRateField) *CThostFtdcQryExchangeMarginRateField {
	ret := new(CThostFtdcQryExchangeMarginRateField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryExchangeMarginRateField) CValue() *C.CThostFtdcQryExchangeMarginRateField {
	ptr := C.newCThostFtdcQryExchangeMarginRateField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryExchangeMarginRateAdjustField struct {
	BrokerID     string
	Reserve1     string
	HedgeFlag    byte
	InstrumentID string
}

func NewCThostFtdcQryExchangeMarginRateAdjustField(p *C.CThostFtdcQryExchangeMarginRateAdjustField) *CThostFtdcQryExchangeMarginRateAdjustField {
	ret := new(CThostFtdcQryExchangeMarginRateAdjustField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryExchangeMarginRateAdjustField) CValue() *C.CThostFtdcQryExchangeMarginRateAdjustField {
	ptr := C.newCThostFtdcQryExchangeMarginRateAdjustField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryExchangeRateField struct {
	BrokerID       string
	FromCurrencyID string
	ToCurrencyID   string
}

func NewCThostFtdcQryExchangeRateField(p *C.CThostFtdcQryExchangeRateField) *CThostFtdcQryExchangeRateField {
	ret := new(CThostFtdcQryExchangeRateField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.FromCurrencyID = c2goStr(&p.FromCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.ToCurrencyID = c2goStr(&p.ToCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ret
}
func (s *CThostFtdcQryExchangeRateField) CValue() *C.CThostFtdcQryExchangeRateField {
	ptr := C.newCThostFtdcQryExchangeRateField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.FromCurrencyID, &ptr.FromCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.ToCurrencyID, &ptr.ToCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ptr
}

type CThostFtdcQrySyncFundMortgageField struct {
	BrokerID      string
	MortgageSeqNo string
}

func NewCThostFtdcQrySyncFundMortgageField(p *C.CThostFtdcQrySyncFundMortgageField) *CThostFtdcQrySyncFundMortgageField {
	ret := new(CThostFtdcQrySyncFundMortgageField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.MortgageSeqNo = c2goStr(&p.MortgageSeqNo[0], C.sizeof_TThostFtdcDepositSeqNoType)
	return ret
}
func (s *CThostFtdcQrySyncFundMortgageField) CValue() *C.CThostFtdcQrySyncFundMortgageField {
	ptr := C.newCThostFtdcQrySyncFundMortgageField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.MortgageSeqNo, &ptr.MortgageSeqNo[0], C.sizeof_TThostFtdcDepositSeqNoType)
	return ptr
}

type CThostFtdcQryHisOrderField struct {
	BrokerID        string
	InvestorID      string
	Reserve1        string
	ExchangeID      string
	OrderSysID      string
	InsertTimeStart string
	InsertTimeEnd   string
	TradingDay      string
	SettlementID    int
	InstrumentID    string
}

func NewCThostFtdcQryHisOrderField(p *C.CThostFtdcQryHisOrderField) *CThostFtdcQryHisOrderField {
	ret := new(CThostFtdcQryHisOrderField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.OrderSysID = c2goStr(&p.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.InsertTimeStart = c2goStr(&p.InsertTimeStart[0], C.sizeof_TThostFtdcTimeType)
	ret.InsertTimeEnd = c2goStr(&p.InsertTimeEnd[0], C.sizeof_TThostFtdcTimeType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryHisOrderField) CValue() *C.CThostFtdcQryHisOrderField {
	ptr := C.newCThostFtdcQryHisOrderField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.OrderSysID, &ptr.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.InsertTimeStart, &ptr.InsertTimeStart[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.InsertTimeEnd, &ptr.InsertTimeEnd[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcOptionInstrMiniMarginField struct {
	Reserve1      string
	InvestorRange byte
	BrokerID      string
	InvestorID    string
	MinMargin     float64
	ValueMethod   byte
	IsRelative    int
	InstrumentID  string
}

func NewCThostFtdcOptionInstrMiniMarginField(p *C.CThostFtdcOptionInstrMiniMarginField) *CThostFtdcOptionInstrMiniMarginField {
	ret := new(CThostFtdcOptionInstrMiniMarginField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.MinMargin = goFloat64(p.MinMargin)
	ret.ValueMethod = byte(p.ValueMethod)
	ret.IsRelative = int(p.IsRelative)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcOptionInstrMiniMarginField) CValue() *C.CThostFtdcOptionInstrMiniMarginField {
	ptr := C.newCThostFtdcOptionInstrMiniMarginField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.MinMargin = C.double(s.MinMargin)
	ptr.ValueMethod = C.char(s.ValueMethod)
	ptr.IsRelative = C.int(s.IsRelative)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcOptionInstrMarginAdjustField struct {
	Reserve1                  string
	InvestorRange             byte
	BrokerID                  string
	InvestorID                string
	SShortMarginRatioByMoney  float64
	SShortMarginRatioByVolume float64
	HShortMarginRatioByMoney  float64
	HShortMarginRatioByVolume float64
	AShortMarginRatioByMoney  float64
	AShortMarginRatioByVolume float64
	IsRelative                int
	MShortMarginRatioByMoney  float64
	MShortMarginRatioByVolume float64
	InstrumentID              string
}

func NewCThostFtdcOptionInstrMarginAdjustField(p *C.CThostFtdcOptionInstrMarginAdjustField) *CThostFtdcOptionInstrMarginAdjustField {
	ret := new(CThostFtdcOptionInstrMarginAdjustField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.SShortMarginRatioByMoney = goFloat64(p.SShortMarginRatioByMoney)
	ret.SShortMarginRatioByVolume = goFloat64(p.SShortMarginRatioByVolume)
	ret.HShortMarginRatioByMoney = goFloat64(p.HShortMarginRatioByMoney)
	ret.HShortMarginRatioByVolume = goFloat64(p.HShortMarginRatioByVolume)
	ret.AShortMarginRatioByMoney = goFloat64(p.AShortMarginRatioByMoney)
	ret.AShortMarginRatioByVolume = goFloat64(p.AShortMarginRatioByVolume)
	ret.IsRelative = int(p.IsRelative)
	ret.MShortMarginRatioByMoney = goFloat64(p.MShortMarginRatioByMoney)
	ret.MShortMarginRatioByVolume = goFloat64(p.MShortMarginRatioByVolume)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcOptionInstrMarginAdjustField) CValue() *C.CThostFtdcOptionInstrMarginAdjustField {
	ptr := C.newCThostFtdcOptionInstrMarginAdjustField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.SShortMarginRatioByMoney = C.double(s.SShortMarginRatioByMoney)
	ptr.SShortMarginRatioByVolume = C.double(s.SShortMarginRatioByVolume)
	ptr.HShortMarginRatioByMoney = C.double(s.HShortMarginRatioByMoney)
	ptr.HShortMarginRatioByVolume = C.double(s.HShortMarginRatioByVolume)
	ptr.AShortMarginRatioByMoney = C.double(s.AShortMarginRatioByMoney)
	ptr.AShortMarginRatioByVolume = C.double(s.AShortMarginRatioByVolume)
	ptr.IsRelative = C.int(s.IsRelative)
	ptr.MShortMarginRatioByMoney = C.double(s.MShortMarginRatioByMoney)
	ptr.MShortMarginRatioByVolume = C.double(s.MShortMarginRatioByVolume)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcOptionInstrCommRateField struct {
	Reserve1                string
	InvestorRange           byte
	BrokerID                string
	InvestorID              string
	OpenRatioByMoney        float64
	OpenRatioByVolume       float64
	CloseRatioByMoney       float64
	CloseRatioByVolume      float64
	CloseTodayRatioByMoney  float64
	CloseTodayRatioByVolume float64
	StrikeRatioByMoney      float64
	StrikeRatioByVolume     float64
	ExchangeID              string
	InvestUnitID            string
	InstrumentID            string
}

func NewCThostFtdcOptionInstrCommRateField(p *C.CThostFtdcOptionInstrCommRateField) *CThostFtdcOptionInstrCommRateField {
	ret := new(CThostFtdcOptionInstrCommRateField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.OpenRatioByMoney = goFloat64(p.OpenRatioByMoney)
	ret.OpenRatioByVolume = goFloat64(p.OpenRatioByVolume)
	ret.CloseRatioByMoney = goFloat64(p.CloseRatioByMoney)
	ret.CloseRatioByVolume = goFloat64(p.CloseRatioByVolume)
	ret.CloseTodayRatioByMoney = goFloat64(p.CloseTodayRatioByMoney)
	ret.CloseTodayRatioByVolume = goFloat64(p.CloseTodayRatioByVolume)
	ret.StrikeRatioByMoney = goFloat64(p.StrikeRatioByMoney)
	ret.StrikeRatioByVolume = goFloat64(p.StrikeRatioByVolume)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcOptionInstrCommRateField) CValue() *C.CThostFtdcOptionInstrCommRateField {
	ptr := C.newCThostFtdcOptionInstrCommRateField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.OpenRatioByMoney = C.double(s.OpenRatioByMoney)
	ptr.OpenRatioByVolume = C.double(s.OpenRatioByVolume)
	ptr.CloseRatioByMoney = C.double(s.CloseRatioByMoney)
	ptr.CloseRatioByVolume = C.double(s.CloseRatioByVolume)
	ptr.CloseTodayRatioByMoney = C.double(s.CloseTodayRatioByMoney)
	ptr.CloseTodayRatioByVolume = C.double(s.CloseTodayRatioByVolume)
	ptr.StrikeRatioByMoney = C.double(s.StrikeRatioByMoney)
	ptr.StrikeRatioByVolume = C.double(s.StrikeRatioByVolume)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcOptionInstrTradeCostField struct {
	BrokerID        string
	InvestorID      string
	Reserve1        string
	HedgeFlag       byte
	FixedMargin     float64
	MiniMargin      float64
	Royalty         float64
	ExchFixedMargin float64
	ExchMiniMargin  float64
	ExchangeID      string
	InvestUnitID    string
	InstrumentID    string
}

func NewCThostFtdcOptionInstrTradeCostField(p *C.CThostFtdcOptionInstrTradeCostField) *CThostFtdcOptionInstrTradeCostField {
	ret := new(CThostFtdcOptionInstrTradeCostField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.FixedMargin = goFloat64(p.FixedMargin)
	ret.MiniMargin = goFloat64(p.MiniMargin)
	ret.Royalty = goFloat64(p.Royalty)
	ret.ExchFixedMargin = goFloat64(p.ExchFixedMargin)
	ret.ExchMiniMargin = goFloat64(p.ExchMiniMargin)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcOptionInstrTradeCostField) CValue() *C.CThostFtdcOptionInstrTradeCostField {
	ptr := C.newCThostFtdcOptionInstrTradeCostField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.FixedMargin = C.double(s.FixedMargin)
	ptr.MiniMargin = C.double(s.MiniMargin)
	ptr.Royalty = C.double(s.Royalty)
	ptr.ExchFixedMargin = C.double(s.ExchFixedMargin)
	ptr.ExchMiniMargin = C.double(s.ExchMiniMargin)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryOptionInstrTradeCostField struct {
	BrokerID        string
	InvestorID      string
	Reserve1        string
	HedgeFlag       byte
	InputPrice      float64
	UnderlyingPrice float64
	ExchangeID      string
	InvestUnitID    string
	InstrumentID    string
}

func NewCThostFtdcQryOptionInstrTradeCostField(p *C.CThostFtdcQryOptionInstrTradeCostField) *CThostFtdcQryOptionInstrTradeCostField {
	ret := new(CThostFtdcQryOptionInstrTradeCostField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.InputPrice = goFloat64(p.InputPrice)
	ret.UnderlyingPrice = goFloat64(p.UnderlyingPrice)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryOptionInstrTradeCostField) CValue() *C.CThostFtdcQryOptionInstrTradeCostField {
	ptr := C.newCThostFtdcQryOptionInstrTradeCostField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.InputPrice = C.double(s.InputPrice)
	ptr.UnderlyingPrice = C.double(s.UnderlyingPrice)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryOptionInstrCommRateField struct {
	BrokerID     string
	InvestorID   string
	Reserve1     string
	ExchangeID   string
	InvestUnitID string
	InstrumentID string
}

func NewCThostFtdcQryOptionInstrCommRateField(p *C.CThostFtdcQryOptionInstrCommRateField) *CThostFtdcQryOptionInstrCommRateField {
	ret := new(CThostFtdcQryOptionInstrCommRateField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryOptionInstrCommRateField) CValue() *C.CThostFtdcQryOptionInstrCommRateField {
	ptr := C.newCThostFtdcQryOptionInstrCommRateField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcIndexPriceField struct {
	BrokerID     string
	Reserve1     string
	ClosePrice   float64
	InstrumentID string
}

func NewCThostFtdcIndexPriceField(p *C.CThostFtdcIndexPriceField) *CThostFtdcIndexPriceField {
	ret := new(CThostFtdcIndexPriceField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ClosePrice = goFloat64(p.ClosePrice)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcIndexPriceField) CValue() *C.CThostFtdcIndexPriceField {
	ptr := C.newCThostFtdcIndexPriceField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.ClosePrice = C.double(s.ClosePrice)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcInputExecOrderField struct {
	BrokerID            string
	InvestorID          string
	Reserve1            string
	ExecOrderRef        string
	UserID              string
	Volume              int
	RequestID           int
	BusinessUnit        string
	OffsetFlag          byte
	HedgeFlag           byte
	ActionType          byte
	PosiDirection       byte
	ReservePositionFlag byte
	CloseFlag           byte
	ExchangeID          string
	InvestUnitID        string
	AccountID           string
	CurrencyID          string
	ClientID            string
	Reserve2            string
	MacAddress          string
	InstrumentID        string
	IPAddress           string
}

func NewCThostFtdcInputExecOrderField(p *C.CThostFtdcInputExecOrderField) *CThostFtdcInputExecOrderField {
	ret := new(CThostFtdcInputExecOrderField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExecOrderRef = c2goStr(&p.ExecOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Volume = int(p.Volume)
	ret.RequestID = int(p.RequestID)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.OffsetFlag = byte(p.OffsetFlag)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.ActionType = byte(p.ActionType)
	ret.PosiDirection = byte(p.PosiDirection)
	ret.ReservePositionFlag = byte(p.ReservePositionFlag)
	ret.CloseFlag = byte(p.CloseFlag)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcInputExecOrderField) CValue() *C.CThostFtdcInputExecOrderField {
	ptr := C.newCThostFtdcInputExecOrderField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExecOrderRef, &ptr.ExecOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.Volume = C.int(s.Volume)
	ptr.RequestID = C.int(s.RequestID)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.OffsetFlag = C.char(s.OffsetFlag)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.ActionType = C.char(s.ActionType)
	ptr.PosiDirection = C.char(s.PosiDirection)
	ptr.ReservePositionFlag = C.char(s.ReservePositionFlag)
	ptr.CloseFlag = C.char(s.CloseFlag)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcInputExecOrderActionField struct {
	BrokerID           string
	InvestorID         string
	ExecOrderActionRef int
	ExecOrderRef       string
	RequestID          int
	FrontID            int
	SessionID          int
	ExchangeID         string
	ExecOrderSysID     string
	ActionFlag         byte
	UserID             string
	Reserve1           string
	InvestUnitID       string
	Reserve2           string
	MacAddress         string
	InstrumentID       string
	IPAddress          string
}

func NewCThostFtdcInputExecOrderActionField(p *C.CThostFtdcInputExecOrderActionField) *CThostFtdcInputExecOrderActionField {
	ret := new(CThostFtdcInputExecOrderActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.ExecOrderActionRef = int(p.ExecOrderActionRef)
	ret.ExecOrderRef = c2goStr(&p.ExecOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.RequestID = int(p.RequestID)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ExecOrderSysID = c2goStr(&p.ExecOrderSysID[0], C.sizeof_TThostFtdcExecOrderSysIDType)
	ret.ActionFlag = byte(p.ActionFlag)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcInputExecOrderActionField) CValue() *C.CThostFtdcInputExecOrderActionField {
	ptr := C.newCThostFtdcInputExecOrderActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.ExecOrderActionRef = C.int(s.ExecOrderActionRef)
	go2cStr(s.ExecOrderRef, &ptr.ExecOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ExecOrderSysID, &ptr.ExecOrderSysID[0], C.sizeof_TThostFtdcExecOrderSysIDType)
	ptr.ActionFlag = C.char(s.ActionFlag)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcExecOrderField struct {
	BrokerID            string
	InvestorID          string
	Reserve1            string
	ExecOrderRef        string
	UserID              string
	Volume              int
	RequestID           int
	BusinessUnit        string
	OffsetFlag          byte
	HedgeFlag           byte
	ActionType          byte
	PosiDirection       byte
	ReservePositionFlag byte
	CloseFlag           byte
	ExecOrderLocalID    string
	ExchangeID          string
	ParticipantID       string
	ClientID            string
	Reserve2            string
	TraderID            string
	InstallID           int
	OrderSubmitStatus   byte
	NotifySequence      int
	TradingDay          string
	SettlementID        int
	ExecOrderSysID      string
	InsertDate          string
	InsertTime          string
	CancelTime          string
	ExecResult          byte
	ClearingPartID      string
	SequenceNo          int
	FrontID             int
	SessionID           int
	UserProductInfo     string
	StatusMsg           string
	ActiveUserID        string
	BrokerExecOrderSeq  int
	BranchID            string
	InvestUnitID        string
	AccountID           string
	CurrencyID          string
	Reserve3            string
	MacAddress          string
	InstrumentID        string
	ExchangeInstID      string
	IPAddress           string
}

func NewCThostFtdcExecOrderField(p *C.CThostFtdcExecOrderField) *CThostFtdcExecOrderField {
	ret := new(CThostFtdcExecOrderField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExecOrderRef = c2goStr(&p.ExecOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Volume = int(p.Volume)
	ret.RequestID = int(p.RequestID)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.OffsetFlag = byte(p.OffsetFlag)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.ActionType = byte(p.ActionType)
	ret.PosiDirection = byte(p.PosiDirection)
	ret.ReservePositionFlag = byte(p.ReservePositionFlag)
	ret.CloseFlag = byte(p.CloseFlag)
	ret.ExecOrderLocalID = c2goStr(&p.ExecOrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.OrderSubmitStatus = byte(p.OrderSubmitStatus)
	ret.NotifySequence = int(p.NotifySequence)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.ExecOrderSysID = c2goStr(&p.ExecOrderSysID[0], C.sizeof_TThostFtdcExecOrderSysIDType)
	ret.InsertDate = c2goStr(&p.InsertDate[0], C.sizeof_TThostFtdcDateType)
	ret.InsertTime = c2goStr(&p.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	ret.CancelTime = c2goStr(&p.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	ret.ExecResult = byte(p.ExecResult)
	ret.ClearingPartID = c2goStr(&p.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.SequenceNo = int(p.SequenceNo)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.UserProductInfo = c2goStr(&p.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.StatusMsg = c2goStr(&p.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.ActiveUserID = c2goStr(&p.ActiveUserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.BrokerExecOrderSeq = int(p.BrokerExecOrderSeq)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.Reserve3 = c2goStr(&p.reserve3[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcExecOrderField) CValue() *C.CThostFtdcExecOrderField {
	ptr := C.newCThostFtdcExecOrderField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExecOrderRef, &ptr.ExecOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.Volume = C.int(s.Volume)
	ptr.RequestID = C.int(s.RequestID)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.OffsetFlag = C.char(s.OffsetFlag)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.ActionType = C.char(s.ActionType)
	ptr.PosiDirection = C.char(s.PosiDirection)
	ptr.ReservePositionFlag = C.char(s.ReservePositionFlag)
	ptr.CloseFlag = C.char(s.CloseFlag)
	go2cStr(s.ExecOrderLocalID, &ptr.ExecOrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.OrderSubmitStatus = C.char(s.OrderSubmitStatus)
	ptr.NotifySequence = C.int(s.NotifySequence)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	go2cStr(s.ExecOrderSysID, &ptr.ExecOrderSysID[0], C.sizeof_TThostFtdcExecOrderSysIDType)
	go2cStr(s.InsertDate, &ptr.InsertDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.InsertTime, &ptr.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.CancelTime, &ptr.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.ExecResult = C.char(s.ExecResult)
	go2cStr(s.ClearingPartID, &ptr.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ptr.SequenceNo = C.int(s.SequenceNo)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.UserProductInfo, &ptr.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.StatusMsg, &ptr.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.ActiveUserID, &ptr.ActiveUserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.BrokerExecOrderSeq = C.int(s.BrokerExecOrderSeq)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.Reserve3, &ptr.reserve3[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcExecOrderActionField struct {
	BrokerID           string
	InvestorID         string
	ExecOrderActionRef int
	ExecOrderRef       string
	RequestID          int
	FrontID            int
	SessionID          int
	ExchangeID         string
	ExecOrderSysID     string
	ActionFlag         byte
	ActionDate         string
	ActionTime         string
	TraderID           string
	InstallID          int
	ExecOrderLocalID   string
	ActionLocalID      string
	ParticipantID      string
	ClientID           string
	BusinessUnit       string
	OrderActionStatus  byte
	UserID             string
	ActionType         byte
	StatusMsg          string
	Reserve1           string
	BranchID           string
	InvestUnitID       string
	Reserve2           string
	MacAddress         string
	InstrumentID       string
	IPAddress          string
}

func NewCThostFtdcExecOrderActionField(p *C.CThostFtdcExecOrderActionField) *CThostFtdcExecOrderActionField {
	ret := new(CThostFtdcExecOrderActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.ExecOrderActionRef = int(p.ExecOrderActionRef)
	ret.ExecOrderRef = c2goStr(&p.ExecOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.RequestID = int(p.RequestID)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ExecOrderSysID = c2goStr(&p.ExecOrderSysID[0], C.sizeof_TThostFtdcExecOrderSysIDType)
	ret.ActionFlag = byte(p.ActionFlag)
	ret.ActionDate = c2goStr(&p.ActionDate[0], C.sizeof_TThostFtdcDateType)
	ret.ActionTime = c2goStr(&p.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.ExecOrderLocalID = c2goStr(&p.ExecOrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ActionLocalID = c2goStr(&p.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.OrderActionStatus = byte(p.OrderActionStatus)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.ActionType = byte(p.ActionType)
	ret.StatusMsg = c2goStr(&p.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcExecOrderActionField) CValue() *C.CThostFtdcExecOrderActionField {
	ptr := C.newCThostFtdcExecOrderActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.ExecOrderActionRef = C.int(s.ExecOrderActionRef)
	go2cStr(s.ExecOrderRef, &ptr.ExecOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ExecOrderSysID, &ptr.ExecOrderSysID[0], C.sizeof_TThostFtdcExecOrderSysIDType)
	ptr.ActionFlag = C.char(s.ActionFlag)
	go2cStr(s.ActionDate, &ptr.ActionDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ActionTime, &ptr.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.ExecOrderLocalID, &ptr.ExecOrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ActionLocalID, &ptr.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.OrderActionStatus = C.char(s.OrderActionStatus)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.ActionType = C.char(s.ActionType)
	go2cStr(s.StatusMsg, &ptr.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryExecOrderField struct {
	BrokerID        string
	InvestorID      string
	Reserve1        string
	ExchangeID      string
	ExecOrderSysID  string
	InsertTimeStart string
	InsertTimeEnd   string
	InstrumentID    string
}

func NewCThostFtdcQryExecOrderField(p *C.CThostFtdcQryExecOrderField) *CThostFtdcQryExecOrderField {
	ret := new(CThostFtdcQryExecOrderField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ExecOrderSysID = c2goStr(&p.ExecOrderSysID[0], C.sizeof_TThostFtdcExecOrderSysIDType)
	ret.InsertTimeStart = c2goStr(&p.InsertTimeStart[0], C.sizeof_TThostFtdcTimeType)
	ret.InsertTimeEnd = c2goStr(&p.InsertTimeEnd[0], C.sizeof_TThostFtdcTimeType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryExecOrderField) CValue() *C.CThostFtdcQryExecOrderField {
	ptr := C.newCThostFtdcQryExecOrderField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ExecOrderSysID, &ptr.ExecOrderSysID[0], C.sizeof_TThostFtdcExecOrderSysIDType)
	go2cStr(s.InsertTimeStart, &ptr.InsertTimeStart[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.InsertTimeEnd, &ptr.InsertTimeEnd[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcExchangeExecOrderField struct {
	Volume              int
	RequestID           int
	BusinessUnit        string
	OffsetFlag          byte
	HedgeFlag           byte
	ActionType          byte
	PosiDirection       byte
	ReservePositionFlag byte
	CloseFlag           byte
	ExecOrderLocalID    string
	ExchangeID          string
	ParticipantID       string
	ClientID            string
	Reserve1            string
	TraderID            string
	InstallID           int
	OrderSubmitStatus   byte
	NotifySequence      int
	TradingDay          string
	SettlementID        int
	ExecOrderSysID      string
	InsertDate          string
	InsertTime          string
	CancelTime          string
	ExecResult          byte
	ClearingPartID      string
	SequenceNo          int
	BranchID            string
	Reserve2            string
	MacAddress          string
	ExchangeInstID      string
	IPAddress           string
}

func NewCThostFtdcExchangeExecOrderField(p *C.CThostFtdcExchangeExecOrderField) *CThostFtdcExchangeExecOrderField {
	ret := new(CThostFtdcExchangeExecOrderField)
	ret.Volume = int(p.Volume)
	ret.RequestID = int(p.RequestID)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.OffsetFlag = byte(p.OffsetFlag)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.ActionType = byte(p.ActionType)
	ret.PosiDirection = byte(p.PosiDirection)
	ret.ReservePositionFlag = byte(p.ReservePositionFlag)
	ret.CloseFlag = byte(p.CloseFlag)
	ret.ExecOrderLocalID = c2goStr(&p.ExecOrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.OrderSubmitStatus = byte(p.OrderSubmitStatus)
	ret.NotifySequence = int(p.NotifySequence)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.ExecOrderSysID = c2goStr(&p.ExecOrderSysID[0], C.sizeof_TThostFtdcExecOrderSysIDType)
	ret.InsertDate = c2goStr(&p.InsertDate[0], C.sizeof_TThostFtdcDateType)
	ret.InsertTime = c2goStr(&p.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	ret.CancelTime = c2goStr(&p.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	ret.ExecResult = byte(p.ExecResult)
	ret.ClearingPartID = c2goStr(&p.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.SequenceNo = int(p.SequenceNo)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcExchangeExecOrderField) CValue() *C.CThostFtdcExchangeExecOrderField {
	ptr := C.newCThostFtdcExchangeExecOrderField()
	ptr.Volume = C.int(s.Volume)
	ptr.RequestID = C.int(s.RequestID)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.OffsetFlag = C.char(s.OffsetFlag)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.ActionType = C.char(s.ActionType)
	ptr.PosiDirection = C.char(s.PosiDirection)
	ptr.ReservePositionFlag = C.char(s.ReservePositionFlag)
	ptr.CloseFlag = C.char(s.CloseFlag)
	go2cStr(s.ExecOrderLocalID, &ptr.ExecOrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.OrderSubmitStatus = C.char(s.OrderSubmitStatus)
	ptr.NotifySequence = C.int(s.NotifySequence)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	go2cStr(s.ExecOrderSysID, &ptr.ExecOrderSysID[0], C.sizeof_TThostFtdcExecOrderSysIDType)
	go2cStr(s.InsertDate, &ptr.InsertDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.InsertTime, &ptr.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.CancelTime, &ptr.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.ExecResult = C.char(s.ExecResult)
	go2cStr(s.ClearingPartID, &ptr.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ptr.SequenceNo = C.int(s.SequenceNo)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryExchangeExecOrderField struct {
	ParticipantID  string
	ClientID       string
	Reserve1       string
	ExchangeID     string
	TraderID       string
	ExchangeInstID string
}

func NewCThostFtdcQryExchangeExecOrderField(p *C.CThostFtdcQryExchangeExecOrderField) *CThostFtdcQryExchangeExecOrderField {
	ret := new(CThostFtdcQryExchangeExecOrderField)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ret
}
func (s *CThostFtdcQryExchangeExecOrderField) CValue() *C.CThostFtdcQryExchangeExecOrderField {
	ptr := C.newCThostFtdcQryExchangeExecOrderField()
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ptr
}

type CThostFtdcQryExecOrderActionField struct {
	BrokerID   string
	InvestorID string
	ExchangeID string
}

func NewCThostFtdcQryExecOrderActionField(p *C.CThostFtdcQryExecOrderActionField) *CThostFtdcQryExecOrderActionField {
	ret := new(CThostFtdcQryExecOrderActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ret
}
func (s *CThostFtdcQryExecOrderActionField) CValue() *C.CThostFtdcQryExecOrderActionField {
	ptr := C.newCThostFtdcQryExecOrderActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ptr
}

type CThostFtdcExchangeExecOrderActionField struct {
	ExchangeID        string
	ExecOrderSysID    string
	ActionFlag        byte
	ActionDate        string
	ActionTime        string
	TraderID          string
	InstallID         int
	ExecOrderLocalID  string
	ActionLocalID     string
	ParticipantID     string
	ClientID          string
	BusinessUnit      string
	OrderActionStatus byte
	UserID            string
	ActionType        byte
	BranchID          string
	Reserve1          string
	MacAddress        string
	Reserve2          string
	Volume            int
	IPAddress         string
	ExchangeInstID    string
}

func NewCThostFtdcExchangeExecOrderActionField(p *C.CThostFtdcExchangeExecOrderActionField) *CThostFtdcExchangeExecOrderActionField {
	ret := new(CThostFtdcExchangeExecOrderActionField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ExecOrderSysID = c2goStr(&p.ExecOrderSysID[0], C.sizeof_TThostFtdcExecOrderSysIDType)
	ret.ActionFlag = byte(p.ActionFlag)
	ret.ActionDate = c2goStr(&p.ActionDate[0], C.sizeof_TThostFtdcDateType)
	ret.ActionTime = c2goStr(&p.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.ExecOrderLocalID = c2goStr(&p.ExecOrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ActionLocalID = c2goStr(&p.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.OrderActionStatus = byte(p.OrderActionStatus)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.ActionType = byte(p.ActionType)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.Volume = int(p.Volume)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ret
}
func (s *CThostFtdcExchangeExecOrderActionField) CValue() *C.CThostFtdcExchangeExecOrderActionField {
	ptr := C.newCThostFtdcExchangeExecOrderActionField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ExecOrderSysID, &ptr.ExecOrderSysID[0], C.sizeof_TThostFtdcExecOrderSysIDType)
	ptr.ActionFlag = C.char(s.ActionFlag)
	go2cStr(s.ActionDate, &ptr.ActionDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ActionTime, &ptr.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.ExecOrderLocalID, &ptr.ExecOrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ActionLocalID, &ptr.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.OrderActionStatus = C.char(s.OrderActionStatus)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.ActionType = C.char(s.ActionType)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ptr.Volume = C.int(s.Volume)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ptr
}

type CThostFtdcQryExchangeExecOrderActionField struct {
	ParticipantID string
	ClientID      string
	ExchangeID    string
	TraderID      string
}

func NewCThostFtdcQryExchangeExecOrderActionField(p *C.CThostFtdcQryExchangeExecOrderActionField) *CThostFtdcQryExchangeExecOrderActionField {
	ret := new(CThostFtdcQryExchangeExecOrderActionField)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	return ret
}
func (s *CThostFtdcQryExchangeExecOrderActionField) CValue() *C.CThostFtdcQryExchangeExecOrderActionField {
	ptr := C.newCThostFtdcQryExchangeExecOrderActionField()
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	return ptr
}

type CThostFtdcErrExecOrderField struct {
	BrokerID            string
	InvestorID          string
	Reserve1            string
	ExecOrderRef        string
	UserID              string
	Volume              int
	RequestID           int
	BusinessUnit        string
	OffsetFlag          byte
	HedgeFlag           byte
	ActionType          byte
	PosiDirection       byte
	ReservePositionFlag byte
	CloseFlag           byte
	ExchangeID          string
	InvestUnitID        string
	AccountID           string
	CurrencyID          string
	ClientID            string
	Reserve2            string
	MacAddress          string
	ErrorID             int
	ErrorMsg            string
	InstrumentID        string
	IPAddress           string
}

func NewCThostFtdcErrExecOrderField(p *C.CThostFtdcErrExecOrderField) *CThostFtdcErrExecOrderField {
	ret := new(CThostFtdcErrExecOrderField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExecOrderRef = c2goStr(&p.ExecOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Volume = int(p.Volume)
	ret.RequestID = int(p.RequestID)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.OffsetFlag = byte(p.OffsetFlag)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.ActionType = byte(p.ActionType)
	ret.PosiDirection = byte(p.PosiDirection)
	ret.ReservePositionFlag = byte(p.ReservePositionFlag)
	ret.CloseFlag = byte(p.CloseFlag)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcErrExecOrderField) CValue() *C.CThostFtdcErrExecOrderField {
	ptr := C.newCThostFtdcErrExecOrderField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExecOrderRef, &ptr.ExecOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.Volume = C.int(s.Volume)
	ptr.RequestID = C.int(s.RequestID)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.OffsetFlag = C.char(s.OffsetFlag)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.ActionType = C.char(s.ActionType)
	ptr.PosiDirection = C.char(s.PosiDirection)
	ptr.ReservePositionFlag = C.char(s.ReservePositionFlag)
	ptr.CloseFlag = C.char(s.CloseFlag)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryErrExecOrderField struct {
	BrokerID   string
	InvestorID string
}

func NewCThostFtdcQryErrExecOrderField(p *C.CThostFtdcQryErrExecOrderField) *CThostFtdcQryErrExecOrderField {
	ret := new(CThostFtdcQryErrExecOrderField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ret
}
func (s *CThostFtdcQryErrExecOrderField) CValue() *C.CThostFtdcQryErrExecOrderField {
	ptr := C.newCThostFtdcQryErrExecOrderField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ptr
}

type CThostFtdcErrExecOrderActionField struct {
	BrokerID           string
	InvestorID         string
	ExecOrderActionRef int
	ExecOrderRef       string
	RequestID          int
	FrontID            int
	SessionID          int
	ExchangeID         string
	ExecOrderSysID     string
	ActionFlag         byte
	UserID             string
	Reserve1           string
	InvestUnitID       string
	Reserve2           string
	MacAddress         string
	ErrorID            int
	ErrorMsg           string
	InstrumentID       string
	IPAddress          string
}

func NewCThostFtdcErrExecOrderActionField(p *C.CThostFtdcErrExecOrderActionField) *CThostFtdcErrExecOrderActionField {
	ret := new(CThostFtdcErrExecOrderActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.ExecOrderActionRef = int(p.ExecOrderActionRef)
	ret.ExecOrderRef = c2goStr(&p.ExecOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.RequestID = int(p.RequestID)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ExecOrderSysID = c2goStr(&p.ExecOrderSysID[0], C.sizeof_TThostFtdcExecOrderSysIDType)
	ret.ActionFlag = byte(p.ActionFlag)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcErrExecOrderActionField) CValue() *C.CThostFtdcErrExecOrderActionField {
	ptr := C.newCThostFtdcErrExecOrderActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.ExecOrderActionRef = C.int(s.ExecOrderActionRef)
	go2cStr(s.ExecOrderRef, &ptr.ExecOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ExecOrderSysID, &ptr.ExecOrderSysID[0], C.sizeof_TThostFtdcExecOrderSysIDType)
	ptr.ActionFlag = C.char(s.ActionFlag)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryErrExecOrderActionField struct {
	BrokerID   string
	InvestorID string
}

func NewCThostFtdcQryErrExecOrderActionField(p *C.CThostFtdcQryErrExecOrderActionField) *CThostFtdcQryErrExecOrderActionField {
	ret := new(CThostFtdcQryErrExecOrderActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ret
}
func (s *CThostFtdcQryErrExecOrderActionField) CValue() *C.CThostFtdcQryErrExecOrderActionField {
	ptr := C.newCThostFtdcQryErrExecOrderActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ptr
}

type CThostFtdcOptionInstrTradingRightField struct {
	Reserve1      string
	InvestorRange byte
	BrokerID      string
	InvestorID    string
	Direction     byte
	TradingRight  byte
	InstrumentID  string
}

func NewCThostFtdcOptionInstrTradingRightField(p *C.CThostFtdcOptionInstrTradingRightField) *CThostFtdcOptionInstrTradingRightField {
	ret := new(CThostFtdcOptionInstrTradingRightField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Direction = byte(p.Direction)
	ret.TradingRight = byte(p.TradingRight)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcOptionInstrTradingRightField) CValue() *C.CThostFtdcOptionInstrTradingRightField {
	ptr := C.newCThostFtdcOptionInstrTradingRightField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.Direction = C.char(s.Direction)
	ptr.TradingRight = C.char(s.TradingRight)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryOptionInstrTradingRightField struct {
	BrokerID     string
	InvestorID   string
	Reserve1     string
	Direction    byte
	InstrumentID string
}

func NewCThostFtdcQryOptionInstrTradingRightField(p *C.CThostFtdcQryOptionInstrTradingRightField) *CThostFtdcQryOptionInstrTradingRightField {
	ret := new(CThostFtdcQryOptionInstrTradingRightField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.Direction = byte(p.Direction)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryOptionInstrTradingRightField) CValue() *C.CThostFtdcQryOptionInstrTradingRightField {
	ptr := C.newCThostFtdcQryOptionInstrTradingRightField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.Direction = C.char(s.Direction)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcInputForQuoteField struct {
	BrokerID     string
	InvestorID   string
	Reserve1     string
	ForQuoteRef  string
	UserID       string
	ExchangeID   string
	InvestUnitID string
	Reserve2     string
	MacAddress   string
	InstrumentID string
	IPAddress    string
}

func NewCThostFtdcInputForQuoteField(p *C.CThostFtdcInputForQuoteField) *CThostFtdcInputForQuoteField {
	ret := new(CThostFtdcInputForQuoteField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ForQuoteRef = c2goStr(&p.ForQuoteRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcInputForQuoteField) CValue() *C.CThostFtdcInputForQuoteField {
	ptr := C.newCThostFtdcInputForQuoteField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ForQuoteRef, &ptr.ForQuoteRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcForQuoteField struct {
	BrokerID         string
	InvestorID       string
	Reserve1         string
	ForQuoteRef      string
	UserID           string
	ForQuoteLocalID  string
	ExchangeID       string
	ParticipantID    string
	ClientID         string
	Reserve2         string
	TraderID         string
	InstallID        int
	InsertDate       string
	InsertTime       string
	ForQuoteStatus   byte
	FrontID          int
	SessionID        int
	StatusMsg        string
	ActiveUserID     string
	BrokerForQutoSeq int
	InvestUnitID     string
	Reserve3         string
	MacAddress       string
	InstrumentID     string
	ExchangeInstID   string
	IPAddress        string
}

func NewCThostFtdcForQuoteField(p *C.CThostFtdcForQuoteField) *CThostFtdcForQuoteField {
	ret := new(CThostFtdcForQuoteField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ForQuoteRef = c2goStr(&p.ForQuoteRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.ForQuoteLocalID = c2goStr(&p.ForQuoteLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.InsertDate = c2goStr(&p.InsertDate[0], C.sizeof_TThostFtdcDateType)
	ret.InsertTime = c2goStr(&p.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	ret.ForQuoteStatus = byte(p.ForQuoteStatus)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.StatusMsg = c2goStr(&p.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.ActiveUserID = c2goStr(&p.ActiveUserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.BrokerForQutoSeq = int(p.BrokerForQutoSeq)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.Reserve3 = c2goStr(&p.reserve3[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcForQuoteField) CValue() *C.CThostFtdcForQuoteField {
	ptr := C.newCThostFtdcForQuoteField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ForQuoteRef, &ptr.ForQuoteRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.ForQuoteLocalID, &ptr.ForQuoteLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.InsertDate, &ptr.InsertDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.InsertTime, &ptr.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.ForQuoteStatus = C.char(s.ForQuoteStatus)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.StatusMsg, &ptr.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.ActiveUserID, &ptr.ActiveUserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.BrokerForQutoSeq = C.int(s.BrokerForQutoSeq)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.Reserve3, &ptr.reserve3[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryForQuoteField struct {
	BrokerID        string
	InvestorID      string
	Reserve1        string
	ExchangeID      string
	InsertTimeStart string
	InsertTimeEnd   string
	InvestUnitID    string
	InstrumentID    string
}

func NewCThostFtdcQryForQuoteField(p *C.CThostFtdcQryForQuoteField) *CThostFtdcQryForQuoteField {
	ret := new(CThostFtdcQryForQuoteField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InsertTimeStart = c2goStr(&p.InsertTimeStart[0], C.sizeof_TThostFtdcTimeType)
	ret.InsertTimeEnd = c2goStr(&p.InsertTimeEnd[0], C.sizeof_TThostFtdcTimeType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryForQuoteField) CValue() *C.CThostFtdcQryForQuoteField {
	ptr := C.newCThostFtdcQryForQuoteField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InsertTimeStart, &ptr.InsertTimeStart[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.InsertTimeEnd, &ptr.InsertTimeEnd[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcExchangeForQuoteField struct {
	ForQuoteLocalID string
	ExchangeID      string
	ParticipantID   string
	ClientID        string
	Reserve1        string
	TraderID        string
	InstallID       int
	InsertDate      string
	InsertTime      string
	ForQuoteStatus  byte
	Reserve2        string
	MacAddress      string
	ExchangeInstID  string
	IPAddress       string
}

func NewCThostFtdcExchangeForQuoteField(p *C.CThostFtdcExchangeForQuoteField) *CThostFtdcExchangeForQuoteField {
	ret := new(CThostFtdcExchangeForQuoteField)
	ret.ForQuoteLocalID = c2goStr(&p.ForQuoteLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.InsertDate = c2goStr(&p.InsertDate[0], C.sizeof_TThostFtdcDateType)
	ret.InsertTime = c2goStr(&p.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	ret.ForQuoteStatus = byte(p.ForQuoteStatus)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcExchangeForQuoteField) CValue() *C.CThostFtdcExchangeForQuoteField {
	ptr := C.newCThostFtdcExchangeForQuoteField()
	go2cStr(s.ForQuoteLocalID, &ptr.ForQuoteLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.InsertDate, &ptr.InsertDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.InsertTime, &ptr.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.ForQuoteStatus = C.char(s.ForQuoteStatus)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryExchangeForQuoteField struct {
	ParticipantID  string
	ClientID       string
	Reserve1       string
	ExchangeID     string
	TraderID       string
	ExchangeInstID string
}

func NewCThostFtdcQryExchangeForQuoteField(p *C.CThostFtdcQryExchangeForQuoteField) *CThostFtdcQryExchangeForQuoteField {
	ret := new(CThostFtdcQryExchangeForQuoteField)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ret
}
func (s *CThostFtdcQryExchangeForQuoteField) CValue() *C.CThostFtdcQryExchangeForQuoteField {
	ptr := C.newCThostFtdcQryExchangeForQuoteField()
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ptr
}

type CThostFtdcInputQuoteField struct {
	BrokerID      string
	InvestorID    string
	Reserve1      string
	QuoteRef      string
	UserID        string
	AskPrice      float64
	BidPrice      float64
	AskVolume     int
	BidVolume     int
	RequestID     int
	BusinessUnit  string
	AskOffsetFlag byte
	BidOffsetFlag byte
	AskHedgeFlag  byte
	BidHedgeFlag  byte
	AskOrderRef   string
	BidOrderRef   string
	ForQuoteSysID string
	ExchangeID    string
	InvestUnitID  string
	ClientID      string
	Reserve2      string
	MacAddress    string
	InstrumentID  string
	IPAddress     string
	ReplaceSysID  string
}

func NewCThostFtdcInputQuoteField(p *C.CThostFtdcInputQuoteField) *CThostFtdcInputQuoteField {
	ret := new(CThostFtdcInputQuoteField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.QuoteRef = c2goStr(&p.QuoteRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.AskPrice = goFloat64(p.AskPrice)
	ret.BidPrice = goFloat64(p.BidPrice)
	ret.AskVolume = int(p.AskVolume)
	ret.BidVolume = int(p.BidVolume)
	ret.RequestID = int(p.RequestID)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.AskOffsetFlag = byte(p.AskOffsetFlag)
	ret.BidOffsetFlag = byte(p.BidOffsetFlag)
	ret.AskHedgeFlag = byte(p.AskHedgeFlag)
	ret.BidHedgeFlag = byte(p.BidHedgeFlag)
	ret.AskOrderRef = c2goStr(&p.AskOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.BidOrderRef = c2goStr(&p.BidOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.ForQuoteSysID = c2goStr(&p.ForQuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	ret.ReplaceSysID = c2goStr(&p.ReplaceSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	return ret
}
func (s *CThostFtdcInputQuoteField) CValue() *C.CThostFtdcInputQuoteField {
	ptr := C.newCThostFtdcInputQuoteField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.QuoteRef, &ptr.QuoteRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.AskPrice = C.double(s.AskPrice)
	ptr.BidPrice = C.double(s.BidPrice)
	ptr.AskVolume = C.int(s.AskVolume)
	ptr.BidVolume = C.int(s.BidVolume)
	ptr.RequestID = C.int(s.RequestID)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.AskOffsetFlag = C.char(s.AskOffsetFlag)
	ptr.BidOffsetFlag = C.char(s.BidOffsetFlag)
	ptr.AskHedgeFlag = C.char(s.AskHedgeFlag)
	ptr.BidHedgeFlag = C.char(s.BidHedgeFlag)
	go2cStr(s.AskOrderRef, &ptr.AskOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.BidOrderRef, &ptr.BidOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.ForQuoteSysID, &ptr.ForQuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	go2cStr(s.ReplaceSysID, &ptr.ReplaceSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	return ptr
}

type CThostFtdcInputQuoteActionField struct {
	BrokerID       string
	InvestorID     string
	QuoteActionRef int
	QuoteRef       string
	RequestID      int
	FrontID        int
	SessionID      int
	ExchangeID     string
	QuoteSysID     string
	ActionFlag     byte
	UserID         string
	Reserve1       string
	InvestUnitID   string
	ClientID       string
	Reserve2       string
	MacAddress     string
	InstrumentID   string
	IPAddress      string
}

func NewCThostFtdcInputQuoteActionField(p *C.CThostFtdcInputQuoteActionField) *CThostFtdcInputQuoteActionField {
	ret := new(CThostFtdcInputQuoteActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.QuoteActionRef = int(p.QuoteActionRef)
	ret.QuoteRef = c2goStr(&p.QuoteRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.RequestID = int(p.RequestID)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.QuoteSysID = c2goStr(&p.QuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ActionFlag = byte(p.ActionFlag)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcInputQuoteActionField) CValue() *C.CThostFtdcInputQuoteActionField {
	ptr := C.newCThostFtdcInputQuoteActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.QuoteActionRef = C.int(s.QuoteActionRef)
	go2cStr(s.QuoteRef, &ptr.QuoteRef[0], C.sizeof_TThostFtdcOrderRefType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.QuoteSysID, &ptr.QuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ptr.ActionFlag = C.char(s.ActionFlag)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQuoteField struct {
	BrokerID          string
	InvestorID        string
	Reserve1          string
	QuoteRef          string
	UserID            string
	AskPrice          float64
	BidPrice          float64
	AskVolume         int
	BidVolume         int
	RequestID         int
	BusinessUnit      string
	AskOffsetFlag     byte
	BidOffsetFlag     byte
	AskHedgeFlag      byte
	BidHedgeFlag      byte
	QuoteLocalID      string
	ExchangeID        string
	ParticipantID     string
	ClientID          string
	Reserve2          string
	TraderID          string
	InstallID         int
	NotifySequence    int
	OrderSubmitStatus byte
	TradingDay        string
	SettlementID      int
	QuoteSysID        string
	InsertDate        string
	InsertTime        string
	CancelTime        string
	QuoteStatus       byte
	ClearingPartID    string
	SequenceNo        int
	AskOrderSysID     string
	BidOrderSysID     string
	FrontID           int
	SessionID         int
	UserProductInfo   string
	StatusMsg         string
	ActiveUserID      string
	BrokerQuoteSeq    int
	AskOrderRef       string
	BidOrderRef       string
	ForQuoteSysID     string
	BranchID          string
	InvestUnitID      string
	AccountID         string
	CurrencyID        string
	Reserve3          string
	MacAddress        string
	InstrumentID      string
	ExchangeInstID    string
	IPAddress         string
	ReplaceSysID      string
}

func NewCThostFtdcQuoteField(p *C.CThostFtdcQuoteField) *CThostFtdcQuoteField {
	ret := new(CThostFtdcQuoteField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.QuoteRef = c2goStr(&p.QuoteRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.AskPrice = goFloat64(p.AskPrice)
	ret.BidPrice = goFloat64(p.BidPrice)
	ret.AskVolume = int(p.AskVolume)
	ret.BidVolume = int(p.BidVolume)
	ret.RequestID = int(p.RequestID)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.AskOffsetFlag = byte(p.AskOffsetFlag)
	ret.BidOffsetFlag = byte(p.BidOffsetFlag)
	ret.AskHedgeFlag = byte(p.AskHedgeFlag)
	ret.BidHedgeFlag = byte(p.BidHedgeFlag)
	ret.QuoteLocalID = c2goStr(&p.QuoteLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.NotifySequence = int(p.NotifySequence)
	ret.OrderSubmitStatus = byte(p.OrderSubmitStatus)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.QuoteSysID = c2goStr(&p.QuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.InsertDate = c2goStr(&p.InsertDate[0], C.sizeof_TThostFtdcDateType)
	ret.InsertTime = c2goStr(&p.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	ret.CancelTime = c2goStr(&p.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	ret.QuoteStatus = byte(p.QuoteStatus)
	ret.ClearingPartID = c2goStr(&p.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.SequenceNo = int(p.SequenceNo)
	ret.AskOrderSysID = c2goStr(&p.AskOrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.BidOrderSysID = c2goStr(&p.BidOrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.UserProductInfo = c2goStr(&p.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.StatusMsg = c2goStr(&p.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.ActiveUserID = c2goStr(&p.ActiveUserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.BrokerQuoteSeq = int(p.BrokerQuoteSeq)
	ret.AskOrderRef = c2goStr(&p.AskOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.BidOrderRef = c2goStr(&p.BidOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.ForQuoteSysID = c2goStr(&p.ForQuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.Reserve3 = c2goStr(&p.reserve3[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	ret.ReplaceSysID = c2goStr(&p.ReplaceSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	return ret
}
func (s *CThostFtdcQuoteField) CValue() *C.CThostFtdcQuoteField {
	ptr := C.newCThostFtdcQuoteField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.QuoteRef, &ptr.QuoteRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.AskPrice = C.double(s.AskPrice)
	ptr.BidPrice = C.double(s.BidPrice)
	ptr.AskVolume = C.int(s.AskVolume)
	ptr.BidVolume = C.int(s.BidVolume)
	ptr.RequestID = C.int(s.RequestID)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.AskOffsetFlag = C.char(s.AskOffsetFlag)
	ptr.BidOffsetFlag = C.char(s.BidOffsetFlag)
	ptr.AskHedgeFlag = C.char(s.AskHedgeFlag)
	ptr.BidHedgeFlag = C.char(s.BidHedgeFlag)
	go2cStr(s.QuoteLocalID, &ptr.QuoteLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.NotifySequence = C.int(s.NotifySequence)
	ptr.OrderSubmitStatus = C.char(s.OrderSubmitStatus)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	go2cStr(s.QuoteSysID, &ptr.QuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.InsertDate, &ptr.InsertDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.InsertTime, &ptr.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.CancelTime, &ptr.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.QuoteStatus = C.char(s.QuoteStatus)
	go2cStr(s.ClearingPartID, &ptr.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ptr.SequenceNo = C.int(s.SequenceNo)
	go2cStr(s.AskOrderSysID, &ptr.AskOrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.BidOrderSysID, &ptr.BidOrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.UserProductInfo, &ptr.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.StatusMsg, &ptr.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.ActiveUserID, &ptr.ActiveUserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.BrokerQuoteSeq = C.int(s.BrokerQuoteSeq)
	go2cStr(s.AskOrderRef, &ptr.AskOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.BidOrderRef, &ptr.BidOrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.ForQuoteSysID, &ptr.ForQuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.Reserve3, &ptr.reserve3[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	go2cStr(s.ReplaceSysID, &ptr.ReplaceSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	return ptr
}

type CThostFtdcQuoteActionField struct {
	BrokerID          string
	InvestorID        string
	QuoteActionRef    int
	QuoteRef          string
	RequestID         int
	FrontID           int
	SessionID         int
	ExchangeID        string
	QuoteSysID        string
	ActionFlag        byte
	ActionDate        string
	ActionTime        string
	TraderID          string
	InstallID         int
	QuoteLocalID      string
	ActionLocalID     string
	ParticipantID     string
	ClientID          string
	BusinessUnit      string
	OrderActionStatus byte
	UserID            string
	StatusMsg         string
	Reserve1          string
	BranchID          string
	InvestUnitID      string
	Reserve2          string
	MacAddress        string
	InstrumentID      string
	IPAddress         string
}

func NewCThostFtdcQuoteActionField(p *C.CThostFtdcQuoteActionField) *CThostFtdcQuoteActionField {
	ret := new(CThostFtdcQuoteActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.QuoteActionRef = int(p.QuoteActionRef)
	ret.QuoteRef = c2goStr(&p.QuoteRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.RequestID = int(p.RequestID)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.QuoteSysID = c2goStr(&p.QuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ActionFlag = byte(p.ActionFlag)
	ret.ActionDate = c2goStr(&p.ActionDate[0], C.sizeof_TThostFtdcDateType)
	ret.ActionTime = c2goStr(&p.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.QuoteLocalID = c2goStr(&p.QuoteLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ActionLocalID = c2goStr(&p.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.OrderActionStatus = byte(p.OrderActionStatus)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.StatusMsg = c2goStr(&p.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcQuoteActionField) CValue() *C.CThostFtdcQuoteActionField {
	ptr := C.newCThostFtdcQuoteActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.QuoteActionRef = C.int(s.QuoteActionRef)
	go2cStr(s.QuoteRef, &ptr.QuoteRef[0], C.sizeof_TThostFtdcOrderRefType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.QuoteSysID, &ptr.QuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ptr.ActionFlag = C.char(s.ActionFlag)
	go2cStr(s.ActionDate, &ptr.ActionDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ActionTime, &ptr.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.QuoteLocalID, &ptr.QuoteLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ActionLocalID, &ptr.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.OrderActionStatus = C.char(s.OrderActionStatus)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.StatusMsg, &ptr.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryQuoteField struct {
	BrokerID        string
	InvestorID      string
	Reserve1        string
	ExchangeID      string
	QuoteSysID      string
	InsertTimeStart string
	InsertTimeEnd   string
	InvestUnitID    string
	InstrumentID    string
}

func NewCThostFtdcQryQuoteField(p *C.CThostFtdcQryQuoteField) *CThostFtdcQryQuoteField {
	ret := new(CThostFtdcQryQuoteField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.QuoteSysID = c2goStr(&p.QuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.InsertTimeStart = c2goStr(&p.InsertTimeStart[0], C.sizeof_TThostFtdcTimeType)
	ret.InsertTimeEnd = c2goStr(&p.InsertTimeEnd[0], C.sizeof_TThostFtdcTimeType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryQuoteField) CValue() *C.CThostFtdcQryQuoteField {
	ptr := C.newCThostFtdcQryQuoteField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.QuoteSysID, &ptr.QuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.InsertTimeStart, &ptr.InsertTimeStart[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.InsertTimeEnd, &ptr.InsertTimeEnd[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcExchangeQuoteField struct {
	AskPrice          float64
	BidPrice          float64
	AskVolume         int
	BidVolume         int
	RequestID         int
	BusinessUnit      string
	AskOffsetFlag     byte
	BidOffsetFlag     byte
	AskHedgeFlag      byte
	BidHedgeFlag      byte
	QuoteLocalID      string
	ExchangeID        string
	ParticipantID     string
	ClientID          string
	Reserve1          string
	TraderID          string
	InstallID         int
	NotifySequence    int
	OrderSubmitStatus byte
	TradingDay        string
	SettlementID      int
	QuoteSysID        string
	InsertDate        string
	InsertTime        string
	CancelTime        string
	QuoteStatus       byte
	ClearingPartID    string
	SequenceNo        int
	AskOrderSysID     string
	BidOrderSysID     string
	ForQuoteSysID     string
	BranchID          string
	Reserve2          string
	MacAddress        string
	ExchangeInstID    string
	IPAddress         string
}

func NewCThostFtdcExchangeQuoteField(p *C.CThostFtdcExchangeQuoteField) *CThostFtdcExchangeQuoteField {
	ret := new(CThostFtdcExchangeQuoteField)
	ret.AskPrice = goFloat64(p.AskPrice)
	ret.BidPrice = goFloat64(p.BidPrice)
	ret.AskVolume = int(p.AskVolume)
	ret.BidVolume = int(p.BidVolume)
	ret.RequestID = int(p.RequestID)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.AskOffsetFlag = byte(p.AskOffsetFlag)
	ret.BidOffsetFlag = byte(p.BidOffsetFlag)
	ret.AskHedgeFlag = byte(p.AskHedgeFlag)
	ret.BidHedgeFlag = byte(p.BidHedgeFlag)
	ret.QuoteLocalID = c2goStr(&p.QuoteLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.NotifySequence = int(p.NotifySequence)
	ret.OrderSubmitStatus = byte(p.OrderSubmitStatus)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.QuoteSysID = c2goStr(&p.QuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.InsertDate = c2goStr(&p.InsertDate[0], C.sizeof_TThostFtdcDateType)
	ret.InsertTime = c2goStr(&p.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	ret.CancelTime = c2goStr(&p.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	ret.QuoteStatus = byte(p.QuoteStatus)
	ret.ClearingPartID = c2goStr(&p.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.SequenceNo = int(p.SequenceNo)
	ret.AskOrderSysID = c2goStr(&p.AskOrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.BidOrderSysID = c2goStr(&p.BidOrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ForQuoteSysID = c2goStr(&p.ForQuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcExchangeQuoteField) CValue() *C.CThostFtdcExchangeQuoteField {
	ptr := C.newCThostFtdcExchangeQuoteField()
	ptr.AskPrice = C.double(s.AskPrice)
	ptr.BidPrice = C.double(s.BidPrice)
	ptr.AskVolume = C.int(s.AskVolume)
	ptr.BidVolume = C.int(s.BidVolume)
	ptr.RequestID = C.int(s.RequestID)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.AskOffsetFlag = C.char(s.AskOffsetFlag)
	ptr.BidOffsetFlag = C.char(s.BidOffsetFlag)
	ptr.AskHedgeFlag = C.char(s.AskHedgeFlag)
	ptr.BidHedgeFlag = C.char(s.BidHedgeFlag)
	go2cStr(s.QuoteLocalID, &ptr.QuoteLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.NotifySequence = C.int(s.NotifySequence)
	ptr.OrderSubmitStatus = C.char(s.OrderSubmitStatus)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	go2cStr(s.QuoteSysID, &ptr.QuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.InsertDate, &ptr.InsertDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.InsertTime, &ptr.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.CancelTime, &ptr.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.QuoteStatus = C.char(s.QuoteStatus)
	go2cStr(s.ClearingPartID, &ptr.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ptr.SequenceNo = C.int(s.SequenceNo)
	go2cStr(s.AskOrderSysID, &ptr.AskOrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.BidOrderSysID, &ptr.BidOrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.ForQuoteSysID, &ptr.ForQuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryExchangeQuoteField struct {
	ParticipantID  string
	ClientID       string
	Reserve1       string
	ExchangeID     string
	TraderID       string
	ExchangeInstID string
}

func NewCThostFtdcQryExchangeQuoteField(p *C.CThostFtdcQryExchangeQuoteField) *CThostFtdcQryExchangeQuoteField {
	ret := new(CThostFtdcQryExchangeQuoteField)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ret
}
func (s *CThostFtdcQryExchangeQuoteField) CValue() *C.CThostFtdcQryExchangeQuoteField {
	ptr := C.newCThostFtdcQryExchangeQuoteField()
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ptr
}

type CThostFtdcQryQuoteActionField struct {
	BrokerID   string
	InvestorID string
	ExchangeID string
}

func NewCThostFtdcQryQuoteActionField(p *C.CThostFtdcQryQuoteActionField) *CThostFtdcQryQuoteActionField {
	ret := new(CThostFtdcQryQuoteActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ret
}
func (s *CThostFtdcQryQuoteActionField) CValue() *C.CThostFtdcQryQuoteActionField {
	ptr := C.newCThostFtdcQryQuoteActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ptr
}

type CThostFtdcExchangeQuoteActionField struct {
	ExchangeID        string
	QuoteSysID        string
	ActionFlag        byte
	ActionDate        string
	ActionTime        string
	TraderID          string
	InstallID         int
	QuoteLocalID      string
	ActionLocalID     string
	ParticipantID     string
	ClientID          string
	BusinessUnit      string
	OrderActionStatus byte
	UserID            string
	Reserve1          string
	MacAddress        string
	IPAddress         string
}

func NewCThostFtdcExchangeQuoteActionField(p *C.CThostFtdcExchangeQuoteActionField) *CThostFtdcExchangeQuoteActionField {
	ret := new(CThostFtdcExchangeQuoteActionField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.QuoteSysID = c2goStr(&p.QuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ActionFlag = byte(p.ActionFlag)
	ret.ActionDate = c2goStr(&p.ActionDate[0], C.sizeof_TThostFtdcDateType)
	ret.ActionTime = c2goStr(&p.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.QuoteLocalID = c2goStr(&p.QuoteLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ActionLocalID = c2goStr(&p.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.OrderActionStatus = byte(p.OrderActionStatus)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcExchangeQuoteActionField) CValue() *C.CThostFtdcExchangeQuoteActionField {
	ptr := C.newCThostFtdcExchangeQuoteActionField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.QuoteSysID, &ptr.QuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ptr.ActionFlag = C.char(s.ActionFlag)
	go2cStr(s.ActionDate, &ptr.ActionDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ActionTime, &ptr.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.QuoteLocalID, &ptr.QuoteLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ActionLocalID, &ptr.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.OrderActionStatus = C.char(s.OrderActionStatus)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryExchangeQuoteActionField struct {
	ParticipantID string
	ClientID      string
	ExchangeID    string
	TraderID      string
}

func NewCThostFtdcQryExchangeQuoteActionField(p *C.CThostFtdcQryExchangeQuoteActionField) *CThostFtdcQryExchangeQuoteActionField {
	ret := new(CThostFtdcQryExchangeQuoteActionField)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	return ret
}
func (s *CThostFtdcQryExchangeQuoteActionField) CValue() *C.CThostFtdcQryExchangeQuoteActionField {
	ptr := C.newCThostFtdcQryExchangeQuoteActionField()
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	return ptr
}

type CThostFtdcOptionInstrDeltaField struct {
	Reserve1      string
	InvestorRange byte
	BrokerID      string
	InvestorID    string
	Delta         float64
	InstrumentID  string
}

func NewCThostFtdcOptionInstrDeltaField(p *C.CThostFtdcOptionInstrDeltaField) *CThostFtdcOptionInstrDeltaField {
	ret := new(CThostFtdcOptionInstrDeltaField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Delta = goFloat64(p.Delta)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcOptionInstrDeltaField) CValue() *C.CThostFtdcOptionInstrDeltaField {
	ptr := C.newCThostFtdcOptionInstrDeltaField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.Delta = C.double(s.Delta)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcForQuoteRspField struct {
	TradingDay    string
	Reserve1      string
	ForQuoteSysID string
	ForQuoteTime  string
	ActionDay     string
	ExchangeID    string
	InstrumentID  string
}

func NewCThostFtdcForQuoteRspField(p *C.CThostFtdcForQuoteRspField) *CThostFtdcForQuoteRspField {
	ret := new(CThostFtdcForQuoteRspField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ForQuoteSysID = c2goStr(&p.ForQuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ForQuoteTime = c2goStr(&p.ForQuoteTime[0], C.sizeof_TThostFtdcTimeType)
	ret.ActionDay = c2goStr(&p.ActionDay[0], C.sizeof_TThostFtdcDateType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcForQuoteRspField) CValue() *C.CThostFtdcForQuoteRspField {
	ptr := C.newCThostFtdcForQuoteRspField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ForQuoteSysID, &ptr.ForQuoteSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.ForQuoteTime, &ptr.ForQuoteTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.ActionDay, &ptr.ActionDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcStrikeOffsetField struct {
	Reserve1      string
	InvestorRange byte
	BrokerID      string
	InvestorID    string
	Offset        float64
	OffsetType    byte
	InstrumentID  string
}

func NewCThostFtdcStrikeOffsetField(p *C.CThostFtdcStrikeOffsetField) *CThostFtdcStrikeOffsetField {
	ret := new(CThostFtdcStrikeOffsetField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Offset = goFloat64(p.Offset)
	ret.OffsetType = byte(p.OffsetType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcStrikeOffsetField) CValue() *C.CThostFtdcStrikeOffsetField {
	ptr := C.newCThostFtdcStrikeOffsetField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.Offset = C.double(s.Offset)
	ptr.OffsetType = C.char(s.OffsetType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryStrikeOffsetField struct {
	BrokerID     string
	InvestorID   string
	Reserve1     string
	InstrumentID string
}

func NewCThostFtdcQryStrikeOffsetField(p *C.CThostFtdcQryStrikeOffsetField) *CThostFtdcQryStrikeOffsetField {
	ret := new(CThostFtdcQryStrikeOffsetField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryStrikeOffsetField) CValue() *C.CThostFtdcQryStrikeOffsetField {
	ptr := C.newCThostFtdcQryStrikeOffsetField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcInputBatchOrderActionField struct {
	BrokerID       string
	InvestorID     string
	OrderActionRef int
	RequestID      int
	FrontID        int
	SessionID      int
	ExchangeID     string
	UserID         string
	InvestUnitID   string
	Reserve1       string
	MacAddress     string
	IPAddress      string
}

func NewCThostFtdcInputBatchOrderActionField(p *C.CThostFtdcInputBatchOrderActionField) *CThostFtdcInputBatchOrderActionField {
	ret := new(CThostFtdcInputBatchOrderActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.OrderActionRef = int(p.OrderActionRef)
	ret.RequestID = int(p.RequestID)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcInputBatchOrderActionField) CValue() *C.CThostFtdcInputBatchOrderActionField {
	ptr := C.newCThostFtdcInputBatchOrderActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.OrderActionRef = C.int(s.OrderActionRef)
	ptr.RequestID = C.int(s.RequestID)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcBatchOrderActionField struct {
	BrokerID          string
	InvestorID        string
	OrderActionRef    int
	RequestID         int
	FrontID           int
	SessionID         int
	ExchangeID        string
	ActionDate        string
	ActionTime        string
	TraderID          string
	InstallID         int
	ActionLocalID     string
	ParticipantID     string
	ClientID          string
	BusinessUnit      string
	OrderActionStatus byte
	UserID            string
	StatusMsg         string
	InvestUnitID      string
	Reserve1          string
	MacAddress        string
	IPAddress         string
}

func NewCThostFtdcBatchOrderActionField(p *C.CThostFtdcBatchOrderActionField) *CThostFtdcBatchOrderActionField {
	ret := new(CThostFtdcBatchOrderActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.OrderActionRef = int(p.OrderActionRef)
	ret.RequestID = int(p.RequestID)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ActionDate = c2goStr(&p.ActionDate[0], C.sizeof_TThostFtdcDateType)
	ret.ActionTime = c2goStr(&p.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.ActionLocalID = c2goStr(&p.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.OrderActionStatus = byte(p.OrderActionStatus)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.StatusMsg = c2goStr(&p.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcBatchOrderActionField) CValue() *C.CThostFtdcBatchOrderActionField {
	ptr := C.newCThostFtdcBatchOrderActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.OrderActionRef = C.int(s.OrderActionRef)
	ptr.RequestID = C.int(s.RequestID)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ActionDate, &ptr.ActionDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ActionTime, &ptr.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.ActionLocalID, &ptr.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.OrderActionStatus = C.char(s.OrderActionStatus)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.StatusMsg, &ptr.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcExchangeBatchOrderActionField struct {
	ExchangeID        string
	ActionDate        string
	ActionTime        string
	TraderID          string
	InstallID         int
	ActionLocalID     string
	ParticipantID     string
	ClientID          string
	BusinessUnit      string
	OrderActionStatus byte
	UserID            string
	Reserve1          string
	MacAddress        string
	IPAddress         string
}

func NewCThostFtdcExchangeBatchOrderActionField(p *C.CThostFtdcExchangeBatchOrderActionField) *CThostFtdcExchangeBatchOrderActionField {
	ret := new(CThostFtdcExchangeBatchOrderActionField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ActionDate = c2goStr(&p.ActionDate[0], C.sizeof_TThostFtdcDateType)
	ret.ActionTime = c2goStr(&p.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.ActionLocalID = c2goStr(&p.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.OrderActionStatus = byte(p.OrderActionStatus)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcExchangeBatchOrderActionField) CValue() *C.CThostFtdcExchangeBatchOrderActionField {
	ptr := C.newCThostFtdcExchangeBatchOrderActionField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ActionDate, &ptr.ActionDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ActionTime, &ptr.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.ActionLocalID, &ptr.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.OrderActionStatus = C.char(s.OrderActionStatus)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryBatchOrderActionField struct {
	BrokerID   string
	InvestorID string
	ExchangeID string
}

func NewCThostFtdcQryBatchOrderActionField(p *C.CThostFtdcQryBatchOrderActionField) *CThostFtdcQryBatchOrderActionField {
	ret := new(CThostFtdcQryBatchOrderActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ret
}
func (s *CThostFtdcQryBatchOrderActionField) CValue() *C.CThostFtdcQryBatchOrderActionField {
	ptr := C.newCThostFtdcQryBatchOrderActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ptr
}

type CThostFtdcCombInstrumentGuardField struct {
	BrokerID     string
	Reserve1     string
	GuarantRatio float64
	ExchangeID   string
	InstrumentID string
}

func NewCThostFtdcCombInstrumentGuardField(p *C.CThostFtdcCombInstrumentGuardField) *CThostFtdcCombInstrumentGuardField {
	ret := new(CThostFtdcCombInstrumentGuardField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.GuarantRatio = goFloat64(p.GuarantRatio)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcCombInstrumentGuardField) CValue() *C.CThostFtdcCombInstrumentGuardField {
	ptr := C.newCThostFtdcCombInstrumentGuardField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.GuarantRatio = C.double(s.GuarantRatio)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryCombInstrumentGuardField struct {
	BrokerID     string
	Reserve1     string
	ExchangeID   string
	InstrumentID string
}

func NewCThostFtdcQryCombInstrumentGuardField(p *C.CThostFtdcQryCombInstrumentGuardField) *CThostFtdcQryCombInstrumentGuardField {
	ret := new(CThostFtdcQryCombInstrumentGuardField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryCombInstrumentGuardField) CValue() *C.CThostFtdcQryCombInstrumentGuardField {
	ptr := C.newCThostFtdcQryCombInstrumentGuardField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcInputCombActionField struct {
	BrokerID      string
	InvestorID    string
	Reserve1      string
	CombActionRef string
	UserID        string
	Direction     byte
	Volume        int
	CombDirection byte
	HedgeFlag     byte
	ExchangeID    string
	Reserve2      string
	MacAddress    string
	InvestUnitID  string
	FrontID       int
	SessionID     int
	InstrumentID  string
	IPAddress     string
}

func NewCThostFtdcInputCombActionField(p *C.CThostFtdcInputCombActionField) *CThostFtdcInputCombActionField {
	ret := new(CThostFtdcInputCombActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.CombActionRef = c2goStr(&p.CombActionRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Direction = byte(p.Direction)
	ret.Volume = int(p.Volume)
	ret.CombDirection = byte(p.CombDirection)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcInputCombActionField) CValue() *C.CThostFtdcInputCombActionField {
	ptr := C.newCThostFtdcInputCombActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.CombActionRef, &ptr.CombActionRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.Direction = C.char(s.Direction)
	ptr.Volume = C.int(s.Volume)
	ptr.CombDirection = C.char(s.CombDirection)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcCombActionField struct {
	BrokerID        string
	InvestorID      string
	Reserve1        string
	CombActionRef   string
	UserID          string
	Direction       byte
	Volume          int
	CombDirection   byte
	HedgeFlag       byte
	ActionLocalID   string
	ExchangeID      string
	ParticipantID   string
	ClientID        string
	Reserve2        string
	TraderID        string
	InstallID       int
	ActionStatus    byte
	NotifySequence  int
	TradingDay      string
	SettlementID    int
	SequenceNo      int
	FrontID         int
	SessionID       int
	UserProductInfo string
	StatusMsg       string
	Reserve3        string
	MacAddress      string
	ComTradeID      string
	BranchID        string
	InvestUnitID    string
	InstrumentID    string
	ExchangeInstID  string
	IPAddress       string
}

func NewCThostFtdcCombActionField(p *C.CThostFtdcCombActionField) *CThostFtdcCombActionField {
	ret := new(CThostFtdcCombActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.CombActionRef = c2goStr(&p.CombActionRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Direction = byte(p.Direction)
	ret.Volume = int(p.Volume)
	ret.CombDirection = byte(p.CombDirection)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.ActionLocalID = c2goStr(&p.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.ActionStatus = byte(p.ActionStatus)
	ret.NotifySequence = int(p.NotifySequence)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.SequenceNo = int(p.SequenceNo)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.UserProductInfo = c2goStr(&p.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.StatusMsg = c2goStr(&p.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.Reserve3 = c2goStr(&p.reserve3[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.ComTradeID = c2goStr(&p.ComTradeID[0], C.sizeof_TThostFtdcTradeIDType)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcCombActionField) CValue() *C.CThostFtdcCombActionField {
	ptr := C.newCThostFtdcCombActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.CombActionRef, &ptr.CombActionRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.Direction = C.char(s.Direction)
	ptr.Volume = C.int(s.Volume)
	ptr.CombDirection = C.char(s.CombDirection)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	go2cStr(s.ActionLocalID, &ptr.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.ActionStatus = C.char(s.ActionStatus)
	ptr.NotifySequence = C.int(s.NotifySequence)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	ptr.SequenceNo = C.int(s.SequenceNo)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.UserProductInfo, &ptr.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.StatusMsg, &ptr.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.Reserve3, &ptr.reserve3[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.ComTradeID, &ptr.ComTradeID[0], C.sizeof_TThostFtdcTradeIDType)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryCombActionField struct {
	BrokerID     string
	InvestorID   string
	Reserve1     string
	ExchangeID   string
	InvestUnitID string
	InstrumentID string
}

func NewCThostFtdcQryCombActionField(p *C.CThostFtdcQryCombActionField) *CThostFtdcQryCombActionField {
	ret := new(CThostFtdcQryCombActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryCombActionField) CValue() *C.CThostFtdcQryCombActionField {
	ptr := C.newCThostFtdcQryCombActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcExchangeCombActionField struct {
	Direction      byte
	Volume         int
	CombDirection  byte
	HedgeFlag      byte
	ActionLocalID  string
	ExchangeID     string
	ParticipantID  string
	ClientID       string
	Reserve1       string
	TraderID       string
	InstallID      int
	ActionStatus   byte
	NotifySequence int
	TradingDay     string
	SettlementID   int
	SequenceNo     int
	Reserve2       string
	MacAddress     string
	ComTradeID     string
	BranchID       string
	ExchangeInstID string
	IPAddress      string
}

func NewCThostFtdcExchangeCombActionField(p *C.CThostFtdcExchangeCombActionField) *CThostFtdcExchangeCombActionField {
	ret := new(CThostFtdcExchangeCombActionField)
	ret.Direction = byte(p.Direction)
	ret.Volume = int(p.Volume)
	ret.CombDirection = byte(p.CombDirection)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.ActionLocalID = c2goStr(&p.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.ActionStatus = byte(p.ActionStatus)
	ret.NotifySequence = int(p.NotifySequence)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.SequenceNo = int(p.SequenceNo)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.ComTradeID = c2goStr(&p.ComTradeID[0], C.sizeof_TThostFtdcTradeIDType)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcExchangeCombActionField) CValue() *C.CThostFtdcExchangeCombActionField {
	ptr := C.newCThostFtdcExchangeCombActionField()
	ptr.Direction = C.char(s.Direction)
	ptr.Volume = C.int(s.Volume)
	ptr.CombDirection = C.char(s.CombDirection)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	go2cStr(s.ActionLocalID, &ptr.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.ActionStatus = C.char(s.ActionStatus)
	ptr.NotifySequence = C.int(s.NotifySequence)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	ptr.SequenceNo = C.int(s.SequenceNo)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.ComTradeID, &ptr.ComTradeID[0], C.sizeof_TThostFtdcTradeIDType)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryExchangeCombActionField struct {
	ParticipantID  string
	ClientID       string
	Reserve1       string
	ExchangeID     string
	TraderID       string
	ExchangeInstID string
}

func NewCThostFtdcQryExchangeCombActionField(p *C.CThostFtdcQryExchangeCombActionField) *CThostFtdcQryExchangeCombActionField {
	ret := new(CThostFtdcQryExchangeCombActionField)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ret
}
func (s *CThostFtdcQryExchangeCombActionField) CValue() *C.CThostFtdcQryExchangeCombActionField {
	ptr := C.newCThostFtdcQryExchangeCombActionField()
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ptr
}

type CThostFtdcProductExchRateField struct {
	Reserve1        string
	QuoteCurrencyID string
	ExchangeRate    float64
	ExchangeID      string
	ProductID       string
}

func NewCThostFtdcProductExchRateField(p *C.CThostFtdcProductExchRateField) *CThostFtdcProductExchRateField {
	ret := new(CThostFtdcProductExchRateField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.QuoteCurrencyID = c2goStr(&p.QuoteCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.ExchangeRate = goFloat64(p.ExchangeRate)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ProductID = c2goStr(&p.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcProductExchRateField) CValue() *C.CThostFtdcProductExchRateField {
	ptr := C.newCThostFtdcProductExchRateField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.QuoteCurrencyID, &ptr.QuoteCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.ExchangeRate = C.double(s.ExchangeRate)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ProductID, &ptr.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryProductExchRateField struct {
	Reserve1   string
	ExchangeID string
	ProductID  string
}

func NewCThostFtdcQryProductExchRateField(p *C.CThostFtdcQryProductExchRateField) *CThostFtdcQryProductExchRateField {
	ret := new(CThostFtdcQryProductExchRateField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ProductID = c2goStr(&p.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryProductExchRateField) CValue() *C.CThostFtdcQryProductExchRateField {
	ptr := C.newCThostFtdcQryProductExchRateField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ProductID, &ptr.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryForQuoteParamField struct {
	BrokerID     string
	Reserve1     string
	ExchangeID   string
	InstrumentID string
}

func NewCThostFtdcQryForQuoteParamField(p *C.CThostFtdcQryForQuoteParamField) *CThostFtdcQryForQuoteParamField {
	ret := new(CThostFtdcQryForQuoteParamField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryForQuoteParamField) CValue() *C.CThostFtdcQryForQuoteParamField {
	ptr := C.newCThostFtdcQryForQuoteParamField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcForQuoteParamField struct {
	BrokerID      string
	Reserve1      string
	ExchangeID    string
	LastPrice     float64
	PriceInterval float64
	InstrumentID  string
}

func NewCThostFtdcForQuoteParamField(p *C.CThostFtdcForQuoteParamField) *CThostFtdcForQuoteParamField {
	ret := new(CThostFtdcForQuoteParamField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.LastPrice = goFloat64(p.LastPrice)
	ret.PriceInterval = goFloat64(p.PriceInterval)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcForQuoteParamField) CValue() *C.CThostFtdcForQuoteParamField {
	ptr := C.newCThostFtdcForQuoteParamField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ptr.LastPrice = C.double(s.LastPrice)
	ptr.PriceInterval = C.double(s.PriceInterval)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcMMOptionInstrCommRateField struct {
	Reserve1                string
	InvestorRange           byte
	BrokerID                string
	InvestorID              string
	OpenRatioByMoney        float64
	OpenRatioByVolume       float64
	CloseRatioByMoney       float64
	CloseRatioByVolume      float64
	CloseTodayRatioByMoney  float64
	CloseTodayRatioByVolume float64
	StrikeRatioByMoney      float64
	StrikeRatioByVolume     float64
	InstrumentID            string
}

func NewCThostFtdcMMOptionInstrCommRateField(p *C.CThostFtdcMMOptionInstrCommRateField) *CThostFtdcMMOptionInstrCommRateField {
	ret := new(CThostFtdcMMOptionInstrCommRateField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.OpenRatioByMoney = goFloat64(p.OpenRatioByMoney)
	ret.OpenRatioByVolume = goFloat64(p.OpenRatioByVolume)
	ret.CloseRatioByMoney = goFloat64(p.CloseRatioByMoney)
	ret.CloseRatioByVolume = goFloat64(p.CloseRatioByVolume)
	ret.CloseTodayRatioByMoney = goFloat64(p.CloseTodayRatioByMoney)
	ret.CloseTodayRatioByVolume = goFloat64(p.CloseTodayRatioByVolume)
	ret.StrikeRatioByMoney = goFloat64(p.StrikeRatioByMoney)
	ret.StrikeRatioByVolume = goFloat64(p.StrikeRatioByVolume)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcMMOptionInstrCommRateField) CValue() *C.CThostFtdcMMOptionInstrCommRateField {
	ptr := C.newCThostFtdcMMOptionInstrCommRateField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.OpenRatioByMoney = C.double(s.OpenRatioByMoney)
	ptr.OpenRatioByVolume = C.double(s.OpenRatioByVolume)
	ptr.CloseRatioByMoney = C.double(s.CloseRatioByMoney)
	ptr.CloseRatioByVolume = C.double(s.CloseRatioByVolume)
	ptr.CloseTodayRatioByMoney = C.double(s.CloseTodayRatioByMoney)
	ptr.CloseTodayRatioByVolume = C.double(s.CloseTodayRatioByVolume)
	ptr.StrikeRatioByMoney = C.double(s.StrikeRatioByMoney)
	ptr.StrikeRatioByVolume = C.double(s.StrikeRatioByVolume)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryMMOptionInstrCommRateField struct {
	BrokerID     string
	InvestorID   string
	Reserve1     string
	InstrumentID string
}

func NewCThostFtdcQryMMOptionInstrCommRateField(p *C.CThostFtdcQryMMOptionInstrCommRateField) *CThostFtdcQryMMOptionInstrCommRateField {
	ret := new(CThostFtdcQryMMOptionInstrCommRateField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryMMOptionInstrCommRateField) CValue() *C.CThostFtdcQryMMOptionInstrCommRateField {
	ptr := C.newCThostFtdcQryMMOptionInstrCommRateField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcMMInstrumentCommissionRateField struct {
	Reserve1                string
	InvestorRange           byte
	BrokerID                string
	InvestorID              string
	OpenRatioByMoney        float64
	OpenRatioByVolume       float64
	CloseRatioByMoney       float64
	CloseRatioByVolume      float64
	CloseTodayRatioByMoney  float64
	CloseTodayRatioByVolume float64
	InstrumentID            string
}

func NewCThostFtdcMMInstrumentCommissionRateField(p *C.CThostFtdcMMInstrumentCommissionRateField) *CThostFtdcMMInstrumentCommissionRateField {
	ret := new(CThostFtdcMMInstrumentCommissionRateField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.OpenRatioByMoney = goFloat64(p.OpenRatioByMoney)
	ret.OpenRatioByVolume = goFloat64(p.OpenRatioByVolume)
	ret.CloseRatioByMoney = goFloat64(p.CloseRatioByMoney)
	ret.CloseRatioByVolume = goFloat64(p.CloseRatioByVolume)
	ret.CloseTodayRatioByMoney = goFloat64(p.CloseTodayRatioByMoney)
	ret.CloseTodayRatioByVolume = goFloat64(p.CloseTodayRatioByVolume)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcMMInstrumentCommissionRateField) CValue() *C.CThostFtdcMMInstrumentCommissionRateField {
	ptr := C.newCThostFtdcMMInstrumentCommissionRateField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.OpenRatioByMoney = C.double(s.OpenRatioByMoney)
	ptr.OpenRatioByVolume = C.double(s.OpenRatioByVolume)
	ptr.CloseRatioByMoney = C.double(s.CloseRatioByMoney)
	ptr.CloseRatioByVolume = C.double(s.CloseRatioByVolume)
	ptr.CloseTodayRatioByMoney = C.double(s.CloseTodayRatioByMoney)
	ptr.CloseTodayRatioByVolume = C.double(s.CloseTodayRatioByVolume)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryMMInstrumentCommissionRateField struct {
	BrokerID     string
	InvestorID   string
	Reserve1     string
	InstrumentID string
}

func NewCThostFtdcQryMMInstrumentCommissionRateField(p *C.CThostFtdcQryMMInstrumentCommissionRateField) *CThostFtdcQryMMInstrumentCommissionRateField {
	ret := new(CThostFtdcQryMMInstrumentCommissionRateField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryMMInstrumentCommissionRateField) CValue() *C.CThostFtdcQryMMInstrumentCommissionRateField {
	ptr := C.newCThostFtdcQryMMInstrumentCommissionRateField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcInstrumentOrderCommRateField struct {
	Reserve1                string
	InvestorRange           byte
	BrokerID                string
	InvestorID              string
	HedgeFlag               byte
	OrderCommByVolume       float64
	OrderActionCommByVolume float64
	ExchangeID              string
	InvestUnitID            string
	InstrumentID            string
	OrderCommByTrade        float64
	OrderActionCommByTrade  float64
}

func NewCThostFtdcInstrumentOrderCommRateField(p *C.CThostFtdcInstrumentOrderCommRateField) *CThostFtdcInstrumentOrderCommRateField {
	ret := new(CThostFtdcInstrumentOrderCommRateField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.OrderCommByVolume = goFloat64(p.OrderCommByVolume)
	ret.OrderActionCommByVolume = goFloat64(p.OrderActionCommByVolume)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.OrderCommByTrade = goFloat64(p.OrderCommByTrade)
	ret.OrderActionCommByTrade = goFloat64(p.OrderActionCommByTrade)
	return ret
}
func (s *CThostFtdcInstrumentOrderCommRateField) CValue() *C.CThostFtdcInstrumentOrderCommRateField {
	ptr := C.newCThostFtdcInstrumentOrderCommRateField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.OrderCommByVolume = C.double(s.OrderCommByVolume)
	ptr.OrderActionCommByVolume = C.double(s.OrderActionCommByVolume)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ptr.OrderCommByTrade = C.double(s.OrderCommByTrade)
	ptr.OrderActionCommByTrade = C.double(s.OrderActionCommByTrade)
	return ptr
}

type CThostFtdcQryInstrumentOrderCommRateField struct {
	BrokerID     string
	InvestorID   string
	Reserve1     string
	InstrumentID string
}

func NewCThostFtdcQryInstrumentOrderCommRateField(p *C.CThostFtdcQryInstrumentOrderCommRateField) *CThostFtdcQryInstrumentOrderCommRateField {
	ret := new(CThostFtdcQryInstrumentOrderCommRateField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryInstrumentOrderCommRateField) CValue() *C.CThostFtdcQryInstrumentOrderCommRateField {
	ptr := C.newCThostFtdcQryInstrumentOrderCommRateField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcTradeParamField struct {
	BrokerID        string
	TradeParamID    byte
	TradeParamValue string
	Memo            string
}

func NewCThostFtdcTradeParamField(p *C.CThostFtdcTradeParamField) *CThostFtdcTradeParamField {
	ret := new(CThostFtdcTradeParamField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.TradeParamID = byte(p.TradeParamID)
	ret.TradeParamValue = c2goStr(&p.TradeParamValue[0], C.sizeof_TThostFtdcSettlementParamValueType)
	ret.Memo = c2goStr(&p.Memo[0], C.sizeof_TThostFtdcMemoType)
	return ret
}
func (s *CThostFtdcTradeParamField) CValue() *C.CThostFtdcTradeParamField {
	ptr := C.newCThostFtdcTradeParamField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ptr.TradeParamID = C.char(s.TradeParamID)
	go2cStr(s.TradeParamValue, &ptr.TradeParamValue[0], C.sizeof_TThostFtdcSettlementParamValueType)
	go2cStr(s.Memo, &ptr.Memo[0], C.sizeof_TThostFtdcMemoType)
	return ptr
}

type CThostFtdcInstrumentMarginRateULField struct {
	Reserve1                 string
	InvestorRange            byte
	BrokerID                 string
	InvestorID               string
	HedgeFlag                byte
	LongMarginRatioByMoney   float64
	LongMarginRatioByVolume  float64
	ShortMarginRatioByMoney  float64
	ShortMarginRatioByVolume float64
	InstrumentID             string
}

func NewCThostFtdcInstrumentMarginRateULField(p *C.CThostFtdcInstrumentMarginRateULField) *CThostFtdcInstrumentMarginRateULField {
	ret := new(CThostFtdcInstrumentMarginRateULField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.LongMarginRatioByMoney = goFloat64(p.LongMarginRatioByMoney)
	ret.LongMarginRatioByVolume = goFloat64(p.LongMarginRatioByVolume)
	ret.ShortMarginRatioByMoney = goFloat64(p.ShortMarginRatioByMoney)
	ret.ShortMarginRatioByVolume = goFloat64(p.ShortMarginRatioByVolume)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcInstrumentMarginRateULField) CValue() *C.CThostFtdcInstrumentMarginRateULField {
	ptr := C.newCThostFtdcInstrumentMarginRateULField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.LongMarginRatioByMoney = C.double(s.LongMarginRatioByMoney)
	ptr.LongMarginRatioByVolume = C.double(s.LongMarginRatioByVolume)
	ptr.ShortMarginRatioByMoney = C.double(s.ShortMarginRatioByMoney)
	ptr.ShortMarginRatioByVolume = C.double(s.ShortMarginRatioByVolume)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcFutureLimitPosiParamField struct {
	InvestorRange  byte
	BrokerID       string
	InvestorID     string
	Reserve1       string
	SpecOpenVolume int
	ArbiOpenVolume int
	OpenVolume     int
	ProductID      string
}

func NewCThostFtdcFutureLimitPosiParamField(p *C.CThostFtdcFutureLimitPosiParamField) *CThostFtdcFutureLimitPosiParamField {
	ret := new(CThostFtdcFutureLimitPosiParamField)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.SpecOpenVolume = int(p.SpecOpenVolume)
	ret.ArbiOpenVolume = int(p.ArbiOpenVolume)
	ret.OpenVolume = int(p.OpenVolume)
	ret.ProductID = c2goStr(&p.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcFutureLimitPosiParamField) CValue() *C.CThostFtdcFutureLimitPosiParamField {
	ptr := C.newCThostFtdcFutureLimitPosiParamField()
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.SpecOpenVolume = C.int(s.SpecOpenVolume)
	ptr.ArbiOpenVolume = C.int(s.ArbiOpenVolume)
	ptr.OpenVolume = C.int(s.OpenVolume)
	go2cStr(s.ProductID, &ptr.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcLoginForbiddenIPField struct {
	Reserve1  string
	IPAddress string
}

func NewCThostFtdcLoginForbiddenIPField(p *C.CThostFtdcLoginForbiddenIPField) *CThostFtdcLoginForbiddenIPField {
	ret := new(CThostFtdcLoginForbiddenIPField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcLoginForbiddenIPField) CValue() *C.CThostFtdcLoginForbiddenIPField {
	ptr := C.newCThostFtdcLoginForbiddenIPField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcIPListField struct {
	Reserve1  string
	IsWhite   int
	IPAddress string
}

func NewCThostFtdcIPListField(p *C.CThostFtdcIPListField) *CThostFtdcIPListField {
	ret := new(CThostFtdcIPListField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.IsWhite = int(p.IsWhite)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcIPListField) CValue() *C.CThostFtdcIPListField {
	ptr := C.newCThostFtdcIPListField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ptr.IsWhite = C.int(s.IsWhite)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcInputOptionSelfCloseField struct {
	BrokerID           string
	InvestorID         string
	Reserve1           string
	OptionSelfCloseRef string
	UserID             string
	Volume             int
	RequestID          int
	BusinessUnit       string
	HedgeFlag          byte
	OptSelfCloseFlag   byte
	ExchangeID         string
	InvestUnitID       string
	AccountID          string
	CurrencyID         string
	ClientID           string
	Reserve2           string
	MacAddress         string
	InstrumentID       string
	IPAddress          string
}

func NewCThostFtdcInputOptionSelfCloseField(p *C.CThostFtdcInputOptionSelfCloseField) *CThostFtdcInputOptionSelfCloseField {
	ret := new(CThostFtdcInputOptionSelfCloseField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.OptionSelfCloseRef = c2goStr(&p.OptionSelfCloseRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Volume = int(p.Volume)
	ret.RequestID = int(p.RequestID)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.OptSelfCloseFlag = byte(p.OptSelfCloseFlag)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcInputOptionSelfCloseField) CValue() *C.CThostFtdcInputOptionSelfCloseField {
	ptr := C.newCThostFtdcInputOptionSelfCloseField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.OptionSelfCloseRef, &ptr.OptionSelfCloseRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.Volume = C.int(s.Volume)
	ptr.RequestID = C.int(s.RequestID)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.OptSelfCloseFlag = C.char(s.OptSelfCloseFlag)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcInputOptionSelfCloseActionField struct {
	BrokerID                 string
	InvestorID               string
	OptionSelfCloseActionRef int
	OptionSelfCloseRef       string
	RequestID                int
	FrontID                  int
	SessionID                int
	ExchangeID               string
	OptionSelfCloseSysID     string
	ActionFlag               byte
	UserID                   string
	Reserve1                 string
	InvestUnitID             string
	Reserve2                 string
	MacAddress               string
	InstrumentID             string
	IPAddress                string
}

func NewCThostFtdcInputOptionSelfCloseActionField(p *C.CThostFtdcInputOptionSelfCloseActionField) *CThostFtdcInputOptionSelfCloseActionField {
	ret := new(CThostFtdcInputOptionSelfCloseActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.OptionSelfCloseActionRef = int(p.OptionSelfCloseActionRef)
	ret.OptionSelfCloseRef = c2goStr(&p.OptionSelfCloseRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.RequestID = int(p.RequestID)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.OptionSelfCloseSysID = c2goStr(&p.OptionSelfCloseSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ActionFlag = byte(p.ActionFlag)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcInputOptionSelfCloseActionField) CValue() *C.CThostFtdcInputOptionSelfCloseActionField {
	ptr := C.newCThostFtdcInputOptionSelfCloseActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.OptionSelfCloseActionRef = C.int(s.OptionSelfCloseActionRef)
	go2cStr(s.OptionSelfCloseRef, &ptr.OptionSelfCloseRef[0], C.sizeof_TThostFtdcOrderRefType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.OptionSelfCloseSysID, &ptr.OptionSelfCloseSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ptr.ActionFlag = C.char(s.ActionFlag)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcOptionSelfCloseField struct {
	BrokerID                 string
	InvestorID               string
	Reserve1                 string
	OptionSelfCloseRef       string
	UserID                   string
	Volume                   int
	RequestID                int
	BusinessUnit             string
	HedgeFlag                byte
	OptSelfCloseFlag         byte
	OptionSelfCloseLocalID   string
	ExchangeID               string
	ParticipantID            string
	ClientID                 string
	Reserve2                 string
	TraderID                 string
	InstallID                int
	OrderSubmitStatus        byte
	NotifySequence           int
	TradingDay               string
	SettlementID             int
	OptionSelfCloseSysID     string
	InsertDate               string
	InsertTime               string
	CancelTime               string
	ExecResult               byte
	ClearingPartID           string
	SequenceNo               int
	FrontID                  int
	SessionID                int
	UserProductInfo          string
	StatusMsg                string
	ActiveUserID             string
	BrokerOptionSelfCloseSeq int
	BranchID                 string
	InvestUnitID             string
	AccountID                string
	CurrencyID               string
	Reserve3                 string
	MacAddress               string
	InstrumentID             string
	ExchangeInstID           string
	IPAddress                string
}

func NewCThostFtdcOptionSelfCloseField(p *C.CThostFtdcOptionSelfCloseField) *CThostFtdcOptionSelfCloseField {
	ret := new(CThostFtdcOptionSelfCloseField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.OptionSelfCloseRef = c2goStr(&p.OptionSelfCloseRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Volume = int(p.Volume)
	ret.RequestID = int(p.RequestID)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.OptSelfCloseFlag = byte(p.OptSelfCloseFlag)
	ret.OptionSelfCloseLocalID = c2goStr(&p.OptionSelfCloseLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.OrderSubmitStatus = byte(p.OrderSubmitStatus)
	ret.NotifySequence = int(p.NotifySequence)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.OptionSelfCloseSysID = c2goStr(&p.OptionSelfCloseSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.InsertDate = c2goStr(&p.InsertDate[0], C.sizeof_TThostFtdcDateType)
	ret.InsertTime = c2goStr(&p.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	ret.CancelTime = c2goStr(&p.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	ret.ExecResult = byte(p.ExecResult)
	ret.ClearingPartID = c2goStr(&p.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.SequenceNo = int(p.SequenceNo)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.UserProductInfo = c2goStr(&p.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.StatusMsg = c2goStr(&p.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.ActiveUserID = c2goStr(&p.ActiveUserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.BrokerOptionSelfCloseSeq = int(p.BrokerOptionSelfCloseSeq)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.Reserve3 = c2goStr(&p.reserve3[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcOptionSelfCloseField) CValue() *C.CThostFtdcOptionSelfCloseField {
	ptr := C.newCThostFtdcOptionSelfCloseField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.OptionSelfCloseRef, &ptr.OptionSelfCloseRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.Volume = C.int(s.Volume)
	ptr.RequestID = C.int(s.RequestID)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.OptSelfCloseFlag = C.char(s.OptSelfCloseFlag)
	go2cStr(s.OptionSelfCloseLocalID, &ptr.OptionSelfCloseLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.OrderSubmitStatus = C.char(s.OrderSubmitStatus)
	ptr.NotifySequence = C.int(s.NotifySequence)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	go2cStr(s.OptionSelfCloseSysID, &ptr.OptionSelfCloseSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.InsertDate, &ptr.InsertDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.InsertTime, &ptr.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.CancelTime, &ptr.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.ExecResult = C.char(s.ExecResult)
	go2cStr(s.ClearingPartID, &ptr.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ptr.SequenceNo = C.int(s.SequenceNo)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.UserProductInfo, &ptr.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.StatusMsg, &ptr.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.ActiveUserID, &ptr.ActiveUserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.BrokerOptionSelfCloseSeq = C.int(s.BrokerOptionSelfCloseSeq)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.Reserve3, &ptr.reserve3[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcOptionSelfCloseActionField struct {
	BrokerID                 string
	InvestorID               string
	OptionSelfCloseActionRef int
	OptionSelfCloseRef       string
	RequestID                int
	FrontID                  int
	SessionID                int
	ExchangeID               string
	OptionSelfCloseSysID     string
	ActionFlag               byte
	ActionDate               string
	ActionTime               string
	TraderID                 string
	InstallID                int
	OptionSelfCloseLocalID   string
	ActionLocalID            string
	ParticipantID            string
	ClientID                 string
	BusinessUnit             string
	OrderActionStatus        byte
	UserID                   string
	StatusMsg                string
	Reserve1                 string
	BranchID                 string
	InvestUnitID             string
	Reserve2                 string
	MacAddress               string
	InstrumentID             string
	IPAddress                string
}

func NewCThostFtdcOptionSelfCloseActionField(p *C.CThostFtdcOptionSelfCloseActionField) *CThostFtdcOptionSelfCloseActionField {
	ret := new(CThostFtdcOptionSelfCloseActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.OptionSelfCloseActionRef = int(p.OptionSelfCloseActionRef)
	ret.OptionSelfCloseRef = c2goStr(&p.OptionSelfCloseRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.RequestID = int(p.RequestID)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.OptionSelfCloseSysID = c2goStr(&p.OptionSelfCloseSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ActionFlag = byte(p.ActionFlag)
	ret.ActionDate = c2goStr(&p.ActionDate[0], C.sizeof_TThostFtdcDateType)
	ret.ActionTime = c2goStr(&p.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.OptionSelfCloseLocalID = c2goStr(&p.OptionSelfCloseLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ActionLocalID = c2goStr(&p.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.OrderActionStatus = byte(p.OrderActionStatus)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.StatusMsg = c2goStr(&p.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcOptionSelfCloseActionField) CValue() *C.CThostFtdcOptionSelfCloseActionField {
	ptr := C.newCThostFtdcOptionSelfCloseActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.OptionSelfCloseActionRef = C.int(s.OptionSelfCloseActionRef)
	go2cStr(s.OptionSelfCloseRef, &ptr.OptionSelfCloseRef[0], C.sizeof_TThostFtdcOrderRefType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.OptionSelfCloseSysID, &ptr.OptionSelfCloseSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ptr.ActionFlag = C.char(s.ActionFlag)
	go2cStr(s.ActionDate, &ptr.ActionDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ActionTime, &ptr.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.OptionSelfCloseLocalID, &ptr.OptionSelfCloseLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ActionLocalID, &ptr.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.OrderActionStatus = C.char(s.OrderActionStatus)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.StatusMsg, &ptr.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryOptionSelfCloseField struct {
	BrokerID             string
	InvestorID           string
	Reserve1             string
	ExchangeID           string
	OptionSelfCloseSysID string
	InsertTimeStart      string
	InsertTimeEnd        string
	InstrumentID         string
}

func NewCThostFtdcQryOptionSelfCloseField(p *C.CThostFtdcQryOptionSelfCloseField) *CThostFtdcQryOptionSelfCloseField {
	ret := new(CThostFtdcQryOptionSelfCloseField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.OptionSelfCloseSysID = c2goStr(&p.OptionSelfCloseSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.InsertTimeStart = c2goStr(&p.InsertTimeStart[0], C.sizeof_TThostFtdcTimeType)
	ret.InsertTimeEnd = c2goStr(&p.InsertTimeEnd[0], C.sizeof_TThostFtdcTimeType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryOptionSelfCloseField) CValue() *C.CThostFtdcQryOptionSelfCloseField {
	ptr := C.newCThostFtdcQryOptionSelfCloseField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.OptionSelfCloseSysID, &ptr.OptionSelfCloseSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.InsertTimeStart, &ptr.InsertTimeStart[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.InsertTimeEnd, &ptr.InsertTimeEnd[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcExchangeOptionSelfCloseField struct {
	Volume                 int
	RequestID              int
	BusinessUnit           string
	HedgeFlag              byte
	OptSelfCloseFlag       byte
	OptionSelfCloseLocalID string
	ExchangeID             string
	ParticipantID          string
	ClientID               string
	Reserve1               string
	TraderID               string
	InstallID              int
	OrderSubmitStatus      byte
	NotifySequence         int
	TradingDay             string
	SettlementID           int
	OptionSelfCloseSysID   string
	InsertDate             string
	InsertTime             string
	CancelTime             string
	ExecResult             byte
	ClearingPartID         string
	SequenceNo             int
	BranchID               string
	Reserve2               string
	MacAddress             string
	ExchangeInstID         string
	IPAddress              string
}

func NewCThostFtdcExchangeOptionSelfCloseField(p *C.CThostFtdcExchangeOptionSelfCloseField) *CThostFtdcExchangeOptionSelfCloseField {
	ret := new(CThostFtdcExchangeOptionSelfCloseField)
	ret.Volume = int(p.Volume)
	ret.RequestID = int(p.RequestID)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.OptSelfCloseFlag = byte(p.OptSelfCloseFlag)
	ret.OptionSelfCloseLocalID = c2goStr(&p.OptionSelfCloseLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.OrderSubmitStatus = byte(p.OrderSubmitStatus)
	ret.NotifySequence = int(p.NotifySequence)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.OptionSelfCloseSysID = c2goStr(&p.OptionSelfCloseSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.InsertDate = c2goStr(&p.InsertDate[0], C.sizeof_TThostFtdcDateType)
	ret.InsertTime = c2goStr(&p.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	ret.CancelTime = c2goStr(&p.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	ret.ExecResult = byte(p.ExecResult)
	ret.ClearingPartID = c2goStr(&p.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.SequenceNo = int(p.SequenceNo)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcExchangeOptionSelfCloseField) CValue() *C.CThostFtdcExchangeOptionSelfCloseField {
	ptr := C.newCThostFtdcExchangeOptionSelfCloseField()
	ptr.Volume = C.int(s.Volume)
	ptr.RequestID = C.int(s.RequestID)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.OptSelfCloseFlag = C.char(s.OptSelfCloseFlag)
	go2cStr(s.OptionSelfCloseLocalID, &ptr.OptionSelfCloseLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.OrderSubmitStatus = C.char(s.OrderSubmitStatus)
	ptr.NotifySequence = C.int(s.NotifySequence)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	go2cStr(s.OptionSelfCloseSysID, &ptr.OptionSelfCloseSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	go2cStr(s.InsertDate, &ptr.InsertDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.InsertTime, &ptr.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.CancelTime, &ptr.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.ExecResult = C.char(s.ExecResult)
	go2cStr(s.ClearingPartID, &ptr.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ptr.SequenceNo = C.int(s.SequenceNo)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryOptionSelfCloseActionField struct {
	BrokerID   string
	InvestorID string
	ExchangeID string
}

func NewCThostFtdcQryOptionSelfCloseActionField(p *C.CThostFtdcQryOptionSelfCloseActionField) *CThostFtdcQryOptionSelfCloseActionField {
	ret := new(CThostFtdcQryOptionSelfCloseActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ret
}
func (s *CThostFtdcQryOptionSelfCloseActionField) CValue() *C.CThostFtdcQryOptionSelfCloseActionField {
	ptr := C.newCThostFtdcQryOptionSelfCloseActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ptr
}

type CThostFtdcExchangeOptionSelfCloseActionField struct {
	ExchangeID             string
	OptionSelfCloseSysID   string
	ActionFlag             byte
	ActionDate             string
	ActionTime             string
	TraderID               string
	InstallID              int
	OptionSelfCloseLocalID string
	ActionLocalID          string
	ParticipantID          string
	ClientID               string
	BusinessUnit           string
	OrderActionStatus      byte
	UserID                 string
	BranchID               string
	Reserve1               string
	MacAddress             string
	Reserve2               string
	OptSelfCloseFlag       byte
	IPAddress              string
	ExchangeInstID         string
}

func NewCThostFtdcExchangeOptionSelfCloseActionField(p *C.CThostFtdcExchangeOptionSelfCloseActionField) *CThostFtdcExchangeOptionSelfCloseActionField {
	ret := new(CThostFtdcExchangeOptionSelfCloseActionField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.OptionSelfCloseSysID = c2goStr(&p.OptionSelfCloseSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ActionFlag = byte(p.ActionFlag)
	ret.ActionDate = c2goStr(&p.ActionDate[0], C.sizeof_TThostFtdcDateType)
	ret.ActionTime = c2goStr(&p.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.OptionSelfCloseLocalID = c2goStr(&p.OptionSelfCloseLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ActionLocalID = c2goStr(&p.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.OrderActionStatus = byte(p.OrderActionStatus)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.OptSelfCloseFlag = byte(p.OptSelfCloseFlag)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ret
}
func (s *CThostFtdcExchangeOptionSelfCloseActionField) CValue() *C.CThostFtdcExchangeOptionSelfCloseActionField {
	ptr := C.newCThostFtdcExchangeOptionSelfCloseActionField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.OptionSelfCloseSysID, &ptr.OptionSelfCloseSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ptr.ActionFlag = C.char(s.ActionFlag)
	go2cStr(s.ActionDate, &ptr.ActionDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ActionTime, &ptr.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.OptionSelfCloseLocalID, &ptr.OptionSelfCloseLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ActionLocalID, &ptr.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.OrderActionStatus = C.char(s.OrderActionStatus)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ptr.OptSelfCloseFlag = C.char(s.OptSelfCloseFlag)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ptr
}

type CThostFtdcSyncDelaySwapField struct {
	DelaySwapSeqNo     string
	BrokerID           string
	InvestorID         string
	FromCurrencyID     string
	FromAmount         float64
	FromFrozenSwap     float64
	FromRemainSwap     float64
	ToCurrencyID       string
	ToAmount           float64
	IsManualSwap       int
	IsAllRemainSetZero int
}

func NewCThostFtdcSyncDelaySwapField(p *C.CThostFtdcSyncDelaySwapField) *CThostFtdcSyncDelaySwapField {
	ret := new(CThostFtdcSyncDelaySwapField)
	ret.DelaySwapSeqNo = c2goStr(&p.DelaySwapSeqNo[0], C.sizeof_TThostFtdcDepositSeqNoType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.FromCurrencyID = c2goStr(&p.FromCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.FromAmount = goFloat64(p.FromAmount)
	ret.FromFrozenSwap = goFloat64(p.FromFrozenSwap)
	ret.FromRemainSwap = goFloat64(p.FromRemainSwap)
	ret.ToCurrencyID = c2goStr(&p.ToCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.ToAmount = goFloat64(p.ToAmount)
	ret.IsManualSwap = int(p.IsManualSwap)
	ret.IsAllRemainSetZero = int(p.IsAllRemainSetZero)
	return ret
}
func (s *CThostFtdcSyncDelaySwapField) CValue() *C.CThostFtdcSyncDelaySwapField {
	ptr := C.newCThostFtdcSyncDelaySwapField()
	go2cStr(s.DelaySwapSeqNo, &ptr.DelaySwapSeqNo[0], C.sizeof_TThostFtdcDepositSeqNoType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.FromCurrencyID, &ptr.FromCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.FromAmount = C.double(s.FromAmount)
	ptr.FromFrozenSwap = C.double(s.FromFrozenSwap)
	ptr.FromRemainSwap = C.double(s.FromRemainSwap)
	go2cStr(s.ToCurrencyID, &ptr.ToCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.ToAmount = C.double(s.ToAmount)
	ptr.IsManualSwap = C.int(s.IsManualSwap)
	ptr.IsAllRemainSetZero = C.int(s.IsAllRemainSetZero)
	return ptr
}

type CThostFtdcQrySyncDelaySwapField struct {
	BrokerID       string
	DelaySwapSeqNo string
}

func NewCThostFtdcQrySyncDelaySwapField(p *C.CThostFtdcQrySyncDelaySwapField) *CThostFtdcQrySyncDelaySwapField {
	ret := new(CThostFtdcQrySyncDelaySwapField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.DelaySwapSeqNo = c2goStr(&p.DelaySwapSeqNo[0], C.sizeof_TThostFtdcDepositSeqNoType)
	return ret
}
func (s *CThostFtdcQrySyncDelaySwapField) CValue() *C.CThostFtdcQrySyncDelaySwapField {
	ptr := C.newCThostFtdcQrySyncDelaySwapField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.DelaySwapSeqNo, &ptr.DelaySwapSeqNo[0], C.sizeof_TThostFtdcDepositSeqNoType)
	return ptr
}

type CThostFtdcInvestUnitField struct {
	BrokerID         string
	InvestorID       string
	InvestUnitID     string
	InvestorUnitName string
	InvestorGroupID  string
	CommModelID      string
	MarginModelID    string
	AccountID        string
	CurrencyID       string
}

func NewCThostFtdcInvestUnitField(p *C.CThostFtdcInvestUnitField) *CThostFtdcInvestUnitField {
	ret := new(CThostFtdcInvestUnitField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InvestorUnitName = c2goStr(&p.InvestorUnitName[0], C.sizeof_TThostFtdcPartyNameType)
	ret.InvestorGroupID = c2goStr(&p.InvestorGroupID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.CommModelID = c2goStr(&p.CommModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.MarginModelID = c2goStr(&p.MarginModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ret
}
func (s *CThostFtdcInvestUnitField) CValue() *C.CThostFtdcInvestUnitField {
	ptr := C.newCThostFtdcInvestUnitField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InvestorUnitName, &ptr.InvestorUnitName[0], C.sizeof_TThostFtdcPartyNameType)
	go2cStr(s.InvestorGroupID, &ptr.InvestorGroupID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.CommModelID, &ptr.CommModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.MarginModelID, &ptr.MarginModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ptr
}

type CThostFtdcQryInvestUnitField struct {
	BrokerID     string
	InvestorID   string
	InvestUnitID string
}

func NewCThostFtdcQryInvestUnitField(p *C.CThostFtdcQryInvestUnitField) *CThostFtdcQryInvestUnitField {
	ret := new(CThostFtdcQryInvestUnitField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ret
}
func (s *CThostFtdcQryInvestUnitField) CValue() *C.CThostFtdcQryInvestUnitField {
	ptr := C.newCThostFtdcQryInvestUnitField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ptr
}

type CThostFtdcSecAgentCheckModeField struct {
	InvestorID       string
	BrokerID         string
	CurrencyID       string
	BrokerSecAgentID string
	CheckSelfAccount int
}

func NewCThostFtdcSecAgentCheckModeField(p *C.CThostFtdcSecAgentCheckModeField) *CThostFtdcSecAgentCheckModeField {
	ret := new(CThostFtdcSecAgentCheckModeField)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.BrokerSecAgentID = c2goStr(&p.BrokerSecAgentID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CheckSelfAccount = int(p.CheckSelfAccount)
	return ret
}
func (s *CThostFtdcSecAgentCheckModeField) CValue() *C.CThostFtdcSecAgentCheckModeField {
	ptr := C.newCThostFtdcSecAgentCheckModeField()
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.BrokerSecAgentID, &ptr.BrokerSecAgentID[0], C.sizeof_TThostFtdcAccountIDType)
	ptr.CheckSelfAccount = C.int(s.CheckSelfAccount)
	return ptr
}

type CThostFtdcSecAgentTradeInfoField struct {
	BrokerID         string
	BrokerSecAgentID string
	InvestorID       string
	LongCustomerName string
}

func NewCThostFtdcSecAgentTradeInfoField(p *C.CThostFtdcSecAgentTradeInfoField) *CThostFtdcSecAgentTradeInfoField {
	ret := new(CThostFtdcSecAgentTradeInfoField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerSecAgentID = c2goStr(&p.BrokerSecAgentID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcSecAgentTradeInfoField) CValue() *C.CThostFtdcSecAgentTradeInfoField {
	ptr := C.newCThostFtdcSecAgentTradeInfoField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerSecAgentID, &ptr.BrokerSecAgentID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcMarketDataField struct {
	TradingDay         string
	Reserve1           string
	ExchangeID         string
	Reserve2           string
	LastPrice          float64
	PreSettlementPrice float64
	PreClosePrice      float64
	PreOpenInterest    float64
	OpenPrice          float64
	HighestPrice       float64
	LowestPrice        float64
	Volume             int
	Turnover           float64
	OpenInterest       float64
	ClosePrice         float64
	SettlementPrice    float64
	UpperLimitPrice    float64
	LowerLimitPrice    float64
	PreDelta           float64
	CurrDelta          float64
	UpdateTime         string
	UpdateMillisec     int
	ActionDay          string
	InstrumentID       string
	ExchangeInstID     string
}

func NewCThostFtdcMarketDataField(p *C.CThostFtdcMarketDataField) *CThostFtdcMarketDataField {
	ret := new(CThostFtdcMarketDataField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.LastPrice = goFloat64(p.LastPrice)
	ret.PreSettlementPrice = goFloat64(p.PreSettlementPrice)
	ret.PreClosePrice = goFloat64(p.PreClosePrice)
	ret.PreOpenInterest = goFloat64(p.PreOpenInterest)
	ret.OpenPrice = goFloat64(p.OpenPrice)
	ret.HighestPrice = goFloat64(p.HighestPrice)
	ret.LowestPrice = goFloat64(p.LowestPrice)
	ret.Volume = int(p.Volume)
	ret.Turnover = goFloat64(p.Turnover)
	ret.OpenInterest = goFloat64(p.OpenInterest)
	ret.ClosePrice = goFloat64(p.ClosePrice)
	ret.SettlementPrice = goFloat64(p.SettlementPrice)
	ret.UpperLimitPrice = goFloat64(p.UpperLimitPrice)
	ret.LowerLimitPrice = goFloat64(p.LowerLimitPrice)
	ret.PreDelta = goFloat64(p.PreDelta)
	ret.CurrDelta = goFloat64(p.CurrDelta)
	ret.UpdateTime = c2goStr(&p.UpdateTime[0], C.sizeof_TThostFtdcTimeType)
	ret.UpdateMillisec = int(p.UpdateMillisec)
	ret.ActionDay = c2goStr(&p.ActionDay[0], C.sizeof_TThostFtdcDateType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ret
}
func (s *CThostFtdcMarketDataField) CValue() *C.CThostFtdcMarketDataField {
	ptr := C.newCThostFtdcMarketDataField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ptr.LastPrice = C.double(s.LastPrice)
	ptr.PreSettlementPrice = C.double(s.PreSettlementPrice)
	ptr.PreClosePrice = C.double(s.PreClosePrice)
	ptr.PreOpenInterest = C.double(s.PreOpenInterest)
	ptr.OpenPrice = C.double(s.OpenPrice)
	ptr.HighestPrice = C.double(s.HighestPrice)
	ptr.LowestPrice = C.double(s.LowestPrice)
	ptr.Volume = C.int(s.Volume)
	ptr.Turnover = C.double(s.Turnover)
	ptr.OpenInterest = C.double(s.OpenInterest)
	ptr.ClosePrice = C.double(s.ClosePrice)
	ptr.SettlementPrice = C.double(s.SettlementPrice)
	ptr.UpperLimitPrice = C.double(s.UpperLimitPrice)
	ptr.LowerLimitPrice = C.double(s.LowerLimitPrice)
	ptr.PreDelta = C.double(s.PreDelta)
	ptr.CurrDelta = C.double(s.CurrDelta)
	go2cStr(s.UpdateTime, &ptr.UpdateTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.UpdateMillisec = C.int(s.UpdateMillisec)
	go2cStr(s.ActionDay, &ptr.ActionDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ptr
}

type CThostFtdcMarketDataBaseField struct {
	TradingDay         string
	PreSettlementPrice float64
	PreClosePrice      float64
	PreOpenInterest    float64
	PreDelta           float64
}

func NewCThostFtdcMarketDataBaseField(p *C.CThostFtdcMarketDataBaseField) *CThostFtdcMarketDataBaseField {
	ret := new(CThostFtdcMarketDataBaseField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.PreSettlementPrice = goFloat64(p.PreSettlementPrice)
	ret.PreClosePrice = goFloat64(p.PreClosePrice)
	ret.PreOpenInterest = goFloat64(p.PreOpenInterest)
	ret.PreDelta = goFloat64(p.PreDelta)
	return ret
}
func (s *CThostFtdcMarketDataBaseField) CValue() *C.CThostFtdcMarketDataBaseField {
	ptr := C.newCThostFtdcMarketDataBaseField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.PreSettlementPrice = C.double(s.PreSettlementPrice)
	ptr.PreClosePrice = C.double(s.PreClosePrice)
	ptr.PreOpenInterest = C.double(s.PreOpenInterest)
	ptr.PreDelta = C.double(s.PreDelta)
	return ptr
}

type CThostFtdcMarketDataStaticField struct {
	OpenPrice       float64
	HighestPrice    float64
	LowestPrice     float64
	ClosePrice      float64
	UpperLimitPrice float64
	LowerLimitPrice float64
	SettlementPrice float64
	CurrDelta       float64
}

func NewCThostFtdcMarketDataStaticField(p *C.CThostFtdcMarketDataStaticField) *CThostFtdcMarketDataStaticField {
	ret := new(CThostFtdcMarketDataStaticField)
	ret.OpenPrice = goFloat64(p.OpenPrice)
	ret.HighestPrice = goFloat64(p.HighestPrice)
	ret.LowestPrice = goFloat64(p.LowestPrice)
	ret.ClosePrice = goFloat64(p.ClosePrice)
	ret.UpperLimitPrice = goFloat64(p.UpperLimitPrice)
	ret.LowerLimitPrice = goFloat64(p.LowerLimitPrice)
	ret.SettlementPrice = goFloat64(p.SettlementPrice)
	ret.CurrDelta = goFloat64(p.CurrDelta)
	return ret
}
func (s *CThostFtdcMarketDataStaticField) CValue() *C.CThostFtdcMarketDataStaticField {
	ptr := C.newCThostFtdcMarketDataStaticField()
	ptr.OpenPrice = C.double(s.OpenPrice)
	ptr.HighestPrice = C.double(s.HighestPrice)
	ptr.LowestPrice = C.double(s.LowestPrice)
	ptr.ClosePrice = C.double(s.ClosePrice)
	ptr.UpperLimitPrice = C.double(s.UpperLimitPrice)
	ptr.LowerLimitPrice = C.double(s.LowerLimitPrice)
	ptr.SettlementPrice = C.double(s.SettlementPrice)
	ptr.CurrDelta = C.double(s.CurrDelta)
	return ptr
}

type CThostFtdcMarketDataLastMatchField struct {
	LastPrice    float64
	Volume       int
	Turnover     float64
	OpenInterest float64
}

func NewCThostFtdcMarketDataLastMatchField(p *C.CThostFtdcMarketDataLastMatchField) *CThostFtdcMarketDataLastMatchField {
	ret := new(CThostFtdcMarketDataLastMatchField)
	ret.LastPrice = goFloat64(p.LastPrice)
	ret.Volume = int(p.Volume)
	ret.Turnover = goFloat64(p.Turnover)
	ret.OpenInterest = goFloat64(p.OpenInterest)
	return ret
}
func (s *CThostFtdcMarketDataLastMatchField) CValue() *C.CThostFtdcMarketDataLastMatchField {
	ptr := C.newCThostFtdcMarketDataLastMatchField()
	ptr.LastPrice = C.double(s.LastPrice)
	ptr.Volume = C.int(s.Volume)
	ptr.Turnover = C.double(s.Turnover)
	ptr.OpenInterest = C.double(s.OpenInterest)
	return ptr
}

type CThostFtdcMarketDataBestPriceField struct {
	BidPrice1  float64
	BidVolume1 int
	AskPrice1  float64
	AskVolume1 int
}

func NewCThostFtdcMarketDataBestPriceField(p *C.CThostFtdcMarketDataBestPriceField) *CThostFtdcMarketDataBestPriceField {
	ret := new(CThostFtdcMarketDataBestPriceField)
	ret.BidPrice1 = goFloat64(p.BidPrice1)
	ret.BidVolume1 = int(p.BidVolume1)
	ret.AskPrice1 = goFloat64(p.AskPrice1)
	ret.AskVolume1 = int(p.AskVolume1)
	return ret
}
func (s *CThostFtdcMarketDataBestPriceField) CValue() *C.CThostFtdcMarketDataBestPriceField {
	ptr := C.newCThostFtdcMarketDataBestPriceField()
	ptr.BidPrice1 = C.double(s.BidPrice1)
	ptr.BidVolume1 = C.int(s.BidVolume1)
	ptr.AskPrice1 = C.double(s.AskPrice1)
	ptr.AskVolume1 = C.int(s.AskVolume1)
	return ptr
}

type CThostFtdcMarketDataBid23Field struct {
	BidPrice2  float64
	BidVolume2 int
	BidPrice3  float64
	BidVolume3 int
}

func NewCThostFtdcMarketDataBid23Field(p *C.CThostFtdcMarketDataBid23Field) *CThostFtdcMarketDataBid23Field {
	ret := new(CThostFtdcMarketDataBid23Field)
	ret.BidPrice2 = goFloat64(p.BidPrice2)
	ret.BidVolume2 = int(p.BidVolume2)
	ret.BidPrice3 = goFloat64(p.BidPrice3)
	ret.BidVolume3 = int(p.BidVolume3)
	return ret
}
func (s *CThostFtdcMarketDataBid23Field) CValue() *C.CThostFtdcMarketDataBid23Field {
	ptr := C.newCThostFtdcMarketDataBid23Field()
	ptr.BidPrice2 = C.double(s.BidPrice2)
	ptr.BidVolume2 = C.int(s.BidVolume2)
	ptr.BidPrice3 = C.double(s.BidPrice3)
	ptr.BidVolume3 = C.int(s.BidVolume3)
	return ptr
}

type CThostFtdcMarketDataAsk23Field struct {
	AskPrice2  float64
	AskVolume2 int
	AskPrice3  float64
	AskVolume3 int
}

func NewCThostFtdcMarketDataAsk23Field(p *C.CThostFtdcMarketDataAsk23Field) *CThostFtdcMarketDataAsk23Field {
	ret := new(CThostFtdcMarketDataAsk23Field)
	ret.AskPrice2 = goFloat64(p.AskPrice2)
	ret.AskVolume2 = int(p.AskVolume2)
	ret.AskPrice3 = goFloat64(p.AskPrice3)
	ret.AskVolume3 = int(p.AskVolume3)
	return ret
}
func (s *CThostFtdcMarketDataAsk23Field) CValue() *C.CThostFtdcMarketDataAsk23Field {
	ptr := C.newCThostFtdcMarketDataAsk23Field()
	ptr.AskPrice2 = C.double(s.AskPrice2)
	ptr.AskVolume2 = C.int(s.AskVolume2)
	ptr.AskPrice3 = C.double(s.AskPrice3)
	ptr.AskVolume3 = C.int(s.AskVolume3)
	return ptr
}

type CThostFtdcMarketDataBid45Field struct {
	BidPrice4  float64
	BidVolume4 int
	BidPrice5  float64
	BidVolume5 int
}

func NewCThostFtdcMarketDataBid45Field(p *C.CThostFtdcMarketDataBid45Field) *CThostFtdcMarketDataBid45Field {
	ret := new(CThostFtdcMarketDataBid45Field)
	ret.BidPrice4 = goFloat64(p.BidPrice4)
	ret.BidVolume4 = int(p.BidVolume4)
	ret.BidPrice5 = goFloat64(p.BidPrice5)
	ret.BidVolume5 = int(p.BidVolume5)
	return ret
}
func (s *CThostFtdcMarketDataBid45Field) CValue() *C.CThostFtdcMarketDataBid45Field {
	ptr := C.newCThostFtdcMarketDataBid45Field()
	ptr.BidPrice4 = C.double(s.BidPrice4)
	ptr.BidVolume4 = C.int(s.BidVolume4)
	ptr.BidPrice5 = C.double(s.BidPrice5)
	ptr.BidVolume5 = C.int(s.BidVolume5)
	return ptr
}

type CThostFtdcMarketDataAsk45Field struct {
	AskPrice4  float64
	AskVolume4 int
	AskPrice5  float64
	AskVolume5 int
}

func NewCThostFtdcMarketDataAsk45Field(p *C.CThostFtdcMarketDataAsk45Field) *CThostFtdcMarketDataAsk45Field {
	ret := new(CThostFtdcMarketDataAsk45Field)
	ret.AskPrice4 = goFloat64(p.AskPrice4)
	ret.AskVolume4 = int(p.AskVolume4)
	ret.AskPrice5 = goFloat64(p.AskPrice5)
	ret.AskVolume5 = int(p.AskVolume5)
	return ret
}
func (s *CThostFtdcMarketDataAsk45Field) CValue() *C.CThostFtdcMarketDataAsk45Field {
	ptr := C.newCThostFtdcMarketDataAsk45Field()
	ptr.AskPrice4 = C.double(s.AskPrice4)
	ptr.AskVolume4 = C.int(s.AskVolume4)
	ptr.AskPrice5 = C.double(s.AskPrice5)
	ptr.AskVolume5 = C.int(s.AskVolume5)
	return ptr
}

type CThostFtdcMarketDataUpdateTimeField struct {
	Reserve1       string
	UpdateTime     string
	UpdateMillisec int
	ActionDay      string
	InstrumentID   string
}

func NewCThostFtdcMarketDataUpdateTimeField(p *C.CThostFtdcMarketDataUpdateTimeField) *CThostFtdcMarketDataUpdateTimeField {
	ret := new(CThostFtdcMarketDataUpdateTimeField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.UpdateTime = c2goStr(&p.UpdateTime[0], C.sizeof_TThostFtdcTimeType)
	ret.UpdateMillisec = int(p.UpdateMillisec)
	ret.ActionDay = c2goStr(&p.ActionDay[0], C.sizeof_TThostFtdcDateType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcMarketDataUpdateTimeField) CValue() *C.CThostFtdcMarketDataUpdateTimeField {
	ptr := C.newCThostFtdcMarketDataUpdateTimeField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.UpdateTime, &ptr.UpdateTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.UpdateMillisec = C.int(s.UpdateMillisec)
	go2cStr(s.ActionDay, &ptr.ActionDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcMarketDataBandingPriceField struct {
	BandingUpperPrice float64
	BandingLowerPrice float64
}

func NewCThostFtdcMarketDataBandingPriceField(p *C.CThostFtdcMarketDataBandingPriceField) *CThostFtdcMarketDataBandingPriceField {
	ret := new(CThostFtdcMarketDataBandingPriceField)
	ret.BandingUpperPrice = goFloat64(p.BandingUpperPrice)
	ret.BandingLowerPrice = goFloat64(p.BandingLowerPrice)
	return ret
}
func (s *CThostFtdcMarketDataBandingPriceField) CValue() *C.CThostFtdcMarketDataBandingPriceField {
	ptr := C.newCThostFtdcMarketDataBandingPriceField()
	ptr.BandingUpperPrice = C.double(s.BandingUpperPrice)
	ptr.BandingLowerPrice = C.double(s.BandingLowerPrice)
	return ptr
}

type CThostFtdcMarketDataExchangeField struct {
	ExchangeID string
}

func NewCThostFtdcMarketDataExchangeField(p *C.CThostFtdcMarketDataExchangeField) *CThostFtdcMarketDataExchangeField {
	ret := new(CThostFtdcMarketDataExchangeField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ret
}
func (s *CThostFtdcMarketDataExchangeField) CValue() *C.CThostFtdcMarketDataExchangeField {
	ptr := C.newCThostFtdcMarketDataExchangeField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ptr
}

type CThostFtdcSpecificInstrumentField struct {
	Reserve1     string
	InstrumentID string
}

func NewCThostFtdcSpecificInstrumentField(p *C.CThostFtdcSpecificInstrumentField) *CThostFtdcSpecificInstrumentField {
	ret := new(CThostFtdcSpecificInstrumentField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcSpecificInstrumentField) CValue() *C.CThostFtdcSpecificInstrumentField {
	ptr := C.newCThostFtdcSpecificInstrumentField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcInstrumentStatusField struct {
	ExchangeID        string
	Reserve1          string
	SettlementGroupID string
	Reserve2          string
	InstrumentStatus  byte
	TradingSegmentSN  int
	EnterTime         string
	EnterReason       byte
	ExchangeInstID    string
	InstrumentID      string
}

func NewCThostFtdcInstrumentStatusField(p *C.CThostFtdcInstrumentStatusField) *CThostFtdcInstrumentStatusField {
	ret := new(CThostFtdcInstrumentStatusField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.SettlementGroupID = c2goStr(&p.SettlementGroupID[0], C.sizeof_TThostFtdcSettlementGroupIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InstrumentStatus = byte(p.InstrumentStatus)
	ret.TradingSegmentSN = int(p.TradingSegmentSN)
	ret.EnterTime = c2goStr(&p.EnterTime[0], C.sizeof_TThostFtdcTimeType)
	ret.EnterReason = byte(p.EnterReason)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcInstrumentStatusField) CValue() *C.CThostFtdcInstrumentStatusField {
	ptr := C.newCThostFtdcInstrumentStatusField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.SettlementGroupID, &ptr.SettlementGroupID[0], C.sizeof_TThostFtdcSettlementGroupIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InstrumentStatus = C.char(s.InstrumentStatus)
	ptr.TradingSegmentSN = C.int(s.TradingSegmentSN)
	go2cStr(s.EnterTime, &ptr.EnterTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.EnterReason = C.char(s.EnterReason)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryInstrumentStatusField struct {
	ExchangeID     string
	Reserve1       string
	ExchangeInstID string
}

func NewCThostFtdcQryInstrumentStatusField(p *C.CThostFtdcQryInstrumentStatusField) *CThostFtdcQryInstrumentStatusField {
	ret := new(CThostFtdcQryInstrumentStatusField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ret
}
func (s *CThostFtdcQryInstrumentStatusField) CValue() *C.CThostFtdcQryInstrumentStatusField {
	ptr := C.newCThostFtdcQryInstrumentStatusField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	return ptr
}

type CThostFtdcInvestorAccountField struct {
	BrokerID   string
	InvestorID string
	AccountID  string
	CurrencyID string
}

func NewCThostFtdcInvestorAccountField(p *C.CThostFtdcInvestorAccountField) *CThostFtdcInvestorAccountField {
	ret := new(CThostFtdcInvestorAccountField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ret
}
func (s *CThostFtdcInvestorAccountField) CValue() *C.CThostFtdcInvestorAccountField {
	ptr := C.newCThostFtdcInvestorAccountField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ptr
}

type CThostFtdcPositionProfitAlgorithmField struct {
	BrokerID   string
	AccountID  string
	Algorithm  byte
	Memo       string
	CurrencyID string
}

func NewCThostFtdcPositionProfitAlgorithmField(p *C.CThostFtdcPositionProfitAlgorithmField) *CThostFtdcPositionProfitAlgorithmField {
	ret := new(CThostFtdcPositionProfitAlgorithmField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Algorithm = byte(p.Algorithm)
	ret.Memo = c2goStr(&p.Memo[0], C.sizeof_TThostFtdcMemoType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ret
}
func (s *CThostFtdcPositionProfitAlgorithmField) CValue() *C.CThostFtdcPositionProfitAlgorithmField {
	ptr := C.newCThostFtdcPositionProfitAlgorithmField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ptr.Algorithm = C.char(s.Algorithm)
	go2cStr(s.Memo, &ptr.Memo[0], C.sizeof_TThostFtdcMemoType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ptr
}

type CThostFtdcDiscountField struct {
	BrokerID      string
	InvestorRange byte
	InvestorID    string
	Discount      float64
}

func NewCThostFtdcDiscountField(p *C.CThostFtdcDiscountField) *CThostFtdcDiscountField {
	ret := new(CThostFtdcDiscountField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Discount = goFloat64(p.Discount)
	return ret
}
func (s *CThostFtdcDiscountField) CValue() *C.CThostFtdcDiscountField {
	ptr := C.newCThostFtdcDiscountField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.Discount = C.double(s.Discount)
	return ptr
}

type CThostFtdcQryTransferBankField struct {
	BankID     string
	BankBrchID string
}

func NewCThostFtdcQryTransferBankField(p *C.CThostFtdcQryTransferBankField) *CThostFtdcQryTransferBankField {
	ret := new(CThostFtdcQryTransferBankField)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBrchID = c2goStr(&p.BankBrchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	return ret
}
func (s *CThostFtdcQryTransferBankField) CValue() *C.CThostFtdcQryTransferBankField {
	ptr := C.newCThostFtdcQryTransferBankField()
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBrchID, &ptr.BankBrchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	return ptr
}

type CThostFtdcTransferBankField struct {
	BankID     string
	BankBrchID string
	BankName   string
	IsActive   int
}

func NewCThostFtdcTransferBankField(p *C.CThostFtdcTransferBankField) *CThostFtdcTransferBankField {
	ret := new(CThostFtdcTransferBankField)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBrchID = c2goStr(&p.BankBrchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BankName = c2goStr(&p.BankName[0], C.sizeof_TThostFtdcBankNameType)
	ret.IsActive = int(p.IsActive)
	return ret
}
func (s *CThostFtdcTransferBankField) CValue() *C.CThostFtdcTransferBankField {
	ptr := C.newCThostFtdcTransferBankField()
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBrchID, &ptr.BankBrchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BankName, &ptr.BankName[0], C.sizeof_TThostFtdcBankNameType)
	ptr.IsActive = C.int(s.IsActive)
	return ptr
}

type CThostFtdcQryInvestorPositionDetailField struct {
	BrokerID     string
	InvestorID   string
	Reserve1     string
	ExchangeID   string
	InvestUnitID string
	InstrumentID string
}

func NewCThostFtdcQryInvestorPositionDetailField(p *C.CThostFtdcQryInvestorPositionDetailField) *CThostFtdcQryInvestorPositionDetailField {
	ret := new(CThostFtdcQryInvestorPositionDetailField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryInvestorPositionDetailField) CValue() *C.CThostFtdcQryInvestorPositionDetailField {
	ptr := C.newCThostFtdcQryInvestorPositionDetailField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcInvestorPositionDetailField struct {
	Reserve1              string
	BrokerID              string
	InvestorID            string
	HedgeFlag             byte
	Direction             byte
	OpenDate              string
	TradeID               string
	Volume                int
	OpenPrice             float64
	TradingDay            string
	SettlementID          int
	TradeType             byte
	Reserve2              string
	ExchangeID            string
	CloseProfitByDate     float64
	CloseProfitByTrade    float64
	PositionProfitByDate  float64
	PositionProfitByTrade float64
	Margin                float64
	ExchMargin            float64
	MarginRateByMoney     float64
	MarginRateByVolume    float64
	LastSettlementPrice   float64
	SettlementPrice       float64
	CloseVolume           int
	CloseAmount           float64
	TimeFirstVolume       int
	InvestUnitID          string
	SpecPosiType          byte
	InstrumentID          string
	CombInstrumentID      string
}

func NewCThostFtdcInvestorPositionDetailField(p *C.CThostFtdcInvestorPositionDetailField) *CThostFtdcInvestorPositionDetailField {
	ret := new(CThostFtdcInvestorPositionDetailField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.Direction = byte(p.Direction)
	ret.OpenDate = c2goStr(&p.OpenDate[0], C.sizeof_TThostFtdcDateType)
	ret.TradeID = c2goStr(&p.TradeID[0], C.sizeof_TThostFtdcTradeIDType)
	ret.Volume = int(p.Volume)
	ret.OpenPrice = goFloat64(p.OpenPrice)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.TradeType = byte(p.TradeType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.CloseProfitByDate = goFloat64(p.CloseProfitByDate)
	ret.CloseProfitByTrade = goFloat64(p.CloseProfitByTrade)
	ret.PositionProfitByDate = goFloat64(p.PositionProfitByDate)
	ret.PositionProfitByTrade = goFloat64(p.PositionProfitByTrade)
	ret.Margin = goFloat64(p.Margin)
	ret.ExchMargin = goFloat64(p.ExchMargin)
	ret.MarginRateByMoney = goFloat64(p.MarginRateByMoney)
	ret.MarginRateByVolume = goFloat64(p.MarginRateByVolume)
	ret.LastSettlementPrice = goFloat64(p.LastSettlementPrice)
	ret.SettlementPrice = goFloat64(p.SettlementPrice)
	ret.CloseVolume = int(p.CloseVolume)
	ret.CloseAmount = goFloat64(p.CloseAmount)
	ret.TimeFirstVolume = int(p.TimeFirstVolume)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.SpecPosiType = byte(p.SpecPosiType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.CombInstrumentID = c2goStr(&p.CombInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcInvestorPositionDetailField) CValue() *C.CThostFtdcInvestorPositionDetailField {
	ptr := C.newCThostFtdcInvestorPositionDetailField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.Direction = C.char(s.Direction)
	go2cStr(s.OpenDate, &ptr.OpenDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.TradeID, &ptr.TradeID[0], C.sizeof_TThostFtdcTradeIDType)
	ptr.Volume = C.int(s.Volume)
	ptr.OpenPrice = C.double(s.OpenPrice)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	ptr.TradeType = C.char(s.TradeType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ptr.CloseProfitByDate = C.double(s.CloseProfitByDate)
	ptr.CloseProfitByTrade = C.double(s.CloseProfitByTrade)
	ptr.PositionProfitByDate = C.double(s.PositionProfitByDate)
	ptr.PositionProfitByTrade = C.double(s.PositionProfitByTrade)
	ptr.Margin = C.double(s.Margin)
	ptr.ExchMargin = C.double(s.ExchMargin)
	ptr.MarginRateByMoney = C.double(s.MarginRateByMoney)
	ptr.MarginRateByVolume = C.double(s.MarginRateByVolume)
	ptr.LastSettlementPrice = C.double(s.LastSettlementPrice)
	ptr.SettlementPrice = C.double(s.SettlementPrice)
	ptr.CloseVolume = C.int(s.CloseVolume)
	ptr.CloseAmount = C.double(s.CloseAmount)
	ptr.TimeFirstVolume = C.int(s.TimeFirstVolume)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ptr.SpecPosiType = C.char(s.SpecPosiType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.CombInstrumentID, &ptr.CombInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcTradingAccountPasswordField struct {
	BrokerID   string
	AccountID  string
	Password   string
	CurrencyID string
}

func NewCThostFtdcTradingAccountPasswordField(p *C.CThostFtdcTradingAccountPasswordField) *CThostFtdcTradingAccountPasswordField {
	ret := new(CThostFtdcTradingAccountPasswordField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ret
}
func (s *CThostFtdcTradingAccountPasswordField) CValue() *C.CThostFtdcTradingAccountPasswordField {
	ptr := C.newCThostFtdcTradingAccountPasswordField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ptr
}

type CThostFtdcMDTraderOfferField struct {
	ExchangeID               string
	TraderID                 string
	ParticipantID            string
	Password                 string
	InstallID                int
	OrderLocalID             string
	TraderConnectStatus      byte
	ConnectRequestDate       string
	ConnectRequestTime       string
	LastReportDate           string
	LastReportTime           string
	ConnectDate              string
	ConnectTime              string
	StartDate                string
	StartTime                string
	TradingDay               string
	BrokerID                 string
	MaxTradeID               string
	MaxOrderMessageReference string
}

func NewCThostFtdcMDTraderOfferField(p *C.CThostFtdcMDTraderOfferField) *CThostFtdcMDTraderOfferField {
	ret := new(CThostFtdcMDTraderOfferField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.InstallID = int(p.InstallID)
	ret.OrderLocalID = c2goStr(&p.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.TraderConnectStatus = byte(p.TraderConnectStatus)
	ret.ConnectRequestDate = c2goStr(&p.ConnectRequestDate[0], C.sizeof_TThostFtdcDateType)
	ret.ConnectRequestTime = c2goStr(&p.ConnectRequestTime[0], C.sizeof_TThostFtdcTimeType)
	ret.LastReportDate = c2goStr(&p.LastReportDate[0], C.sizeof_TThostFtdcDateType)
	ret.LastReportTime = c2goStr(&p.LastReportTime[0], C.sizeof_TThostFtdcTimeType)
	ret.ConnectDate = c2goStr(&p.ConnectDate[0], C.sizeof_TThostFtdcDateType)
	ret.ConnectTime = c2goStr(&p.ConnectTime[0], C.sizeof_TThostFtdcTimeType)
	ret.StartDate = c2goStr(&p.StartDate[0], C.sizeof_TThostFtdcDateType)
	ret.StartTime = c2goStr(&p.StartTime[0], C.sizeof_TThostFtdcTimeType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.MaxTradeID = c2goStr(&p.MaxTradeID[0], C.sizeof_TThostFtdcTradeIDType)
	ret.MaxOrderMessageReference = c2goStr(&p.MaxOrderMessageReference[0], C.sizeof_TThostFtdcReturnCodeType)
	return ret
}
func (s *CThostFtdcMDTraderOfferField) CValue() *C.CThostFtdcMDTraderOfferField {
	ptr := C.newCThostFtdcMDTraderOfferField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.OrderLocalID, &ptr.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ptr.TraderConnectStatus = C.char(s.TraderConnectStatus)
	go2cStr(s.ConnectRequestDate, &ptr.ConnectRequestDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ConnectRequestTime, &ptr.ConnectRequestTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.LastReportDate, &ptr.LastReportDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.LastReportTime, &ptr.LastReportTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.ConnectDate, &ptr.ConnectDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ConnectTime, &ptr.ConnectTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.StartDate, &ptr.StartDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.StartTime, &ptr.StartTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.MaxTradeID, &ptr.MaxTradeID[0], C.sizeof_TThostFtdcTradeIDType)
	go2cStr(s.MaxOrderMessageReference, &ptr.MaxOrderMessageReference[0], C.sizeof_TThostFtdcReturnCodeType)
	return ptr
}

type CThostFtdcQryMDTraderOfferField struct {
	ExchangeID    string
	ParticipantID string
	TraderID      string
}

func NewCThostFtdcQryMDTraderOfferField(p *C.CThostFtdcQryMDTraderOfferField) *CThostFtdcQryMDTraderOfferField {
	ret := new(CThostFtdcQryMDTraderOfferField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	return ret
}
func (s *CThostFtdcQryMDTraderOfferField) CValue() *C.CThostFtdcQryMDTraderOfferField {
	ptr := C.newCThostFtdcQryMDTraderOfferField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	return ptr
}

type CThostFtdcQryNoticeField struct {
	BrokerID string
}

func NewCThostFtdcQryNoticeField(p *C.CThostFtdcQryNoticeField) *CThostFtdcQryNoticeField {
	ret := new(CThostFtdcQryNoticeField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	return ret
}
func (s *CThostFtdcQryNoticeField) CValue() *C.CThostFtdcQryNoticeField {
	ptr := C.newCThostFtdcQryNoticeField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	return ptr
}

type CThostFtdcNoticeField struct {
	BrokerID      string
	Content       string
	SequenceLabel string
}

func NewCThostFtdcNoticeField(p *C.CThostFtdcNoticeField) *CThostFtdcNoticeField {
	ret := new(CThostFtdcNoticeField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.Content = c2goStr(&p.Content[0], C.sizeof_TThostFtdcContentType)
	ret.SequenceLabel = c2goStr(&p.SequenceLabel[0], C.sizeof_TThostFtdcSequenceLabelType)
	return ret
}
func (s *CThostFtdcNoticeField) CValue() *C.CThostFtdcNoticeField {
	ptr := C.newCThostFtdcNoticeField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.Content, &ptr.Content[0], C.sizeof_TThostFtdcContentType)
	go2cStr(s.SequenceLabel, &ptr.SequenceLabel[0], C.sizeof_TThostFtdcSequenceLabelType)
	return ptr
}

type CThostFtdcUserRightField struct {
	BrokerID      string
	UserID        string
	UserRightType byte
	IsForbidden   int
}

func NewCThostFtdcUserRightField(p *C.CThostFtdcUserRightField) *CThostFtdcUserRightField {
	ret := new(CThostFtdcUserRightField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.UserRightType = byte(p.UserRightType)
	ret.IsForbidden = int(p.IsForbidden)
	return ret
}
func (s *CThostFtdcUserRightField) CValue() *C.CThostFtdcUserRightField {
	ptr := C.newCThostFtdcUserRightField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.UserRightType = C.char(s.UserRightType)
	ptr.IsForbidden = C.int(s.IsForbidden)
	return ptr
}

type CThostFtdcQrySettlementInfoConfirmField struct {
	BrokerID   string
	InvestorID string
	AccountID  string
	CurrencyID string
}

func NewCThostFtdcQrySettlementInfoConfirmField(p *C.CThostFtdcQrySettlementInfoConfirmField) *CThostFtdcQrySettlementInfoConfirmField {
	ret := new(CThostFtdcQrySettlementInfoConfirmField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ret
}
func (s *CThostFtdcQrySettlementInfoConfirmField) CValue() *C.CThostFtdcQrySettlementInfoConfirmField {
	ptr := C.newCThostFtdcQrySettlementInfoConfirmField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ptr
}

type CThostFtdcLoadSettlementInfoField struct {
	BrokerID string
}

func NewCThostFtdcLoadSettlementInfoField(p *C.CThostFtdcLoadSettlementInfoField) *CThostFtdcLoadSettlementInfoField {
	ret := new(CThostFtdcLoadSettlementInfoField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	return ret
}
func (s *CThostFtdcLoadSettlementInfoField) CValue() *C.CThostFtdcLoadSettlementInfoField {
	ptr := C.newCThostFtdcLoadSettlementInfoField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	return ptr
}

type CThostFtdcBrokerWithdrawAlgorithmField struct {
	BrokerID                string
	WithdrawAlgorithm       byte
	UsingRatio              float64
	IncludeCloseProfit      byte
	AllWithoutTrade         byte
	AvailIncludeCloseProfit byte
	IsBrokerUserEvent       int
	CurrencyID              string
	FundMortgageRatio       float64
	BalanceAlgorithm        byte
}

func NewCThostFtdcBrokerWithdrawAlgorithmField(p *C.CThostFtdcBrokerWithdrawAlgorithmField) *CThostFtdcBrokerWithdrawAlgorithmField {
	ret := new(CThostFtdcBrokerWithdrawAlgorithmField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.WithdrawAlgorithm = byte(p.WithdrawAlgorithm)
	ret.UsingRatio = goFloat64(p.UsingRatio)
	ret.IncludeCloseProfit = byte(p.IncludeCloseProfit)
	ret.AllWithoutTrade = byte(p.AllWithoutTrade)
	ret.AvailIncludeCloseProfit = byte(p.AvailIncludeCloseProfit)
	ret.IsBrokerUserEvent = int(p.IsBrokerUserEvent)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.FundMortgageRatio = goFloat64(p.FundMortgageRatio)
	ret.BalanceAlgorithm = byte(p.BalanceAlgorithm)
	return ret
}
func (s *CThostFtdcBrokerWithdrawAlgorithmField) CValue() *C.CThostFtdcBrokerWithdrawAlgorithmField {
	ptr := C.newCThostFtdcBrokerWithdrawAlgorithmField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ptr.WithdrawAlgorithm = C.char(s.WithdrawAlgorithm)
	ptr.UsingRatio = C.double(s.UsingRatio)
	ptr.IncludeCloseProfit = C.char(s.IncludeCloseProfit)
	ptr.AllWithoutTrade = C.char(s.AllWithoutTrade)
	ptr.AvailIncludeCloseProfit = C.char(s.AvailIncludeCloseProfit)
	ptr.IsBrokerUserEvent = C.int(s.IsBrokerUserEvent)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.FundMortgageRatio = C.double(s.FundMortgageRatio)
	ptr.BalanceAlgorithm = C.char(s.BalanceAlgorithm)
	return ptr
}

type CThostFtdcTradingAccountPasswordUpdateV1Field struct {
	BrokerID    string
	InvestorID  string
	OldPassword string
	NewPassword string
}

func NewCThostFtdcTradingAccountPasswordUpdateV1Field(p *C.CThostFtdcTradingAccountPasswordUpdateV1Field) *CThostFtdcTradingAccountPasswordUpdateV1Field {
	ret := new(CThostFtdcTradingAccountPasswordUpdateV1Field)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.OldPassword = c2goStr(&p.OldPassword[0], C.sizeof_TThostFtdcPasswordType)
	ret.NewPassword = c2goStr(&p.NewPassword[0], C.sizeof_TThostFtdcPasswordType)
	return ret
}
func (s *CThostFtdcTradingAccountPasswordUpdateV1Field) CValue() *C.CThostFtdcTradingAccountPasswordUpdateV1Field {
	ptr := C.newCThostFtdcTradingAccountPasswordUpdateV1Field()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.OldPassword, &ptr.OldPassword[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.NewPassword, &ptr.NewPassword[0], C.sizeof_TThostFtdcPasswordType)
	return ptr
}

type CThostFtdcTradingAccountPasswordUpdateField struct {
	BrokerID    string
	AccountID   string
	OldPassword string
	NewPassword string
	CurrencyID  string
}

func NewCThostFtdcTradingAccountPasswordUpdateField(p *C.CThostFtdcTradingAccountPasswordUpdateField) *CThostFtdcTradingAccountPasswordUpdateField {
	ret := new(CThostFtdcTradingAccountPasswordUpdateField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.OldPassword = c2goStr(&p.OldPassword[0], C.sizeof_TThostFtdcPasswordType)
	ret.NewPassword = c2goStr(&p.NewPassword[0], C.sizeof_TThostFtdcPasswordType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ret
}
func (s *CThostFtdcTradingAccountPasswordUpdateField) CValue() *C.CThostFtdcTradingAccountPasswordUpdateField {
	ptr := C.newCThostFtdcTradingAccountPasswordUpdateField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.OldPassword, &ptr.OldPassword[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.NewPassword, &ptr.NewPassword[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ptr
}

type CThostFtdcQryCombinationLegField struct {
	Reserve1         string
	LegID            int
	Reserve2         string
	CombInstrumentID string
	LegInstrumentID  string
}

func NewCThostFtdcQryCombinationLegField(p *C.CThostFtdcQryCombinationLegField) *CThostFtdcQryCombinationLegField {
	ret := new(CThostFtdcQryCombinationLegField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.LegID = int(p.LegID)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.CombInstrumentID = c2goStr(&p.CombInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.LegInstrumentID = c2goStr(&p.LegInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryCombinationLegField) CValue() *C.CThostFtdcQryCombinationLegField {
	ptr := C.newCThostFtdcQryCombinationLegField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.LegID = C.int(s.LegID)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.CombInstrumentID, &ptr.CombInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.LegInstrumentID, &ptr.LegInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQrySyncStatusField struct {
	TradingDay string
}

func NewCThostFtdcQrySyncStatusField(p *C.CThostFtdcQrySyncStatusField) *CThostFtdcQrySyncStatusField {
	ret := new(CThostFtdcQrySyncStatusField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	return ret
}
func (s *CThostFtdcQrySyncStatusField) CValue() *C.CThostFtdcQrySyncStatusField {
	ptr := C.newCThostFtdcQrySyncStatusField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	return ptr
}

type CThostFtdcCombinationLegField struct {
	Reserve1         string
	LegID            int
	Reserve2         string
	Direction        byte
	LegMultiple      int
	ImplyLevel       int
	CombInstrumentID string
	LegInstrumentID  string
}

func NewCThostFtdcCombinationLegField(p *C.CThostFtdcCombinationLegField) *CThostFtdcCombinationLegField {
	ret := new(CThostFtdcCombinationLegField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.LegID = int(p.LegID)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.Direction = byte(p.Direction)
	ret.LegMultiple = int(p.LegMultiple)
	ret.ImplyLevel = int(p.ImplyLevel)
	ret.CombInstrumentID = c2goStr(&p.CombInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.LegInstrumentID = c2goStr(&p.LegInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcCombinationLegField) CValue() *C.CThostFtdcCombinationLegField {
	ptr := C.newCThostFtdcCombinationLegField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.LegID = C.int(s.LegID)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.Direction = C.char(s.Direction)
	ptr.LegMultiple = C.int(s.LegMultiple)
	ptr.ImplyLevel = C.int(s.ImplyLevel)
	go2cStr(s.CombInstrumentID, &ptr.CombInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.LegInstrumentID, &ptr.LegInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcSyncStatusField struct {
	TradingDay     string
	DataSyncStatus byte
}

func NewCThostFtdcSyncStatusField(p *C.CThostFtdcSyncStatusField) *CThostFtdcSyncStatusField {
	ret := new(CThostFtdcSyncStatusField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.DataSyncStatus = byte(p.DataSyncStatus)
	return ret
}
func (s *CThostFtdcSyncStatusField) CValue() *C.CThostFtdcSyncStatusField {
	ptr := C.newCThostFtdcSyncStatusField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.DataSyncStatus = C.char(s.DataSyncStatus)
	return ptr
}

type CThostFtdcQryLinkManField struct {
	BrokerID   string
	InvestorID string
}

func NewCThostFtdcQryLinkManField(p *C.CThostFtdcQryLinkManField) *CThostFtdcQryLinkManField {
	ret := new(CThostFtdcQryLinkManField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ret
}
func (s *CThostFtdcQryLinkManField) CValue() *C.CThostFtdcQryLinkManField {
	ptr := C.newCThostFtdcQryLinkManField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ptr
}

type CThostFtdcLinkManField struct {
	BrokerID           string
	InvestorID         string
	PersonType         byte
	IdentifiedCardType byte
	IdentifiedCardNo   string
	PersonName         string
	Telephone          string
	Address            string
	ZipCode            string
	Priority           int
	UOAZipCode         string
	PersonFullName     string
}

func NewCThostFtdcLinkManField(p *C.CThostFtdcLinkManField) *CThostFtdcLinkManField {
	ret := new(CThostFtdcLinkManField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.PersonType = byte(p.PersonType)
	ret.IdentifiedCardType = byte(p.IdentifiedCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.PersonName = c2goStr(&p.PersonName[0], C.sizeof_TThostFtdcPartyNameType)
	ret.Telephone = c2goStr(&p.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	ret.Address = c2goStr(&p.Address[0], C.sizeof_TThostFtdcAddressType)
	ret.ZipCode = c2goStr(&p.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	ret.Priority = int(p.Priority)
	ret.UOAZipCode = c2goStr(&p.UOAZipCode[0], C.sizeof_TThostFtdcUOAZipCodeType)
	ret.PersonFullName = c2goStr(&p.PersonFullName[0], C.sizeof_TThostFtdcInvestorFullNameType)
	return ret
}
func (s *CThostFtdcLinkManField) CValue() *C.CThostFtdcLinkManField {
	ptr := C.newCThostFtdcLinkManField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.PersonType = C.char(s.PersonType)
	ptr.IdentifiedCardType = C.char(s.IdentifiedCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	go2cStr(s.PersonName, &ptr.PersonName[0], C.sizeof_TThostFtdcPartyNameType)
	go2cStr(s.Telephone, &ptr.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	go2cStr(s.Address, &ptr.Address[0], C.sizeof_TThostFtdcAddressType)
	go2cStr(s.ZipCode, &ptr.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	ptr.Priority = C.int(s.Priority)
	go2cStr(s.UOAZipCode, &ptr.UOAZipCode[0], C.sizeof_TThostFtdcUOAZipCodeType)
	go2cStr(s.PersonFullName, &ptr.PersonFullName[0], C.sizeof_TThostFtdcInvestorFullNameType)
	return ptr
}

type CThostFtdcQryBrokerUserEventField struct {
	BrokerID      string
	UserID        string
	UserEventType byte
}

func NewCThostFtdcQryBrokerUserEventField(p *C.CThostFtdcQryBrokerUserEventField) *CThostFtdcQryBrokerUserEventField {
	ret := new(CThostFtdcQryBrokerUserEventField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.UserEventType = byte(p.UserEventType)
	return ret
}
func (s *CThostFtdcQryBrokerUserEventField) CValue() *C.CThostFtdcQryBrokerUserEventField {
	ptr := C.newCThostFtdcQryBrokerUserEventField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.UserEventType = C.char(s.UserEventType)
	return ptr
}

type CThostFtdcBrokerUserEventField struct {
	BrokerID        string
	UserID          string
	UserEventType   byte
	EventSequenceNo int
	EventDate       string
	EventTime       string
	UserEventInfo   string
	InvestorID      string
	Reserve1        string
	InstrumentID    string
}

func NewCThostFtdcBrokerUserEventField(p *C.CThostFtdcBrokerUserEventField) *CThostFtdcBrokerUserEventField {
	ret := new(CThostFtdcBrokerUserEventField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.UserEventType = byte(p.UserEventType)
	ret.EventSequenceNo = int(p.EventSequenceNo)
	ret.EventDate = c2goStr(&p.EventDate[0], C.sizeof_TThostFtdcDateType)
	ret.EventTime = c2goStr(&p.EventTime[0], C.sizeof_TThostFtdcTimeType)
	ret.UserEventInfo = c2goStr(&p.UserEventInfo[0], C.sizeof_TThostFtdcUserEventInfoType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcBrokerUserEventField) CValue() *C.CThostFtdcBrokerUserEventField {
	ptr := C.newCThostFtdcBrokerUserEventField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.UserEventType = C.char(s.UserEventType)
	ptr.EventSequenceNo = C.int(s.EventSequenceNo)
	go2cStr(s.EventDate, &ptr.EventDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.EventTime, &ptr.EventTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.UserEventInfo, &ptr.UserEventInfo[0], C.sizeof_TThostFtdcUserEventInfoType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryContractBankField struct {
	BrokerID   string
	BankID     string
	BankBrchID string
}

func NewCThostFtdcQryContractBankField(p *C.CThostFtdcQryContractBankField) *CThostFtdcQryContractBankField {
	ret := new(CThostFtdcQryContractBankField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBrchID = c2goStr(&p.BankBrchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	return ret
}
func (s *CThostFtdcQryContractBankField) CValue() *C.CThostFtdcQryContractBankField {
	ptr := C.newCThostFtdcQryContractBankField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBrchID, &ptr.BankBrchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	return ptr
}

type CThostFtdcContractBankField struct {
	BrokerID   string
	BankID     string
	BankBrchID string
	BankName   string
}

func NewCThostFtdcContractBankField(p *C.CThostFtdcContractBankField) *CThostFtdcContractBankField {
	ret := new(CThostFtdcContractBankField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBrchID = c2goStr(&p.BankBrchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BankName = c2goStr(&p.BankName[0], C.sizeof_TThostFtdcBankNameType)
	return ret
}
func (s *CThostFtdcContractBankField) CValue() *C.CThostFtdcContractBankField {
	ptr := C.newCThostFtdcContractBankField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBrchID, &ptr.BankBrchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BankName, &ptr.BankName[0], C.sizeof_TThostFtdcBankNameType)
	return ptr
}

type CThostFtdcInvestorPositionCombineDetailField struct {
	TradingDay         string
	OpenDate           string
	ExchangeID         string
	SettlementID       int
	BrokerID           string
	InvestorID         string
	ComTradeID         string
	TradeID            string
	Reserve1           string
	HedgeFlag          byte
	Direction          byte
	TotalAmt           int
	Margin             float64
	ExchMargin         float64
	MarginRateByMoney  float64
	MarginRateByVolume float64
	LegID              int
	LegMultiple        int
	Reserve2           string
	TradeGroupID       int
	InvestUnitID       string
	InstrumentID       string
	CombInstrumentID   string
}

func NewCThostFtdcInvestorPositionCombineDetailField(p *C.CThostFtdcInvestorPositionCombineDetailField) *CThostFtdcInvestorPositionCombineDetailField {
	ret := new(CThostFtdcInvestorPositionCombineDetailField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.OpenDate = c2goStr(&p.OpenDate[0], C.sizeof_TThostFtdcDateType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.SettlementID = int(p.SettlementID)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.ComTradeID = c2goStr(&p.ComTradeID[0], C.sizeof_TThostFtdcTradeIDType)
	ret.TradeID = c2goStr(&p.TradeID[0], C.sizeof_TThostFtdcTradeIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.Direction = byte(p.Direction)
	ret.TotalAmt = int(p.TotalAmt)
	ret.Margin = goFloat64(p.Margin)
	ret.ExchMargin = goFloat64(p.ExchMargin)
	ret.MarginRateByMoney = goFloat64(p.MarginRateByMoney)
	ret.MarginRateByVolume = goFloat64(p.MarginRateByVolume)
	ret.LegID = int(p.LegID)
	ret.LegMultiple = int(p.LegMultiple)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.TradeGroupID = int(p.TradeGroupID)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.CombInstrumentID = c2goStr(&p.CombInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcInvestorPositionCombineDetailField) CValue() *C.CThostFtdcInvestorPositionCombineDetailField {
	ptr := C.newCThostFtdcInvestorPositionCombineDetailField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.OpenDate, &ptr.OpenDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ptr.SettlementID = C.int(s.SettlementID)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.ComTradeID, &ptr.ComTradeID[0], C.sizeof_TThostFtdcTradeIDType)
	go2cStr(s.TradeID, &ptr.TradeID[0], C.sizeof_TThostFtdcTradeIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.Direction = C.char(s.Direction)
	ptr.TotalAmt = C.int(s.TotalAmt)
	ptr.Margin = C.double(s.Margin)
	ptr.ExchMargin = C.double(s.ExchMargin)
	ptr.MarginRateByMoney = C.double(s.MarginRateByMoney)
	ptr.MarginRateByVolume = C.double(s.MarginRateByVolume)
	ptr.LegID = C.int(s.LegID)
	ptr.LegMultiple = C.int(s.LegMultiple)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.TradeGroupID = C.int(s.TradeGroupID)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.CombInstrumentID, &ptr.CombInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcParkedOrderField struct {
	BrokerID            string
	InvestorID          string
	Reserve1            string
	OrderRef            string
	UserID              string
	OrderPriceType      byte
	Direction           byte
	CombOffsetFlag      string
	CombHedgeFlag       string
	LimitPrice          float64
	VolumeTotalOriginal int
	TimeCondition       byte
	GTDDate             string
	VolumeCondition     byte
	MinVolume           int
	ContingentCondition byte
	StopPrice           float64
	ForceCloseReason    byte
	IsAutoSuspend       int
	BusinessUnit        string
	RequestID           int
	UserForceClose      int
	ExchangeID          string
	ParkedOrderID       string
	UserType            byte
	Status              byte
	ErrorID             int
	ErrorMsg            string
	IsSwapOrder         int
	AccountID           string
	CurrencyID          string
	ClientID            string
	InvestUnitID        string
	Reserve2            string
	MacAddress          string
	InstrumentID        string
	IPAddress           string
}

func NewCThostFtdcParkedOrderField(p *C.CThostFtdcParkedOrderField) *CThostFtdcParkedOrderField {
	ret := new(CThostFtdcParkedOrderField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.OrderRef = c2goStr(&p.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.OrderPriceType = byte(p.OrderPriceType)
	ret.Direction = byte(p.Direction)
	ret.CombOffsetFlag = c2goStr(&p.CombOffsetFlag[0], C.sizeof_TThostFtdcCombOffsetFlagType)
	ret.CombHedgeFlag = c2goStr(&p.CombHedgeFlag[0], C.sizeof_TThostFtdcCombHedgeFlagType)
	ret.LimitPrice = goFloat64(p.LimitPrice)
	ret.VolumeTotalOriginal = int(p.VolumeTotalOriginal)
	ret.TimeCondition = byte(p.TimeCondition)
	ret.GTDDate = c2goStr(&p.GTDDate[0], C.sizeof_TThostFtdcDateType)
	ret.VolumeCondition = byte(p.VolumeCondition)
	ret.MinVolume = int(p.MinVolume)
	ret.ContingentCondition = byte(p.ContingentCondition)
	ret.StopPrice = goFloat64(p.StopPrice)
	ret.ForceCloseReason = byte(p.ForceCloseReason)
	ret.IsAutoSuspend = int(p.IsAutoSuspend)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.RequestID = int(p.RequestID)
	ret.UserForceClose = int(p.UserForceClose)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParkedOrderID = c2goStr(&p.ParkedOrderID[0], C.sizeof_TThostFtdcParkedOrderIDType)
	ret.UserType = byte(p.UserType)
	ret.Status = byte(p.Status)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.IsSwapOrder = int(p.IsSwapOrder)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcParkedOrderField) CValue() *C.CThostFtdcParkedOrderField {
	ptr := C.newCThostFtdcParkedOrderField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.OrderRef, &ptr.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.OrderPriceType = C.char(s.OrderPriceType)
	ptr.Direction = C.char(s.Direction)
	go2cStr(s.CombOffsetFlag, &ptr.CombOffsetFlag[0], C.sizeof_TThostFtdcCombOffsetFlagType)
	go2cStr(s.CombHedgeFlag, &ptr.CombHedgeFlag[0], C.sizeof_TThostFtdcCombHedgeFlagType)
	ptr.LimitPrice = C.double(s.LimitPrice)
	ptr.VolumeTotalOriginal = C.int(s.VolumeTotalOriginal)
	ptr.TimeCondition = C.char(s.TimeCondition)
	go2cStr(s.GTDDate, &ptr.GTDDate[0], C.sizeof_TThostFtdcDateType)
	ptr.VolumeCondition = C.char(s.VolumeCondition)
	ptr.MinVolume = C.int(s.MinVolume)
	ptr.ContingentCondition = C.char(s.ContingentCondition)
	ptr.StopPrice = C.double(s.StopPrice)
	ptr.ForceCloseReason = C.char(s.ForceCloseReason)
	ptr.IsAutoSuspend = C.int(s.IsAutoSuspend)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.UserForceClose = C.int(s.UserForceClose)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParkedOrderID, &ptr.ParkedOrderID[0], C.sizeof_TThostFtdcParkedOrderIDType)
	ptr.UserType = C.char(s.UserType)
	ptr.Status = C.char(s.Status)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ptr.IsSwapOrder = C.int(s.IsSwapOrder)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcParkedOrderActionField struct {
	BrokerID            string
	InvestorID          string
	OrderActionRef      int
	OrderRef            string
	RequestID           int
	FrontID             int
	SessionID           int
	ExchangeID          string
	OrderSysID          string
	ActionFlag          byte
	LimitPrice          float64
	VolumeChange        int
	UserID              string
	Reserve1            string
	ParkedOrderActionID string
	UserType            byte
	Status              byte
	ErrorID             int
	ErrorMsg            string
	InvestUnitID        string
	Reserve2            string
	MacAddress          string
	InstrumentID        string
	IPAddress           string
}

func NewCThostFtdcParkedOrderActionField(p *C.CThostFtdcParkedOrderActionField) *CThostFtdcParkedOrderActionField {
	ret := new(CThostFtdcParkedOrderActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.OrderActionRef = int(p.OrderActionRef)
	ret.OrderRef = c2goStr(&p.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.RequestID = int(p.RequestID)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.OrderSysID = c2goStr(&p.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ActionFlag = byte(p.ActionFlag)
	ret.LimitPrice = goFloat64(p.LimitPrice)
	ret.VolumeChange = int(p.VolumeChange)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ParkedOrderActionID = c2goStr(&p.ParkedOrderActionID[0], C.sizeof_TThostFtdcParkedOrderActionIDType)
	ret.UserType = byte(p.UserType)
	ret.Status = byte(p.Status)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcParkedOrderActionField) CValue() *C.CThostFtdcParkedOrderActionField {
	ptr := C.newCThostFtdcParkedOrderActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.OrderActionRef = C.int(s.OrderActionRef)
	go2cStr(s.OrderRef, &ptr.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.OrderSysID, &ptr.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ptr.ActionFlag = C.char(s.ActionFlag)
	ptr.LimitPrice = C.double(s.LimitPrice)
	ptr.VolumeChange = C.int(s.VolumeChange)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ParkedOrderActionID, &ptr.ParkedOrderActionID[0], C.sizeof_TThostFtdcParkedOrderActionIDType)
	ptr.UserType = C.char(s.UserType)
	ptr.Status = C.char(s.Status)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryParkedOrderField struct {
	BrokerID     string
	InvestorID   string
	Reserve1     string
	ExchangeID   string
	InvestUnitID string
	InstrumentID string
}

func NewCThostFtdcQryParkedOrderField(p *C.CThostFtdcQryParkedOrderField) *CThostFtdcQryParkedOrderField {
	ret := new(CThostFtdcQryParkedOrderField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryParkedOrderField) CValue() *C.CThostFtdcQryParkedOrderField {
	ptr := C.newCThostFtdcQryParkedOrderField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryParkedOrderActionField struct {
	BrokerID     string
	InvestorID   string
	Reserve1     string
	ExchangeID   string
	InvestUnitID string
	InstrumentID string
}

func NewCThostFtdcQryParkedOrderActionField(p *C.CThostFtdcQryParkedOrderActionField) *CThostFtdcQryParkedOrderActionField {
	ret := new(CThostFtdcQryParkedOrderActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryParkedOrderActionField) CValue() *C.CThostFtdcQryParkedOrderActionField {
	ptr := C.newCThostFtdcQryParkedOrderActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcRemoveParkedOrderField struct {
	BrokerID      string
	InvestorID    string
	ParkedOrderID string
	InvestUnitID  string
}

func NewCThostFtdcRemoveParkedOrderField(p *C.CThostFtdcRemoveParkedOrderField) *CThostFtdcRemoveParkedOrderField {
	ret := new(CThostFtdcRemoveParkedOrderField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.ParkedOrderID = c2goStr(&p.ParkedOrderID[0], C.sizeof_TThostFtdcParkedOrderIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ret
}
func (s *CThostFtdcRemoveParkedOrderField) CValue() *C.CThostFtdcRemoveParkedOrderField {
	ptr := C.newCThostFtdcRemoveParkedOrderField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.ParkedOrderID, &ptr.ParkedOrderID[0], C.sizeof_TThostFtdcParkedOrderIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ptr
}

type CThostFtdcRemoveParkedOrderActionField struct {
	BrokerID            string
	InvestorID          string
	ParkedOrderActionID string
	InvestUnitID        string
}

func NewCThostFtdcRemoveParkedOrderActionField(p *C.CThostFtdcRemoveParkedOrderActionField) *CThostFtdcRemoveParkedOrderActionField {
	ret := new(CThostFtdcRemoveParkedOrderActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.ParkedOrderActionID = c2goStr(&p.ParkedOrderActionID[0], C.sizeof_TThostFtdcParkedOrderActionIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ret
}
func (s *CThostFtdcRemoveParkedOrderActionField) CValue() *C.CThostFtdcRemoveParkedOrderActionField {
	ptr := C.newCThostFtdcRemoveParkedOrderActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.ParkedOrderActionID, &ptr.ParkedOrderActionID[0], C.sizeof_TThostFtdcParkedOrderActionIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ptr
}

type CThostFtdcInvestorWithdrawAlgorithmField struct {
	BrokerID          string
	InvestorRange     byte
	InvestorID        string
	UsingRatio        float64
	CurrencyID        string
	FundMortgageRatio float64
}

func NewCThostFtdcInvestorWithdrawAlgorithmField(p *C.CThostFtdcInvestorWithdrawAlgorithmField) *CThostFtdcInvestorWithdrawAlgorithmField {
	ret := new(CThostFtdcInvestorWithdrawAlgorithmField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.UsingRatio = goFloat64(p.UsingRatio)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.FundMortgageRatio = goFloat64(p.FundMortgageRatio)
	return ret
}
func (s *CThostFtdcInvestorWithdrawAlgorithmField) CValue() *C.CThostFtdcInvestorWithdrawAlgorithmField {
	ptr := C.newCThostFtdcInvestorWithdrawAlgorithmField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.UsingRatio = C.double(s.UsingRatio)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.FundMortgageRatio = C.double(s.FundMortgageRatio)
	return ptr
}

type CThostFtdcQryInvestorPositionCombineDetailField struct {
	BrokerID         string
	InvestorID       string
	Reserve1         string
	ExchangeID       string
	InvestUnitID     string
	CombInstrumentID string
}

func NewCThostFtdcQryInvestorPositionCombineDetailField(p *C.CThostFtdcQryInvestorPositionCombineDetailField) *CThostFtdcQryInvestorPositionCombineDetailField {
	ret := new(CThostFtdcQryInvestorPositionCombineDetailField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.CombInstrumentID = c2goStr(&p.CombInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryInvestorPositionCombineDetailField) CValue() *C.CThostFtdcQryInvestorPositionCombineDetailField {
	ptr := C.newCThostFtdcQryInvestorPositionCombineDetailField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.CombInstrumentID, &ptr.CombInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcMarketDataAveragePriceField struct {
	AveragePrice float64
}

func NewCThostFtdcMarketDataAveragePriceField(p *C.CThostFtdcMarketDataAveragePriceField) *CThostFtdcMarketDataAveragePriceField {
	ret := new(CThostFtdcMarketDataAveragePriceField)
	ret.AveragePrice = goFloat64(p.AveragePrice)
	return ret
}
func (s *CThostFtdcMarketDataAveragePriceField) CValue() *C.CThostFtdcMarketDataAveragePriceField {
	ptr := C.newCThostFtdcMarketDataAveragePriceField()
	ptr.AveragePrice = C.double(s.AveragePrice)
	return ptr
}

type CThostFtdcVerifyInvestorPasswordField struct {
	BrokerID   string
	InvestorID string
	Password   string
}

func NewCThostFtdcVerifyInvestorPasswordField(p *C.CThostFtdcVerifyInvestorPasswordField) *CThostFtdcVerifyInvestorPasswordField {
	ret := new(CThostFtdcVerifyInvestorPasswordField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	return ret
}
func (s *CThostFtdcVerifyInvestorPasswordField) CValue() *C.CThostFtdcVerifyInvestorPasswordField {
	ptr := C.newCThostFtdcVerifyInvestorPasswordField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	return ptr
}

type CThostFtdcUserIPField struct {
	BrokerID   string
	UserID     string
	Reserve1   string
	Reserve2   string
	MacAddress string
	IPAddress  string
	IPMask     string
}

func NewCThostFtdcUserIPField(p *C.CThostFtdcUserIPField) *CThostFtdcUserIPField {
	ret := new(CThostFtdcUserIPField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	ret.IPMask = c2goStr(&p.IPMask[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcUserIPField) CValue() *C.CThostFtdcUserIPField {
	ptr := C.newCThostFtdcUserIPField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	go2cStr(s.IPMask, &ptr.IPMask[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcTradingNoticeInfoField struct {
	BrokerID       string
	InvestorID     string
	SendTime       string
	FieldContent   string
	SequenceSeries int16
	SequenceNo     int
	InvestUnitID   string
}

func NewCThostFtdcTradingNoticeInfoField(p *C.CThostFtdcTradingNoticeInfoField) *CThostFtdcTradingNoticeInfoField {
	ret := new(CThostFtdcTradingNoticeInfoField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.SendTime = c2goStr(&p.SendTime[0], C.sizeof_TThostFtdcTimeType)
	ret.FieldContent = c2goStr(&p.FieldContent[0], C.sizeof_TThostFtdcContentType)
	ret.SequenceSeries = int16(p.SequenceSeries)
	ret.SequenceNo = int(p.SequenceNo)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ret
}
func (s *CThostFtdcTradingNoticeInfoField) CValue() *C.CThostFtdcTradingNoticeInfoField {
	ptr := C.newCThostFtdcTradingNoticeInfoField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.SendTime, &ptr.SendTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.FieldContent, &ptr.FieldContent[0], C.sizeof_TThostFtdcContentType)
	ptr.SequenceSeries = C.short(s.SequenceSeries)
	ptr.SequenceNo = C.int(s.SequenceNo)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ptr
}

type CThostFtdcTradingNoticeField struct {
	BrokerID       string
	InvestorRange  byte
	InvestorID     string
	SequenceSeries int16
	UserID         string
	SendTime       string
	SequenceNo     int
	FieldContent   string
	InvestUnitID   string
}

func NewCThostFtdcTradingNoticeField(p *C.CThostFtdcTradingNoticeField) *CThostFtdcTradingNoticeField {
	ret := new(CThostFtdcTradingNoticeField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.SequenceSeries = int16(p.SequenceSeries)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.SendTime = c2goStr(&p.SendTime[0], C.sizeof_TThostFtdcTimeType)
	ret.SequenceNo = int(p.SequenceNo)
	ret.FieldContent = c2goStr(&p.FieldContent[0], C.sizeof_TThostFtdcContentType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ret
}
func (s *CThostFtdcTradingNoticeField) CValue() *C.CThostFtdcTradingNoticeField {
	ptr := C.newCThostFtdcTradingNoticeField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.SequenceSeries = C.short(s.SequenceSeries)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.SendTime, &ptr.SendTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.SequenceNo = C.int(s.SequenceNo)
	go2cStr(s.FieldContent, &ptr.FieldContent[0], C.sizeof_TThostFtdcContentType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ptr
}

type CThostFtdcQryTradingNoticeField struct {
	BrokerID     string
	InvestorID   string
	InvestUnitID string
}

func NewCThostFtdcQryTradingNoticeField(p *C.CThostFtdcQryTradingNoticeField) *CThostFtdcQryTradingNoticeField {
	ret := new(CThostFtdcQryTradingNoticeField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ret
}
func (s *CThostFtdcQryTradingNoticeField) CValue() *C.CThostFtdcQryTradingNoticeField {
	ptr := C.newCThostFtdcQryTradingNoticeField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ptr
}

type CThostFtdcQryErrOrderField struct {
	BrokerID   string
	InvestorID string
}

func NewCThostFtdcQryErrOrderField(p *C.CThostFtdcQryErrOrderField) *CThostFtdcQryErrOrderField {
	ret := new(CThostFtdcQryErrOrderField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ret
}
func (s *CThostFtdcQryErrOrderField) CValue() *C.CThostFtdcQryErrOrderField {
	ptr := C.newCThostFtdcQryErrOrderField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ptr
}

type CThostFtdcErrOrderField struct {
	BrokerID            string
	InvestorID          string
	Reserve1            string
	OrderRef            string
	UserID              string
	OrderPriceType      byte
	Direction           byte
	CombOffsetFlag      string
	CombHedgeFlag       string
	LimitPrice          float64
	VolumeTotalOriginal int
	TimeCondition       byte
	GTDDate             string
	VolumeCondition     byte
	MinVolume           int
	ContingentCondition byte
	StopPrice           float64
	ForceCloseReason    byte
	IsAutoSuspend       int
	BusinessUnit        string
	RequestID           int
	UserForceClose      int
	ErrorID             int
	ErrorMsg            string
	IsSwapOrder         int
	ExchangeID          string
	InvestUnitID        string
	AccountID           string
	CurrencyID          string
	ClientID            string
	Reserve2            string
	MacAddress          string
	InstrumentID        string
	IPAddress           string
}

func NewCThostFtdcErrOrderField(p *C.CThostFtdcErrOrderField) *CThostFtdcErrOrderField {
	ret := new(CThostFtdcErrOrderField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.OrderRef = c2goStr(&p.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.OrderPriceType = byte(p.OrderPriceType)
	ret.Direction = byte(p.Direction)
	ret.CombOffsetFlag = c2goStr(&p.CombOffsetFlag[0], C.sizeof_TThostFtdcCombOffsetFlagType)
	ret.CombHedgeFlag = c2goStr(&p.CombHedgeFlag[0], C.sizeof_TThostFtdcCombHedgeFlagType)
	ret.LimitPrice = goFloat64(p.LimitPrice)
	ret.VolumeTotalOriginal = int(p.VolumeTotalOriginal)
	ret.TimeCondition = byte(p.TimeCondition)
	ret.GTDDate = c2goStr(&p.GTDDate[0], C.sizeof_TThostFtdcDateType)
	ret.VolumeCondition = byte(p.VolumeCondition)
	ret.MinVolume = int(p.MinVolume)
	ret.ContingentCondition = byte(p.ContingentCondition)
	ret.StopPrice = goFloat64(p.StopPrice)
	ret.ForceCloseReason = byte(p.ForceCloseReason)
	ret.IsAutoSuspend = int(p.IsAutoSuspend)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.RequestID = int(p.RequestID)
	ret.UserForceClose = int(p.UserForceClose)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.IsSwapOrder = int(p.IsSwapOrder)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcErrOrderField) CValue() *C.CThostFtdcErrOrderField {
	ptr := C.newCThostFtdcErrOrderField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.OrderRef, &ptr.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.OrderPriceType = C.char(s.OrderPriceType)
	ptr.Direction = C.char(s.Direction)
	go2cStr(s.CombOffsetFlag, &ptr.CombOffsetFlag[0], C.sizeof_TThostFtdcCombOffsetFlagType)
	go2cStr(s.CombHedgeFlag, &ptr.CombHedgeFlag[0], C.sizeof_TThostFtdcCombHedgeFlagType)
	ptr.LimitPrice = C.double(s.LimitPrice)
	ptr.VolumeTotalOriginal = C.int(s.VolumeTotalOriginal)
	ptr.TimeCondition = C.char(s.TimeCondition)
	go2cStr(s.GTDDate, &ptr.GTDDate[0], C.sizeof_TThostFtdcDateType)
	ptr.VolumeCondition = C.char(s.VolumeCondition)
	ptr.MinVolume = C.int(s.MinVolume)
	ptr.ContingentCondition = C.char(s.ContingentCondition)
	ptr.StopPrice = C.double(s.StopPrice)
	ptr.ForceCloseReason = C.char(s.ForceCloseReason)
	ptr.IsAutoSuspend = C.int(s.IsAutoSuspend)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.UserForceClose = C.int(s.UserForceClose)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ptr.IsSwapOrder = C.int(s.IsSwapOrder)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcErrorConditionalOrderField struct {
	BrokerID             string
	InvestorID           string
	Reserve1             string
	OrderRef             string
	UserID               string
	OrderPriceType       byte
	Direction            byte
	CombOffsetFlag       string
	CombHedgeFlag        string
	LimitPrice           float64
	VolumeTotalOriginal  int
	TimeCondition        byte
	GTDDate              string
	VolumeCondition      byte
	MinVolume            int
	ContingentCondition  byte
	StopPrice            float64
	ForceCloseReason     byte
	IsAutoSuspend        int
	BusinessUnit         string
	RequestID            int
	OrderLocalID         string
	ExchangeID           string
	ParticipantID        string
	ClientID             string
	Reserve2             string
	TraderID             string
	InstallID            int
	OrderSubmitStatus    byte
	NotifySequence       int
	TradingDay           string
	SettlementID         int
	OrderSysID           string
	OrderSource          byte
	OrderStatus          byte
	OrderType            byte
	VolumeTraded         int
	VolumeTotal          int
	InsertDate           string
	InsertTime           string
	ActiveTime           string
	SuspendTime          string
	UpdateTime           string
	CancelTime           string
	ActiveTraderID       string
	ClearingPartID       string
	SequenceNo           int
	FrontID              int
	SessionID            int
	UserProductInfo      string
	StatusMsg            string
	UserForceClose       int
	ActiveUserID         string
	BrokerOrderSeq       int
	RelativeOrderSysID   string
	ZCETotalTradedVolume int
	ErrorID              int
	ErrorMsg             string
	IsSwapOrder          int
	BranchID             string
	InvestUnitID         string
	AccountID            string
	CurrencyID           string
	Reserve3             string
	MacAddress           string
	InstrumentID         string
	ExchangeInstID       string
	IPAddress            string
}

func NewCThostFtdcErrorConditionalOrderField(p *C.CThostFtdcErrorConditionalOrderField) *CThostFtdcErrorConditionalOrderField {
	ret := new(CThostFtdcErrorConditionalOrderField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.OrderRef = c2goStr(&p.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.OrderPriceType = byte(p.OrderPriceType)
	ret.Direction = byte(p.Direction)
	ret.CombOffsetFlag = c2goStr(&p.CombOffsetFlag[0], C.sizeof_TThostFtdcCombOffsetFlagType)
	ret.CombHedgeFlag = c2goStr(&p.CombHedgeFlag[0], C.sizeof_TThostFtdcCombHedgeFlagType)
	ret.LimitPrice = goFloat64(p.LimitPrice)
	ret.VolumeTotalOriginal = int(p.VolumeTotalOriginal)
	ret.TimeCondition = byte(p.TimeCondition)
	ret.GTDDate = c2goStr(&p.GTDDate[0], C.sizeof_TThostFtdcDateType)
	ret.VolumeCondition = byte(p.VolumeCondition)
	ret.MinVolume = int(p.MinVolume)
	ret.ContingentCondition = byte(p.ContingentCondition)
	ret.StopPrice = goFloat64(p.StopPrice)
	ret.ForceCloseReason = byte(p.ForceCloseReason)
	ret.IsAutoSuspend = int(p.IsAutoSuspend)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.RequestID = int(p.RequestID)
	ret.OrderLocalID = c2goStr(&p.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.OrderSubmitStatus = byte(p.OrderSubmitStatus)
	ret.NotifySequence = int(p.NotifySequence)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.OrderSysID = c2goStr(&p.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.OrderSource = byte(p.OrderSource)
	ret.OrderStatus = byte(p.OrderStatus)
	ret.OrderType = byte(p.OrderType)
	ret.VolumeTraded = int(p.VolumeTraded)
	ret.VolumeTotal = int(p.VolumeTotal)
	ret.InsertDate = c2goStr(&p.InsertDate[0], C.sizeof_TThostFtdcDateType)
	ret.InsertTime = c2goStr(&p.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	ret.ActiveTime = c2goStr(&p.ActiveTime[0], C.sizeof_TThostFtdcTimeType)
	ret.SuspendTime = c2goStr(&p.SuspendTime[0], C.sizeof_TThostFtdcTimeType)
	ret.UpdateTime = c2goStr(&p.UpdateTime[0], C.sizeof_TThostFtdcTimeType)
	ret.CancelTime = c2goStr(&p.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	ret.ActiveTraderID = c2goStr(&p.ActiveTraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.ClearingPartID = c2goStr(&p.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.SequenceNo = int(p.SequenceNo)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.UserProductInfo = c2goStr(&p.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.StatusMsg = c2goStr(&p.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.UserForceClose = int(p.UserForceClose)
	ret.ActiveUserID = c2goStr(&p.ActiveUserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.BrokerOrderSeq = int(p.BrokerOrderSeq)
	ret.RelativeOrderSysID = c2goStr(&p.RelativeOrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ZCETotalTradedVolume = int(p.ZCETotalTradedVolume)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.IsSwapOrder = int(p.IsSwapOrder)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.Reserve3 = c2goStr(&p.reserve3[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcErrorConditionalOrderField) CValue() *C.CThostFtdcErrorConditionalOrderField {
	ptr := C.newCThostFtdcErrorConditionalOrderField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.OrderRef, &ptr.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.OrderPriceType = C.char(s.OrderPriceType)
	ptr.Direction = C.char(s.Direction)
	go2cStr(s.CombOffsetFlag, &ptr.CombOffsetFlag[0], C.sizeof_TThostFtdcCombOffsetFlagType)
	go2cStr(s.CombHedgeFlag, &ptr.CombHedgeFlag[0], C.sizeof_TThostFtdcCombHedgeFlagType)
	ptr.LimitPrice = C.double(s.LimitPrice)
	ptr.VolumeTotalOriginal = C.int(s.VolumeTotalOriginal)
	ptr.TimeCondition = C.char(s.TimeCondition)
	go2cStr(s.GTDDate, &ptr.GTDDate[0], C.sizeof_TThostFtdcDateType)
	ptr.VolumeCondition = C.char(s.VolumeCondition)
	ptr.MinVolume = C.int(s.MinVolume)
	ptr.ContingentCondition = C.char(s.ContingentCondition)
	ptr.StopPrice = C.double(s.StopPrice)
	ptr.ForceCloseReason = C.char(s.ForceCloseReason)
	ptr.IsAutoSuspend = C.int(s.IsAutoSuspend)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.RequestID = C.int(s.RequestID)
	go2cStr(s.OrderLocalID, &ptr.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldExchangeInstIDType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.OrderSubmitStatus = C.char(s.OrderSubmitStatus)
	ptr.NotifySequence = C.int(s.NotifySequence)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	go2cStr(s.OrderSysID, &ptr.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ptr.OrderSource = C.char(s.OrderSource)
	ptr.OrderStatus = C.char(s.OrderStatus)
	ptr.OrderType = C.char(s.OrderType)
	ptr.VolumeTraded = C.int(s.VolumeTraded)
	ptr.VolumeTotal = C.int(s.VolumeTotal)
	go2cStr(s.InsertDate, &ptr.InsertDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.InsertTime, &ptr.InsertTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.ActiveTime, &ptr.ActiveTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.SuspendTime, &ptr.SuspendTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.UpdateTime, &ptr.UpdateTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.CancelTime, &ptr.CancelTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.ActiveTraderID, &ptr.ActiveTraderID[0], C.sizeof_TThostFtdcTraderIDType)
	go2cStr(s.ClearingPartID, &ptr.ClearingPartID[0], C.sizeof_TThostFtdcParticipantIDType)
	ptr.SequenceNo = C.int(s.SequenceNo)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.UserProductInfo, &ptr.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.StatusMsg, &ptr.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ptr.UserForceClose = C.int(s.UserForceClose)
	go2cStr(s.ActiveUserID, &ptr.ActiveUserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.BrokerOrderSeq = C.int(s.BrokerOrderSeq)
	go2cStr(s.RelativeOrderSysID, &ptr.RelativeOrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ptr.ZCETotalTradedVolume = C.int(s.ZCETotalTradedVolume)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ptr.IsSwapOrder = C.int(s.IsSwapOrder)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.Reserve3, &ptr.reserve3[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryErrOrderActionField struct {
	BrokerID   string
	InvestorID string
}

func NewCThostFtdcQryErrOrderActionField(p *C.CThostFtdcQryErrOrderActionField) *CThostFtdcQryErrOrderActionField {
	ret := new(CThostFtdcQryErrOrderActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ret
}
func (s *CThostFtdcQryErrOrderActionField) CValue() *C.CThostFtdcQryErrOrderActionField {
	ptr := C.newCThostFtdcQryErrOrderActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ptr
}

type CThostFtdcErrOrderActionField struct {
	BrokerID          string
	InvestorID        string
	OrderActionRef    int
	OrderRef          string
	RequestID         int
	FrontID           int
	SessionID         int
	ExchangeID        string
	OrderSysID        string
	ActionFlag        byte
	LimitPrice        float64
	VolumeChange      int
	ActionDate        string
	ActionTime        string
	TraderID          string
	InstallID         int
	OrderLocalID      string
	ActionLocalID     string
	ParticipantID     string
	ClientID          string
	BusinessUnit      string
	OrderActionStatus byte
	UserID            string
	StatusMsg         string
	Reserve1          string
	BranchID          string
	InvestUnitID      string
	Reserve2          string
	MacAddress        string
	ErrorID           int
	ErrorMsg          string
	InstrumentID      string
	IPAddress         string
}

func NewCThostFtdcErrOrderActionField(p *C.CThostFtdcErrOrderActionField) *CThostFtdcErrOrderActionField {
	ret := new(CThostFtdcErrOrderActionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.OrderActionRef = int(p.OrderActionRef)
	ret.OrderRef = c2goStr(&p.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ret.RequestID = int(p.RequestID)
	ret.FrontID = int(p.FrontID)
	ret.SessionID = int(p.SessionID)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.OrderSysID = c2goStr(&p.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ret.ActionFlag = byte(p.ActionFlag)
	ret.LimitPrice = goFloat64(p.LimitPrice)
	ret.VolumeChange = int(p.VolumeChange)
	ret.ActionDate = c2goStr(&p.ActionDate[0], C.sizeof_TThostFtdcDateType)
	ret.ActionTime = c2goStr(&p.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	ret.TraderID = c2goStr(&p.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ret.InstallID = int(p.InstallID)
	ret.OrderLocalID = c2goStr(&p.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ActionLocalID = c2goStr(&p.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ClientID = c2goStr(&p.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	ret.BusinessUnit = c2goStr(&p.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ret.OrderActionStatus = byte(p.OrderActionStatus)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.StatusMsg = c2goStr(&p.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.BranchID = c2goStr(&p.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcErrOrderActionField) CValue() *C.CThostFtdcErrOrderActionField {
	ptr := C.newCThostFtdcErrOrderActionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.OrderActionRef = C.int(s.OrderActionRef)
	go2cStr(s.OrderRef, &ptr.OrderRef[0], C.sizeof_TThostFtdcOrderRefType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.FrontID = C.int(s.FrontID)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.OrderSysID, &ptr.OrderSysID[0], C.sizeof_TThostFtdcOrderSysIDType)
	ptr.ActionFlag = C.char(s.ActionFlag)
	ptr.LimitPrice = C.double(s.LimitPrice)
	ptr.VolumeChange = C.int(s.VolumeChange)
	go2cStr(s.ActionDate, &ptr.ActionDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ActionTime, &ptr.ActionTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.TraderID, &ptr.TraderID[0], C.sizeof_TThostFtdcTraderIDType)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.OrderLocalID, &ptr.OrderLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ActionLocalID, &ptr.ActionLocalID[0], C.sizeof_TThostFtdcOrderLocalIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ClientID, &ptr.ClientID[0], C.sizeof_TThostFtdcClientIDType)
	go2cStr(s.BusinessUnit, &ptr.BusinessUnit[0], C.sizeof_TThostFtdcBusinessUnitType)
	ptr.OrderActionStatus = C.char(s.OrderActionStatus)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.StatusMsg, &ptr.StatusMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.BranchID, &ptr.BranchID[0], C.sizeof_TThostFtdcBranchIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryExchangeSequenceField struct {
	ExchangeID string
}

func NewCThostFtdcQryExchangeSequenceField(p *C.CThostFtdcQryExchangeSequenceField) *CThostFtdcQryExchangeSequenceField {
	ret := new(CThostFtdcQryExchangeSequenceField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ret
}
func (s *CThostFtdcQryExchangeSequenceField) CValue() *C.CThostFtdcQryExchangeSequenceField {
	ptr := C.newCThostFtdcQryExchangeSequenceField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ptr
}

type CThostFtdcExchangeSequenceField struct {
	ExchangeID   string
	SequenceNo   int
	MarketStatus byte
}

func NewCThostFtdcExchangeSequenceField(p *C.CThostFtdcExchangeSequenceField) *CThostFtdcExchangeSequenceField {
	ret := new(CThostFtdcExchangeSequenceField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.SequenceNo = int(p.SequenceNo)
	ret.MarketStatus = byte(p.MarketStatus)
	return ret
}
func (s *CThostFtdcExchangeSequenceField) CValue() *C.CThostFtdcExchangeSequenceField {
	ptr := C.newCThostFtdcExchangeSequenceField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ptr.SequenceNo = C.int(s.SequenceNo)
	ptr.MarketStatus = C.char(s.MarketStatus)
	return ptr
}

type CThostFtdcQryMaxOrderVolumeWithPriceField struct {
	BrokerID     string
	InvestorID   string
	Reserve1     string
	Direction    byte
	OffsetFlag   byte
	HedgeFlag    byte
	MaxVolume    int
	Price        float64
	ExchangeID   string
	InvestUnitID string
	InstrumentID string
}

func NewCThostFtdcQryMaxOrderVolumeWithPriceField(p *C.CThostFtdcQryMaxOrderVolumeWithPriceField) *CThostFtdcQryMaxOrderVolumeWithPriceField {
	ret := new(CThostFtdcQryMaxOrderVolumeWithPriceField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.Direction = byte(p.Direction)
	ret.OffsetFlag = byte(p.OffsetFlag)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.MaxVolume = int(p.MaxVolume)
	ret.Price = goFloat64(p.Price)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryMaxOrderVolumeWithPriceField) CValue() *C.CThostFtdcQryMaxOrderVolumeWithPriceField {
	ptr := C.newCThostFtdcQryMaxOrderVolumeWithPriceField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.Direction = C.char(s.Direction)
	ptr.OffsetFlag = C.char(s.OffsetFlag)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.MaxVolume = C.int(s.MaxVolume)
	ptr.Price = C.double(s.Price)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryBrokerTradingParamsField struct {
	BrokerID   string
	InvestorID string
	CurrencyID string
	AccountID  string
}

func NewCThostFtdcQryBrokerTradingParamsField(p *C.CThostFtdcQryBrokerTradingParamsField) *CThostFtdcQryBrokerTradingParamsField {
	ret := new(CThostFtdcQryBrokerTradingParamsField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	return ret
}
func (s *CThostFtdcQryBrokerTradingParamsField) CValue() *C.CThostFtdcQryBrokerTradingParamsField {
	ptr := C.newCThostFtdcQryBrokerTradingParamsField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	return ptr
}

type CThostFtdcBrokerTradingParamsField struct {
	BrokerID                string
	InvestorID              string
	MarginPriceType         byte
	Algorithm               byte
	AvailIncludeCloseProfit byte
	CurrencyID              string
	OptionRoyaltyPriceType  byte
	AccountID               string
}

func NewCThostFtdcBrokerTradingParamsField(p *C.CThostFtdcBrokerTradingParamsField) *CThostFtdcBrokerTradingParamsField {
	ret := new(CThostFtdcBrokerTradingParamsField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.MarginPriceType = byte(p.MarginPriceType)
	ret.Algorithm = byte(p.Algorithm)
	ret.AvailIncludeCloseProfit = byte(p.AvailIncludeCloseProfit)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.OptionRoyaltyPriceType = byte(p.OptionRoyaltyPriceType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	return ret
}
func (s *CThostFtdcBrokerTradingParamsField) CValue() *C.CThostFtdcBrokerTradingParamsField {
	ptr := C.newCThostFtdcBrokerTradingParamsField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.MarginPriceType = C.char(s.MarginPriceType)
	ptr.Algorithm = C.char(s.Algorithm)
	ptr.AvailIncludeCloseProfit = C.char(s.AvailIncludeCloseProfit)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.OptionRoyaltyPriceType = C.char(s.OptionRoyaltyPriceType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	return ptr
}

type CThostFtdcQryBrokerTradingAlgosField struct {
	BrokerID     string
	ExchangeID   string
	Reserve1     string
	InstrumentID string
}

func NewCThostFtdcQryBrokerTradingAlgosField(p *C.CThostFtdcQryBrokerTradingAlgosField) *CThostFtdcQryBrokerTradingAlgosField {
	ret := new(CThostFtdcQryBrokerTradingAlgosField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryBrokerTradingAlgosField) CValue() *C.CThostFtdcQryBrokerTradingAlgosField {
	ptr := C.newCThostFtdcQryBrokerTradingAlgosField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcBrokerTradingAlgosField struct {
	BrokerID                   string
	ExchangeID                 string
	Reserve1                   string
	HandlePositionAlgoID       byte
	FindMarginRateAlgoID       byte
	HandleTradingAccountAlgoID byte
	InstrumentID               string
}

func NewCThostFtdcBrokerTradingAlgosField(p *C.CThostFtdcBrokerTradingAlgosField) *CThostFtdcBrokerTradingAlgosField {
	ret := new(CThostFtdcBrokerTradingAlgosField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.HandlePositionAlgoID = byte(p.HandlePositionAlgoID)
	ret.FindMarginRateAlgoID = byte(p.FindMarginRateAlgoID)
	ret.HandleTradingAccountAlgoID = byte(p.HandleTradingAccountAlgoID)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcBrokerTradingAlgosField) CValue() *C.CThostFtdcBrokerTradingAlgosField {
	ptr := C.newCThostFtdcBrokerTradingAlgosField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.HandlePositionAlgoID = C.char(s.HandlePositionAlgoID)
	ptr.FindMarginRateAlgoID = C.char(s.FindMarginRateAlgoID)
	ptr.HandleTradingAccountAlgoID = C.char(s.HandleTradingAccountAlgoID)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQueryBrokerDepositField struct {
	BrokerID   string
	ExchangeID string
}

func NewCThostFtdcQueryBrokerDepositField(p *C.CThostFtdcQueryBrokerDepositField) *CThostFtdcQueryBrokerDepositField {
	ret := new(CThostFtdcQueryBrokerDepositField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ret
}
func (s *CThostFtdcQueryBrokerDepositField) CValue() *C.CThostFtdcQueryBrokerDepositField {
	ptr := C.newCThostFtdcQueryBrokerDepositField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	return ptr
}

type CThostFtdcBrokerDepositField struct {
	TradingDay    string
	BrokerID      string
	ParticipantID string
	ExchangeID    string
	PreBalance    float64
	CurrMargin    float64
	CloseProfit   float64
	Balance       float64
	Deposit       float64
	Withdraw      float64
	Available     float64
	Reserve       float64
	FrozenMargin  float64
}

func NewCThostFtdcBrokerDepositField(p *C.CThostFtdcBrokerDepositField) *CThostFtdcBrokerDepositField {
	ret := new(CThostFtdcBrokerDepositField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.PreBalance = goFloat64(p.PreBalance)
	ret.CurrMargin = goFloat64(p.CurrMargin)
	ret.CloseProfit = goFloat64(p.CloseProfit)
	ret.Balance = goFloat64(p.Balance)
	ret.Deposit = goFloat64(p.Deposit)
	ret.Withdraw = goFloat64(p.Withdraw)
	ret.Available = goFloat64(p.Available)
	ret.Reserve = goFloat64(p.Reserve)
	ret.FrozenMargin = goFloat64(p.FrozenMargin)
	return ret
}
func (s *CThostFtdcBrokerDepositField) CValue() *C.CThostFtdcBrokerDepositField {
	ptr := C.newCThostFtdcBrokerDepositField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ptr.PreBalance = C.double(s.PreBalance)
	ptr.CurrMargin = C.double(s.CurrMargin)
	ptr.CloseProfit = C.double(s.CloseProfit)
	ptr.Balance = C.double(s.Balance)
	ptr.Deposit = C.double(s.Deposit)
	ptr.Withdraw = C.double(s.Withdraw)
	ptr.Available = C.double(s.Available)
	ptr.Reserve = C.double(s.Reserve)
	ptr.FrozenMargin = C.double(s.FrozenMargin)
	return ptr
}

type CThostFtdcQryCFMMCBrokerKeyField struct {
	BrokerID string
}

func NewCThostFtdcQryCFMMCBrokerKeyField(p *C.CThostFtdcQryCFMMCBrokerKeyField) *CThostFtdcQryCFMMCBrokerKeyField {
	ret := new(CThostFtdcQryCFMMCBrokerKeyField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	return ret
}
func (s *CThostFtdcQryCFMMCBrokerKeyField) CValue() *C.CThostFtdcQryCFMMCBrokerKeyField {
	ptr := C.newCThostFtdcQryCFMMCBrokerKeyField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	return ptr
}

type CThostFtdcCFMMCBrokerKeyField struct {
	BrokerID      string
	ParticipantID string
	CreateDate    string
	CreateTime    string
	KeyID         int
	CurrentKey    string
	KeyKind       byte
}

func NewCThostFtdcCFMMCBrokerKeyField(p *C.CThostFtdcCFMMCBrokerKeyField) *CThostFtdcCFMMCBrokerKeyField {
	ret := new(CThostFtdcCFMMCBrokerKeyField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.CreateDate = c2goStr(&p.CreateDate[0], C.sizeof_TThostFtdcDateType)
	ret.CreateTime = c2goStr(&p.CreateTime[0], C.sizeof_TThostFtdcTimeType)
	ret.KeyID = int(p.KeyID)
	ret.CurrentKey = c2goStr(&p.CurrentKey[0], C.sizeof_TThostFtdcCFMMCKeyType)
	ret.KeyKind = byte(p.KeyKind)
	return ret
}
func (s *CThostFtdcCFMMCBrokerKeyField) CValue() *C.CThostFtdcCFMMCBrokerKeyField {
	ptr := C.newCThostFtdcCFMMCBrokerKeyField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.CreateDate, &ptr.CreateDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.CreateTime, &ptr.CreateTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.KeyID = C.int(s.KeyID)
	go2cStr(s.CurrentKey, &ptr.CurrentKey[0], C.sizeof_TThostFtdcCFMMCKeyType)
	ptr.KeyKind = C.char(s.KeyKind)
	return ptr
}

type CThostFtdcCFMMCTradingAccountKeyField struct {
	BrokerID      string
	ParticipantID string
	AccountID     string
	KeyID         int
	CurrentKey    string
}

func NewCThostFtdcCFMMCTradingAccountKeyField(p *C.CThostFtdcCFMMCTradingAccountKeyField) *CThostFtdcCFMMCTradingAccountKeyField {
	ret := new(CThostFtdcCFMMCTradingAccountKeyField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.KeyID = int(p.KeyID)
	ret.CurrentKey = c2goStr(&p.CurrentKey[0], C.sizeof_TThostFtdcCFMMCKeyType)
	return ret
}
func (s *CThostFtdcCFMMCTradingAccountKeyField) CValue() *C.CThostFtdcCFMMCTradingAccountKeyField {
	ptr := C.newCThostFtdcCFMMCTradingAccountKeyField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ptr.KeyID = C.int(s.KeyID)
	go2cStr(s.CurrentKey, &ptr.CurrentKey[0], C.sizeof_TThostFtdcCFMMCKeyType)
	return ptr
}

type CThostFtdcQryCFMMCTradingAccountKeyField struct {
	BrokerID   string
	InvestorID string
}

func NewCThostFtdcQryCFMMCTradingAccountKeyField(p *C.CThostFtdcQryCFMMCTradingAccountKeyField) *CThostFtdcQryCFMMCTradingAccountKeyField {
	ret := new(CThostFtdcQryCFMMCTradingAccountKeyField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ret
}
func (s *CThostFtdcQryCFMMCTradingAccountKeyField) CValue() *C.CThostFtdcQryCFMMCTradingAccountKeyField {
	ptr := C.newCThostFtdcQryCFMMCTradingAccountKeyField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ptr
}

type CThostFtdcBrokerUserOTPParamField struct {
	BrokerID     string
	UserID       string
	OTPVendorsID string
	SerialNumber string
	AuthKey      string
	LastDrift    int
	LastSuccess  int
	OTPType      byte
}

func NewCThostFtdcBrokerUserOTPParamField(p *C.CThostFtdcBrokerUserOTPParamField) *CThostFtdcBrokerUserOTPParamField {
	ret := new(CThostFtdcBrokerUserOTPParamField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.OTPVendorsID = c2goStr(&p.OTPVendorsID[0], C.sizeof_TThostFtdcOTPVendorsIDType)
	ret.SerialNumber = c2goStr(&p.SerialNumber[0], C.sizeof_TThostFtdcSerialNumberType)
	ret.AuthKey = c2goStr(&p.AuthKey[0], C.sizeof_TThostFtdcAuthKeyType)
	ret.LastDrift = int(p.LastDrift)
	ret.LastSuccess = int(p.LastSuccess)
	ret.OTPType = byte(p.OTPType)
	return ret
}
func (s *CThostFtdcBrokerUserOTPParamField) CValue() *C.CThostFtdcBrokerUserOTPParamField {
	ptr := C.newCThostFtdcBrokerUserOTPParamField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.OTPVendorsID, &ptr.OTPVendorsID[0], C.sizeof_TThostFtdcOTPVendorsIDType)
	go2cStr(s.SerialNumber, &ptr.SerialNumber[0], C.sizeof_TThostFtdcSerialNumberType)
	go2cStr(s.AuthKey, &ptr.AuthKey[0], C.sizeof_TThostFtdcAuthKeyType)
	ptr.LastDrift = C.int(s.LastDrift)
	ptr.LastSuccess = C.int(s.LastSuccess)
	ptr.OTPType = C.char(s.OTPType)
	return ptr
}

type CThostFtdcManualSyncBrokerUserOTPField struct {
	BrokerID  string
	UserID    string
	OTPType   byte
	FirstOTP  string
	SecondOTP string
}

func NewCThostFtdcManualSyncBrokerUserOTPField(p *C.CThostFtdcManualSyncBrokerUserOTPField) *CThostFtdcManualSyncBrokerUserOTPField {
	ret := new(CThostFtdcManualSyncBrokerUserOTPField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.OTPType = byte(p.OTPType)
	ret.FirstOTP = c2goStr(&p.FirstOTP[0], C.sizeof_TThostFtdcPasswordType)
	ret.SecondOTP = c2goStr(&p.SecondOTP[0], C.sizeof_TThostFtdcPasswordType)
	return ret
}
func (s *CThostFtdcManualSyncBrokerUserOTPField) CValue() *C.CThostFtdcManualSyncBrokerUserOTPField {
	ptr := C.newCThostFtdcManualSyncBrokerUserOTPField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.OTPType = C.char(s.OTPType)
	go2cStr(s.FirstOTP, &ptr.FirstOTP[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.SecondOTP, &ptr.SecondOTP[0], C.sizeof_TThostFtdcPasswordType)
	return ptr
}

type CThostFtdcCommRateModelField struct {
	BrokerID      string
	CommModelID   string
	CommModelName string
}

func NewCThostFtdcCommRateModelField(p *C.CThostFtdcCommRateModelField) *CThostFtdcCommRateModelField {
	ret := new(CThostFtdcCommRateModelField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.CommModelID = c2goStr(&p.CommModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.CommModelName = c2goStr(&p.CommModelName[0], C.sizeof_TThostFtdcCommModelNameType)
	return ret
}
func (s *CThostFtdcCommRateModelField) CValue() *C.CThostFtdcCommRateModelField {
	ptr := C.newCThostFtdcCommRateModelField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.CommModelID, &ptr.CommModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.CommModelName, &ptr.CommModelName[0], C.sizeof_TThostFtdcCommModelNameType)
	return ptr
}

type CThostFtdcQryCommRateModelField struct {
	BrokerID    string
	CommModelID string
}

func NewCThostFtdcQryCommRateModelField(p *C.CThostFtdcQryCommRateModelField) *CThostFtdcQryCommRateModelField {
	ret := new(CThostFtdcQryCommRateModelField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.CommModelID = c2goStr(&p.CommModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ret
}
func (s *CThostFtdcQryCommRateModelField) CValue() *C.CThostFtdcQryCommRateModelField {
	ptr := C.newCThostFtdcQryCommRateModelField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.CommModelID, &ptr.CommModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ptr
}

type CThostFtdcMarginModelField struct {
	BrokerID        string
	MarginModelID   string
	MarginModelName string
}

func NewCThostFtdcMarginModelField(p *C.CThostFtdcMarginModelField) *CThostFtdcMarginModelField {
	ret := new(CThostFtdcMarginModelField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.MarginModelID = c2goStr(&p.MarginModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.MarginModelName = c2goStr(&p.MarginModelName[0], C.sizeof_TThostFtdcCommModelNameType)
	return ret
}
func (s *CThostFtdcMarginModelField) CValue() *C.CThostFtdcMarginModelField {
	ptr := C.newCThostFtdcMarginModelField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.MarginModelID, &ptr.MarginModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.MarginModelName, &ptr.MarginModelName[0], C.sizeof_TThostFtdcCommModelNameType)
	return ptr
}

type CThostFtdcQryMarginModelField struct {
	BrokerID      string
	MarginModelID string
}

func NewCThostFtdcQryMarginModelField(p *C.CThostFtdcQryMarginModelField) *CThostFtdcQryMarginModelField {
	ret := new(CThostFtdcQryMarginModelField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.MarginModelID = c2goStr(&p.MarginModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ret
}
func (s *CThostFtdcQryMarginModelField) CValue() *C.CThostFtdcQryMarginModelField {
	ptr := C.newCThostFtdcQryMarginModelField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.MarginModelID, &ptr.MarginModelID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ptr
}

type CThostFtdcEWarrantOffsetField struct {
	TradingDay   string
	BrokerID     string
	InvestorID   string
	ExchangeID   string
	Reserve1     string
	Direction    byte
	HedgeFlag    byte
	Volume       int
	InvestUnitID string
	InstrumentID string
}

func NewCThostFtdcEWarrantOffsetField(p *C.CThostFtdcEWarrantOffsetField) *CThostFtdcEWarrantOffsetField {
	ret := new(CThostFtdcEWarrantOffsetField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.Direction = byte(p.Direction)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.Volume = int(p.Volume)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcEWarrantOffsetField) CValue() *C.CThostFtdcEWarrantOffsetField {
	ptr := C.newCThostFtdcEWarrantOffsetField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.Direction = C.char(s.Direction)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.Volume = C.int(s.Volume)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryEWarrantOffsetField struct {
	BrokerID     string
	InvestorID   string
	ExchangeID   string
	Reserve1     string
	InvestUnitID string
	InstrumentID string
}

func NewCThostFtdcQryEWarrantOffsetField(p *C.CThostFtdcQryEWarrantOffsetField) *CThostFtdcQryEWarrantOffsetField {
	ret := new(CThostFtdcQryEWarrantOffsetField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryEWarrantOffsetField) CValue() *C.CThostFtdcQryEWarrantOffsetField {
	ptr := C.newCThostFtdcQryEWarrantOffsetField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryInvestorProductGroupMarginField struct {
	BrokerID       string
	InvestorID     string
	Reserve1       string
	HedgeFlag      byte
	ExchangeID     string
	InvestUnitID   string
	ProductGroupID string
}

func NewCThostFtdcQryInvestorProductGroupMarginField(p *C.CThostFtdcQryInvestorProductGroupMarginField) *CThostFtdcQryInvestorProductGroupMarginField {
	ret := new(CThostFtdcQryInvestorProductGroupMarginField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.ProductGroupID = c2goStr(&p.ProductGroupID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryInvestorProductGroupMarginField) CValue() *C.CThostFtdcQryInvestorProductGroupMarginField {
	ptr := C.newCThostFtdcQryInvestorProductGroupMarginField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.ProductGroupID, &ptr.ProductGroupID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcInvestorProductGroupMarginField struct {
	Reserve1              string
	BrokerID              string
	InvestorID            string
	TradingDay            string
	SettlementID          int
	FrozenMargin          float64
	LongFrozenMargin      float64
	ShortFrozenMargin     float64
	UseMargin             float64
	LongUseMargin         float64
	ShortUseMargin        float64
	ExchMargin            float64
	LongExchMargin        float64
	ShortExchMargin       float64
	CloseProfit           float64
	FrozenCommission      float64
	Commission            float64
	FrozenCash            float64
	CashIn                float64
	PositionProfit        float64
	OffsetAmount          float64
	LongOffsetAmount      float64
	ShortOffsetAmount     float64
	ExchOffsetAmount      float64
	LongExchOffsetAmount  float64
	ShortExchOffsetAmount float64
	HedgeFlag             byte
	ExchangeID            string
	InvestUnitID          string
	ProductGroupID        string
}

func NewCThostFtdcInvestorProductGroupMarginField(p *C.CThostFtdcInvestorProductGroupMarginField) *CThostFtdcInvestorProductGroupMarginField {
	ret := new(CThostFtdcInvestorProductGroupMarginField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.FrozenMargin = goFloat64(p.FrozenMargin)
	ret.LongFrozenMargin = goFloat64(p.LongFrozenMargin)
	ret.ShortFrozenMargin = goFloat64(p.ShortFrozenMargin)
	ret.UseMargin = goFloat64(p.UseMargin)
	ret.LongUseMargin = goFloat64(p.LongUseMargin)
	ret.ShortUseMargin = goFloat64(p.ShortUseMargin)
	ret.ExchMargin = goFloat64(p.ExchMargin)
	ret.LongExchMargin = goFloat64(p.LongExchMargin)
	ret.ShortExchMargin = goFloat64(p.ShortExchMargin)
	ret.CloseProfit = goFloat64(p.CloseProfit)
	ret.FrozenCommission = goFloat64(p.FrozenCommission)
	ret.Commission = goFloat64(p.Commission)
	ret.FrozenCash = goFloat64(p.FrozenCash)
	ret.CashIn = goFloat64(p.CashIn)
	ret.PositionProfit = goFloat64(p.PositionProfit)
	ret.OffsetAmount = goFloat64(p.OffsetAmount)
	ret.LongOffsetAmount = goFloat64(p.LongOffsetAmount)
	ret.ShortOffsetAmount = goFloat64(p.ShortOffsetAmount)
	ret.ExchOffsetAmount = goFloat64(p.ExchOffsetAmount)
	ret.LongExchOffsetAmount = goFloat64(p.LongExchOffsetAmount)
	ret.ShortExchOffsetAmount = goFloat64(p.ShortExchOffsetAmount)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.ProductGroupID = c2goStr(&p.ProductGroupID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcInvestorProductGroupMarginField) CValue() *C.CThostFtdcInvestorProductGroupMarginField {
	ptr := C.newCThostFtdcInvestorProductGroupMarginField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	ptr.FrozenMargin = C.double(s.FrozenMargin)
	ptr.LongFrozenMargin = C.double(s.LongFrozenMargin)
	ptr.ShortFrozenMargin = C.double(s.ShortFrozenMargin)
	ptr.UseMargin = C.double(s.UseMargin)
	ptr.LongUseMargin = C.double(s.LongUseMargin)
	ptr.ShortUseMargin = C.double(s.ShortUseMargin)
	ptr.ExchMargin = C.double(s.ExchMargin)
	ptr.LongExchMargin = C.double(s.LongExchMargin)
	ptr.ShortExchMargin = C.double(s.ShortExchMargin)
	ptr.CloseProfit = C.double(s.CloseProfit)
	ptr.FrozenCommission = C.double(s.FrozenCommission)
	ptr.Commission = C.double(s.Commission)
	ptr.FrozenCash = C.double(s.FrozenCash)
	ptr.CashIn = C.double(s.CashIn)
	ptr.PositionProfit = C.double(s.PositionProfit)
	ptr.OffsetAmount = C.double(s.OffsetAmount)
	ptr.LongOffsetAmount = C.double(s.LongOffsetAmount)
	ptr.ShortOffsetAmount = C.double(s.ShortOffsetAmount)
	ptr.ExchOffsetAmount = C.double(s.ExchOffsetAmount)
	ptr.LongExchOffsetAmount = C.double(s.LongExchOffsetAmount)
	ptr.ShortExchOffsetAmount = C.double(s.ShortExchOffsetAmount)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	go2cStr(s.ProductGroupID, &ptr.ProductGroupID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQueryCFMMCTradingAccountTokenField struct {
	BrokerID     string
	InvestorID   string
	InvestUnitID string
}

func NewCThostFtdcQueryCFMMCTradingAccountTokenField(p *C.CThostFtdcQueryCFMMCTradingAccountTokenField) *CThostFtdcQueryCFMMCTradingAccountTokenField {
	ret := new(CThostFtdcQueryCFMMCTradingAccountTokenField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ret
}
func (s *CThostFtdcQueryCFMMCTradingAccountTokenField) CValue() *C.CThostFtdcQueryCFMMCTradingAccountTokenField {
	ptr := C.newCThostFtdcQueryCFMMCTradingAccountTokenField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	return ptr
}

type CThostFtdcCFMMCTradingAccountTokenField struct {
	BrokerID      string
	ParticipantID string
	AccountID     string
	KeyID         int
	Token         string
}

func NewCThostFtdcCFMMCTradingAccountTokenField(p *C.CThostFtdcCFMMCTradingAccountTokenField) *CThostFtdcCFMMCTradingAccountTokenField {
	ret := new(CThostFtdcCFMMCTradingAccountTokenField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.ParticipantID = c2goStr(&p.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.KeyID = int(p.KeyID)
	ret.Token = c2goStr(&p.Token[0], C.sizeof_TThostFtdcCFMMCTokenType)
	return ret
}
func (s *CThostFtdcCFMMCTradingAccountTokenField) CValue() *C.CThostFtdcCFMMCTradingAccountTokenField {
	ptr := C.newCThostFtdcCFMMCTradingAccountTokenField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.ParticipantID, &ptr.ParticipantID[0], C.sizeof_TThostFtdcParticipantIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ptr.KeyID = C.int(s.KeyID)
	go2cStr(s.Token, &ptr.Token[0], C.sizeof_TThostFtdcCFMMCTokenType)
	return ptr
}

type CThostFtdcQryProductGroupField struct {
	Reserve1   string
	ExchangeID string
	ProductID  string
}

func NewCThostFtdcQryProductGroupField(p *C.CThostFtdcQryProductGroupField) *CThostFtdcQryProductGroupField {
	ret := new(CThostFtdcQryProductGroupField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ProductID = c2goStr(&p.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryProductGroupField) CValue() *C.CThostFtdcQryProductGroupField {
	ptr := C.newCThostFtdcQryProductGroupField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ProductID, &ptr.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcProductGroupField struct {
	Reserve1       string
	ExchangeID     string
	Reserve2       string
	ProductID      string
	ProductGroupID string
}

func NewCThostFtdcProductGroupField(p *C.CThostFtdcProductGroupField) *CThostFtdcProductGroupField {
	ret := new(CThostFtdcProductGroupField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.Reserve2 = c2goStr(&p.reserve2[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.ProductID = c2goStr(&p.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ProductGroupID = c2goStr(&p.ProductGroupID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcProductGroupField) CValue() *C.CThostFtdcProductGroupField {
	ptr := C.newCThostFtdcProductGroupField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.Reserve2, &ptr.reserve2[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.ProductID, &ptr.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ProductGroupID, &ptr.ProductGroupID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcBulletinField struct {
	ExchangeID  string
	TradingDay  string
	BulletinID  int
	SequenceNo  int
	NewsType    string
	NewsUrgency byte
	SendTime    string
	Abstract    string
	ComeFrom    string
	Content     string
	URLLink     string
	MarketID    string
}

func NewCThostFtdcBulletinField(p *C.CThostFtdcBulletinField) *CThostFtdcBulletinField {
	ret := new(CThostFtdcBulletinField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.BulletinID = int(p.BulletinID)
	ret.SequenceNo = int(p.SequenceNo)
	ret.NewsType = c2goStr(&p.NewsType[0], C.sizeof_TThostFtdcNewsTypeType)
	ret.NewsUrgency = byte(p.NewsUrgency)
	ret.SendTime = c2goStr(&p.SendTime[0], C.sizeof_TThostFtdcTimeType)
	ret.Abstract = c2goStr(&p.Abstract[0], C.sizeof_TThostFtdcAbstractType)
	ret.ComeFrom = c2goStr(&p.ComeFrom[0], C.sizeof_TThostFtdcComeFromType)
	ret.Content = c2goStr(&p.Content[0], C.sizeof_TThostFtdcContentType)
	ret.URLLink = c2goStr(&p.URLLink[0], C.sizeof_TThostFtdcURLLinkType)
	ret.MarketID = c2goStr(&p.MarketID[0], C.sizeof_TThostFtdcMarketIDType)
	return ret
}
func (s *CThostFtdcBulletinField) CValue() *C.CThostFtdcBulletinField {
	ptr := C.newCThostFtdcBulletinField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.BulletinID = C.int(s.BulletinID)
	ptr.SequenceNo = C.int(s.SequenceNo)
	go2cStr(s.NewsType, &ptr.NewsType[0], C.sizeof_TThostFtdcNewsTypeType)
	ptr.NewsUrgency = C.char(s.NewsUrgency)
	go2cStr(s.SendTime, &ptr.SendTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.Abstract, &ptr.Abstract[0], C.sizeof_TThostFtdcAbstractType)
	go2cStr(s.ComeFrom, &ptr.ComeFrom[0], C.sizeof_TThostFtdcComeFromType)
	go2cStr(s.Content, &ptr.Content[0], C.sizeof_TThostFtdcContentType)
	go2cStr(s.URLLink, &ptr.URLLink[0], C.sizeof_TThostFtdcURLLinkType)
	go2cStr(s.MarketID, &ptr.MarketID[0], C.sizeof_TThostFtdcMarketIDType)
	return ptr
}

type CThostFtdcQryBulletinField struct {
	ExchangeID  string
	BulletinID  int
	SequenceNo  int
	NewsType    string
	NewsUrgency byte
}

func NewCThostFtdcQryBulletinField(p *C.CThostFtdcQryBulletinField) *CThostFtdcQryBulletinField {
	ret := new(CThostFtdcQryBulletinField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.BulletinID = int(p.BulletinID)
	ret.SequenceNo = int(p.SequenceNo)
	ret.NewsType = c2goStr(&p.NewsType[0], C.sizeof_TThostFtdcNewsTypeType)
	ret.NewsUrgency = byte(p.NewsUrgency)
	return ret
}
func (s *CThostFtdcQryBulletinField) CValue() *C.CThostFtdcQryBulletinField {
	ptr := C.newCThostFtdcQryBulletinField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ptr.BulletinID = C.int(s.BulletinID)
	ptr.SequenceNo = C.int(s.SequenceNo)
	go2cStr(s.NewsType, &ptr.NewsType[0], C.sizeof_TThostFtdcNewsTypeType)
	ptr.NewsUrgency = C.char(s.NewsUrgency)
	return ptr
}

type CThostFtdcMulticastInstrumentField struct {
	TopicID        int
	Reserve1       string
	InstrumentNo   int
	CodePrice      float64
	VolumeMultiple int
	PriceTick      float64
	InstrumentID   string
}

func NewCThostFtdcMulticastInstrumentField(p *C.CThostFtdcMulticastInstrumentField) *CThostFtdcMulticastInstrumentField {
	ret := new(CThostFtdcMulticastInstrumentField)
	ret.TopicID = int(p.TopicID)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InstrumentNo = int(p.InstrumentNo)
	ret.CodePrice = goFloat64(p.CodePrice)
	ret.VolumeMultiple = int(p.VolumeMultiple)
	ret.PriceTick = goFloat64(p.PriceTick)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcMulticastInstrumentField) CValue() *C.CThostFtdcMulticastInstrumentField {
	ptr := C.newCThostFtdcMulticastInstrumentField()
	ptr.TopicID = C.int(s.TopicID)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ptr.InstrumentNo = C.int(s.InstrumentNo)
	ptr.CodePrice = C.double(s.CodePrice)
	ptr.VolumeMultiple = C.int(s.VolumeMultiple)
	ptr.PriceTick = C.double(s.PriceTick)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryMulticastInstrumentField struct {
	TopicID      int
	Reserve1     string
	InstrumentID string
}

func NewCThostFtdcQryMulticastInstrumentField(p *C.CThostFtdcQryMulticastInstrumentField) *CThostFtdcQryMulticastInstrumentField {
	ret := new(CThostFtdcQryMulticastInstrumentField)
	ret.TopicID = int(p.TopicID)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryMulticastInstrumentField) CValue() *C.CThostFtdcQryMulticastInstrumentField {
	ptr := C.newCThostFtdcQryMulticastInstrumentField()
	ptr.TopicID = C.int(s.TopicID)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldInstrumentIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcAppIDAuthAssignField struct {
	BrokerID     string
	AppID        string
	DRIdentityID int
}

func NewCThostFtdcAppIDAuthAssignField(p *C.CThostFtdcAppIDAuthAssignField) *CThostFtdcAppIDAuthAssignField {
	ret := new(CThostFtdcAppIDAuthAssignField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.AppID = c2goStr(&p.AppID[0], C.sizeof_TThostFtdcAppIDType)
	ret.DRIdentityID = int(p.DRIdentityID)
	return ret
}
func (s *CThostFtdcAppIDAuthAssignField) CValue() *C.CThostFtdcAppIDAuthAssignField {
	ptr := C.newCThostFtdcAppIDAuthAssignField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.AppID, &ptr.AppID[0], C.sizeof_TThostFtdcAppIDType)
	ptr.DRIdentityID = C.int(s.DRIdentityID)
	return ptr
}

type CThostFtdcReqOpenAccountField struct {
	TradeCode          string
	BankID             string
	BankBranchID       string
	BrokerID           string
	BrokerBranchID     string
	TradeDate          string
	TradeTime          string
	BankSerial         string
	TradingDay         string
	PlateSerial        int
	LastFragment       byte
	SessionID          int
	CustomerName       string
	IdCardType         byte
	IdentifiedCardNo   string
	Gender             byte
	CountryCode        string
	CustType           byte
	Address            string
	ZipCode            string
	Telephone          string
	MobilePhone        string
	Fax                string
	EMail              string
	MoneyAccountStatus byte
	BankAccount        string
	BankPassWord       string
	AccountID          string
	Password           string
	InstallID          int
	VerifyCertNoFlag   byte
	CurrencyID         string
	CashExchangeCode   byte
	Digest             string
	BankAccType        byte
	DeviceID           string
	BankSecuAccType    byte
	BrokerIDByBank     string
	BankSecuAcc        string
	BankPwdFlag        byte
	SecuPwdFlag        byte
	OperNo             string
	TID                int
	UserID             string
	LongCustomerName   string
}

func NewCThostFtdcReqOpenAccountField(p *C.CThostFtdcReqOpenAccountField) *CThostFtdcReqOpenAccountField {
	ret := new(CThostFtdcReqOpenAccountField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.Gender = byte(p.Gender)
	ret.CountryCode = c2goStr(&p.CountryCode[0], C.sizeof_TThostFtdcCountryCodeType)
	ret.CustType = byte(p.CustType)
	ret.Address = c2goStr(&p.Address[0], C.sizeof_TThostFtdcAddressType)
	ret.ZipCode = c2goStr(&p.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	ret.Telephone = c2goStr(&p.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	ret.MobilePhone = c2goStr(&p.MobilePhone[0], C.sizeof_TThostFtdcMobilePhoneType)
	ret.Fax = c2goStr(&p.Fax[0], C.sizeof_TThostFtdcFaxType)
	ret.EMail = c2goStr(&p.EMail[0], C.sizeof_TThostFtdcEMailType)
	ret.MoneyAccountStatus = byte(p.MoneyAccountStatus)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.InstallID = int(p.InstallID)
	ret.VerifyCertNoFlag = byte(p.VerifyCertNoFlag)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.CashExchangeCode = byte(p.CashExchangeCode)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.BankAccType = byte(p.BankAccType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BankSecuAccType = byte(p.BankSecuAccType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.BankSecuAcc = c2goStr(&p.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPwdFlag = byte(p.BankPwdFlag)
	ret.SecuPwdFlag = byte(p.SecuPwdFlag)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.TID = int(p.TID)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcReqOpenAccountField) CValue() *C.CThostFtdcReqOpenAccountField {
	ptr := C.newCThostFtdcReqOpenAccountField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.Gender = C.char(s.Gender)
	go2cStr(s.CountryCode, &ptr.CountryCode[0], C.sizeof_TThostFtdcCountryCodeType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.Address, &ptr.Address[0], C.sizeof_TThostFtdcAddressType)
	go2cStr(s.ZipCode, &ptr.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	go2cStr(s.Telephone, &ptr.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	go2cStr(s.MobilePhone, &ptr.MobilePhone[0], C.sizeof_TThostFtdcMobilePhoneType)
	go2cStr(s.Fax, &ptr.Fax[0], C.sizeof_TThostFtdcFaxType)
	go2cStr(s.EMail, &ptr.EMail[0], C.sizeof_TThostFtdcEMailType)
	ptr.MoneyAccountStatus = C.char(s.MoneyAccountStatus)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.VerifyCertNoFlag = C.char(s.VerifyCertNoFlag)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.CashExchangeCode = C.char(s.CashExchangeCode)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	ptr.BankAccType = C.char(s.BankAccType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ptr.BankSecuAccType = C.char(s.BankSecuAccType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.BankSecuAcc, &ptr.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ptr.BankPwdFlag = C.char(s.BankPwdFlag)
	ptr.SecuPwdFlag = C.char(s.SecuPwdFlag)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.TID = C.int(s.TID)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcReqCancelAccountField struct {
	TradeCode          string
	BankID             string
	BankBranchID       string
	BrokerID           string
	BrokerBranchID     string
	TradeDate          string
	TradeTime          string
	BankSerial         string
	TradingDay         string
	PlateSerial        int
	LastFragment       byte
	SessionID          int
	CustomerName       string
	IdCardType         byte
	IdentifiedCardNo   string
	Gender             byte
	CountryCode        string
	CustType           byte
	Address            string
	ZipCode            string
	Telephone          string
	MobilePhone        string
	Fax                string
	EMail              string
	MoneyAccountStatus byte
	BankAccount        string
	BankPassWord       string
	AccountID          string
	Password           string
	InstallID          int
	VerifyCertNoFlag   byte
	CurrencyID         string
	CashExchangeCode   byte
	Digest             string
	BankAccType        byte
	DeviceID           string
	BankSecuAccType    byte
	BrokerIDByBank     string
	BankSecuAcc        string
	BankPwdFlag        byte
	SecuPwdFlag        byte
	OperNo             string
	TID                int
	UserID             string
	LongCustomerName   string
}

func NewCThostFtdcReqCancelAccountField(p *C.CThostFtdcReqCancelAccountField) *CThostFtdcReqCancelAccountField {
	ret := new(CThostFtdcReqCancelAccountField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.Gender = byte(p.Gender)
	ret.CountryCode = c2goStr(&p.CountryCode[0], C.sizeof_TThostFtdcCountryCodeType)
	ret.CustType = byte(p.CustType)
	ret.Address = c2goStr(&p.Address[0], C.sizeof_TThostFtdcAddressType)
	ret.ZipCode = c2goStr(&p.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	ret.Telephone = c2goStr(&p.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	ret.MobilePhone = c2goStr(&p.MobilePhone[0], C.sizeof_TThostFtdcMobilePhoneType)
	ret.Fax = c2goStr(&p.Fax[0], C.sizeof_TThostFtdcFaxType)
	ret.EMail = c2goStr(&p.EMail[0], C.sizeof_TThostFtdcEMailType)
	ret.MoneyAccountStatus = byte(p.MoneyAccountStatus)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.InstallID = int(p.InstallID)
	ret.VerifyCertNoFlag = byte(p.VerifyCertNoFlag)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.CashExchangeCode = byte(p.CashExchangeCode)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.BankAccType = byte(p.BankAccType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BankSecuAccType = byte(p.BankSecuAccType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.BankSecuAcc = c2goStr(&p.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPwdFlag = byte(p.BankPwdFlag)
	ret.SecuPwdFlag = byte(p.SecuPwdFlag)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.TID = int(p.TID)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcReqCancelAccountField) CValue() *C.CThostFtdcReqCancelAccountField {
	ptr := C.newCThostFtdcReqCancelAccountField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.Gender = C.char(s.Gender)
	go2cStr(s.CountryCode, &ptr.CountryCode[0], C.sizeof_TThostFtdcCountryCodeType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.Address, &ptr.Address[0], C.sizeof_TThostFtdcAddressType)
	go2cStr(s.ZipCode, &ptr.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	go2cStr(s.Telephone, &ptr.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	go2cStr(s.MobilePhone, &ptr.MobilePhone[0], C.sizeof_TThostFtdcMobilePhoneType)
	go2cStr(s.Fax, &ptr.Fax[0], C.sizeof_TThostFtdcFaxType)
	go2cStr(s.EMail, &ptr.EMail[0], C.sizeof_TThostFtdcEMailType)
	ptr.MoneyAccountStatus = C.char(s.MoneyAccountStatus)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.VerifyCertNoFlag = C.char(s.VerifyCertNoFlag)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.CashExchangeCode = C.char(s.CashExchangeCode)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	ptr.BankAccType = C.char(s.BankAccType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ptr.BankSecuAccType = C.char(s.BankSecuAccType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.BankSecuAcc, &ptr.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ptr.BankPwdFlag = C.char(s.BankPwdFlag)
	ptr.SecuPwdFlag = C.char(s.SecuPwdFlag)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.TID = C.int(s.TID)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcReqChangeAccountField struct {
	TradeCode          string
	BankID             string
	BankBranchID       string
	BrokerID           string
	BrokerBranchID     string
	TradeDate          string
	TradeTime          string
	BankSerial         string
	TradingDay         string
	PlateSerial        int
	LastFragment       byte
	SessionID          int
	CustomerName       string
	IdCardType         byte
	IdentifiedCardNo   string
	Gender             byte
	CountryCode        string
	CustType           byte
	Address            string
	ZipCode            string
	Telephone          string
	MobilePhone        string
	Fax                string
	EMail              string
	MoneyAccountStatus byte
	BankAccount        string
	BankPassWord       string
	NewBankAccount     string
	NewBankPassWord    string
	AccountID          string
	Password           string
	BankAccType        byte
	InstallID          int
	VerifyCertNoFlag   byte
	CurrencyID         string
	BrokerIDByBank     string
	BankPwdFlag        byte
	SecuPwdFlag        byte
	TID                int
	Digest             string
	LongCustomerName   string
}

func NewCThostFtdcReqChangeAccountField(p *C.CThostFtdcReqChangeAccountField) *CThostFtdcReqChangeAccountField {
	ret := new(CThostFtdcReqChangeAccountField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.Gender = byte(p.Gender)
	ret.CountryCode = c2goStr(&p.CountryCode[0], C.sizeof_TThostFtdcCountryCodeType)
	ret.CustType = byte(p.CustType)
	ret.Address = c2goStr(&p.Address[0], C.sizeof_TThostFtdcAddressType)
	ret.ZipCode = c2goStr(&p.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	ret.Telephone = c2goStr(&p.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	ret.MobilePhone = c2goStr(&p.MobilePhone[0], C.sizeof_TThostFtdcMobilePhoneType)
	ret.Fax = c2goStr(&p.Fax[0], C.sizeof_TThostFtdcFaxType)
	ret.EMail = c2goStr(&p.EMail[0], C.sizeof_TThostFtdcEMailType)
	ret.MoneyAccountStatus = byte(p.MoneyAccountStatus)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.NewBankAccount = c2goStr(&p.NewBankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.NewBankPassWord = c2goStr(&p.NewBankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.BankAccType = byte(p.BankAccType)
	ret.InstallID = int(p.InstallID)
	ret.VerifyCertNoFlag = byte(p.VerifyCertNoFlag)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.BankPwdFlag = byte(p.BankPwdFlag)
	ret.SecuPwdFlag = byte(p.SecuPwdFlag)
	ret.TID = int(p.TID)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcReqChangeAccountField) CValue() *C.CThostFtdcReqChangeAccountField {
	ptr := C.newCThostFtdcReqChangeAccountField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.Gender = C.char(s.Gender)
	go2cStr(s.CountryCode, &ptr.CountryCode[0], C.sizeof_TThostFtdcCountryCodeType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.Address, &ptr.Address[0], C.sizeof_TThostFtdcAddressType)
	go2cStr(s.ZipCode, &ptr.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	go2cStr(s.Telephone, &ptr.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	go2cStr(s.MobilePhone, &ptr.MobilePhone[0], C.sizeof_TThostFtdcMobilePhoneType)
	go2cStr(s.Fax, &ptr.Fax[0], C.sizeof_TThostFtdcFaxType)
	go2cStr(s.EMail, &ptr.EMail[0], C.sizeof_TThostFtdcEMailType)
	ptr.MoneyAccountStatus = C.char(s.MoneyAccountStatus)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.NewBankAccount, &ptr.NewBankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.NewBankPassWord, &ptr.NewBankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	ptr.BankAccType = C.char(s.BankAccType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.VerifyCertNoFlag = C.char(s.VerifyCertNoFlag)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ptr.BankPwdFlag = C.char(s.BankPwdFlag)
	ptr.SecuPwdFlag = C.char(s.SecuPwdFlag)
	ptr.TID = C.int(s.TID)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcReqTransferField struct {
	TradeCode         string
	BankID            string
	BankBranchID      string
	BrokerID          string
	BrokerBranchID    string
	TradeDate         string
	TradeTime         string
	BankSerial        string
	TradingDay        string
	PlateSerial       int
	LastFragment      byte
	SessionID         int
	CustomerName      string
	IdCardType        byte
	IdentifiedCardNo  string
	CustType          byte
	BankAccount       string
	BankPassWord      string
	AccountID         string
	Password          string
	InstallID         int
	FutureSerial      int
	UserID            string
	VerifyCertNoFlag  byte
	CurrencyID        string
	TradeAmount       float64
	FutureFetchAmount float64
	FeePayFlag        byte
	CustFee           float64
	BrokerFee         float64
	Message           string
	Digest            string
	BankAccType       byte
	DeviceID          string
	BankSecuAccType   byte
	BrokerIDByBank    string
	BankSecuAcc       string
	BankPwdFlag       byte
	SecuPwdFlag       byte
	OperNo            string
	RequestID         int
	TID               int
	TransferStatus    byte
	LongCustomerName  string
}

func NewCThostFtdcReqTransferField(p *C.CThostFtdcReqTransferField) *CThostFtdcReqTransferField {
	ret := new(CThostFtdcReqTransferField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.CustType = byte(p.CustType)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.InstallID = int(p.InstallID)
	ret.FutureSerial = int(p.FutureSerial)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.VerifyCertNoFlag = byte(p.VerifyCertNoFlag)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.TradeAmount = goFloat64(p.TradeAmount)
	ret.FutureFetchAmount = goFloat64(p.FutureFetchAmount)
	ret.FeePayFlag = byte(p.FeePayFlag)
	ret.CustFee = goFloat64(p.CustFee)
	ret.BrokerFee = goFloat64(p.BrokerFee)
	ret.Message = c2goStr(&p.Message[0], C.sizeof_TThostFtdcAddInfoType)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.BankAccType = byte(p.BankAccType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BankSecuAccType = byte(p.BankSecuAccType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.BankSecuAcc = c2goStr(&p.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPwdFlag = byte(p.BankPwdFlag)
	ret.SecuPwdFlag = byte(p.SecuPwdFlag)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.RequestID = int(p.RequestID)
	ret.TID = int(p.TID)
	ret.TransferStatus = byte(p.TransferStatus)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcReqTransferField) CValue() *C.CThostFtdcReqTransferField {
	ptr := C.newCThostFtdcReqTransferField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.FutureSerial = C.int(s.FutureSerial)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.VerifyCertNoFlag = C.char(s.VerifyCertNoFlag)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.TradeAmount = C.double(s.TradeAmount)
	ptr.FutureFetchAmount = C.double(s.FutureFetchAmount)
	ptr.FeePayFlag = C.char(s.FeePayFlag)
	ptr.CustFee = C.double(s.CustFee)
	ptr.BrokerFee = C.double(s.BrokerFee)
	go2cStr(s.Message, &ptr.Message[0], C.sizeof_TThostFtdcAddInfoType)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	ptr.BankAccType = C.char(s.BankAccType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ptr.BankSecuAccType = C.char(s.BankSecuAccType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.BankSecuAcc, &ptr.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ptr.BankPwdFlag = C.char(s.BankPwdFlag)
	ptr.SecuPwdFlag = C.char(s.SecuPwdFlag)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.TID = C.int(s.TID)
	ptr.TransferStatus = C.char(s.TransferStatus)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcRspTransferField struct {
	TradeCode         string
	BankID            string
	BankBranchID      string
	BrokerID          string
	BrokerBranchID    string
	TradeDate         string
	TradeTime         string
	BankSerial        string
	TradingDay        string
	PlateSerial       int
	LastFragment      byte
	SessionID         int
	CustomerName      string
	IdCardType        byte
	IdentifiedCardNo  string
	CustType          byte
	BankAccount       string
	BankPassWord      string
	AccountID         string
	Password          string
	InstallID         int
	FutureSerial      int
	UserID            string
	VerifyCertNoFlag  byte
	CurrencyID        string
	TradeAmount       float64
	FutureFetchAmount float64
	FeePayFlag        byte
	CustFee           float64
	BrokerFee         float64
	Message           string
	Digest            string
	BankAccType       byte
	DeviceID          string
	BankSecuAccType   byte
	BrokerIDByBank    string
	BankSecuAcc       string
	BankPwdFlag       byte
	SecuPwdFlag       byte
	OperNo            string
	RequestID         int
	TID               int
	TransferStatus    byte
	ErrorID           int
	ErrorMsg          string
	LongCustomerName  string
}

func NewCThostFtdcRspTransferField(p *C.CThostFtdcRspTransferField) *CThostFtdcRspTransferField {
	ret := new(CThostFtdcRspTransferField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.CustType = byte(p.CustType)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.InstallID = int(p.InstallID)
	ret.FutureSerial = int(p.FutureSerial)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.VerifyCertNoFlag = byte(p.VerifyCertNoFlag)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.TradeAmount = goFloat64(p.TradeAmount)
	ret.FutureFetchAmount = goFloat64(p.FutureFetchAmount)
	ret.FeePayFlag = byte(p.FeePayFlag)
	ret.CustFee = goFloat64(p.CustFee)
	ret.BrokerFee = goFloat64(p.BrokerFee)
	ret.Message = c2goStr(&p.Message[0], C.sizeof_TThostFtdcAddInfoType)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.BankAccType = byte(p.BankAccType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BankSecuAccType = byte(p.BankSecuAccType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.BankSecuAcc = c2goStr(&p.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPwdFlag = byte(p.BankPwdFlag)
	ret.SecuPwdFlag = byte(p.SecuPwdFlag)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.RequestID = int(p.RequestID)
	ret.TID = int(p.TID)
	ret.TransferStatus = byte(p.TransferStatus)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcRspTransferField) CValue() *C.CThostFtdcRspTransferField {
	ptr := C.newCThostFtdcRspTransferField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.FutureSerial = C.int(s.FutureSerial)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.VerifyCertNoFlag = C.char(s.VerifyCertNoFlag)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.TradeAmount = C.double(s.TradeAmount)
	ptr.FutureFetchAmount = C.double(s.FutureFetchAmount)
	ptr.FeePayFlag = C.char(s.FeePayFlag)
	ptr.CustFee = C.double(s.CustFee)
	ptr.BrokerFee = C.double(s.BrokerFee)
	go2cStr(s.Message, &ptr.Message[0], C.sizeof_TThostFtdcAddInfoType)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	ptr.BankAccType = C.char(s.BankAccType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ptr.BankSecuAccType = C.char(s.BankSecuAccType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.BankSecuAcc, &ptr.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ptr.BankPwdFlag = C.char(s.BankPwdFlag)
	ptr.SecuPwdFlag = C.char(s.SecuPwdFlag)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.TID = C.int(s.TID)
	ptr.TransferStatus = C.char(s.TransferStatus)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcReqRepealField struct {
	RepealTimeInterval int
	RepealedTimes      int
	BankRepealFlag     byte
	BrokerRepealFlag   byte
	PlateRepealSerial  int
	BankRepealSerial   string
	FutureRepealSerial int
	TradeCode          string
	BankID             string
	BankBranchID       string
	BrokerID           string
	BrokerBranchID     string
	TradeDate          string
	TradeTime          string
	BankSerial         string
	TradingDay         string
	PlateSerial        int
	LastFragment       byte
	SessionID          int
	CustomerName       string
	IdCardType         byte
	IdentifiedCardNo   string
	CustType           byte
	BankAccount        string
	BankPassWord       string
	AccountID          string
	Password           string
	InstallID          int
	FutureSerial       int
	UserID             string
	VerifyCertNoFlag   byte
	CurrencyID         string
	TradeAmount        float64
	FutureFetchAmount  float64
	FeePayFlag         byte
	CustFee            float64
	BrokerFee          float64
	Message            string
	Digest             string
	BankAccType        byte
	DeviceID           string
	BankSecuAccType    byte
	BrokerIDByBank     string
	BankSecuAcc        string
	BankPwdFlag        byte
	SecuPwdFlag        byte
	OperNo             string
	RequestID          int
	TID                int
	TransferStatus     byte
	LongCustomerName   string
}

func NewCThostFtdcReqRepealField(p *C.CThostFtdcReqRepealField) *CThostFtdcReqRepealField {
	ret := new(CThostFtdcReqRepealField)
	ret.RepealTimeInterval = int(p.RepealTimeInterval)
	ret.RepealedTimes = int(p.RepealedTimes)
	ret.BankRepealFlag = byte(p.BankRepealFlag)
	ret.BrokerRepealFlag = byte(p.BrokerRepealFlag)
	ret.PlateRepealSerial = int(p.PlateRepealSerial)
	ret.BankRepealSerial = c2goStr(&p.BankRepealSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.FutureRepealSerial = int(p.FutureRepealSerial)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.CustType = byte(p.CustType)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.InstallID = int(p.InstallID)
	ret.FutureSerial = int(p.FutureSerial)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.VerifyCertNoFlag = byte(p.VerifyCertNoFlag)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.TradeAmount = goFloat64(p.TradeAmount)
	ret.FutureFetchAmount = goFloat64(p.FutureFetchAmount)
	ret.FeePayFlag = byte(p.FeePayFlag)
	ret.CustFee = goFloat64(p.CustFee)
	ret.BrokerFee = goFloat64(p.BrokerFee)
	ret.Message = c2goStr(&p.Message[0], C.sizeof_TThostFtdcAddInfoType)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.BankAccType = byte(p.BankAccType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BankSecuAccType = byte(p.BankSecuAccType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.BankSecuAcc = c2goStr(&p.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPwdFlag = byte(p.BankPwdFlag)
	ret.SecuPwdFlag = byte(p.SecuPwdFlag)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.RequestID = int(p.RequestID)
	ret.TID = int(p.TID)
	ret.TransferStatus = byte(p.TransferStatus)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcReqRepealField) CValue() *C.CThostFtdcReqRepealField {
	ptr := C.newCThostFtdcReqRepealField()
	ptr.RepealTimeInterval = C.int(s.RepealTimeInterval)
	ptr.RepealedTimes = C.int(s.RepealedTimes)
	ptr.BankRepealFlag = C.char(s.BankRepealFlag)
	ptr.BrokerRepealFlag = C.char(s.BrokerRepealFlag)
	ptr.PlateRepealSerial = C.int(s.PlateRepealSerial)
	go2cStr(s.BankRepealSerial, &ptr.BankRepealSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ptr.FutureRepealSerial = C.int(s.FutureRepealSerial)
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.FutureSerial = C.int(s.FutureSerial)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.VerifyCertNoFlag = C.char(s.VerifyCertNoFlag)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.TradeAmount = C.double(s.TradeAmount)
	ptr.FutureFetchAmount = C.double(s.FutureFetchAmount)
	ptr.FeePayFlag = C.char(s.FeePayFlag)
	ptr.CustFee = C.double(s.CustFee)
	ptr.BrokerFee = C.double(s.BrokerFee)
	go2cStr(s.Message, &ptr.Message[0], C.sizeof_TThostFtdcAddInfoType)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	ptr.BankAccType = C.char(s.BankAccType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ptr.BankSecuAccType = C.char(s.BankSecuAccType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.BankSecuAcc, &ptr.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ptr.BankPwdFlag = C.char(s.BankPwdFlag)
	ptr.SecuPwdFlag = C.char(s.SecuPwdFlag)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.TID = C.int(s.TID)
	ptr.TransferStatus = C.char(s.TransferStatus)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcRspRepealField struct {
	RepealTimeInterval int
	RepealedTimes      int
	BankRepealFlag     byte
	BrokerRepealFlag   byte
	PlateRepealSerial  int
	BankRepealSerial   string
	FutureRepealSerial int
	TradeCode          string
	BankID             string
	BankBranchID       string
	BrokerID           string
	BrokerBranchID     string
	TradeDate          string
	TradeTime          string
	BankSerial         string
	TradingDay         string
	PlateSerial        int
	LastFragment       byte
	SessionID          int
	CustomerName       string
	IdCardType         byte
	IdentifiedCardNo   string
	CustType           byte
	BankAccount        string
	BankPassWord       string
	AccountID          string
	Password           string
	InstallID          int
	FutureSerial       int
	UserID             string
	VerifyCertNoFlag   byte
	CurrencyID         string
	TradeAmount        float64
	FutureFetchAmount  float64
	FeePayFlag         byte
	CustFee            float64
	BrokerFee          float64
	Message            string
	Digest             string
	BankAccType        byte
	DeviceID           string
	BankSecuAccType    byte
	BrokerIDByBank     string
	BankSecuAcc        string
	BankPwdFlag        byte
	SecuPwdFlag        byte
	OperNo             string
	RequestID          int
	TID                int
	TransferStatus     byte
	ErrorID            int
	ErrorMsg           string
	LongCustomerName   string
}

func NewCThostFtdcRspRepealField(p *C.CThostFtdcRspRepealField) *CThostFtdcRspRepealField {
	ret := new(CThostFtdcRspRepealField)
	ret.RepealTimeInterval = int(p.RepealTimeInterval)
	ret.RepealedTimes = int(p.RepealedTimes)
	ret.BankRepealFlag = byte(p.BankRepealFlag)
	ret.BrokerRepealFlag = byte(p.BrokerRepealFlag)
	ret.PlateRepealSerial = int(p.PlateRepealSerial)
	ret.BankRepealSerial = c2goStr(&p.BankRepealSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.FutureRepealSerial = int(p.FutureRepealSerial)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.CustType = byte(p.CustType)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.InstallID = int(p.InstallID)
	ret.FutureSerial = int(p.FutureSerial)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.VerifyCertNoFlag = byte(p.VerifyCertNoFlag)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.TradeAmount = goFloat64(p.TradeAmount)
	ret.FutureFetchAmount = goFloat64(p.FutureFetchAmount)
	ret.FeePayFlag = byte(p.FeePayFlag)
	ret.CustFee = goFloat64(p.CustFee)
	ret.BrokerFee = goFloat64(p.BrokerFee)
	ret.Message = c2goStr(&p.Message[0], C.sizeof_TThostFtdcAddInfoType)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.BankAccType = byte(p.BankAccType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BankSecuAccType = byte(p.BankSecuAccType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.BankSecuAcc = c2goStr(&p.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPwdFlag = byte(p.BankPwdFlag)
	ret.SecuPwdFlag = byte(p.SecuPwdFlag)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.RequestID = int(p.RequestID)
	ret.TID = int(p.TID)
	ret.TransferStatus = byte(p.TransferStatus)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcRspRepealField) CValue() *C.CThostFtdcRspRepealField {
	ptr := C.newCThostFtdcRspRepealField()
	ptr.RepealTimeInterval = C.int(s.RepealTimeInterval)
	ptr.RepealedTimes = C.int(s.RepealedTimes)
	ptr.BankRepealFlag = C.char(s.BankRepealFlag)
	ptr.BrokerRepealFlag = C.char(s.BrokerRepealFlag)
	ptr.PlateRepealSerial = C.int(s.PlateRepealSerial)
	go2cStr(s.BankRepealSerial, &ptr.BankRepealSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ptr.FutureRepealSerial = C.int(s.FutureRepealSerial)
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.FutureSerial = C.int(s.FutureSerial)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.VerifyCertNoFlag = C.char(s.VerifyCertNoFlag)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.TradeAmount = C.double(s.TradeAmount)
	ptr.FutureFetchAmount = C.double(s.FutureFetchAmount)
	ptr.FeePayFlag = C.char(s.FeePayFlag)
	ptr.CustFee = C.double(s.CustFee)
	ptr.BrokerFee = C.double(s.BrokerFee)
	go2cStr(s.Message, &ptr.Message[0], C.sizeof_TThostFtdcAddInfoType)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	ptr.BankAccType = C.char(s.BankAccType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ptr.BankSecuAccType = C.char(s.BankSecuAccType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.BankSecuAcc, &ptr.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ptr.BankPwdFlag = C.char(s.BankPwdFlag)
	ptr.SecuPwdFlag = C.char(s.SecuPwdFlag)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.TID = C.int(s.TID)
	ptr.TransferStatus = C.char(s.TransferStatus)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcReqQueryAccountField struct {
	TradeCode        string
	BankID           string
	BankBranchID     string
	BrokerID         string
	BrokerBranchID   string
	TradeDate        string
	TradeTime        string
	BankSerial       string
	TradingDay       string
	PlateSerial      int
	LastFragment     byte
	SessionID        int
	CustomerName     string
	IdCardType       byte
	IdentifiedCardNo string
	CustType         byte
	BankAccount      string
	BankPassWord     string
	AccountID        string
	Password         string
	FutureSerial     int
	InstallID        int
	UserID           string
	VerifyCertNoFlag byte
	CurrencyID       string
	Digest           string
	BankAccType      byte
	DeviceID         string
	BankSecuAccType  byte
	BrokerIDByBank   string
	BankSecuAcc      string
	BankPwdFlag      byte
	SecuPwdFlag      byte
	OperNo           string
	RequestID        int
	TID              int
	LongCustomerName string
}

func NewCThostFtdcReqQueryAccountField(p *C.CThostFtdcReqQueryAccountField) *CThostFtdcReqQueryAccountField {
	ret := new(CThostFtdcReqQueryAccountField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.CustType = byte(p.CustType)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.FutureSerial = int(p.FutureSerial)
	ret.InstallID = int(p.InstallID)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.VerifyCertNoFlag = byte(p.VerifyCertNoFlag)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.BankAccType = byte(p.BankAccType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BankSecuAccType = byte(p.BankSecuAccType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.BankSecuAcc = c2goStr(&p.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPwdFlag = byte(p.BankPwdFlag)
	ret.SecuPwdFlag = byte(p.SecuPwdFlag)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.RequestID = int(p.RequestID)
	ret.TID = int(p.TID)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcReqQueryAccountField) CValue() *C.CThostFtdcReqQueryAccountField {
	ptr := C.newCThostFtdcReqQueryAccountField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	ptr.FutureSerial = C.int(s.FutureSerial)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.VerifyCertNoFlag = C.char(s.VerifyCertNoFlag)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	ptr.BankAccType = C.char(s.BankAccType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ptr.BankSecuAccType = C.char(s.BankSecuAccType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.BankSecuAcc, &ptr.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ptr.BankPwdFlag = C.char(s.BankPwdFlag)
	ptr.SecuPwdFlag = C.char(s.SecuPwdFlag)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.TID = C.int(s.TID)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcRspQueryAccountField struct {
	TradeCode        string
	BankID           string
	BankBranchID     string
	BrokerID         string
	BrokerBranchID   string
	TradeDate        string
	TradeTime        string
	BankSerial       string
	TradingDay       string
	PlateSerial      int
	LastFragment     byte
	SessionID        int
	CustomerName     string
	IdCardType       byte
	IdentifiedCardNo string
	CustType         byte
	BankAccount      string
	BankPassWord     string
	AccountID        string
	Password         string
	FutureSerial     int
	InstallID        int
	UserID           string
	VerifyCertNoFlag byte
	CurrencyID       string
	Digest           string
	BankAccType      byte
	DeviceID         string
	BankSecuAccType  byte
	BrokerIDByBank   string
	BankSecuAcc      string
	BankPwdFlag      byte
	SecuPwdFlag      byte
	OperNo           string
	RequestID        int
	TID              int
	BankUseAmount    float64
	BankFetchAmount  float64
	LongCustomerName string
}

func NewCThostFtdcRspQueryAccountField(p *C.CThostFtdcRspQueryAccountField) *CThostFtdcRspQueryAccountField {
	ret := new(CThostFtdcRspQueryAccountField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.CustType = byte(p.CustType)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.FutureSerial = int(p.FutureSerial)
	ret.InstallID = int(p.InstallID)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.VerifyCertNoFlag = byte(p.VerifyCertNoFlag)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.BankAccType = byte(p.BankAccType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BankSecuAccType = byte(p.BankSecuAccType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.BankSecuAcc = c2goStr(&p.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPwdFlag = byte(p.BankPwdFlag)
	ret.SecuPwdFlag = byte(p.SecuPwdFlag)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.RequestID = int(p.RequestID)
	ret.TID = int(p.TID)
	ret.BankUseAmount = goFloat64(p.BankUseAmount)
	ret.BankFetchAmount = goFloat64(p.BankFetchAmount)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcRspQueryAccountField) CValue() *C.CThostFtdcRspQueryAccountField {
	ptr := C.newCThostFtdcRspQueryAccountField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	ptr.FutureSerial = C.int(s.FutureSerial)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.VerifyCertNoFlag = C.char(s.VerifyCertNoFlag)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	ptr.BankAccType = C.char(s.BankAccType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ptr.BankSecuAccType = C.char(s.BankSecuAccType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.BankSecuAcc, &ptr.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ptr.BankPwdFlag = C.char(s.BankPwdFlag)
	ptr.SecuPwdFlag = C.char(s.SecuPwdFlag)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.TID = C.int(s.TID)
	ptr.BankUseAmount = C.double(s.BankUseAmount)
	ptr.BankFetchAmount = C.double(s.BankFetchAmount)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcFutureSignIOField struct {
	TradeCode      string
	BankID         string
	BankBranchID   string
	BrokerID       string
	BrokerBranchID string
	TradeDate      string
	TradeTime      string
	BankSerial     string
	TradingDay     string
	PlateSerial    int
	LastFragment   byte
	SessionID      int
	InstallID      int
	UserID         string
	Digest         string
	CurrencyID     string
	DeviceID       string
	BrokerIDByBank string
	OperNo         string
	RequestID      int
	TID            int
}

func NewCThostFtdcFutureSignIOField(p *C.CThostFtdcFutureSignIOField) *CThostFtdcFutureSignIOField {
	ret := new(CThostFtdcFutureSignIOField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.InstallID = int(p.InstallID)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.RequestID = int(p.RequestID)
	ret.TID = int(p.TID)
	return ret
}
func (s *CThostFtdcFutureSignIOField) CValue() *C.CThostFtdcFutureSignIOField {
	ptr := C.newCThostFtdcFutureSignIOField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.TID = C.int(s.TID)
	return ptr
}

type CThostFtdcRspFutureSignInField struct {
	TradeCode      string
	BankID         string
	BankBranchID   string
	BrokerID       string
	BrokerBranchID string
	TradeDate      string
	TradeTime      string
	BankSerial     string
	TradingDay     string
	PlateSerial    int
	LastFragment   byte
	SessionID      int
	InstallID      int
	UserID         string
	Digest         string
	CurrencyID     string
	DeviceID       string
	BrokerIDByBank string
	OperNo         string
	RequestID      int
	TID            int
	ErrorID        int
	ErrorMsg       string
	PinKey         string
	MacKey         string
}

func NewCThostFtdcRspFutureSignInField(p *C.CThostFtdcRspFutureSignInField) *CThostFtdcRspFutureSignInField {
	ret := new(CThostFtdcRspFutureSignInField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.InstallID = int(p.InstallID)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.RequestID = int(p.RequestID)
	ret.TID = int(p.TID)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.PinKey = c2goStr(&p.PinKey[0], C.sizeof_TThostFtdcPasswordKeyType)
	ret.MacKey = c2goStr(&p.MacKey[0], C.sizeof_TThostFtdcPasswordKeyType)
	return ret
}
func (s *CThostFtdcRspFutureSignInField) CValue() *C.CThostFtdcRspFutureSignInField {
	ptr := C.newCThostFtdcRspFutureSignInField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.TID = C.int(s.TID)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.PinKey, &ptr.PinKey[0], C.sizeof_TThostFtdcPasswordKeyType)
	go2cStr(s.MacKey, &ptr.MacKey[0], C.sizeof_TThostFtdcPasswordKeyType)
	return ptr
}

type CThostFtdcReqFutureSignOutField struct {
	TradeCode      string
	BankID         string
	BankBranchID   string
	BrokerID       string
	BrokerBranchID string
	TradeDate      string
	TradeTime      string
	BankSerial     string
	TradingDay     string
	PlateSerial    int
	LastFragment   byte
	SessionID      int
	InstallID      int
	UserID         string
	Digest         string
	CurrencyID     string
	DeviceID       string
	BrokerIDByBank string
	OperNo         string
	RequestID      int
	TID            int
}

func NewCThostFtdcReqFutureSignOutField(p *C.CThostFtdcReqFutureSignOutField) *CThostFtdcReqFutureSignOutField {
	ret := new(CThostFtdcReqFutureSignOutField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.InstallID = int(p.InstallID)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.RequestID = int(p.RequestID)
	ret.TID = int(p.TID)
	return ret
}
func (s *CThostFtdcReqFutureSignOutField) CValue() *C.CThostFtdcReqFutureSignOutField {
	ptr := C.newCThostFtdcReqFutureSignOutField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.TID = C.int(s.TID)
	return ptr
}

type CThostFtdcRspFutureSignOutField struct {
	TradeCode      string
	BankID         string
	BankBranchID   string
	BrokerID       string
	BrokerBranchID string
	TradeDate      string
	TradeTime      string
	BankSerial     string
	TradingDay     string
	PlateSerial    int
	LastFragment   byte
	SessionID      int
	InstallID      int
	UserID         string
	Digest         string
	CurrencyID     string
	DeviceID       string
	BrokerIDByBank string
	OperNo         string
	RequestID      int
	TID            int
	ErrorID        int
	ErrorMsg       string
}

func NewCThostFtdcRspFutureSignOutField(p *C.CThostFtdcRspFutureSignOutField) *CThostFtdcRspFutureSignOutField {
	ret := new(CThostFtdcRspFutureSignOutField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.InstallID = int(p.InstallID)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.RequestID = int(p.RequestID)
	ret.TID = int(p.TID)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ret
}
func (s *CThostFtdcRspFutureSignOutField) CValue() *C.CThostFtdcRspFutureSignOutField {
	ptr := C.newCThostFtdcRspFutureSignOutField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.TID = C.int(s.TID)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ptr
}

type CThostFtdcReqQueryTradeResultBySerialField struct {
	TradeCode          string
	BankID             string
	BankBranchID       string
	BrokerID           string
	BrokerBranchID     string
	TradeDate          string
	TradeTime          string
	BankSerial         string
	TradingDay         string
	PlateSerial        int
	LastFragment       byte
	SessionID          int
	Reference          int
	RefrenceIssureType byte
	RefrenceIssure     string
	CustomerName       string
	IdCardType         byte
	IdentifiedCardNo   string
	CustType           byte
	BankAccount        string
	BankPassWord       string
	AccountID          string
	Password           string
	CurrencyID         string
	TradeAmount        float64
	Digest             string
	LongCustomerName   string
}

func NewCThostFtdcReqQueryTradeResultBySerialField(p *C.CThostFtdcReqQueryTradeResultBySerialField) *CThostFtdcReqQueryTradeResultBySerialField {
	ret := new(CThostFtdcReqQueryTradeResultBySerialField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.Reference = int(p.Reference)
	ret.RefrenceIssureType = byte(p.RefrenceIssureType)
	ret.RefrenceIssure = c2goStr(&p.RefrenceIssure[0], C.sizeof_TThostFtdcOrganCodeType)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.CustType = byte(p.CustType)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.TradeAmount = goFloat64(p.TradeAmount)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcReqQueryTradeResultBySerialField) CValue() *C.CThostFtdcReqQueryTradeResultBySerialField {
	ptr := C.newCThostFtdcReqQueryTradeResultBySerialField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	ptr.Reference = C.int(s.Reference)
	ptr.RefrenceIssureType = C.char(s.RefrenceIssureType)
	go2cStr(s.RefrenceIssure, &ptr.RefrenceIssure[0], C.sizeof_TThostFtdcOrganCodeType)
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.TradeAmount = C.double(s.TradeAmount)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcRspQueryTradeResultBySerialField struct {
	TradeCode                    string
	BankID                       string
	BankBranchID                 string
	BrokerID                     string
	BrokerBranchID               string
	TradeDate                    string
	TradeTime                    string
	BankSerial                   string
	TradingDay                   string
	PlateSerial                  int
	LastFragment                 byte
	SessionID                    int
	ErrorID                      int
	ErrorMsg                     string
	Reference                    int
	RefrenceIssureType           byte
	RefrenceIssure               string
	OriginReturnCode             string
	OriginDescrInfoForReturnCode string
	BankAccount                  string
	BankPassWord                 string
	AccountID                    string
	Password                     string
	CurrencyID                   string
	TradeAmount                  float64
	Digest                       string
}

func NewCThostFtdcRspQueryTradeResultBySerialField(p *C.CThostFtdcRspQueryTradeResultBySerialField) *CThostFtdcRspQueryTradeResultBySerialField {
	ret := new(CThostFtdcRspQueryTradeResultBySerialField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.Reference = int(p.Reference)
	ret.RefrenceIssureType = byte(p.RefrenceIssureType)
	ret.RefrenceIssure = c2goStr(&p.RefrenceIssure[0], C.sizeof_TThostFtdcOrganCodeType)
	ret.OriginReturnCode = c2goStr(&p.OriginReturnCode[0], C.sizeof_TThostFtdcReturnCodeType)
	ret.OriginDescrInfoForReturnCode = c2goStr(&p.OriginDescrInfoForReturnCode[0], C.sizeof_TThostFtdcDescrInfoForReturnCodeType)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.TradeAmount = goFloat64(p.TradeAmount)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	return ret
}
func (s *CThostFtdcRspQueryTradeResultBySerialField) CValue() *C.CThostFtdcRspQueryTradeResultBySerialField {
	ptr := C.newCThostFtdcRspQueryTradeResultBySerialField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ptr.Reference = C.int(s.Reference)
	ptr.RefrenceIssureType = C.char(s.RefrenceIssureType)
	go2cStr(s.RefrenceIssure, &ptr.RefrenceIssure[0], C.sizeof_TThostFtdcOrganCodeType)
	go2cStr(s.OriginReturnCode, &ptr.OriginReturnCode[0], C.sizeof_TThostFtdcReturnCodeType)
	go2cStr(s.OriginDescrInfoForReturnCode, &ptr.OriginDescrInfoForReturnCode[0], C.sizeof_TThostFtdcDescrInfoForReturnCodeType)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.TradeAmount = C.double(s.TradeAmount)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	return ptr
}

type CThostFtdcReqDayEndFileReadyField struct {
	TradeCode        string
	BankID           string
	BankBranchID     string
	BrokerID         string
	BrokerBranchID   string
	TradeDate        string
	TradeTime        string
	BankSerial       string
	TradingDay       string
	PlateSerial      int
	LastFragment     byte
	SessionID        int
	FileBusinessCode byte
	Digest           string
}

func NewCThostFtdcReqDayEndFileReadyField(p *C.CThostFtdcReqDayEndFileReadyField) *CThostFtdcReqDayEndFileReadyField {
	ret := new(CThostFtdcReqDayEndFileReadyField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.FileBusinessCode = byte(p.FileBusinessCode)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	return ret
}
func (s *CThostFtdcReqDayEndFileReadyField) CValue() *C.CThostFtdcReqDayEndFileReadyField {
	ptr := C.newCThostFtdcReqDayEndFileReadyField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	ptr.FileBusinessCode = C.char(s.FileBusinessCode)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	return ptr
}

type CThostFtdcReturnResultField struct {
	ReturnCode             string
	DescrInfoForReturnCode string
}

func NewCThostFtdcReturnResultField(p *C.CThostFtdcReturnResultField) *CThostFtdcReturnResultField {
	ret := new(CThostFtdcReturnResultField)
	ret.ReturnCode = c2goStr(&p.ReturnCode[0], C.sizeof_TThostFtdcReturnCodeType)
	ret.DescrInfoForReturnCode = c2goStr(&p.DescrInfoForReturnCode[0], C.sizeof_TThostFtdcDescrInfoForReturnCodeType)
	return ret
}
func (s *CThostFtdcReturnResultField) CValue() *C.CThostFtdcReturnResultField {
	ptr := C.newCThostFtdcReturnResultField()
	go2cStr(s.ReturnCode, &ptr.ReturnCode[0], C.sizeof_TThostFtdcReturnCodeType)
	go2cStr(s.DescrInfoForReturnCode, &ptr.DescrInfoForReturnCode[0], C.sizeof_TThostFtdcDescrInfoForReturnCodeType)
	return ptr
}

type CThostFtdcVerifyFuturePasswordField struct {
	TradeCode      string
	BankID         string
	BankBranchID   string
	BrokerID       string
	BrokerBranchID string
	TradeDate      string
	TradeTime      string
	BankSerial     string
	TradingDay     string
	PlateSerial    int
	LastFragment   byte
	SessionID      int
	AccountID      string
	Password       string
	BankAccount    string
	BankPassWord   string
	InstallID      int
	TID            int
	CurrencyID     string
}

func NewCThostFtdcVerifyFuturePasswordField(p *C.CThostFtdcVerifyFuturePasswordField) *CThostFtdcVerifyFuturePasswordField {
	ret := new(CThostFtdcVerifyFuturePasswordField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.InstallID = int(p.InstallID)
	ret.TID = int(p.TID)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ret
}
func (s *CThostFtdcVerifyFuturePasswordField) CValue() *C.CThostFtdcVerifyFuturePasswordField {
	ptr := C.newCThostFtdcVerifyFuturePasswordField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.TID = C.int(s.TID)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ptr
}

type CThostFtdcVerifyCustInfoField struct {
	CustomerName     string
	IdCardType       byte
	IdentifiedCardNo string
	CustType         byte
	LongCustomerName string
}

func NewCThostFtdcVerifyCustInfoField(p *C.CThostFtdcVerifyCustInfoField) *CThostFtdcVerifyCustInfoField {
	ret := new(CThostFtdcVerifyCustInfoField)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.CustType = byte(p.CustType)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcVerifyCustInfoField) CValue() *C.CThostFtdcVerifyCustInfoField {
	ptr := C.newCThostFtdcVerifyCustInfoField()
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcVerifyFuturePasswordAndCustInfoField struct {
	CustomerName     string
	IdCardType       byte
	IdentifiedCardNo string
	CustType         byte
	AccountID        string
	Password         string
	CurrencyID       string
	LongCustomerName string
}

func NewCThostFtdcVerifyFuturePasswordAndCustInfoField(p *C.CThostFtdcVerifyFuturePasswordAndCustInfoField) *CThostFtdcVerifyFuturePasswordAndCustInfoField {
	ret := new(CThostFtdcVerifyFuturePasswordAndCustInfoField)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.CustType = byte(p.CustType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcVerifyFuturePasswordAndCustInfoField) CValue() *C.CThostFtdcVerifyFuturePasswordAndCustInfoField {
	ptr := C.newCThostFtdcVerifyFuturePasswordAndCustInfoField()
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcDepositResultInformField struct {
	DepositSeqNo           string
	BrokerID               string
	InvestorID             string
	Deposit                float64
	RequestID              int
	ReturnCode             string
	DescrInfoForReturnCode string
}

func NewCThostFtdcDepositResultInformField(p *C.CThostFtdcDepositResultInformField) *CThostFtdcDepositResultInformField {
	ret := new(CThostFtdcDepositResultInformField)
	ret.DepositSeqNo = c2goStr(&p.DepositSeqNo[0], C.sizeof_TThostFtdcDepositSeqNoType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.Deposit = goFloat64(p.Deposit)
	ret.RequestID = int(p.RequestID)
	ret.ReturnCode = c2goStr(&p.ReturnCode[0], C.sizeof_TThostFtdcReturnCodeType)
	ret.DescrInfoForReturnCode = c2goStr(&p.DescrInfoForReturnCode[0], C.sizeof_TThostFtdcDescrInfoForReturnCodeType)
	return ret
}
func (s *CThostFtdcDepositResultInformField) CValue() *C.CThostFtdcDepositResultInformField {
	ptr := C.newCThostFtdcDepositResultInformField()
	go2cStr(s.DepositSeqNo, &ptr.DepositSeqNo[0], C.sizeof_TThostFtdcDepositSeqNoType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.Deposit = C.double(s.Deposit)
	ptr.RequestID = C.int(s.RequestID)
	go2cStr(s.ReturnCode, &ptr.ReturnCode[0], C.sizeof_TThostFtdcReturnCodeType)
	go2cStr(s.DescrInfoForReturnCode, &ptr.DescrInfoForReturnCode[0], C.sizeof_TThostFtdcDescrInfoForReturnCodeType)
	return ptr
}

type CThostFtdcReqSyncKeyField struct {
	TradeCode      string
	BankID         string
	BankBranchID   string
	BrokerID       string
	BrokerBranchID string
	TradeDate      string
	TradeTime      string
	BankSerial     string
	TradingDay     string
	PlateSerial    int
	LastFragment   byte
	SessionID      int
	InstallID      int
	UserID         string
	Message        string
	DeviceID       string
	BrokerIDByBank string
	OperNo         string
	RequestID      int
	TID            int
}

func NewCThostFtdcReqSyncKeyField(p *C.CThostFtdcReqSyncKeyField) *CThostFtdcReqSyncKeyField {
	ret := new(CThostFtdcReqSyncKeyField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.InstallID = int(p.InstallID)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Message = c2goStr(&p.Message[0], C.sizeof_TThostFtdcAddInfoType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.RequestID = int(p.RequestID)
	ret.TID = int(p.TID)
	return ret
}
func (s *CThostFtdcReqSyncKeyField) CValue() *C.CThostFtdcReqSyncKeyField {
	ptr := C.newCThostFtdcReqSyncKeyField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Message, &ptr.Message[0], C.sizeof_TThostFtdcAddInfoType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.TID = C.int(s.TID)
	return ptr
}

type CThostFtdcRspSyncKeyField struct {
	TradeCode      string
	BankID         string
	BankBranchID   string
	BrokerID       string
	BrokerBranchID string
	TradeDate      string
	TradeTime      string
	BankSerial     string
	TradingDay     string
	PlateSerial    int
	LastFragment   byte
	SessionID      int
	InstallID      int
	UserID         string
	Message        string
	DeviceID       string
	BrokerIDByBank string
	OperNo         string
	RequestID      int
	TID            int
	ErrorID        int
	ErrorMsg       string
}

func NewCThostFtdcRspSyncKeyField(p *C.CThostFtdcRspSyncKeyField) *CThostFtdcRspSyncKeyField {
	ret := new(CThostFtdcRspSyncKeyField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.InstallID = int(p.InstallID)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Message = c2goStr(&p.Message[0], C.sizeof_TThostFtdcAddInfoType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.RequestID = int(p.RequestID)
	ret.TID = int(p.TID)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ret
}
func (s *CThostFtdcRspSyncKeyField) CValue() *C.CThostFtdcRspSyncKeyField {
	ptr := C.newCThostFtdcRspSyncKeyField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Message, &ptr.Message[0], C.sizeof_TThostFtdcAddInfoType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.TID = C.int(s.TID)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ptr
}

type CThostFtdcNotifyQueryAccountField struct {
	TradeCode        string
	BankID           string
	BankBranchID     string
	BrokerID         string
	BrokerBranchID   string
	TradeDate        string
	TradeTime        string
	BankSerial       string
	TradingDay       string
	PlateSerial      int
	LastFragment     byte
	SessionID        int
	CustomerName     string
	IdCardType       byte
	IdentifiedCardNo string
	CustType         byte
	BankAccount      string
	BankPassWord     string
	AccountID        string
	Password         string
	FutureSerial     int
	InstallID        int
	UserID           string
	VerifyCertNoFlag byte
	CurrencyID       string
	Digest           string
	BankAccType      byte
	DeviceID         string
	BankSecuAccType  byte
	BrokerIDByBank   string
	BankSecuAcc      string
	BankPwdFlag      byte
	SecuPwdFlag      byte
	OperNo           string
	RequestID        int
	TID              int
	BankUseAmount    float64
	BankFetchAmount  float64
	ErrorID          int
	ErrorMsg         string
	LongCustomerName string
}

func NewCThostFtdcNotifyQueryAccountField(p *C.CThostFtdcNotifyQueryAccountField) *CThostFtdcNotifyQueryAccountField {
	ret := new(CThostFtdcNotifyQueryAccountField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.CustType = byte(p.CustType)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.FutureSerial = int(p.FutureSerial)
	ret.InstallID = int(p.InstallID)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.VerifyCertNoFlag = byte(p.VerifyCertNoFlag)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.BankAccType = byte(p.BankAccType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BankSecuAccType = byte(p.BankSecuAccType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.BankSecuAcc = c2goStr(&p.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPwdFlag = byte(p.BankPwdFlag)
	ret.SecuPwdFlag = byte(p.SecuPwdFlag)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.RequestID = int(p.RequestID)
	ret.TID = int(p.TID)
	ret.BankUseAmount = goFloat64(p.BankUseAmount)
	ret.BankFetchAmount = goFloat64(p.BankFetchAmount)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcNotifyQueryAccountField) CValue() *C.CThostFtdcNotifyQueryAccountField {
	ptr := C.newCThostFtdcNotifyQueryAccountField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	ptr.FutureSerial = C.int(s.FutureSerial)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.VerifyCertNoFlag = C.char(s.VerifyCertNoFlag)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	ptr.BankAccType = C.char(s.BankAccType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ptr.BankSecuAccType = C.char(s.BankSecuAccType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.BankSecuAcc, &ptr.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ptr.BankPwdFlag = C.char(s.BankPwdFlag)
	ptr.SecuPwdFlag = C.char(s.SecuPwdFlag)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.TID = C.int(s.TID)
	ptr.BankUseAmount = C.double(s.BankUseAmount)
	ptr.BankFetchAmount = C.double(s.BankFetchAmount)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcTransferSerialField struct {
	PlateSerial      int
	TradeDate        string
	TradingDay       string
	TradeTime        string
	TradeCode        string
	SessionID        int
	BankID           string
	BankBranchID     string
	BankAccType      byte
	BankAccount      string
	BankSerial       string
	BrokerID         string
	BrokerBranchID   string
	FutureAccType    byte
	AccountID        string
	InvestorID       string
	FutureSerial     int
	IdCardType       byte
	IdentifiedCardNo string
	CurrencyID       string
	TradeAmount      float64
	CustFee          float64
	BrokerFee        float64
	AvailabilityFlag byte
	OperatorCode     string
	BankNewAccount   string
	ErrorID          int
	ErrorMsg         string
}

func NewCThostFtdcTransferSerialField(p *C.CThostFtdcTransferSerialField) *CThostFtdcTransferSerialField {
	ret := new(CThostFtdcTransferSerialField)
	ret.PlateSerial = int(p.PlateSerial)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.SessionID = int(p.SessionID)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BankAccType = byte(p.BankAccType)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.FutureAccType = byte(p.FutureAccType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.FutureSerial = int(p.FutureSerial)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.TradeAmount = goFloat64(p.TradeAmount)
	ret.CustFee = goFloat64(p.CustFee)
	ret.BrokerFee = goFloat64(p.BrokerFee)
	ret.AvailabilityFlag = byte(p.AvailabilityFlag)
	ret.OperatorCode = c2goStr(&p.OperatorCode[0], C.sizeof_TThostFtdcOperatorCodeType)
	ret.BankNewAccount = c2goStr(&p.BankNewAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ret
}
func (s *CThostFtdcTransferSerialField) CValue() *C.CThostFtdcTransferSerialField {
	ptr := C.newCThostFtdcTransferSerialField()
	ptr.PlateSerial = C.int(s.PlateSerial)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ptr.BankAccType = C.char(s.BankAccType)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ptr.FutureAccType = C.char(s.FutureAccType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.FutureSerial = C.int(s.FutureSerial)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.TradeAmount = C.double(s.TradeAmount)
	ptr.CustFee = C.double(s.CustFee)
	ptr.BrokerFee = C.double(s.BrokerFee)
	ptr.AvailabilityFlag = C.char(s.AvailabilityFlag)
	go2cStr(s.OperatorCode, &ptr.OperatorCode[0], C.sizeof_TThostFtdcOperatorCodeType)
	go2cStr(s.BankNewAccount, &ptr.BankNewAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ptr
}

type CThostFtdcQryTransferSerialField struct {
	BrokerID   string
	AccountID  string
	BankID     string
	CurrencyID string
}

func NewCThostFtdcQryTransferSerialField(p *C.CThostFtdcQryTransferSerialField) *CThostFtdcQryTransferSerialField {
	ret := new(CThostFtdcQryTransferSerialField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ret
}
func (s *CThostFtdcQryTransferSerialField) CValue() *C.CThostFtdcQryTransferSerialField {
	ptr := C.newCThostFtdcQryTransferSerialField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ptr
}

type CThostFtdcNotifyFutureSignInField struct {
	TradeCode      string
	BankID         string
	BankBranchID   string
	BrokerID       string
	BrokerBranchID string
	TradeDate      string
	TradeTime      string
	BankSerial     string
	TradingDay     string
	PlateSerial    int
	LastFragment   byte
	SessionID      int
	InstallID      int
	UserID         string
	Digest         string
	CurrencyID     string
	DeviceID       string
	BrokerIDByBank string
	OperNo         string
	RequestID      int
	TID            int
	ErrorID        int
	ErrorMsg       string
	PinKey         string
	MacKey         string
}

func NewCThostFtdcNotifyFutureSignInField(p *C.CThostFtdcNotifyFutureSignInField) *CThostFtdcNotifyFutureSignInField {
	ret := new(CThostFtdcNotifyFutureSignInField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.InstallID = int(p.InstallID)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.RequestID = int(p.RequestID)
	ret.TID = int(p.TID)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.PinKey = c2goStr(&p.PinKey[0], C.sizeof_TThostFtdcPasswordKeyType)
	ret.MacKey = c2goStr(&p.MacKey[0], C.sizeof_TThostFtdcPasswordKeyType)
	return ret
}
func (s *CThostFtdcNotifyFutureSignInField) CValue() *C.CThostFtdcNotifyFutureSignInField {
	ptr := C.newCThostFtdcNotifyFutureSignInField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.TID = C.int(s.TID)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.PinKey, &ptr.PinKey[0], C.sizeof_TThostFtdcPasswordKeyType)
	go2cStr(s.MacKey, &ptr.MacKey[0], C.sizeof_TThostFtdcPasswordKeyType)
	return ptr
}

type CThostFtdcNotifyFutureSignOutField struct {
	TradeCode      string
	BankID         string
	BankBranchID   string
	BrokerID       string
	BrokerBranchID string
	TradeDate      string
	TradeTime      string
	BankSerial     string
	TradingDay     string
	PlateSerial    int
	LastFragment   byte
	SessionID      int
	InstallID      int
	UserID         string
	Digest         string
	CurrencyID     string
	DeviceID       string
	BrokerIDByBank string
	OperNo         string
	RequestID      int
	TID            int
	ErrorID        int
	ErrorMsg       string
}

func NewCThostFtdcNotifyFutureSignOutField(p *C.CThostFtdcNotifyFutureSignOutField) *CThostFtdcNotifyFutureSignOutField {
	ret := new(CThostFtdcNotifyFutureSignOutField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.InstallID = int(p.InstallID)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.RequestID = int(p.RequestID)
	ret.TID = int(p.TID)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ret
}
func (s *CThostFtdcNotifyFutureSignOutField) CValue() *C.CThostFtdcNotifyFutureSignOutField {
	ptr := C.newCThostFtdcNotifyFutureSignOutField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.TID = C.int(s.TID)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ptr
}

type CThostFtdcNotifySyncKeyField struct {
	TradeCode      string
	BankID         string
	BankBranchID   string
	BrokerID       string
	BrokerBranchID string
	TradeDate      string
	TradeTime      string
	BankSerial     string
	TradingDay     string
	PlateSerial    int
	LastFragment   byte
	SessionID      int
	InstallID      int
	UserID         string
	Message        string
	DeviceID       string
	BrokerIDByBank string
	OperNo         string
	RequestID      int
	TID            int
	ErrorID        int
	ErrorMsg       string
}

func NewCThostFtdcNotifySyncKeyField(p *C.CThostFtdcNotifySyncKeyField) *CThostFtdcNotifySyncKeyField {
	ret := new(CThostFtdcNotifySyncKeyField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.InstallID = int(p.InstallID)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Message = c2goStr(&p.Message[0], C.sizeof_TThostFtdcAddInfoType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.RequestID = int(p.RequestID)
	ret.TID = int(p.TID)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ret
}
func (s *CThostFtdcNotifySyncKeyField) CValue() *C.CThostFtdcNotifySyncKeyField {
	ptr := C.newCThostFtdcNotifySyncKeyField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	ptr.InstallID = C.int(s.InstallID)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Message, &ptr.Message[0], C.sizeof_TThostFtdcAddInfoType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.RequestID = C.int(s.RequestID)
	ptr.TID = C.int(s.TID)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ptr
}

type CThostFtdcQryAccountregisterField struct {
	BrokerID     string
	AccountID    string
	BankID       string
	BankBranchID string
	CurrencyID   string
}

func NewCThostFtdcQryAccountregisterField(p *C.CThostFtdcQryAccountregisterField) *CThostFtdcQryAccountregisterField {
	ret := new(CThostFtdcQryAccountregisterField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ret
}
func (s *CThostFtdcQryAccountregisterField) CValue() *C.CThostFtdcQryAccountregisterField {
	ptr := C.newCThostFtdcQryAccountregisterField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ptr
}

type CThostFtdcAccountregisterField struct {
	TradeDay         string
	BankID           string
	BankBranchID     string
	BankAccount      string
	BrokerID         string
	BrokerBranchID   string
	AccountID        string
	IdCardType       byte
	IdentifiedCardNo string
	CustomerName     string
	CurrencyID       string
	OpenOrDestroy    byte
	RegDate          string
	OutDate          string
	TID              int
	CustType         byte
	BankAccType      byte
	LongCustomerName string
}

func NewCThostFtdcAccountregisterField(p *C.CThostFtdcAccountregisterField) *CThostFtdcAccountregisterField {
	ret := new(CThostFtdcAccountregisterField)
	ret.TradeDay = c2goStr(&p.TradeDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.OpenOrDestroy = byte(p.OpenOrDestroy)
	ret.RegDate = c2goStr(&p.RegDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.OutDate = c2goStr(&p.OutDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TID = int(p.TID)
	ret.CustType = byte(p.CustType)
	ret.BankAccType = byte(p.BankAccType)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcAccountregisterField) CValue() *C.CThostFtdcAccountregisterField {
	ptr := C.newCThostFtdcAccountregisterField()
	go2cStr(s.TradeDay, &ptr.TradeDay[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.OpenOrDestroy = C.char(s.OpenOrDestroy)
	go2cStr(s.RegDate, &ptr.RegDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.OutDate, &ptr.OutDate[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.TID = C.int(s.TID)
	ptr.CustType = C.char(s.CustType)
	ptr.BankAccType = C.char(s.BankAccType)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcOpenAccountField struct {
	TradeCode          string
	BankID             string
	BankBranchID       string
	BrokerID           string
	BrokerBranchID     string
	TradeDate          string
	TradeTime          string
	BankSerial         string
	TradingDay         string
	PlateSerial        int
	LastFragment       byte
	SessionID          int
	CustomerName       string
	IdCardType         byte
	IdentifiedCardNo   string
	Gender             byte
	CountryCode        string
	CustType           byte
	Address            string
	ZipCode            string
	Telephone          string
	MobilePhone        string
	Fax                string
	EMail              string
	MoneyAccountStatus byte
	BankAccount        string
	BankPassWord       string
	AccountID          string
	Password           string
	InstallID          int
	VerifyCertNoFlag   byte
	CurrencyID         string
	CashExchangeCode   byte
	Digest             string
	BankAccType        byte
	DeviceID           string
	BankSecuAccType    byte
	BrokerIDByBank     string
	BankSecuAcc        string
	BankPwdFlag        byte
	SecuPwdFlag        byte
	OperNo             string
	TID                int
	UserID             string
	ErrorID            int
	ErrorMsg           string
	LongCustomerName   string
}

func NewCThostFtdcOpenAccountField(p *C.CThostFtdcOpenAccountField) *CThostFtdcOpenAccountField {
	ret := new(CThostFtdcOpenAccountField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.Gender = byte(p.Gender)
	ret.CountryCode = c2goStr(&p.CountryCode[0], C.sizeof_TThostFtdcCountryCodeType)
	ret.CustType = byte(p.CustType)
	ret.Address = c2goStr(&p.Address[0], C.sizeof_TThostFtdcAddressType)
	ret.ZipCode = c2goStr(&p.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	ret.Telephone = c2goStr(&p.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	ret.MobilePhone = c2goStr(&p.MobilePhone[0], C.sizeof_TThostFtdcMobilePhoneType)
	ret.Fax = c2goStr(&p.Fax[0], C.sizeof_TThostFtdcFaxType)
	ret.EMail = c2goStr(&p.EMail[0], C.sizeof_TThostFtdcEMailType)
	ret.MoneyAccountStatus = byte(p.MoneyAccountStatus)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.InstallID = int(p.InstallID)
	ret.VerifyCertNoFlag = byte(p.VerifyCertNoFlag)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.CashExchangeCode = byte(p.CashExchangeCode)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.BankAccType = byte(p.BankAccType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BankSecuAccType = byte(p.BankSecuAccType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.BankSecuAcc = c2goStr(&p.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPwdFlag = byte(p.BankPwdFlag)
	ret.SecuPwdFlag = byte(p.SecuPwdFlag)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.TID = int(p.TID)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcOpenAccountField) CValue() *C.CThostFtdcOpenAccountField {
	ptr := C.newCThostFtdcOpenAccountField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.Gender = C.char(s.Gender)
	go2cStr(s.CountryCode, &ptr.CountryCode[0], C.sizeof_TThostFtdcCountryCodeType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.Address, &ptr.Address[0], C.sizeof_TThostFtdcAddressType)
	go2cStr(s.ZipCode, &ptr.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	go2cStr(s.Telephone, &ptr.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	go2cStr(s.MobilePhone, &ptr.MobilePhone[0], C.sizeof_TThostFtdcMobilePhoneType)
	go2cStr(s.Fax, &ptr.Fax[0], C.sizeof_TThostFtdcFaxType)
	go2cStr(s.EMail, &ptr.EMail[0], C.sizeof_TThostFtdcEMailType)
	ptr.MoneyAccountStatus = C.char(s.MoneyAccountStatus)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.VerifyCertNoFlag = C.char(s.VerifyCertNoFlag)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.CashExchangeCode = C.char(s.CashExchangeCode)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	ptr.BankAccType = C.char(s.BankAccType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ptr.BankSecuAccType = C.char(s.BankSecuAccType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.BankSecuAcc, &ptr.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ptr.BankPwdFlag = C.char(s.BankPwdFlag)
	ptr.SecuPwdFlag = C.char(s.SecuPwdFlag)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.TID = C.int(s.TID)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcCancelAccountField struct {
	TradeCode          string
	BankID             string
	BankBranchID       string
	BrokerID           string
	BrokerBranchID     string
	TradeDate          string
	TradeTime          string
	BankSerial         string
	TradingDay         string
	PlateSerial        int
	LastFragment       byte
	SessionID          int
	CustomerName       string
	IdCardType         byte
	IdentifiedCardNo   string
	Gender             byte
	CountryCode        string
	CustType           byte
	Address            string
	ZipCode            string
	Telephone          string
	MobilePhone        string
	Fax                string
	EMail              string
	MoneyAccountStatus byte
	BankAccount        string
	BankPassWord       string
	AccountID          string
	Password           string
	InstallID          int
	VerifyCertNoFlag   byte
	CurrencyID         string
	CashExchangeCode   byte
	Digest             string
	BankAccType        byte
	DeviceID           string
	BankSecuAccType    byte
	BrokerIDByBank     string
	BankSecuAcc        string
	BankPwdFlag        byte
	SecuPwdFlag        byte
	OperNo             string
	TID                int
	UserID             string
	ErrorID            int
	ErrorMsg           string
	LongCustomerName   string
}

func NewCThostFtdcCancelAccountField(p *C.CThostFtdcCancelAccountField) *CThostFtdcCancelAccountField {
	ret := new(CThostFtdcCancelAccountField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.Gender = byte(p.Gender)
	ret.CountryCode = c2goStr(&p.CountryCode[0], C.sizeof_TThostFtdcCountryCodeType)
	ret.CustType = byte(p.CustType)
	ret.Address = c2goStr(&p.Address[0], C.sizeof_TThostFtdcAddressType)
	ret.ZipCode = c2goStr(&p.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	ret.Telephone = c2goStr(&p.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	ret.MobilePhone = c2goStr(&p.MobilePhone[0], C.sizeof_TThostFtdcMobilePhoneType)
	ret.Fax = c2goStr(&p.Fax[0], C.sizeof_TThostFtdcFaxType)
	ret.EMail = c2goStr(&p.EMail[0], C.sizeof_TThostFtdcEMailType)
	ret.MoneyAccountStatus = byte(p.MoneyAccountStatus)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.InstallID = int(p.InstallID)
	ret.VerifyCertNoFlag = byte(p.VerifyCertNoFlag)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.CashExchangeCode = byte(p.CashExchangeCode)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.BankAccType = byte(p.BankAccType)
	ret.DeviceID = c2goStr(&p.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ret.BankSecuAccType = byte(p.BankSecuAccType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.BankSecuAcc = c2goStr(&p.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPwdFlag = byte(p.BankPwdFlag)
	ret.SecuPwdFlag = byte(p.SecuPwdFlag)
	ret.OperNo = c2goStr(&p.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ret.TID = int(p.TID)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcCancelAccountField) CValue() *C.CThostFtdcCancelAccountField {
	ptr := C.newCThostFtdcCancelAccountField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.Gender = C.char(s.Gender)
	go2cStr(s.CountryCode, &ptr.CountryCode[0], C.sizeof_TThostFtdcCountryCodeType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.Address, &ptr.Address[0], C.sizeof_TThostFtdcAddressType)
	go2cStr(s.ZipCode, &ptr.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	go2cStr(s.Telephone, &ptr.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	go2cStr(s.MobilePhone, &ptr.MobilePhone[0], C.sizeof_TThostFtdcMobilePhoneType)
	go2cStr(s.Fax, &ptr.Fax[0], C.sizeof_TThostFtdcFaxType)
	go2cStr(s.EMail, &ptr.EMail[0], C.sizeof_TThostFtdcEMailType)
	ptr.MoneyAccountStatus = C.char(s.MoneyAccountStatus)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.VerifyCertNoFlag = C.char(s.VerifyCertNoFlag)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.CashExchangeCode = C.char(s.CashExchangeCode)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	ptr.BankAccType = C.char(s.BankAccType)
	go2cStr(s.DeviceID, &ptr.DeviceID[0], C.sizeof_TThostFtdcDeviceIDType)
	ptr.BankSecuAccType = C.char(s.BankSecuAccType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	go2cStr(s.BankSecuAcc, &ptr.BankSecuAcc[0], C.sizeof_TThostFtdcBankAccountType)
	ptr.BankPwdFlag = C.char(s.BankPwdFlag)
	ptr.SecuPwdFlag = C.char(s.SecuPwdFlag)
	go2cStr(s.OperNo, &ptr.OperNo[0], C.sizeof_TThostFtdcOperNoType)
	ptr.TID = C.int(s.TID)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcChangeAccountField struct {
	TradeCode          string
	BankID             string
	BankBranchID       string
	BrokerID           string
	BrokerBranchID     string
	TradeDate          string
	TradeTime          string
	BankSerial         string
	TradingDay         string
	PlateSerial        int
	LastFragment       byte
	SessionID          int
	CustomerName       string
	IdCardType         byte
	IdentifiedCardNo   string
	Gender             byte
	CountryCode        string
	CustType           byte
	Address            string
	ZipCode            string
	Telephone          string
	MobilePhone        string
	Fax                string
	EMail              string
	MoneyAccountStatus byte
	BankAccount        string
	BankPassWord       string
	NewBankAccount     string
	NewBankPassWord    string
	AccountID          string
	Password           string
	BankAccType        byte
	InstallID          int
	VerifyCertNoFlag   byte
	CurrencyID         string
	BrokerIDByBank     string
	BankPwdFlag        byte
	SecuPwdFlag        byte
	TID                int
	Digest             string
	ErrorID            int
	ErrorMsg           string
	LongCustomerName   string
}

func NewCThostFtdcChangeAccountField(p *C.CThostFtdcChangeAccountField) *CThostFtdcChangeAccountField {
	ret := new(CThostFtdcChangeAccountField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.Gender = byte(p.Gender)
	ret.CountryCode = c2goStr(&p.CountryCode[0], C.sizeof_TThostFtdcCountryCodeType)
	ret.CustType = byte(p.CustType)
	ret.Address = c2goStr(&p.Address[0], C.sizeof_TThostFtdcAddressType)
	ret.ZipCode = c2goStr(&p.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	ret.Telephone = c2goStr(&p.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	ret.MobilePhone = c2goStr(&p.MobilePhone[0], C.sizeof_TThostFtdcMobilePhoneType)
	ret.Fax = c2goStr(&p.Fax[0], C.sizeof_TThostFtdcFaxType)
	ret.EMail = c2goStr(&p.EMail[0], C.sizeof_TThostFtdcEMailType)
	ret.MoneyAccountStatus = byte(p.MoneyAccountStatus)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.NewBankAccount = c2goStr(&p.NewBankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.NewBankPassWord = c2goStr(&p.NewBankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.BankAccType = byte(p.BankAccType)
	ret.InstallID = int(p.InstallID)
	ret.VerifyCertNoFlag = byte(p.VerifyCertNoFlag)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.BankPwdFlag = byte(p.BankPwdFlag)
	ret.SecuPwdFlag = byte(p.SecuPwdFlag)
	ret.TID = int(p.TID)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	ret.LongCustomerName = c2goStr(&p.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ret
}
func (s *CThostFtdcChangeAccountField) CValue() *C.CThostFtdcChangeAccountField {
	ptr := C.newCThostFtdcChangeAccountField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.Gender = C.char(s.Gender)
	go2cStr(s.CountryCode, &ptr.CountryCode[0], C.sizeof_TThostFtdcCountryCodeType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.Address, &ptr.Address[0], C.sizeof_TThostFtdcAddressType)
	go2cStr(s.ZipCode, &ptr.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	go2cStr(s.Telephone, &ptr.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	go2cStr(s.MobilePhone, &ptr.MobilePhone[0], C.sizeof_TThostFtdcMobilePhoneType)
	go2cStr(s.Fax, &ptr.Fax[0], C.sizeof_TThostFtdcFaxType)
	go2cStr(s.EMail, &ptr.EMail[0], C.sizeof_TThostFtdcEMailType)
	ptr.MoneyAccountStatus = C.char(s.MoneyAccountStatus)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.NewBankAccount, &ptr.NewBankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.NewBankPassWord, &ptr.NewBankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	ptr.BankAccType = C.char(s.BankAccType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.VerifyCertNoFlag = C.char(s.VerifyCertNoFlag)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ptr.BankPwdFlag = C.char(s.BankPwdFlag)
	ptr.SecuPwdFlag = C.char(s.SecuPwdFlag)
	ptr.TID = C.int(s.TID)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	go2cStr(s.LongCustomerName, &ptr.LongCustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	return ptr
}

type CThostFtdcSecAgentACIDMapField struct {
	BrokerID         string
	UserID           string
	AccountID        string
	CurrencyID       string
	BrokerSecAgentID string
}

func NewCThostFtdcSecAgentACIDMapField(p *C.CThostFtdcSecAgentACIDMapField) *CThostFtdcSecAgentACIDMapField {
	ret := new(CThostFtdcSecAgentACIDMapField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.BrokerSecAgentID = c2goStr(&p.BrokerSecAgentID[0], C.sizeof_TThostFtdcAccountIDType)
	return ret
}
func (s *CThostFtdcSecAgentACIDMapField) CValue() *C.CThostFtdcSecAgentACIDMapField {
	ptr := C.newCThostFtdcSecAgentACIDMapField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.BrokerSecAgentID, &ptr.BrokerSecAgentID[0], C.sizeof_TThostFtdcAccountIDType)
	return ptr
}

type CThostFtdcQrySecAgentACIDMapField struct {
	BrokerID   string
	UserID     string
	AccountID  string
	CurrencyID string
}

func NewCThostFtdcQrySecAgentACIDMapField(p *C.CThostFtdcQrySecAgentACIDMapField) *CThostFtdcQrySecAgentACIDMapField {
	ret := new(CThostFtdcQrySecAgentACIDMapField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ret
}
func (s *CThostFtdcQrySecAgentACIDMapField) CValue() *C.CThostFtdcQrySecAgentACIDMapField {
	ptr := C.newCThostFtdcQrySecAgentACIDMapField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ptr
}

type CThostFtdcUserRightsAssignField struct {
	BrokerID     string
	UserID       string
	DRIdentityID int
}

func NewCThostFtdcUserRightsAssignField(p *C.CThostFtdcUserRightsAssignField) *CThostFtdcUserRightsAssignField {
	ret := new(CThostFtdcUserRightsAssignField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.DRIdentityID = int(p.DRIdentityID)
	return ret
}
func (s *CThostFtdcUserRightsAssignField) CValue() *C.CThostFtdcUserRightsAssignField {
	ptr := C.newCThostFtdcUserRightsAssignField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.DRIdentityID = C.int(s.DRIdentityID)
	return ptr
}

type CThostFtdcBrokerUserRightAssignField struct {
	BrokerID     string
	DRIdentityID int
	Tradeable    int
}

func NewCThostFtdcBrokerUserRightAssignField(p *C.CThostFtdcBrokerUserRightAssignField) *CThostFtdcBrokerUserRightAssignField {
	ret := new(CThostFtdcBrokerUserRightAssignField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.DRIdentityID = int(p.DRIdentityID)
	ret.Tradeable = int(p.Tradeable)
	return ret
}
func (s *CThostFtdcBrokerUserRightAssignField) CValue() *C.CThostFtdcBrokerUserRightAssignField {
	ptr := C.newCThostFtdcBrokerUserRightAssignField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ptr.DRIdentityID = C.int(s.DRIdentityID)
	ptr.Tradeable = C.int(s.Tradeable)
	return ptr
}

type CThostFtdcDRTransferField struct {
	OrigDRIdentityID int
	DestDRIdentityID int
	OrigBrokerID     string
	DestBrokerID     string
}

func NewCThostFtdcDRTransferField(p *C.CThostFtdcDRTransferField) *CThostFtdcDRTransferField {
	ret := new(CThostFtdcDRTransferField)
	ret.OrigDRIdentityID = int(p.OrigDRIdentityID)
	ret.DestDRIdentityID = int(p.DestDRIdentityID)
	ret.OrigBrokerID = c2goStr(&p.OrigBrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.DestBrokerID = c2goStr(&p.DestBrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	return ret
}
func (s *CThostFtdcDRTransferField) CValue() *C.CThostFtdcDRTransferField {
	ptr := C.newCThostFtdcDRTransferField()
	ptr.OrigDRIdentityID = C.int(s.OrigDRIdentityID)
	ptr.DestDRIdentityID = C.int(s.DestDRIdentityID)
	go2cStr(s.OrigBrokerID, &ptr.OrigBrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.DestBrokerID, &ptr.DestBrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	return ptr
}

type CThostFtdcFensUserInfoField struct {
	BrokerID  string
	UserID    string
	LoginMode byte
}

func NewCThostFtdcFensUserInfoField(p *C.CThostFtdcFensUserInfoField) *CThostFtdcFensUserInfoField {
	ret := new(CThostFtdcFensUserInfoField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.LoginMode = byte(p.LoginMode)
	return ret
}
func (s *CThostFtdcFensUserInfoField) CValue() *C.CThostFtdcFensUserInfoField {
	ptr := C.newCThostFtdcFensUserInfoField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.LoginMode = C.char(s.LoginMode)
	return ptr
}

type CThostFtdcCurrTransferIdentityField struct {
	IdentityID int
}

func NewCThostFtdcCurrTransferIdentityField(p *C.CThostFtdcCurrTransferIdentityField) *CThostFtdcCurrTransferIdentityField {
	ret := new(CThostFtdcCurrTransferIdentityField)
	ret.IdentityID = int(p.IdentityID)
	return ret
}
func (s *CThostFtdcCurrTransferIdentityField) CValue() *C.CThostFtdcCurrTransferIdentityField {
	ptr := C.newCThostFtdcCurrTransferIdentityField()
	ptr.IdentityID = C.int(s.IdentityID)
	return ptr
}

type CThostFtdcLoginForbiddenUserField struct {
	BrokerID  string
	UserID    string
	Reserve1  string
	IPAddress string
}

func NewCThostFtdcLoginForbiddenUserField(p *C.CThostFtdcLoginForbiddenUserField) *CThostFtdcLoginForbiddenUserField {
	ret := new(CThostFtdcLoginForbiddenUserField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcLoginForbiddenUserField) CValue() *C.CThostFtdcLoginForbiddenUserField {
	ptr := C.newCThostFtdcLoginForbiddenUserField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryLoginForbiddenUserField struct {
	BrokerID string
	UserID   string
}

func NewCThostFtdcQryLoginForbiddenUserField(p *C.CThostFtdcQryLoginForbiddenUserField) *CThostFtdcQryLoginForbiddenUserField {
	ret := new(CThostFtdcQryLoginForbiddenUserField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ret
}
func (s *CThostFtdcQryLoginForbiddenUserField) CValue() *C.CThostFtdcQryLoginForbiddenUserField {
	ptr := C.newCThostFtdcQryLoginForbiddenUserField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ptr
}

type CThostFtdcTradingAccountReserveField struct {
	BrokerID   string
	AccountID  string
	Reserve    float64
	CurrencyID string
}

func NewCThostFtdcTradingAccountReserveField(p *C.CThostFtdcTradingAccountReserveField) *CThostFtdcTradingAccountReserveField {
	ret := new(CThostFtdcTradingAccountReserveField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Reserve = goFloat64(p.Reserve)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ret
}
func (s *CThostFtdcTradingAccountReserveField) CValue() *C.CThostFtdcTradingAccountReserveField {
	ptr := C.newCThostFtdcTradingAccountReserveField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ptr.Reserve = C.double(s.Reserve)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ptr
}

type CThostFtdcQryLoginForbiddenIPField struct {
	Reserve1  string
	IPAddress string
}

func NewCThostFtdcQryLoginForbiddenIPField(p *C.CThostFtdcQryLoginForbiddenIPField) *CThostFtdcQryLoginForbiddenIPField {
	ret := new(CThostFtdcQryLoginForbiddenIPField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcQryLoginForbiddenIPField) CValue() *C.CThostFtdcQryLoginForbiddenIPField {
	ptr := C.newCThostFtdcQryLoginForbiddenIPField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryIPListField struct {
	Reserve1  string
	IPAddress string
}

func NewCThostFtdcQryIPListField(p *C.CThostFtdcQryIPListField) *CThostFtdcQryIPListField {
	ret := new(CThostFtdcQryIPListField)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcQryIPListField) CValue() *C.CThostFtdcQryIPListField {
	ptr := C.newCThostFtdcQryIPListField()
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryUserRightsAssignField struct {
	BrokerID string
	UserID   string
}

func NewCThostFtdcQryUserRightsAssignField(p *C.CThostFtdcQryUserRightsAssignField) *CThostFtdcQryUserRightsAssignField {
	ret := new(CThostFtdcQryUserRightsAssignField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ret
}
func (s *CThostFtdcQryUserRightsAssignField) CValue() *C.CThostFtdcQryUserRightsAssignField {
	ptr := C.newCThostFtdcQryUserRightsAssignField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ptr
}

type CThostFtdcReserveOpenAccountConfirmField struct {
	TradeCode          string
	BankID             string
	BankBranchID       string
	BrokerID           string
	BrokerBranchID     string
	TradeDate          string
	TradeTime          string
	BankSerial         string
	TradingDay         string
	PlateSerial        int
	LastFragment       byte
	SessionID          int
	CustomerName       string
	IdCardType         byte
	IdentifiedCardNo   string
	Gender             byte
	CountryCode        string
	CustType           byte
	Address            string
	ZipCode            string
	Telephone          string
	MobilePhone        string
	Fax                string
	EMail              string
	MoneyAccountStatus byte
	BankAccount        string
	BankPassWord       string
	InstallID          int
	VerifyCertNoFlag   byte
	CurrencyID         string
	Digest             string
	BankAccType        byte
	BrokerIDByBank     string
	TID                int
	AccountID          string
	Password           string
	BankReserveOpenSeq string
	BookDate           string
	BookPsw            string
	ErrorID            int
	ErrorMsg           string
}

func NewCThostFtdcReserveOpenAccountConfirmField(p *C.CThostFtdcReserveOpenAccountConfirmField) *CThostFtdcReserveOpenAccountConfirmField {
	ret := new(CThostFtdcReserveOpenAccountConfirmField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.Gender = byte(p.Gender)
	ret.CountryCode = c2goStr(&p.CountryCode[0], C.sizeof_TThostFtdcCountryCodeType)
	ret.CustType = byte(p.CustType)
	ret.Address = c2goStr(&p.Address[0], C.sizeof_TThostFtdcAddressType)
	ret.ZipCode = c2goStr(&p.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	ret.Telephone = c2goStr(&p.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	ret.MobilePhone = c2goStr(&p.MobilePhone[0], C.sizeof_TThostFtdcMobilePhoneType)
	ret.Fax = c2goStr(&p.Fax[0], C.sizeof_TThostFtdcFaxType)
	ret.EMail = c2goStr(&p.EMail[0], C.sizeof_TThostFtdcEMailType)
	ret.MoneyAccountStatus = byte(p.MoneyAccountStatus)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.InstallID = int(p.InstallID)
	ret.VerifyCertNoFlag = byte(p.VerifyCertNoFlag)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.BankAccType = byte(p.BankAccType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.TID = int(p.TID)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.BankReserveOpenSeq = c2goStr(&p.BankReserveOpenSeq[0], C.sizeof_TThostFtdcBankSerialType)
	ret.BookDate = c2goStr(&p.BookDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.BookPsw = c2goStr(&p.BookPsw[0], C.sizeof_TThostFtdcPasswordType)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ret
}
func (s *CThostFtdcReserveOpenAccountConfirmField) CValue() *C.CThostFtdcReserveOpenAccountConfirmField {
	ptr := C.newCThostFtdcReserveOpenAccountConfirmField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.Gender = C.char(s.Gender)
	go2cStr(s.CountryCode, &ptr.CountryCode[0], C.sizeof_TThostFtdcCountryCodeType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.Address, &ptr.Address[0], C.sizeof_TThostFtdcAddressType)
	go2cStr(s.ZipCode, &ptr.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	go2cStr(s.Telephone, &ptr.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	go2cStr(s.MobilePhone, &ptr.MobilePhone[0], C.sizeof_TThostFtdcMobilePhoneType)
	go2cStr(s.Fax, &ptr.Fax[0], C.sizeof_TThostFtdcFaxType)
	go2cStr(s.EMail, &ptr.EMail[0], C.sizeof_TThostFtdcEMailType)
	ptr.MoneyAccountStatus = C.char(s.MoneyAccountStatus)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.VerifyCertNoFlag = C.char(s.VerifyCertNoFlag)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	ptr.BankAccType = C.char(s.BankAccType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ptr.TID = C.int(s.TID)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.BankReserveOpenSeq, &ptr.BankReserveOpenSeq[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.BookDate, &ptr.BookDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.BookPsw, &ptr.BookPsw[0], C.sizeof_TThostFtdcPasswordType)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ptr
}

type CThostFtdcReserveOpenAccountField struct {
	TradeCode          string
	BankID             string
	BankBranchID       string
	BrokerID           string
	BrokerBranchID     string
	TradeDate          string
	TradeTime          string
	BankSerial         string
	TradingDay         string
	PlateSerial        int
	LastFragment       byte
	SessionID          int
	CustomerName       string
	IdCardType         byte
	IdentifiedCardNo   string
	Gender             byte
	CountryCode        string
	CustType           byte
	Address            string
	ZipCode            string
	Telephone          string
	MobilePhone        string
	Fax                string
	EMail              string
	MoneyAccountStatus byte
	BankAccount        string
	BankPassWord       string
	InstallID          int
	VerifyCertNoFlag   byte
	CurrencyID         string
	Digest             string
	BankAccType        byte
	BrokerIDByBank     string
	TID                int
	ReserveOpenAccStas byte
	ErrorID            int
	ErrorMsg           string
}

func NewCThostFtdcReserveOpenAccountField(p *C.CThostFtdcReserveOpenAccountField) *CThostFtdcReserveOpenAccountField {
	ret := new(CThostFtdcReserveOpenAccountField)
	ret.TradeCode = c2goStr(&p.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankBranchID = c2goStr(&p.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerBranchID = c2goStr(&p.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	ret.TradeDate = c2goStr(&p.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	ret.TradeTime = c2goStr(&p.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	ret.BankSerial = c2goStr(&p.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.PlateSerial = int(p.PlateSerial)
	ret.LastFragment = byte(p.LastFragment)
	ret.SessionID = int(p.SessionID)
	ret.CustomerName = c2goStr(&p.CustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	ret.IdCardType = byte(p.IdCardType)
	ret.IdentifiedCardNo = c2goStr(&p.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ret.Gender = byte(p.Gender)
	ret.CountryCode = c2goStr(&p.CountryCode[0], C.sizeof_TThostFtdcCountryCodeType)
	ret.CustType = byte(p.CustType)
	ret.Address = c2goStr(&p.Address[0], C.sizeof_TThostFtdcAddressType)
	ret.ZipCode = c2goStr(&p.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	ret.Telephone = c2goStr(&p.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	ret.MobilePhone = c2goStr(&p.MobilePhone[0], C.sizeof_TThostFtdcMobilePhoneType)
	ret.Fax = c2goStr(&p.Fax[0], C.sizeof_TThostFtdcFaxType)
	ret.EMail = c2goStr(&p.EMail[0], C.sizeof_TThostFtdcEMailType)
	ret.MoneyAccountStatus = byte(p.MoneyAccountStatus)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.BankPassWord = c2goStr(&p.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ret.InstallID = int(p.InstallID)
	ret.VerifyCertNoFlag = byte(p.VerifyCertNoFlag)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.Digest = c2goStr(&p.Digest[0], C.sizeof_TThostFtdcDigestType)
	ret.BankAccType = byte(p.BankAccType)
	ret.BrokerIDByBank = c2goStr(&p.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ret.TID = int(p.TID)
	ret.ReserveOpenAccStas = byte(p.ReserveOpenAccStas)
	ret.ErrorID = int(p.ErrorID)
	ret.ErrorMsg = c2goStr(&p.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ret
}
func (s *CThostFtdcReserveOpenAccountField) CValue() *C.CThostFtdcReserveOpenAccountField {
	ptr := C.newCThostFtdcReserveOpenAccountField()
	go2cStr(s.TradeCode, &ptr.TradeCode[0], C.sizeof_TThostFtdcTradeCodeType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankBranchID, &ptr.BankBranchID[0], C.sizeof_TThostFtdcBankBrchIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerBranchID, &ptr.BrokerBranchID[0], C.sizeof_TThostFtdcFutureBranchIDType)
	go2cStr(s.TradeDate, &ptr.TradeDate[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.TradeTime, &ptr.TradeTime[0], C.sizeof_TThostFtdcTradeTimeType)
	go2cStr(s.BankSerial, &ptr.BankSerial[0], C.sizeof_TThostFtdcBankSerialType)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ptr.PlateSerial = C.int(s.PlateSerial)
	ptr.LastFragment = C.char(s.LastFragment)
	ptr.SessionID = C.int(s.SessionID)
	go2cStr(s.CustomerName, &ptr.CustomerName[0], C.sizeof_TThostFtdcLongIndividualNameType)
	ptr.IdCardType = C.char(s.IdCardType)
	go2cStr(s.IdentifiedCardNo, &ptr.IdentifiedCardNo[0], C.sizeof_TThostFtdcIdentifiedCardNoType)
	ptr.Gender = C.char(s.Gender)
	go2cStr(s.CountryCode, &ptr.CountryCode[0], C.sizeof_TThostFtdcCountryCodeType)
	ptr.CustType = C.char(s.CustType)
	go2cStr(s.Address, &ptr.Address[0], C.sizeof_TThostFtdcAddressType)
	go2cStr(s.ZipCode, &ptr.ZipCode[0], C.sizeof_TThostFtdcZipCodeType)
	go2cStr(s.Telephone, &ptr.Telephone[0], C.sizeof_TThostFtdcTelephoneType)
	go2cStr(s.MobilePhone, &ptr.MobilePhone[0], C.sizeof_TThostFtdcMobilePhoneType)
	go2cStr(s.Fax, &ptr.Fax[0], C.sizeof_TThostFtdcFaxType)
	go2cStr(s.EMail, &ptr.EMail[0], C.sizeof_TThostFtdcEMailType)
	ptr.MoneyAccountStatus = C.char(s.MoneyAccountStatus)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.BankPassWord, &ptr.BankPassWord[0], C.sizeof_TThostFtdcPasswordType)
	ptr.InstallID = C.int(s.InstallID)
	ptr.VerifyCertNoFlag = C.char(s.VerifyCertNoFlag)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	go2cStr(s.Digest, &ptr.Digest[0], C.sizeof_TThostFtdcDigestType)
	ptr.BankAccType = C.char(s.BankAccType)
	go2cStr(s.BrokerIDByBank, &ptr.BrokerIDByBank[0], C.sizeof_TThostFtdcBankCodingForFutureType)
	ptr.TID = C.int(s.TID)
	ptr.ReserveOpenAccStas = C.char(s.ReserveOpenAccStas)
	ptr.ErrorID = C.int(s.ErrorID)
	go2cStr(s.ErrorMsg, &ptr.ErrorMsg[0], C.sizeof_TThostFtdcErrorMsgType)
	return ptr
}

type CThostFtdcAccountPropertyField struct {
	BrokerID          string
	AccountID         string
	BankID            string
	BankAccount       string
	OpenName          string
	OpenBank          string
	IsActive          int
	AccountSourceType byte
	OpenDate          string
	CancelDate        string
	OperatorID        string
	OperateDate       string
	OperateTime       string
	CurrencyID        string
}

func NewCThostFtdcAccountPropertyField(p *C.CThostFtdcAccountPropertyField) *CThostFtdcAccountPropertyField {
	ret := new(CThostFtdcAccountPropertyField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.BankID = c2goStr(&p.BankID[0], C.sizeof_TThostFtdcBankIDType)
	ret.BankAccount = c2goStr(&p.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	ret.OpenName = c2goStr(&p.OpenName[0], C.sizeof_TThostFtdcInvestorFullNameType)
	ret.OpenBank = c2goStr(&p.OpenBank[0], C.sizeof_TThostFtdcOpenBankType)
	ret.IsActive = int(p.IsActive)
	ret.AccountSourceType = byte(p.AccountSourceType)
	ret.OpenDate = c2goStr(&p.OpenDate[0], C.sizeof_TThostFtdcDateType)
	ret.CancelDate = c2goStr(&p.CancelDate[0], C.sizeof_TThostFtdcDateType)
	ret.OperatorID = c2goStr(&p.OperatorID[0], C.sizeof_TThostFtdcOperatorIDType)
	ret.OperateDate = c2goStr(&p.OperateDate[0], C.sizeof_TThostFtdcDateType)
	ret.OperateTime = c2goStr(&p.OperateTime[0], C.sizeof_TThostFtdcTimeType)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ret
}
func (s *CThostFtdcAccountPropertyField) CValue() *C.CThostFtdcAccountPropertyField {
	ptr := C.newCThostFtdcAccountPropertyField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	go2cStr(s.BankID, &ptr.BankID[0], C.sizeof_TThostFtdcBankIDType)
	go2cStr(s.BankAccount, &ptr.BankAccount[0], C.sizeof_TThostFtdcBankAccountType)
	go2cStr(s.OpenName, &ptr.OpenName[0], C.sizeof_TThostFtdcInvestorFullNameType)
	go2cStr(s.OpenBank, &ptr.OpenBank[0], C.sizeof_TThostFtdcOpenBankType)
	ptr.IsActive = C.int(s.IsActive)
	ptr.AccountSourceType = C.char(s.AccountSourceType)
	go2cStr(s.OpenDate, &ptr.OpenDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.CancelDate, &ptr.CancelDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.OperatorID, &ptr.OperatorID[0], C.sizeof_TThostFtdcOperatorIDType)
	go2cStr(s.OperateDate, &ptr.OperateDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.OperateTime, &ptr.OperateTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	return ptr
}

type CThostFtdcQryCurrDRIdentityField struct {
	DRIdentityID int
}

func NewCThostFtdcQryCurrDRIdentityField(p *C.CThostFtdcQryCurrDRIdentityField) *CThostFtdcQryCurrDRIdentityField {
	ret := new(CThostFtdcQryCurrDRIdentityField)
	ret.DRIdentityID = int(p.DRIdentityID)
	return ret
}
func (s *CThostFtdcQryCurrDRIdentityField) CValue() *C.CThostFtdcQryCurrDRIdentityField {
	ptr := C.newCThostFtdcQryCurrDRIdentityField()
	ptr.DRIdentityID = C.int(s.DRIdentityID)
	return ptr
}

type CThostFtdcCurrDRIdentityField struct {
	DRIdentityID int
}

func NewCThostFtdcCurrDRIdentityField(p *C.CThostFtdcCurrDRIdentityField) *CThostFtdcCurrDRIdentityField {
	ret := new(CThostFtdcCurrDRIdentityField)
	ret.DRIdentityID = int(p.DRIdentityID)
	return ret
}
func (s *CThostFtdcCurrDRIdentityField) CValue() *C.CThostFtdcCurrDRIdentityField {
	ptr := C.newCThostFtdcCurrDRIdentityField()
	ptr.DRIdentityID = C.int(s.DRIdentityID)
	return ptr
}

type CThostFtdcQrySecAgentCheckModeField struct {
	BrokerID   string
	InvestorID string
}

func NewCThostFtdcQrySecAgentCheckModeField(p *C.CThostFtdcQrySecAgentCheckModeField) *CThostFtdcQrySecAgentCheckModeField {
	ret := new(CThostFtdcQrySecAgentCheckModeField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ret
}
func (s *CThostFtdcQrySecAgentCheckModeField) CValue() *C.CThostFtdcQrySecAgentCheckModeField {
	ptr := C.newCThostFtdcQrySecAgentCheckModeField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ptr
}

type CThostFtdcQrySecAgentTradeInfoField struct {
	BrokerID         string
	BrokerSecAgentID string
}

func NewCThostFtdcQrySecAgentTradeInfoField(p *C.CThostFtdcQrySecAgentTradeInfoField) *CThostFtdcQrySecAgentTradeInfoField {
	ret := new(CThostFtdcQrySecAgentTradeInfoField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.BrokerSecAgentID = c2goStr(&p.BrokerSecAgentID[0], C.sizeof_TThostFtdcAccountIDType)
	return ret
}
func (s *CThostFtdcQrySecAgentTradeInfoField) CValue() *C.CThostFtdcQrySecAgentTradeInfoField {
	ptr := C.newCThostFtdcQrySecAgentTradeInfoField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.BrokerSecAgentID, &ptr.BrokerSecAgentID[0], C.sizeof_TThostFtdcAccountIDType)
	return ptr
}

type CThostFtdcReqUserAuthMethodField struct {
	TradingDay string
	BrokerID   string
	UserID     string
}

func NewCThostFtdcReqUserAuthMethodField(p *C.CThostFtdcReqUserAuthMethodField) *CThostFtdcReqUserAuthMethodField {
	ret := new(CThostFtdcReqUserAuthMethodField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ret
}
func (s *CThostFtdcReqUserAuthMethodField) CValue() *C.CThostFtdcReqUserAuthMethodField {
	ptr := C.newCThostFtdcReqUserAuthMethodField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ptr
}

type CThostFtdcRspUserAuthMethodField struct {
	UsableAuthMethod int
}

func NewCThostFtdcRspUserAuthMethodField(p *C.CThostFtdcRspUserAuthMethodField) *CThostFtdcRspUserAuthMethodField {
	ret := new(CThostFtdcRspUserAuthMethodField)
	ret.UsableAuthMethod = int(p.UsableAuthMethod)
	return ret
}
func (s *CThostFtdcRspUserAuthMethodField) CValue() *C.CThostFtdcRspUserAuthMethodField {
	ptr := C.newCThostFtdcRspUserAuthMethodField()
	ptr.UsableAuthMethod = C.int(s.UsableAuthMethod)
	return ptr
}

type CThostFtdcReqGenUserCaptchaField struct {
	TradingDay string
	BrokerID   string
	UserID     string
}

func NewCThostFtdcReqGenUserCaptchaField(p *C.CThostFtdcReqGenUserCaptchaField) *CThostFtdcReqGenUserCaptchaField {
	ret := new(CThostFtdcReqGenUserCaptchaField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ret
}
func (s *CThostFtdcReqGenUserCaptchaField) CValue() *C.CThostFtdcReqGenUserCaptchaField {
	ptr := C.newCThostFtdcReqGenUserCaptchaField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ptr
}

type CThostFtdcRspGenUserCaptchaField struct {
	BrokerID       string
	UserID         string
	CaptchaInfoLen int
	CaptchaInfo    string
}

func NewCThostFtdcRspGenUserCaptchaField(p *C.CThostFtdcRspGenUserCaptchaField) *CThostFtdcRspGenUserCaptchaField {
	ret := new(CThostFtdcRspGenUserCaptchaField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.CaptchaInfoLen = int(p.CaptchaInfoLen)
	ret.CaptchaInfo = c2goStr(&p.CaptchaInfo[0], C.sizeof_TThostFtdcCaptchaInfoType)
	return ret
}
func (s *CThostFtdcRspGenUserCaptchaField) CValue() *C.CThostFtdcRspGenUserCaptchaField {
	ptr := C.newCThostFtdcRspGenUserCaptchaField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.CaptchaInfoLen = C.int(s.CaptchaInfoLen)
	go2cStr(s.CaptchaInfo, &ptr.CaptchaInfo[0], C.sizeof_TThostFtdcCaptchaInfoType)
	return ptr
}

type CThostFtdcReqGenUserTextField struct {
	TradingDay string
	BrokerID   string
	UserID     string
}

func NewCThostFtdcReqGenUserTextField(p *C.CThostFtdcReqGenUserTextField) *CThostFtdcReqGenUserTextField {
	ret := new(CThostFtdcReqGenUserTextField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ret
}
func (s *CThostFtdcReqGenUserTextField) CValue() *C.CThostFtdcReqGenUserTextField {
	ptr := C.newCThostFtdcReqGenUserTextField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	return ptr
}

type CThostFtdcRspGenUserTextField struct {
	UserTextSeq int
}

func NewCThostFtdcRspGenUserTextField(p *C.CThostFtdcRspGenUserTextField) *CThostFtdcRspGenUserTextField {
	ret := new(CThostFtdcRspGenUserTextField)
	ret.UserTextSeq = int(p.UserTextSeq)
	return ret
}
func (s *CThostFtdcRspGenUserTextField) CValue() *C.CThostFtdcRspGenUserTextField {
	ptr := C.newCThostFtdcRspGenUserTextField()
	ptr.UserTextSeq = C.int(s.UserTextSeq)
	return ptr
}

type CThostFtdcReqUserLoginWithCaptchaField struct {
	TradingDay           string
	BrokerID             string
	UserID               string
	Password             string
	UserProductInfo      string
	InterfaceProductInfo string
	ProtocolInfo         string
	MacAddress           string
	Reserve1             string
	LoginRemark          string
	Captcha              string
	ClientIPPort         int
	ClientIPAddress      string
}

func NewCThostFtdcReqUserLoginWithCaptchaField(p *C.CThostFtdcReqUserLoginWithCaptchaField) *CThostFtdcReqUserLoginWithCaptchaField {
	ret := new(CThostFtdcReqUserLoginWithCaptchaField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.UserProductInfo = c2goStr(&p.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.InterfaceProductInfo = c2goStr(&p.InterfaceProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.ProtocolInfo = c2goStr(&p.ProtocolInfo[0], C.sizeof_TThostFtdcProtocolInfoType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.LoginRemark = c2goStr(&p.LoginRemark[0], C.sizeof_TThostFtdcLoginRemarkType)
	ret.Captcha = c2goStr(&p.Captcha[0], C.sizeof_TThostFtdcPasswordType)
	ret.ClientIPPort = int(p.ClientIPPort)
	ret.ClientIPAddress = c2goStr(&p.ClientIPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcReqUserLoginWithCaptchaField) CValue() *C.CThostFtdcReqUserLoginWithCaptchaField {
	ptr := C.newCThostFtdcReqUserLoginWithCaptchaField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.UserProductInfo, &ptr.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.InterfaceProductInfo, &ptr.InterfaceProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.ProtocolInfo, &ptr.ProtocolInfo[0], C.sizeof_TThostFtdcProtocolInfoType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.LoginRemark, &ptr.LoginRemark[0], C.sizeof_TThostFtdcLoginRemarkType)
	go2cStr(s.Captcha, &ptr.Captcha[0], C.sizeof_TThostFtdcPasswordType)
	ptr.ClientIPPort = C.int(s.ClientIPPort)
	go2cStr(s.ClientIPAddress, &ptr.ClientIPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcReqUserLoginWithTextField struct {
	TradingDay           string
	BrokerID             string
	UserID               string
	Password             string
	UserProductInfo      string
	InterfaceProductInfo string
	ProtocolInfo         string
	MacAddress           string
	Reserve1             string
	LoginRemark          string
	Text                 string
	ClientIPPort         int
	ClientIPAddress      string
}

func NewCThostFtdcReqUserLoginWithTextField(p *C.CThostFtdcReqUserLoginWithTextField) *CThostFtdcReqUserLoginWithTextField {
	ret := new(CThostFtdcReqUserLoginWithTextField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.UserProductInfo = c2goStr(&p.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.InterfaceProductInfo = c2goStr(&p.InterfaceProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.ProtocolInfo = c2goStr(&p.ProtocolInfo[0], C.sizeof_TThostFtdcProtocolInfoType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.LoginRemark = c2goStr(&p.LoginRemark[0], C.sizeof_TThostFtdcLoginRemarkType)
	ret.Text = c2goStr(&p.Text[0], C.sizeof_TThostFtdcPasswordType)
	ret.ClientIPPort = int(p.ClientIPPort)
	ret.ClientIPAddress = c2goStr(&p.ClientIPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcReqUserLoginWithTextField) CValue() *C.CThostFtdcReqUserLoginWithTextField {
	ptr := C.newCThostFtdcReqUserLoginWithTextField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.UserProductInfo, &ptr.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.InterfaceProductInfo, &ptr.InterfaceProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.ProtocolInfo, &ptr.ProtocolInfo[0], C.sizeof_TThostFtdcProtocolInfoType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.LoginRemark, &ptr.LoginRemark[0], C.sizeof_TThostFtdcLoginRemarkType)
	go2cStr(s.Text, &ptr.Text[0], C.sizeof_TThostFtdcPasswordType)
	ptr.ClientIPPort = C.int(s.ClientIPPort)
	go2cStr(s.ClientIPAddress, &ptr.ClientIPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcReqUserLoginWithOTPField struct {
	TradingDay           string
	BrokerID             string
	UserID               string
	Password             string
	UserProductInfo      string
	InterfaceProductInfo string
	ProtocolInfo         string
	MacAddress           string
	Reserve1             string
	LoginRemark          string
	OTPPassword          string
	ClientIPPort         int
	ClientIPAddress      string
}

func NewCThostFtdcReqUserLoginWithOTPField(p *C.CThostFtdcReqUserLoginWithOTPField) *CThostFtdcReqUserLoginWithOTPField {
	ret := new(CThostFtdcReqUserLoginWithOTPField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.Password = c2goStr(&p.Password[0], C.sizeof_TThostFtdcPasswordType)
	ret.UserProductInfo = c2goStr(&p.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.InterfaceProductInfo = c2goStr(&p.InterfaceProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	ret.ProtocolInfo = c2goStr(&p.ProtocolInfo[0], C.sizeof_TThostFtdcProtocolInfoType)
	ret.MacAddress = c2goStr(&p.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.LoginRemark = c2goStr(&p.LoginRemark[0], C.sizeof_TThostFtdcLoginRemarkType)
	ret.OTPPassword = c2goStr(&p.OTPPassword[0], C.sizeof_TThostFtdcPasswordType)
	ret.ClientIPPort = int(p.ClientIPPort)
	ret.ClientIPAddress = c2goStr(&p.ClientIPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcReqUserLoginWithOTPField) CValue() *C.CThostFtdcReqUserLoginWithOTPField {
	ptr := C.newCThostFtdcReqUserLoginWithOTPField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	go2cStr(s.Password, &ptr.Password[0], C.sizeof_TThostFtdcPasswordType)
	go2cStr(s.UserProductInfo, &ptr.UserProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.InterfaceProductInfo, &ptr.InterfaceProductInfo[0], C.sizeof_TThostFtdcProductInfoType)
	go2cStr(s.ProtocolInfo, &ptr.ProtocolInfo[0], C.sizeof_TThostFtdcProtocolInfoType)
	go2cStr(s.MacAddress, &ptr.MacAddress[0], C.sizeof_TThostFtdcMacAddressType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	go2cStr(s.LoginRemark, &ptr.LoginRemark[0], C.sizeof_TThostFtdcLoginRemarkType)
	go2cStr(s.OTPPassword, &ptr.OTPPassword[0], C.sizeof_TThostFtdcPasswordType)
	ptr.ClientIPPort = C.int(s.ClientIPPort)
	go2cStr(s.ClientIPAddress, &ptr.ClientIPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcReqApiHandshakeField struct {
	CryptoKeyVersion string
}

func NewCThostFtdcReqApiHandshakeField(p *C.CThostFtdcReqApiHandshakeField) *CThostFtdcReqApiHandshakeField {
	ret := new(CThostFtdcReqApiHandshakeField)
	ret.CryptoKeyVersion = c2goStr(&p.CryptoKeyVersion[0], C.sizeof_TThostFtdcCryptoKeyVersionType)
	return ret
}
func (s *CThostFtdcReqApiHandshakeField) CValue() *C.CThostFtdcReqApiHandshakeField {
	ptr := C.newCThostFtdcReqApiHandshakeField()
	go2cStr(s.CryptoKeyVersion, &ptr.CryptoKeyVersion[0], C.sizeof_TThostFtdcCryptoKeyVersionType)
	return ptr
}

type CThostFtdcRspApiHandshakeField struct {
	FrontHandshakeDataLen int
	FrontHandshakeData    string
	IsApiAuthEnabled      int
}

func NewCThostFtdcRspApiHandshakeField(p *C.CThostFtdcRspApiHandshakeField) *CThostFtdcRspApiHandshakeField {
	ret := new(CThostFtdcRspApiHandshakeField)
	ret.FrontHandshakeDataLen = int(p.FrontHandshakeDataLen)
	ret.FrontHandshakeData = c2goStr(&p.FrontHandshakeData[0], C.sizeof_TThostFtdcHandshakeDataType)
	ret.IsApiAuthEnabled = int(p.IsApiAuthEnabled)
	return ret
}
func (s *CThostFtdcRspApiHandshakeField) CValue() *C.CThostFtdcRspApiHandshakeField {
	ptr := C.newCThostFtdcRspApiHandshakeField()
	ptr.FrontHandshakeDataLen = C.int(s.FrontHandshakeDataLen)
	go2cStr(s.FrontHandshakeData, &ptr.FrontHandshakeData[0], C.sizeof_TThostFtdcHandshakeDataType)
	ptr.IsApiAuthEnabled = C.int(s.IsApiAuthEnabled)
	return ptr
}

type CThostFtdcReqVerifyApiKeyField struct {
	ApiHandshakeDataLen int
	ApiHandshakeData    string
}

func NewCThostFtdcReqVerifyApiKeyField(p *C.CThostFtdcReqVerifyApiKeyField) *CThostFtdcReqVerifyApiKeyField {
	ret := new(CThostFtdcReqVerifyApiKeyField)
	ret.ApiHandshakeDataLen = int(p.ApiHandshakeDataLen)
	ret.ApiHandshakeData = c2goStr(&p.ApiHandshakeData[0], C.sizeof_TThostFtdcHandshakeDataType)
	return ret
}
func (s *CThostFtdcReqVerifyApiKeyField) CValue() *C.CThostFtdcReqVerifyApiKeyField {
	ptr := C.newCThostFtdcReqVerifyApiKeyField()
	ptr.ApiHandshakeDataLen = C.int(s.ApiHandshakeDataLen)
	go2cStr(s.ApiHandshakeData, &ptr.ApiHandshakeData[0], C.sizeof_TThostFtdcHandshakeDataType)
	return ptr
}

type CThostFtdcDepartmentUserField struct {
	BrokerID      string
	UserID        string
	InvestorRange byte
	InvestorID    string
}

func NewCThostFtdcDepartmentUserField(p *C.CThostFtdcDepartmentUserField) *CThostFtdcDepartmentUserField {
	ret := new(CThostFtdcDepartmentUserField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ret
}
func (s *CThostFtdcDepartmentUserField) CValue() *C.CThostFtdcDepartmentUserField {
	ptr := C.newCThostFtdcDepartmentUserField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	return ptr
}

type CThostFtdcQueryFreqField struct {
	QueryFreq int
}

func NewCThostFtdcQueryFreqField(p *C.CThostFtdcQueryFreqField) *CThostFtdcQueryFreqField {
	ret := new(CThostFtdcQueryFreqField)
	ret.QueryFreq = int(p.QueryFreq)
	return ret
}
func (s *CThostFtdcQueryFreqField) CValue() *C.CThostFtdcQueryFreqField {
	ptr := C.newCThostFtdcQueryFreqField()
	ptr.QueryFreq = C.int(s.QueryFreq)
	return ptr
}

type CThostFtdcAuthForbiddenIPField struct {
	IPAddress string
}

func NewCThostFtdcAuthForbiddenIPField(p *C.CThostFtdcAuthForbiddenIPField) *CThostFtdcAuthForbiddenIPField {
	ret := new(CThostFtdcAuthForbiddenIPField)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcAuthForbiddenIPField) CValue() *C.CThostFtdcAuthForbiddenIPField {
	ptr := C.newCThostFtdcAuthForbiddenIPField()
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryAuthForbiddenIPField struct {
	IPAddress string
}

func NewCThostFtdcQryAuthForbiddenIPField(p *C.CThostFtdcQryAuthForbiddenIPField) *CThostFtdcQryAuthForbiddenIPField {
	ret := new(CThostFtdcQryAuthForbiddenIPField)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcQryAuthForbiddenIPField) CValue() *C.CThostFtdcQryAuthForbiddenIPField {
	ptr := C.newCThostFtdcQryAuthForbiddenIPField()
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcSyncDelaySwapFrozenField struct {
	DelaySwapSeqNo string
	BrokerID       string
	InvestorID     string
	FromCurrencyID string
	FromRemainSwap float64
	IsManualSwap   int
}

func NewCThostFtdcSyncDelaySwapFrozenField(p *C.CThostFtdcSyncDelaySwapFrozenField) *CThostFtdcSyncDelaySwapFrozenField {
	ret := new(CThostFtdcSyncDelaySwapFrozenField)
	ret.DelaySwapSeqNo = c2goStr(&p.DelaySwapSeqNo[0], C.sizeof_TThostFtdcDepositSeqNoType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.FromCurrencyID = c2goStr(&p.FromCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.FromRemainSwap = goFloat64(p.FromRemainSwap)
	ret.IsManualSwap = int(p.IsManualSwap)
	return ret
}
func (s *CThostFtdcSyncDelaySwapFrozenField) CValue() *C.CThostFtdcSyncDelaySwapFrozenField {
	ptr := C.newCThostFtdcSyncDelaySwapFrozenField()
	go2cStr(s.DelaySwapSeqNo, &ptr.DelaySwapSeqNo[0], C.sizeof_TThostFtdcDepositSeqNoType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.FromCurrencyID, &ptr.FromCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.FromRemainSwap = C.double(s.FromRemainSwap)
	ptr.IsManualSwap = C.int(s.IsManualSwap)
	return ptr
}

type CThostFtdcUserSystemInfoField struct {
	BrokerID            string
	UserID              string
	ClientSystemInfoLen int
	ClientSystemInfo    string
	Reserve1            string
	ClientIPPort        int
	ClientLoginTime     string
	ClientAppID         string
	ClientPublicIP      string
	ClientLoginRemark   string
}

func NewCThostFtdcUserSystemInfoField(p *C.CThostFtdcUserSystemInfoField) *CThostFtdcUserSystemInfoField {
	ret := new(CThostFtdcUserSystemInfoField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.ClientSystemInfoLen = int(p.ClientSystemInfoLen)
	ret.ClientSystemInfo = c2goStr(&p.ClientSystemInfo[0], C.sizeof_TThostFtdcClientSystemInfoType)
	ret.Reserve1 = c2goStr(&p.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ret.ClientIPPort = int(p.ClientIPPort)
	ret.ClientLoginTime = c2goStr(&p.ClientLoginTime[0], C.sizeof_TThostFtdcTimeType)
	ret.ClientAppID = c2goStr(&p.ClientAppID[0], C.sizeof_TThostFtdcAppIDType)
	ret.ClientPublicIP = c2goStr(&p.ClientPublicIP[0], C.sizeof_TThostFtdcIPAddressType)
	ret.ClientLoginRemark = c2goStr(&p.ClientLoginRemark[0], C.sizeof_TThostFtdcClientLoginRemarkType)
	return ret
}
func (s *CThostFtdcUserSystemInfoField) CValue() *C.CThostFtdcUserSystemInfoField {
	ptr := C.newCThostFtdcUserSystemInfoField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.ClientSystemInfoLen = C.int(s.ClientSystemInfoLen)
	go2cStr(s.ClientSystemInfo, &ptr.ClientSystemInfo[0], C.sizeof_TThostFtdcClientSystemInfoType)
	go2cStr(s.Reserve1, &ptr.reserve1[0], C.sizeof_TThostFtdcOldIPAddressType)
	ptr.ClientIPPort = C.int(s.ClientIPPort)
	go2cStr(s.ClientLoginTime, &ptr.ClientLoginTime[0], C.sizeof_TThostFtdcTimeType)
	go2cStr(s.ClientAppID, &ptr.ClientAppID[0], C.sizeof_TThostFtdcAppIDType)
	go2cStr(s.ClientPublicIP, &ptr.ClientPublicIP[0], C.sizeof_TThostFtdcIPAddressType)
	go2cStr(s.ClientLoginRemark, &ptr.ClientLoginRemark[0], C.sizeof_TThostFtdcClientLoginRemarkType)
	return ptr
}

type CThostFtdcAuthUserIDField struct {
	BrokerID string
	AppID    string
	UserID   string
	AuthType byte
}

func NewCThostFtdcAuthUserIDField(p *C.CThostFtdcAuthUserIDField) *CThostFtdcAuthUserIDField {
	ret := new(CThostFtdcAuthUserIDField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.AppID = c2goStr(&p.AppID[0], C.sizeof_TThostFtdcAppIDType)
	ret.UserID = c2goStr(&p.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ret.AuthType = byte(p.AuthType)
	return ret
}
func (s *CThostFtdcAuthUserIDField) CValue() *C.CThostFtdcAuthUserIDField {
	ptr := C.newCThostFtdcAuthUserIDField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.AppID, &ptr.AppID[0], C.sizeof_TThostFtdcAppIDType)
	go2cStr(s.UserID, &ptr.UserID[0], C.sizeof_TThostFtdcUserIDType)
	ptr.AuthType = C.char(s.AuthType)
	return ptr
}

type CThostFtdcAuthIPField struct {
	BrokerID  string
	AppID     string
	IPAddress string
}

func NewCThostFtdcAuthIPField(p *C.CThostFtdcAuthIPField) *CThostFtdcAuthIPField {
	ret := new(CThostFtdcAuthIPField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.AppID = c2goStr(&p.AppID[0], C.sizeof_TThostFtdcAppIDType)
	ret.IPAddress = c2goStr(&p.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ret
}
func (s *CThostFtdcAuthIPField) CValue() *C.CThostFtdcAuthIPField {
	ptr := C.newCThostFtdcAuthIPField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.AppID, &ptr.AppID[0], C.sizeof_TThostFtdcAppIDType)
	go2cStr(s.IPAddress, &ptr.IPAddress[0], C.sizeof_TThostFtdcIPAddressType)
	return ptr
}

type CThostFtdcQryClassifiedInstrumentField struct {
	InstrumentID   string
	ExchangeID     string
	ExchangeInstID string
	ProductID      string
	TradingType    byte
	ClassType      byte
}

func NewCThostFtdcQryClassifiedInstrumentField(p *C.CThostFtdcQryClassifiedInstrumentField) *CThostFtdcQryClassifiedInstrumentField {
	ret := new(CThostFtdcQryClassifiedInstrumentField)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.ProductID = c2goStr(&p.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.TradingType = byte(p.TradingType)
	ret.ClassType = byte(p.ClassType)
	return ret
}
func (s *CThostFtdcQryClassifiedInstrumentField) CValue() *C.CThostFtdcQryClassifiedInstrumentField {
	ptr := C.newCThostFtdcQryClassifiedInstrumentField()
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	go2cStr(s.ProductID, &ptr.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ptr.TradingType = C.char(s.TradingType)
	ptr.ClassType = C.char(s.ClassType)
	return ptr
}

type CThostFtdcQryCombPromotionParamField struct {
	ExchangeID   string
	InstrumentID string
}

func NewCThostFtdcQryCombPromotionParamField(p *C.CThostFtdcQryCombPromotionParamField) *CThostFtdcQryCombPromotionParamField {
	ret := new(CThostFtdcQryCombPromotionParamField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryCombPromotionParamField) CValue() *C.CThostFtdcQryCombPromotionParamField {
	ptr := C.newCThostFtdcQryCombPromotionParamField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcCombPromotionParamField struct {
	ExchangeID    string
	InstrumentID  string
	CombHedgeFlag string
	Xparameter    float64
}

func NewCThostFtdcCombPromotionParamField(p *C.CThostFtdcCombPromotionParamField) *CThostFtdcCombPromotionParamField {
	ret := new(CThostFtdcCombPromotionParamField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.CombHedgeFlag = c2goStr(&p.CombHedgeFlag[0], C.sizeof_TThostFtdcCombHedgeFlagType)
	ret.Xparameter = goFloat64(p.Xparameter)
	return ret
}
func (s *CThostFtdcCombPromotionParamField) CValue() *C.CThostFtdcCombPromotionParamField {
	ptr := C.newCThostFtdcCombPromotionParamField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.CombHedgeFlag, &ptr.CombHedgeFlag[0], C.sizeof_TThostFtdcCombHedgeFlagType)
	ptr.Xparameter = C.double(s.Xparameter)
	return ptr
}

type CThostFtdcQryRiskSettleInvstPositionField struct {
	BrokerID     string
	InvestorID   string
	InstrumentID string
}

func NewCThostFtdcQryRiskSettleInvstPositionField(p *C.CThostFtdcQryRiskSettleInvstPositionField) *CThostFtdcQryRiskSettleInvstPositionField {
	ret := new(CThostFtdcQryRiskSettleInvstPositionField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryRiskSettleInvstPositionField) CValue() *C.CThostFtdcQryRiskSettleInvstPositionField {
	ptr := C.newCThostFtdcQryRiskSettleInvstPositionField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcQryRiskSettleProductStatusField struct {
	ProductID string
}

func NewCThostFtdcQryRiskSettleProductStatusField(p *C.CThostFtdcQryRiskSettleProductStatusField) *CThostFtdcQryRiskSettleProductStatusField {
	ret := new(CThostFtdcQryRiskSettleProductStatusField)
	ret.ProductID = c2goStr(&p.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ret
}
func (s *CThostFtdcQryRiskSettleProductStatusField) CValue() *C.CThostFtdcQryRiskSettleProductStatusField {
	ptr := C.newCThostFtdcQryRiskSettleProductStatusField()
	go2cStr(s.ProductID, &ptr.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	return ptr
}

type CThostFtdcRiskSettleInvstPositionField struct {
	InstrumentID       string
	BrokerID           string
	InvestorID         string
	PosiDirection      byte
	HedgeFlag          byte
	PositionDate       byte
	YdPosition         int
	Position           int
	LongFrozen         int
	ShortFrozen        int
	LongFrozenAmount   float64
	ShortFrozenAmount  float64
	OpenVolume         int
	CloseVolume        int
	OpenAmount         float64
	CloseAmount        float64
	PositionCost       float64
	PreMargin          float64
	UseMargin          float64
	FrozenMargin       float64
	FrozenCash         float64
	FrozenCommission   float64
	CashIn             float64
	Commission         float64
	CloseProfit        float64
	PositionProfit     float64
	PreSettlementPrice float64
	SettlementPrice    float64
	TradingDay         string
	SettlementID       int
	OpenCost           float64
	ExchangeMargin     float64
	CombPosition       int
	CombLongFrozen     int
	CombShortFrozen    int
	CloseProfitByDate  float64
	CloseProfitByTrade float64
	TodayPosition      int
	MarginRateByMoney  float64
	MarginRateByVolume float64
	StrikeFrozen       int
	StrikeFrozenAmount float64
	AbandonFrozen      int
	ExchangeID         string
	YdStrikeFrozen     int
	InvestUnitID       string
	PositionCostOffset float64
	TasPosition        int
	TasPositionCost    float64
}

func NewCThostFtdcRiskSettleInvstPositionField(p *C.CThostFtdcRiskSettleInvstPositionField) *CThostFtdcRiskSettleInvstPositionField {
	ret := new(CThostFtdcRiskSettleInvstPositionField)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.PosiDirection = byte(p.PosiDirection)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.PositionDate = byte(p.PositionDate)
	ret.YdPosition = int(p.YdPosition)
	ret.Position = int(p.Position)
	ret.LongFrozen = int(p.LongFrozen)
	ret.ShortFrozen = int(p.ShortFrozen)
	ret.LongFrozenAmount = goFloat64(p.LongFrozenAmount)
	ret.ShortFrozenAmount = goFloat64(p.ShortFrozenAmount)
	ret.OpenVolume = int(p.OpenVolume)
	ret.CloseVolume = int(p.CloseVolume)
	ret.OpenAmount = goFloat64(p.OpenAmount)
	ret.CloseAmount = goFloat64(p.CloseAmount)
	ret.PositionCost = goFloat64(p.PositionCost)
	ret.PreMargin = goFloat64(p.PreMargin)
	ret.UseMargin = goFloat64(p.UseMargin)
	ret.FrozenMargin = goFloat64(p.FrozenMargin)
	ret.FrozenCash = goFloat64(p.FrozenCash)
	ret.FrozenCommission = goFloat64(p.FrozenCommission)
	ret.CashIn = goFloat64(p.CashIn)
	ret.Commission = goFloat64(p.Commission)
	ret.CloseProfit = goFloat64(p.CloseProfit)
	ret.PositionProfit = goFloat64(p.PositionProfit)
	ret.PreSettlementPrice = goFloat64(p.PreSettlementPrice)
	ret.SettlementPrice = goFloat64(p.SettlementPrice)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.OpenCost = goFloat64(p.OpenCost)
	ret.ExchangeMargin = goFloat64(p.ExchangeMargin)
	ret.CombPosition = int(p.CombPosition)
	ret.CombLongFrozen = int(p.CombLongFrozen)
	ret.CombShortFrozen = int(p.CombShortFrozen)
	ret.CloseProfitByDate = goFloat64(p.CloseProfitByDate)
	ret.CloseProfitByTrade = goFloat64(p.CloseProfitByTrade)
	ret.TodayPosition = int(p.TodayPosition)
	ret.MarginRateByMoney = goFloat64(p.MarginRateByMoney)
	ret.MarginRateByVolume = goFloat64(p.MarginRateByVolume)
	ret.StrikeFrozen = int(p.StrikeFrozen)
	ret.StrikeFrozenAmount = goFloat64(p.StrikeFrozenAmount)
	ret.AbandonFrozen = int(p.AbandonFrozen)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.YdStrikeFrozen = int(p.YdStrikeFrozen)
	ret.InvestUnitID = c2goStr(&p.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ret.PositionCostOffset = goFloat64(p.PositionCostOffset)
	ret.TasPosition = int(p.TasPosition)
	ret.TasPositionCost = goFloat64(p.TasPositionCost)
	return ret
}
func (s *CThostFtdcRiskSettleInvstPositionField) CValue() *C.CThostFtdcRiskSettleInvstPositionField {
	ptr := C.newCThostFtdcRiskSettleInvstPositionField()
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.PosiDirection = C.char(s.PosiDirection)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.PositionDate = C.char(s.PositionDate)
	ptr.YdPosition = C.int(s.YdPosition)
	ptr.Position = C.int(s.Position)
	ptr.LongFrozen = C.int(s.LongFrozen)
	ptr.ShortFrozen = C.int(s.ShortFrozen)
	ptr.LongFrozenAmount = C.double(s.LongFrozenAmount)
	ptr.ShortFrozenAmount = C.double(s.ShortFrozenAmount)
	ptr.OpenVolume = C.int(s.OpenVolume)
	ptr.CloseVolume = C.int(s.CloseVolume)
	ptr.OpenAmount = C.double(s.OpenAmount)
	ptr.CloseAmount = C.double(s.CloseAmount)
	ptr.PositionCost = C.double(s.PositionCost)
	ptr.PreMargin = C.double(s.PreMargin)
	ptr.UseMargin = C.double(s.UseMargin)
	ptr.FrozenMargin = C.double(s.FrozenMargin)
	ptr.FrozenCash = C.double(s.FrozenCash)
	ptr.FrozenCommission = C.double(s.FrozenCommission)
	ptr.CashIn = C.double(s.CashIn)
	ptr.Commission = C.double(s.Commission)
	ptr.CloseProfit = C.double(s.CloseProfit)
	ptr.PositionProfit = C.double(s.PositionProfit)
	ptr.PreSettlementPrice = C.double(s.PreSettlementPrice)
	ptr.SettlementPrice = C.double(s.SettlementPrice)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	ptr.OpenCost = C.double(s.OpenCost)
	ptr.ExchangeMargin = C.double(s.ExchangeMargin)
	ptr.CombPosition = C.int(s.CombPosition)
	ptr.CombLongFrozen = C.int(s.CombLongFrozen)
	ptr.CombShortFrozen = C.int(s.CombShortFrozen)
	ptr.CloseProfitByDate = C.double(s.CloseProfitByDate)
	ptr.CloseProfitByTrade = C.double(s.CloseProfitByTrade)
	ptr.TodayPosition = C.int(s.TodayPosition)
	ptr.MarginRateByMoney = C.double(s.MarginRateByMoney)
	ptr.MarginRateByVolume = C.double(s.MarginRateByVolume)
	ptr.StrikeFrozen = C.int(s.StrikeFrozen)
	ptr.StrikeFrozenAmount = C.double(s.StrikeFrozenAmount)
	ptr.AbandonFrozen = C.int(s.AbandonFrozen)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ptr.YdStrikeFrozen = C.int(s.YdStrikeFrozen)
	go2cStr(s.InvestUnitID, &ptr.InvestUnitID[0], C.sizeof_TThostFtdcInvestUnitIDType)
	ptr.PositionCostOffset = C.double(s.PositionCostOffset)
	ptr.TasPosition = C.int(s.TasPosition)
	ptr.TasPositionCost = C.double(s.TasPositionCost)
	return ptr
}

type CThostFtdcRiskSettleProductStatusField struct {
	ExchangeID    string
	ProductID     string
	ProductStatus byte
}

func NewCThostFtdcRiskSettleProductStatusField(p *C.CThostFtdcRiskSettleProductStatusField) *CThostFtdcRiskSettleProductStatusField {
	ret := new(CThostFtdcRiskSettleProductStatusField)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ProductID = c2goStr(&p.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ProductStatus = byte(p.ProductStatus)
	return ret
}
func (s *CThostFtdcRiskSettleProductStatusField) CValue() *C.CThostFtdcRiskSettleProductStatusField {
	ptr := C.newCThostFtdcRiskSettleProductStatusField()
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ProductID, &ptr.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ptr.ProductStatus = C.char(s.ProductStatus)
	return ptr
}

type CThostFtdcSyncDeltaInfoField struct {
	SyncDeltaSequenceNo int
	SyncDeltaStatus     byte
	SyncDescription     string
	IsOnlyTrdDelta      int
}

func NewCThostFtdcSyncDeltaInfoField(p *C.CThostFtdcSyncDeltaInfoField) *CThostFtdcSyncDeltaInfoField {
	ret := new(CThostFtdcSyncDeltaInfoField)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	ret.SyncDeltaStatus = byte(p.SyncDeltaStatus)
	ret.SyncDescription = c2goStr(&p.SyncDescription[0], C.sizeof_TThostFtdcSyncDescriptionType)
	ret.IsOnlyTrdDelta = int(p.IsOnlyTrdDelta)
	return ret
}
func (s *CThostFtdcSyncDeltaInfoField) CValue() *C.CThostFtdcSyncDeltaInfoField {
	ptr := C.newCThostFtdcSyncDeltaInfoField()
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	ptr.SyncDeltaStatus = C.char(s.SyncDeltaStatus)
	go2cStr(s.SyncDescription, &ptr.SyncDescription[0], C.sizeof_TThostFtdcSyncDescriptionType)
	ptr.IsOnlyTrdDelta = C.int(s.IsOnlyTrdDelta)
	return ptr
}

type CThostFtdcSyncDeltaProductStatusField struct {
	SyncDeltaSequenceNo int
	ExchangeID          string
	ProductID           string
	ProductStatus       byte
}

func NewCThostFtdcSyncDeltaProductStatusField(p *C.CThostFtdcSyncDeltaProductStatusField) *CThostFtdcSyncDeltaProductStatusField {
	ret := new(CThostFtdcSyncDeltaProductStatusField)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ProductID = c2goStr(&p.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ProductStatus = byte(p.ProductStatus)
	return ret
}
func (s *CThostFtdcSyncDeltaProductStatusField) CValue() *C.CThostFtdcSyncDeltaProductStatusField {
	ptr := C.newCThostFtdcSyncDeltaProductStatusField()
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ProductID, &ptr.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ptr.ProductStatus = C.char(s.ProductStatus)
	return ptr
}

type CThostFtdcSyncDeltaInvstPosDtlField struct {
	InstrumentID          string
	BrokerID              string
	InvestorID            string
	HedgeFlag             byte
	Direction             byte
	OpenDate              string
	TradeID               string
	Volume                int
	OpenPrice             float64
	TradingDay            string
	SettlementID          int
	TradeType             byte
	CombInstrumentID      string
	ExchangeID            string
	CloseProfitByDate     float64
	CloseProfitByTrade    float64
	PositionProfitByDate  float64
	PositionProfitByTrade float64
	Margin                float64
	ExchMargin            float64
	MarginRateByMoney     float64
	MarginRateByVolume    float64
	LastSettlementPrice   float64
	SettlementPrice       float64
	CloseVolume           int
	CloseAmount           float64
	TimeFirstVolume       int
	SpecPosiType          byte
	ActionDirection       byte
	SyncDeltaSequenceNo   int
}

func NewCThostFtdcSyncDeltaInvstPosDtlField(p *C.CThostFtdcSyncDeltaInvstPosDtlField) *CThostFtdcSyncDeltaInvstPosDtlField {
	ret := new(CThostFtdcSyncDeltaInvstPosDtlField)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.Direction = byte(p.Direction)
	ret.OpenDate = c2goStr(&p.OpenDate[0], C.sizeof_TThostFtdcDateType)
	ret.TradeID = c2goStr(&p.TradeID[0], C.sizeof_TThostFtdcTradeIDType)
	ret.Volume = int(p.Volume)
	ret.OpenPrice = goFloat64(p.OpenPrice)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.TradeType = byte(p.TradeType)
	ret.CombInstrumentID = c2goStr(&p.CombInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.CloseProfitByDate = goFloat64(p.CloseProfitByDate)
	ret.CloseProfitByTrade = goFloat64(p.CloseProfitByTrade)
	ret.PositionProfitByDate = goFloat64(p.PositionProfitByDate)
	ret.PositionProfitByTrade = goFloat64(p.PositionProfitByTrade)
	ret.Margin = goFloat64(p.Margin)
	ret.ExchMargin = goFloat64(p.ExchMargin)
	ret.MarginRateByMoney = goFloat64(p.MarginRateByMoney)
	ret.MarginRateByVolume = goFloat64(p.MarginRateByVolume)
	ret.LastSettlementPrice = goFloat64(p.LastSettlementPrice)
	ret.SettlementPrice = goFloat64(p.SettlementPrice)
	ret.CloseVolume = int(p.CloseVolume)
	ret.CloseAmount = goFloat64(p.CloseAmount)
	ret.TimeFirstVolume = int(p.TimeFirstVolume)
	ret.SpecPosiType = byte(p.SpecPosiType)
	ret.ActionDirection = byte(p.ActionDirection)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	return ret
}
func (s *CThostFtdcSyncDeltaInvstPosDtlField) CValue() *C.CThostFtdcSyncDeltaInvstPosDtlField {
	ptr := C.newCThostFtdcSyncDeltaInvstPosDtlField()
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.Direction = C.char(s.Direction)
	go2cStr(s.OpenDate, &ptr.OpenDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.TradeID, &ptr.TradeID[0], C.sizeof_TThostFtdcTradeIDType)
	ptr.Volume = C.int(s.Volume)
	ptr.OpenPrice = C.double(s.OpenPrice)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	ptr.TradeType = C.char(s.TradeType)
	go2cStr(s.CombInstrumentID, &ptr.CombInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ptr.CloseProfitByDate = C.double(s.CloseProfitByDate)
	ptr.CloseProfitByTrade = C.double(s.CloseProfitByTrade)
	ptr.PositionProfitByDate = C.double(s.PositionProfitByDate)
	ptr.PositionProfitByTrade = C.double(s.PositionProfitByTrade)
	ptr.Margin = C.double(s.Margin)
	ptr.ExchMargin = C.double(s.ExchMargin)
	ptr.MarginRateByMoney = C.double(s.MarginRateByMoney)
	ptr.MarginRateByVolume = C.double(s.MarginRateByVolume)
	ptr.LastSettlementPrice = C.double(s.LastSettlementPrice)
	ptr.SettlementPrice = C.double(s.SettlementPrice)
	ptr.CloseVolume = C.int(s.CloseVolume)
	ptr.CloseAmount = C.double(s.CloseAmount)
	ptr.TimeFirstVolume = C.int(s.TimeFirstVolume)
	ptr.SpecPosiType = C.char(s.SpecPosiType)
	ptr.ActionDirection = C.char(s.ActionDirection)
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	return ptr
}

type CThostFtdcSyncDeltaInvstPosCombDtlField struct {
	TradingDay          string
	OpenDate            string
	ExchangeID          string
	SettlementID        int
	BrokerID            string
	InvestorID          string
	ComTradeID          string
	TradeID             string
	InstrumentID        string
	HedgeFlag           byte
	Direction           byte
	TotalAmt            int
	Margin              float64
	ExchMargin          float64
	MarginRateByMoney   float64
	MarginRateByVolume  float64
	LegID               int
	LegMultiple         int
	TradeGroupID        int
	ActionDirection     byte
	SyncDeltaSequenceNo int
}

func NewCThostFtdcSyncDeltaInvstPosCombDtlField(p *C.CThostFtdcSyncDeltaInvstPosCombDtlField) *CThostFtdcSyncDeltaInvstPosCombDtlField {
	ret := new(CThostFtdcSyncDeltaInvstPosCombDtlField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.OpenDate = c2goStr(&p.OpenDate[0], C.sizeof_TThostFtdcDateType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.SettlementID = int(p.SettlementID)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.ComTradeID = c2goStr(&p.ComTradeID[0], C.sizeof_TThostFtdcTradeIDType)
	ret.TradeID = c2goStr(&p.TradeID[0], C.sizeof_TThostFtdcTradeIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.Direction = byte(p.Direction)
	ret.TotalAmt = int(p.TotalAmt)
	ret.Margin = goFloat64(p.Margin)
	ret.ExchMargin = goFloat64(p.ExchMargin)
	ret.MarginRateByMoney = goFloat64(p.MarginRateByMoney)
	ret.MarginRateByVolume = goFloat64(p.MarginRateByVolume)
	ret.LegID = int(p.LegID)
	ret.LegMultiple = int(p.LegMultiple)
	ret.TradeGroupID = int(p.TradeGroupID)
	ret.ActionDirection = byte(p.ActionDirection)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	return ret
}
func (s *CThostFtdcSyncDeltaInvstPosCombDtlField) CValue() *C.CThostFtdcSyncDeltaInvstPosCombDtlField {
	ptr := C.newCThostFtdcSyncDeltaInvstPosCombDtlField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.OpenDate, &ptr.OpenDate[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ptr.SettlementID = C.int(s.SettlementID)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.ComTradeID, &ptr.ComTradeID[0], C.sizeof_TThostFtdcTradeIDType)
	go2cStr(s.TradeID, &ptr.TradeID[0], C.sizeof_TThostFtdcTradeIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.Direction = C.char(s.Direction)
	ptr.TotalAmt = C.int(s.TotalAmt)
	ptr.Margin = C.double(s.Margin)
	ptr.ExchMargin = C.double(s.ExchMargin)
	ptr.MarginRateByMoney = C.double(s.MarginRateByMoney)
	ptr.MarginRateByVolume = C.double(s.MarginRateByVolume)
	ptr.LegID = C.int(s.LegID)
	ptr.LegMultiple = C.int(s.LegMultiple)
	ptr.TradeGroupID = C.int(s.TradeGroupID)
	ptr.ActionDirection = C.char(s.ActionDirection)
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	return ptr
}

type CThostFtdcSyncDeltaTradingAccountField struct {
	BrokerID                       string
	AccountID                      string
	PreMortgage                    float64
	PreCredit                      float64
	PreDeposit                     float64
	PreBalance                     float64
	PreMargin                      float64
	InterestBase                   float64
	Interest                       float64
	Deposit                        float64
	Withdraw                       float64
	FrozenMargin                   float64
	FrozenCash                     float64
	FrozenCommission               float64
	CurrMargin                     float64
	CashIn                         float64
	Commission                     float64
	CloseProfit                    float64
	PositionProfit                 float64
	Balance                        float64
	Available                      float64
	WithdrawQuota                  float64
	Reserve                        float64
	TradingDay                     string
	SettlementID                   int
	Credit                         float64
	Mortgage                       float64
	ExchangeMargin                 float64
	DeliveryMargin                 float64
	ExchangeDeliveryMargin         float64
	ReserveBalance                 float64
	CurrencyID                     string
	PreFundMortgageIn              float64
	PreFundMortgageOut             float64
	FundMortgageIn                 float64
	FundMortgageOut                float64
	FundMortgageAvailable          float64
	MortgageableFund               float64
	SpecProductMargin              float64
	SpecProductFrozenMargin        float64
	SpecProductCommission          float64
	SpecProductFrozenCommission    float64
	SpecProductPositionProfit      float64
	SpecProductCloseProfit         float64
	SpecProductPositionProfitByAlg float64
	SpecProductExchangeMargin      float64
	FrozenSwap                     float64
	RemainSwap                     float64
	SyncDeltaSequenceNo            int
}

func NewCThostFtdcSyncDeltaTradingAccountField(p *C.CThostFtdcSyncDeltaTradingAccountField) *CThostFtdcSyncDeltaTradingAccountField {
	ret := new(CThostFtdcSyncDeltaTradingAccountField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.AccountID = c2goStr(&p.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ret.PreMortgage = goFloat64(p.PreMortgage)
	ret.PreCredit = goFloat64(p.PreCredit)
	ret.PreDeposit = goFloat64(p.PreDeposit)
	ret.PreBalance = goFloat64(p.PreBalance)
	ret.PreMargin = goFloat64(p.PreMargin)
	ret.InterestBase = goFloat64(p.InterestBase)
	ret.Interest = goFloat64(p.Interest)
	ret.Deposit = goFloat64(p.Deposit)
	ret.Withdraw = goFloat64(p.Withdraw)
	ret.FrozenMargin = goFloat64(p.FrozenMargin)
	ret.FrozenCash = goFloat64(p.FrozenCash)
	ret.FrozenCommission = goFloat64(p.FrozenCommission)
	ret.CurrMargin = goFloat64(p.CurrMargin)
	ret.CashIn = goFloat64(p.CashIn)
	ret.Commission = goFloat64(p.Commission)
	ret.CloseProfit = goFloat64(p.CloseProfit)
	ret.PositionProfit = goFloat64(p.PositionProfit)
	ret.Balance = goFloat64(p.Balance)
	ret.Available = goFloat64(p.Available)
	ret.WithdrawQuota = goFloat64(p.WithdrawQuota)
	ret.Reserve = goFloat64(p.Reserve)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.SettlementID = int(p.SettlementID)
	ret.Credit = goFloat64(p.Credit)
	ret.Mortgage = goFloat64(p.Mortgage)
	ret.ExchangeMargin = goFloat64(p.ExchangeMargin)
	ret.DeliveryMargin = goFloat64(p.DeliveryMargin)
	ret.ExchangeDeliveryMargin = goFloat64(p.ExchangeDeliveryMargin)
	ret.ReserveBalance = goFloat64(p.ReserveBalance)
	ret.CurrencyID = c2goStr(&p.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.PreFundMortgageIn = goFloat64(p.PreFundMortgageIn)
	ret.PreFundMortgageOut = goFloat64(p.PreFundMortgageOut)
	ret.FundMortgageIn = goFloat64(p.FundMortgageIn)
	ret.FundMortgageOut = goFloat64(p.FundMortgageOut)
	ret.FundMortgageAvailable = goFloat64(p.FundMortgageAvailable)
	ret.MortgageableFund = goFloat64(p.MortgageableFund)
	ret.SpecProductMargin = goFloat64(p.SpecProductMargin)
	ret.SpecProductFrozenMargin = goFloat64(p.SpecProductFrozenMargin)
	ret.SpecProductCommission = goFloat64(p.SpecProductCommission)
	ret.SpecProductFrozenCommission = goFloat64(p.SpecProductFrozenCommission)
	ret.SpecProductPositionProfit = goFloat64(p.SpecProductPositionProfit)
	ret.SpecProductCloseProfit = goFloat64(p.SpecProductCloseProfit)
	ret.SpecProductPositionProfitByAlg = goFloat64(p.SpecProductPositionProfitByAlg)
	ret.SpecProductExchangeMargin = goFloat64(p.SpecProductExchangeMargin)
	ret.FrozenSwap = goFloat64(p.FrozenSwap)
	ret.RemainSwap = goFloat64(p.RemainSwap)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	return ret
}
func (s *CThostFtdcSyncDeltaTradingAccountField) CValue() *C.CThostFtdcSyncDeltaTradingAccountField {
	ptr := C.newCThostFtdcSyncDeltaTradingAccountField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.AccountID, &ptr.AccountID[0], C.sizeof_TThostFtdcAccountIDType)
	ptr.PreMortgage = C.double(s.PreMortgage)
	ptr.PreCredit = C.double(s.PreCredit)
	ptr.PreDeposit = C.double(s.PreDeposit)
	ptr.PreBalance = C.double(s.PreBalance)
	ptr.PreMargin = C.double(s.PreMargin)
	ptr.InterestBase = C.double(s.InterestBase)
	ptr.Interest = C.double(s.Interest)
	ptr.Deposit = C.double(s.Deposit)
	ptr.Withdraw = C.double(s.Withdraw)
	ptr.FrozenMargin = C.double(s.FrozenMargin)
	ptr.FrozenCash = C.double(s.FrozenCash)
	ptr.FrozenCommission = C.double(s.FrozenCommission)
	ptr.CurrMargin = C.double(s.CurrMargin)
	ptr.CashIn = C.double(s.CashIn)
	ptr.Commission = C.double(s.Commission)
	ptr.CloseProfit = C.double(s.CloseProfit)
	ptr.PositionProfit = C.double(s.PositionProfit)
	ptr.Balance = C.double(s.Balance)
	ptr.Available = C.double(s.Available)
	ptr.WithdrawQuota = C.double(s.WithdrawQuota)
	ptr.Reserve = C.double(s.Reserve)
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ptr.SettlementID = C.int(s.SettlementID)
	ptr.Credit = C.double(s.Credit)
	ptr.Mortgage = C.double(s.Mortgage)
	ptr.ExchangeMargin = C.double(s.ExchangeMargin)
	ptr.DeliveryMargin = C.double(s.DeliveryMargin)
	ptr.ExchangeDeliveryMargin = C.double(s.ExchangeDeliveryMargin)
	ptr.ReserveBalance = C.double(s.ReserveBalance)
	go2cStr(s.CurrencyID, &ptr.CurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.PreFundMortgageIn = C.double(s.PreFundMortgageIn)
	ptr.PreFundMortgageOut = C.double(s.PreFundMortgageOut)
	ptr.FundMortgageIn = C.double(s.FundMortgageIn)
	ptr.FundMortgageOut = C.double(s.FundMortgageOut)
	ptr.FundMortgageAvailable = C.double(s.FundMortgageAvailable)
	ptr.MortgageableFund = C.double(s.MortgageableFund)
	ptr.SpecProductMargin = C.double(s.SpecProductMargin)
	ptr.SpecProductFrozenMargin = C.double(s.SpecProductFrozenMargin)
	ptr.SpecProductCommission = C.double(s.SpecProductCommission)
	ptr.SpecProductFrozenCommission = C.double(s.SpecProductFrozenCommission)
	ptr.SpecProductPositionProfit = C.double(s.SpecProductPositionProfit)
	ptr.SpecProductCloseProfit = C.double(s.SpecProductCloseProfit)
	ptr.SpecProductPositionProfitByAlg = C.double(s.SpecProductPositionProfitByAlg)
	ptr.SpecProductExchangeMargin = C.double(s.SpecProductExchangeMargin)
	ptr.FrozenSwap = C.double(s.FrozenSwap)
	ptr.RemainSwap = C.double(s.RemainSwap)
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	return ptr
}

type CThostFtdcSyncDeltaInitInvstMarginField struct {
	BrokerID                     string
	InvestorID                   string
	LastRiskTotalInvstMargin     float64
	LastRiskTotalExchMargin      float64
	ThisSyncInvstMargin          float64
	ThisSyncExchMargin           float64
	RemainRiskInvstMargin        float64
	RemainRiskExchMargin         float64
	LastRiskSpecTotalInvstMargin float64
	LastRiskSpecTotalExchMargin  float64
	ThisSyncSpecInvstMargin      float64
	ThisSyncSpecExchMargin       float64
	RemainRiskSpecInvstMargin    float64
	RemainRiskSpecExchMargin     float64
	SyncDeltaSequenceNo          int
}

func NewCThostFtdcSyncDeltaInitInvstMarginField(p *C.CThostFtdcSyncDeltaInitInvstMarginField) *CThostFtdcSyncDeltaInitInvstMarginField {
	ret := new(CThostFtdcSyncDeltaInitInvstMarginField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.LastRiskTotalInvstMargin = goFloat64(p.LastRiskTotalInvstMargin)
	ret.LastRiskTotalExchMargin = goFloat64(p.LastRiskTotalExchMargin)
	ret.ThisSyncInvstMargin = goFloat64(p.ThisSyncInvstMargin)
	ret.ThisSyncExchMargin = goFloat64(p.ThisSyncExchMargin)
	ret.RemainRiskInvstMargin = goFloat64(p.RemainRiskInvstMargin)
	ret.RemainRiskExchMargin = goFloat64(p.RemainRiskExchMargin)
	ret.LastRiskSpecTotalInvstMargin = goFloat64(p.LastRiskSpecTotalInvstMargin)
	ret.LastRiskSpecTotalExchMargin = goFloat64(p.LastRiskSpecTotalExchMargin)
	ret.ThisSyncSpecInvstMargin = goFloat64(p.ThisSyncSpecInvstMargin)
	ret.ThisSyncSpecExchMargin = goFloat64(p.ThisSyncSpecExchMargin)
	ret.RemainRiskSpecInvstMargin = goFloat64(p.RemainRiskSpecInvstMargin)
	ret.RemainRiskSpecExchMargin = goFloat64(p.RemainRiskSpecExchMargin)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	return ret
}
func (s *CThostFtdcSyncDeltaInitInvstMarginField) CValue() *C.CThostFtdcSyncDeltaInitInvstMarginField {
	ptr := C.newCThostFtdcSyncDeltaInitInvstMarginField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.LastRiskTotalInvstMargin = C.double(s.LastRiskTotalInvstMargin)
	ptr.LastRiskTotalExchMargin = C.double(s.LastRiskTotalExchMargin)
	ptr.ThisSyncInvstMargin = C.double(s.ThisSyncInvstMargin)
	ptr.ThisSyncExchMargin = C.double(s.ThisSyncExchMargin)
	ptr.RemainRiskInvstMargin = C.double(s.RemainRiskInvstMargin)
	ptr.RemainRiskExchMargin = C.double(s.RemainRiskExchMargin)
	ptr.LastRiskSpecTotalInvstMargin = C.double(s.LastRiskSpecTotalInvstMargin)
	ptr.LastRiskSpecTotalExchMargin = C.double(s.LastRiskSpecTotalExchMargin)
	ptr.ThisSyncSpecInvstMargin = C.double(s.ThisSyncSpecInvstMargin)
	ptr.ThisSyncSpecExchMargin = C.double(s.ThisSyncSpecExchMargin)
	ptr.RemainRiskSpecInvstMargin = C.double(s.RemainRiskSpecInvstMargin)
	ptr.RemainRiskSpecExchMargin = C.double(s.RemainRiskSpecExchMargin)
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	return ptr
}

type CThostFtdcSyncDeltaDceCombInstrumentField struct {
	CombInstrumentID    string
	ExchangeID          string
	ExchangeInstID      string
	TradeGroupID        int
	CombHedgeFlag       byte
	CombinationType     byte
	Direction           byte
	ProductID           string
	Xparameter          float64
	ActionDirection     byte
	SyncDeltaSequenceNo int
}

func NewCThostFtdcSyncDeltaDceCombInstrumentField(p *C.CThostFtdcSyncDeltaDceCombInstrumentField) *CThostFtdcSyncDeltaDceCombInstrumentField {
	ret := new(CThostFtdcSyncDeltaDceCombInstrumentField)
	ret.CombInstrumentID = c2goStr(&p.CombInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.TradeGroupID = int(p.TradeGroupID)
	ret.CombHedgeFlag = byte(p.CombHedgeFlag)
	ret.CombinationType = byte(p.CombinationType)
	ret.Direction = byte(p.Direction)
	ret.ProductID = c2goStr(&p.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.Xparameter = goFloat64(p.Xparameter)
	ret.ActionDirection = byte(p.ActionDirection)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	return ret
}
func (s *CThostFtdcSyncDeltaDceCombInstrumentField) CValue() *C.CThostFtdcSyncDeltaDceCombInstrumentField {
	ptr := C.newCThostFtdcSyncDeltaDceCombInstrumentField()
	go2cStr(s.CombInstrumentID, &ptr.CombInstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ptr.TradeGroupID = C.int(s.TradeGroupID)
	ptr.CombHedgeFlag = C.char(s.CombHedgeFlag)
	ptr.CombinationType = C.char(s.CombinationType)
	ptr.Direction = C.char(s.Direction)
	go2cStr(s.ProductID, &ptr.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ptr.Xparameter = C.double(s.Xparameter)
	ptr.ActionDirection = C.char(s.ActionDirection)
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	return ptr
}

type CThostFtdcSyncDeltaInvstMarginRateField struct {
	InstrumentID             string
	InvestorRange            byte
	BrokerID                 string
	InvestorID               string
	HedgeFlag                byte
	LongMarginRatioByMoney   float64
	LongMarginRatioByVolume  float64
	ShortMarginRatioByMoney  float64
	ShortMarginRatioByVolume float64
	IsRelative               int
	ActionDirection          byte
	SyncDeltaSequenceNo      int
}

func NewCThostFtdcSyncDeltaInvstMarginRateField(p *C.CThostFtdcSyncDeltaInvstMarginRateField) *CThostFtdcSyncDeltaInvstMarginRateField {
	ret := new(CThostFtdcSyncDeltaInvstMarginRateField)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.LongMarginRatioByMoney = goFloat64(p.LongMarginRatioByMoney)
	ret.LongMarginRatioByVolume = goFloat64(p.LongMarginRatioByVolume)
	ret.ShortMarginRatioByMoney = goFloat64(p.ShortMarginRatioByMoney)
	ret.ShortMarginRatioByVolume = goFloat64(p.ShortMarginRatioByVolume)
	ret.IsRelative = int(p.IsRelative)
	ret.ActionDirection = byte(p.ActionDirection)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	return ret
}
func (s *CThostFtdcSyncDeltaInvstMarginRateField) CValue() *C.CThostFtdcSyncDeltaInvstMarginRateField {
	ptr := C.newCThostFtdcSyncDeltaInvstMarginRateField()
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.LongMarginRatioByMoney = C.double(s.LongMarginRatioByMoney)
	ptr.LongMarginRatioByVolume = C.double(s.LongMarginRatioByVolume)
	ptr.ShortMarginRatioByMoney = C.double(s.ShortMarginRatioByMoney)
	ptr.ShortMarginRatioByVolume = C.double(s.ShortMarginRatioByVolume)
	ptr.IsRelative = C.int(s.IsRelative)
	ptr.ActionDirection = C.char(s.ActionDirection)
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	return ptr
}

type CThostFtdcSyncDeltaExchMarginRateField struct {
	BrokerID                 string
	InstrumentID             string
	HedgeFlag                byte
	LongMarginRatioByMoney   float64
	LongMarginRatioByVolume  float64
	ShortMarginRatioByMoney  float64
	ShortMarginRatioByVolume float64
	ActionDirection          byte
	SyncDeltaSequenceNo      int
}

func NewCThostFtdcSyncDeltaExchMarginRateField(p *C.CThostFtdcSyncDeltaExchMarginRateField) *CThostFtdcSyncDeltaExchMarginRateField {
	ret := new(CThostFtdcSyncDeltaExchMarginRateField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.LongMarginRatioByMoney = goFloat64(p.LongMarginRatioByMoney)
	ret.LongMarginRatioByVolume = goFloat64(p.LongMarginRatioByVolume)
	ret.ShortMarginRatioByMoney = goFloat64(p.ShortMarginRatioByMoney)
	ret.ShortMarginRatioByVolume = goFloat64(p.ShortMarginRatioByVolume)
	ret.ActionDirection = byte(p.ActionDirection)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	return ret
}
func (s *CThostFtdcSyncDeltaExchMarginRateField) CValue() *C.CThostFtdcSyncDeltaExchMarginRateField {
	ptr := C.newCThostFtdcSyncDeltaExchMarginRateField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.LongMarginRatioByMoney = C.double(s.LongMarginRatioByMoney)
	ptr.LongMarginRatioByVolume = C.double(s.LongMarginRatioByVolume)
	ptr.ShortMarginRatioByMoney = C.double(s.ShortMarginRatioByMoney)
	ptr.ShortMarginRatioByVolume = C.double(s.ShortMarginRatioByVolume)
	ptr.ActionDirection = C.char(s.ActionDirection)
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	return ptr
}

type CThostFtdcSyncDeltaOptExchMarginField struct {
	BrokerID                  string
	InstrumentID              string
	SShortMarginRatioByMoney  float64
	SShortMarginRatioByVolume float64
	HShortMarginRatioByMoney  float64
	HShortMarginRatioByVolume float64
	AShortMarginRatioByMoney  float64
	AShortMarginRatioByVolume float64
	MShortMarginRatioByMoney  float64
	MShortMarginRatioByVolume float64
	ActionDirection           byte
	SyncDeltaSequenceNo       int
}

func NewCThostFtdcSyncDeltaOptExchMarginField(p *C.CThostFtdcSyncDeltaOptExchMarginField) *CThostFtdcSyncDeltaOptExchMarginField {
	ret := new(CThostFtdcSyncDeltaOptExchMarginField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.SShortMarginRatioByMoney = goFloat64(p.SShortMarginRatioByMoney)
	ret.SShortMarginRatioByVolume = goFloat64(p.SShortMarginRatioByVolume)
	ret.HShortMarginRatioByMoney = goFloat64(p.HShortMarginRatioByMoney)
	ret.HShortMarginRatioByVolume = goFloat64(p.HShortMarginRatioByVolume)
	ret.AShortMarginRatioByMoney = goFloat64(p.AShortMarginRatioByMoney)
	ret.AShortMarginRatioByVolume = goFloat64(p.AShortMarginRatioByVolume)
	ret.MShortMarginRatioByMoney = goFloat64(p.MShortMarginRatioByMoney)
	ret.MShortMarginRatioByVolume = goFloat64(p.MShortMarginRatioByVolume)
	ret.ActionDirection = byte(p.ActionDirection)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	return ret
}
func (s *CThostFtdcSyncDeltaOptExchMarginField) CValue() *C.CThostFtdcSyncDeltaOptExchMarginField {
	ptr := C.newCThostFtdcSyncDeltaOptExchMarginField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ptr.SShortMarginRatioByMoney = C.double(s.SShortMarginRatioByMoney)
	ptr.SShortMarginRatioByVolume = C.double(s.SShortMarginRatioByVolume)
	ptr.HShortMarginRatioByMoney = C.double(s.HShortMarginRatioByMoney)
	ptr.HShortMarginRatioByVolume = C.double(s.HShortMarginRatioByVolume)
	ptr.AShortMarginRatioByMoney = C.double(s.AShortMarginRatioByMoney)
	ptr.AShortMarginRatioByVolume = C.double(s.AShortMarginRatioByVolume)
	ptr.MShortMarginRatioByMoney = C.double(s.MShortMarginRatioByMoney)
	ptr.MShortMarginRatioByVolume = C.double(s.MShortMarginRatioByVolume)
	ptr.ActionDirection = C.char(s.ActionDirection)
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	return ptr
}

type CThostFtdcSyncDeltaOptInvstMarginField struct {
	InstrumentID              string
	InvestorRange             byte
	BrokerID                  string
	InvestorID                string
	SShortMarginRatioByMoney  float64
	SShortMarginRatioByVolume float64
	HShortMarginRatioByMoney  float64
	HShortMarginRatioByVolume float64
	AShortMarginRatioByMoney  float64
	AShortMarginRatioByVolume float64
	IsRelative                int
	MShortMarginRatioByMoney  float64
	MShortMarginRatioByVolume float64
	ActionDirection           byte
	SyncDeltaSequenceNo       int
}

func NewCThostFtdcSyncDeltaOptInvstMarginField(p *C.CThostFtdcSyncDeltaOptInvstMarginField) *CThostFtdcSyncDeltaOptInvstMarginField {
	ret := new(CThostFtdcSyncDeltaOptInvstMarginField)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.SShortMarginRatioByMoney = goFloat64(p.SShortMarginRatioByMoney)
	ret.SShortMarginRatioByVolume = goFloat64(p.SShortMarginRatioByVolume)
	ret.HShortMarginRatioByMoney = goFloat64(p.HShortMarginRatioByMoney)
	ret.HShortMarginRatioByVolume = goFloat64(p.HShortMarginRatioByVolume)
	ret.AShortMarginRatioByMoney = goFloat64(p.AShortMarginRatioByMoney)
	ret.AShortMarginRatioByVolume = goFloat64(p.AShortMarginRatioByVolume)
	ret.IsRelative = int(p.IsRelative)
	ret.MShortMarginRatioByMoney = goFloat64(p.MShortMarginRatioByMoney)
	ret.MShortMarginRatioByVolume = goFloat64(p.MShortMarginRatioByVolume)
	ret.ActionDirection = byte(p.ActionDirection)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	return ret
}
func (s *CThostFtdcSyncDeltaOptInvstMarginField) CValue() *C.CThostFtdcSyncDeltaOptInvstMarginField {
	ptr := C.newCThostFtdcSyncDeltaOptInvstMarginField()
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.SShortMarginRatioByMoney = C.double(s.SShortMarginRatioByMoney)
	ptr.SShortMarginRatioByVolume = C.double(s.SShortMarginRatioByVolume)
	ptr.HShortMarginRatioByMoney = C.double(s.HShortMarginRatioByMoney)
	ptr.HShortMarginRatioByVolume = C.double(s.HShortMarginRatioByVolume)
	ptr.AShortMarginRatioByMoney = C.double(s.AShortMarginRatioByMoney)
	ptr.AShortMarginRatioByVolume = C.double(s.AShortMarginRatioByVolume)
	ptr.IsRelative = C.int(s.IsRelative)
	ptr.MShortMarginRatioByMoney = C.double(s.MShortMarginRatioByMoney)
	ptr.MShortMarginRatioByVolume = C.double(s.MShortMarginRatioByVolume)
	ptr.ActionDirection = C.char(s.ActionDirection)
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	return ptr
}

type CThostFtdcSyncDeltaInvstMarginRateULField struct {
	InstrumentID             string
	InvestorRange            byte
	BrokerID                 string
	InvestorID               string
	HedgeFlag                byte
	LongMarginRatioByMoney   float64
	LongMarginRatioByVolume  float64
	ShortMarginRatioByMoney  float64
	ShortMarginRatioByVolume float64
	ActionDirection          byte
	SyncDeltaSequenceNo      int
}

func NewCThostFtdcSyncDeltaInvstMarginRateULField(p *C.CThostFtdcSyncDeltaInvstMarginRateULField) *CThostFtdcSyncDeltaInvstMarginRateULField {
	ret := new(CThostFtdcSyncDeltaInvstMarginRateULField)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.LongMarginRatioByMoney = goFloat64(p.LongMarginRatioByMoney)
	ret.LongMarginRatioByVolume = goFloat64(p.LongMarginRatioByVolume)
	ret.ShortMarginRatioByMoney = goFloat64(p.ShortMarginRatioByMoney)
	ret.ShortMarginRatioByVolume = goFloat64(p.ShortMarginRatioByVolume)
	ret.ActionDirection = byte(p.ActionDirection)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	return ret
}
func (s *CThostFtdcSyncDeltaInvstMarginRateULField) CValue() *C.CThostFtdcSyncDeltaInvstMarginRateULField {
	ptr := C.newCThostFtdcSyncDeltaInvstMarginRateULField()
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.LongMarginRatioByMoney = C.double(s.LongMarginRatioByMoney)
	ptr.LongMarginRatioByVolume = C.double(s.LongMarginRatioByVolume)
	ptr.ShortMarginRatioByMoney = C.double(s.ShortMarginRatioByMoney)
	ptr.ShortMarginRatioByVolume = C.double(s.ShortMarginRatioByVolume)
	ptr.ActionDirection = C.char(s.ActionDirection)
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	return ptr
}

type CThostFtdcSyncDeltaOptInvstCommRateField struct {
	InstrumentID            string
	InvestorRange           byte
	BrokerID                string
	InvestorID              string
	OpenRatioByMoney        float64
	OpenRatioByVolume       float64
	CloseRatioByMoney       float64
	CloseRatioByVolume      float64
	CloseTodayRatioByMoney  float64
	CloseTodayRatioByVolume float64
	StrikeRatioByMoney      float64
	StrikeRatioByVolume     float64
	ActionDirection         byte
	SyncDeltaSequenceNo     int
}

func NewCThostFtdcSyncDeltaOptInvstCommRateField(p *C.CThostFtdcSyncDeltaOptInvstCommRateField) *CThostFtdcSyncDeltaOptInvstCommRateField {
	ret := new(CThostFtdcSyncDeltaOptInvstCommRateField)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.OpenRatioByMoney = goFloat64(p.OpenRatioByMoney)
	ret.OpenRatioByVolume = goFloat64(p.OpenRatioByVolume)
	ret.CloseRatioByMoney = goFloat64(p.CloseRatioByMoney)
	ret.CloseRatioByVolume = goFloat64(p.CloseRatioByVolume)
	ret.CloseTodayRatioByMoney = goFloat64(p.CloseTodayRatioByMoney)
	ret.CloseTodayRatioByVolume = goFloat64(p.CloseTodayRatioByVolume)
	ret.StrikeRatioByMoney = goFloat64(p.StrikeRatioByMoney)
	ret.StrikeRatioByVolume = goFloat64(p.StrikeRatioByVolume)
	ret.ActionDirection = byte(p.ActionDirection)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	return ret
}
func (s *CThostFtdcSyncDeltaOptInvstCommRateField) CValue() *C.CThostFtdcSyncDeltaOptInvstCommRateField {
	ptr := C.newCThostFtdcSyncDeltaOptInvstCommRateField()
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.OpenRatioByMoney = C.double(s.OpenRatioByMoney)
	ptr.OpenRatioByVolume = C.double(s.OpenRatioByVolume)
	ptr.CloseRatioByMoney = C.double(s.CloseRatioByMoney)
	ptr.CloseRatioByVolume = C.double(s.CloseRatioByVolume)
	ptr.CloseTodayRatioByMoney = C.double(s.CloseTodayRatioByMoney)
	ptr.CloseTodayRatioByVolume = C.double(s.CloseTodayRatioByVolume)
	ptr.StrikeRatioByMoney = C.double(s.StrikeRatioByMoney)
	ptr.StrikeRatioByVolume = C.double(s.StrikeRatioByVolume)
	ptr.ActionDirection = C.char(s.ActionDirection)
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	return ptr
}

type CThostFtdcSyncDeltaInvstCommRateField struct {
	InstrumentID            string
	InvestorRange           byte
	BrokerID                string
	InvestorID              string
	OpenRatioByMoney        float64
	OpenRatioByVolume       float64
	CloseRatioByMoney       float64
	CloseRatioByVolume      float64
	CloseTodayRatioByMoney  float64
	CloseTodayRatioByVolume float64
	ActionDirection         byte
	SyncDeltaSequenceNo     int
}

func NewCThostFtdcSyncDeltaInvstCommRateField(p *C.CThostFtdcSyncDeltaInvstCommRateField) *CThostFtdcSyncDeltaInvstCommRateField {
	ret := new(CThostFtdcSyncDeltaInvstCommRateField)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.InvestorRange = byte(p.InvestorRange)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.OpenRatioByMoney = goFloat64(p.OpenRatioByMoney)
	ret.OpenRatioByVolume = goFloat64(p.OpenRatioByVolume)
	ret.CloseRatioByMoney = goFloat64(p.CloseRatioByMoney)
	ret.CloseRatioByVolume = goFloat64(p.CloseRatioByVolume)
	ret.CloseTodayRatioByMoney = goFloat64(p.CloseTodayRatioByMoney)
	ret.CloseTodayRatioByVolume = goFloat64(p.CloseTodayRatioByVolume)
	ret.ActionDirection = byte(p.ActionDirection)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	return ret
}
func (s *CThostFtdcSyncDeltaInvstCommRateField) CValue() *C.CThostFtdcSyncDeltaInvstCommRateField {
	ptr := C.newCThostFtdcSyncDeltaInvstCommRateField()
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ptr.InvestorRange = C.char(s.InvestorRange)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ptr.OpenRatioByMoney = C.double(s.OpenRatioByMoney)
	ptr.OpenRatioByVolume = C.double(s.OpenRatioByVolume)
	ptr.CloseRatioByMoney = C.double(s.CloseRatioByMoney)
	ptr.CloseRatioByVolume = C.double(s.CloseRatioByVolume)
	ptr.CloseTodayRatioByMoney = C.double(s.CloseTodayRatioByMoney)
	ptr.CloseTodayRatioByVolume = C.double(s.CloseTodayRatioByVolume)
	ptr.ActionDirection = C.char(s.ActionDirection)
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	return ptr
}

type CThostFtdcSyncDeltaProductExchRateField struct {
	ProductID           string
	QuoteCurrencyID     string
	ExchangeRate        float64
	ActionDirection     byte
	SyncDeltaSequenceNo int
}

func NewCThostFtdcSyncDeltaProductExchRateField(p *C.CThostFtdcSyncDeltaProductExchRateField) *CThostFtdcSyncDeltaProductExchRateField {
	ret := new(CThostFtdcSyncDeltaProductExchRateField)
	ret.ProductID = c2goStr(&p.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.QuoteCurrencyID = c2goStr(&p.QuoteCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ret.ExchangeRate = goFloat64(p.ExchangeRate)
	ret.ActionDirection = byte(p.ActionDirection)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	return ret
}
func (s *CThostFtdcSyncDeltaProductExchRateField) CValue() *C.CThostFtdcSyncDeltaProductExchRateField {
	ptr := C.newCThostFtdcSyncDeltaProductExchRateField()
	go2cStr(s.ProductID, &ptr.ProductID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.QuoteCurrencyID, &ptr.QuoteCurrencyID[0], C.sizeof_TThostFtdcCurrencyIDType)
	ptr.ExchangeRate = C.double(s.ExchangeRate)
	ptr.ActionDirection = C.char(s.ActionDirection)
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	return ptr
}

type CThostFtdcSyncDeltaDepthMarketDataField struct {
	TradingDay          string
	InstrumentID        string
	ExchangeID          string
	ExchangeInstID      string
	LastPrice           float64
	PreSettlementPrice  float64
	PreClosePrice       float64
	PreOpenInterest     float64
	OpenPrice           float64
	HighestPrice        float64
	LowestPrice         float64
	Volume              int
	Turnover            float64
	OpenInterest        float64
	ClosePrice          float64
	SettlementPrice     float64
	UpperLimitPrice     float64
	LowerLimitPrice     float64
	PreDelta            float64
	CurrDelta           float64
	UpdateTime          string
	UpdateMillisec      int
	BidPrice1           float64
	BidVolume1          int
	AskPrice1           float64
	AskVolume1          int
	BidPrice2           float64
	BidVolume2          int
	AskPrice2           float64
	AskVolume2          int
	BidPrice3           float64
	BidVolume3          int
	AskPrice3           float64
	AskVolume3          int
	BidPrice4           float64
	BidVolume4          int
	AskPrice4           float64
	AskVolume4          int
	BidPrice5           float64
	BidVolume5          int
	AskPrice5           float64
	AskVolume5          int
	AveragePrice        float64
	ActionDay           string
	BandingUpperPrice   float64
	BandingLowerPrice   float64
	ActionDirection     byte
	SyncDeltaSequenceNo int
}

func NewCThostFtdcSyncDeltaDepthMarketDataField(p *C.CThostFtdcSyncDeltaDepthMarketDataField) *CThostFtdcSyncDeltaDepthMarketDataField {
	ret := new(CThostFtdcSyncDeltaDepthMarketDataField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcDateType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.ExchangeInstID = c2goStr(&p.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ret.LastPrice = goFloat64(p.LastPrice)
	ret.PreSettlementPrice = goFloat64(p.PreSettlementPrice)
	ret.PreClosePrice = goFloat64(p.PreClosePrice)
	ret.PreOpenInterest = goFloat64(p.PreOpenInterest)
	ret.OpenPrice = goFloat64(p.OpenPrice)
	ret.HighestPrice = goFloat64(p.HighestPrice)
	ret.LowestPrice = goFloat64(p.LowestPrice)
	ret.Volume = int(p.Volume)
	ret.Turnover = goFloat64(p.Turnover)
	ret.OpenInterest = goFloat64(p.OpenInterest)
	ret.ClosePrice = goFloat64(p.ClosePrice)
	ret.SettlementPrice = goFloat64(p.SettlementPrice)
	ret.UpperLimitPrice = goFloat64(p.UpperLimitPrice)
	ret.LowerLimitPrice = goFloat64(p.LowerLimitPrice)
	ret.PreDelta = goFloat64(p.PreDelta)
	ret.CurrDelta = goFloat64(p.CurrDelta)
	ret.UpdateTime = c2goStr(&p.UpdateTime[0], C.sizeof_TThostFtdcTimeType)
	ret.UpdateMillisec = int(p.UpdateMillisec)
	ret.BidPrice1 = goFloat64(p.BidPrice1)
	ret.BidVolume1 = int(p.BidVolume1)
	ret.AskPrice1 = goFloat64(p.AskPrice1)
	ret.AskVolume1 = int(p.AskVolume1)
	ret.BidPrice2 = goFloat64(p.BidPrice2)
	ret.BidVolume2 = int(p.BidVolume2)
	ret.AskPrice2 = goFloat64(p.AskPrice2)
	ret.AskVolume2 = int(p.AskVolume2)
	ret.BidPrice3 = goFloat64(p.BidPrice3)
	ret.BidVolume3 = int(p.BidVolume3)
	ret.AskPrice3 = goFloat64(p.AskPrice3)
	ret.AskVolume3 = int(p.AskVolume3)
	ret.BidPrice4 = goFloat64(p.BidPrice4)
	ret.BidVolume4 = int(p.BidVolume4)
	ret.AskPrice4 = goFloat64(p.AskPrice4)
	ret.AskVolume4 = int(p.AskVolume4)
	ret.BidPrice5 = goFloat64(p.BidPrice5)
	ret.BidVolume5 = int(p.BidVolume5)
	ret.AskPrice5 = goFloat64(p.AskPrice5)
	ret.AskVolume5 = int(p.AskVolume5)
	ret.AveragePrice = goFloat64(p.AveragePrice)
	ret.ActionDay = c2goStr(&p.ActionDay[0], C.sizeof_TThostFtdcDateType)
	ret.BandingUpperPrice = goFloat64(p.BandingUpperPrice)
	ret.BandingLowerPrice = goFloat64(p.BandingLowerPrice)
	ret.ActionDirection = byte(p.ActionDirection)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	return ret
}
func (s *CThostFtdcSyncDeltaDepthMarketDataField) CValue() *C.CThostFtdcSyncDeltaDepthMarketDataField {
	ptr := C.newCThostFtdcSyncDeltaDepthMarketDataField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcDateType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.ExchangeInstID, &ptr.ExchangeInstID[0], C.sizeof_TThostFtdcExchangeInstIDType)
	ptr.LastPrice = C.double(s.LastPrice)
	ptr.PreSettlementPrice = C.double(s.PreSettlementPrice)
	ptr.PreClosePrice = C.double(s.PreClosePrice)
	ptr.PreOpenInterest = C.double(s.PreOpenInterest)
	ptr.OpenPrice = C.double(s.OpenPrice)
	ptr.HighestPrice = C.double(s.HighestPrice)
	ptr.LowestPrice = C.double(s.LowestPrice)
	ptr.Volume = C.int(s.Volume)
	ptr.Turnover = C.double(s.Turnover)
	ptr.OpenInterest = C.double(s.OpenInterest)
	ptr.ClosePrice = C.double(s.ClosePrice)
	ptr.SettlementPrice = C.double(s.SettlementPrice)
	ptr.UpperLimitPrice = C.double(s.UpperLimitPrice)
	ptr.LowerLimitPrice = C.double(s.LowerLimitPrice)
	ptr.PreDelta = C.double(s.PreDelta)
	ptr.CurrDelta = C.double(s.CurrDelta)
	go2cStr(s.UpdateTime, &ptr.UpdateTime[0], C.sizeof_TThostFtdcTimeType)
	ptr.UpdateMillisec = C.int(s.UpdateMillisec)
	ptr.BidPrice1 = C.double(s.BidPrice1)
	ptr.BidVolume1 = C.int(s.BidVolume1)
	ptr.AskPrice1 = C.double(s.AskPrice1)
	ptr.AskVolume1 = C.int(s.AskVolume1)
	ptr.BidPrice2 = C.double(s.BidPrice2)
	ptr.BidVolume2 = C.int(s.BidVolume2)
	ptr.AskPrice2 = C.double(s.AskPrice2)
	ptr.AskVolume2 = C.int(s.AskVolume2)
	ptr.BidPrice3 = C.double(s.BidPrice3)
	ptr.BidVolume3 = C.int(s.BidVolume3)
	ptr.AskPrice3 = C.double(s.AskPrice3)
	ptr.AskVolume3 = C.int(s.AskVolume3)
	ptr.BidPrice4 = C.double(s.BidPrice4)
	ptr.BidVolume4 = C.int(s.BidVolume4)
	ptr.AskPrice4 = C.double(s.AskPrice4)
	ptr.AskVolume4 = C.int(s.AskVolume4)
	ptr.BidPrice5 = C.double(s.BidPrice5)
	ptr.BidVolume5 = C.int(s.BidVolume5)
	ptr.AskPrice5 = C.double(s.AskPrice5)
	ptr.AskVolume5 = C.int(s.AskVolume5)
	ptr.AveragePrice = C.double(s.AveragePrice)
	go2cStr(s.ActionDay, &ptr.ActionDay[0], C.sizeof_TThostFtdcDateType)
	ptr.BandingUpperPrice = C.double(s.BandingUpperPrice)
	ptr.BandingLowerPrice = C.double(s.BandingLowerPrice)
	ptr.ActionDirection = C.char(s.ActionDirection)
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	return ptr
}

type CThostFtdcSyncDeltaIndexPriceField struct {
	BrokerID            string
	InstrumentID        string
	ClosePrice          float64
	ActionDirection     byte
	SyncDeltaSequenceNo int
}

func NewCThostFtdcSyncDeltaIndexPriceField(p *C.CThostFtdcSyncDeltaIndexPriceField) *CThostFtdcSyncDeltaIndexPriceField {
	ret := new(CThostFtdcSyncDeltaIndexPriceField)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.ClosePrice = goFloat64(p.ClosePrice)
	ret.ActionDirection = byte(p.ActionDirection)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	return ret
}
func (s *CThostFtdcSyncDeltaIndexPriceField) CValue() *C.CThostFtdcSyncDeltaIndexPriceField {
	ptr := C.newCThostFtdcSyncDeltaIndexPriceField()
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ptr.ClosePrice = C.double(s.ClosePrice)
	ptr.ActionDirection = C.char(s.ActionDirection)
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	return ptr
}

type CThostFtdcSyncDeltaEWarrantOffsetField struct {
	TradingDay          string
	BrokerID            string
	InvestorID          string
	ExchangeID          string
	InstrumentID        string
	Direction           byte
	HedgeFlag           byte
	Volume              int
	ActionDirection     byte
	SyncDeltaSequenceNo int
}

func NewCThostFtdcSyncDeltaEWarrantOffsetField(p *C.CThostFtdcSyncDeltaEWarrantOffsetField) *CThostFtdcSyncDeltaEWarrantOffsetField {
	ret := new(CThostFtdcSyncDeltaEWarrantOffsetField)
	ret.TradingDay = c2goStr(&p.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	ret.BrokerID = c2goStr(&p.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	ret.InvestorID = c2goStr(&p.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	ret.ExchangeID = c2goStr(&p.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	ret.InstrumentID = c2goStr(&p.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ret.Direction = byte(p.Direction)
	ret.HedgeFlag = byte(p.HedgeFlag)
	ret.Volume = int(p.Volume)
	ret.ActionDirection = byte(p.ActionDirection)
	ret.SyncDeltaSequenceNo = int(p.SyncDeltaSequenceNo)
	return ret
}
func (s *CThostFtdcSyncDeltaEWarrantOffsetField) CValue() *C.CThostFtdcSyncDeltaEWarrantOffsetField {
	ptr := C.newCThostFtdcSyncDeltaEWarrantOffsetField()
	go2cStr(s.TradingDay, &ptr.TradingDay[0], C.sizeof_TThostFtdcTradeDateType)
	go2cStr(s.BrokerID, &ptr.BrokerID[0], C.sizeof_TThostFtdcBrokerIDType)
	go2cStr(s.InvestorID, &ptr.InvestorID[0], C.sizeof_TThostFtdcInvestorIDType)
	go2cStr(s.ExchangeID, &ptr.ExchangeID[0], C.sizeof_TThostFtdcExchangeIDType)
	go2cStr(s.InstrumentID, &ptr.InstrumentID[0], C.sizeof_TThostFtdcInstrumentIDType)
	ptr.Direction = C.char(s.Direction)
	ptr.HedgeFlag = C.char(s.HedgeFlag)
	ptr.Volume = C.int(s.Volume)
	ptr.ActionDirection = C.char(s.ActionDirection)
	ptr.SyncDeltaSequenceNo = C.int(s.SyncDeltaSequenceNo)
	return ptr
}
