# Add ssh module for running commands over ssh

**yorinasub17** commented *Mar 10, 2021*

This ports over the `ssh` module from `terratest` so that we can use it in our CLI tooling to SSH to instances. An example use case is SSH-ing to a deployed jenkins instance in the Reference Architecture to retrieve the admin token for configuring Jenkins over the API.

Note that this implementation has several changes over the `terratest` version:

- The `terratest` functions only supported a single hop for jump hosts. This implementation supports arbitrary hops by treating the `JumpHost` parameter as a linked list.
- Connection state tracking is simplified to a single `Stack` data structure that tracks all the items that need to be closed.
- `HostKeyCallback` is now configurable at the top level `Host` data structure.
<br />
***


**yorinasub17** commented *Mar 10, 2021*

Example use case: https://github.com/gruntwork-io/terraform-aws-architecture-catalog/pull/155
***

**yorinasub17** commented *Mar 11, 2021*

UPDATE:

- Better error handling so that the `sshAgent` can be stopped if there is an error loading keys to it: 72c9d76
- Add support for adding non-RSA private PEM keys to ssh agent: 77f1a09
***

**yorinasub17** commented *Mar 12, 2021*

Thanks for review! Merging this in now!
***

