FROM golang:1.19-alpine AS builder
RUN adduser -S goruntime
WORKDIR /work
COPY ./src /work/
RUN chmod -R 777 /work
USER goruntime 
RUN cd cmd/web/ && CGO_ENABLED=0 go build -o /work/binary

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
WORKDIR /work
USER goruntime
COPY --from=builder /work /work
EXPOSE 8000
ENTRYPOINT [ "/work/binary" ]