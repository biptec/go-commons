# Configurable stack traces

**yorinasub17** commented *Nov 28, 2018*

The stack traces are very useful, but as a customer facing CLI, it can be rather verbose and is useless unless they know the source code.

This introduces the concept of a debug mode for configuring when to print out the stack trace. Now the command will only print out stack traces when run with `GRUNTWORK_DEBUG=1`.

Note: I couldn't find a good way to unit test this, given that `checkForErrorsAndExit` kills the process. For now, here is the manual test demo:

Without env:
```
%~> go run main.go                                                                                                                                        [0]
ERRO[2018-11-28T11:31:10-08:00] --aws-region is required                      error="--aws-region is required"
exit status 1
```

With env:
```
%~> GRUNTWORK_DEBUG=1 go run main.go                                                                                                                      [1]
ERRO[2018-11-28T11:31:21-08:00] *errors.errorString --aws-region is required
/Users/yoriy/go/src/github.com/gruntwork-io/package-k8s/modules/k8s-binaries/cmd/eks-configure-kubectl/main.go:111 (0x114eb33)
        run: return gruntworkErrors.WithStackTrace(err)
/Users/yoriy/go/src/github.com/gruntwork-io/package-k8s/modules/k8s-binaries/vendor/github.com/urfave/cli/app.go:490 (0x11225c8)
        HandleAction: return a(context)
/Users/yoriy/go/src/github.com/gruntwork-io/package-k8s/modules/k8s-binaries/vendor/github.com/urfave/cli/app.go:264 (0x112083d)
        (*App).Run: err = HandleAction(a.Action, context)
/Users/yoriy/go/src/github.com/gruntwork-io/package-k8s/modules/k8s-binaries/vendor/github.com/gruntwork-io/gruntwork-cli/entrypoint/entrypoint.go:33 (0x11440f6)
        RunApp: err := app.Run(os.Args)
/Users/yoriy/go/src/github.com/gruntwork-io/package-k8s/modules/k8s-binaries/cmd/eks-configure-kubectl/main.go:174 (0x114ef3c)
        main: entrypoint.RunApp(app)
/usr/local/Cellar/go/1.11.2/libexec/src/runtime/proc.go:201 (0x102aba7)
        main: fn()
/usr/local/Cellar/go/1.11.2/libexec/src/runtime/asm_amd64.s:1333 (0x1055a51)
        goexit: BYTE    $0x90   // NOP  error="--aws-region is required"
exit status 1
```
<br />
***


**josh-padnick** commented *Nov 28, 2018*

This is great! Should we just reduce the default logging to:

```
%~> go run main.go                                                                                                                                        
ERROR: --aws-region is required
```
***

**yorinasub17** commented *Nov 28, 2018*

+1 to @josh-padnick suggestion, so implemented it! I think this now closes https://github.com/gruntwork-io/gruntwork-cli/issues/7
***

**yorinasub17** commented *Nov 28, 2018*

Ok merging. Will release this with the other 2.
***

