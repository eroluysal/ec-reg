# Basic EditorConfig CLI tool

- [Arguments](#arguments)
- [Usage](#usage)
- [License](#license)

<a name="#arguments"></a>
### Arguments

**`-c` charset:**
Set to latin1, utf-8, utf-8-bom, utf-16be or utf-16le to control the character set. Use of utf-8-bom is discouraged. 
*(default "UTF-8")*

**`-e` end_of_line:**
Set to lf, cr, or crlf to control how line breaks are represented. 
*(default "lf")*

**`-f` insert_final_newline:**
Set to true to ensure file ends with a newline when saving and false to ensure it doesn't. 
**(default true)**

**`-i` indent_size:**
A whole number defining the number of columns used for each indentation level and the width of soft tabs (when supported). When set to tab, the value of tab_width (if specified) will be used. 
*(default 4)*

**`-r` root:**
Special property that should be specified at the top of the file outside of any sections. Set to true to stop .editorconfig files search on current file. 
*(default true)*

**`-s` indent_style:**
Set to tab or space to use hard tabs or soft tabs respectively. 
*(default "space")*

**`-t` trim_trailing_whitespace:**
Set to true to remove any whitespace characters preceding newline characters and false to ensure it doesn't. 
*(default true)*

<a name="#license"></a>
### Usage


<a name="#license"></a>
### License

    MIT License
    
    Copyright (c) 2017 Erol Uysal
    
    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:
    
    The above copyright notice and this permission notice shall be included in all
    copies or substantial portions of the Software.
    
    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
    SOFTWARE.
