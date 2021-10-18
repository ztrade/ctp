#include<string.h>
#include "ctp.h"
#include "include/ThostFtdcMdApi.h"
#include "mdspi.h"

const char *mdapi_get_api_version() { return CThostFtdcMdApi::GetApiVersion(); }

MdApi mdapi_create(const char *pszFlowPath, int bIsUsingUdp,
                   int bIsMulticast) {
  CThostFtdcMdApi* mdApi = CThostFtdcMdApi::CreateFtdcMdApi(pszFlowPath, bIsMulticast, bIsMulticast);
  return mdApi;

}
void mdapi_free(MdApi a) {
  CThostFtdcMdApi* mdApi = (CThostFtdcMdApi*) a;
  mdApi->Release();
}

SpiPtr mdapi_register_spi(MdApi a, uint64_t ptr) {
  CThostFtdcMdApi* mdApi = (CThostFtdcMdApi*) a;
  MdSpi* spi = new MdSpi(ptr);
  mdApi->RegisterSpi(spi);
  return (SpiPtr)spi;
}

void mdapi_init(MdApi a) {
  CThostFtdcMdApi* mdApi = (CThostFtdcMdApi*) a;
  mdApi->Init();
}

int mdapi_join(MdApi a) {
  CThostFtdcMdApi* mdApi = (CThostFtdcMdApi*) a;
  return mdApi->Join();
}

const char *mdapi_get_trading_day(MdApi a) {
  CThostFtdcMdApi* mdApi = (CThostFtdcMdApi*) a;
  return mdApi->GetTradingDay();
}

void mdapi_register_front(MdApi a, char *pszFrontAddress) {
  CThostFtdcMdApi* mdApi = (CThostFtdcMdApi*) a;
  mdApi->RegisterFront(pszFrontAddress);
}
void mdapi_register_name_server(MdApi a, char *pszNsAddress) {
  CThostFtdcMdApi* mdApi = (CThostFtdcMdApi*) a;
  mdApi->RegisterNameServer(pszNsAddress);
}
void mdapi_register_fens_user_info(MdApi a,CThostFtdcFensUserInfoField *pFensUserInfo){
    CThostFtdcMdApi* mdApi = (CThostFtdcMdApi*) a;
    mdApi->RegisterFensUserInfo(pFensUserInfo);

}
int mdapi_subscribe_market_data(MdApi a, char *ppInstrumentID[], int nCount) {
  CThostFtdcMdApi* mdApi = (CThostFtdcMdApi*) a;
  return mdApi->SubscribeMarketData(ppInstrumentID, nCount);
}

int mdapi_unsubscribe_market_data(MdApi a, char *ppInstrumentID[], int nCount) {
  CThostFtdcMdApi* mdApi = (CThostFtdcMdApi*) a;
  return mdApi->UnSubscribeMarketData(ppInstrumentID, nCount);
}

int mdapi_subscribe_for_quote_rsp(MdApi a, char *ppInstrumentID[], int nCount) {
  CThostFtdcMdApi* mdApi = (CThostFtdcMdApi*) a;
  return mdApi->SubscribeForQuoteRsp(ppInstrumentID, nCount);
}

int mdapi_unsubscribe_for_quote_rsp(MdApi a, char *ppInstrumentID[],
                                    int nCount) {
  CThostFtdcMdApi* mdApi = (CThostFtdcMdApi*) a;
  return mdApi->UnSubscribeForQuoteRsp(ppInstrumentID, nCount);
}

int mdapi_req_user_login(MdApi a,
                         CThostFtdcReqUserLoginField *pReqUserLoginField,
                         int nRequestID) {
  CThostFtdcMdApi* mdApi = (CThostFtdcMdApi*) a;
  return mdApi->ReqUserLogin(pReqUserLoginField, nRequestID);
}

int mdapi_req_user_logout(MdApi a,
                          CThostFtdcUserLogoutField *pUserLogout,
                          int nRequestID) {
  CThostFtdcMdApi* mdApi = (CThostFtdcMdApi*) a;
  return mdApi->ReqUserLogout(pUserLogout, nRequestID);
}

int mdapi_req_qry_multicast_instrument(
    MdApi a, CThostFtdcQryMulticastInstrumentField *pQryMulticastInstrument,
    int nRequestID) {
  CThostFtdcMdApi* mdApi = (CThostFtdcMdApi*) a;
  return mdApi->ReqQryMulticastInstrument(pQryMulticastInstrument, nRequestID);
}
