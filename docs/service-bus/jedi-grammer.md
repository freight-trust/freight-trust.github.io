# 

expression = sub-expression / index-expression / comparator-expression
expression =/ or-expression / identifier expression =/ and-expression /
not-expression / paren-expression expression =/ "*" / multi-select-list
/ multi-select-hash / literal expression =/ function-expression /
pipe-expression / raw-string expression =/ current-node sub-expression =
expression "." ( identifier / multi-select-list / multi-select-hash /
function-expression / "*" ) pipe-expression = expression "|" expression
or-expression = expression "||" expression and-expression = expression
"&&" expression not-expression = "\!" expression paren-expression = "("
expression ")" index-expression = expression bracket-specifier /
bracket-specifier multi-select-list = "\[" ( expression *( ","
expression ) ) "\]" multi-select-hash = "{" ( keyval-expr \*( ","
keyval-expr ) ) "}" keyval-expr = identifier ":" expression
bracket-specifier = "\[" (number / "*" / slice-expression) "\]" / "\[\]"
bracket-specifier =/ "\[?" expression "\]" comparator-expression =
expression comparator expression slice-expression = \[number\] ":"
\[number\] \[ ":" \[number\] \] comparator = "\<" / "\<=" / "==" / "\>="
/ "\>" / "\!=" function-expression = unquoted-string ( no-args /
one-or-more-args ) no-args = "(" ")" one-or-more-args = "(" (
function-arg \*( "," function-arg ) ) ")" function-arg = expression /
expression-type current-node = "@" expression-type = "&" expression

raw-string = "'" *raw-string-char "'" raw-string-char = (%x20-26 /
%x28-5B / %x5D-10FFFF) / preserved-escape / raw-string-escape
preserved-escape = escape (%x20-26 / %28-5B / %x5D-10FFFF)
raw-string-escape = escape ("'" / escape) literal = "`" json-value "`"
unescaped-literal = %x20-21 / ; space \! %x23-5B / ; \# - \[ %x5D-5F / ;
\] ^ \_ %x61-7A ; a-z %x7C-10FFFF ; |}\~ …​ escaped-literal =
escaped-char / (escape %x60) number = \["-"\]1*digit digit = %x30-39
identifier = unquoted-string / quoted-string unquoted-string = (%x41-5A
/ %x61-7A / %x5F) *( ; A-Za-z* %x30-39 / ; 0-9 %x41-5A / ; A-Z %x5F / ;
\_ %x61-7A) ; a-z quoted-string = quote 1\_(unescaped-char /
escaped-char) quote unescaped-char = %x20-21 / %x23-5B / %x5D-10FFFF
escape = %x5C ; Back slash: \\ quote = %x22 ; Double quote: '"'
escaped-char = escape ( %x22 / ; " quotation mark U+0022 %x5C / ; \\
reverse solidus U+005C %x2F / ; / solidus U+002F %x62 / ; b backspace
U+0008 %x66 / ; f form feed U+000C %x6E / ; n line feed U+000A %x72 / ;
r carriage return U+000D %x74 / ; t tab U+0009 %x75 4HEXDIG ) ; uXXXX
U+XXXX

; The `json-value` is any valid JSON value with the one exception that
the ; `%x60` character must be escaped. While it’s encouraged that
implementations ; use any existing JSON parser for this grammar rule
(after handling the escaped ; literal characters), the grammar rule is
shown below for completeness::

json-value = false / null / true / json-object / json-array /
json-number / json-quoted-string false = %x66.61.6c.73.65 ; false null =
%x6e.75.6c.6c ; null true = %x74.72.75.65 ; true json-quoted-string =
%x22 1*(unescaped-literal / escaped-literal) %x22 begin-array = ws %x5B
ws ; \[ left square bracket begin-object = ws %x7B ws ; { left curly
bracket end-array = ws %x5D ws ; \] right square bracket end-object = ws
%x7D ws ; } right curly bracket name-separator = ws %x3A ws ; : colon
value-separator = ws %x2C ws ; , comma ws = \*(%x20 / ; Space %x09 / ;
Horizontal tab %x0A / ; Line feed or New line %x0D ; Carriage return )
json-object = begin-object \[ member \*( value-separator member ) \]
end-object member = quoted-string name-separator json-value json-array =
begin-array \[ json-value \*( value-separator json-value ) \] end-array
json-number = \[ minus \] int \[ frac \] \[ exp \] decimal-point = %x2E
; . digit1-9 = %x31-39 ; 1-9 e = %x65 / %x45 ; e E exp = e \[ minus /
plus \] 1*DIGIT frac = decimal-point 1\*DIGIT int = zero / ( digit1-9
\*DIGIT ) minus = %x2D ; - plus = %x2B ; + zero = %x30 ; 0
