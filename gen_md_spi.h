// This file was automatically generated by ctpgen
#ifndef GEN_MD_SPI_H
#define GEN_MD_SPI_H
#include <stdint.h>
#include "include/ThostFtdcMdApi.h"
class CThostFtdcMdSpiImpl : public CThostFtdcMdSpi
{
public:
  CThostFtdcMdSpiImpl(uint64_t ptr);
  virtual ~CThostFtdcMdSpiImpl();

  virtual void OnFrontConnected();

  virtual void OnFrontDisconnected(int nReason);

  virtual void OnHeartBeatWarning(int nTimeLapse);

  virtual void OnRspUserLogin(CThostFtdcRspUserLoginField * pRspUserLogin,CThostFtdcRspInfoField * pRspInfo,int nRequestID,bool bIsLast);

  virtual void OnRspUserLogout(CThostFtdcUserLogoutField * pUserLogout,CThostFtdcRspInfoField * pRspInfo,int nRequestID,bool bIsLast);

  virtual void OnRspQryMulticastInstrument(CThostFtdcMulticastInstrumentField * pMulticastInstrument,CThostFtdcRspInfoField * pRspInfo,int nRequestID,bool bIsLast);

  virtual void OnRspError(CThostFtdcRspInfoField * pRspInfo,int nRequestID,bool bIsLast);

  virtual void OnRspSubMarketData(CThostFtdcSpecificInstrumentField * pSpecificInstrument,CThostFtdcRspInfoField * pRspInfo,int nRequestID,bool bIsLast);

  virtual void OnRspUnSubMarketData(CThostFtdcSpecificInstrumentField * pSpecificInstrument,CThostFtdcRspInfoField * pRspInfo,int nRequestID,bool bIsLast);

  virtual void OnRspSubForQuoteRsp(CThostFtdcSpecificInstrumentField * pSpecificInstrument,CThostFtdcRspInfoField * pRspInfo,int nRequestID,bool bIsLast);

  virtual void OnRspUnSubForQuoteRsp(CThostFtdcSpecificInstrumentField * pSpecificInstrument,CThostFtdcRspInfoField * pRspInfo,int nRequestID,bool bIsLast);

  virtual void OnRtnDepthMarketData(CThostFtdcDepthMarketDataField * pDepthMarketData);

  virtual void OnRtnForQuoteRsp(CThostFtdcForQuoteRspField * pForQuoteRsp);


private:
  uint64_t ptr = 0;
};
#endif  // GEN_MD_SPI_H