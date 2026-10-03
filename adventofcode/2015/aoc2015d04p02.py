from hashlib import md5

def calculate(input, salt):
    input = f'{input}{salt}'.encode()
    return md5(input).hexdigest()


def logic(input):
    salt = 1
    while True:
        result = calculate(input, salt)
        if result[:6] == '000000':
            return salt
        salt+=1
        

def execute(input):
    print("\n=====================================================")
    print('--- Day 4: The Ideal Stocking Stuffer - Part two ---')
    print(' Now find one that starts with six zeroes.')
    
    return logic(input)
        