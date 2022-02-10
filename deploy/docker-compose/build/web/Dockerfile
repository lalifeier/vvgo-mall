# build stage
FROM node:lts-alpine AS build-stage

RUN npm install -g pnpm
RUN pnpm config set registry https://registry.npmmirror.com

WORKDIR /app

COPY .npmrc package.json pnpm-lock.yaml .pnpmfile.cjs ./
RUN pnpm install --frozen-lockfile --prod
COPY . .
RUN pnpm build

# production stage
FROM nginx:stable-alpine AS production-stage
LABEL maintainer="lalifeier <lalifeier@gmail.com>"

COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY docker/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
