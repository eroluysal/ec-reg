# Basic EditorConfig CLI tool

    -c string
        set to latin1, utf-8, utf-8-bom, utf-16be or utf-16le to control the character set. Use of utf-8-bom is discouraged. (default "UTF-8")
    -e string
        set to lf, cr, or crlf to control how line breaks are represented. (default "lf")
    -f	set to true to ensure file ends with a newline when saving and false to ensure it doesn't. (default true)
    -i int
        a whole number defining the number of columns used for each indentation level and the width of soft tabs (when supported). When set to tab, the value of tab_width (if specified) will be used. (default 4)
    -r	special property that should be specified at the top of the file outside of any sections. Set to true to stop .editorconfig files search on current file. (default true)
    -s string
        set to tab or space to use hard tabs or soft tabs respectively. (default "space")
    -t	set to true to remove any whitespace characters preceding newline characters and false to ensure it doesn't. (default true)