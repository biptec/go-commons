# Add concept of sensitive args

**yorinasub17** commented *Nov 28, 2018*

Right now we log the command string when we call commands using `RunShellCommand`, which is really nice for debugging purposes. However, sometimes this should be suppressed because the args might include sensitive data. For example, when scripting calls to `openssl` to create certs, the password can only be passed in via the command line when avoiding the prompt.

The solution proposed and implemented here is to add a new shell option that marks the commands as having sensitive args, which is used to suppress logging the args.
<br />
***


**yorinasub17** commented *Nov 28, 2018*

> It would be even better if we could mark specific arguments as sensitive, but that's prob too error prone.

Yea I thought about this, but it was hard to come up with a good API for the general case when the target command could be anything.
***

**yorinasub17** commented *Nov 28, 2018*

Merging and releasing!
***

