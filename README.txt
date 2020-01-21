This project can be built with the golang-builder
https://hub.docker.com/r/centurylink/golang-builder (https://github.com/CenturyLinkLabs/golang-builder)

add this function to your bashrc after building it with build.sh
build.sh will create a docker container and a binary that you can use directly

list_instances() { 
    docker run -it -e AWS_PROFILE=$AWS_PROFILE -v $HOME/.aws:/root/.aws eedgar/list_instances:latest ./list_instances $*
}


