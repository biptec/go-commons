# logging.GetLogger does not do what we think it does

**yorinasub17** commented *Nov 28, 2018*

`WithField` returns a [`logrus.Entry`](https://github.com/sirupsen/logrus/blob/master/entry.go#L50) which contains the field information. We then return the attached `Logger` object, but [`logrus.Logger`](https://github.com/sirupsen/logrus/blob/master/logger.go#L11) does not actually point back to the `Entry`. So we lose the Field information on return.
<br />
***


