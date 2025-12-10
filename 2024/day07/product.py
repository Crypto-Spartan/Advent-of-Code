def product(*iterables, repeat=1):
    # product('ABCD', 'xy') → Ax Ay Bx By Cx Cy Dx Dy
    # product(range(2), repeat=3) → 000 001 010 011 100 101 110 111

    if repeat < 0:
        raise ValueError('repeat argument cannot be negative')
    pools = [tuple(pool) for pool in iterables] * repeat
    print("pools:", pools)

    result = [[]]
    for pool in pools:
        print(f"{result=}")
        # result = [x+[y] for x in result for y in pool]
        temp = []
        for x in result:
            for y in pool:
                temp.append(x+[y])
        result = temp
        print(f"{result=}")
        print(f"----------------------")

    for prod in result:
        yield tuple(prod)

prod = tuple(product(range(3), repeat=1))
print(f"{len(prod)=}")

prod = tuple(product(range(3), repeat=2))
print(f"{len(prod)=}")

prod = tuple(product(range(3), repeat=3))
print(f"{len(prod)=}")

prod = tuple(product(range(3), repeat=4))
print(f"{len(prod)=}")

prod = tuple(product(range(3), repeat=5))
print(f"{len(prod)=}")
# product(range(3), repeat=4)
