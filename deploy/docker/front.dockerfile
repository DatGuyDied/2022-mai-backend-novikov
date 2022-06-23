ARG NODE_VERSION=17-bullseye-slim
ARG NGINX_VERSION=1.21.6

FROM docker.io/nginx:${NGINX_VERSION}

RUN rm /etc/nginx/conf.d/*

COPY /front/build /opt/app/
COPY /deploy/nginx /etc/nginx/templates