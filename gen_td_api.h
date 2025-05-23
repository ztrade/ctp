// This file was automatically generated by ctpgen
#ifndef GEN_TD_API_H
#define GEN_TD_API_H
#include <stdint.h>
#include "gen_types.h"
#ifdef  __cplusplus
extern "C"{
#endif

typedef void* tdApi;
typedef void* tdSpi;

tdSpi td_new_spi(uint64_t value);
void td_spi_free(tdSpi p);


tdApi td_create_ftdc_trader_api(char * pszFlowPath);
const char * td_get_api_version();

void td_release(tdApi a);
void td_init(tdApi a);
int td_join(tdApi a);
const char * td_get_trading_day(tdApi a);
void td_get_front_info(tdApi a, CThostFtdcFrontInfoField * pFrontInfo);
void td_register_front(tdApi a, char * pszFrontAddress);
void td_register_name_server(tdApi a, char * pszNsAddress);
void td_register_fens_user_info(tdApi a, CThostFtdcFensUserInfoField * pFensUserInfo);
void td_register_spi(tdApi a, tdSpi s);
void td_subscribe_private_topic(tdApi a, THOST_TE_RESUME_TYPE nResumeType);
void td_subscribe_public_topic(tdApi a, THOST_TE_RESUME_TYPE nResumeType);
int td_req_authenticate(tdApi a, CThostFtdcReqAuthenticateField * pReqAuthenticateField, int nRequestID);
int td_register_user_system_info(tdApi a, CThostFtdcUserSystemInfoField * pUserSystemInfo);
int td_submit_user_system_info(tdApi a, CThostFtdcUserSystemInfoField * pUserSystemInfo);
int td_req_user_login(tdApi a, CThostFtdcReqUserLoginField * pReqUserLoginField, int nRequestID);
int td_req_user_logout(tdApi a, CThostFtdcUserLogoutField * pUserLogout, int nRequestID);
int td_req_user_password_update(tdApi a, CThostFtdcUserPasswordUpdateField * pUserPasswordUpdate, int nRequestID);
int td_req_trading_account_password_update(tdApi a, CThostFtdcTradingAccountPasswordUpdateField * pTradingAccountPasswordUpdate, int nRequestID);
int td_req_user_auth_method(tdApi a, CThostFtdcReqUserAuthMethodField * pReqUserAuthMethod, int nRequestID);
int td_req_gen_user_captcha(tdApi a, CThostFtdcReqGenUserCaptchaField * pReqGenUserCaptcha, int nRequestID);
int td_req_gen_user_text(tdApi a, CThostFtdcReqGenUserTextField * pReqGenUserText, int nRequestID);
int td_req_user_login_with_captcha(tdApi a, CThostFtdcReqUserLoginWithCaptchaField * pReqUserLoginWithCaptcha, int nRequestID);
int td_req_user_login_with_text(tdApi a, CThostFtdcReqUserLoginWithTextField * pReqUserLoginWithText, int nRequestID);
int td_req_user_login_with_o_t_p(tdApi a, CThostFtdcReqUserLoginWithOTPField * pReqUserLoginWithOTP, int nRequestID);
int td_req_order_insert(tdApi a, CThostFtdcInputOrderField * pInputOrder, int nRequestID);
int td_req_parked_order_insert(tdApi a, CThostFtdcParkedOrderField * pParkedOrder, int nRequestID);
int td_req_parked_order_action(tdApi a, CThostFtdcParkedOrderActionField * pParkedOrderAction, int nRequestID);
int td_req_order_action(tdApi a, CThostFtdcInputOrderActionField * pInputOrderAction, int nRequestID);
int td_req_qry_max_order_volume(tdApi a, CThostFtdcQryMaxOrderVolumeField * pQryMaxOrderVolume, int nRequestID);
int td_req_settlement_info_confirm(tdApi a, CThostFtdcSettlementInfoConfirmField * pSettlementInfoConfirm, int nRequestID);
int td_req_remove_parked_order(tdApi a, CThostFtdcRemoveParkedOrderField * pRemoveParkedOrder, int nRequestID);
int td_req_remove_parked_order_action(tdApi a, CThostFtdcRemoveParkedOrderActionField * pRemoveParkedOrderAction, int nRequestID);
int td_req_exec_order_insert(tdApi a, CThostFtdcInputExecOrderField * pInputExecOrder, int nRequestID);
int td_req_exec_order_action(tdApi a, CThostFtdcInputExecOrderActionField * pInputExecOrderAction, int nRequestID);
int td_req_for_quote_insert(tdApi a, CThostFtdcInputForQuoteField * pInputForQuote, int nRequestID);
int td_req_quote_insert(tdApi a, CThostFtdcInputQuoteField * pInputQuote, int nRequestID);
int td_req_quote_action(tdApi a, CThostFtdcInputQuoteActionField * pInputQuoteAction, int nRequestID);
int td_req_batch_order_action(tdApi a, CThostFtdcInputBatchOrderActionField * pInputBatchOrderAction, int nRequestID);
int td_req_option_self_close_insert(tdApi a, CThostFtdcInputOptionSelfCloseField * pInputOptionSelfClose, int nRequestID);
int td_req_option_self_close_action(tdApi a, CThostFtdcInputOptionSelfCloseActionField * pInputOptionSelfCloseAction, int nRequestID);
int td_req_comb_action_insert(tdApi a, CThostFtdcInputCombActionField * pInputCombAction, int nRequestID);
int td_req_qry_order(tdApi a, CThostFtdcQryOrderField * pQryOrder, int nRequestID);
int td_req_qry_trade(tdApi a, CThostFtdcQryTradeField * pQryTrade, int nRequestID);
int td_req_qry_investor_position(tdApi a, CThostFtdcQryInvestorPositionField * pQryInvestorPosition, int nRequestID);
int td_req_qry_trading_account(tdApi a, CThostFtdcQryTradingAccountField * pQryTradingAccount, int nRequestID);
int td_req_qry_investor(tdApi a, CThostFtdcQryInvestorField * pQryInvestor, int nRequestID);
int td_req_qry_trading_code(tdApi a, CThostFtdcQryTradingCodeField * pQryTradingCode, int nRequestID);
int td_req_qry_instrument_margin_rate(tdApi a, CThostFtdcQryInstrumentMarginRateField * pQryInstrumentMarginRate, int nRequestID);
int td_req_qry_instrument_commission_rate(tdApi a, CThostFtdcQryInstrumentCommissionRateField * pQryInstrumentCommissionRate, int nRequestID);
int td_req_qry_exchange(tdApi a, CThostFtdcQryExchangeField * pQryExchange, int nRequestID);
int td_req_qry_product(tdApi a, CThostFtdcQryProductField * pQryProduct, int nRequestID);
int td_req_qry_instrument(tdApi a, CThostFtdcQryInstrumentField * pQryInstrument, int nRequestID);
int td_req_qry_depth_market_data(tdApi a, CThostFtdcQryDepthMarketDataField * pQryDepthMarketData, int nRequestID);
int td_req_qry_trader_offer(tdApi a, CThostFtdcQryTraderOfferField * pQryTraderOffer, int nRequestID);
int td_req_qry_settlement_info(tdApi a, CThostFtdcQrySettlementInfoField * pQrySettlementInfo, int nRequestID);
int td_req_qry_transfer_bank(tdApi a, CThostFtdcQryTransferBankField * pQryTransferBank, int nRequestID);
int td_req_qry_investor_position_detail(tdApi a, CThostFtdcQryInvestorPositionDetailField * pQryInvestorPositionDetail, int nRequestID);
int td_req_qry_notice(tdApi a, CThostFtdcQryNoticeField * pQryNotice, int nRequestID);
int td_req_qry_settlement_info_confirm(tdApi a, CThostFtdcQrySettlementInfoConfirmField * pQrySettlementInfoConfirm, int nRequestID);
int td_req_qry_investor_position_combine_detail(tdApi a, CThostFtdcQryInvestorPositionCombineDetailField * pQryInvestorPositionCombineDetail, int nRequestID);
int td_req_qry_c_f_m_m_c_trading_account_key(tdApi a, CThostFtdcQryCFMMCTradingAccountKeyField * pQryCFMMCTradingAccountKey, int nRequestID);
int td_req_qry_e_warrant_offset(tdApi a, CThostFtdcQryEWarrantOffsetField * pQryEWarrantOffset, int nRequestID);
int td_req_qry_investor_product_group_margin(tdApi a, CThostFtdcQryInvestorProductGroupMarginField * pQryInvestorProductGroupMargin, int nRequestID);
int td_req_qry_exchange_margin_rate(tdApi a, CThostFtdcQryExchangeMarginRateField * pQryExchangeMarginRate, int nRequestID);
int td_req_qry_exchange_margin_rate_adjust(tdApi a, CThostFtdcQryExchangeMarginRateAdjustField * pQryExchangeMarginRateAdjust, int nRequestID);
int td_req_qry_exchange_rate(tdApi a, CThostFtdcQryExchangeRateField * pQryExchangeRate, int nRequestID);
int td_req_qry_sec_agent_a_c_i_d_map(tdApi a, CThostFtdcQrySecAgentACIDMapField * pQrySecAgentACIDMap, int nRequestID);
int td_req_qry_product_exch_rate(tdApi a, CThostFtdcQryProductExchRateField * pQryProductExchRate, int nRequestID);
int td_req_qry_product_group(tdApi a, CThostFtdcQryProductGroupField * pQryProductGroup, int nRequestID);
int td_req_qry_m_m_instrument_commission_rate(tdApi a, CThostFtdcQryMMInstrumentCommissionRateField * pQryMMInstrumentCommissionRate, int nRequestID);
int td_req_qry_m_m_option_instr_comm_rate(tdApi a, CThostFtdcQryMMOptionInstrCommRateField * pQryMMOptionInstrCommRate, int nRequestID);
int td_req_qry_instrument_order_comm_rate(tdApi a, CThostFtdcQryInstrumentOrderCommRateField * pQryInstrumentOrderCommRate, int nRequestID);
int td_req_qry_sec_agent_trading_account(tdApi a, CThostFtdcQryTradingAccountField * pQryTradingAccount, int nRequestID);
int td_req_qry_sec_agent_check_mode(tdApi a, CThostFtdcQrySecAgentCheckModeField * pQrySecAgentCheckMode, int nRequestID);
int td_req_qry_sec_agent_trade_info(tdApi a, CThostFtdcQrySecAgentTradeInfoField * pQrySecAgentTradeInfo, int nRequestID);
int td_req_qry_option_instr_trade_cost(tdApi a, CThostFtdcQryOptionInstrTradeCostField * pQryOptionInstrTradeCost, int nRequestID);
int td_req_qry_option_instr_comm_rate(tdApi a, CThostFtdcQryOptionInstrCommRateField * pQryOptionInstrCommRate, int nRequestID);
int td_req_qry_exec_order(tdApi a, CThostFtdcQryExecOrderField * pQryExecOrder, int nRequestID);
int td_req_qry_for_quote(tdApi a, CThostFtdcQryForQuoteField * pQryForQuote, int nRequestID);
int td_req_qry_quote(tdApi a, CThostFtdcQryQuoteField * pQryQuote, int nRequestID);
int td_req_qry_option_self_close(tdApi a, CThostFtdcQryOptionSelfCloseField * pQryOptionSelfClose, int nRequestID);
int td_req_qry_invest_unit(tdApi a, CThostFtdcQryInvestUnitField * pQryInvestUnit, int nRequestID);
int td_req_qry_comb_instrument_guard(tdApi a, CThostFtdcQryCombInstrumentGuardField * pQryCombInstrumentGuard, int nRequestID);
int td_req_qry_comb_action(tdApi a, CThostFtdcQryCombActionField * pQryCombAction, int nRequestID);
int td_req_qry_transfer_serial(tdApi a, CThostFtdcQryTransferSerialField * pQryTransferSerial, int nRequestID);
int td_req_qry_accountregister(tdApi a, CThostFtdcQryAccountregisterField * pQryAccountregister, int nRequestID);
int td_req_qry_contract_bank(tdApi a, CThostFtdcQryContractBankField * pQryContractBank, int nRequestID);
int td_req_qry_parked_order(tdApi a, CThostFtdcQryParkedOrderField * pQryParkedOrder, int nRequestID);
int td_req_qry_parked_order_action(tdApi a, CThostFtdcQryParkedOrderActionField * pQryParkedOrderAction, int nRequestID);
int td_req_qry_trading_notice(tdApi a, CThostFtdcQryTradingNoticeField * pQryTradingNotice, int nRequestID);
int td_req_qry_broker_trading_params(tdApi a, CThostFtdcQryBrokerTradingParamsField * pQryBrokerTradingParams, int nRequestID);
int td_req_qry_broker_trading_algos(tdApi a, CThostFtdcQryBrokerTradingAlgosField * pQryBrokerTradingAlgos, int nRequestID);
int td_req_query_c_f_m_m_c_trading_account_token(tdApi a, CThostFtdcQueryCFMMCTradingAccountTokenField * pQueryCFMMCTradingAccountToken, int nRequestID);
int td_req_from_bank_to_future_by_future(tdApi a, CThostFtdcReqTransferField * pReqTransfer, int nRequestID);
int td_req_from_future_to_bank_by_future(tdApi a, CThostFtdcReqTransferField * pReqTransfer, int nRequestID);
int td_req_query_bank_account_money_by_future(tdApi a, CThostFtdcReqQueryAccountField * pReqQueryAccount, int nRequestID);
int td_req_qry_classified_instrument(tdApi a, CThostFtdcQryClassifiedInstrumentField * pQryClassifiedInstrument, int nRequestID);
int td_req_qry_comb_promotion_param(tdApi a, CThostFtdcQryCombPromotionParamField * pQryCombPromotionParam, int nRequestID);
int td_req_qry_risk_settle_invst_position(tdApi a, CThostFtdcQryRiskSettleInvstPositionField * pQryRiskSettleInvstPosition, int nRequestID);
int td_req_qry_risk_settle_product_status(tdApi a, CThostFtdcQryRiskSettleProductStatusField * pQryRiskSettleProductStatus, int nRequestID);
int td_req_qry_s_p_b_m_future_parameter(tdApi a, CThostFtdcQrySPBMFutureParameterField * pQrySPBMFutureParameter, int nRequestID);
int td_req_qry_s_p_b_m_option_parameter(tdApi a, CThostFtdcQrySPBMOptionParameterField * pQrySPBMOptionParameter, int nRequestID);
int td_req_qry_s_p_b_m_intra_parameter(tdApi a, CThostFtdcQrySPBMIntraParameterField * pQrySPBMIntraParameter, int nRequestID);
int td_req_qry_s_p_b_m_inter_parameter(tdApi a, CThostFtdcQrySPBMInterParameterField * pQrySPBMInterParameter, int nRequestID);
int td_req_qry_s_p_b_m_portf_definition(tdApi a, CThostFtdcQrySPBMPortfDefinitionField * pQrySPBMPortfDefinition, int nRequestID);
int td_req_qry_s_p_b_m_investor_portf_def(tdApi a, CThostFtdcQrySPBMInvestorPortfDefField * pQrySPBMInvestorPortfDef, int nRequestID);
int td_req_qry_investor_portf_margin_ratio(tdApi a, CThostFtdcQryInvestorPortfMarginRatioField * pQryInvestorPortfMarginRatio, int nRequestID);
int td_req_qry_investor_prod_s_p_b_m_detail(tdApi a, CThostFtdcQryInvestorProdSPBMDetailField * pQryInvestorProdSPBMDetail, int nRequestID);
int td_req_qry_investor_commodity_s_p_m_m_margin(tdApi a, CThostFtdcQryInvestorCommoditySPMMMarginField * pQryInvestorCommoditySPMMMargin, int nRequestID);
int td_req_qry_investor_commodity_group_s_p_m_m_margin(tdApi a, CThostFtdcQryInvestorCommodityGroupSPMMMarginField * pQryInvestorCommodityGroupSPMMMargin, int nRequestID);
int td_req_qry_s_p_m_m_inst_param(tdApi a, CThostFtdcQrySPMMInstParamField * pQrySPMMInstParam, int nRequestID);
int td_req_qry_s_p_m_m_product_param(tdApi a, CThostFtdcQrySPMMProductParamField * pQrySPMMProductParam, int nRequestID);
int td_req_qry_s_p_b_m_add_on_inter_parameter(tdApi a, CThostFtdcQrySPBMAddOnInterParameterField * pQrySPBMAddOnInterParameter, int nRequestID);
int td_req_qry_r_c_a_m_s_comb_product_info(tdApi a, CThostFtdcQryRCAMSCombProductInfoField * pQryRCAMSCombProductInfo, int nRequestID);
int td_req_qry_r_c_a_m_s_instr_parameter(tdApi a, CThostFtdcQryRCAMSInstrParameterField * pQryRCAMSInstrParameter, int nRequestID);
int td_req_qry_r_c_a_m_s_intra_parameter(tdApi a, CThostFtdcQryRCAMSIntraParameterField * pQryRCAMSIntraParameter, int nRequestID);
int td_req_qry_r_c_a_m_s_inter_parameter(tdApi a, CThostFtdcQryRCAMSInterParameterField * pQryRCAMSInterParameter, int nRequestID);
int td_req_qry_r_c_a_m_s_short_opt_adjust_param(tdApi a, CThostFtdcQryRCAMSShortOptAdjustParamField * pQryRCAMSShortOptAdjustParam, int nRequestID);
int td_req_qry_r_c_a_m_s_investor_comb_position(tdApi a, CThostFtdcQryRCAMSInvestorCombPositionField * pQryRCAMSInvestorCombPosition, int nRequestID);
int td_req_qry_investor_prod_r_c_a_m_s_margin(tdApi a, CThostFtdcQryInvestorProdRCAMSMarginField * pQryInvestorProdRCAMSMargin, int nRequestID);
int td_req_qry_r_u_l_e_instr_parameter(tdApi a, CThostFtdcQryRULEInstrParameterField * pQryRULEInstrParameter, int nRequestID);
int td_req_qry_r_u_l_e_intra_parameter(tdApi a, CThostFtdcQryRULEIntraParameterField * pQryRULEIntraParameter, int nRequestID);
int td_req_qry_r_u_l_e_inter_parameter(tdApi a, CThostFtdcQryRULEInterParameterField * pQryRULEInterParameter, int nRequestID);
int td_req_qry_investor_prod_r_u_l_e_margin(tdApi a, CThostFtdcQryInvestorProdRULEMarginField * pQryInvestorProdRULEMargin, int nRequestID);
int td_req_qry_investor_portf_setting(tdApi a, CThostFtdcQryInvestorPortfSettingField * pQryInvestorPortfSetting, int nRequestID);
int td_req_qry_investor_info_comm_rec(tdApi a, CThostFtdcQryInvestorInfoCommRecField * pQryInvestorInfoCommRec, int nRequestID);
int td_req_qry_comb_leg(tdApi a, CThostFtdcQryCombLegField * pQryCombLeg, int nRequestID);

#ifdef  __cplusplus
};
#endif
#endif  // GEN_TD_API_H