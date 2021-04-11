'''

lag: the lag parameter determines how much your data will be smoothed and how adaptive the algorithm is to changes in the long-term average of the data. The more stationary your data is, the more lags you should include (this should improve the robustness of the algorithm). If your data contains time-varying trends, you should consider how quickly you want the algorithm to adapt to these trends. I.e., if you put lag at 10, it takes 10 'periods' before the algorithm's treshold is adjusted to any systematic changes in the long-term average. So choose the lag parameter based on the trending behavior of your data and how adaptive you want the algorithm to be.

influence: this parameter determines the influence of signals on the algorithm's detection threshold. If put at 0, signals have no influence on the threshold, such that future signals are detected based on a threshold that is calculated with a mean and standard deviation that is not influenced by past signals. If put at 0.5, signals have half the influence of normal data points. Another way to think about this is that if you put the influence at 0, you implicitly assume stationarity (i.e. no matter how many signals there are, you always expect the time series to return to the same average over the long term). If this is not the case, you should put the influence parameter somewhere between 0 and 1, depending on the extent to which signals can systematically influence the time-varying trend of the data. E.g., if signals lead to a structural break of the long-term average of the time series, the influence parameter should be put high (close to 1) so the threshold can react to structural breaks quickly.

threshold: the threshold parameter is the number of standard deviations from the moving mean above which the algorithm will classify a new datapoint as being a signal. For example, if a new datapoint is 4.0 standard deviations above the moving mean and the threshold parameter is set as 3.5, the algorithm will identify the datapoint as a signal. This parameter should be set based on how many signals you expect. For example, if your data is normally distributed, a threshold (or: z-score) of 3.5 corresponds to a signaling probability of 0.00047 (from this table), which implies that you expect a signal once every 2128 datapoints (1/0.00047). The threshold therefore directly influences how sensitive the algorithm is and thereby also determines how often the algorithm signals. Examine your own data and choose a sensible threshold that makes the algorithm signal when you want it to (some trial-and-error might be needed here to get to a good threshold for your purpose).

'''            


# Implementation of algorithm from https://stackoverflow.com/a/22640362/6029703
import numpy as np


def thresholding_algo(y, lag, threshold, influence):
    signals = np.zeros(len(y))
    filteredY = np.array(y)
    avgFilter = [0]*len(y)
    stdFilter = [0]*len(y)
    avgFilter[lag - 1] = np.mean(y[0:lag])
    stdFilter[lag - 1] = np.std(y[0:lag])
    for i in range(lag, len(y)):
        if abs(y[i] - avgFilter[i-1]) > threshold * stdFilter [i-1]:
            if y[i] > avgFilter[i-1]:
                signals[i] = 1
            else:
                signals[i] = -1

            filteredY[i] = influence * y[i] + (1 - influence) * filteredY[i-1]
            avgFilter[i] = np.mean(filteredY[(i-lag+1):i+1])
            stdFilter[i] = np.std(filteredY[(i-lag+1):i+1])
        else:
            signals[i] = 0
            filteredY[i] = y[i]
            avgFilter[i] = np.mean(filteredY[(i-lag+1):i+1])
            stdFilter[i] = np.std(filteredY[(i-lag+1):i+1])

    return dict(signals = np.asarray(signals),
                avgFilter = np.asarray(avgFilter),
                stdFilter = np.asarray(stdFilter))
# Data
y = np.array([1,1,1.1,1,0.9,1,1,1.1,1,0.9,1,1.1,1,1,0.9,1,1,1.1,1,1,1,1,1.1,0.9,1,1.1,1,1,0.9,
       1,1.1,1,1,1.1,1,0.8,0.9,1,1.2,0.9,1,1,1.1,1.2,1,1.5,1,3,2,5,3,2,1,1,1,0.9,1,1,3,
       2.6,4,3,3.2,2,1,1,0.8,4,4,2,2.5,1,1,1])

# Settings: lag = 30, threshold = 5, influence = 0
lag = 30
threshold =  2 #5
influence = 0

# Run algo with settings from above
result = thresholding_algo(y, lag=lag, threshold=threshold, influence=influence)
for i in range(0, len(y)):
  print("i",i,y[i],result["signals"][i],result["avgFilter"][i],result["stdFilter"][i])
