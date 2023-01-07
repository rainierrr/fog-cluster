#!/bin/bash
docker build -t fog-app ./app/fog-app
docker tag fog-app:latest public.ecr.aws/f9w6d5a2/fog-app:latest
docker push public.ecr.aws/f9w6d5a2/fog-app:latest
