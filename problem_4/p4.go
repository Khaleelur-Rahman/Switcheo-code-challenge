package p4

// Solution A
// Time complexity : 0(n)
// Space complexity : 0(1)
func sum_to_n_a(n int) int {
    if n <= 0  {
        return 0
    }
	sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum
}

// Solution B
// Time complexity : 0(1)
// Space complexity : 0(1)
func sum_to_n_b(n int) int {
    if n <= 0  {
        return 0
    }

	return n * (n + 1) / 2;
}

// Solution C
// Time complexity : 0(1)
// Space complexity : 0(1)
func sum_to_n_c(n int) int {
    if n <= 0 {
        return 0
    }

    sum := 0

    start, end := 1, n

    for start <= end {
        sum += start + end
        start++
        end--
    }

    if n%2 == 1 {
        sum -= n / 2 + 1
    }

    return sum
}




