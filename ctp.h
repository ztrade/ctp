#ifndef CTP_H
#define CTP_H
#include <stdint.h>
#include "include/ThostFtdcUserApiStruct.h"
#ifdef  __cplusplus
extern "C"{
#endif

typedef void *MdApi;
typedef void *SpiPtr;
typedef struct CThostFtdcRspUserLoginField CThostFtdcRspUserLoginField;
typedef struct CThostFtdcRspInfoField CThostFtdcRspInfoField;
typedef struct CThostFtdcUserLogoutField CThostFtdcUserLogoutField;
typedef struct CThostFtdcMulticastInstrumentField CThostFtdcMulticastInstrumentField;
typedef struct CThostFtdcSpecificInstrumentField CThostFtdcSpecificInstrumentField;
typedef struct CThostFtdcDepthMarketDataField CThostFtdcDepthMarketDataField;
typedef struct CThostFtdcForQuoteRspField CThostFtdcForQuoteRspField;
typedef struct CThostFtdcReqUserLoginField CThostFtdcReqUserLoginField;
typedef struct CThostFtdcUserLogoutField CThostFtdcUserLogoutField;
typedef struct CThostFtdcQryMulticastInstrumentField CThostFtdcQryMulticastInstrumentField;
typedef struct CThostFtdcFensUserInfoField CThostFtdcFensUserInfoField;
/* struct CThostFtdcReqUserLoginField; */
/* struct CThostFtdcUserLogoutField; */
/* struct CThostFtdcQryMulticastInstrumentField; */

const char *mdapi_get_api_version();


MdApi mdapi_create(const char *pszFlowPath, int bIsUsingUdp,
                      int bIsMulticast);
void mdapi_free(MdApi a);

SpiPtr mdapi_register_spi(MdApi a,uint64_t ptr);
void mdapi_init(MdApi a);

int mdapi_join(MdApi a);
const char *mdapi_get_trading_day(MdApi a);
void mdapi_register_front(MdApi a, char *pszFrontAddress);
void mdapi_register_name_server(MdApi a, char *pszNsAddress);
void mdapi_register_fens_user_info(MdApi a,CThostFtdcFensUserInfoField *pFensUserInfo);
int mdapi_subscribe_market_data(MdApi a, char *ppInstrumentID[], int nCount);
int mdapi_unsubscribe_market_data(MdApi a, char *ppInstrumentID[], int nCount);
int mdapi_subscribe_for_quote_rsp(MdApi a, char *ppInstrumentID[], int nCount);
int mdapi_unsubscribe_for_quote_rsp(MdApi a, char *ppInstrumentID[], int nCount);
int mdapi_req_user_login(MdApi a, CThostFtdcReqUserLoginField *pReqUserLoginField, int
nRequestID);
int mdapi_req_user_logout(MdApi a, CThostFtdcUserLogoutField *pReqUserLoginField,
int nRequestID);
int mdapi_req_qry_multicast_instrument(MdApi a, CThostFtdcQryMulticastInstrumentField *pQryMulticastInstrument, int nRequestID);


#ifdef  __cplusplus
};
#endif


#endif /* CTP_H */
