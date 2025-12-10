def one():
    count = 0
    with open("input.txt", "r") as file:
        word_search = [line.strip() for line in file.readlines()]

    length = len(word_search)
    directions = [(0, -1), (0, 1), (-1, 0), (1, 0), (-1, -1), (1, 1), (-1, 1), (1, -1)]

    for i, x in enumerate(word_search):
        for j, y in enumerate(x):
            if y == "X":
                print("Found X at index", i, j)
                for di, dj in directions:
                    if 0 <= i + 3 * di < length and 0 <= j + 3 * dj < length and \
                            word_search[i + di][j + dj] == "M" and \
                            word_search[i + 2 * di][j + 2 * dj] == "A" and \
                            word_search[i + 3 * di][j + 3 * dj] == "S":
                        print("Found XMAS")
                        count += 1
    print("found", count, "christmases", "meow")


one()