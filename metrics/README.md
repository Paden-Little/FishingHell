Metrics will be reported by each individual service into a redis cache which is then read from and reported into a time series database

Each service will use a unique buffer within the redis cache, so each service will have its own buffer, and the metrics processing app knows how to process what from where

This service will also display a web based admin dashboard serving metric graphs to show usage 

# Flexible-mq
1. req/sec
	This is raw requests into the system through the Discord API
2. handle time
	This will record a time at request begin, note the service it will be pathed too and then record the time at return, and send those three datapoints to the Metrics cache
3. latency
	Amount of time it takes for the system to begin processing the request, will send time to begin handling request
4. response time
	What the user sees, I.E latency + handle time. Calculate response time within metrics app not from MQ
