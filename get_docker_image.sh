
if [ $# -ne 1 ]; then
    echo "Usage: $0 {dir}"
    exit 1
fi

echo "docker image is 312013929070.dkr.ecr.ap-northeast-1.amazonaws.com/panda-cronjob:{dir}-{commit_hash}"
echo "please input your dir"

dir=`echo $1 | sed 's#/##g'`

aws_ecr="312013929070.dkr.ecr.ap-northeast-1.amazonaws.com/panda-cronjob"
commit_hash=`git rev-parse --short HEAD`
tag=${dir}-${commit_hash}
image=$aws_ecr:$tag

echo "------------------------------------------------"
echo "docker image: ${image}"