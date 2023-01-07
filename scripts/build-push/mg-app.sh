#!/bin/bash
docker build -t mg-app ./app/mg-app
docker tag mg-app:latest public.ecr.aws/f9w6d5a2/mg-app:latest
docker push public.ecr.aws/f9w6d5a2/mg-app:latest
