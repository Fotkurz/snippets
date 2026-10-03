
def coord(x, y):
    return f"{x}:{y}"


def logic(input):
    how_many_houses = 1

    init_pos = {
        coord(0, 0): 1,
        'x_pos': 0,
        'y_pos': 0,
    }
    
    santa = {
        'name': 'santa',
        **init_pos
    }

    robo_santa = {
        'name': 'robo_santa',
        **init_pos,
    }

    # print(f"starting at {x}:{y}")
    for i in range(len(input)):
        # if pair, santa moves, else robo_santa moves
        if i % 2 == 0:
            turn = santa
        else:
            turn = robo_santa

        x, y = turn['x_pos'], turn['y_pos']
        
        actual_pos = coord(x, y)
        
        if input[i] == '^':
            y += 1
        elif input[i] == 'v':
            y -= 1
        elif input[i] == '>':
            x += 1
        elif input[i] == '<':
            x -= 1

        new_pos = coord(x, y)
        
        # print(f"Moving {turn['name']} from {actual_pos} to {new_pos}")
        
        turn['x_pos'], turn['y_pos'] = x, y

        pos = turn.get(new_pos, None)
        if pos is None:
            # print("new house")
            turn[new_pos] = 1
        else:
            # print("already been here")
            turn[new_pos] += 1
        # if pair, santa moves, else robo_santa moves
        if i % 2 == 0:
            santa = turn
        else:
            robo_santa = turn

    merge = robo_santa | santa

    how_many_houses = len(merge) - 3 # counting number of keys - (name, x_pos and y_pos)
    return how_many_houses


def execute(input):
    print("\n=====================================================")
    print('--- Day 3: Perfectly Spherical Houses in a Vacuum - Part two ---')
    print(' How many houses receive at least one present?')
    
    return logic(input)
    