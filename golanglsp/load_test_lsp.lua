
local client=vim.lsp.start_client {
    name="golanglsp",
    cmd={"./main"},

}


if not client then 
    vim.notify "hey you didn't start the server thingy"
   return 
end

local buf_attach = function()
  vim.lsp.buf_attach_client(0, client)
end

vim.api.nvim_create_autocmd('FileType', {
  pattern = "markdown",
  callback = buf_attach,
})
