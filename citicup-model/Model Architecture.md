# Model Architecture

![](D:\Study\Project\citi_cup\Networker\model\机器学习模型.jpg)

## Model Introduction

### 1、TCN (Temporal Convolutional Network)

![tcn](D:\Study\Project\citi_cup\Networker\model\tcn.jpg)

- Flexible receptive field size
- Convolution layers increase causal correlation
- Convolutions can be done in parallel, thus improve efficiency



### 2、LSTM (Long Short-Term Memory)

![lstm](D:\Study\Project\citi_cup\Networker\model\lstm.jpg)



- Avoiding the problem of gradient disappearing in RNN
- Extracting dependencies between time series data



### 3、DNN

![resnet](D:\Study\Project\citi_cup\Networker\model\resnet.jpg)

- using residual neural network
- mitigating the problem of gradient disappearing when the network is too deep



### 4、LightGBM

-  based on Gradient Boosting Decision Tree (GBDT)
-  LightGBM uses histogram-based algorithms. This speeds up training and reduces memory usage
- Leaf-wise tree growth algorithms with max depth limitation, which avoiding over-fit