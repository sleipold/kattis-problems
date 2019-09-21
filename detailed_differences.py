num_cases = input()
cases = []

for _ in range(int(num_cases)):
    cases.append((input(), input()))

for a, b in cases:
    result = ''
    print(a + '\n' + b)

    for i in range(len(a)):
        if a[i] == b[i]:
            result += '.'
        else:
            result += '*'

    print(result)
    print()
