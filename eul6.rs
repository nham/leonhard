fn sumsqn(n: int) -> int {
    let mut sum = 0;
    let mut i = 1;

    while i <= n {
        sum += i*i;
        i += 1;
    }
    return sum;
}

fn sqsumn(n: int) -> int {
    let sum = n * (n+1) / 2;
    return sum * sum;
}

fn main() {
    let x = sumsqn(100);
    let y = sqsumn(100);
    println!("{} {} {}", x, y, y - x);
}
