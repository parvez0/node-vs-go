# Node.js vs Go
Benchmarking rest api of Node.js and GO

I was working on a Node.js micro-service which handles a lot of synchronous api calls to external services and 
started facing a lot of performance issues at peak load, we were seeing around 20% - 30% of request drops
and there was no change in the resources allocated to the services. So we decided to make all the apis assynchronous.
After this change we started seeing a lot of improvement in our services, and I started wondering what will happen if I
replace the Node.js with Go as this is new trend going on these days, so I decided to do a little experiment with 2 simple 
micro-services one written in Node.js and other in Go and the results were shocking.

