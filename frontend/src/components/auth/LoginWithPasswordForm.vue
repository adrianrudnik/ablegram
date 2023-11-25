<template>
  <FormGrid @submit="onFormSubmit" class="mb-3">
    <FormRow>
      <!-- Password forms should have (optionally hidden) username fields for accessibility: (More info: https://goo.gl/9p2vKq) -->
      <!-- @see https://stackoverflow.com/a/77542473/7527858 -->
      <input type="text" name="username" autocomplete="username" class="hidden" />

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
import { useRouter } from 'vue-router'

const { t } = useI18n()
const router = useRouter()
const { reconsider } = useSessionStore()

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

    await reconsider()

    await router.push({ name: 'app' })
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
