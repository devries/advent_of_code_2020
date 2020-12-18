#!/usr/bin/env python
import time

samples = ['1 + 2 * 3 + 4 * 5 + 6',
        '1 + (2 * 3) + (4 * (5 + 6))',
        '2 * 3 + (4 * 5)',
        '5 + (8 * 3 + 9 + 3 * 4 * 3)',
        '5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))',
        '((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2']

smallsleep = 0.5
bigsleep = 1.0

steps = 0

def main():
    #for s in samples:
    #    t = tokenize(s)
    #    pf = infix_to_postfix(t)
    #    print(repr(t), repr(pf), eval_postfix(pf))
    with open('input.txt','r') as f:
        lines = [l.strip() for l in f.readlines()]

    total = 0
    for l in lines:
        t = tokenize(l)
        print('\033[2J')
        print('\033[3J\033[H')
        pf = infix_to_postfix(t, total)
        result = eval_postfix(pf)
        total+=result
        detok = ' '.join(str(v) for v in pf)
        #print(f'{l} = {detok} = {result}')

    print(total)

def tokenize(statement):
    tokens = []

    number = []
    for c in statement:
        if c>='0' and c<='9':
            number.append(c)
        else:
            if len(number)>0:
                tokens.append(int(''.join(number)))
                number = []
            if c=='+':
                tokens.append('+')
            elif c=='*':
                tokens.append('*')
            elif c=='(':
                tokens.append('(')
            elif c==')':
                tokens.append(')')

    if len(number)>0:
        tokens.append(int(''.join(number)))

    return tokens

def infix_to_postfix(tokens, total):
    global smallsleep
    global bigsleep
    global steps
    opstack = []
    output = []

    for p, tok in enumerate(tokens):
        print('\033[2J')
        print('\033[3J\033[H')
        poutput = ' '.join(str(o) for o in output)
        pstack = ' '.join(str(s) for s in opstack)
        ptok = ' '.join(str(t) for t in tokens[p:])
        print("Total:        "+str(total))
        print("Postfix Form: "+poutput)
        print("Op Stack:     "+pstack)
        print("Statement:    "+ptok)
        time.sleep(smallsleep)
        if isinstance(tok, int):
            output.append(tok)
        elif tok=='(':
            opstack.append(tok)
        elif tok==')':
            while opstack:
                v = opstack.pop()
                if v=='(':
                    break
                output.append(v)
        elif tok=='*':
            while opstack:
                v = opstack.pop()
                if v=='(': # Pop everything of higher presidence onto output
                    opstack.append(v)
                    break
                output.append(v)
            opstack.append(tok)
        elif tok=='+':
            while opstack:
                v = opstack.pop()
                if v=='(' or v=='*': # Pop everything of higher presidence onto output,
                                     # Multiply is lower presidence
                    opstack.append(v)
                    break
                output.append(v)
            opstack.append(tok)

    while opstack:
        output.append(opstack.pop())

    print('\033[2J')
    print('\033[3J\033[H')
    poutput = ' '.join(str(o) for o in output)
    pstack = ' '.join(str(s) for s in opstack)
    ptok = ' '.join(str(t) for t in tokens[p:])
    print("Total:        "+str(total))
    print("Postfix Form: "+poutput)
    print("Op Stack:     "+pstack)
    print("Statement:")
    time.sleep(bigsleep)

    steps+=1
    if steps>=3:
        smallsleep=0.01
        bigsleep=0.1
    elif steps>=1:
        smallsleep=0.1
        bigsleep=0.2

    return output

def eval_postfix(tokens):
    stack = []

    for tok in tokens:
        if isinstance(tok, int):
            stack.append(tok)
        elif tok=='+':
            res = stack.pop()+stack.pop()
            stack.append(res)
        elif tok=='*':
            res = stack.pop()*stack.pop()
            stack.append(res)

    return stack.pop()

if __name__ == '__main__':
    main()
