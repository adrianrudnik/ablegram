<template>
  <TabView class="SearchExamples border-2 border-round">
    <TabPanel :header="t('search-examples.advanced.title')">
      <p>
        {{ t('search-examples.advanced.examples.description') }}
      </p>

      <ul>
        <li>
          <i18n-t keypath="search-examples.advanced.examples.list.1" tag="label">
            <template v-slot:e1>
              <code>+type:AbletonAudioTrack</code>
            </template>
            <template v-slot:e2>
              <code>AbletonMidiTrack</code>
            </template>
            <template v-slot:e3>
              <code>AbletonLiveSet</code>
            </template>
          </i18n-t>
        </li>

        <li>
          <i18n-t keypath="search-examples.advanced.examples.list.2" tag="label">
            <code>-type:AbletonMidiTrack</code>
          </i18n-t>
        </li>

        <li>
          <i18n-t keypath="search-examples.advanced.examples.list.3" tag="label">
            <code>type:AbletonMidiTrack +fx</code>
          </i18n-t>
        </li>

        <li>
          <i18n-t keypath="search-examples.advanced.examples.list.4" tag="label">
            <code>+bpm:>=110 +bpm:&lt;=120</code>
          </i18n-t>
        </li>

        <li>
          <i18n-t keypath="search-examples.advanced.examples.list.5" tag="label">
            <code>+filename: "Classic Disco-Flat vs Bumpy-118bpm.als"</code>
          </i18n-t>
        </li>

        <li>
          <i18n-t keypath="search-examples.advanced.examples.list.6" tag="label">
            <template v-slot:e1>
              <code>-pathFolder:"/home/adrian/Downloads/ablegram/samples"</code>
            </template>

            <template v-slot:e2>
              <code>pathAbsolute</code>
            </template>
          </i18n-t>
        </li>

        <li>
          <i18n-t keypath="search-examples.advanced.examples.list.7" tag="label">
            <template v-slot:e1>
              <code>annotation:MID</code>
            </template>

            <template v-slot:e2>
              <code>MID</code>
            </template>

            <template v-slot:e3>
              <code>MIDI</code>
            </template>
          </i18n-t>
        </li>

        <li>
          <i18n-t keypath="search-examples.advanced.examples.list.8" tag="label">
            <template v-slot:e1>
              <code>*</code>
            </template>

            <template v-slot:e2>
              <code>annotation:*</code>
            </template>
          </i18n-t>
        </li>
      </ul>

      <p>
        <a :href="apiTagUrl" target="_blank">
          {{ t('search-examples.advanced.tag-api-link') }}
        </a>
      </p>
    </TabPanel>
  </TabView>
</template>

<script setup lang="ts">
import TabView from 'primevue/tabview'
import TabPanel from 'primevue/tabpanel'
import { useI18n } from 'vue-i18n'
import { fetchApi } from '@/plugins/api'

const { t } = useI18n()

const apiTagUrl = import.meta.env.VITE_API_URL + '/api/tags?verbose=true'
</script>

<style lang="scss">
.SearchExamples {
  .p-tabview-panels {
    p,
    li {
      line-height: 1.8;
    }

    code {
      white-space: nowrap;
      font-size: 0.9em;
      margin-right: 0.2em;
      background-color: black;
      color: white;
      padding: 0.25rem;
      border-radius: 0.25rem;
    }

    li {
      padding-bottom: 6px;
    }
  }
}
</style>
