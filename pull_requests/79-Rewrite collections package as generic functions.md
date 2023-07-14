# Rewrite collections package as generic functions

**infraredgirl** commented *Nov 8, 2022*

This also bumps the Go version to 1.18.

Note that this would be a backwards incompatible change due to Go v1.18 now being required.

~~Unfortunately there's a failing test that I haven't had time to resolve.~~ Solved: tests are now passing.
<br />
***


**infraredgirl** commented *Nov 9, 2022*

Ok, tests are now passing, and I added a test case to verify that using a custom type works.

This is ready for a final review.
***

**infraredgirl** commented *Nov 11, 2022*

Thanks, will merge and release.
***

