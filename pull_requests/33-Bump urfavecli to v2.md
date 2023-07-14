# Bump urfave/cli to v2

**yorinasub17** commented *Aug 29, 2020*

This bumps `urfave/cli` to v2 track. This is an opportunistic upgrade (we're not using any of the features yet) so that we can stay up to date as bugs are fixed.
<br />
***


**yorinasub17** commented *Aug 31, 2020*

> Is that worth taking on right now?

We don't really have to take on migrating right now (after all, we could just hold the version back in each respective CLI). I thought it would be nice to do this though so that new CLIs will be on the v2 version.
***

**yorinasub17** commented *Aug 31, 2020*

> args vs options order

I believe most of our commands/CLI that we have that use this will not be sensitive to this. But I could be wrong.
***

**osulli** commented *Oct 28, 2020*

@yorinasub17 what do you feel your next step will be?

I'm on the fence of using entrypoint for my own Terraform CLI wrapper as it is still using urfave/cli/v1 as opposed to v2. 

E: Using this branch
```
# go.mod
require (
	github.com/gruntwork-io/gruntwork-cli v0.7.1-0.20200831164626-978768fef544 // https://github.com/gruntwork-io/gruntwork-cli/pull/33
)
```
***

**yorinasub17** commented *Jun 22, 2021*

UPDATE: rebased on `master`.
***

**infraredgirl** commented *Feb 28, 2022*

We are looking to upgrade the Patcher CLI to `v2` of `urfave/cli`, and it seems this PR needs to get merged first. @yorinasub17 can I help push this over the line? Is only rebasing and fixing conflicts needed or something else as well? Happy to take over this PR and shepherd it to completion!
***

**yorinasub17** commented *Feb 28, 2022*

Just rebased on `master` and it looks like build passed and no more merge conflict, so this is ready for merge again! (provided you don't find anything else that needs updating).
***

**yorinasub17** commented *Feb 28, 2022*

Thanks for review! Merging now.
***

**yorinasub17** commented *Feb 28, 2022*

Released as https://github.com/gruntwork-io/go-commons/releases/tag/v0.11.0
***

