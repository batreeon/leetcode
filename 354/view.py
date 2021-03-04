import matplotlib.pyplot as plt
import numpy as np

points = np.asarray([[30,15],[14,43],[17,18],[35,34],[2,7],[1,12],[11,14],[23,5],[45,13],[18,8],[24,31],[44,31],[47,45],[2,11],[6,44],[3,18],[36,45],[9,28],[30,13],[3,32],[8,20],[27,31],[25,23],[28,5],[5,26],[30,33],[18,46],[15,2],[6,24],[9,4],[5,12],[7,2],[2,27],[23,22]])
x = points[:,0]
print(x)
y = points[:,1]
print(y)
plt.scatter(x,y)

ax = plt.gca()
ax.set_xlim(0, 55)
miloc = plt.MultipleLocator(1)
ax.xaxis.set_minor_locator(miloc)
ax.grid(axis='x', which='minor')

ay = plt.gca()
ay.set_xlim(0, 55)
miloc = plt.MultipleLocator(1)
ay.yaxis.set_minor_locator(miloc)
ay.grid(axis='y', which='minor')

plt.grid()
plt.show()