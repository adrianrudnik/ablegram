import { defineStore } from 'pinia'
import { ref } from 'vue'
import { fetchApi } from '@/plugins/api'

export interface Config {
  about: {
    version: string
    commit: string
    build: string
  }

  log: LogConfig

  behaviour: {
    autostart_webservice: boolean
    autostart_browser: boolean
    show_gui: boolean
  }

  collector: {
    worker_count: number
    worker_delay_in_milliseconds: number
    searchable_paths: string[]
    exclude_system_folders: boolean
  }

  parser: {
    worker_count: number
    worker_delay_in_milliseconds: number
  }

  indexer: {
    worker_delay_in_milliseconds: number
  }

  webservice: {
    try_ports: number[]
  }
}

export interface LogConfig {
  level: 'debug' | 'info'
  enable_runtime_logfile: boolean
  enable_processed_logfile: boolean
  readonly runtime_logfile_path: string
  readonly process_logfile_path: string
}

export const useConfigStore = defineStore('config', () => {
  // Initialize the current settings with a fallback config that will be replaced once we load the URL
  const current = ref<Config>({
    about: {
      version: '',
      commit: '',
      build: ''
    },
    behaviour: {
      autostart_browser: false,
      autostart_webservice: false,
      show_gui: true
    },
    collector: {
      worker_count: 4,
      worker_delay_in_milliseconds: 100,
      searchable_paths: [],
      exclude_system_folders: true
    },
    indexer: {
      worker_delay_in_milliseconds: 100
    },
    log: {
      level: 'info',
      enable_runtime_logfile: false,
      enable_processed_logfile: false,
      runtime_logfile_path: '',
      process_logfile_path: ''
    },
    parser: {
      worker_count: 4,
      worker_delay_in_milliseconds: 100
    },
    webservice: {
      try_ports: [8080, 8081, 8082, 8083, 8084, 8085, 8086]
    }
  })

  const load = async () => {
    current.value = await fetchApi<Config>('/config')
  }

  return {
    current,
    load
  }
})
