use crate::subprocess::model::Error::{SpawnFailed, StdoutReadFailed, WrongHttpPortBytesAmount};
use crate::subprocess::model::{Error, Subprocess};
use std::collections::HashMap;
use std::io::{Read, Write};
use std::path::PathBuf;
use std::process::{Child, Command, Stdio};
use std::sync::Mutex;

const SUBPROCESS_LOCATION: &str = "../copies-detector-backend/build/copies-detector-backend";

struct StoredSubprocess {
    child: Child,
    http_port: u16,
}

impl StoredSubprocess {
    pub fn new(child: Child, http_port: u16) -> Self {
        Self { child, http_port }
    }

    pub fn to_model(&self) -> Subprocess {
        Subprocess {
            pid: self.child.id(),
            http_port: self.http_port,
        }
    }
}

pub struct Store {
    lock: Mutex<HashMap<u32, Box<StoredSubprocess>>>,
}

impl Store {
    pub fn new() -> Self {
        Self {
            lock: Mutex::new(HashMap::new()),
        }
    }

    pub fn run_child(&self) -> Result<Subprocess, Error> {
        let subprocess_path = PathBuf::from(SUBPROCESS_LOCATION);

        log::info!("running subprocess {}", subprocess_path.display());

        let mut child = Command::new(subprocess_path)
            .stdin(Stdio::piped())
            .stdout(Stdio::piped())
            .spawn()
            .map_err(SpawnFailed)?;

        log::info!("spawned subprocess {}", child.id());

        let mut out = child.stdout.take().expect("no piped stdout from child");

        let mut buffer = Vec::new();
        out.read_to_end(&mut buffer).map_err(StdoutReadFailed)?;

        if buffer.len() != 2 {
            return Err(WrongHttpPortBytesAmount(buffer.len()));
        }

        let pid = child.id();
        let http_port = u16::from_be_bytes([buffer[0], buffer[1]]);

        log::info!("subprocess {} has http port {}", pid, http_port);

        let stored_subprocess = Box::new(StoredSubprocess::new(child, http_port));

        let model_subprocess = stored_subprocess.to_model();

        //TODO: think about potential previous stored value (unreal case)
        self.lock.lock().unwrap().insert(pid, stored_subprocess);

        Ok(model_subprocess)
    }

    pub fn end(&self, id: u32) {
        log::info!("ending subprocess {}", id);

        let stored_value_opt = self.lock.lock().unwrap().remove(&id);

        if stored_value_opt.is_some() {
            let stored_value = stored_value_opt.unwrap();
            let buffer: Vec<u8> = vec![1];
            let write_result = stored_value.child.stdin.unwrap().write(buffer.as_slice());
            if write_result.is_err() {
                let err = write_result.unwrap_err();
                log::warn!("failed to send {} bytes to subprocess {} as stop notification: {}", buffer.len(), id, err)
            } else {
                log::info!("sent {} bytes to subprocess {} as stop notification", buffer.len(), id);
            }
        } else {
            log::warn!("subprocess {} not found", id);
        }
    }

    pub fn kill(&self, id: u32) {
        log::info!("killing subprocess {}", id);

        let stored_value_opt = self.lock.lock().unwrap().remove(&id);

        if stored_value_opt.is_some() {
            let mut stored_value = stored_value_opt.unwrap();
            let kill_result = stored_value.child.kill();
            if kill_result.is_err() {
                let err = kill_result.unwrap_err();
                log::warn!("failed to kill subprocess {}: {}", id, err);
            } else {
                log::info!("subprocess {} was killed", id);
            }
        } else {
            log::warn!("subprocess {} not found", id);
        }
    }

    pub fn get(&self, pid: u32) -> Option<Subprocess> {
        self.lock
            .lock()
            .unwrap()
            .get(&pid)
            .map(move |v| v.to_model())
    }

    pub fn get_all(&self) -> Vec<Subprocess> {
        self.lock
            .lock()
            .unwrap()
            .values()
            .map(|v| v.to_model())
            .collect()
    }
}
