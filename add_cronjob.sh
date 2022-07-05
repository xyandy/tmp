set -e

echo "please input your cronjob name"
if [ $# -ne 1 ]; then
    echo "Usage: $0 {cronjob_name}"
    exit 1
fi
cronjob_name=`echo $1 | sed 's#/##g'`
echo "your cronjob name is ${cronjob_name}"

dir=$cronjob_name
mkdir ${dir}

aws_ecr="312013929070.dkr.ecr.ap-northeast-1.amazonaws.com/panda-cronjob"
commit_hash=`git rev-parse --short HEAD`
tag=${dir}-${commit_hash}
docker_image=$aws_ecr:$tag
echo "your docker image is ${docker_image}"
echo "note that this image will be deployed to k8s cluster after you make this pr"


cp Dockerfile_sample ./${cronjob_name}/Dockerfile
cp sample.py ./${cronjob_name}/sample.py


ns="test"
schedule_expression='"*/1 * * * *"'
suspend=false

cat ./deploy_sample.yaml \
| sed "s#{ns}#${ns}#g" \
| sed "s#{cronjob_name}#${cronjob_name}#g" \
| sed "s#{docker_image}#${docker_image}#g" \
| sed "s#{schedule_expression}#${schedule_expression}#g" \
| sed "s#{suspend}#${suspend}#g" > ${dir}/deploy.yaml

echo "create dir ${dir} successfully, now you can update Dockerfile, codes and deploy.yaml according to your requirements"


