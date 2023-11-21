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
  behaviour: BehaviorConfig
  collector: CollectorConfig

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

export interface BehaviorConfig {
  autostart_webservice: boolean
  open_browser_on_start: boolean
  show_service_gui: boolean
}

export interface CollectorConfig {
  targets: CollectorTargetConfig[]
}

export interface CollectorTargetConfig {
  id: string
  type: 'filesystem'
  uri: string
  parser_performance: 'low' | 'default' | 'high'
  parser_delay: number
  exclude_system_folders: boolean
  exclude_dot_folders: boolean
}

export function defaultCollectorTargetConfig(): CollectorTargetConfig {
  return {
    id: '',
    type: 'filesystem',
    uri: '',
    parser_performance: 'default',
    parser_delay: 0,
    exclude_system_folders: true,
    exclude_dot_folders: true
  }
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
      autostart_webservice: true,
      open_browser_on_start: true,
      show_service_gui: true
    },
    collector: {
      targets: []
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
    webservice: {
      try_ports: [8080, 8081, 8082, 8083, 8084, 8085, 8086]
    }
  })

  const load = async () => {
    try {
      current.value = await fetchApi<Config>('/api/config')
    } catch (e) {
      console.error('Failed to load config', e)
    }
  }

  return {
    current,
    load
  }
})
