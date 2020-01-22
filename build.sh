docker run --rm -v $(pwd):/src -v /var/run/docker.sock:/var/run/docker.sock centurylink/golang-builder eedgar/list_instances:latest
docker tag eedgar/list_instances:latest smartthings-docker-deploy.jfrog.io/docker-deploy/list_instances:latest

TAG=$(git log -1 --pretty=%h)
docker tag eedgar/list_instances:latest smartthings-docker-deploy.jfrog.io/docker-deploy/list_instances:${TAG}
