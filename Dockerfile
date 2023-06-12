FROM golang:1.18
WORKDIR /app
COPY . .
RUN go mod download
RUN go get github.com/jafari-mohammad-reza/hotel-reservation.git/api/handlers
RUN go get github.com/jafari-mohammad-reza/hotel-reservation.git/api/routes
RUN go get github.com/jafari-mohammad-reza/hotel-reservation.git/db
RUN make seed
EXPOSE 5000
CMD ["make" , "run"]
