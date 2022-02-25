FROM golang:1.15.2-alpine

ARG firestore_emulator_host
ARG gcloud_project_id

RUN apk update && apk add git
RUN mkdir /go/src/app
WORKDIR /go/src/app

COPY go.* ./
RUN go mod download

ADD . /go/src/app

ENV FIRESTORE_EMULATOR_HOST $firestore_emulator_host
ENV GCLOUD_PROJECT_ID $gcloud_project_id

EXPOSE 8000
CMD ["go", "run", "main.go"]

#RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
#    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air
#CMD air
