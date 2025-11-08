import { invoke } from "@tauri-apps/api/core";
import { Subprocess } from "../model/subprocess";

interface SubprocessDto {
	pid: number,
	http_port: number,
}

const toModel = (dto: SubprocessDto): Subprocess => ({
	pid: dto.pid,
	httpPort: dto.http_port,
})

interface GetAllResponse {
	subprocesses: SubprocessDto[],
}

export const getAll = async (): Promise<Subprocess[]> => {
	const dto = await invoke<GetAllResponse>("subprocess_get_all")
	return dto.subprocesses.map(toModel)
}

export const run = async () => {
	const dto = await invoke<SubprocessDto>("subprocess_run")
	
	return toModel(dto)
}

export const end = async (pid: number) => {
	await invoke("subprocess_end", { pid: pid} )
}

export const kill = async (pid: number) => {
	await invoke("subprocess_kill", { pid: pid })
}

