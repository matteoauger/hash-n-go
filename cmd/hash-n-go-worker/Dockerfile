# Base : go image
FROM golang

# Git credentials
# --build-arg USERNAME={your username}
# --build-arg PASSWORD={your password}
ARG USERNAME 
ARG PASSWORD

# Websocket uri
ENV WS_URI ws://localhost:8080

WORKDIR $GOPATH/src

# Should have used SSH keys here but ULH blocks SSH connections
# Blame 'em ;)
# Clone repo
RUN git clone \
    https://$USERNAME:$PASSWORD@www-apps.univ-lehavre.fr/forge/hacheurs/hash-n-go \
    gitlab.com/hacheurs/hash-n-go

# TODO remove both lines on merge to master
WORKDIR $GOPATH/src/gitlab.com/hacheurs/hash-n-go
RUN git checkout 3-docker-image-master

# Install go dependencies & app
RUN make worker

CMD hash-n-go-worker $WS_URI
