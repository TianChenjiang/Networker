import numpy as np
from sklearn.preprocessing import MinMaxScaler

a = np.arange(11*5*5).reshape(11, 5, 5)
scalar = MinMaxScaler()
a = scalar.fit_transform(a)
print(a)
