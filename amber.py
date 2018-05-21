#!/usr/bin/python

from os import system
from sys import argv

def lex(data):
    i = 0
    tok = ""
    state = 0
    tokens = []
    string = ""
    while i < len(data):
        tok += data[i]
        if tok == " " and state == 0:
            tok = ""
        elif tok == "\n":
            tok = ""
        elif tok.lower() == "print":
            tokens.append("[PRINT]")
            tok = ""
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
        elif state == 1:
            string += tok
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
                    

        i+=1
    code += " halt "
    return code


def run():
    if len(argv) < 3:
        print "Error, missing input or output file name."
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
            code = parse(toks)
            try:
                open(argv[2], "w").write(code)
                system("chmod +x " + argv[2])
            except:
                print "Error : couldn't write executable file named \"" + argv[2] + "\""


run()