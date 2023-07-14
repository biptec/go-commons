# Make GetPathRelativeTo() resolve symbolic links

**infraredgirl** commented *Oct 25, 2022*

Related to #75.

Without this change the relative paths are not computed correctly in case when `basePath` is a subfolder of `path`, and one  of them is a symbolic link of the other.

Note that I had to create fixtures for test cases of `GetPathRelativeTo()`, because the `filepath.EvalSymlinks()` method that we're now using is actually looking for files on disk.
<br />
***


**infraredgirl** commented *Oct 25, 2022*

Thanks for the review, Zack! Merging.
***

