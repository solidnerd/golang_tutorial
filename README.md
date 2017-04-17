# Some go sample files for a small go tutorial

All samples can be run on a Raspberry Pi and all run on any system with one exception: trafficLight uses GPIOs which are available on a Raspberry Pi only.

1. blinkAllLEDs - Sample which turns on and off LEDs connected to RaspberryPi GPIOs
2. functions - Sample for go functions
3. geoLocationSunriseSunset - Sample using go http client to retrieve the lattitude and longitude of a passed location and calculation of the local
4. gofunc - go functions used to demonstrate the dining philosophers problem using goroutines, channels and mutex
5. helloWorld - Simple Hello world program
6. highLow - Simple high/low game on command window level
7. highLowGameServer - Simple high/low game server using templates and go http server. IBM Bluemix artefacts to deploy the server on bluemix are included
8. httpsServer - Sample go https server which uses TLS and generates a self signed certificate
9. interfaces - Sample code which demonstrates go interfaces  
10. ipLocation - Retrieve location information about the current used IP address via REST call sunrise and sunset time via REST calls
11. loginFritz - Login into an AVM Fritzbox and handle the challenge response. The code retrieves the internet usage counters
12. methods - Sample go code which demonstrates usage om methods
13. trafficLight - Simple traffix light simulation using channels and go routines. Can be used on Raspberry when LEDs are connected to GPIOs
14. types - Samples for the different types in go
15. utf8 - Sample for utf8 handling in go
