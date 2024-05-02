FROM python:3.11-alpine

WORKDIR /opt/notigram/

ADD app app
ADD requirements.txt requirements.txt

RUN pip install -r requirements.txt

CMD gunicorn -k uvicorn.workers.UvicornWorker -b 0.0.0.0:8000 app:app
