<template>
  <Card class="surface-200">
    <template #title>{{ t('behavior-settings.title') }}</template>
    <template #content>
      <p>
        {{ t('behavior-settings.description') }}
      </p>

      <FormGrid @submit="onFormSubmit">
        <FormRow>
          <CheckboxInput
            name="open_browser_on_start"
            :label="t('behavior-settings.form.open_browser_on_start.title')"
          />

          <p>{{ t('behavior-settings.form.open_browser_on_start.description') }}</p>
        </FormRow>

        <FormRow>
          <CheckboxInput
            name="autostart_webservice"
            :label="t('behavior-settings.form.autostart_webservice.title')"
          />

          <p>{{ t('behavior-settings.form.autostart_webservice.description') }}</p>
        </FormRow>

        <FormRow>
          <CheckboxInput
            name="show_service_gui"
            :label="t('behavior-settings.form.show_service_gui.title')"
          />

          <p>{{ t('behavior-settings.form.show_service_gui.description') }}</p>
        </FormRow>

        <SubmitButton :label="t('button.save')" :loading="isSubmitting" v-model="isSaved" />
      </FormGrid>
    </template>
  </Card>
</template>

<script setup lang="ts">
import type { BehaviorConfig, Config, LogConfig } from '@/stores/config'
import { useConfigStore } from '@/stores/config'
import FormRow from '@/components/structure/form/FormRow.vue'
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
const { handleSubmit, isSubmitting, values } = useForm<BehaviorConfig>({
  initialValues: useConfigStore().current.behaviour
})

const onFormSubmit = handleSubmit(async (v) => {
  useConfigStore().current = await fetchApi<Config>('/api/config/behaviour', {
    method: 'PUT',
    body: JSON.stringify({
      autostart_webservice: v.autostart_webservice,
      open_browser_on_start: v.open_browser_on_start,
      show_service_gui: v.show_service_gui
    })
  })

  isSaved.value = true
})
</script>
