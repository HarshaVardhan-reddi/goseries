# Context

This is package part of standard library package in golang

## Use cases
<!-- (Many), following are crucial -->
### 1.Specifying timeouts / time bounds for a go routine / a block of code to execute

Eg: Adding time out whenever doing external api calls from the application

### 2. Cancelling a go routine

Eg: If there are any go routines are running, those can be killed by using contexts

### 3. Use as a global variable that can be used accross the application

Eg: Access current user details accross the go application
