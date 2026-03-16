programa
{
	funcao inicio()
	{
		real n1, n2, n3, media

		leia(n1, n2, n3)

		media = (n1 + n2 + n3) / 3

		escreva("MEDIA = ", media, "\n")

		se (media >= 6)
		{
			escreva("APROVADO\n")
		}
		senao
		{
			escreva("REPROVADO\n")
		}
	}
}