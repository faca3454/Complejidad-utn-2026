// 1
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



dx: arreglo[1..8] de entero := [-2,-1,1,2,2,1,-1,-2] //movimientos del caballo en x
dy: arreglo[1..8] de entero := [1,2,2,1,-1,-2,-2,-1] //moviemientos del caballo en y

//Ej: con x+dx[1] y+dx[1] seria el movimiento uno en la posicion xy es decir
// x se mueve a la izquiera e y va hacia arriba (el movimiento en l del caballo

VAR globales: xg, yg //Sino imposible actualizar tablero jaja

funcion caballo_de_mrd (tablero:tablero ; x,y,n:entero):booleano
	Ambiente
		i:entero
	Iniciar_tablero(tablero,n)//Pongo las casillas a 0 (0=no visitado)
	xg:=x
	yg:=y
	Para i=1 hasta (n*n) hacer
		tablero(xg,yg):=i
		Si Buscar_Mov(tablero,xg,yg,n)=falso y i<(n*n)-1 entonces  //cumplio el objetivo?
			Retornar falso
		Fin Si
	Fin para
Fin funcion

funcion Buscar_Mov (tablero:tablero ; x,y,n:entero):booleano //criterio de seleccion
	Ambiente
		i:entero
		menor:entero
		aux:entero
		nx,ny:entero
	menor:=9
	Para i:=1 hasta 8 hacer//candidatos
	 nx:=x+dx[i]
	 ny:=y+dy[i]
	 Si Puedo_saltar(tablero ,nx ,ny,n) =true entonces
		 aux:=movimiento_proximos(tablero,nx,ny,n,i)
		 Si aux<menor and aux>0 entonces //criterio de seleccion
				menor=aux
				prox_x=nx 
				prox_y=ny 
		 Fin Si
	 Fin Si
	Fin Para

	xg:=prox_x
	yg:=prox_y 
fin funcion

funcion Puedo_saltar (tablero:tablero ; x,y,n:entero):booleano
	Puedo_saltar= x,y<n y x,y>1 y tablero(x,y)=0 //resultado booleano
fin funcion

funcion movimiento_proximos(tablero:tablero ; x,y,n:enteros)
	Ambiente
		i:entero
		aux,nx,ny:entero
	aux=0
	Para i=1 hasta 8 hacer
		nx:=x+dx[i]
	 	ny:=y+dy[i]
		Si Puedo_saltar(tablero, nx,ny,n)=true entonces
			aux:=aux+1
		Fin
	Fin Para
	Retornar aux
Fin funcion
	

//3	
"El problema consiste en, teniendo unos objetos que se desean meter en una mochila, siendo que la mochila solo puede llevar un peso máximo y si queremos meter todos los objetos se sobrepasa el peso máximo de la mochila, seleccionar los objetos que maximicen la ganancia del vendedor. Para ello cada objeto tiene un peso y un valor material asociado."


Conjunto candidatos: arreglo de objetos a meter a la mochila
Criterio seleccion: objeto cuyo valor respecto a su peso sea mayor
Factibilidad:el objeto debe no romper la capacidad maxima de la moxila
Funcion objetivo: maximar el valor dentro de la mochila
funcion condicion: el peso de la mochila esta al maximo o no entra ningun objeto mas

funcion mochila (objetos,pesos,valores:arreglo(1..n); n,capacidad):entero
	Ambiente
		peso_total,i,aux,valor_total,mayor,cual:entero
	valor_total:=0
	Para i=1 hasta n hacer
		mayor=0
		Para i=1 hasta n hacer
			aux:= valores[i]/peso[i]
			Si aux>mayor entonces
				mayor=aux
				cual:=i
			Fin si
		Fin Para
		Si peso_total+peso[cual]<capacidad entonces
			valor_total:=valor_total+aux
			valores[cual]=0
		Sino
			peso[cual]-capacidad:=

			//lara y facu se aman

	
	