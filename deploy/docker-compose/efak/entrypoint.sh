#!/bin/bash

cd ${KE_HOME}
# sed -i s/cluster1.zk.list=tdn1:2181,tdn2:2181,tdn3:2181/cluster1.zk.list=${ZK_HOSTS}/ conf/system-config.properties

sh ./bin/ke.sh start
tail -f ./logs/log.log
