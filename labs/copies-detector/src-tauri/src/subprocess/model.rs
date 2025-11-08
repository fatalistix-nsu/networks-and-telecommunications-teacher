use std::fmt::{Display, Formatter};

#[derive(Copy, Clone)]
pub struct Subprocess {
    pub pid: u32,
    pub http_port: u16,
}

impl Subprocess {
    pub fn new(pid: u32, http_port: u16) -> Self {
        Self { pid, http_port }
    }
}

pub enum Error {
    SpawnFailed(std::io::Error),
    StdoutReadFailed(std::io::Error),
    WrongHttpPortBytesAmount(usize),
    NotFound(u32),
}

impl Display for Error {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        match self {
            Error::SpawnFailed(err) => {
                f.write_str(format!("Failed to spawn subprocess: {}", err).as_str())
            }
            Error::StdoutReadFailed(err) => {
                f.write_str(format!("Failed to read from subprocess stdout: {}", err).as_str())
            }
            Error::WrongHttpPortBytesAmount(amount) => {
                f.write_str(format!("Expected 2 bytes, but got {}", amount).as_str())
            },
            Error::NotFound(pid) => {
                f.write_str(format!("Subprocess {} does not exists", pid).as_str())
            }
        }
    }
}
