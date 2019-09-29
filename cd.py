#while True:
#    n, m = [int(x) for x in input().split()]
#
#    if n == 0 and m == 0:
#        break
#
#    cds_jack = set(int(input()) for _ in range(n))
#    cds_jill = set(int(input()) for _ in range(m))
#
#    print(len(cds_jack.intersection(cds_jill)))
#
# had to solve it with python2 for some reason
while True:
    n, m = map(int, raw_input().split())
    if n == 0 and m == 0:
        break
    cds_jack = set(int(raw_input()) for _ in range(n))
    ans = 0
    for _ in range(m):
        if int(raw_input()) in cds_jack:
            ans += 1
    print(ans)