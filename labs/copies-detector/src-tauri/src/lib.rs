use crate::subprocess::SubprocessTauriBuilder;

pub mod subprocess;

#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() {
    tauri::Builder::default()
        .plugin(tauri_plugin_log::Builder::new().build())
        .plugin(tauri_plugin_opener::init())
        .register_subprocess_services()
        .invoke_handler(tauri::generate_handler![
                subprocess::command::subprocess_run,
                subprocess::command::subprocess_end,
                subprocess::command::subprocess_kill,
                subprocess::command::subprocess_get_all,
        ])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
