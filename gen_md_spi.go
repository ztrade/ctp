// This file was automatically generated by ctpgen
package ctp

/*
#include <stdint.h>
#include "gen_types.h"
*/
import "C"

type CThostFtdcMdSpi interface {
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

func getCThostFtdcMdSpi(ptr uint64) CThostFtdcMdSpi {
	value := getGoPtr(ptr)
	if value == nil {
		return nil
	}
	v := value.(CThostFtdcMdSpi)
	return v
}

//export mdOnFrontConnected
func mdOnFrontConnected(ptr uint64) {
	p := getCThostFtdcMdSpi(ptr)
	if p != nil {
		p.OnFrontConnected()
	}
}

//export mdOnFrontDisconnected
func mdOnFrontDisconnected(ptr uint64, nReason int) {
	p := getCThostFtdcMdSpi(ptr)
	if p != nil {
		gonReason := int(nReason)
		p.OnFrontDisconnected(gonReason)
	}
}

//export mdOnHeartBeatWarning
func mdOnHeartBeatWarning(ptr uint64, nTimeLapse int) {
	p := getCThostFtdcMdSpi(ptr)
	if p != nil {
		gonTimeLapse := int(nTimeLapse)
		p.OnHeartBeatWarning(gonTimeLapse)
	}
}

//export mdOnRspUserLogin
func mdOnRspUserLogin(ptr uint64, pRspUserLogin *C.CThostFtdcRspUserLoginField, pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast C.int8_t) {
	p := getCThostFtdcMdSpi(ptr)
	if p != nil {
		gopRspUserLogin := NewCThostFtdcRspUserLoginField(pRspUserLogin)
		gopRspInfo := NewCThostFtdcRspInfoField(pRspInfo)
		gonRequestID := int(nRequestID)
		gobIsLast := c2goBool(bIsLast)
		p.OnRspUserLogin(gopRspUserLogin, gopRspInfo, gonRequestID, gobIsLast)
	}
}

//export mdOnRspUserLogout
func mdOnRspUserLogout(ptr uint64, pUserLogout *C.CThostFtdcUserLogoutField, pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast C.int8_t) {
	p := getCThostFtdcMdSpi(ptr)
	if p != nil {
		gopUserLogout := NewCThostFtdcUserLogoutField(pUserLogout)
		gopRspInfo := NewCThostFtdcRspInfoField(pRspInfo)
		gonRequestID := int(nRequestID)
		gobIsLast := c2goBool(bIsLast)
		p.OnRspUserLogout(gopUserLogout, gopRspInfo, gonRequestID, gobIsLast)
	}
}

//export mdOnRspQryMulticastInstrument
func mdOnRspQryMulticastInstrument(ptr uint64, pMulticastInstrument *C.CThostFtdcMulticastInstrumentField, pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast C.int8_t) {
	p := getCThostFtdcMdSpi(ptr)
	if p != nil {
		gopMulticastInstrument := NewCThostFtdcMulticastInstrumentField(pMulticastInstrument)
		gopRspInfo := NewCThostFtdcRspInfoField(pRspInfo)
		gonRequestID := int(nRequestID)
		gobIsLast := c2goBool(bIsLast)
		p.OnRspQryMulticastInstrument(gopMulticastInstrument, gopRspInfo, gonRequestID, gobIsLast)
	}
}

//export mdOnRspError
func mdOnRspError(ptr uint64, pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast C.int8_t) {
	p := getCThostFtdcMdSpi(ptr)
	if p != nil {
		gopRspInfo := NewCThostFtdcRspInfoField(pRspInfo)
		gonRequestID := int(nRequestID)
		gobIsLast := c2goBool(bIsLast)
		p.OnRspError(gopRspInfo, gonRequestID, gobIsLast)
	}
}

//export mdOnRspSubMarketData
func mdOnRspSubMarketData(ptr uint64, pSpecificInstrument *C.CThostFtdcSpecificInstrumentField, pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast C.int8_t) {
	p := getCThostFtdcMdSpi(ptr)
	if p != nil {
		gopSpecificInstrument := NewCThostFtdcSpecificInstrumentField(pSpecificInstrument)
		gopRspInfo := NewCThostFtdcRspInfoField(pRspInfo)
		gonRequestID := int(nRequestID)
		gobIsLast := c2goBool(bIsLast)
		p.OnRspSubMarketData(gopSpecificInstrument, gopRspInfo, gonRequestID, gobIsLast)
	}
}

//export mdOnRspUnSubMarketData
func mdOnRspUnSubMarketData(ptr uint64, pSpecificInstrument *C.CThostFtdcSpecificInstrumentField, pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast C.int8_t) {
	p := getCThostFtdcMdSpi(ptr)
	if p != nil {
		gopSpecificInstrument := NewCThostFtdcSpecificInstrumentField(pSpecificInstrument)
		gopRspInfo := NewCThostFtdcRspInfoField(pRspInfo)
		gonRequestID := int(nRequestID)
		gobIsLast := c2goBool(bIsLast)
		p.OnRspUnSubMarketData(gopSpecificInstrument, gopRspInfo, gonRequestID, gobIsLast)
	}
}

//export mdOnRspSubForQuoteRsp
func mdOnRspSubForQuoteRsp(ptr uint64, pSpecificInstrument *C.CThostFtdcSpecificInstrumentField, pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast C.int8_t) {
	p := getCThostFtdcMdSpi(ptr)
	if p != nil {
		gopSpecificInstrument := NewCThostFtdcSpecificInstrumentField(pSpecificInstrument)
		gopRspInfo := NewCThostFtdcRspInfoField(pRspInfo)
		gonRequestID := int(nRequestID)
		gobIsLast := c2goBool(bIsLast)
		p.OnRspSubForQuoteRsp(gopSpecificInstrument, gopRspInfo, gonRequestID, gobIsLast)
	}
}

//export mdOnRspUnSubForQuoteRsp
func mdOnRspUnSubForQuoteRsp(ptr uint64, pSpecificInstrument *C.CThostFtdcSpecificInstrumentField, pRspInfo *C.CThostFtdcRspInfoField, nRequestID int, bIsLast C.int8_t) {
	p := getCThostFtdcMdSpi(ptr)
	if p != nil {
		gopSpecificInstrument := NewCThostFtdcSpecificInstrumentField(pSpecificInstrument)
		gopRspInfo := NewCThostFtdcRspInfoField(pRspInfo)
		gonRequestID := int(nRequestID)
		gobIsLast := c2goBool(bIsLast)
		p.OnRspUnSubForQuoteRsp(gopSpecificInstrument, gopRspInfo, gonRequestID, gobIsLast)
	}
}

//export mdOnRtnDepthMarketData
func mdOnRtnDepthMarketData(ptr uint64, pDepthMarketData *C.CThostFtdcDepthMarketDataField) {
	p := getCThostFtdcMdSpi(ptr)
	if p != nil {
		gopDepthMarketData := NewCThostFtdcDepthMarketDataField(pDepthMarketData)
		p.OnRtnDepthMarketData(gopDepthMarketData)
	}
}

//export mdOnRtnForQuoteRsp
func mdOnRtnForQuoteRsp(ptr uint64, pForQuoteRsp *C.CThostFtdcForQuoteRspField) {
	p := getCThostFtdcMdSpi(ptr)
	if p != nil {
		gopForQuoteRsp := NewCThostFtdcForQuoteRspField(pForQuoteRsp)
		p.OnRtnForQuoteRsp(gopForQuoteRsp)
	}
}