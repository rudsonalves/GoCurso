/*
 Copyright ...
*/

/*
Pacote statistics possui funções de estatísica.
Neste pacote existem as funções:

  - Min;
  - Max;
  - ...

E a documentação continua...

# Isto é um título

...

	func Max(e ...float64) float64 {
		max := e[0]
		for i := 1; i < len(e); i++ {
			if e[i] > max {
				max = e[i]
			}
		}

		return max
	}

Lista numérica:
 1. elemento 1
 2. elemento 2
*/
package statistics

// Max retorna o máximo dos elementos passados.
func Max(e ...float64) float64 {
	max := e[0]
	for i := 1; i < len(e); i++ {
		// Este comentário está dentro do Max
		if e[i] > max {
			max = e[i]
		}
	}

	return max
}

func Min(e ...float64) float64 {
	min := e[0]
	for i := 1; i < len(e); i++ {
		if e[i] < min {
			min = e[i]
		}
	}

	return min
}

func MidRange(e ...float64) float64 {
	return (Max(e...) + Min(e...)) / 2.
}

func ArithmeticMean(e ...float64) float64 {
	sum := 0.
	for _, v := range e {
		sum += v
	}

	return sum / float64(len(e))
}

func zero() float64 {
	return 0.
}
