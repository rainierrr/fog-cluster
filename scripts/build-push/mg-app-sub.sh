#!/bin/bash
docker build -t mg-app-sub ./app/mg-app-sub
docker tag mg-app-sub:latest public.ecr.aws/f9w6d5a2/mg-app-sub:latest
docker push public.ecr.aws/f9w6d5a2/mg-app-sub:latest
