
def logic(input):
    graph = {
        '0:0': 1
    }
    x, y = 0, 0
    print(f"starting at {x}:{y}")
    how_many_houses = 1
    for d in input:
        if d == '^':
            y += 1
        if d == 'v':
            y -= 1
        if d == '>':
            x += 1
        if d == '<':
            x -= 1

        coord = f'{x}:{y}'
        print(f"Moving to {coord}")
        pos = graph.get(coord, None)
        
        if pos is None:
            print("new house")
            how_many_houses+=1
            graph[f"{x}:{y}"] = 1
        else:
            print("already been here")
            graph[f"{x}:{y}"] += 1
            
    return how_many_houses

def execute(input):
    print("\n=====================================================")
    print('--- Day 3: Perfectly Spherical Houses in a Vacuum - Part One ---')
    print(' How many houses receive at least one present?')

    return logic(input)
    