# Need an option for printing "simple" error messages.

**josh-padnick** commented *Dec 14, 2017*

While writing [health-checker](http://github.com/gruntwork-io/health-checker) (which is based on gruntwork-cli), I wanted to validate that all required arguments were passed in. Gruntwork-cli supports this by capturing errors and [logging them](https://github.com/gruntwork-io/gruntwork-cli/blob/master/entrypoint/entrypoint.go#L30), but the output seems verbose for a simple argument validation:

```
ERRO[2017-12-14T14:51:59-07:00] Missing required parameter --port             error="Missing required parameter --port"
```

I'd prefer at least the option to output something simpler configurable in my program (health-checker) that doesn't force me to step outside of the error handling provided by gruntwork-cli:

```
Missing required parameter --port
```
<br />
***


**brikis98** commented *Dec 14, 2017*

Agreed! The default error handling right now is optimized for debugging (as a developer) rather than fixing the issue (as an end user).
***

**yorinasub17** commented *Nov 28, 2018*

Closed with https://github.com/gruntwork-io/gruntwork-cli/pull/15
***

