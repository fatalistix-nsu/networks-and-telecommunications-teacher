use crate::subprocess::model::{Error, Subprocess};
use crate::subprocess::store::Store;

pub struct Service {
    store: Store,
}

impl Service {
    pub fn new(store: Store) -> Self {
        Self { store }
    }

    pub fn run(&self) -> Result<Subprocess, Error> {
        self.store.run_child()
    }

    pub fn end(&self, pid: u32) {
        self.store.end(pid)
    }

    pub fn kill(&self, pid: u32) {
        self.store.kill(pid)
    }

    pub fn get_all(&self) -> Vec<Subprocess> {
        self.store.get_all()
    }
}
