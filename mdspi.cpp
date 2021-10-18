#include "mdspi.h"

extern "C" {
  void goMdOnFrontConnected(uint64_t ptr);
  void goMdOnFrontDisconnected(uint64_t ptr, int nReason);
  void goMdOnHeartBeatWarning(uint64_t ptr, int nTimeLapse);
  void goMdOnRspUserLogin(uint64_t ptr, CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
  void goMdOnRspUserLogout(uint64_t ptr, CThostFtdcUserLogoutField *pUserLogout, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

  void goMdOnRspQryMulticastInstrument(uint64_t ptr, CThostFtdcMulticastInstrumentField *pMulticastInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
  void goMdOnRspError(uint64_t ptr, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
  void goMdOnRspSubMarketData(uint64_t ptr, CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
  void goMdOnRspUnSubMarketData(uint64_t ptr,CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
  void goMdOnRspSubForQuoteRsp(uint64_t ptr,CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
  void goMdOnRspUnSubForQuoteRsp(uint64_t ptr,CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
  void goMdOnRtnDepthMarketData(uint64_t ptr,CThostFtdcDepthMarketDataField *pDepthMarketData);
  void goMdOnRtnForQuoteRsp(uint64_t ptr, CThostFtdcForQuoteRspField *pForQuoteRsp);
};


MdSpi::MdSpi(uint64_t ptr):ptr(ptr) {
}

MdSpi::~MdSpi() {}

void MdSpi::OnFrontConnected() {
  goMdOnFrontConnected(ptr);
}

///当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
///@param nReason 错误原因
///        0x1001 网络读失败
///        0x1002 网络写失败
///        0x2001 接收心跳超时
///        0x2002 发送心跳失败
///        0x2003 收到错误报文
void MdSpi::OnFrontDisconnected(int nReason) {
  goMdOnFrontDisconnected(ptr, nReason);
}

///心跳超时警告。当长时间未收到报文时，该方法被调用。
///@param nTimeLapse 距离上次接收报文的时间
void MdSpi::OnHeartBeatWarning(int nTimeLapse) {
  goMdOnHeartBeatWarning(ptr, nTimeLapse);
}


///登录请求响应
void MdSpi::OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin,
                           CThostFtdcRspInfoField *pRspInfo, int nRequestID,
                           bool bIsLast) {
  goMdOnRspUserLogin(ptr, pRspUserLogin, pRspInfo, nRequestID, bIsLast);
}


///登出请求响应
void MdSpi::OnRspUserLogout(CThostFtdcUserLogoutField *pUserLogout,
                            CThostFtdcRspInfoField *pRspInfo, int nRequestID,
                            bool bIsLast) {

  goMdOnRspUserLogout(ptr, pUserLogout, pRspInfo, nRequestID, bIsLast);
}

///请求查询组播合约响应
void MdSpi::OnRspQryMulticastInstrument(
    CThostFtdcMulticastInstrumentField *pMulticastInstrument,
    CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {

  goMdOnRspQryMulticastInstrument(ptr, pMulticastInstrument, pRspInfo, nRequestID, bIsLast);
}

///错误应答
void MdSpi::OnRspError(CThostFtdcRspInfoField *pRspInfo, int nRequestID,
                       bool bIsLast) {
  goMdOnRspError(ptr, pRspInfo, nRequestID, bIsLast);
}

///订阅行情应答
void MdSpi::OnRspSubMarketData(
    CThostFtdcSpecificInstrumentField *pSpecificInstrument,
    CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
   goMdOnRspSubMarketData(ptr, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
}

///取消订阅行情应答
void MdSpi::OnRspUnSubMarketData(
    CThostFtdcSpecificInstrumentField *pSpecificInstrument,
    CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
  goMdOnRspUnSubMarketData(ptr, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
}

///订阅询价应答
void MdSpi::OnRspSubForQuoteRsp(
    CThostFtdcSpecificInstrumentField *pSpecificInstrument,
    CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
  goMdOnRspSubForQuoteRsp(ptr, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
}

///取消订阅询价应答
void MdSpi::OnRspUnSubForQuoteRsp(
    CThostFtdcSpecificInstrumentField *pSpecificInstrument,
    CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
  goMdOnRspUnSubForQuoteRsp(ptr, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
}

///深度行情通知
void MdSpi::OnRtnDepthMarketData(
    CThostFtdcDepthMarketDataField *pDepthMarketData) {
  goMdOnRtnDepthMarketData(ptr, pDepthMarketData);
}

///询价通知
void MdSpi::OnRtnForQuoteRsp(CThostFtdcForQuoteRspField *pForQuoteRsp) {
  goMdOnRtnForQuoteRsp(ptr, pForQuoteRsp);
}
