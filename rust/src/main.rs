use clap::Parser;
use years::year_2021::run_year;

use crate::utils::cli::Opts;

pub mod utils;
pub mod years;

fn main() {
    // Get options on CLI command
    let opts = Opts::parse();
    let Opts { year, day } = opts;

    get_solution(year, day)
}

fn get_solution(year: usize, day: usize) {
    match year {
        2021 => run_year(day),
        _ => println!("There are no solutins for that year"),
    }
}
