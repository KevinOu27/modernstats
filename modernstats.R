library(boot)

medianfunction = function(data, indices){
  d = data[indices]
  return(median(d))
}

data = rnorm(100)

data

results = boot(data, medianfunction, R = 1000)
print(results)

