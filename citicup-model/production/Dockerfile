FROM tensorflow/tensorflow:1.14.0-py3

WORKDIR /usr/app
COPY ./requirements.txt ./
RUN pip install -i https://pypi.tuna.tsinghua.edu.cn/simple -r requirements.txt

RUN  sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN  apt-get clean
ENV TZ=Asia/Shanghai

RUN echo "${TZ}" > /etc/timezone \
    && ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime \
    && apt update \
    && apt install -y tzdata \
    && rm -rf /var/lib/apt/lists/*

COPY ./ ./

CMD ["gunicorn", "-c", "gunicorn.py", "wsgi:app"]