# Base : go image
FROM golang

# Git credentials
# --build-arg USERNAME={your username}
# --build-arg PASSWORD={your password}
ARG USERNAME 
ARG PASSWORD

WORKDIR $GOPATH/src

# Should have used SSH keys here but ULH blocks SSH connections
# Blame 'em ;)
# Clone repo
RUN git clone \
    https://$USERNAME:$PASSWORD@www-apps.univ-lehavre.fr/forge/hacheurs/hash-n-go \
    gitlab.com/hacheurs/hash-n-go

# TODO remove 
WORKDIR $GOPATH/src/gitlab.com/hacheurs/hash-n-go
RUN git checkout 3-docker-image-master

# Install go dependencies & app
RUN make worker

ENTRYPOINT ["hash-n-go-worker"]
