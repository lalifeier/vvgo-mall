FROM php:8.1-fpm

LABEL maintainer="lalifeier <lalifeier@gmail.com>"

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# RUN sed -i s/archive.ubuntu.com/mirrors.aliyun.com/g /etc/apt/sources.list \
#     && sed -i s/security.ubuntu.com/mirrors.aliyun.com/g /etc/apt/sources.list \
#     && apt-get update

RUN sed -i s/deb.debian.org/mirrors.aliyun.com/g /etc/apt/sources.list \
    && apt-get update

RUN apt-get update; \
    apt-get -y --no-install-recommends install \
        git; \
        # php8.1-mysql \
        # php8.1-redis \
        # php8.1-gd \
        # php8.1-imagick \
        # php8.1-xdebug \
    apt-get clean; \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* /usr/share/doc/*

RUN curl -sS https://getcomposer.org/installer | php && mv composer.phar /usr/local/bin/composer

RUN composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/
