FROM node:14.15.2-alpine3.10

# set working directory
WORKDIR /app

ADD ./entrypoint.sh /entrypoint.sh

# for window
RUN cat /entrypoint.sh | tr -d '\r' > /entrypoint.sh

EXPOSE 3000