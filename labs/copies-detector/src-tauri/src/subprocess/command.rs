use crate::subprocess::model::Subprocess;
use crate::subprocess::service::Service;
use log::log;
use tauri::State;

#[derive(serde::Serialize)]
pub struct SubprocessDto {
    pid: u32,
    http_port: u16,
}

impl SubprocessDto {
    pub fn new(pid: u32, http_port: u16) -> Self {
        Self { pid, http_port }
    }

    pub fn from(model: Subprocess) -> Self {
        Self {
            pid: model.pid,
            http_port: model.http_port,
        }
    }
}

pub enum Error {}

#[tauri::command]
pub fn subprocess_end(service: State<'_, Service>, pid: u32) {
    service.end(pid);
}

#[tauri::command]
pub fn subprocess_kill(service: State<'_, Service>, pid: u32) {
    service.kill(pid);
}

#[derive(serde::Serialize)]
pub struct SubprocessGetAllResponse {
    subprocesses: Vec<SubprocessDto>,
}

#[tauri::command]
pub fn subprocess_get_all(service: State<'_, Service>) -> SubprocessGetAllResponse {
    let subprocesses = service.get_all();

    log::info!("{} active subprocesses found", subprocesses.len());

    SubprocessGetAllResponse {
        subprocesses: service
            .get_all()
            .into_iter()
            .map(SubprocessDto::from)
            .collect(),
    }
}

#[tauri::command]
pub fn subprocess_run(service: State<'_, Service>) -> Result<SubprocessDto, String> {
    service
        .run()
        .map(SubprocessDto::from)
        .map_err(move |e| e.to_string())
}
