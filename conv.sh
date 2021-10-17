#!/bin/sh
mkdir tmp
iconv -f GBK -t UTF-8 include/ThostFtdcMdApi.h -o tmp/ThostFtdcMdApi.h
iconv -f GBK -t UTF-8 include/ThostFtdcTraderApi.h -o tmp/ThostFtdcTraderApi.h
iconv -f GBK -t UTF-8 include/ThostFtdcUserApiDataType.h -o  tmp/ThostFtdcUserApiDataType.h
iconv -f GBK -t UTF-8 include/ThostFtdcUserApiStruct.h -o tmp/ThostFtdcUserApiStruct.h

#sed -i 's/'"'"'102001'"'"'/"102001"/g' tmp/ThostFtdcUserApiDataType.h
#sed -i 's/'"'"'102002'"'"'/"102002"/g' tmp/ThostFtdcUserApiDataType.h
#sed -i 's/'"'"'202001'"'"'/"202001"/g' tmp/ThostFtdcUserApiDataType.h
#sed -i 's/'"'"'202002'"'"'/"202002"/g' tmp/ThostFtdcUserApiDataType.h
mv include include_bak
mv tmp include
