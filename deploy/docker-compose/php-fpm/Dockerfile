FROM phpdockerio/php:8.1-fpm

LABEL maintainer="lalifeier <lalifeier@gmail.com>"

RUN apt-get update; \
    apt-get -y --no-install-recommends install \
        git \
        php8.1-mysql \
        php8.1-redis \
        php8.1-gd \
        php8.1-imagick \
        php8.1-xdebug; \
    apt-get clean; \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* /usr/share/doc/*

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
