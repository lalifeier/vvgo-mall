FROM node:lts-alpine

LABEL maintainer="lalifeier <lalifeier@gmail.com>"

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN npm install -g pnpm

RUN pnpm config set registry https://registry.npmmirror.com
