# Add version package

**levkohimins** commented *Dec 14, 2020*

I would want to add only one package `version`. To be honest I don't like keep VERSION in main.go.
Reasons:
1. VERSION - uppercase 
2. This forces to always add an excess arg `cli.NewApp(VERSION)`

It's easy to make DevOps compatible with both implementations. Just try to rewrite in both places:
`-ldflags "-X main.VERSION=v1.0.0 -X github.com/gruntwork-io/gruntwork-cli/version.VERSION=v1.0.0"`
<br />
***


**levkohimins** commented *Dec 14, 2020*

Exactly. When it's in a separate package, we have the ability to call `version.Version()` from different parts of the code, without worrying about where the magic actually happens.

Done!
***

**brikis98** commented *Dec 14, 2020*

Tests passed! Merging now.
***

**brikis98** commented *Dec 14, 2020*

https://github.com/gruntwork-io/gruntwork-cli/releases/tag/v0.7.1
***

