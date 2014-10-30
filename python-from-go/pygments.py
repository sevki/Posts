from pygments import highlight
from pygments.lexers import get_lexer_by_name
// START1 OMIT
from pygments.formatters import HtmlFormatter
// END1 OMIT

lexer = get_lexer_by_name("python")
// START2 OMIT
formatter = HtmlFormatter(linenos=True, encoding="utf-8")
// END2 OMIT
// START3 OMIT
result = highlight(code, lexer, formatter)
// END3 OMIT
