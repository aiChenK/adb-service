FROM nginx:latest

COPY ./docker/default.conf /etc/nginx/conf.d/
COPY ./docker/enterpoint.sh /app/
COPY ./docker/platform-tools/ /app/platform-tools/
COPY ./docker/adb_server /app/
COPY ./docker/dist /usr/share/nginx/html

RUN chmod +x /app/*
ENV PLATFORM_TOOLS_PATH=/app/platform-tools
ENV PATH=$PLATFORM_TOOLS_PATH:$PATH

WORKDIR /app

EXPOSE 80

CMD ["./enterpoint.sh"]