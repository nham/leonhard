pub fn is_prime(x: int) -> bool {
    if x == 1 {
        return false;
    } else if x == 2 {
        return true;
    }

    if x % 2 == 0 {
        return false;
    }

    let mut j = 3;
    while j*j <= x {
        if x % j == 0 {
            return false;
        }

        j += 2;
    }

    return true;
}

#[test]
fn test_is_prime() {
    assert!(!is_prime(1));
    assert!(is_prime(2));
    assert!(is_prime(3));
    assert!(!is_prime(4));
    assert!(is_prime(5));
    assert!(is_prime(-17));
    assert!(!is_prime(-18));
    assert!(!is_prime(9));
}
