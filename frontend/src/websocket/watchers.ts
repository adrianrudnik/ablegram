import { watchThrottled } from '@vueuse/core'
import { computed } from 'vue'
import { useStatStore } from '@/stores/stats'
import { hydrateTags } from '@/stores/tags'

export function startWebsocketWatchers() {
  // Watcher used to update tags on a regular basis while processing is still going on
  watchThrottled(
    computed(() => useStatStore().processRoutineCount),
    async (v) => {
      await hydrateTags()
    },
    { throttle: 1000 }
  )
}
