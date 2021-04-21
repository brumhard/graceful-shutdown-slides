# use with:
# GOARCH=amd64 GOOS=linux go build -o example example.go
# docker build -t example:latest .
# docker run example:latest

# START OMIT
FROM scratch

COPY ./example ./example

CMD ["./example"]
# END OMIT
