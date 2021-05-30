### Ralali Go Error Wapped
We have wrapped the go error interface to make debugging easy and have standary error response through out all golang micro services.
This also support __on-demand__ stack trace of error.  To get the stack_trace of the error you need to pass __"debug"__ key in context of the error.
You can also change the stack trace depth by passing the __trace_depth__ in context with you desired depth value.   
In __rl-go-microservice__ all the API endpoint will pass there two keys in context if that exists in __API Header__.

### Sample Header
![Screen Shot 2020-03-05 at 15 55 59](https://user-images.githubusercontent.com/59592686/75964456-d44e7280-5ef9-11ea-914d-16bbbac31e43.png)

### Sample Response
![Screen Shot 2020-03-05 at 15 50 17](https://user-images.githubusercontent.com/59592686/75964036-1cb96080-5ef9-11ea-8eda-366ad95d8550.png)
