def ternary_full_adder(a, b, c):
    """
    Implementa um somador completo ternário balanceado.
    Entradas: a, b, c (valores no conjunto {-1, 0, 1})
    Saídas: soma (bit de soma), carry (bit de propagação)
    """
    # Verifica se as entradas são válidas
    if a not in [-1, 0, 1] or b not in [-1, 0, 1] or c not in [-1, 0, 1]:
        raise ValueError("As entradas devem ser valores no conjunto {-1, 0, 1}.")

    # Soma inicial
    soma_inicial = a + b + c

    # Calcula o bit de soma e o carry
    if soma_inicial > 1:
        soma = soma_inicial - 3  # Ajusta para o intervalo {-1, 0, 1}
        carry = 1
    elif soma_inicial < -1:
        soma = soma_inicial + 3  # Ajusta para o intervalo {-1, 0, 1}
        carry = -1
    else:
        soma = soma_inicial
        carry = 0

    return soma, carry


# Testando o somador completo ternário balanceado
if __name__ == "__main__":
    # Exemplos de teste
    trit = [-1,0,1]
       

    print("Entradas (a, b, c) -> Soma, Carry")
    for a in trit:
        for b in trit:
            for c in trit:
                soma, carry = ternary_full_adder(a, b, c)
                print(f"{a},{b},{c},{soma},{carry}")
