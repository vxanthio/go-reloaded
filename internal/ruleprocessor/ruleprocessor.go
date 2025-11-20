package ruleprocessor

import "strconv"

func ProcessTokens(tokens []string) []string {
result:=[]string{}
for _,t:=range tokens   {

if t!="(hex)" {
	result:=append(result,t)

		continue
	}
}

	previous:=tokens[i-1]
	valid:=true
	if t=="(hex)" {
	for  _,p:=range previous{
if !(p>='0' && p<='9') || !(p>='A' && p<='F') {

	valid=false
	break
}
	}
	 
	value,err:=strconv.ParseInt(previous,16,64)
	if err!=nil {
		tokens[i]=""
		continue
	}
}
tokens[i]:=strconv.FormatInt(value,10)
tokens[i]
continue

	}
}

