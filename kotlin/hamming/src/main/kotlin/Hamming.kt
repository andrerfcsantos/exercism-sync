object Hamming {

    fun compute(leftStrand: String, rightStrand: String): Int {
        require (leftStrand.length == rightStrand.length){
            "left and right strands must be of equal length"
        }

        var dist = 0

        leftStrand.forEachIndexed { i, elem ->
            if (elem != rightStrand[i]){
                dist++
            }
        }

        return dist
    }

    fun compute2(leftStrand: String, rightStrand: String): Int {
        require (leftStrand.length == rightStrand.length){
            "left and right strands must be of equal length"
        }

        val size = leftStrand.length
        var dist = 0

        for (i in 0 until size) {
            if (leftStrand[i] != rightStrand[i]) dist++
        }

        return dist
    }
}
