#!/bin/bash
docker build -t super-app ./app/super-app
docker tag super-app:latest public.ecr.aws/f9w6d5a2/super-app:latest
docker push public.ecr.aws/f9w6d5a2/super-app:latest
