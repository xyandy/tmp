
pwd_dir=`pwd`
dir=`basename ${pwd_dir}`

aws_ecr="312013929070.dkr.ecr.ap-northeast-1.amazonaws.com/panda-cronjob"
commit_hash=`git rev-parse --short HEAD`
tag=${dir}-${commit_hash}
docker_image=$aws_ecr:$tag

echo "your docker image is ${docker_image}"
echo "note that this image will be deployed to k8s cluster after you make this pr"
