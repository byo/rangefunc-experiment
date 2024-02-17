fn main() {
    let before = Some("body");
    let after = before
        .map(|x| x.to_uppercase())
        .map(|x| x + "!")
        .unwrap_or("default".to_string());
    
    println!("{after}")
}
