package ctp

// #include "ctp.h"
import "C"
import (
	"unsafe"
)

type MdApi struct {
	p C.MdApi
}

func NewMdApi(flowPath string, bIsUsingUdp, bIsMulticast bool) *MdApi {
	md := new(MdApi)
	cPath := C.CString(flowPath)
	var cIsUsingUdp, cIsMulticast C.int
	if bIsMulticast {
		cIsUsingUdp = 1
	}
	if bIsMulticast {
		cIsMulticast = 1
	}
	md.p = C.mdapi_create(cPath, cIsUsingUdp, cIsMulticast)
	return md
}

func GetMdApiVersion() string {
	ver := C.mdapi_get_api_version()
	return C.GoString(ver)
}

func (m *MdApi) RegisterSpi(spi MdSpi) {
	ptr := uintptr(unsafe.Pointer(&spi))
	C.mdapi_register_spi(m.p, C.int(ptr))
}

func (m *MdApi) Init() {
	C.mdapi_init(m.p)
}

func (m *MdApi) Join() int {
	ret := C.mdapi_join(m.p)
	return int(ret)
}

func (m *MdApi) GetTradingDay() string {
	ret := C.mdapi_get_trading_day(m.p)
	return C.GoString(ret)
}

func (m *MdApi) RegisterFront(front string) {
	cFront := C.CString(front)
	C.mdapi_register_front(m.p, cFront)
}

func (m *MdApi) RegisterNameServer(ns string) {
	cNs := C.CString(ns)
	C.mdapi_register_name_server(m.p, cNs)
}

func (m *MdApi) RegisterFensUserInfo(pFensUserInfo *CThostFtdcFensUserInfoField) {
	var fenUserInfo *C.CThostFtdcFensUserInfoField
	C.mdapi_register_fens_user_info(m.p, fenUserInfo)
}

func (m *MdApi) SubscribeMarketData(instrumentIDs []string) int {
	var ppInstrumentID **C.char
	var nCount C.int
	ret := C.mdapi_subscribe_market_data(m.p, ppInstrumentID, nCount)
	return int(ret)
}

func (m *MdApi) UnSubscribeMarketData(instrumentIDs []string) int {
	var ppInstrumentID **C.char
	var nCount C.int
	ret := C.mdapi_unsubscribe_market_data(m.p, ppInstrumentID, nCount)
	return int(ret)
}

func (m *MdApi) SubForQuoteRsp(instrumentIDs []string) int {
	var ppInstrumentID **C.char
	var nCount C.int
	ret := C.mdapi_subscribe_for_quote_rsp(m.p, ppInstrumentID, nCount)
	return int(ret)
}

func (m *MdApi) UnSubForQuoteRsp(instrumentIDs []string) int {
	var ppInstrumentID **C.char
	var nCount C.int
	ret := C.mdapi_unsubscribe_for_quote_rsp(m.p, ppInstrumentID, nCount)
	return int(ret)
}

func (m *MdApi) ReqUserLogin(reqUserLoginField *CThostFtdcReqUserLoginField, reqID int) int {
	var pReqUserLoginField *C.CThostFtdcReqUserLoginField
	var nReqID C.int
	ret := C.mdapi_req_user_login(m.p, pReqUserLoginField, nReqID)
	return int(ret)
}

func (m *MdApi) ReqUserLogout(reqUserLogoutField *CThostFtdcUserLogoutField, reqID int) int {
	var pReqUserLogoutField *C.CThostFtdcUserLogoutField
	var nReqID C.int
	ret := C.mdapi_req_user_logout(m.p, pReqUserLogoutField, nReqID)
	return int(ret)
}

func (m *MdApi) ReqQryMulticastInstrument(qryMulticastInstrument *CThostFtdcQryMulticastInstrumentField, reqID int) int {
	var pQryMulticastInstrument *C.CThostFtdcQryMulticastInstrumentField
	var nReqID C.int
	ret := C.mdapi_req_qry_multicast_instrument(m.p, pQryMulticastInstrument, nReqID)
	return int(ret)
}
