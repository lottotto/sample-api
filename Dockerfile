# FROM alpine:latest
# COPY main .

# EXPOSE 1323
# ENTRYPOINT ./main

# #This is the "builder" stage
FROM golang:1.14 as builder
WORKDIR /my-go-app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build main.go
#This is the final stage, and we copy artifacts from "builder"
FROM gcr.io/distroless/base-debian10:debug
COPY --from=builder /my-go-app/main /bin/main
EXPOSE 1323
ENTRYPOINT ["/bin/main"]