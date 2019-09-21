ratings = []
final_ratings = {}

for _ in range(5):
    ratings.append(input().split())

for i in range(len(ratings)):
    sum_rating = 0
    for grade in ratings[i]:
        sum_rating += int(grade)
    final_ratings[i+1] = sum_rating

max_key = 1
max_val = final_ratings.get(max_key)

for j in range(len(final_ratings)):
    if final_ratings.get(j+1) > max_val:
        max_key = j+1
        max_val = final_ratings.get(max_key)

print(str(max_key) + " " + str(max_val))
