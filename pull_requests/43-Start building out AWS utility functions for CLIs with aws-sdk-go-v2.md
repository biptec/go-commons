# Start building out AWS utility functions for CLIs, with aws-sdk-go-v2

**yorinasub17** commented *Apr 16, 2021*

I am starting to build out a new CLI for a hackday project, and in the process I realized that we don't have a central library for AWS access in production CLIs (resulting in duplication: e.g., see [refarch-deployer](https://github.com/gruntwork-io/terraform-aws-architecture-catalog/tree/main/tools/refarch-deployer/aws) and [kubergrunt](https://github.com/gruntwork-io/kubergrunt/tree/master/eksawshelper)). This is the start of building out a common aws helper library for use in CLIs.

In addition, [aws-sdk-go v2 is now out](https://aws.amazon.com/blogs/developer/aws-sdk-for-go-version-2-v2-release-candidate/). We should start to slowly transition to using the v2 line, as I expect v1 to be deprecated soon.

In that regard, I propose we:

- Start to build out common AWS functionality here.
- For new CLI tools that interact with AWS, use go-commons' aws interface as much as possible.
- Implicitly, new CLI tools should use `aws-sdk-go-v2`.
- In ticket jam, slowly upgrade existing tools to switch to `aws-sdk-go-v2`. For common libraries like `terratest`, this should be done by adding a `v2` folder (and drop v1 when we reach parity). Note that this should be done in parallel with the switch to `go-commons`.

This PR is the start of that.
<br />
***


**yorinasub17** commented *Apr 28, 2021*

Thanks for review! Merging this in!
***

