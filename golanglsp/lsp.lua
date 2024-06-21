vim.lsp.start({
	name = "golanglsp",
	cmd = { "./main" },
	root_dir = vim.fn.getcwd(),
})



