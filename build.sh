docker build --no-cache -t physicalgraph/list_instances:latest .
docker run --rm -v $(pwd):/src physicalgraph/list_instances:latest cp /list_instances /src

docker tag physicalgraph/list_instances:latest smartthings-docker-deploy.jfrog.io/docker-deploy/list_instances:latest

TAG=$(git log -1 --pretty=%h)
docker tag physicalgraph/list_instances:latest smartthings-docker-deploy.jfrog.io/docker-deploy/list_instances:${TAG}
