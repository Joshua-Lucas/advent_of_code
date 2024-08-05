use std::{env, fs};

pub fn run_solution() {
    // Get the current working directory
    let current_dir = env::current_dir().expect("no current dir");

    // Construct the path to the nested file
    let file_path = current_dir
        .join("src/years/year_2016/days")
        .join("day_01")
        .join("input.txt");

    // Parse input instructions
    let input =
        fs::read_to_string(&file_path).expect("Can't read input file, to figure out the solution");

    // -- Part One -- //

    // We start facing North
    let starting_point = Position {
        facing: "N",
        latt: 0,
        long: 0,
    };

    // We process the directions to determine the final location.
    let final_location = input
        .trim()
        .split(',')
        .map(|dir| {
            let direction = dir.trim().split_at(1);
            Direction::new(direction.0, direction.1)
        })
        .fold(starting_point, |acc, x| -> Position {
            navigate_blocks(acc, x)
        });

    // We calculate the blocks away.
    let blocks_away = final_location.long + final_location.latt;

    println!("Part One Anser:{:?}", blocks_away);

    // -- Part Two -- //

    println!("Part Two Anser:{:?}", "")
}

#[derive(Debug, PartialEq)]
struct Position<'a> {
    facing: &'a str,
    latt: i32,
    long: i32,
}

#[derive(Debug)]
struct Direction {
    plot: char,
    steps: i32,
}

impl Direction {
    fn new(plot: &str, steps: &str) -> Self {
        let plot_char = plot.chars().next().unwrap_or('\0');
        let step = steps.parse::<i32>().unwrap_or(0);
        return Direction {
            plot: plot_char,
            steps: step,
        };
    }
}

// Returns the position after the directions are processed.
fn navigate_blocks(current_position: Position, direction: Direction) -> Position {
    // Directions come as right or left, so we find the current facing position
    // and progress off that data.
    match current_position.facing {
        "N" => {
            if direction.plot == 'L' {
                Position {
                    facing: "W",
                    latt: current_position.latt,
                    long: current_position.long - direction.steps,
                }
            } else {
                Position {
                    facing: "E",
                    latt: current_position.latt,
                    long: current_position.long + direction.steps,
                }
            }
        }

        "S" => {
            if direction.plot == 'L' {
                Position {
                    facing: "E",
                    latt: current_position.latt,
                    long: current_position.long + direction.steps,
                }
            } else {
                Position {
                    facing: "W",
                    latt: current_position.latt,
                    long: current_position.long - direction.steps,
                }
            }
        }

        "W" => {
            if direction.plot == 'L' {
                Position {
                    facing: "S",
                    latt: current_position.latt - direction.steps,
                    long: current_position.long,
                }
            } else {
                Position {
                    facing: "N",
                    latt: current_position.latt + direction.steps,
                    long: current_position.long,
                }
            }
        }

        "E" => {
            if direction.plot == 'L' {
                Position {
                    facing: "N",
                    latt: current_position.latt + direction.steps,
                    long: current_position.long,
                }
            } else {
                Position {
                    facing: "S",
                    latt: current_position.latt - direction.steps,
                    long: current_position.long,
                }
            }
        }
        _ => current_position,
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_navigate_blocks() {
        let test_cases = vec![
            (
                Position {
                    facing: "N",
                    latt: 0,
                    long: 0,
                },
                Direction {
                    plot: 'R',
                    steps: 2,
                },
                Position {
                    facing: "E",
                    latt: 0,
                    long: 2,
                },
            ),
            (
                Position {
                    facing: "N",
                    latt: 0,
                    long: 1,
                },
                Direction {
                    plot: 'L',
                    steps: 3,
                },
                Position {
                    facing: "W",
                    latt: 0,
                    long: -2,
                },
            ),
            (
                Position {
                    facing: "E",
                    latt: 0,
                    long: 0,
                },
                Direction {
                    plot: 'R',
                    steps: 2,
                },
                Position {
                    facing: "S",
                    latt: -2,
                    long: 0,
                },
            ),
        ];

        for (current, dir, expected) in test_cases.into_iter() {
            assert_eq!(navigate_blocks(current, dir), expected)
        }
    }
}
