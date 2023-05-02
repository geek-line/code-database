FROM node:16.15.0

WORKDIR /frontend

RUN apt update -y && \
    apt upgrade -y && \
    # Necessary packages for npm to work
    apt install -y build-essential