# The Anscombe Quartet (Python)

# demonstration data from
# Anscombe, F. J. 1973, February. Graphs in statistical analysis. 
#  The American Statistician 27: 17–21.

# prepare for Python version 3x features and functions
from __future__ import division, print_function

# import packages for Anscombe Quartet demonstration
import pandas as pd  # data frame operations
import numpy as np  # arrays and math functions
import statsmodels.api as sm  # statistical models (including regression)
import time

# define the anscombe data frame using dictionary of equal-length lists
anscombe = pd.DataFrame({'x1' : [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
    'x2' : [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
    'x3' : [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
    'x4' : [8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8],
    'y1' : [8.04, 6.95,  7.58, 8.81, 8.33, 9.96, 7.24, 4.26,10.84, 4.82, 5.68],
    'y2' : [9.14, 8.14,  8.74, 8.77, 9.26, 8.1, 6.13, 3.1,  9.13, 7.26, 4.74],
    'y3' : [7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73],
    'y4' : [6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89]})

def time_regression(x, y, n_runs):
    total_time = 0.0

    for _ in range(n_runs):
        start = time.perf_counter()

        X = sm.add_constant(x)
        model = sm.OLS(y, X)
        model.fit
        
        end = time.perf_counter()
        total_time+=(end - start)

    return total_time/n_runs

# fit linear regression models by ordinary least squares
n=1000

avg_I = time_regression(anscombe['x1'], anscombe['y1'], n)
avg_II = time_regression(anscombe['x2'], anscombe['y2'], n)
avg_III = time_regression(anscombe['x3'], anscombe['y3'], n)
avg_IV = time_regression(anscombe['x4'], anscombe['y4'], n)

print(f"Set I avg runtime:  {avg_I:.6f} sec")
print(f"Set II avg runtime: {avg_II:.6f} sec")
print(f"Set III avg runtime:{avg_III:.6f} sec")
print(f"Set IV avg runtime: {avg_IV:.6f} sec")
            
# Suggestions for the student:
# See if you can develop a quartet of your own, 
# or perhaps just a duet, two very different data sets 
# with the same fitted model.