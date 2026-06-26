"Resolver el siguiente problema utilizando un algoritmo voraz: En un tablero de ajedrez (de
tamaño n x n) partimos de una casilla inicial (x,y). Tenemos una ficha de un caballo, que puede
realizar los mismos movimientos que en el ajedrez. El objetivo es, partiendo de la posición inicial,
visitar todas las casillas del tablero, sin repetir ninguna."

Conjunto candidatos: los movimientos disponibles del caballo (reglas de ajedrez)
Criterio seleccion: el movimiento que menor cantidad de movimientos proximos disponibles tiene
Factibilidad:el movimiento tiene que ser dentro del tablero , ademas de no poder ser un movimiento a una casilla ya visitada
Funcion objetivo: no maxmimza ni minimiza sino se busca encontrar la posibilidad de recorrer el tablero
funcion condicion: cuando todas las casillas fueron visitadas sin repetirse ninguna


//El ejemplo de catedra hace uso aparente de variables globales que a la vez son locales y no tiene sentido

//cuando eso pase lo pongo comentario
//Variables globales en el codifo de mrd:tablero y todas las cordenadas

funcion caballo_de_mrd (tablero:tablero , x,y,n:entero):booleano
	Ambiente
		i:entero
	Iniciar_tablero(tablero,n)//Pongo las casillas a 0 (0=no visitado)
	Para i=1 hasta (n*n) hacer
		tablero(x,y):=i
		Si Buscar_Mov(tablero,x,y,n)=falso y i<(n*n)-1 entonces  
			Retornar falso
		Fin Si
	Fin para
Fin funcion

funcion Buscar_Mov (tablero:tablero , x,y,n:entero):booleano //criterio de seleccion
	Ambiente
		i:entero
		menor:entero
		aux:entero
	menor:=9
	Para i:=1 hasta 8 hacer
	 Si Puedo_saltar(tablero ,x ,y,n ,i ) =true entonces
	 	aux:=movimiento_proximos(tablero,x,y,n,i)
		Si aux<menor entonces
			menor=aux
			prox_x=nueva_x // global
			prox_y=nueva_y //

	

	
