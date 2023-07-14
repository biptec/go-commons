# Consolidate multierror packages

**infraredgirl** commented *Nov 11, 2020*

Currently we have several implementations of the multierror functionality spread out across different repos:

- [terragrunt](https://github.com/gruntwork-io/terragrunt/blob/master/errors/multierror.go)
- [terratest](https://github.com/gruntwork-io/terratest/blob/master/modules/customerrors/multierror.go)
- [cloud-nuke](https://github.com/gruntwork-io/cloud-nuke/blob/master/util/errors.go)

This is bad for many reasons not least of which being maintenance. Ideally we would avoid reinventing the wheel by using an existing open source implementation of the multierror concept, such as [hashicorp/go-multierror](https://github.com/hashicorp/go-multierror). However, this is currently not an option due to the licensing limitations. 

Therefore we have two options:

1. If we conclude to allow MPL licences (see https://github.com/gruntwork-io/company/issues/48 for more context), we should get rid of all the above custom implementations of multierror and replace them with `go-multierror`.
1. If the above is not possible, we should migrate the implementation of multierror to `gruntwork-cli`.


<br />
***


**infraredgirl** commented *Jan 6, 2021*

Update: Since we've been [cleared to use code licensed with MPL](https://gruntwork-io.slack.com/archives/CBUAJT85V/p1609923499010900), we should proceed with option 1 above.
***

**infraredgirl** commented *Jan 6, 2021*

I've opened relevant issues in corresponding repos: 
- https://github.com/gruntwork-io/terragrunt/issues/1482
- https://github.com/gruntwork-io/terratest/issues/751
- https://github.com/gruntwork-io/cloud-nuke/issues/174

Since no change will be done in this repo, I'm closing this issue.
***

