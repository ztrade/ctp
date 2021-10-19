package ctp

// #include "ctp.h"
import "C"
import (
	"unsafe"
)

type MdApi struct {
	spi *SpiWrapper
	p   C.MdApi
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
	m.spi = &SpiWrapper{MdSpi: spi}
	ptr := uintptr(unsafe.Pointer(m.spi))
	m.spi.ptr = C.mdapi_register_spi(m.p, C.uint64_t(ptr))
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
	value := go2cStrArray(instrumentIDs)
	ppInstrumentID = &value[0]
	nCount = C.int(len(value))
	ret := C.mdapi_subscribe_market_data(m.p, ppInstrumentID, nCount)
	return int(ret)
}

func (m *MdApi) UnSubscribeMarketData(instrumentIDs []string) int {
	var ppInstrumentID **C.char
	var nCount C.int
	value := go2cStrArray(instrumentIDs)
	ppInstrumentID = &value[0]
	nCount = C.int(len(value))
	ret := C.mdapi_unsubscribe_market_data(m.p, ppInstrumentID, nCount)
	return int(ret)
}

func (m *MdApi) SubForQuoteRsp(instrumentIDs []string) int {
	var ppInstrumentID **C.char
	var nCount C.int
	value := go2cStrArray(instrumentIDs)
	ppInstrumentID = &value[0]
	nCount = C.int(len(value))
	ret := C.mdapi_subscribe_for_quote_rsp(m.p, ppInstrumentID, nCount)
	return int(ret)
}

func (m *MdApi) UnSubForQuoteRsp(instrumentIDs []string) int {
	var ppInstrumentID **C.char
	var nCount C.int
	value := go2cStrArray(instrumentIDs)
	ppInstrumentID = &value[0]
	nCount = C.int(len(value))
	ret := C.mdapi_unsubscribe_for_quote_rsp(m.p, ppInstrumentID, nCount)
	return int(ret)
}

func (m *MdApi) ReqUserLogin(reqUserLoginField *CThostFtdcReqUserLoginField, reqID int) int {
	nReqID := C.int(reqID)
	ret := C.mdapi_req_user_login(m.p, reqUserLoginField.CValue(), nReqID)
	return int(ret)
}

func (m *MdApi) ReqUserLogout(reqUserLogoutField *CThostFtdcUserLogoutField, reqID int) int {
	nReqID := C.int(reqID)
	ret := C.mdapi_req_user_logout(m.p, reqUserLogoutField.CValue(), nReqID)
	return int(ret)
}

func (m *MdApi) ReqQryMulticastInstrument(qryMulticastInstrument *CThostFtdcQryMulticastInstrumentField, reqID int) int {
	nReqID := C.int(reqID)
	ret := C.mdapi_req_qry_multicast_instrument(m.p, qryMulticastInstrument.CValue(), nReqID)
	return int(ret)
}
