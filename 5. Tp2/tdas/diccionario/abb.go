package diccionario

import TDAPila "tdas/pila"

type nodoAb[K comparable, T any] struct {
	izq   *nodoAb[K, T]
	der   *nodoAb[K, T]
	clave K
	dato  T
}

type arbol[K comparable, T any] struct {
	raiz     *nodoAb[K, T]
	cmp      func(K, K) int
	cantidad int
}

type iteradorExt[K comparable, T any] struct {
	abb        *arbol[K, T]
	nodoActual *nodoAb[K, T]
	pila       TDAPila.Pila[*nodoAb[K, T]]
	desde      *K
	hasta      *K
}

func CrearABB[K comparable, T any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, T] {
	arbol := new(arbol[K, T])
	arbol.cmp = funcion_cmp
	return arbol
}

func crearNodo[K comparable, T any](clave K, valor T) *nodoAb[K, T] {
	nodo := new(nodoAb[K, T])
	nodo.clave = clave
	nodo.dato = valor
	return nodo
}

func (a *arbol[K, T]) buscarClave(clave K, nodo, anterior *nodoAb[K, T]) (*nodoAb[K, T], *nodoAb[K, T]) {
	if nodo == nil {
		return nil, anterior
	}
	comparacion := a.cmp(clave, nodo.clave)
	if comparacion == 0 {
		return nodo, anterior
	}
	anterior = nodo
	if comparacion > 0 {
		return a.buscarClave(clave, nodo.der, anterior)
	} else {
		return a.buscarClave(clave, nodo.izq, anterior)
	}
}

func (a *arbol[K, T]) esHoja(nodo *nodoAb[K, T]) bool {
	return nodo.izq == nil && nodo.der == nil
}

func (a *arbol[K, T]) Guardar(clave K, dato T) {
	nuevoNodo := crearNodo(clave, dato)
	if a.cantidad == 0 {
		a.raiz = nuevoNodo
		a.cantidad++
		return
	}

	actual, anterior := a.buscarClave(clave, a.raiz, nil)

	if actual == nil {
		comparacion := a.cmp(clave, anterior.clave)
		if comparacion > 0 {
			anterior.der = nuevoNodo
		} else if comparacion < 0 {
			anterior.izq = nuevoNodo
		}
		a.cantidad++
		return
	}

	actual.dato = dato
}

func (a *arbol[K, T]) Pertenece(clave K) bool {
	actual, _ := a.buscarClave(clave, a.raiz, nil)
	return claveExiste(actual, clave)
}

func claveExiste[K comparable, T any](actual *nodoAb[K, T], clave K) bool {
	if actual == nil {
		return false
	}
	return actual.clave == clave
}

func (a *arbol[K, T]) Obtener(clave K) T {
	actual, _ := a.buscarClave(clave, a.raiz, nil)
	if !claveExiste(actual, clave) {
		panic("La clave no pertenece al diccionario")
	}
	return actual.dato
}

func (a *arbol[K, T]) Cantidad() int {
	return a.cantidad
}

func (a *arbol[K, T]) esRaiz(clave K) bool {
	return a.cmp(clave, a.raiz.clave) == 0
}

func (a *arbol[K, T]) Borrar(clave K) T {
	actual, anterior := a.buscarClave(clave, a.raiz, nil)
	if !claveExiste(actual, clave) {
		panic("La clave no pertenece al diccionario")
	}
	a.cantidad--
	aux := actual.dato
	padre := &a.raiz

	if !a.esRaiz(actual.clave) {
		padre = a.yoSoyTuPadre(actual, anterior) // Que la fuerza apunte a vosotros
	}

	if a.esHoja(actual) {
		*padre = nil
		return aux
	}
	if actual.der == nil {
		*padre = actual.izq
		return aux
	}
	if actual.izq == nil {
		*padre = actual.der
		return aux
	}

	nodoAux, anterior := a.buscarAuxiliar(actual.der, actual) // Acá nos movemos una a la derecha
	padre = a.yoSoyTuPadre(nodoAux, anterior)
	actual.clave = nodoAux.clave
	actual.dato = nodoAux.dato

	if nodoAux.der != nil {
		*padre = nodoAux.der
	} else {
		*padre = nil
	}

	return aux
}

func (a *arbol[K, T]) yoSoyTuPadre(nodo, anterior *nodoAb[K, T]) **nodoAb[K, T] {
	comp := a.cmp(anterior.clave, nodo.clave)
	if comp < 0 {
		return &anterior.der
	} else { // No puede ser  == 0
		return &anterior.izq
	}
}

// Para esta funcion tomaremos la logica "me muevo uno a la derecha y luego todo a la izquierda"
func (a *arbol[K, T]) buscarAuxiliar(nodo, anterior *nodoAb[K, T]) (*nodoAb[K, T], *nodoAb[K, T]) {
	if nodo == nil {
		return nil, anterior
	}
	if nodo.izq == nil {
		return nodo, anterior
	}
	anterior = nodo
	return a.buscarAuxiliar(nodo.izq, anterior)
}

func (a *arbol[K, T]) Iterar(visitar func(clave K, dato T) bool) {
	a.raiz.iterarRango(nil, nil, visitar, a.cmp)
}

func (a *arbol[K, T]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato T) bool) {
	a.raiz.iterarRango(desde, hasta, visitar, a.cmp)
}

func (nodo *nodoAb[K, T]) iterarRango(desde *K, hasta *K, visitar func(clave K, dato T) bool, cmp func(K, K) int) bool {
	if nodo == nil {
		return true
	}
	continuar := true
	var compDesde, compHasta int
	if desde != nil {
		compDesde = cmp(nodo.clave, *desde)
	}
	if hasta != nil {
		compHasta = cmp(nodo.clave, *hasta)
	}
	if continuar && compDesde >= 0 {
		continuar = nodo.izq.iterarRango(desde, hasta, visitar, cmp)
	}
	if continuar && compDesde < 0 {
		continuar = nodo.der.iterarRango(desde, hasta, visitar, cmp)
	}
	if continuar && compDesde >= 0 && compHasta <= 0 {
		continuar = visitar(nodo.clave, nodo.dato)
	}
	if continuar && compDesde >= 0 && compHasta <= 0 {
		continuar = nodo.der.iterarRango(desde, hasta, visitar, cmp)
	}
	return continuar
}

func (a *arbol[K, T]) Iterador() IterDiccionario[K, T] {
	return a.IteradorRango(nil, nil)
}

func (a *arbol[K, T]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, T] {
	iterador := new(iteradorExt[K, T])
	iterador.abb = a
	iterador.pila = TDAPila.CrearPilaDinamica[*nodoAb[K, T]]()
	iterador.desde = desde
	iterador.hasta = hasta
	iterador.pilaRecursivaConRango(a.raiz)
	return iterador
}

func (i *iteradorExt[K, T]) pilaRecursivaConRango(nodo *nodoAb[K, T]) {
	if nodo == nil {
		return
	}
	var compDesde, compHasta int
	if i.desde != nil {
		compDesde = i.abb.cmp(nodo.clave, *i.desde)
	}
	if i.hasta != nil {
		compHasta = i.abb.cmp(nodo.clave, *i.hasta)
	}
	if compDesde >= 0 && compHasta <= 0 {
		i.pila.Apilar(nodo)
	}
	if compDesde < 0 {
		i.pilaRecursivaConRango(nodo.der)
	}
	i.pilaRecursivaConRango(nodo.izq)
}

func (i *iteradorExt[K, T]) HaySiguiente() bool {
	return !i.pila.EstaVacia()
}

func (i *iteradorExt[K, T]) Siguiente() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodoAnt := i.pila.Desapilar()
	if nodoAnt.der != nil {
		i.pilaRecursivaConRango(nodoAnt.der)
	}
}

func (i *iteradorExt[K, T]) VerActual() (K, T) {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return i.pila.VerTope().clave, i.pila.VerTope().dato
}
