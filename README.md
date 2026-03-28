# sqlc-gen-go extensions

This fork of sqlc-gen-go introduces multiple options to export generated files to different directories and packages. It also allows sqlc to generate the `params` and `rows` separately, which was not possible before.

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
