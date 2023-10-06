FROM nginx:latest

COPY ./docker/dist /usr/share/nginx/html
COPY ./docker/adb_server /app/
COPY ./docker/enterpoint.sh /app/
COPY ./docker/default.conf /etc/nginx/conf.d/

RUN chmod +x /app/*

WORKDIR /app

EXPOSE 80

CMD ["./enterpoint.sh"]