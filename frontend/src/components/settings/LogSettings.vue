<template>
  <Card class="surface-200">
    <template #title>{{ t('log-settings.title') }}</template>
    <template #content>
      <p>
        {{ t('log-settings.description') }}
      </p>

      <FormGrid @submit="onFormSubmit">
        <FormRow> </FormRow>

        <FormRow>
          <CheckboxInput
            name="enable_runtime_logfile"
            :label="t('log-settings.form.runtime-logfile.title')"
          />

          <i18n-t keypath="log-settings.form.runtime-logfile.location" tag="p">
            <template v-slot:path>
              <code>{{ values.runtime_logfile_path }}</code>
            </template>
          </i18n-t>
        </FormRow>

        <FormRow v-show="values.enable_runtime_logfile">
          <FormRadioGroup>
            <RadioInput
              name="level"
              :label="t('log-settings.form.level.info.title')"
              radio-value="info"
            />
            <RadioInput
              name="level"
              :label="t('log-settings.form.level.debug.title')"
              radio-value="debug"
            />
          </FormRadioGroup>

          <p v-if="values.level === 'info'">{{ t('log-settings.form.level.info.description') }}</p>
          <p v-if="values.level === 'debug'">
            {{ t('log-settings.form.level.debug.description') }}
          </p>
        </FormRow>

        <FormRow v-show="values.enable_runtime_logfile">
          <CheckboxInput name="enable_processed_logfile" label="Log scanned folders to log file" />

          <i18n-t keypath="log-settings.form.processed-logfile.location" tag="p">
            <template v-slot:path>
              <code>{{ values.process_logfile_path }}</code>
            </template>
          </i18n-t>
        </FormRow>

        <SubmitButton :label="t('button.save')" :loading="isSubmitting" v-model="isSaved" />
      </FormGrid>
    </template>
  </Card>
</template>

<script setup lang="ts">
import type { Config, LogConfig } from '@/stores/config'
import { useConfigStore } from '@/stores/config'
import FormRow from '@/components/structure/form/FormRow.vue'
import FormRadioGroup from '@/components/structure/form/FormRadioGroup.vue'
import RadioInput from '@/components/structure/form/RadioInput.vue'
import { useForm } from 'vee-validate'
import SubmitButton from '@/components/structure/form/SubmitButton.vue'
import CheckboxInput from '@/components/structure/form/CheckboxInput.vue'
import { useI18n } from 'vue-i18n'
import FormGrid from '@/components/structure/form/FormGrid.vue'
import Card from 'primevue/card'
import { fetchApi } from '@/plugins/api'
import { ref } from 'vue'

const { t } = useI18n()
const isSaved = ref<boolean>(false)

// Build the form
const { handleSubmit, isSubmitting, values } = useForm<LogConfig>({
  initialValues: useConfigStore().current.log
})

const onFormSubmit = handleSubmit(async (v) => {
  useConfigStore().current = await fetchApi<Config>('/api/config/log', {
    method: 'PUT',
    body: JSON.stringify({
      level: v.level,
      enable_runtime_logfile: v.enable_runtime_logfile,
      enable_processed_logfile: v.enable_processed_logfile
    })
  })

  isSaved.value = true
})
</script>
