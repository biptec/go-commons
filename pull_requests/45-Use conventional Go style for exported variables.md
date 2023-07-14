# Use conventional Go style for exported variables

**bwhaley** commented *Apr 23, 2021*


<br />
***


**yorinasub17** commented *Apr 23, 2021*

Looks like there is a build error:
```
# github.com/gruntwork-io/go-commons/version
version/version.go:16:6: Version redeclared in this block
	previous declaration at version/version.go:12:2
version/version.go:17:2: cannot use Version (type func() string) as type string in return argument
```
***

**bwhaley** commented *Apr 23, 2021*

Oh, duh. 🤦‍♂️

Any objection to renaming this function to `GetVersion()`?
***

**bwhaley** commented *Apr 23, 2021*

Thanks for the reviews and for being patient with my neurotic Go syntax obsession.
***

**bwhaley** commented *Apr 23, 2021*

https://github.com/gruntwork-io/go-commons/releases/tag/v0.9.0
***

