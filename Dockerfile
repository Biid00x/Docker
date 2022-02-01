# --------------- DÉBUT COUCHE OS -------------------
FROM golang:1.17
# --------------- FIN COUCHE OS ---------------------

ENV GO111MODULE=auto

# RÉPERTOIRE DE TRAVAIL

WORKDIR /go/src/app

# ON COPIE le fichier main.go
COPY . .


RUN go build
# ON EXECUTE LE FICHIER MAIN.GO
CMD go run main.go


# OUVERTURE DU PORT HTTP

EXPOSE 8080
