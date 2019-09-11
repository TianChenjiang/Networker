FROM tensorflow/tensorflow:latest-py3
WORKDIR /usr/app
COPY ./requirements.txt ./
RUN pip install -r requirements.txt
COPY ./ ./

CMD ["gunicorn", "-c", "gunicorn.py", "wsgi:app"]