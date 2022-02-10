FROM nginx:stable-alpine
# FROM nginx:stable

LABEL maintainer="lalifeier <lalifeier@gmail.com>"

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# RUN apt-get update; \
#     apt-get -y --no-install-recommends install \
#         gosu

# COPY entrypoint.sh /usr/local/bin/entrypoint.sh
# RUN chmod +x /usr/local/bin/entrypoint.sh
# ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
