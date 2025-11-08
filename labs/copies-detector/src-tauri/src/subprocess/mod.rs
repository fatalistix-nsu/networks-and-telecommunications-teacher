use crate::subprocess::service::Service;
use crate::subprocess::store::Store;
use tauri::Wry;

pub mod command;
pub mod model;
pub mod service;
pub mod store;

pub trait SubprocessTauriBuilder {
    fn register_subprocess_services(self) -> tauri::Builder<Wry>;
}

impl SubprocessTauriBuilder for tauri::Builder<Wry> {
    fn register_subprocess_services(self) -> tauri::Builder<Wry> {
        let store = Store::new();
        let service = Service::new(store);

        self.manage(service)
    }
}
