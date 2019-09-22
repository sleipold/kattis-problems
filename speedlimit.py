n = int(input())

while n != -1:
    val_pairs = []
    for i in range(n):
        val_pairs.append(input().split())
    miles = 0
    s = int(val_pairs[0][0])
    t = int(val_pairs[0][1])
    miles = s * t
    for j in range(len(val_pairs)):
        if j+1 < len(val_pairs):
            t = int(val_pairs[j+1][1]) - int(val_pairs[j][1])
            s = int(val_pairs[j+1][0])
            miles += (s * t)
    print(str(miles) + " miles")
    n = int(input())
