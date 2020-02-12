FROM golang as build-stage
COPY greeter.go /
RUN cd / && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o greeter

FROM alpine
COPY --from=build-stage /greeter /
CMD ["/greeter"]