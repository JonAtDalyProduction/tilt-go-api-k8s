# -*- mode: Python -*-

# For more on Extensions, see: https://docs.tilt.dev/extensions.html
load('ext://restart_process', 'docker_build_with_restart')

compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/tilt-example-api ./cmd/api'

local_resource(
  'example-go-compile',
  compile_cmd,
  deps=['./cmd/api/main.go'])

docker_build_with_restart(
  'example-go-api-image',
  '.',
  entrypoint=['/app/build/tilt-example-api'],
  dockerfile='Dockerfile',
  only=[
    './build',
  ],
  live_update=[
    sync('./build', '/app/build'),
  ],
)

k8s_yaml('deployments/api-deploy.yaml')
k8s_resource('example-go-api', port_forwards=3333,
             resource_deps=['example-go-compile'])