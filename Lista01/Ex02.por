programa
{
	funcao inicio()
	{
		inteiro t
		leia(t)

		para(inteiro i = 1; i <= t; i++)
		{
			real N, popular, geral, arquibancada, cadeiras
			leia(N, popular, geral, arquibancada, cadeiras)

			real renda
			renda = (N * popular / 100) * 1.00 +
			        (N * geral / 100) * 5.00 +
			        (N * arquibancada / 100) * 10.00 +
			        (N * cadeiras / 100) * 20.00

			escreva("A RENDA DO JOGO N. ", i, " E = ", renda, "\n")
		}
	}
}