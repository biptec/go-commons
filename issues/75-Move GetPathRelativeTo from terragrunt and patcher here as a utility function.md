# Move "GetPathRelativeTo" from terragrunt and patcher here as a utility function

**ina-stoyanova** commented *Oct 13, 2022*

So far, we have had to copy and paste the `GetPathRelativeTo` function from `terragrunt` repo to the `patcher/util` repo package. We should instead have this command available here so we avoid duplication.
<br />
***


**infraredgirl** commented *Oct 19, 2022*

This function already exists in `go-commons`: https://github.com/gruntwork-io/go-commons/blob/master/files/paths.go#L70.

However, it seems that a [bug fix](https://github.com/gruntwork-io/terragrunt/pull/100) has been implemented in the function in Terragrunt, and hasn't been ported here.

Therefore the right course of action is:

- [x] Port the above bug fix from Terragrunt to `go-commons`
- [ ] Remove the definition of `GetPathRelativeTo` from Terragrunt (and any other repos that currently define their own versions of it) and use the function from `go-commons` instead

***

**infraredgirl** commented *Oct 25, 2022*

Another bug in `GetPathRelativeTo()` was discovered and I just posted a fix in #78. After this is merged, I'll release another version of `go-commons` and ensure that all the repos that start to make use of `GetPathRelativeTo()` use this latest version.
***

**infraredgirl** commented *Oct 25, 2022*

> Another bug in `GetPathRelativeTo()` was discovered and I just posted a fix in #78. After this is merged, I'll release another version of `go-commons` and ensure that all the repos that start to make use of `GetPathRelativeTo()` use this latest version.

This has been done.
***

