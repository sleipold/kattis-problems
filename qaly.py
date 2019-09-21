description_period = []
sum_QALY = 0

num_periods = input()

for _ in range(int(num_periods)):
    description_period.append(input().split())

for period in description_period:
    sum_QALY += float(period[0]) * float(period[1])

print(sum_QALY)