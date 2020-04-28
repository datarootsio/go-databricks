# Configuration

> see https://aka.ms/autorest

``` yaml
azure-arm: false
```

## Clusters

``` yaml $(clusters)
input-file: clusters.yaml

go:
  namespace: clusters
  output-folder: clusters
  package-name: github.com/innovationnorway/go-databricks/clusters
  user-agent: 'go-databricks/version clusters'
```

## DBFS

``` yaml $(dbfs)
input-file: dbfs.yaml

go:
  namespace: dbfs
  output-folder: dbfs
  package-name: github.com/innovationnorway/go-databricks/dbfs
  user-agent: 'go-databricks/version dbfs'
```

## Groups

``` yaml $(groups)
input-file: groups.yaml

go:
  namespace: groups
  output-folder: groups
  package-name: github.com/innovationnorway/go-databricks/groups
  user-agent: 'go-databricks/version groups'
```

## Workspace

``` yaml $(workspace)
input-file: workspace.yaml

go:
  namespace: workspace
  output-folder: workspace
  package-name: github.com/innovationnorway/go-databricks/workspace
  user-agent: 'go-databricks/version workspace'
```

## Secrets

``` yaml $(secrets)
input-file: secrets.yaml

go:
  namespace: secrets
  output-folder: secrets
  package-name: github.com/innovationnorway/go-databricks/secrets
  user-agent: 'go-databricks/version secrets'
```

## Libraries

``` yaml $(libraries)
input-file: libraries.yaml

go:
  namespace: libraries
  output-folder: libraries
  package-name: github.com/innovationnorway/go-databricks/libraries
  user-agent: 'go-databricks/version libraries'
```

## Jobs

``` yaml $(jobs)
input-file: jobs.yaml

go:
  namespace: jobs
  output-folder: jobs
  package-name: github.com/innovationnorway/go-databricks/jobs
  user-agent: 'go-databricks/version jobs'
```
