pub mod days;

pub(crate) fn run_year(day: usize) {
    match day {
        2 => days::day02::answer::run_solution(),
        _ => println!("There is no solution for this day: {}", day),
    }
}
