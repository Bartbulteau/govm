#!/usr/bin/python

from os import system
from sys import argv
from re import match

def lex(data):
    i = 0
    tok = ""
    state = 0
    tokens = []
    string = ""
    num = ""
    while i < len(data):
        tok += data[i]
        # I/O
        if tok.lower() == "print":
            tokens.append("[PRINT]")
            tok = ""
        

        # NUMBERS
        elif match("([0-9]+)", tok) != None and state == 0:
            state = 2
            num += tok
            tok = ""

        # STRINGS
        elif tok == "\"":
            if state == 0:
                state = 1
                tok = ""
            elif state == 1:
                state = 0
                tokens.append("[STRING]")
                tokens.append(string)
                string = ""
                tok = ""
        
        # STATES
        elif state == 1:
            string += tok
            tok = ""
        elif state == 2:
            if tok == " " or tok == "\n" or i == len(data) - 1:
                state = 0
                tok = ""
                tokens.append("[NUM]")
                tokens.append(num)
                num = ""
            else:
                num += tok
                tok = ""
        
        elif tok == " " and state == 0:
            tok = ""
        elif tok == "\n" and state == 0:
            tok = ""

        i+=1
    return tokens


def parse(toks):
    i = 0
    code = "#!./govm\n"
    while i < len(toks):

        if toks[i] == "[PRINT]":
            i += 1
            if toks[i] == "[STRING]":
                i += 1
                code += "mov ax 1 "
                for char in toks[i]:
                    code = code + "mov bx " + str(ord(char)) + " printr bx "
                code = code + "mov bx 10 printr bx "
            if toks[i] == "[NUM]":
                i += 1
                code += "mov ax 0 "
                code = code + "mov bx " + toks[i] + " printr bx mov ax 1 mov bx 10 printr bx "
                    

        i+=1
    code += " halt "
    return code


def run():
    if len(argv) < 2:
        print "Error, missing input file name."
    else:
        if argv[1] == "help":
            print "Use : amber [source] [target]"
        else:
            try:
                file_handler = open(argv[1], "r")
            except:
                print "Error invalid file name : \"" + argv[1] + "\""
            
            
            data = file_handler.read()
            toks = lex(data)
            print toks
            code = parse(toks)
            try:
                if len(argv) > 2:
                    open(argv[2], "w").write(code)
                    system("chmod +x " + argv[2])
                else:
                    open("a.out", "w").write(code)
                    system("chmod +x a.out")
            except:
                print "Error : couldn't write executable file named \"" + argv[2] + "\""


run()