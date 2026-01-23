package golang

type OutputFile string

const (
	OutputFileModel     OutputFile = "modelFile"
	OutputFileQuery     OutputFile = "queryFile"
	OutputFileDb        OutputFile = "dbFile"
	OutputFileInterface OutputFile = "interfaceFile"
	OutputFileCopyfrom  OutputFile = "copyfromFile"
	OutputFileBatch     OutputFile = "batchFile"
	OutputFileParams    OutputFile = "paramsFile"
	OutputFileRowResult OutputFile = "rowResultFile"
)
