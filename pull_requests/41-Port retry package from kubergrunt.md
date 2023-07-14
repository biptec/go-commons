# Port retry package from kubergrunt

**yorinasub17** commented *Mar 17, 2021*

This DRYs the `retry` package from `kubergrunt` so that it can be reused across CLIs (e.g., most recent use case is `cloud-nuke`: https://github.com/gruntwork-io/cloud-nuke/pull/183)
<br />
***


**yorinasub17** commented *Mar 17, 2021*

These are exactly the same as the terratest functions (the `E` flavors). I only ported the things that we need to avoid adding and maintaining code we don't use.
***

**brikis98** commented *Mar 17, 2021*

Hm, wouldn't that be more confusing? That is, if we port all the retry functions, then we can remove them from all the other places right away. Whereas if we port only a few, then when someone goes to update the other locations, some items will be missing, and it won't be obvious if that's intentional or not, or if there are other differences, and we might never migrate over...
***

**yorinasub17** commented *Mar 17, 2021*

> That is, if we port all the retry functions, then we can remove them from all the other places right away.

This is where I am confused. AFAIK we don't have these functions outside of `kubergrunt` (which was why it was copied over there) in a CLI context, and I decided to move the kubergrunt functions here because of https://github.com/gruntwork-io/cloud-nuke/pull/183.
***

**yorinasub17** commented *Mar 17, 2021*

By the way, the reason I am pushing back on this is that it is some amount of non-zero effort to port the functions over from terratest, since we have to sanitize the logging and the `t` object. Vs copying over from `kubergrunt`, which already sanitized these, is much easier.

I don't quite have the time to do more than what's here before I am off and I didn't want to block the `cloud-nuke` PR, which was why I did the minimal effort here.

Just wanted to provide the context. Happy to keep this PR open until we've sanitized and ported all the functions from terratest if you feel strongly, just that that might make this PR linger.
***

**yorinasub17** commented *Mar 17, 2021*

> AFAIK we don't have these functions outside of kubergrunt

Ok I lied. We have one in terragrunt as well: https://github.com/gruntwork-io/terragrunt/blob/master/util/retry.go, but that should be covered by this function.

But outside of that use case, I don't see any reference for these funcs outside testing contexts from my locally checked out repos (which is almost all repos in our org) or github org search.
***

**yorinasub17** commented *Mar 17, 2021*

Thanks for review and sharing the concerns and sanity check! Going to merge and release this now.
***

