<template>
  <FormGrid @submit="onFormSubmit" class="mb-3">
    <FormRow>
      <PasswordInput
        :disabled="isSubmitting"
        name="password"
        :label="t('user-avatar.form.master-password.title')"
        :help="t('user-avatar.form.master-password.help')"
      />
    </FormRow>

    <SubmitButton :label="t('user-avatar.login')" :loading="isSubmitting" v-model="isSaved" />
  </FormGrid>
</template>

<script setup lang="ts">
import FormRow from '@/components/structure/form/FormRow.vue'
import FormGrid from '@/components/structure/form/FormGrid.vue'
import SubmitButton from '@/components/structure/form/SubmitButton.vue'
import { useForm } from 'vee-validate'
import { object, string } from 'yup'
import { useI18n } from 'vue-i18n'
import { ref } from 'vue'
import { ApiError, fetchApi } from '@/plugins/api'
import { useSessionStore } from '@/stores/session'
import PasswordInput from '@/components/structure/form/PasswordInput.vue'

const { t } = useI18n()
const { hello } = useSessionStore()

const isSaved = ref<boolean>(false)

const { handleSubmit, isSubmitting, setFieldError } = useForm<{ password: string }>({
  initialValues: {
    password: ''
  },
  validationSchema: object().shape({
    password: string().required().label(t('user-avatar.form.master-password.title'))
  })
})

const onFormSubmit = handleSubmit(async (v) => {
  try {
    await fetchApi<{ password: string }>('/api/auth/password', {
      method: 'POST',
      body: JSON.stringify(v)
    })

    await hello()
  } catch (e) {
    if (e instanceof ApiError) {
      switch (e.statusCode) {
        case 403:
          setFieldError('password', t('user-avatar.errors.password-invalid'))
          return
        case 405:
          setFieldError('password', t('user-avatar.errors.master-password-not-set'))
          return
        default:
          setFieldError('password', t('user-avatar.errors.unknown'))
          return
      }
    }
  }
})
</script>
