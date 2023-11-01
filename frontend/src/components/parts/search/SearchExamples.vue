<template>
  <TabView class="SearchExamples border-2">
    <TabPanel :header="t('search-examples.basic.title')">
      <p>
        {{ t('search-examples.basic.rules.description') }}
      </p>

      <ul>
        <li>
          <i18n-t keypath="search-examples.basic.rules.list.1" tag="label">
            <code>[tag] free text [tag]</code>
          </i18n-t>
        </li>

        <li>
          <i18n-t keypath="search-examples.basic.rules.list.2.text" tag="label">
            <template v-slot:e1>
              <code>{{ t('search-examples.basic.rules.list.2.should') }}</code>
            </template>

            <template v-slot:e2>
              <code>{{ t('search-examples.basic.rules.list.2.must') }}</code>
            </template>
          </i18n-t>
        </li>

        <li>
          {{ t('search-examples.basic.rules.list.3') }}
        </li>
      </ul>

      <p>
        {{ t('search-examples.basic.rules.example.intro') }}
      </p>

      <i18n-t keypath="search-examples.basic.rules.example.explain.text" tag="p">
        <template v-slot:e1>
          <code>
            [{{ t('common.label.tag.s') }} {{ t('common.label.beats-per-minute.nc', { count: 110 }) }}]
            [{{ t('common.label.tag.s') }} {{ t('common.label.beats-per-minute.nc', { count: 120 }) }}] Taste
          </code>
        </template>

        <template v-slot:e2>
          <code>
            {{ t('common.label.beats-per-minute.nc', { count: 110 }) }}
          </code>
        </template>

        <template v-slot:e3>
          <code>
            {{ t('common.label.beats-per-minute.nc', { count: 120 }) }}
          </code>
        </template>

        <template v-slot:e4>
          <code>Taste</code>
        </template>

        <template v-slot:e5>
          <code>Taste of Marmalade</code>
        </template>
      </i18n-t>
    </TabPanel>

    <TabPanel :header="t('search-examples.basic.advanced.title')">
      <p>
        {{ t('search-examples.basic.advanced.examples.description') }}
      </p>

      <ul>
        <li>
          <i18n-t keypath="search-examples.basic.advanced.examples.list.1" tag="label">
            <template v-slot:e1>
              <code>+type:AudioTrack</code>
            </template>
            <template v-slot:e2>
              <code>MidiTrack</code>
            </template>
            <template v-slot:e3>
              <code>LiveSet</code>
            </template>
          </i18n-t>
        </li>

        <li>
          <i18n-t keypath="search-examples.basic.advanced.examples.list.2" tag="label">
            <code>-type:MidiTrack</code>
          </i18n-t>
        </li>

        <li>
          <i18n-t keypath="search-examples.basic.advanced.examples.list.3" tag="label">
            <code>type:MidiTrack +fx</code>
          </i18n-t>
        </li>

        <li>
          <i18n-t keypath="search-examples.basic.advanced.examples.list.4" tag="label">
            <code>+bpm:>=110 +bpm:&lt;=120</code>
          </i18n-t>
        </li>

        <li>
          <i18n-t keypath="search-examples.basic.advanced.examples.list.5" tag="label">
            <code>+filename: "Classic Disco-Flat vs Bumpy-118bpm.als"</code>
          </i18n-t>
        </li>

        <li>
          <i18n-t keypath="search-examples.basic.advanced.examples.list.6" tag="label">
            <template v-slot:e1>
              <code>-pathFolder:"/home/adrian/Downloads/ablegram/samples"</code>
            </template>

            <template v-slot:e2>
              <code>pathAbsolute</code>
            </template>
          </i18n-t>
        </li>

        <li>
          <i18n-t keypath="search-examples.basic.advanced.examples.list.7" tag="label">
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
      </ul>
    </TabPanel>
  </TabView>
</template>

<script setup lang="ts">
import TabView from 'primevue/tabview'
import TabPanel from 'primevue/tabpanel'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
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
      margin-right: 0.20em;
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
