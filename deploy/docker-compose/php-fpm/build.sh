#!/usr/bin/env bash

# docker cp php-fpm:/usr/local/etc/php-fpm.d/www.conf www.conf
docker cp php-fpm:/etc/php/8.1/fpm/php-fpm.conf www.conf
docker cp php-fpm:/usr/lib/php/8.1/php.ini-production php.ini
