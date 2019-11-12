FROM tensorflow/tensorflow:1.14.0-py3
WORKDIR /usr/app
COPY ./requirements.txt ./
RUN pip install -i https://pypi.tuna.tsinghua.edu.cn/simple -r requirements.txt
COPY ./ ./

CMD ["gunicorn", "-c", "gunicorn.py", "wsgi:app"]