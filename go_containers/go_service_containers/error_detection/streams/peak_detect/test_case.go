
go stat package
https://github.com/montanaflynn/stats
https://github.com/montanaflynn/stats.git

/*
Settings (the ones below are examples: choose what is best for your data)
set lag to 5;          # lag 5 for the smoothing functions
set threshold to 3.5;  # 3.5 standard deviations for signal
set influence to 0.5;  # between 0 and 1, where 1 is normal influence, 0.5 is half
*/

// ZScore on 16bit WAV samples
func ZScore(samples []int16, lag int, threshold float64, influence float64) (signals []int16) {
    //lag := 20
    //threshold := 3.5
    //influence := 0.5

    signals = make([]int16, len(samples))
    filteredY := make([]int16, len(samples))
    for i, sample := range samples[0:lag] {
        filteredY[i] = sample
    }
    avgFilter := make([]int16, len(samples))
    stdFilter := make([]int16, len(samples))

    avgFilter[lag] = Average(samples[0:lag])
    stdFilter[lag] = Std(samples[0:lag])

    for i := lag + 1; i < len(samples); i++ {

        f := float64(samples[i])

        if float64(Abs(samples[i]-avgFilter[i-1])) > threshold*float64(stdFilter[i-1]) {
            if samples[i] > avgFilter[i-1] {
                signals[i] = 1
            } else {
                signals[i] = -1
            }
            filteredY[i] = int16(influence*f + (1-influence)*float64(filteredY[i-1]))
            avgFilter[i] = Average(filteredY[(i - lag):i])
            stdFilter[i] = Std(filteredY[(i - lag):i])
        } else {
            signals[i] = 0
            filteredY[i] = samples[i]
            avgFilter[i] = Average(filteredY[(i - lag):i])
            stdFilter[i] = Std(filteredY[(i - lag):i])
        }
    }

    return
}

// Average a chunk of values
func Average(chunk []int16) (avg int16) {
    var sum int64
    for _, sample := range chunk {
        if sample < 0 {
            sample *= -1
        }
        sum += int64(sample)
    }
    return int16(sum / int64(len(chunk)))
}