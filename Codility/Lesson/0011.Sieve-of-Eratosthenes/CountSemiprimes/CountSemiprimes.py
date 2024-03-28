# https://github.com/Anfany/Codility-Lessons-By-Python3/blob/master/L11_Sieve%20of%20Eratosthenes/11.2%20CountSemiprimes.md
def solution(N, P, Q):
    """
    返回由数组P、Q的元素组成的区间内，不大于N的半素数的个数, 时间复杂度O(N * log(log(N)) + M)
    :param N: 半素数的最大值
    :param P: 数组
    :param Q: 数组
    :return: 每次查询,得到的半素数的个数
    """
    #  半素数只有3或4个因子，并且不能是素数的立方，例如(1, 3, 9, 27)(1, 5, 25, 125)这种情况
    #  首先计算出不大于N的半素数列表，是半素数的为其值，不是的为0
    semi_prime = []
    k =0
    for i in range(1, N + 1):
        factor_count = 0
        sign = 0
        for j in range(1, int(i ** 0.5) + 1):
            if i % j == 0:
                factor_count += 1
                f = i / j
                if f != j:
                    if f == j ** 2:
                        sign = 1
                        semi_prime.append(0)
                        break
                    else:
                        factor_count += 1
            if factor_count > 4:
                sign = 1
                semi_prime.append(0)
                break
        if sign != 1:
            if factor_count >= 3:
                semi_prime.append(i)
            else:
                semi_prime.append(0)

    index_dict = {}  # 得出当前数值以及前面一共有几个半素数
    semi_dict = {}  # 如果是半素数，则添加到字典中
    count = 0
    for index, value in enumerate(semi_prime):
        if value != 0:
            count += 1
            index_dict[value] = count
            semi_dict[value] = 0
        else:
            index_dict[index + 1] = count
    # index_dict {1: 0, 2: 0, 3: 0, 4: 1, 5: 1, 6: 2, 7: 2, 8: 2, 9: 3, 10: 4, 11: 4, 12: 4, 13: 4, 14: 5, 15: 6, 16: 6, 17: 6, 18: 6, 19: 6, 20: 6, 21: 7, 22: 8, 23: 8, 24: 8, 25: 9, 26: 10}
    #semi_dict {4: 0, 6: 0, 9: 0, 10: 0, 14: 0, 15: 0, 21: 0, 22: 0, 25: 0, 26: 0}
    print("index_dict",index_dict)
    print("semi_dict",semi_dict)

    result_list = []  # 开始计算，在指定区间内有几个半素数
    for i, j in zip(P, Q):
        if i in semi_dict:
            result_list.append(index_dict[j] - index_dict[i] + 1)
        else:
            result_list.append(index_dict[j] - index_dict[i])

    return result_list


if __name__ == '__main__':
    solution(26,[1, 4, 16],[26, 10, 20])