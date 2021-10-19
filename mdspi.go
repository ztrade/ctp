package ctp

/*
 #include "include/ThostFtdcUserApiStruct.h"
 #include "ctp.h"
 #cgo linux amd64 LDFLAGS: -fPIC -L${SRCDIR}/libs/linux64/ -l:thostmduserapi_se.so -l:thosttraderapi_se.so -lstdc++
*/
import "C"
import (
	"unsafe"
)

type MdSpi interface {
	OnFrontConnected()
	OnFrontDisconnected(nReason int)
	OnHeartBeatWarning(nTimeLapse int)
	OnRspUserLogin(pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	OnRspUserLogout(pUserLogout *CThostFtdcUserLogoutField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	OnRspQryMulticastInstrument(pMulticastInstrument *CThostFtdcMulticastInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	OnRspError(pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	OnRspSubMarketData(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	OnRspUnSubMarketData(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	OnRspSubForQuoteRsp(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	OnRspUnSubForQuoteRsp(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	OnRtnDepthMarketData(pDepthMarketData *CThostFtdcDepthMarketDataField)
	OnRtnForQuoteRsp(pForQuoteRsp *CThostFtdcForQuoteRspField)
}

type SpiWrapper struct {
	Name string
	MdSpi
	ptr C.SpiPtr
}

func getMdSpi(ptr uint64) MdSpi {
	var v *SpiWrapper
	v = (*SpiWrapper)(unsafe.Pointer(uintptr(ptr)))
	return v.MdSpi
}

//export goMdOnFrontConnected
func goMdOnFrontConnected(ptr uint64) {
	getMdSpi(ptr).OnFrontConnected()
}

//export goMdOnFrontDisconnected
func goMdOnFrontDisconnected(ptr uint64, nReason int) {
	getMdSpi(ptr).OnFrontDisconnected(nReason)
}

//export goMdOnHeartBeatWarning
func goMdOnHeartBeatWarning(ptr uint64, nTimeLapse int) {
	getMdSpi(ptr).OnHeartBeatWarning(nTimeLapse)
}

//export goMdOnRspUserLogin
func goMdOnRspUserLogin(ptr uint64, pRspUserLogin *C.CThostFtdcRspUserLoginField, pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	rspUserLogin := NewCThostFtdcRspUserLoginField(*pRspUserLogin)
	rspInfo := NewCThostFtdcRspInfoField(pRspInfo)
	getMdSpi(ptr).OnRspUserLogin(rspUserLogin, rspInfo, nRequestID, bIsLast)
}

//export goMdOnRspUserLogout
func goMdOnRspUserLogout(ptr uint64, pUserLogout *C.CThostFtdcUserLogoutField, pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	rspUserLogout := NewCThostFtdcUserLogoutField(pUserLogout)
	rspInfo := NewCThostFtdcRspInfoField(pRspInfo)
	getMdSpi(ptr).OnRspUserLogout(rspUserLogout, rspInfo, nRequestID, bIsLast)
}

//export goMdOnRspQryMulticastInstrument
func goMdOnRspQryMulticastInstrument(ptr uint64, pMulticastInstrument *C.CThostFtdcMulticastInstrumentField, pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	multicastInstrument := NewCThostFtdcMulticastInstrumentField(pMulticastInstrument)
	rspInfo := NewCThostFtdcRspInfoField(pRspInfo)
	getMdSpi(ptr).OnRspQryMulticastInstrument(multicastInstrument, rspInfo, nRequestID, bIsLast)
}

//export goMdOnRspError
func goMdOnRspError(ptr uint64, pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	rspInfo := NewCThostFtdcRspInfoField(pRspInfo)
	getMdSpi(ptr).OnRspError(rspInfo, nRequestID, bIsLast)
}

//export goMdOnRspSubMarketData
func goMdOnRspSubMarketData(ptr uint64, pSpecificInstrument *C.CThostFtdcSpecificInstrumentField, pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	specificInstrument := NewCThostFtdcSpecificInstrumentField(pSpecificInstrument)
	rspInfo := NewCThostFtdcRspInfoField(pRspInfo)
	getMdSpi(ptr).OnRspSubMarketData(specificInstrument, rspInfo, nRequestID, bIsLast)
}

//export goMdOnRspUnSubMarketData
func goMdOnRspUnSubMarketData(ptr uint64, pSpecificInstrument *C.CThostFtdcSpecificInstrumentField, pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	specificInstrument := NewCThostFtdcSpecificInstrumentField(pSpecificInstrument)
	rspInfo := NewCThostFtdcRspInfoField(pRspInfo)
	getMdSpi(ptr).OnRspUnSubMarketData(specificInstrument, rspInfo, nRequestID, bIsLast)
}

//export goMdOnRspSubForQuoteRsp
func goMdOnRspSubForQuoteRsp(ptr uint64, pSpecificInstrument *C.CThostFtdcSpecificInstrumentField, pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	specificInstrument := NewCThostFtdcSpecificInstrumentField(pSpecificInstrument)
	rspInfo := NewCThostFtdcRspInfoField(pRspInfo)
	getMdSpi(ptr).OnRspSubForQuoteRsp(specificInstrument, rspInfo, nRequestID, bIsLast)
}

//export goMdOnRspUnSubForQuoteRsp
func goMdOnRspUnSubForQuoteRsp(ptr uint64, pSpecificInstrument *C.CThostFtdcSpecificInstrumentField, pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	specificInstrument := NewCThostFtdcSpecificInstrumentField(pSpecificInstrument)
	rspInfo := NewCThostFtdcRspInfoField(pRspInfo)
	getMdSpi(ptr).OnRspUnSubForQuoteRsp(specificInstrument, rspInfo, nRequestID, bIsLast)
}

//export goMdOnRtnDepthMarketData
func goMdOnRtnDepthMarketData(ptr uint64, pDepthMarketData *C.CThostFtdcDepthMarketDataField) {
	depthMarketData := NewCThostFtdcDepthMarketDataField(pDepthMarketData)
	getMdSpi(ptr).OnRtnDepthMarketData(depthMarketData)
}

//export goMdOnRtnForQuoteRsp
func goMdOnRtnForQuoteRsp(ptr uint64, pForQuoteRsp *C.CThostFtdcForQuoteRspField) {
	forQuoteRsp := NewCThostFtdcForQuoteRspField(pForQuoteRsp)
	getMdSpi(ptr).OnRtnForQuoteRsp(forQuoteRsp)
}
