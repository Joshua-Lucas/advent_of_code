use clap::Parser;

#[derive(Parser, Debug)]
#[clap()]
pub struct Opts {
    #[clap(short = 'y', long = "year")]
    pub year: usize,

    #[clap(short = 'd', long = "day")]
    pub day: usize,
}
