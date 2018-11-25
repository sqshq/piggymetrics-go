# PiggyMetrics Golang edition
Light and simple version of the [PiggyMetrics](https://github.com/sqshq/PiggyMetrics) project, written in Go with embedded database under the hood. Used for minimum-cost deployment to the cloud. 


### Build and run

``` bash
docker build -t sqshq/piggymetrics .
docker run -p 8080:8080  --rm sqshq/piggymetrics
```