FROM python:3.10-slim

WORKDIR /src
COPY ./requirements.txt /src/requirements.txt
COPY ./src /src
RUN apt-get update
RUN pip install -r requirements.txt

ENTRYPOINT ["locust"]

