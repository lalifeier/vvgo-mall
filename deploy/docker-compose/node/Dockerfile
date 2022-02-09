FROM node:lts-alpine

LABEL maintainer="lalifeier <lalifeier@gmail.com>"

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN npm install -g pnpm

RUN pnpm config set registry https://registry.npmmirror.com

# COPY .npmrc package.json pnpm-lock.yaml .pnpmfile.cjs ./

# ARG NODE_ENV
# RUN if [ "$NODE_ENV" = "development" ]; \
#     then pnpm install; \
#     else pnpm install --frozen-lockfile --prod; \
#     fi

# COPY . .

# RUN pnpm build
