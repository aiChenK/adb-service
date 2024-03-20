FROM debian:latest

COPY ./docker/platform-tools/ /app/platform-tools/
COPY ./docker/adb-service /app/

RUN chmod +x /app/*
ENV PLATFORM_TOOLS_PATH=/app/platform-tools
ENV PATH=$PLATFORM_TOOLS_PATH:$PATH

WORKDIR /app

EXPOSE 8088

CMD ["./adb-service"]