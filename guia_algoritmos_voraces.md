# Guía de razonamiento — Algoritmos Voraces (Ávidos / Greedy)
### Practico 4 — Complejidad y Técnicas de Diseño de Algoritmos

Esta guía no te da "la solución" de memoria: te da **el método para que la
encuentres vos mismo**, en cualquier ejercicio voraz que te tomen (de este
práctico o de un parcial). La idea es que aprendas a hacerte siempre las
mismas 5 preguntas.

---

## 1. La idea central de "voraz"

Un algoritmo voraz construye la solución **paso a paso**, y en cada paso
toma la decisión que parece mejor **en ese momento**, sin volver atrás
nunca a reconsiderarla. Es lo opuesto a backtracking (que sí vuelve atrás)
y a programación dinámica (que evalúa todas las combinaciones posibles de
subproblemas).

La trampa — y lo que realmente te van a evaluar — es esto:

> **Tomar siempre la mejor decisión local NO siempre produce la mejor
> solución global.** Parte del trabajo en cada ejercicio es justificar si
> la voracidad funciona (da el óptimo) o si es solo una heurística rápida
> que puede fallar.

---

## 2. El marco de las 5 preguntas (úsalo en TODOS los ejercicios)

Para cualquier problema voraz, antes de escribir una línea de pseudocódigo,
respondé esto:

| Pregunta | Qué significa |
|---|---|
| **1. Candidatos (C)** | ¿Cuáles son los "ladrillos" con los que se construye la solución? (objetos, aristas, movimientos, ciudades…) |
| **2. Criterio de selección** | De todos los candidatos que quedan, ¿cuál es la regla para elegir "el mejor por ahora"? (esto es el corazón del algoritmo) |
| **3. Factibilidad** | Si agrego ese candidato, ¿la solución parcial sigue siendo válida? (no se rompe una restricción) |
| **4. Función objetivo** | ¿Qué estoy maximizando o minimizando? |
| **5. Condición de solución** | ¿Cuándo digo "ya terminé"? |

Con esas 5 respuestas, **cualquier** algoritmo voraz se escribe con el mismo
esqueleto:

```
Voraz(C):
    S := ∅                         // solución parcial, vacía al inicio
    Ordenar/organizar C según el criterio de selección (pregunta 2)
    Mientras C ≠ ∅ y no es_solucion(S):     // pregunta 5
        x := mejor_candidato(C)             // pregunta 2
        C := C - {x}
        Si es_factible(S ∪ {x}):            // pregunta 3
            S := S ∪ {x}
    Retornar S
```

Todo el "arte" de diseñar el algoritmo está en la pregunta 2: **encontrar
el criterio correcto**. Eso es lo que cambia de ejercicio a ejercicio.

---

## 3. Cómo razonar si la voracidad realmente funciona

No existe una receta universal para demostrarlo, pero hay un razonamiento
que sirve para sospechar/justificar en el 90% de los casos: el
**argumento de intercambio**.

> Supongamos que existe una solución óptima S\* que en su primer paso
> elige algo distinto de lo que elige tu regla voraz. Si podés demostrar
> que **siempre se puede intercambiar** ese elemento por el elegido
> voraz, sin empeorar el resultado, entonces existe una solución óptima
> que arranca igual que tu algoritmo voraz. Si además el problema que
> queda después de esa elección es "el mismo problema en chico"
> (subestructura óptima), por inducción tu algoritmo completo es óptimo.

Y al revés: si encontrás **un solo contraejemplo** donde la elección local
óptima te deja en una situación peor que otra elección, la voracidad
**no es óptima** — es una heurística. Eso no la invalida como algoritmo
(puede ser rápida y útil), pero tenés que decirlo en la respuesta.

Con esto en mente, vamos ejercicio por ejercicio.

---

## 4. Ejercicio por ejercicio

### Ejercicio 1 — Recorrido del caballo (tour completo desde (x,y))

**Releer el problema:** desde una casilla inicial, mover un caballo de
ajedrez visitando las n² casillas sin repetir ninguna.

Razonalo con las 5 preguntas:

1. **Candidatos:** en cada posición, los movimientos de caballo válidos
   desde ahí (hasta 8, pero menos cerca de los bordes).
2. **Criterio de selección (la parte interesante):** ¿cuál de los
   movimientos posibles conviene elegir *ahora*? Pensalo así: si te movés
   a una casilla con muchas salidas disponibles, no pasa nada, la podés
   visitar después igual. Pero si hay una casilla con **pocas salidas**,
   y no la visitás ahora, corre el riesgo de quedar "aislada"
   (rodeada de casillas ya visitadas) y nunca podrás llegar a ella. → La
   regla voraz (conocida como **heurística de Warnsdorff**) es: **moverte
   siempre a la casilla accesible con menor cantidad de salidas futuras
   disponibles** (la más "comprometida" primero).
3. **Factibilidad:** la casilla destino debe estar dentro del tablero y
   no haber sido visitada.
4. **Función objetivo:** no hay que maximizar/minimizar un valor, sino
   completar el recorrido.
5. **Condición de solución:** se visitaron las n² casillas.

Pseudocódigo (completalo vos, la estructura es esta):

```
Caballo_Voraz(tablero n x n, posición inicial (x,y)):
    visitado[x][y] := verdadero
    actual := (x,y)
    Para paso := 2 hasta n*n:
        candidatos := movimientos_validos(actual)   // hasta 8, sin salir y sin repetir
        Si candidatos = ∅:
            Retornar "no se pudo completar desde aquí"   // ¡importante decirlo!
        siguiente := el candidato con MENOR cantidad de movimientos válidos
                     futuros (grado de accesibilidad mínimo)
        visitado[siguiente] := verdadero
        actual := siguiente
    Retornar recorrido completo
```

**Lo que tenés que escribir en la respuesta (y es lo que más se evalúa):**
esta heurística funciona muy bien en la práctica, pero **no garantiza
siempre** terminar el recorrido — puede quedar "atrapada" sin salidas
antes de visitar todas las casillas. Por eso el mismo problema reaparece
en el Práctico 5 resuelto con **backtracking**: backtracking sí garantiza
encontrar una solución si existe (porque puede retroceder), a costa de
ser más lento.

---

### Ejercicio 2 — ¿Qué casillas iniciales permiten recorrer todo el tablero?

Es el mismo algoritmo del ejercicio 1, pero usado como **subrutina** para
explorar todos los puntos de partida posibles:

1. **Candidatos:** ahora los candidatos son las n² casillas como posibles
   *puntos de partida* (no movimientos).
2. **Criterio:** para cada casilla (x,y), correr el algoritmo voraz del
   ejercicio 1 empezando ahí.
3. **Factibilidad / condición de solución:** una casilla "sirve" como
   inicio si el algoritmo del ejercicio 1, lanzado desde ella, logra
   visitar las n² casillas (no exige que el camino sea cerrado, es decir,
   no exige volver al punto de partida — esa es la diferencia con el
   "tour cerrado" clásico).

```
Para cada casilla (x,y) del tablero:
    resultado := Caballo_Voraz(tablero, (x,y))
    Si resultado visitó las n*n casillas:
        marcar (x,y) como "casilla inicial válida"
```

**Nota de razonamiento:** como la heurística de Warnsdorff no es perfecta,
este algoritmo puede dar **falsos negativos** (decir que una casilla no
sirve cuando en realidad con otra estrategia sí serviría). Es importante
que lo aclares: el algoritmo voraz te da una *condición suficiente* pero
no necesaria.

---

### Ejercicio 3 — La Mochila (Knapsack)

Este es **el** ejercicio clásico para entender los límites de la
voracidad.

1. **Candidatos:** los objetos, cada uno con peso `p_i` y valor `v_i`.
2. **Criterio de selección:** ordenar los objetos por **razón
   valor/peso** `v_i / p_i`, de mayor a menor. Intuición: querés meter
   primero lo que te da más "ganancia por kilo".
3. **Factibilidad:** el peso acumulado + peso del objeto no debe superar
   la capacidad máxima.
4. **Función objetivo:** maximizar la suma de valores de lo metido.
5. **Condición de solución:** se recorrieron todos los objetos (o ya no
   entra ninguno más).

```
Mochila_Voraz(objetos[1..n], capacidad W):
    ordenar objetos por v_i/p_i descendente
    peso_actual := 0; valor_total := 0
    Para cada objeto i en ese orden:
        Si peso_actual + p_i <= W:
            incluir objeto i        // (o una fracción, ver abajo)
            peso_actual += p_i
            valor_total += v_i
    Retornar valor_total
```

**Ahora la parte que tenés que razonar (no memorizar):** ¿se puede partir
un objeto o no?

- **Mochila fraccionaria** (podés meter el 70% de un objeto): el
  algoritmo voraz por `v_i/p_i` es **siempre óptimo**. Se puede probar con
  argumento de intercambio: si dejaras afuera algo de mayor razón
  valor/peso para meter algo de menor razón, podrías cambiar un poco de
  peso de uno al otro y mejorar el resultado.
- **Mochila 0/1** (el objeto entra entero o no entra): el mismo algoritmo
  **NO garantiza el óptimo**. Contraejemplo clásico para que lo tengas
  como prueba/razonamiento:

  | Objeto | Peso | Valor | Razón v/p |
  |---|---|---|---|
  | A | 10 | 60 | 6 |
  | B | 20 | 100 | 5 |
  | C | 30 | 120 | 4 |

  Capacidad = 50. El voraz elige A (queda capacidad 40), luego B (queda
  capacidad 20), y C ya no entra → valor = **160**.
  Pero la combinación **B + C** pesa exactamente 50 y vale **220**, que es
  mejor. → El criterio "mejor razón local" llevó a una solución peor.

  **Conclusión que debés escribir:** en la mochila 0/1, el algoritmo
  voraz por razón valor/peso es solo una **aproximación rápida**, no la
  solución exacta (la solución exacta de la 0/1 se ve con programación
  dinámica o backtracking, que son justamente los próximos prácticos).

---

### Ejercicio 4 — "Pepe el Broker" (inversión en bolsa)

Antes de calcular nada, **traducí el problema al lenguaje de mochila**:

- El "peso" de invertir `x_i` euros en el valor `i` es simplemente `x_i`
  euros (el recurso limitado es el dinero `M`, igual unidad para todos).
- El "valor" obtenido al invertir `x_i` euros en `i` es `x_i · B_i · p_i`
  (dado en el enunciado).
- Cada valor bursátil tiene además un tope individual `C_i` (no podés
  invertir más de `C_i` en el mismo).

Ahora la pregunta clave: **¿la ganancia por euro invertido en `i` es
constante o decrece a medida que invierto más?** Mirá la fórmula: la
ganancia por euro es `B_i · p_i`, un número que **no depende de cuánto
invertís** (es lineal). Eso es exactamente lo que pasa en la mochila
**fraccionaria** (no en la 0/1) — y ahí el voraz sí es óptimo.

1. **Candidatos:** los n valores bursátiles.
2. **Criterio de selección:** ordenar por `B_i · p_i` descendente (mayor
   beneficio esperado por euro primero).
3. **Factibilidad:** dinero invertido acumulado ≤ M.
4. **Función objetivo:** maximizar `Σ x_i · B_i · p_i`.
5. **Condición de solución:** se agotó el dinero M o se recorrieron todos
   los valores.

```
Inversion_Voraz(valores[1..n] con (C_i, B_i, p_i), dinero M):
    ordenar valores por B_i*p_i descendente
    restante := M
    Para cada valor i en ese orden, mientras restante > 0:
        invertir_i := mínimo(C_i, restante)
        x_i := invertir_i
        restante := restante - invertir_i
    Retornar las x_i asignadas
```

**Por qué es óptimo (para justificarlo en la respuesta):** al ser la
ganancia por euro constante por cada valor (no decreciente), es
estructuralmente idéntico a la mochila fraccionaria, así que aplica el
mismo argumento de intercambio que la hace óptima.

---

### Ejercicios 5 y 6 — Árbol de Recubrimiento Mínimo (Prim y Kruskal)

Estos dos resuelven el **mismo problema** (encontrar el árbol que conecta
todos los vértices con el menor peso total posible) pero con criterios de
selección distintos. Es un buen ejemplo de que puede haber **más de un
criterio voraz correcto** para el mismo problema.

**Por qué funcionan ambos (la idea de fondo, vale para los dos):**
en cualquier corte del grafo (una forma de separar los vértices en dos
grupos), la arista de menor peso que cruza ese corte **siempre puede
formar parte de algún árbol de recubrimiento mínimo**. Esa es la
propiedad ("propiedad del corte") que sostiene a ambos algoritmos.

#### Prim

1. **Candidatos:** aristas que conectan un vértice ya incluido en el
   árbol con uno que todavía no está.
2. **Criterio de selección:** la de **menor peso** entre esas.
3. **Factibilidad:** automática — por construcción, nunca se forma un
   ciclo (uno de los dos extremos siempre es nuevo).
4. **Función objetivo:** minimizar la suma de pesos de las aristas
   elegidas.
5. **Condición de solución:** se incluyeron todos los vértices (n-1
   aristas para n vértices).

```
Prim(grafo G, vértice inicial r):
    incluidos := {r}
    T := ∅   // aristas del árbol
    Mientras incluidos ≠ V(G):
        e := arista de menor peso con un extremo en "incluidos"
             y el otro fuera
        T := T ∪ {e}
        incluidos := incluidos ∪ {vértice nuevo de e}
    Retornar T
```

(En la práctica se implementa con una cola de prioridad o, como en
Dijkstra, manteniendo para cada vértice fuera la menor arista que lo
conecta a "incluidos", y actualizándola cuando entra un vértice nuevo.)

#### Kruskal

1. **Candidatos:** **todas** las aristas del grafo (no solo las que
   tocan lo ya construido).
2. **Criterio de selección:** procesarlas en **orden creciente de
   peso** (se ordenan una sola vez, al principio).
3. **Factibilidad:** la arista se acepta solo si **no forma un ciclo**
   con las ya elegidas. Para chequear esto de forma eficiente se usa una
   estructura de **conjuntos disjuntos (Union-Find)**: cada componente
   conectada es un conjunto; una arista forma ciclo si sus dos extremos
   ya están en el mismo conjunto.
4. **Función objetivo:** igual que Prim, minimizar el peso total.
5. **Condición de solución:** se eligieron n-1 aristas, o se recorrieron
   todas.

```
Kruskal(grafo G):
    ordenar aristas de G por peso creciente
    T := ∅
    Para cada arista (u,v) en ese orden:
        Si encontrar(u) ≠ encontrar(v):     // no están en el mismo conjunto
            T := T ∪ {(u,v)}
            unir(u, v)                       // fusiona los dos conjuntos
    Retornar T
```

**Diferencia clave entre los dos (para que la sepas explicar, no solo
copiar):** Prim crece **un único árbol conectado** desde adentro hacia
afuera; Kruskal puede tener, en un momento dado, **varios fragmentos de
árbol desconectados entre sí** que se van fusionando. Por eso Prim no
necesita chequear ciclos (siempre agrega un vértice nuevo) y Kruskal sí.

---

### Ejercicio 7 — Red eléctrica de costo mínimo

No es un problema nuevo: es exactamente el problema de **árbol de
recubrimiento mínimo** disfrazado con otro enunciado. La traducción es:

- vértices = ciudades
- aristas = enlaces factibles entre ciudades
- peso de la arista = costo del enlace

Una vez que identificás esa equivalencia, aplicás **Prim o Kruskal**
(cualquiera de los dos algoritmos de los ejercicios 5 y 6, tal cual). Lo
que se evalúa acá es justamente que **reconozcas el patrón** detrás de un
enunciado distinto — esa es una habilidad clave para el parcial: muchos
problemas "nuevos" son en realidad MST, mochila, o caminos mínimos
disfrazados.

---

### Ejercicio 8 — Rutas de la agencia de turismo (Dijkstra modificado)

Esto es el problema de **camino más corto desde un origen único** (City
Aburrida) hacia todos los demás destinos. Es Dijkstra de manual, con dos
agregados pedidos en el enunciado: la **ruta** (no solo la distancia) y
el **combustible necesario**.

1. **Candidatos:** vértices aún no "cerrados" (procesados), cada uno con
   una distancia tentativa desde el origen.
2. **Criterio de selección:** el vértice no procesado con **menor**
   distancia tentativa acumulada — esa distancia ya no puede mejorar, así
   que se "cierra".
3. **Factibilidad / actualización (relajación):** al cerrar un vértice
   `u`, para cada vecino `v`: si `dist[u] + peso(u,v) < dist[v]`,
   actualizar `dist[v]` **y guardar `predecesor[v] := u`** (esto es lo
   que te permite reconstruir la ruta completa al final, no solo la
   distancia).
4. **Función objetivo:** minimizar la distancia (y de ahí derivar el
   combustible).
5. **Condición de solución:** se cerraron todos los vértices de la lista
   de destinos.

```
Dijkstra_Modificado(grafo G, origen = CityAburrida, destinos[1..k]):
    Para cada vértice v: dist[v] := infinito; predecesor[v] := nulo
    dist[origen] := 0
    no_procesados := todos los vértices
    Mientras no_procesados ≠ ∅:
        u := vértice de no_procesados con menor dist[u]
        sacar u de no_procesados
        Para cada vecino v de u:
            Si dist[u] + peso(u,v) < dist[v]:
                dist[v] := dist[u] + peso(u,v)
                predecesor[v] := u
    Para cada destino d en destinos:
        ruta := reconstruir desde d siguiendo predecesor[] hasta origen
        combustible[d] := dist[d] / rendimiento_del_vehículo   // o la
                           // suma de consumo por tramo, si cada arista
                           // tiene un consumo propio en vez de ser
                           // proporcional a la distancia
    Retornar dist[], predecesor[], combustible[] para cada destino
```

**La "modificación" que pide el enunciado** no es un criterio voraz
distinto (sigue siendo "elegir el vértice no cerrado de menor distancia
acumulada") — es agregar **bookkeeping extra** en cada relajación:
guardar de dónde vine (para la ruta) y, si el consumo de combustible no
es simplemente proporcional a la distancia, llevar una segunda etiqueta
de combustible que se actualiza junto con la distancia en cada
relajación.

---

## 5. Tabla resumen — para repasar de un vistazo

| Ej. | Problema | Criterio voraz | ¿Garantiza el óptimo? |
|---|---|---|---|
| 1 | Tour del caballo | Ir a la casilla con menos salidas futuras (Warnsdorff) | No — heurística, puede trabarse |
| 2 | Casillas iniciales válidas | Probar el algoritmo del 1 desde cada casilla | No — depende del 1 |
| 3 | Mochila | Mayor razón valor/peso primero | Sí si es fraccionaria; **no** si es 0/1 |
| 4 | Inversión (Pepe el Broker) | Mayor `B_i·p_i` primero | Sí (equivalente a mochila fraccionaria) |
| 5 | MST — Prim | Arista de menor peso que conecta el árbol con el resto | Sí |
| 6 | MST — Kruskal | Aristas en orden creciente de peso, sin formar ciclo | Sí |
| 7 | Red eléctrica mínima | Es MST (usar Prim o Kruskal) | Sí |
| 8 | Rutas + combustible | Vértice no cerrado de menor distancia acumulada | Sí (con pesos no negativos) |

---

## 6. Cómo usar esta guía para resolver el práctico

1. Para cada ejercicio, completá primero las **5 preguntas** de la
   sección 2 con tus propias palabras, antes de mirar el desarrollo de
   esta guía.
2. Escribí el pseudocódigo con el esqueleto genérico de la sección 2.
3. Preguntate siempre: ¿este criterio da el óptimo, o es una heurística?
   Si podés pensar un contraejemplo (como el de la mochila), inclúyelo —
   eso suele valer más puntos que el pseudocódigo en sí.
4. Si te trabás en cuál es "el criterio correcto", pensá qué pasaría si
   eligieras *cualquier otra cosa* primero: ¿te deja en peor posición que
   la elección voraz? Si la respuesta es "nunca", encontraste el
   criterio.

¡Éxitos con el práctico!
