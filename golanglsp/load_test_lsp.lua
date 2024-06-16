local client=vim.lsp.start_client {
    name="golanglsp",
    cmd={"/home/yabsera/Documents/github/golang-project/golanglsp/main"},
}

if not client then
    vim.notify "hey, you didn't do the client thingy "
    return
end

vim.api.nvim_create_autocmd("filename",{
    pattern="markdown",
    callback=function ()
        vim.lsp.buf_attach_client(0,client)
    end,
})