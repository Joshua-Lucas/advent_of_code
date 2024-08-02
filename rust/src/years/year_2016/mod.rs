pub mod days;

pub(crate) fn run_year(day: usize) {
    match day {
        1 => days::day_01::answer::run_solution(),
        _ => println!("There is no solution for this day: {}", day),
    }
}
