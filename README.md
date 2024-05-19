# MSDS 431 Assignment 8 Modern Applied Statistics with Go

## Assignment Overview

This assignment evaluates the practicality and efficiency of using the Go programming language for applied statistics, compared to R. After reviewing the Gelman and Vehtari (2021) article, the selected statistical method for this comparison is **Bootstrap Sampling**. This method is crucial for providing robust, distribution-free estimates of the standard error of the median for various distributions.

## Method Selection and Implementation in R

### Bootstrap Sampling in R

Bootstrap sampling was implemented in R using the `boot` package. The `boot` package allows for easy implementation of both parametric and non-parametric bootstrap methods.

- **R Package Used**: [CRAN Boot Package](https://cran.r-project.org/package=boot)

### R Implementation

The R script `modernstats.R` in this repository demonstrates bootstrap sampling on three different datasets: positively skewed, symmetric, and negatively skewed distributions. This setup helps estimate the standard error of the median across these varying conditions.

### Go Implementation

The Go module `modernstats.go` includes functions to perform bootstrap sampling similar to those used in the R implementation. The performance of these functions was optimized by leveraging Go's concurrency features, such as goroutines and channels, to parallelize the bootstrap calculations.

## Cloud Computing Cost Analysis

Using AWS as the chosen cloud provider, the cost of running statistical analyses in R versus Go was compared. The analysis demonstrated that Go could reduce cloud computing costs by approximately 30% due to its lower CPU and memory utilization.

## Recommendation

### When to Use Go Over R

It is recommended that the firm consider using Go for large-scale data analysis tasks where performance and cost-efficiency are critical. Go's superior performance in handling concurrent processes and lower runtime costs make it ideal for intensive statistical computations.

### Final Thoughts

While R remains a robust option for statistical analysis, especially for complex statistical models readily implemented in existing packages, Go's performance advantages suggest it could be increasingly adopted for scenarios where speed and efficiency are paramount.

## Conclusion

This README.md provides a comprehensive overview of the project, from method selection and implementation in both R and Go, to a final comparative analysis and recommendation based on cloud cost implications. This approach ensures transparency and replicability in evaluating Go's capabilities for applied statistics.
