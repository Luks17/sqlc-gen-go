# sqlc-gen-go separate models file

This fork of sqlc-gen-go introduces exporting models file to a different package.

There is a related issue on the sqlc repository:
[github.com/sqlc-dev/sqlc/issues/835](https://github.com/sqlc-dev/sqlc/issues/835)

## Added Options

- `base_import_path`:
  - Import path of the go module including the base directory for the generated files (out). Required when using separate packages for any file.
- `output_directory`:
  - Directory path for the generated files. Used when we want to extend the `out` option. Defaults to the value of `out` option.
- `output_models_directory`:
  - Directory path for the models file. Used when models file will be placed in a different directory than `output_directory`. Defaults to the value of `output_directory` option.
- `output_models_package`:
  - Package name of the models file. Used when models file is in a different package. Defaults to value of `package` option.
- `output_params_file_name`:
  - File name for params structs file. Defaults to `params.go`.
- `output_params_directory`:
  - Directory path for the params structs file. Defaults to the value of `output_directory` option.
- `output_params_package`:
  - Package name of the params structs file. Defaults to value of `package` option.
- `output_row_results_file_name`:
  - File name for row result structs file. Defaults to `row_results.go`.
- `output_row_results_directory`:
  - Directory path for the row result structs file. Defaults to the value of `output_directory` option.
- `output_row_results_package`:
  - Package name of the row result structs file. Defaults to value of `package` option.
- `output_querier_directory`:
  - Directory path for the querier file. Used when querier file will be placed in a different directory than `output_directory`. Defaults to the value of `output_directory` option.
- `output_querier_package`:
- Package name of the querier file. Used when querier file is in a different package. Defaults to value of `package` option.

## How to use for separate models file

Lets say you want to export models to `internal/business/entities/database.go` file and keep other
generated files in `internal/sqlcrepo/` directory. You can use the following configuration:

```yaml
version: "2"
plugins:
  - name: golang
    wasm:
      url: https://github.com/berk-karaal/sqlc-gen-go/releases/download/v1.5.1/berk-karaal-sqlc-gen-go_1.5.1.wasm
      sha256: 95bc2009c94bdac0f8c5af8207bd1cf43723f8e33fb17e06fce1d96b83da242e
sql:
  - engine: "postgresql"
    queries: "query.sql"
    schema: "schema.sql"
    codegen:
      - plugin: golang
        out: "internal"  # This is the base directory for the generated files
        options:
          sql_package: "pgx/v5"
          package: "sqlc"  # Default package name for the generated files
          emit_interface: true
          output_directory: "sqlc" # This is the directory for the generated files extending `out`, resulting in `internal/sqlc`
          output_models_package: "entities"  # Package name that should be used in `output_models_file_name` file
          output_models_directory: "business/entities" # Directory path for the models file, extends `out` resulting in `internal/business/entities`
          output_querier_package: "queries"  # Package name that should be used in `output_querier_file_name` file
          output_querier_directory: "business/repository" # Directory path for the querier file, extends `out` resulting in `internal/business/repository`
```
