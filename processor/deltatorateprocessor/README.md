# Delta to Rate Processor
<!-- status autogenerated section -->
| Status        |           |
| ------------- |-----------|
| Stability     | [development]: metrics   |
| Distributions | [contrib], [aws], [observiq], [sumo] |
| Issues        | ![Open issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aopen%20label%3Aprocessor%2Fdeltatorate%20&label=open&color=orange&logo=opentelemetry) ![Closed issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aclosed%20label%3Aprocessor%2Fdeltatorate%20&label=closed&color=blue&logo=opentelemetry) |

[development]: https://github.com/open-telemetry/opentelemetry-collector#development
[contrib]: https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib
[aws]: https://github.com/aws-observability/aws-otel-collector
[observiq]: https://github.com/observIQ/observiq-otel-collector
[sumo]: https://github.com/SumoLogic/sumologic-otel-collector
<!-- end autogenerated section -->

**Status: under development; Not recommended for production usage.**

## Description

The delta to rate processor (`deltatorateprocessor`) converts delta sum metrics to rate metrics. This rate is a gauge. 

## Configuration

Configuration is specified through a list of metrics. The processor uses metric names to identify a set of delta sum metrics and calculates the rates which are gauges.

```yaml
processors:
    # processor name: deltatorate
    deltatorate:

        # list the delta sum metrics to calculate the rate. This is a required field.
        metrics:
            - <metric_1_name>
            - <metric_2_name>
            .
            .
            - <metric_n_name>
```

[development]: https://github.com/open-telemetry/opentelemetry-collector#development
[contrib]:https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib
