#!/usr/bin/env python

samples = ['1 + 2 * 3 + 4 * 5 + 6',
        '1 + (2 * 3) + (4 * (5 + 6))',
        '2 * 3 + (4 * 5)',
        '5 + (8 * 3 + 9 + 3 * 4 * 3)',
        '5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))',
        '((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2']

def main():
    #for s in samples:
    #    t = tokenize(s)
    #    pf = infix_to_postfix(t)
    #    print(repr(t), repr(pf), eval_postfix(pf))
    with open('input.txt','r') as f:
        lines = [l.strip() for l in f.readlines()]

    sum = 0
    for l in lines:
        t = tokenize(l)
        pf = infix_to_postfix(t)
        sum+=eval_postfix(pf)

    print(sum)

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

def infix_to_postfix(tokens):
    opstack = []
    output = []

    for tok in tokens:
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
        elif tok=='+' or tok=='*':
            while opstack:
                v = opstack.pop()
                if v=='(':
                    opstack.append(v)
                    break
                output.append(v)
            opstack.append(tok)

    while opstack:
        output.append(opstack.pop())

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
