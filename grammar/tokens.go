package token

type TokenType string

type Token struct{
	Type TokenType
	Value string
}

const (
	
	EOF = "EOF"
	ILLEGAL ="ILLEGAL"//illegal token
	SEMICOLON = ";"
	
	LBRACKET ="["//left bracket
	RBRACKET ="]"
	LBRACE = "{"
    RBRACE = "}"
	LPAREN = "("
	RPAREN = ")"
	
	//OPERATORS
    EQ  = "==" //equal sign
	PL_EQ = "+="  //plus equal
    MI_EQ  = "-= "//minus equal
	PP = "++"  //plus plus
    MM = "--" //minus minus
    NE = "!=" //not equal
	MULT_EQ = "*="
	DEVIDE_EQ = "/="
	POWER_EQ = "*="
	
	//logical operators
    AND = "&& "// and &&
    OR = "||"  //or ||
	
	//conditional operators
    IF = "if"//if
    ELSE = "else"//else
   
	//for iteration
	FOR  = "for"//for
    WHILE = "while"//while
    FOREACH  = "foreach"//for each
	
	//trulability
	TRUE = "true"
	FALSE = "false"


	FUNCTION = "FUNCTION"//function
	LET = "LET"//a useless keyword which I saw in tutorial and it fucking impressed me))

	//Operatos
	ASIGN = "="
	PLUS = "+"
	MINUS = "-"
	MULTIPLY = "*"
	DIVISION = "/"
	NOT = "!"
	POWER = "^"

	//Identifiers
	IDENT = "IDENT"
	INT = "INT"
	STRING = "STRING"
	RETURN = "RETURN"
	
	COMMA = ","
	DOT = "."

	//comparators
	SMALLER = "<"
	BIGGER = ">"
	SMALLEREQUAL = "<="
	BIGGEREQUAL = ">="
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let":LET,
	"if":IF,
	"else":ELSE,
	"while":WHILE,
	"for":FOR,
	"foreach":FOREACH,
	"true":TRUE,
	"false":FALSE,
	"return":RETURN}

//check where a keyword is a token or return identifier if token not present
func LookUpIdent(ident string) TokenType{
	if tok, ok := keywords[ident];ok{
		return tok
	}
	return IDENT
}
