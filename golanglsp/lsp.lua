vim.api.nvim_create_autocmd('FileType', {
  -- This handler will fire when the buffer's 'filetype' is "python"
  pattern = 'markdown',
  callback = function(ev)
    vim.lsp.start({
      name = 'golanglsp',
      cmd = {'/home/yabsera/Documents/github/golang-project/golanglsp/main'},
      -- Set the "root directory" to the parent directory of the file in the
      -- current buffer (`ev.buf`) that contains either a "setup.py" or a
      -- "pyproject.toml" file. Files that share a root directory will reuse
      -- the connection to the same LSP server.
      
    })
  end,
})
