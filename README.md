workglow is a Go library for defining and running go workflows. Every step in the workglow
should be mapped to a workglow Go handler, created in a way it accepts
a subset of workglow step's named parameters as input arguments, i.e:

```
workglow:
  name: example_workflow
  steps:
    initialize:
      name: workglow/step1.go
    configure_iam_role:
      name: workglow/aws/iam.go
      role_arn: ${example_workflow.initialize.arn} 
      policy: ${example_workflow.initialize.arn}
    configure_k8s_rbac:
      role_arn: ${example_workflow.initialize.arn} 
```

All workglow steps are ran concurrently and simultaneously unless a step has
an implicit dependency on output parameters of other steps, ie

```
role_arn: ${example_workflow.initialize.arn}
```

In this case the step is ran afer output parameter becomes available, ie
after step runs that produce all the needed paramters.

Steps Go handlers are implemented as Go plugins with all the limitations for 
this featutre. Please read https://pkg.go.dev/plugin for more information.


