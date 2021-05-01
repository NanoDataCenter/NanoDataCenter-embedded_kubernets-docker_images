Stream results store in sql like db with fields

path -- string encoded path
key  -- string
fields -- string encoded list
lag     -- int
number_of_deviations float
influence float

measurement results stored in db
moving average
moving std
moving average_lag_samples  -- msgpack data
moving std_lag_samples -- msgpack data
min  average
max average
min std
max std
current_samples   int -- set to zero to reset stream

