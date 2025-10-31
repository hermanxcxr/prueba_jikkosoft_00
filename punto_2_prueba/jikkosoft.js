// Se ingresan dos parámetros:
var intFinal = 5 // número final
var listInt  = [2, 0, 2, 1, 3, 4, 5] // lista de números enteros

// Se usa el complemento y map para no hacer un doble ciclo for
function sumCoincidentPairs(intFinal, listInt) {
    if (!Array.isArray(listInt)) throw new TypeError('listInt debe ser un array');
    if (Array.isArray(listInt)  && listInt.length < 2) throw new TypeError('listInt debe tener al menos 2 elementos');
    if (typeof intFinal !== 'number') throw new TypeError('intFinal debe ser número');

    let Indexes = new Map()
    let selectedPairs = []

    for (let i = 0; i < listInt.length; i++) {
        const num = listInt[i]
        const complement = intFinal - num

        if (Indexes.has(complement)) {
            selectedPairs.push([Indexes.get(complement), i])
        }
        else {
            Indexes.set(num, i)
        }       
        
    }    
    return selectedPairs
}

var indexesCollection = sumCoincidentPairs(intFinal,listInt);
console.log(indexesCollection)