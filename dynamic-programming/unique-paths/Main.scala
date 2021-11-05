object Main {

  def main(args: Array[String]) = {
    println(uniquePaths(3, 7))
  }

  def uniquePaths(m: Int, n: Int): Int = {

    var pathMatrix = Array.ofDim[Int](m, n)

    for (i <- 0 until m) pathMatrix(i)(0) = 1
    for (i <- 0 until n) pathMatrix(0)(i) = 1

    for (i <-1 until m) {
      for(j <- 1 until n) {
        pathMatrix(i)(j) = pathMatrix(i-1)(j) + pathMatrix(i)(j-1)
      }
    }
    return pathMatrix(m-1)(n-1)
  }
}


