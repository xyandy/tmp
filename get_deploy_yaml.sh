ns="test"
cronjob_name="test-cronjob"
cronjob_cmd='["/bin/sh", "-c", "sleep 10; echo hello world"]'
docker_image="alpine:latest"
schedule_expression='"*/1 * * * *"'
suspend=false

cat ./deploy_template.yaml \
| sed "s#{ns}#${ns}#g" \
| sed "s#{cronjob_name}#${cronjob_name}#g" \
| sed "s#{cronjob_cmd}#${cronjob_cmd}#g" \
| sed "s#{docker_image}#${docker_image}#g" \
| sed "s#{schedule_expression}#${schedule_expression}#g" \
| sed "s#{suspend}#${suspend}#g"