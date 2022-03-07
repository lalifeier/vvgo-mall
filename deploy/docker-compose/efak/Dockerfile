FROM java:8-alpine

LABEL maintainer="lalifeier <lalifeier@gmail.com>"

ENV TIME_ZONE=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

ENV EAGLE_VERSION=2.1.0
ENV KE_HOME=/usr/local/efak
ENV PATH=$PATH:$KE_HOME/bin

RUN apk update && \
    apk add curl sqlite && \
    mkdir -p ${KE_HOME} /hadoop/kafka-eagle/db && \
    curl -OL "https://github.com/smartloli/kafka-eagle-bin/archive/v${EAGLE_VERSION}.tar.gz"  && \
    tar -xzvf v${EAGLE_VERSION}.tar.gz && \
    tar -xzvf kafka-eagle-bin-${EAGLE_VERSION}/efak-web-${EAGLE_VERSION}-bin.tar.gz && \
    mv efak-web-${EAGLE_VERSION}/* ${KE_HOME} && \
    chmod +x ${KE_HOME}/bin/ke.sh && \
    rm -rf v${EAGLE_VERSION}.tar.gz kafka-eagle-bin-${EAGLE_VERSION}

WORKDIR ${KE_HOME}
COPY ./conf/system-config.properties ./conf/system-config.properties
COPY entrypoint.sh ./bin/entrypoint.sh

CMD ["sh", "./bin/entrypoint.sh"]
