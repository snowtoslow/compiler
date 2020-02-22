type Lexer struct{
	
	input string
	position int//current position 
	readPosition int//to read next char of after input
	char byte//points to char in input
}


func New(input string) *Lexer{
	
	lexer := &Lexer{input : input}
	lexer.readChar()
	return lexer

}

//read input string and iterate over it to check if its null otherwise iterate over input;
func (lexer *Lexer) readChar(){
	
	if lexer.readPosition >= len(lexer.input){
		
		lexer.ch = 0
	
	}else{

		lexer.ch = lexer.input[l.readPosition]
	}

	lexer.position = lexer.readPosition

	lexer.readPosition += 1
}


//check if the tokens exist,default it try to chekc if its a identifier else return null;
func (lexer *Lexer) nextToken() token.Token{
	
	var tok token.Token
	
	lexer.skipWhitespace();

	switch lexer.ch{
	
		case '[':
			tok = newToken(token.LBRACKET,lexer.ch)
		
		case ']':
			tok = newToken(token.RBRACKET,lexer.ch)
		
		case '{':
			tok = newToken(tokeN.LBRACE,lexer.ch)
		
		case '}':
			tok = newToken(token.RBRACE,lexer.ch)
		
		case '(':
			tok = newToken(token.LPAREN,lexer.ch)
		
		case ')':
			tok = newToken(token.RPAREN,lexer.ch)


		case '=':
			if peekChar() == '='{
				ch := lexer.ch
				lexer.readChar()
				tok = token.Token{Type : token.EQ, Value : string(ch)+string(lexer.ch)}
			
			}
			else {
				tok = newToken(toke.ASSIGN,lexer.ch)
			
			}
		
		case '+':
			if peekChar() == '='{
				ch := lexer.ch
				lexer.readChar()
				tok = token.Token{Type : token.PL_EQ,Value : string(ch)+string(lexer.ch)
			}else if peekChar() == '+'{
				ch := lexer.ch
				lexer.readChar()
				tok = token.Token{Type : token.PP,Value : string(ch) + string(lexer.ch)}
			}else{

				tok = newToken(toke.PLUS,lexer.ch)
		
			}
		}
		
		case '-':
			if peekChar() == '='{
				ch := lexer.ch
				lexer.readChar()
				tok = token.Token{Type : token.MI_EQ, Value : string(ch)+string(lexer.ch)
			}else if peekChar == '-'{
				ch := lexer.ch
				lexer.readChar()
				tok = token.Token{Type : token.MM, Value : string(ch)+string(lexer.ch)}
			}
			else{
				
				tok = newToken(token.MINUS,lexer.ch)
			
			}
			
		}

		case '*':
			if peekChar() == '='{
				ch := lexer.ch
				lexer.readChar()
				tok = token.Token{Type : token.MULT_EQ,Value : string(ch)+string(lexer.ch)
			}else{
				tok = newToken(token.MULTIPLY,lexer.ch)
			}
		}
			
		
		case '/':
			if peekChar() == '='{
				ch := lexer.ch
				lexer.readChar()
				tok = token.Token{Type : token.DEVIDE_EQ, Value : string(ch) + string(lexer.ch)
			}else{
				tok = newToken(token.DIVISION,lexer.ch)
			}
		}
			

		case ',':
			tok = newToken(token.COMMA,lexer.ch)

		case '.':
			tok = newToken(token.DOT,lexer.ch)
		
		case '!':
			if lexer.peekChar() == '='{
				ch := lexer.ch
				lexer.readChar()
				tok = token.Token{Type : token.NE, Value : string(ch)+string(lexer.ch)
			}else{
				
				tok = newToken(token.NOT,lexer.ch)
		
			}
		}
			
		case '<':
			if peekChar() == '='{
				ch := lexer.ch
				lexer.readChar()
				tok = token.Token{Type : token.SMALLEREQUAL, Value : string(ch)+string(lexer.ch)
			}else{
				tok = newToken(token.SMALLER,lexer.ch)
			}
		}

			tok = newToken(token.SMALLER,lexer.ch)
		
		case '>':
			if peekChar() == '='{
				ch := lexer.ch
				lexer.readChar()
				tok = token.Token{Type : token.BIGGEREQUAL, Value : string(ch) + string(lexer.ch)}
			}else{
				tok = newToken(toke.BIGGER,lexer.ch)
			}
		}
		case '^':
			if peekChar() == '='{
				ch := lexer.ch
				lexer.readChar()
				tok = token.Token{Type : token.POWER_EQ, Value : string(ch) + string(lexer.ch)
			}else{
				tok = newToken(toke.POWER,lexer.ch)
			}
		}
		
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF		
		
		}
	default:
		if ifLetter(lexer.ch){
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookUpIdent(tok.Literal)
			return tok
		} else if isDigit(lexer.ch) {  
			
			tok.Type = token.INT
			
			tok.Literal = lexer.readNumber()
			
			return tok
		} else{
			
			tok = newToken(token.ILLEGAL,l.ch)

		}

	}

	lexer.readChar()
		
	return token
}

//function to create tokens
func newToken(tokenType token.TokenType,ch byte) token.Token{
	return token.Token{Type:tokenType, Value:string(ch)}
}


//read Identifier
func(lexer *Lexer) readIdentifier() string{
	
	position := lexer.position

	for isLetter(lexer.ch){
		lexer.readChar()
	}

	return lexer.input[position : lexer.position]
}

//check if it is a letter;
func isLetter(ch byte) bool{
	
	return 'a' <= ch && ch <= 'z' || 'A' <= ch <= 'Z' || ch == '_'

}

func(lexer *Lexer) skipWhitespace(){
	for lexer.ch == ' ' || lexer.ch == "\t" || lexer.ch == "\r"{

		lexer.readChar()

	}

}

func (lexer *Lexer) readNumber() string{
	
	position : = lexer.position

	for isDigit(lexer.ch){
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]

}

//check if character if a digit
func isDigit(ch byte) bool {
	return '0'<=ch && ch <='9'
}

//helper function to work with operators like !=,==,-= etc
//help us to identify difference between a single sign and where is mroe then one char  
func (lexer *Lexer) peekChar(){
	
	if lexer.readPosition >= len(lexer.input){
		
		return 0
	
	}else{
	
		return lexer.input[lexer.readPosition]
	
	}


}

