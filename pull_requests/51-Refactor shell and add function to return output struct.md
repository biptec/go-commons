# Refactor shell and add function to return output struct

**yorinasub17** commented *Jul 15, 2021*

This PR refactors the functions defined in the `shell` module to use a common runner. In addition, a new function that returns the stdout/stderr/interleaved outputs as a struct is introduced. The latter will be used in a feature I am working on where I need to piece out the three so they can be stored as files.

NOTE: This was implemented by porting over similar functionality from terratest, which used to use a similar approach as what we have, but was updated after we encountered some bugs and unreliabilities.
<br />
***


**yorinasub17** commented *Jul 16, 2021*

Thanks for review! Will merge this in now.
***

