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

WORKDIR $GOPATH/src/gitlab.com/hacheurs/hash-n-go

# Install go dependencies & app
RUN make master

ENTRYPOINT ["hash-n-go"]

# Should use FROM golang as intermediate and copy executable to docker in docker image
# But as time lacks, we simply removed swarm managment in the program.
# You can find the docker-swarm code in pkg/swarm/swarm.go