<template>
  <div class="flex flex-column sm:flex-row gap-4">
    <ProgressSpinner v-if="!errorMessage" />

    <SectionHeadline :title="t('otp-auth-view.title')">
      <template #description>
        <p v-if="!errorMessage">{{ t('otp-auth-view.description') }}</p>

        <Message severity="error" v-if="errorMessage" :closable="false">
          {{ t('otp-auth-view.error', { message: errorMessage }) }}
        </Message>
      </template>
    </SectionHeadline>
  </div>
</template>

<script setup lang="ts">
import SectionHeadline from '@/components/structure/SectionHeadline.vue'
import { useI18n } from 'vue-i18n'
import ProgressSpinner from 'primevue/progressspinner'
import Message from 'primevue/message'
import { onMounted, ref } from 'vue'
import { fetchApi } from '@/plugins/api'
import { useRoute, useRouter } from 'vue-router'

const { t } = useI18n()

const errorMessage = ref<string | null>(null)

const route = useRoute()
const router = useRouter()

onMounted(async () => {
  try {
    // We only expect a 200 response, auth itself is handled by httpOnly cookies
    await fetchApi<{}>('/api/auth/otp', {
      method: 'POST',
      body: JSON.stringify({
        token: route.query.token
      })
    })

    await router.replace({ name: 'app' })
  } catch (e: any) {
    errorMessage.value = e.message
  }
})
</script>
