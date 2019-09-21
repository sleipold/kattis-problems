import math

input_splitted = input().split()
num_matches = int(input_splitted[0])
width = int(input_splitted[1])
height = int(input_splitted[2])
diagonal = math.sqrt(width*width + height*height)

for _ in range(num_matches):
    len_match = int(input())
    if len_match <= diagonal:
        print('DA')
    elif len_match <= width:
        print('DA')
    elif len_match <= height:
        print('DA')
    else:
        print('NE')