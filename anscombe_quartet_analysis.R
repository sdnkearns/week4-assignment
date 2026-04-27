# The Anscombe Quartet (R)

# demonstration data from
# Anscombe, F. J. 1973, February. Graphs in statistical analysis. 
#  The American Statistician 27: 17–21.

# define the anscombe data frame
anscombe <- data.frame(
    x1 = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
    x2 = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
    x3 = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
    x4 = c(8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8),
    y1 = c(8.04, 6.95,  7.58, 8.81, 8.33, 9.96, 7.24, 4.26,10.84, 4.82, 5.68),
    y2 = c(9.14, 8.14,  8.74, 8.77, 9.26, 8.1, 6.13, 3.1,  9.13, 7.26, 4.74),
    y3 = c(7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73),
    y4 = c(6.58, 5.76,  7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89))

time_regression <- function(formula, data, n_runs=1000) {
    total_time <- 0

    for (i in 1:n_runs) {
        t <- system.time({
            model <- lm(formula, data = data)
        })
        total_time <- total_time + t["elapsed"]
    }

    return(total_time/n_runs)
}

n <- 1000

avg_I   <- time_regression(y1 ~ x1, anscombe, n)
avg_II  <- time_regression(y2 ~ x2, anscombe, n)
avg_III <- time_regression(y3 ~ x3, anscombe, n)
avg_IV  <- time_regression(y4 ~ x4, anscombe, n)

cat(sprintf("Set I avg runtime:   %.6f sec\n", avg_I))
cat(sprintf("Set II avg runtime:  %.6f sec\n", avg_II))
cat(sprintf("Set III avg runtime: %.6f sec\n", avg_III))
cat(sprintf("Set IV avg runtime:  %.6f sec\n", avg_IV))