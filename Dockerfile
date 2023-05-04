FROM golang:1.20

# Copy Files
RUN mkdir /home/bagshare
COPY . /home/bagshare

# install dependencies
WORKDIR /home/bagshare
RUN go mod download

# Go Build
RUN CGO_ENABLED=0 go build

# Open 5000 port
EXPOSE 5000

# Run the Gin server
CMD ["/home/bagshare/bag-share"]