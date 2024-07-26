FROM golang:1.22.2
WORKDIR /app-golang
COPY . .
ARG PORT_BUILD=3000
ENV PORT=$PORT_BUILD
EXPOSE $PORT_BUILD
RUN go build -o loja
RUN rm -r controllers
RUN rm -r db
RUN rm *.go
RUN rm *.sum
RUN rm -r routes
RUN rm -r models
RUN rm *.mod
RUN rm -r Dockerfile
ENTRYPOINT ./loja