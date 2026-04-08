fn main() {
    println!("Rust service running");
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_sum() {
        assert_eq!(2 + 2, 4);
    }
}