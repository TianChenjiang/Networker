# 预测模型使用手册

### 1、生产环境下的预测模型放在了production文件夹下

### 2、构建镜像

`docker build -t greenw00d/citicup-model .`

### 3、绑定在本地6324端口运行

`docker run -p 6324:5000 greenw00d/citicup-model`

### 4、启用Swagger文档

在浏览器中访问 http://localhost:6324/ 即可

### 5、访问接口

curl -X POST "http://localhost:6324/api/PredictRisk/predict?code=000001.SZ&forecast_close_line=13" -H "accept: application/json"

#### 当然也可在postman里设置body传参，flask都支持





## 如果出现跨域之类的问题，请联系greenwood