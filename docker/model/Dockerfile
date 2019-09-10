FROM tensorflow/tensorflow:latest-py3
WORKDIR /usr/app
COPY citicup-model/production/requirements.txt ./
RUN pip install -i https://pypi.tuna.tsinghua.edu.cn/simple -r requirements.txt
COPY citicup-model/production ./
ARG EXPOSED_PORT=5000

EXPOSE ${EXPOSED_PORT}
CMD ["gunicorn", "-c", "gunicorn.py", "wsgi:app"]
