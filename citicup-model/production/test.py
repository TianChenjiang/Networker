import numpy as np

arr = np.arange(27).reshape((3, 3, 3))
arr = arr.reshape((-1, 3))
print(arr)