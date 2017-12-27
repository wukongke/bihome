FROM golang:1.9.2
RUN mkdir /app
WORKDIR /app

ENV PORT=1324
ENV LOG_LEVEL=info

COPY . /app

WORKDIR app

RUN go build -o app

CMD ["app/app -mode='dev'"]
EXPOSE 1324